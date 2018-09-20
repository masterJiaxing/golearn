package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"-"`        //-此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOK     bool     `json:",string"`
	Price    float64  `json:",string"`
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"rabo", []string{"go", "c++", "大数据"}, true, 555.666}

	Marshal(s)
	MarshalIndent(s)
}

//生成json格式数据
func Marshal(s IT) {
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("buf=", string(buf))
}

//生产json格式数据并使用缩进格式化输出
func MarshalIndent(s IT) {
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("buf=", string(buf))
}
