package lang

import (
	"go-admin/conf"
)

var langMap = map[string]map[string]string{
"cn": cn,
"en": en,
}

func Get(value string) string{
	langKey := ""
	if conf.Cfg.Language=="" {
		langKey = "cn"
	}else{
		langKey = conf.Cfg.Language
	}
	if msg,ok :=langMap[langKey][value];ok {
		return msg
	}
	return value
}