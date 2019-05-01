package scheduler

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-go-go/src/api_query"
	"go-go-go/src/data"
	"go-go-go/src/utils"
)

type HotSpotRunner struct {
}

func (runner HotSpotRunner) Run() {
	GetWeiBo()
	GetBaiDu()
	GetZhiHu()
}

func GetWeiBo() {
	ts, today, hour, min := utils.GetNowTime()
	datas := api_query.GetWeibo()

	if len(datas) == 0 {
		return
	}
	tb := data.WeiBoHotSpotMinuteReport{}.TableName()
	cols := []string{
		"create_time",
		"update_time",

		"log_date",
		"log_hour",
		"log_min",

		"content",
		"rank",

		"request_ts",
	}
	values := make([][]string, 0)
	keys := []string{"log_date", "log_hour", "log_min", "content"}
	updateCols := utils.GetUpdateTail(cols)
	for _, v1 := range datas {
		temp := []string{
			"now()",
			"now()",

			utils.GetStr(&today),
			utils.GetInt(&hour),
			utils.GetInt(&min),

			utils.GetStr(&v1.Content),
			utils.GetInt(&v1.Rank),

			fmt.Sprintf("to_timestamp(%d)", ts),
		}
		values = append(values, temp)
	}
	if len(values) == 0 {
		return
	}
	sqlStr := utils.CreateBatchSql(tb, cols, values, keys, updateCols)

	db := data.GetDataDB("default")
	err := db.Exec(sqlStr).Error
	if err != nil {
		log.Error().Err(err).Str("sql", sqlStr)
	}
	return
}

func GetBaiDu() {
	ts, today, hour, min := utils.GetNowTime()
	datas := api_query.GetBaidu()

	if len(datas) == 0 {
		return
	}
	tb := data.BaiDuHotSpotMinuteReport{}.TableName()
	cols := []string{
		"create_time",
		"update_time",

		"log_date",
		"log_hour",
		"log_min",

		"content",
		"rank",
		"searches",

		"request_ts",
	}
	values := make([][]string, 0)
	keys := []string{"log_date", "log_hour", "log_min", "content"}
	updateCols := utils.GetUpdateTail(cols)
	for _, v1 := range datas {
		temp := []string{
			"now()",
			"now()",

			utils.GetStr(&today),
			utils.GetInt(&hour),
			utils.GetInt(&min),

			utils.GetStr(&v1.Content),
			utils.GetInt(&v1.Rank),
			utils.GetInt(&v1.Searches),

			fmt.Sprintf("to_timestamp(%d)", ts),
		}
		values = append(values, temp)
	}
	if len(values) == 0 {
		return
	}
	sqlStr := utils.CreateBatchSql(tb, cols, values, keys, updateCols)

	db := data.GetDataDB("default")
	err := db.Exec(sqlStr).Error
	if err != nil {
		log.Error().Err(err).Str("sql", sqlStr)
	}
	return
}

func GetZhiHu() {
	ts, today, hour, min := utils.GetNowTime()
	datas := api_query.GetZhihu()

	if len(datas) == 0 {
		return
	}
	tb := data.ZhiHuHotSpotMinuteReport{}.TableName()
	cols := []string{
		"create_time",
		"update_time",

		"log_date",
		"log_hour",
		"log_min",

		"content",
		"rank",
		"score",
		"answer",
		"link",

		"request_ts",
	}
	values := make([][]string, 0)
	keys := []string{"log_date", "log_hour", "log_min", "content"}
	updateCols := utils.GetUpdateTail(cols)
	for _, v1 := range datas {
		temp := []string{
			"now()",
			"now()",

			utils.GetStr(&today),
			utils.GetInt(&hour),
			utils.GetInt(&min),

			utils.GetStr(&v1.Content),
			utils.GetInt(&v1.Rank),
			utils.GetFloat(&v1.Score),
			utils.GetInt(&v1.AnswerCount),
			utils.GetStr(&v1.Link),

			fmt.Sprintf("to_timestamp(%d)", ts),
		}
		values = append(values, temp)
	}
	if len(values) == 0 {
		return
	}
	sqlStr := utils.CreateBatchSql(tb, cols, values, keys, updateCols)

	db := data.GetDataDB("default")
	err := db.Exec(sqlStr).Error
	if err != nil {
		log.Error().Err(err).Str("sql", sqlStr)
	}
	return
}
