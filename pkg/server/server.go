package server

import (
	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPServer struct {
	engine        *gin.Engine
	routerManager *RouterManager
	errorHandler  *ErrorHandler
	serverConfig  *ServerConfig
}

type ServerConfig struct {
	Address          string        `json:"address" yaml:"address"`
	ReadTimeout      time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout     time.Duration `json:"write_timeout" yaml:"write_timeout"`
	MaxHeaderBytes   int           `json:"max_header_bytes" yaml:"max_header_bytes"`
	Mode             string        `json:"mode" yaml:"mode"`
	LogOutput        bool          `json:"log_output" yaml:"log_output"`
	PprofEnabled     bool          `json:"pprof_enabled" yaml:"pprof_enabled"`
	PanicHandlerFunc gin.HandlerFunc
}

type RouterManager struct {
	engine *gin.Engine
}

type ErrorHandler struct {
	writer    ErrorHandlerWriter
	handlerFn gin.HandlerFunc
}

type ErrorHandlerWriter interface {
	Write([]byte) (int, error)
}

type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (int, error) {
	errStr := string(b)
	err := errors.New(errStr)
	return len(errStr), err
}

func NewHTTPServer(config *ServerConfig) *HTTPServer {
	gin.SetMode(config.Mode)
	if !config.LogOutput {
		gin.DefaultWriter = ioutil.Discard
	}

	engine := gin.New()
	engine.Use(gin.Logger(), config.PanicHandlerFunc)

	routerManager := &RouterManager{
		engine: engine,
	}

	return &HTTPServer{
		engine:        engine,
		routerManager: routerManager,
		errorHandler:  NewErrorHandler(&PanicExceptionRecord{}),
		serverConfig:  config,
	}
}

func (s *HTTPServer) Start() {
	svc := endless.NewServer(s.serverConfig.Address, s.engine)
	svc.ReadHeaderTimeout = s.serverConfig.ReadTimeout
	svc.WriteTimeout = s.serverConfig.WriteTimeout
	svc.MaxHeaderBytes = s.serverConfig.MaxHeaderBytes

	if s.serverConfig.PprofEnabled {
		pprof.Register(s.engine)
	}

	svc.ListenAndServe()
}

func (s *HTTPServer) AddRoute(httpMethod, route string, handler ...gin.HandlerFunc) {
	s.routerManager.AddRoute(httpMethod, route, handler...)
}
func (s *HTTPServer) AddRoutes(routes []Route) {
	s.routerManager.AddRoutes(routes)
}

func (s *HTTPServer) UseMiddleware(handlerFunc ...gin.HandlerFunc) {
	s.engine.Use(handlerFunc...)
}

func NewErrorHandler(writer ErrorHandlerWriter) *ErrorHandler {
	handlerFn := func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := errors.New(cast.ToString(err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": s.Error(),
				})
			}
		}()

		c.Next()
	}

	return &ErrorHandler{
		writer:    writer,
		handlerFn: handlerFn,
	}
}

func (e *ErrorHandler) GinHandlerFunc() gin.HandlerFunc {
	return e.handlerFn
}

func (r *RouterManager) AddRoute(httpMethod, route string, handler ...gin.HandlerFunc) {
	r.engine.Handle(httpMethod, route, handler...)
}

func (r *RouterManager) AddRoutes(routes []Route) {
	for _, route := range routes {
		r.engine.Handle(route.Method, route.Path, route.Handlers...)
	}
}

type Route struct {
	Method   string
	Path     string
	Handlers []gin.HandlerFunc
}
