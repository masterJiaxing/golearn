package main

//分别通过if和switch实现断言  判断接口类型
import "fmt"

type Student4 struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "adc"
	i[2] = Student4{"mike", 5}
	//index下标 对应下表值data i[0], i[1], i[2]
	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d]  类型为int，内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string，内容为%s\n", index, value)
		} else if value, ok := data.(Student4); ok == true {
			fmt.Printf("x[%d] 类型为Student4，内容为%s,id=%d\n", index, value.name, value.id)
		}
	}

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]  类型为int，内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为string，内容为%s\n", index, value)
		case Student4:
			fmt.Printf("x[%d] 类型为Student4，内容为%s,id=%d\n", index, value.name, value.id)
		}
	}
}
