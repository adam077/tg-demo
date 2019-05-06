package test

import (
	"github.com/chenjiandongx/go-echarts/charts"
	"github.com/gin-gonic/gin"
	"go-go-go/src/utils"
	"os"
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
