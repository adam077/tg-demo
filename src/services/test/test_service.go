package test

import (
	"github.com/chenjiandongx/go-echarts/charts"
	"github.com/gin-gonic/gin"
	"go-go-go/src/data"
	"go-go-go/src/scheduler"
	"go-go-go/src/utils"
	"net/http"
	"sort"
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
		AddYAxis("test1", []int{1, 2, 3})
	bar.XYReversal().Render(c.Writer)
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
		bar.AddYAxis(content, rankData)
	}
	bar.Render(c.Writer)
}

func GetRank(rankMap map[int]int, hour int) int {
	if rank, ok := rankMap[hour]; !ok {
		return 0
	} else {
		return 50 - rank
	}
}

func EatWhat(c *gin.Context) {
	result := make([]map[string]string, 0)
	params := struct {
		Eat   string `form:"eat"`
		Reset string `form:"reset"`
	}{}
	c.ShouldBindQuery(&params)
	if params.Reset != "" {
		scheduler.Reset()
	} else {
		scheduler.Do2(params.Eat)
	}
	utils.SuccessResp(c, "", result)
}

func SeeEatWhat(c *gin.Context) {
	names := data.GetEatNames()
	if len(names) == 0 {
		return
	}

	result := scheduler.GetSortedEats(names)
	name := make([]string, 0)
	num := make([]int, 0)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count < result[j].Count
	})
	for x := range result {
		name = append(name, result[x].Name)
		num = append(num, result[x].Count)
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "吃啥呢"})
	bar.AddXAxis(name).
		AddYAxis("票数", num)
	bar.XYReversal().Render(c.Writer)
}
