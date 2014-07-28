package crawlers

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/dukex/moraraqui/models"
)

const (
	imovelwebbaseurl = "http://www.imovelweb.com.br"
)

type ImovelWebBot struct {
}

func (i *ImovelWebBot) FirstRun(channel chan *Property, state, city, neighborhood string) int {
	url := i.urlFor("1", state, city, neighborhood)
	doc := i.parserPage(channel, url)

	lastPageS := doc.Find(".box-pagging .bt-pagging-num--p").Last().Text()
	lastPage, _ := strconv.Atoi(lastPageS)
	return lastPage
}

func (i *ImovelWebBot) Get(channel chan *Property, page int, state, city, neighborhood string) {
	pageS := strconv.Itoa(page)

	url := i.urlFor(pageS, state, city, neighborhood)
	i.parserPage(channel, url)
}

// START OMIT
func (i *ImovelWebBot) parserPage(channel chan *Property, url string) *goquery.Document {
	log.Println(" Parsing", url, "...") // OMIT
	// OMIT
	// ...
	doc, err := goquery.NewDocument(url) // OMIT
	if err != nil {                      // OMIT
		return doc // OMIT
	} // OMIT
	// OMIT
	doc.Find("ul[itemtype='http://www.schema.org/RealEstateAgent'] > li").Each(func(_ int, s *goquery.Selection) {

		pType := i.getType(s.Find(".busca-item-heading2").Text())
		// OMIT
		if pType > 0 {
			url, _ := s.Find("a").Attr("href") // OMIT
			var property Property
			property.Title = s.Find(".busca-item-heading1").Text()
			property.Address = s.Find(".busca-item-endereco").Text()
			// ...
			property.Url = imovelwebbaseurl + url                           // OMIT
			property.Type = pType                                           // OMIT
			property.Value = i.getValue(s.Find(".busca-item-preco").Text()) // OMIT
			property.Neighborhood = strings.Split(property.Title, ",")[0]   // OMIT
			channel <- &property
		}
	})

	return doc
}

// END OMIT

func (i *ImovelWebBot) getValue(text string) float64 {
	text = strings.Replace(text, "R$ ", "", -1)
	text = strings.Replace(text, ".", "", -1)
	value, _ := strconv.ParseFloat(text, 64)
	return value
}

func (i *ImovelWebBot) getType(text string) int {
	text = strings.TrimSpace(text)
	splited := strings.Split(text, " ")

	switch splited[0] {
	case "Casa":
		return House
	case "Apartamento":
		return Apartament
	}

	return Undefined
}

func (i *ImovelWebBot) urlFor(page, state, city, neighborhood string) string {
	base := "aluguel"
	query := "?pg=" + page
	return strings.Join([]string{imovelwebbaseurl, base, state, city, neighborhood, query}, "/")
}
