package sql

import (
	"encoding/xml"
	"fmt"
	"github.com/go-tempest/tempest/config"
	"io/ioutil"
	"strings"
)

var SQL map[string]string

type Xml struct {
	Mapper xml.Name `xml:"mapper"` //读取xml节点
	Sql    []Sql    `xml:"sql"`    //读取sql标签下到内容
}

type Sql struct {
	Id     string `xml:"id,attr"`   //读取id属性
	Script string `xml:",innerxml"` //读取 <![CDATA[ xxx ]]> 数据
}

func Initialize() {
	data, err := ioutil.ReadFile(config.TempestConfig.SqlXml.FileUrl)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Readload sql xml file has error:%s", err))
		return
	}
	sqlXml := Xml{}
	err = xml.Unmarshal(data, &sqlXml)
	if err != nil {
		fmt.Sprintf("reader sql xml has :%s", err)
		return
	}

	sqlList := make(map[string]string)
	for _, b := range sqlXml.Sql {
		sqlList[b.Id] = convStrForXml(b.Script)
	}
	SQL = sqlList
}

func convStrForXml(script string) string {
	script = strings.ReplaceAll(script, "&lt;", "<")
	script = strings.ReplaceAll(script, "&gt;", ">")
	return script
}
