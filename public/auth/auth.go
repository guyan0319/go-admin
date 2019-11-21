package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
)

func GetUid(c *gin.Context) interface{} {
	session := sessions.Default(c)
	v := session.Get(conf.Cfg.Token)
	if v == nil {
		return nil
	}
	return session.Get(v)
}