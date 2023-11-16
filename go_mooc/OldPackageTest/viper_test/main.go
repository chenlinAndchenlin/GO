package main

//type ServerConfig struct {
//	ServerName string `mapstructure:"name"`
//	Port       int    `mapstructure:"port"`
//}
//
//func main() {
//	v := viper.New()
//	v.SetConfigFile("config.yaml")
//	err := v.ReadInConfig()
//	if err != nil {
//		panic(err)
//	}
//	serverConfig := ServerConfig{}
//	err = v.Unmarshal(&serverConfig)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(v.Get("name"))
//
//	fmt.Println(serverConfig)
//
//}
