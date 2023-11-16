package config

type UsersrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Port        int           `mapstructure:"port"`
	UsersrvInfo UsersrvConfig `mapstructure:"user_srv""`
}
