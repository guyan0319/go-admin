package conf
type DbConfig struct {
	DriverName string
	Dsn string
	MaxIdle int
	MaxOpen int
}

var Db = map[string]DbConfig{
	"db1": {
		DriverName: "mysql",
		Dsn:        "root:123456@tcp(127.0.0.1:3306)/systemdb?charset=utf8mb4&parseTime=true&loc=Local",
		MaxIdle:    10,
		MaxOpen:    200,
	},
}

