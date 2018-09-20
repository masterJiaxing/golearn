package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	Marshal()
	Unmarshal()
	Buffer()
}

//RawMessage类型是一个保持原本编码的json对象。本类型实现了Marshaler和Unmarshaler接口，用于延迟json的解码或者预计算json的编码
func Unmarshal() {

	type Color struct {
		Space string
		Point json.RawMessage //map[string]int // delay parsing until we know the color space
	}

	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}
	var j = []byte(`[
		{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
		{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
	]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}
	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
			fmt.Printf("dstRGB=%+v\n", dst.(*RGB)) //空接口转化为RGB结构体
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}

		fmt.Printf("c=%+v\n", c.Space)
		fmt.Printf("dst=%+v\n", dst)

		//fmt.Printf("c=%+v\n",c.Point["Y"])
	}
}

//结构体转json
func Marshal() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.MarshalIndent(group, "", "	")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func Buffer() {
	file := os.NewFile(1, "a.txt")
	file.WriteString("aefawefawf")

}

//打印数据类型
func isType(x interface{}) {
	fmt.Printf("type:%+v\n", reflect.TypeOf(x))
}
