package lib

import (
	"net/smtp"
	"testing"
)

func TestMail(t *testing.T) {
	err := smtp.SendMail(
		"smtp.163.com:465",
		smtp.PlainAuth(
			"",
			"jeyrce@163.com",
			"MNQTGTBIQDOSGKZM",
			"smtp.163.com",
		),
		"jeyrce@163.com",
		[]string{"jianxin.lu@woqutech.com"},
		[]byte("lets go"),
	)
	if err != nil {
		t.Fatal(err)
	}
}
