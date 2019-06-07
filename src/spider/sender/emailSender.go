package sender

import (
	"fmt"

	"github.com/go-gomail/gomail"
	"github.com/op/go-logging"
)

const (
	fromAddress string = "xxx@xxx.com"
	fromName    string = "from spider"
	toAddress   string = "xxx@xxx.com"
	toName      string = "to spider"
	emailHost   string = "smtp.xxx.com"
	port        int    = 465
	userName    string = "xxx@xxx.com"
	password    string = "xxx"
)

type EmailSender struct {
	Subject string // required
	Content string // support html
}

var emailLog = logging.MustGetLogger("emailSender")

func (sender *EmailSender) Send() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", fromAddress, fromName)
	m.SetHeader("To", m.FormatAddress(toAddress, toName))
	m.SetHeader("Subject", sender.Subject)
	m.SetBody("text/html", sender.Content)

	d := gomail.NewPlainDialer(emailHost, port, userName, password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("emailSend notic failed: ", err.Error())
		ftqqLog.Error("emailSend notic failed: ", err.Error())
	}
}
