package api_query

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

type BaiduRank struct {
	Rank     int    `json:"rank"`
	Content  string `json:"content"`
	Searches int    `json:"-"`
}

type BaiduQuery struct {
	Data []struct {
		Keyword  string `josn:"keyword"`
		Searches int    `json:"searches"`
	} `json:"data"`
}

func GetBaidu() []BaiduRank {
	var (
		err error
	)
	requestConfig := &RequestConfig{
		Method: http.MethodGet,
		Url:    "https://zhidao.baidu.com/question/api/hotword?rn=50",
	}

	result := make([]BaiduRank, 0)

	var tempStruct BaiduQuery
	if err = GetResponseData(requestConfig, &tempStruct); err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	i := 0
	for _, data := range tempStruct.Data {
		i++
		result = append(result, BaiduRank{
			Rank:     i,
			Content:  data.Keyword,
			Searches: data.Searches,
		})
	}

	return result
}
