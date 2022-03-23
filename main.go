package main

import (
	"fmt"
	_ "go-admin/docs"
	"go-admin/http"
	"go-admin/lib/common"
)

/*
 * 首页
 * author Guo Zhiqiang
 * datetime 2019/10/21 11:53
 */
func main() {
	err := common.Load()
	if err != nil {
		fmt.Println(err)
	}
	//mail
	//example.MailTest()

	//cron 定时任务
	//cron.Start()

	//redis示例
	//example.RedisTest()

	//es, err := es.NewDefaultClient()
	//fmt.Println(es.Info, err)
	//es
	//example.EsTest()

	//rabbitmq
	//example.RabbitTest()
	//kafka
	//example.TestKafka()

	//数据库示例
	//example.DbTest()

	//支付
	//example.AlipayTest()

	//极光推送
	//example.Jpush()

	//启动web服务
	//http.Start()
	//启动后台web服务
	http.AdminStart()

	//启动rpc
	//Rpcserver.Start()
}
