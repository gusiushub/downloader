package scraper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/google/uuid"
)

type Jut struct {
}

func (j *Jut) GetUrl(season int, episode int) string {

	c := colly.NewCollector()

	var url string

	c.OnHTML("video", func(e *colly.HTMLElement) {
		url = e.ChildAttr("source", "src")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	headers := make(http.Header)
	headers["user-agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.1.1140 Yowser/2.5 Safari/537.36"}
	c.Request("GET", fmt.Sprintf("https://jut.su/boku-hero-academia/season-%d/episode-%d.html", season, episode), nil, nil, headers) //@todo убрать название сериала в мапу

	return url
}

func (j *Jut) DownloadFile(URL, fileName string) error {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, URL, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}

	//@todo убрать в env
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.1.1140 Yowser/2.5 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		return errors.New("Received non 200 response code")
	}

	file, err := os.Create(j.getPath())
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (j *Jut) getPath() string {
	uuid := uuid.NewString()
	re := regexp.MustCompile(`(?is)^(?P<f1>\w{1})\w+\-(?P<f2>\w{1})\w+\-`)
	matches := re.FindStringSubmatch(uuid)

	path := filepath.Join("downloads", matches[1], matches[2])

	err := os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path + "/" + uuid + ".mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return path + "/" + uuid + ".mp4"
}
