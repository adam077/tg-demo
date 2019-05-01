package api_query

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type WeiboRank struct {
	Rank    int    `json:"rank"`
	Content string `json:"content"`
}

func GetWeibo() []WeiboRank {
	requestConfig := &RequestConfig{
		Method: http.MethodGet,
		Url:    "https://s.weibo.com/top/summary?Refer=top_hot&topnav=1&wvr=6",
	}

	ret := make([]WeiboRank, 0)

	if result, err := GetResponseDataString(requestConfig); err != nil {
		log.Error().Msg(err.Error())
	} else {
		rank := 0
		s1, s2 := "Refer=top\" target=\"_blank\">", "</a>"
		//fmt.Println(result)
		a1 := strings.Split(result, "\n")
		for _, str1 := range a1 {
			if strings.Contains(str1, s1) && strings.Contains(str1, s2) {
				rank++
				a2 := strings.Split(str1, s1)
				a3 := strings.Split(a2[1], s2)
				ret = append(ret, WeiboRank{
					Rank:    rank,
					Content: a3[0],
				})
			}
		}
	}
	return ret
}
