package scraper

func MakeScraper(stype string) IScraper {
	switch stype {
	case "jut":
		return &Jut{}
	default:
		return nil
	}
}
