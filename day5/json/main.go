package main

//go语言中map转换成json
import (
	"encoding/json"
	"fmt"
)

type IT1 struct {
	Company  string   `json:"-"`        //-此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOK     bool     `json:",string"`
	Price    float64  `json:",string"`
}

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "51ct0"
	m["subjects"] = []string{"go", "c++", "pathon", "Test"}
	m["isok"] = true
	m["price"] = 555.555
	//编码成json
	Marshal1(m)
	MarshalIndent1(m)
}

//生成json格式数据
func Marshal1(s interface{}) {
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("buf=", string(buf))
}

//生产json格式数据并使用缩进格式化输出
func MarshalIndent1(s interface{}) {
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("buf=", string(buf))
}
