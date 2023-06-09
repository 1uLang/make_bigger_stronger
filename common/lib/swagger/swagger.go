/*
 * @Author: Young
 * @Date: 2021-05-28 16:34:29
 * @LastEditors: Young
 * @LastEditTime: 2021-05-28 18:59:42
 * @FilePath: /beecld/src/lib/swagger/swagger.go
 */
package swagger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/alecthomas/template"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/swag"
)

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

var doc string

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func LoadDoc(file string) (gin.HandlerFunc, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	doc = string(data)
	swag.Register(swag.Name, &s{})

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	return ginSwagger.WrapHandler(swaggerFiles.Handler, url), nil
}
