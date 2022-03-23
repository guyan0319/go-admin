package common

import (
	"fmt"
	"github.com/smallnest/rpcx/log"
	"github.com/spf13/viper"
	"go-admin/conf"
)

var Conf conf.Config

func init() {
	v := viper.New()
	v.AddConfigPath("./etc/")
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error: config init  %s", err.Error())
	}
	if err := v.Unmarshal(&Conf); err != nil {
		log.Fatalf("error: config init  %s", err.Error())
	}
	//fmt.Println("in func(viper): ", Conf)
}
func Load() error {
	return nil
}
func ParseConfig(name string, s interface{}) error {
	v := viper.New()
	v.AddConfigPath("./etc/")
	v.SetConfigType("yaml")
	v.SetConfigName(name)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	fmt.Println("in func(viper): ", s)
	return nil
}
func Mysql() {

}
func Url() string {
	return "http://" + Conf.Host + ":" + Conf.Port
}
