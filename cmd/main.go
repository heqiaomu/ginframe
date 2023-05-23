/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Heqiaomu/ginframe/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/heqiaomu/gtools/gviper"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	cfgFile     string
	projectBase string
	userLicense string
	runMode     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GinFrame",
	Short: "Gin的框架",
	Long:  `一个基于Gin的Web框架`,
	Run: func(cmd *cobra.Command, args []string) {
		config := &server.ServerConfig{
			Address:        "0.0.0.0:" + cast.ToString(gviper.GetV().GetString("server.port")),
			ReadTimeout:    time.Second * 60,
			WriteTimeout:   time.Second * 60,
			MaxHeaderBytes: 1 << 20,
			Mode:           runMode,
			LogOutput:      true,
			PprofEnabled:   viper.GetBool("server.pprof"),
		}

		httpServer := server.NewHTTPServer(config)
		httpServer.AddRoutes(server.DefaultRouter()) // 设置默认路由

		httpServer.UseMiddleware(gin.Logger())
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/cfg.yaml", "config file (default is $HOME/.ginframe.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "pb", "b", "github.com/Heiqoamu/ginframe", "base project directory eg. github.com/Heiqoamu/")
	rootCmd.PersistentFlags().StringP("author", "a", "Heqiaomu", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().StringVarP(&runMode, "mode", "m", "release", "运行模式: release/debug 默认（release）")

	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("pb", rootCmd.PersistentFlags().Lookup("pb"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
}

// initConfig 初始化操作
func initConfig() {
	if cfgFile == "" {
		cfgFile = "config/cfg.yaml"
	}
	err := gviper.New(cfgFile)
	if err != nil {
		panic(err)
	}

}

func main() {
	Execute()
}
