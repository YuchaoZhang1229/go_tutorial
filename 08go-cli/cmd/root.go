package cmd

import (
	"fmt" // 标准库，格式化I/O
	"os"  // 标准库，操作系统功能

	"github.com/spf13/cobra" // CLI框架
	"github.com/spf13/viper" // 配置管理库
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "myapp",                                                                           // 命令名称
	Short: "MyApp is a sample Go CLI application",                                            // 简短描述
	Long:  `A sample Go CLI application using Cobra and Viper for configuration management.`, // 详细描述
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./configs/config.yaml)")
	rootCmd.PersistentFlags().String("log-level", "info", "Log level (debug, info, warn, error)")

	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // 使用指定的配置文件
	} else {
		viper.AddConfigPath("./configs") // 添加配置搜索路径
		viper.AddConfigPath(".")         // 添加当前目录为搜索路径
		viper.SetConfigName("config")    // 设置配置文件名（不含扩展名）
	}

	viper.SetEnvPrefix("MYAPP") // 设置环境变量前缀 这样环境变量MYAPP_LOG_LEVEL将对应配置项log.level
	viper.AutomaticEnv()        // 自动绑定环境变量 自动转换为小写，_转换为.

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed()) // 打印实际使用的配置文件路径
	}
}
