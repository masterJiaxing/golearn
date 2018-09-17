package main
//方法的继承和重写
import "fmt"
type Person2 struct {
	name   string
	sex    byte
	age int
}

func (p *Person2) PrintInfo(){
	fmt.Printf("name=%s,sex=%c,age=%d\n", p.name,p.sex,p.age)
}

//有一个学生，继承person字段，成员和方法都会继承
type Student struct{
	Person2
	id   int
	addr string
}
//重写方法
func (p *Student) PrintInfo(){
	fmt.Printf("addr=%s,sex=%c,age=%d\n", p.addr,p.sex,p.age)
}

func main(){
	s := Student{Person2{"mike",'n',18},666,"bj"}
	s.PrintInfo()
	s.Person2.PrintInfo()

	//第二阶段测试题
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := data[7:]

	s2 := data[:5]

	copy(s2, s1)
	fmt.Println(data)
}
