package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "sync",
	Short: "root server.",
	Long:  "root server.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("cfgFile:", cfgFile)
}

func init() {
	cobra.OnInitialize(initConfig)
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&cfgFile, "config", "c", "", "config file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// 从主目录下搜索后缀名为 ".toml" 文件 (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("config_import")
	}
	viper.AutomaticEnv() // 读取匹配的环境变量
	viper.SetConfigType("toml")
	viper.SetEnvPrefix("EasySwap")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// 读取找到的配置文件
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}
}
