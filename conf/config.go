package conf

import (
	"sync"
)

type Config struct {
	Language string
	Token string
	ShowSql bool
	ShowExecTime bool
}
var (
	Cfg  Config
	mutex   sync.Mutex
	declare sync.Once
)

func  Set(cfg Config) {
	mutex.Lock()
	Cfg.Language=setDefault(cfg.Language,"","cn")
	Cfg.Token=setDefault(cfg.Token,"","token")
	Cfg.ShowSql=cfg.ShowSql
	Cfg.ShowExecTime=cfg.ShowExecTime
	mutex.Unlock()
}
func setDefault( value,def ,defValue string) string {
	if value==def {
		return defValue
	}
	return value
}
