package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"spider/repository"
	"spider/sender"

	"github.com/PuerkitoBio/goquery"
	"github.com/op/go-logging"
)

// TODO: adjust to config
const link string = "https://www.douban.com/group/nanshanzufang/discussion?start="

var log = logging.MustGetLogger("spider")

func main() {
	repository.InitDB()

	once()
	time.Sleep(time.Second)

	for {
		count := 0
		rents := []repository.Rent{}
		for count <= 100 {
			source := link + strconv.Itoa(count)
			fmt.Println("load data from:", source)
			data := fetch(source)
			if len(data) > 0 {
				rents = append(rents, data...)
			}

			count += 25
		}
		send(rents, email)

		// TODO: more configable schedule
		time.Sleep(time.Minute * 20)
	}
}

func once() {
	count := 0
	rents := []repository.Rent{}
	for count <= 100 {
		source := link + strconv.Itoa(count)
		fmt.Println("load data from:", source)
		data := fetch(source)
		if len(data) > 0 {
			rents = append(rents, data...)
		}

		count += 25
	}
	send(rents, ftqq)
}

func fetch(link string) []repository.Rent {
	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("fail get data with httpCode:", resp.StatusCode)
		return nil
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error(err)
		return nil
	}

	rents := []repository.Rent{}
	doc.Find("tr[class!=th]").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			// skip title row
			return
		}

		src := s.Find("td.title").Find("a")
		title, _ := src.Attr("title")

		// match key words
		if !checkMatch(title) {
			return
		}

		url, _ := src.Attr("href")
		url = strings.TrimRight(url, "/")
		id := url[strings.LastIndex(url, "/")+1:]

		isExist := repository.CheckExist(id)
		if isExist {
			fmt.Printf("record with id: %s already exist\n", id)
			return
		}

		author := s.Find("td").Eq(1).Text()

		var rent repository.Rent
		rent.Url = url
		rent.Title = title
		rent.Author = author
		rent.Id, _ = strconv.Atoi(id)
		fmt.Println(rent)
		rents = append(rents, rent)
		repository.Insert(&rent)
	})

	return rents
}

func checkMatch(str string) bool {
	// TODO: adjust to config
	keys := [8]string{"两房", "二房", "2房", "两室", "二室", "2室", "2间", "二间"}
	for _, k := range keys {
		if strings.Contains(str, k) {
			return true
		}
	}
	return false
}

const (
	email = iota
	ftqq
)

func send(rents []repository.Rent, channel int) {
	if len(rents) > 0 {
		subject := fmt.Sprintf("%d new record: %s", len(rents), time.Now().Format("2006-01-02 15:04:05"))
		content := ""
		if channel == ftqq {
			for _, v := range rents {
				content = content + "- [" + v.Title + "](" + v.Url + ")  \n"
			}
			sender := sender.FtqqSender{
				Subject: subject,
				Content: content,
			}
			sender.Send()
		} else {
			for i, v := range rents {
				content = content + fmt.Sprintf("<p><a href=\"%s\">%d. %s</a></p>", v.Url, i, v.Title)
			}
			sender := sender.EmailSender{
				Subject: subject,
				Content: content,
			}
			sender.Send()
		}
	}
}
