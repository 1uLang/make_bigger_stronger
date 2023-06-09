/*
 * @Author: Young
 * @Date: 2022-05-11 10:58:40
 * @LastEditors: lihy lihy@zhiannet.com
 * @LastEditTime: 2022-10-31 16:35:08
 * @FilePath: /zero-trust/console/IAM/common/lib/response/reponse.go
 */

package response

import (
	"net/http"
	"time"

	"github.com/1uLang/make_bigger_stronger/common/errors"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code uint32      `json:"code"`
	Time int64       `json:"time"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	body := Body{Data: resp, Time: time.Now().Unix()}

	if err != nil {
		errObj, _ := err.(*errors.Error)
		if errObj != nil {
			body.Code = uint32(errObj.Code())
			body.Msg = errObj.Msg()
		} else {
			// 未知错误码
			body.Code = errors.ErrCodeSystem
			body.Msg = "系统错误"
		}
	}

	httpx.OkJson(w, body)
}
func ResponseRarestone(w http.ResponseWriter, resp interface{}, err error) {

	httpx.OkJson(w, resp)
}

func ResponseRaw(w http.ResponseWriter, resp interface{}) {
	var body []byte
	switch v := resp.(type) {
	case []byte:
		body = v
	case string:
		body = []byte(v)
	default:
		body = []byte("1")
	}

	w.Write(body)
	w.WriteHeader(http.StatusOK)
}
