package main

import (
	"fmt"
	"hosting/services/scraper"
)

func main() {

	scrap := scraper.MakeScraper("jut")
	i := 1
	for {
		j := 1
		for {
			url := scrap.GetUrl(i, j)
			fmt.Println(url)
			//разбить по потокам
			scrap.DownloadFile(url, fmt.Sprintf("boku-hero-academia-%d-%d.mp4", i, j))
			j++
		}
		i++
	}

}
