package conf

import (
	"sync"
)

type Config struct {
	language string
}
var (
	Cfg  Config
	mutex   sync.Mutex
	declare sync.Once
)

func  Set(cfg Config) {
	mutex.Lock()
	Cfg.language=setDefault(cfg.language,"","cn")

	mutex.Unlock()
}
func setDefault( value,def ,defValue string) string {
	if value==def {
		return defValue
	}
	return value

}
