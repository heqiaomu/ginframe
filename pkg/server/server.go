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
	engine       *gin.Engine
	errorHandler *ErrorHandler
	cfg          *Config
}

type Config struct {
	Address          string        `json:"address" yaml:"address"`
	ReadTimeout      time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout     time.Duration `json:"write_timeout" yaml:"write_timeout"`
	MaxHeaderBytes   int           `json:"max_header_bytes" yaml:"max_header_bytes"`
	Mode             string        `json:"mode" yaml:"mode"`
	LogOutput        bool          `json:"log_output" yaml:"log_output"`
	PprofEnabled     bool          `json:"pprof_enabled" yaml:"pprof_enabled"`
	Prefix           string        `json:"prefix" yaml:"prefix"`
	PanicHandlerFunc gin.HandlerFunc
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

func NewHTTPServer(config *Config) *HTTPServer {
	gin.SetMode(config.Mode)
	if !config.LogOutput {
		gin.DefaultWriter = ioutil.Discard
	}

	engine := gin.New()
	engine.Use(gin.Logger(), config.PanicHandlerFunc)

	return &HTTPServer{
		engine:       engine,
		errorHandler: NewErrorHandler(&PanicExceptionRecord{}),
		cfg:          config,
	}
}

func (s *HTTPServer) Start() {
	svc := endless.NewServer(s.cfg.Address, s.engine)
	svc.ReadHeaderTimeout = s.cfg.ReadTimeout
	svc.WriteTimeout = s.cfg.WriteTimeout
	svc.MaxHeaderBytes = s.cfg.MaxHeaderBytes
	if s.cfg.PprofEnabled {
		pprof.Register(s.engine)
	}
	svc.ListenAndServe()
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

func (s *HTTPServer) GetEngine() *gin.Engine {
	return s.engine
}

func (s *HTTPServer) GetRouteGroup() *gin.RouterGroup {
	if s.cfg.Prefix != "" {
		return s.engine.Group(s.cfg.Prefix)
	}
	return s.engine.Group("/")
}
