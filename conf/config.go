package conf

type Config struct {
	Name     string
	Host     string
	Port     string
	Token    string
	RedisPre string
	Auth     struct {
		AccessSecret string
		AccessExpire int64
	}
	Db         map[string]MysqlConf
	CacheRedis map[string]RedisConf
	Rpc        struct {
		Addr string
	}
	Kafka  []string
	Rabbit []string
	Es     struct {
		Addresses []string
		Username  string
		Password  string
	}
}

type (
	// A RedisConf is a redis config.
	RedisConf struct {
		Host []string
		Type string `json:",default=node,options=node|cluster"`
		Pass string `json:",optional"`
		Tls  bool   `json:",default=false,options=true|false"`
	}

	MysqlConf struct {
		Dsn string
	}
)
