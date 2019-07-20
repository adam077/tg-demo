package ding_talk

import (
	"bytes"
	"github.com/json-iterator/go"
	"go-go-go/src/api_query"
	"net/http"
)

type aa struct {
	Content string `json:"content"`
}

func SendDingMessage(chatId, message string) error {
	if chatId == "" {
		return nil
	}
	body := struct {
		MsgType string `json:"msgtype"`
		Text    aa     `json:"text"`
	}{
		MsgType: "text",
		Text: aa{
			Content: message,
		},
	}

	requestData, err := jsoniter.Marshal(&body)
	if err != nil {
		return err
	}

	requestConfig := &api_query.RequestConfig{
		Method: http.MethodPost,
		Url:    "https://oapi.dingtalk.com/robot/send",
		Params: map[string]string{
			"access_token": chatId,
		},
		Body: bytes.NewReader(requestData),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	if err := api_query.GetResponseData(requestConfig, nil); err != nil {
		return err
	}
	return nil
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

func SendDingLink(chatId string, link Link) error {
	if chatId == "" {
		return nil
	}
	body := struct {
		MsgType string `json:"msgtype"`
		Link    Link   `json:"link"`
	}{
		MsgType: "link",
		Link:    link,
	}

	requestData, err := jsoniter.Marshal(&body)
	if err != nil {
		return err
	}

	requestConfig := &api_query.RequestConfig{
		Method: http.MethodPost,
		Url:    "https://oapi.dingtalk.com/robot/send",
		Params: map[string]string{
			"access_token": chatId,
		},
		Body: bytes.NewReader(requestData),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	if err := api_query.GetResponseData(requestConfig, nil); err != nil {
		return err
	}
	return nil
}
