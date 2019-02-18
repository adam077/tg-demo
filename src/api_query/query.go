package api_query

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-go-go/src/utils"
	"net/http"
	"time"
)

type AutoGenerated struct {
	Time     string `json:"time"`
	CityInfo struct {
		City       string `json:"city"`
		CityID     string `json:"cityId"`
		Parent     string `json:"parent"`
		UpdateTime string `json:"updateTime"`
	} `json:"cityInfo"`
	Date    string `json:"date"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    struct {
		Shidu    string  `json:"shidu"`
		Pm25     float64 `json:"pm25"`
		Pm10     float64 `json:"pm10"`
		Quality  string  `json:"quality"`
		Wendu    string  `json:"wendu"`
		Ganmao   string  `json:"ganmao"`
		Forecast []struct {
			Date    string  `json:"date"`
			Sunrise string  `json:"sunrise"`
			High    string  `json:"high"`
			Low     string  `json:"low"`
			Sunset  string  `json:"sunset"`
			Aqi     float64 `json:"aqi"`
			Ymd     string  `json:"ymd"`
			Week    string  `json:"week"`
			Fx      string  `json:"fx"`
			Fl      string  `json:"fl"`
			Type    string  `json:"type"`
			Notice  string  `json:"notice"`
		} `json:"forecast"`
	} `json:"data"`
}

func GetWeather() (*AutoGenerated, error) {
	var (
		err error
	)
	today := time.Now()
	params := map[string]string{
		"beginDate": utils.TimeToDate(today.AddDate(0, 0, -15)),
		"endDate":   utils.TimeToDate(today.AddDate(0, 0, -1)),
	}
	requestConfig := &RequestConfig{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/%d", MultiApi.WeatherApi, 101030100),
		Params: params,
	}

	var tempStruct AutoGenerated
	if err = GetResponseData(requestConfig, &tempStruct); err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return &tempStruct, nil
}
