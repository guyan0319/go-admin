package conf

var Db = map[string]map[string]string{
	"db1": {
		"driverName": "mysql",
		"dsn":        "root:123456@tcp(127.0.0.1:3306)/systemdb?charset=utf8mb4&parseTime=true&loc=Local",
		"maxIdle":    "10",
		"maxOpen":    "200",
	},
	"db1_slave": {
		"driverName": "mysql",
		"dsn":        "root:123456@tcp(127.0.0.1:3306)/systemdb?charset=utf8mb4&parseTime=true&loc=Local",
		"maxIdle":    "10",
		"maxOpen":    "200",
	},
}
