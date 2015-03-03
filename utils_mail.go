package goutils

import (
	"fmt"
	"net/smtp"
	"strings"
)

/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */

func SendMail(user, password, host, to, subject, body string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	content_type = "Content-Type: text/plain" + "; charset=UTF-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	if len(send_to) == 0 {
		return nil
	}

	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

// to "jixiuf@qq.com;jixiuf@gmail.com" //
func SendGMMail(to, subject, body string) (err error) {
	user := "devops-sky@3ciyuan.com"
	password := "ouluopssky2"
	host := "smtp.exmail.qq.com:25"
	err = SendMail(user, password, host, to, subject, body)
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
	return
}
