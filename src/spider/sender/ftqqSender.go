package sender

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/op/go-logging"
)

const (
	ftqqHost     string = "https://sc.ftqq.com/%s.send?text=%s&desp=%s"
	ftqqPostHost string = "https://sc.ftqq.com/%s.send"
	ftqqKey      string = "xxx"
)

type FtqqSender struct {
	Subject string `json:"text"` // required
	Content string `json:"desp"` // support markDown
}

var ftqqLog = logging.MustGetLogger("ftqqSender")

func (sender *FtqqSender) send() {
	reqUrl := fmt.Sprintf(
		ftqqHost, ftqqKey,
		url.QueryEscape(sender.Subject),
		url.QueryEscape(sender.Content))

	// Mark: limit the length of url by the httpGet
	resp, err := http.Get(reqUrl)
	if err != nil {
		ftqqLog.Error("ftqqSend notic failed: ", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("ftqqSend notic failed: ", err.Error())
		ftqqLog.Error("ftqqSend notic failed: ", doc.Text())
	}
}

func (sender *FtqqSender) Send() {
	reqUrl := fmt.Sprintf(ftqqPostHost, ftqqKey)

	form := url.Values{
		"text":{sender.Subject},
		"desp":{sender.Content},
	}

	resp, err := http.PostForm(reqUrl, form)
	if err != nil {
		fmt.Println("ftqqSend notic failed: ", err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
