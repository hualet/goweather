package goweather

import (
	"net/http"

	"strings"

	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func parseContent(res *http.Response) (info *WeatherInfo, err error) {
	info = &WeatherInfo{}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	info.City = doc.Find("div.sk h1 span a").Text()
	info.Weather = []WeatherData{}
	doc.Find("div.days7 ul li").Each(func(i int, s *goquery.Selection) {
		data := WeatherData{}

		img := s.Find("i img").First()
		data.Code = weatherCodeFromImageURL(img.AttrOr("src", ""))
		data.Description = translateCode(data.Code)
		data.CDescription = img.AttrOr("alt", "")
		data.MinTemperature, data.MaxTemperature, _ = parseTemperature(s.Find("span").Text())

		info.Weather = append(info.Weather, data)
	})

	return info, nil
}

func weatherCodeFromImageURL(url string) string {
	if url == "" {
		return ""
	}

	tokens := strings.Split(url, "/")
	tokens = strings.Split(tokens[len(tokens)-1], ".")
	return tokens[0]
}

func parseTemperature(token string) (min, max int, err error) {
	extraChar := "℃"

	ts := strings.Split(token, "/")

	min, err = strconv.Atoi(strings.Trim(ts[0], extraChar))
	if err != nil {
		return 0, 0, ErrParseTemperature
	}

	if len(ts) == 2 {
		max, err = strconv.Atoi(strings.Trim(ts[1], extraChar))
		if err != nil {
			return min, 0, ErrParseTemperature
		}
		return min, max, nil
	}

	return min, min, nil
}
