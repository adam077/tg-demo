package api_query

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

type ZhihuRank struct {
	Rank        int     `json:"rank"`
	Content     string  `json:"content"`
	Link        string  `json:"-"`
	Score       float64 `json:"-"`
	AnswerCount int     `json:"-"`
}

type ZhihuQuery struct {
	Data []struct {
		StyleType    string `json:"style_type"`
		FeedSpecific struct {
			Trend       int     `json:"trend"`
			Score       float64 `json:"score"`
			Debut       bool    `json:"debut"`
			AnswerCount int     `json:"answer_count"`
		} `json:"feed_specific"`
		Target struct {
			LabelArea struct {
				Trend       int    `json:"trend"`
				Type        string `json:"type"`
				NightColor  string `json:"night_color"`
				NormalColor string `json:"normal_color"`
			} `json:"label_area"`
			MetricsArea struct {
				Text string `json:"text"`
			} `json:"metrics_area"`
			TitleArea struct {
				Text string `json:"text"`
			} `json:"title_area"`
			ExcerptArea struct {
				Text string `json:"text"`
			} `json:"excerpt_area"`
			ImageArea struct {
				URL string `json:"url"`
			} `json:"image_area"`
			Link struct {
				URL string `json:"url"`
			} `json:"link"`
		} `json:"target"`
		CardID       string `json:"card_id"`
		AttachedInfo string `json:"attached_info"`
		Type         string `json:"type"`
		ID           string `json:"id"`
	} `json:"data"`
}

func GetZhihu() []ZhihuRank {
	var (
		err error
	)
	requestConfig := &RequestConfig{
		Method: http.MethodGet,
		Url:    "https://www.zhihu.com/api/v3/feed/topstory/hot-list-web?limit=50&desktop=true",
	}

	result := make([]ZhihuRank, 0)

	var tempStruct ZhihuQuery
	if err = GetResponseData(requestConfig, &tempStruct); err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	i := 0
	for _, data := range tempStruct.Data {
		i++
		result = append(result, ZhihuRank{
			Rank:        i,
			Content:     data.Target.TitleArea.Text,
			Link:        data.Target.Link.URL,
			Score:       data.FeedSpecific.Score,
			AnswerCount: data.FeedSpecific.AnswerCount,
		})
	}

	return result
}
