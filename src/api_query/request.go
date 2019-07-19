package api_query

import (
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

type RequestConfig struct {
	Method                string
	Url                   string
	Headers               map[string]string
	Params                map[string]string
	ListParams            map[string][]string
	Body                  io.Reader
	ErrorMessageOverrider map[int]string // code - message 对应关系，用于重载返回的message
}

var HttpClient = &http.Client{
	Timeout: time.Minute,
}

func GetResponseData(config *RequestConfig, respData interface{}) error {
	/*
		请求构建的RequestConfig，将返回结构中的data字段解析成respData(若respData不为nil)
	*/
	var (
		err      error
		req      = new(http.Request)
		resp     = new(http.Response)
		rawBytes []byte
	)

	if req, err = http.NewRequest(config.Method, config.Url, config.Body); err != nil {
		log.Error().Err(err).Msg("无法构建请求")
		return errors.New("无法构建请求")
	}
	// 不让BOSTEN_DATA_CONTEXT使用gzip传递数据
	req.Header.Set("Accept-Encoding", "identity")
	for k, v := range config.Headers {
		req.Header.Set(k, v)
	}

	// 构建请求的参数
	q := req.URL.Query()
	for k, v := range config.ListParams {
		if len(v) > 0 {
			for _, value := range v {
				q.Add(k, value)
			}
		}
	}
	for k, v := range config.Params {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()

	if resp, err = HttpClient.Do(req); err != nil {
		message, _ := httputil.DumpRequest(req, true)
		log.Error().Err(err).Str("message", string(message)).Msg("数据后端返回错误")
		return errors.New("无法请求数据后台")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		message, _ := httputil.DumpRequest(req, true)
		log.Error().Int("statusCode", resp.StatusCode).Str("message", string(message)).Msg("数据后端返回HTTP_STATUS_CODE")
		fmt.Println(string(message))
		return errors.New("无法请求数据后台")
	}
	if rawBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		message, _ := httputil.DumpRequest(req, true)
		log.Error().Err(err).Str("message", string(message)).Msg("读取数据后台返回结果错误")
		return errors.New("读取数据后台返回结果错误")
	}
	if respData != nil {
		if err = jsoniter.Unmarshal(rawBytes, respData); err != nil {
			message, _ := httputil.DumpRequest(req, true)
			log.Error().Err(err).Str("message", string(message)).Msg("解析数据后台返回结果错误")
			return errors.New("无法解析后台数据")
		}
	}
	return nil
}

func GetResponseDataString(config *RequestConfig) (string, error) {
	/*
		请求构建的RequestConfig，将返回结构中的data字段解析成respData(若respData不为nil)
	*/
	var (
		err      error
		req      = new(http.Request)
		resp     = new(http.Response)
		rawBytes []byte
	)

	if req, err = http.NewRequest(config.Method, config.Url, config.Body); err != nil {
		log.Error().Err(err).Msg("无法构建请求")
		return "", errors.New("无法构建请求")
	}
	// 不让BOSTEN_DATA_CONTEXT使用gzip传递数据
	req.Header.Set("Accept-Encoding", "identity")
	for k, v := range config.Headers {
		req.Header.Set(k, v)
	}

	// 构建请求的参数
	q := req.URL.Query()
	for k, v := range config.ListParams {
		if len(v) > 0 {
			for _, value := range v {
				q.Add(k, value)
			}
		}
	}
	for k, v := range config.Params {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()

	if resp, err = HttpClient.Do(req); err != nil {
		message, _ := httputil.DumpRequest(req, true)
		log.Error().Err(err).Str("message", string(message)).Msg("数据后端返回错误")
		return "", errors.New("无法请求数据后台")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		message, _ := httputil.DumpRequest(req, true)
		log.Error().Int("statusCode", resp.StatusCode).Str("message", string(message)).Msg("数据后端返回HTTP_STATUS_CODE")
		fmt.Println(string(message))
		return "", errors.New("无法请求数据后台")
	}
	if rawBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		message, _ := httputil.DumpRequest(req, true)
		log.Error().Err(err).Str("message", string(message)).Msg("读取数据后台返回结果错误")
		return "", errors.New("读取数据后台返回结果错误")
	}
	return string(rawBytes), nil
}
