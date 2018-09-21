package main

//json解析成结构体
import (
	"encoding/json"
	"fmt"
)

type IT2 struct {
	Company  string   `json:"-"`        //-此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOK     bool     `json:",string"`
	Price    float64  `json:",string"`
}

func main() {
	jsonBuf := `
		{
			"company":"53cc",
			"isok":"true",
			"price":"666.55",
			"subjects":[
				"go",
				"c++",
				"pathon",
				"Test"
			]
		}
	`

	var tmp IT2                                  //定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp) //将json转化为结构体
	if err != nil {
		fmt.Println("err=", err)
	}

	fmt.Printf("tmp=%+v\n", tmp)

	type IT3 struct {
		Subjects []string `json:"subjects"`
	}

	var temp2 IT3
	err = json.Unmarshal([]byte(jsonBuf), &temp2)
	if err != nil {
		fmt.Println("err=", err)
	}

	fmt.Printf("temp2=%+v\n", temp2)
}
