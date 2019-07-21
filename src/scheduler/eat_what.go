package scheduler

import (
	"fmt"
	"go-go-go/src/data"
	"go-go-go/src/ding-talk"
	"math/rand"
	"sort"
)

type EatWhat struct {
	do string
}

const (
	Choose      = "choose"
	Result      = "result"
	ResetResult = "reset"
)

var eatMap = map[string][]string{}

// service
// 发送投票选项  Task1
// 投票
// 发送投票结果 Task2
// 清空投票结果 ResetTask

// task
// 11、17点清空投票结果+发送投票选项
// 12、18点发送投票结果

func (runner EatWhat) Run() {
	switch runner.do {
	case Choose:
		ResetTask()
		Task1()
	case Result:
		Task2()
	}
}

func Reset() {
	eatMap = make(map[string][]string, 0)
}

func Task1() {
	// 发送投票选项
	names := data.GetEatNames()
	chatId := ""
	var dings = data.GetDingChatId("eat_what")
	if len(dings) > 0 {
		chatId = dings[0].ChatId
	}
	if chatId != "" {
		//go ding_talk.SendDingMessage(chatId, resultStr)
		choose := make([]ding_talk.DingChoose, 0)
		choose = append(choose, ding_talk.DingChoose{
			Title:     "完全随机",
			ActionURL: data.Env.SelfUrl + "/lv1/lv2/set_eat_what?eat=",
		})
		eatCountMap := make(map[string]int)
		for _, eats := range eatMap {
			for _, eat := range eats {
				eatCountMap[eat] = eatCountMap[eat] + 1
			}
		}

		for x := range names {
			name := names[x].Name
			if c, ok := eatCountMap[name]; ok {
				name = fmt.Sprintf("%s (%d)", name, c)
			}
			choose = append(choose, ding_talk.DingChoose{
				Title:     name,
				ActionURL: data.Env.SelfUrl + "/lv1/lv2/set_eat_what?eat=" + names[x].Name,
			})

		}
		go ding_talk.SendDingChoose(chatId, "投票啦", "每人限制点一次，或随机投出5张，或固定1张+随机4张", choose)
	}
}

func Task2() {
	// 发送投票结果
	chatId := ""
	var dings = data.GetDingChatId("eat_what")
	if len(dings) > 0 {
		chatId = dings[0].ChatId
	}
	if chatId != "" {
		go ding_talk.SendDingLink(chatId, ding_talk.Link{
			Text:       "点击查看结果",
			Title:      "吃啥好呢",
			MessageUrl: data.Env.SelfUrl + "/lv1/lv2/eat_what",
		})
	}
}

func ResetTask() {
	// 清空投票结果
	eatMap = make(map[string][]string, 0)
}

func EnrichEatMap(ip string, eat string) bool {
	if _, ok := eatMap[ip]; ok {
		return false
	}
	names := data.GetEatNames()
	eats := make([]string, 0)
	if eat == "" {
		ind := rand.Intn(len(names))
		eats = append(eats, names[ind].Name)
	} else {
		eats = append(eats, eat)
	}
	for i := 0; i < 4; i++ {
		ind := rand.Intn(len(names))
		eats = append(eats, names[ind].Name)
	}
	eatMap[ip] = eats
	return true
}

func GetSortedEats(names []*data.EatWhatTable) ([]Eat, []string) {
	result := make([]Eat, 0)
	eatCountMap := make(map[string]int)
	eatIpMap := make(map[string]map[string]int)
	ipList := make([]string, 0)
	for ip, eats := range eatMap {
		ipList = append(ipList, ip)
		for _, eat := range eats {
			eatCountMap[eat] = eatCountMap[eat] + 1
			if _, ok := eatIpMap[eat]; !ok {
				eatIpMap[eat] = make(map[string]int)
			}
			eatIpMap[eat][ip] = eatIpMap[eat][ip] + 1
		}
	}
	for x := range names {
		eat := names[x].Name
		result = append(result, Eat{
			Name:    eat,
			Count:   eatCountMap[eat],
			IpCount: eatIpMap[eat],
		})
	}
	sort.Sort(EatWrapper{result, EatWrapperOrder})
	return result, ipList

}

func EatWrapperOrder(p, q Eat) bool {
	return p.Count > q.Count
}

type Eat struct {
	Name    string
	Count   int
	IpCount map[string]int
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
