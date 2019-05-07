package test

import (
	"github.com/chenjiandongx/go-echarts/charts"
	"github.com/gin-gonic/gin"
	"go-go-go/src/data"
	"go-go-go/src/utils"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Person struct {
	Name          string `json:"name"`
	Score         int    `json:"score"`
	ScoreTHisTurn *int   `json:"scoreThisTurn"`
}

func GetPersons(c *gin.Context) {
	result := make([]Person, 0)
	asd := 15
	result = append(result, Person{
		Name:          "Adam Zhao",
		Score:         40,
		ScoreTHisTurn: &asd,
	})
	result = append(result, Person{
		Name:  "Adam Qian",
		Score: 30,
	})
	result = append(result, Person{
		Name:  "Adam Sun",
		Score: 20,
	})
	result = append(result, Person{
		Name:  "Adam Li",
		Score: 5,
	})
	utils.SuccessResp(c, "", result)
}

func GetEcharts(c *gin.Context) {
	n := []string{"hhh", "ddd", "asdf"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "yohaha"})
	bar.AddXAxis(n).
		AddYAxis("test1", []int{1, 2, 3}).
		AddYAxis("test2", []int{2, 3, 4})
	f, err := os.Create("bar.html")
	if err != nil {
		panic(err)
	}
	bar.Render(c.Writer, f)
}

func GetZhihuEcharts(c *gin.Context) {
	params := struct {
		LogDate string `form:"logDate"`
		Limit   int    `form:"limit"`
	}{}
	if err := c.ShouldBindQuery(&params); err != nil {
		utils.ErrorResp(c, http.StatusBadRequest, 0, "")
		return
	}
	if params.LogDate == "" {
		params.LogDate = utils.GetTimeDateString(time.Now())
	}
	if params.Limit == 0 {
		params.Limit = 3
	}
	contents, zhihuDatas := data.GetZhihuData(params.LogDate, params.Limit)
	hourList := make([]string, 0)
	for hour := 0; hour < 24; hour++ {
		hourList = append(hourList, strconv.Itoa(hour))
	}
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "zhihu"})
	zhihuMap := make(map[string]map[int]int)
	for _, ontData := range zhihuDatas {
		if _, ok := zhihuMap[ontData.Content]; !ok {
			zhihuMap[ontData.Content] = make(map[int]int)
		}
		zhihuMap[ontData.Content][ontData.LogHour] = ontData.Rank
	}
	bar.AddXAxis(hourList)
	for _, content := range contents {
		hourMap := zhihuMap[content]
		rankData := make([]int, 0)
		for hour := 0; hour < 24; hour++ {
			rankData = append(rankData, GetRank(hourMap, hour))
		}
		AddToBar(bar, content, rankData)
	}
	f, err := os.Create("bar.html")
	if err != nil {
		panic(err)
	}
	bar.Render(c.Writer, f)
}

func AddToBar(bar *charts.Line, name string, rankData []int) {
	bar.AddYAxis(name, rankData)
}

func GetRank(rankMap map[int]int, hour int) int {
	if rank, ok := rankMap[hour]; !ok {
		return 0
	} else {
		return 50 - rank
	}
}
