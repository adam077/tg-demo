package scheduler

import (
	"fmt"
	"go-go-go/src/data"
	"go-go-go/src/ding-talk"
	"math/rand"
	"sort"
	"time"
)

type EatWhat struct {
}

var eatMap = map[string]int{}

func (runner EatWhat) Run() {
	if time.Now().Hour() == 19 {
		Reset()
	}
	if data.GetConfig(data.RandomEat) != "" {
		Do2("")
	}
}

func Reset() {
	eatMap = make(map[string]int, 0)
}

func Do2(eat string) {
	names := data.GetEatNames()

	if eat == "" {
		if len(names) == 0 {
			return
		}
		ind := rand.Intn(len(names))
		eat = names[ind].Name
	}
	eatMap[eat] = eatMap[eat] + 1

	result := GetSortedEats(names)
	resultStr := "" //  "票数|吃啥\n"
	for x := range result {
		resultStr = resultStr + fmt.Sprintf("%d : %s \n", result[x].Count, result[x].Name)
	}
	chatId := ""
	var dings = data.GetDingChatId("eat_what")
	if len(dings) > 0 {
		chatId = dings[0].ChatId
	}
	if chatId != "" {
		//go ding_talk.SendDingMessage(chatId, resultStr)
		go ding_talk.SendDingLink(chatId, ding_talk.Link{
			Text:       resultStr,
			Title:      "吃啥好呢",
			MessageUrl: data.Env.SelfUrl + "/lv1/lv2/eat_what",
		})
	}

}

func GetSortedEats(names []*data.EatWhatTable) []Eat {
	result := make([]Eat, 0)
	for x := range names {
		result = append(result, Eat{
			Name:  names[x].Name,
			Count: eatMap[names[x].Name],
		})
	}
	sort.Sort(EatWrapper{result, EatWrapperOrder})
	return result

}

func EatWrapperOrder(p, q Eat) bool {
	return p.Count > q.Count
}

type Eat struct {
	Name  string
	Count int
}

type EatWrapper struct {
	Items []Eat
	By    func(p, q Eat) bool
}

func (wp EatWrapper) Len() int { // 重写 Len() 方法
	return len(wp.Items)
}
func (wp EatWrapper) Swap(i, j int) { // 重写 Swap() 方法
	wp.Items[i], wp.Items[j] = wp.Items[j], wp.Items[i]
}
func (wp EatWrapper) Less(i, j int) bool { // 重写 Less() 方法
	return wp.By(wp.Items[i], wp.Items[j])
}
