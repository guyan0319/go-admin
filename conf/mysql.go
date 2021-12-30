package conf
type DbConfig struct {
	DriverName string
	Dsn string
	ShowSql bool
	ShowExecTime bool
	MaxIdle int
	MaxOpen int
}

var Db = map[string]DbConfig{
	"db1": {
		DriverName: "mysql",
<<<<<<< HEAD
		Dsn:        "root:Myyd@159@tcp(127.0.0.1:3306)/systemrdb?charset=utf8mb4&parseTime=true&loc=Local",
=======
		Dsn:        "root:123456@tcp(127.0.0.1:3306)/systemdb?charset=utf8mb4&parseTime=true&loc=Local",
>>>>>>> 94b83ae67378b8f339e274bfcb7699ae765af7df
		ShowSql:    true,
		ShowExecTime:    false,
		MaxIdle:    10,
		MaxOpen:    200,
	},
}

