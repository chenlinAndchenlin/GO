package initialize

import (
	"fmt"
	"mxshop_api/uer-web/global"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func InitConfig() {
	debug := GetEnvInfo("MXSHOOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("uer-web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("uer-web/%s-dubug.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//serverConfig := config.ServerConfig{}
	err = v.Unmarshal(global.UserServerConfig)
	if err != nil {
		panic(err)
	}
	//fmt.Println(v.Get("name"))
	//
	//fmt.Println(global.UserServerConfig)
	zap.S().Infof("配置信息：%v", global.UserServerConfig)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		//fmt.Println("config file changed:",e.Name)
		zap.S().Infof("配置信息产生变化: %v", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.UserServerConfig)
		//fmt.Println(global.UserServerConfig)
		zap.S().Infof("配置信息：%v", global.UserServerConfig)
	})
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
