package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"spider/repository"

	"github.com/PuerkitoBio/goquery"
	"github.com/op/go-logging"
)

// TODO: adjust to config
const link string = "https://www.douban.com/group/nanshanzufang/"

var log = logging.MustGetLogger("example")

func main() {
	repository.InitDB()

	// TODO: set schedule
	start()
}

func start() {
	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("fail get data with httpCode:", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error(err)
	}

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
		repository.Insert(&rent)
	})
}

func checkMatch(str string) bool {
	// TODO: adjust to config
	keys := [5]string{"两房", "2房", "两室", "2室", "2间"}
	for _, k := range keys {
		if strings.Contains(str, k) {
			return true
		}
	}
	return false
}
