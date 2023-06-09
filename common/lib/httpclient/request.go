/*
 * @Author: Young
 * @Date: 2021-05-26 19:11:06
 * @LastEditors: lihy lihy@zhiannet.com
 * @LastEditTime: 2022-10-25 15:24:21
 * @FilePath: /zero-trust/console/IAM/common/lib/httpclient/request.go
 */
package httpclient

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/1uLang/make_bigger_stronger/common/lib/json"
	"go.uber.org/zap"
)

var (
	Transport = &http.Transport{
		MaxIdleConnsPerHost: 100,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	logger  *zap.Logger
	isDebug bool = false
)

type HttpClient struct {
	Url        string
	Headers    map[string]string
	Query      url.Values
	Body       []byte
	client     *http.Client
	err        error
	Response   *http.Response
	StatusCode int
	method     string
}

func NewRequest() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Transport: Transport,
			Timeout:   30 * time.Second,
		},
		Headers: make(map[string]string),
	}
}

func SetLogger(pLogger *zap.Logger) {
	logger = pLogger
}

func SetDebug(debug bool) {
	isDebug = debug
}

func (hc *HttpClient) SetQuery(query url.Values) *HttpClient {
	hc.Query = query
	return hc
}

func (hc *HttpClient) SetUrl(url string) *HttpClient {
	hc.Url = url
	return hc
}

func (hc *HttpClient) SetHeader(headers map[string]string) *HttpClient {
	for k, v := range headers {
		hc.Headers[k] = v
	}
	return hc
}

func (hc *HttpClient) SetHeaderKV(key string, val string) *HttpClient {
	hc.Headers[key] = val
	return hc
}

func (hc *HttpClient) Error() error {
	return hc.err
}

/**
 * @author: Young
 * @param {interface{}} data 请求参数，只允许传入struct
 * @description: post请求，可携带消息体
 * @return {*HttpClient}
 */
func (hc *HttpClient) Post(data interface{}) *HttpClient {

	var (
		bodyReader io.Reader
		err        error
	)
	if data != nil {
		bodyReader, err = json.MarshalReader(data)
		if err != nil {
			hc.err = err
			return hc
		}
	}

	hc.method = "POST"
	req, err := http.NewRequest(hc.method, hc.Url, bodyReader)
	if err != nil {
		hc.err = err
		return hc
	}

	for k, v := range hc.Headers {
		req.Header.Set(k, v)
	}

	if hc.Query != nil {
		req.URL.RawQuery = hc.Query.Encode()
	}

	hc.Response, err = hc.client.Do(req)
	if err != nil {
		hc.err = err
		return hc
	}
	hc.StatusCode = hc.Response.StatusCode

	defer hc.Response.Body.Close()
	body, err := ioutil.ReadAll(hc.Response.Body)
	if err != nil {
		hc.err = err
		return hc
	}

	hc.Body = body
	hc.debugLogger()
	return hc
}

func (hc *HttpClient) GetAuth(data interface{}, auth string) *HttpClient {

	var (
		bodyReader io.Reader
		err        error
	)
	if data != nil {
		bodyReader, err = json.MarshalReader(data)
		if err != nil {
			hc.err = err
			return hc
		}
	}

	hc.method = "POST"
	req, err := http.NewRequest("POST", hc.Url, bodyReader)
	if err != nil {
		hc.err = err
		return hc
	}

	for k, v := range hc.Headers {
		req.Header.Set(k, v)
	}

	if hc.Query != nil {
		req.URL.RawQuery = hc.Query.Encode()
	}

	response, err := hc.client.Do(req)
	if err != nil {
		hc.err = err
		return hc
	}
	hc.StatusCode = response.StatusCode

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		hc.err = err
		return hc
	}

	hc.Headers[auth] = response.Header.Get(auth)
	hc.Body = body
	hc.debugLogger()
	return hc
}

func (hc *HttpClient) Get() *HttpClient {
	hc.method = "GET"
	req, err := http.NewRequest("GET", hc.Url, nil)
	if err != nil {
		hc.err = err
		return hc
	}

	for k, v := range hc.Headers {
		req.Header.Set(k, v)
	}

	if hc.Query != nil {
		req.URL.RawQuery = hc.Query.Encode()
	}

	hc.Response, err = hc.client.Do(req)
	if err != nil {
		hc.err = err
		return hc
	}
	hc.StatusCode = hc.Response.StatusCode

	defer hc.Response.Body.Close()
	body, err := ioutil.ReadAll(hc.Response.Body)
	if err != nil {
		hc.err = err
		return hc
	}

	hc.Body = body
	hc.debugLogger()
	return hc
}

func (hc *HttpClient) Do(methd string, reqData interface{}) *HttpClient {

	var (
		bodyReader io.Reader
		err        error
	)
	if reqData != nil {
		switch d := reqData.(type) {
		case string:
			bodyReader = strings.NewReader(d)
		default:
			bodyReader, err = json.MarshalReader(reqData)
			if err != nil {
				hc.err = err
				return hc
			}
		}
	}

	hc.method = methd
	req, err := http.NewRequest(methd, hc.Url, bodyReader)
	if err != nil {
		hc.err = err
		return hc
	}

	for k, v := range hc.Headers {
		req.Header.Set(k, v)
	}

	if hc.Query != nil {
		req.URL.RawQuery = hc.Query.Encode()
	}

	hc.Response, err = hc.client.Do(req)
	if err != nil {
		hc.err = err
		return hc
	}
	hc.StatusCode = hc.Response.StatusCode

	defer hc.Response.Body.Close()
	body, err := ioutil.ReadAll(hc.Response.Body)
	if err != nil {
		hc.err = err
		return hc
	}

	hc.Body = body
	hc.debugLogger()
	return hc
}

func (hc *HttpClient) Put() *HttpClient {
	hc.method = "PUT"
	req, err := http.NewRequest(hc.method, hc.Url, nil)
	if err != nil {
		hc.err = err
		return hc
	}

	for k, v := range hc.Headers {
		req.Header.Set(k, v)
	}

	if hc.Query != nil {
		req.URL.RawQuery = hc.Query.Encode()
	}

	hc.Response, err = hc.client.Do(req)
	if err != nil {
		hc.err = err
		return hc
	}
	hc.StatusCode = hc.Response.StatusCode

	defer hc.Response.Body.Close()
	body, err := ioutil.ReadAll(hc.Response.Body)
	if err != nil {
		hc.err = err
		return hc
	}

	hc.Body = body
	hc.debugLogger()
	return hc
}

func (hc *HttpClient) GetHeader(key string) (string, error) {
	if hc.err != nil {
		return "", hc.err
	}
	if hc.Response != nil {
		return hc.Response.Header.Get(key), nil
	}
	return "", nil
}

func (hc *HttpClient) BindJSON(val interface{}) error {
	if hc.err != nil {
		return hc.err
	}
	err := json.Unmarshal(hc.Body, val)
	if err != nil {
		hc.err = err
	}
	return err
}

func (hc *HttpClient) debugLogger() {
	if isDebug {
		headerInfo, _ := json.Marshal(hc.Headers)
		logger.Debug("http request info", zap.String("method", hc.method), zap.String("url", hc.Url), zap.String("header", string(headerInfo)))
	}
}
