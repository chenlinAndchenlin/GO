package global

import (
	"mxshop_api/uer-web/config"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans            ut.Translator
	UserServerConfig *config.ServerConfig = &config.ServerConfig{}
)
