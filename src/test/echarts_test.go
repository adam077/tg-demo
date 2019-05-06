package scheduler

import (
	"github.com/chenjiandongx/go-echarts/charts"
	"net/http"
	"os"
	"testing"
)

func hand(w http.ResponseWriter, _ *http.Request) {
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
	bar.Render(w, f)
}

func TestHa(t *testing.T) {
	http.HandleFunc("/", hand)
	http.ListenAndServe(":8080", nil)
}
