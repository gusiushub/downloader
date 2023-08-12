package scraper

type IScraper interface {
	getPath() string
	GetUrl(season int, episode int) string
	DownloadFile(URL, fileName string) error
}
