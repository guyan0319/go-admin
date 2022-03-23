package mail

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
	"strings"
)

type MailMap struct {
	Mail map[string]MailConf
}
type MailConf struct {
	Username string
	Passwd   string
	Port     string
	Host     string
}

var conf MailMap

func init() {
	v := viper.New()
	v.AddConfigPath("./etc/")
	v.SetConfigType("yaml")
	v.SetConfigName("mail")
	if err := v.ReadInConfig(); err != nil {
		log.Printf("error: mail config init  %s", err.Error())
		return
	}
	if err := v.Unmarshal(&conf); err != nil {
		log.Printf("error: mail config init  %s", err.Error())
		return
	}
}
func sendaAuth(m MailConf) (smtp.Auth, error) {
	return smtp.PlainAuth("", m.Username, m.Passwd, m.Host), nil
}
func Send(t string, to []string, subject, body, mailType string) error {
	mail := conf.Mail
	fmt.Println(mail)
	m, ok := mail[t]
	if !ok {
		return errors.New("mail config is error")
	}
	var content_type string
	if mailType == "html" {
		content_type = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	auth, err := sendaAuth(m)
	if err != nil {
		return err
	}
	addr := fmt.Sprintf("%s:%s", m.Host, m.Port)
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + m.Username +
		"<" + m.Username + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err = smtp.SendMail(addr, auth, m.Username, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
	return nil
}
