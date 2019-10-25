package conf

import (
	"sync"
)

type Config struct {
	Language string
}
var (
	Cfg  Config
	mutex   sync.Mutex
	declare sync.Once
)

func  Set(cfg Config) {
	mutex.Lock()
	Cfg.Language=setDefault(cfg.Language,"","cn")

	mutex.Unlock()
}
func setDefault( value,def ,defValue string) string {
	if value==def {
		return defValue
	}
	return value

}
