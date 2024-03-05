/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/Heqiaomu/ginframe/internal/api"
	"os"
	"time"

	"github.com/Heqiaomu/ginframe/internal/config"
	"github.com/Heqiaomu/ginframe/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	projectBase string
	userLicense string
	runMode     string
)
var rootCmd = &cobra.Command{
	Use:   "luka",
	Short: "luka",
	Long:  `Luka的Go实现版本`,
	Run: func(cmd *cobra.Command, args []string) {
		config := &server.Config{
			Address:        fmt.Sprintf("0.0.0.0:%d", config.GetConfig().Server.Port),
			ReadTimeout:    time.Second * 60,
			WriteTimeout:   time.Second * 60,
			MaxHeaderBytes: 1 << 20,
			Mode:           runMode,
			LogOutput:      true,
			PprofEnabled:   config.GetConfig().Server.Pprof,
			PanicHandlerFunc: func(ctx *gin.Context) {
			},
			Prefix: config.GetConfig().Server.Prefix,
		}

		httpServer := server.NewHTTPServer(config)
		api.Router(httpServer.GetRouteGroup())
		httpServer.Start()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "conf/cfg.toml", "config file (default is $HOME/.ginframe.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "pb", "b", "github.com/Heiqoamu/ginframe", "base project directory eg. github.com/Heiqoamu/")
	rootCmd.PersistentFlags().StringP("author", "a", "SunYang", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().StringVarP(&runMode, "mode", "m", "release", "运行模式: release/debug 默认（release）")
}

// initConfig 初始化操作
func initConfig() {
	if cfgFile == "" {
		cfgFile = "conf/cfg.toml"
	}

	if _, err := config.InitConfig(cfgFile); err != nil {
		panic(err)
	}

}

func main() {
	Execute()
}
