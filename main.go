package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const DATE_FORMAT = "2006-01-02 15:04:05"
//获取之前时间
func getTime(times int) (start string) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start = thisMonth.AddDate(0, -times, 0).Format(DATE_FORMAT)
	//end := thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	//timeRange := fmt.Sprintf("%s~%s", start, end)
	return
}

func main (){
	curl()
	//isValidate("2018-08-01 00:00:00")
}

func isValidate(t string)(res bool){
	str := scan()//获取时间输入
	timeInt, e := strconv.Atoi(str)
	if e !=nil{
		fmt.Println("时间转换失败")
		return
	}
	token := scan1()//获取token输入
	if token == "" {
		fmt.Println("token获取失败")
		return
	}
	
	timeStr := getTime(timeInt)//获取输入的之前时间
	return timein(timeStr, t)
}
func timein(ti string, ti2 string)(res bool){
	timer :=toTime(ti2)//接口获取
	timerend :=toTime(ti)//输入
	return timer.Before(timerend)
}

func toTime(toBeCharge string) (st time.Time) {
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	st, err := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	if err!=nil{
		fmt.Println("读取时间异常")
	}
	return
}

func scan() (startTime string)  {
	startTime = ""
	input := bufio.NewScanner(os.Stdin)
	println("请输入删除几个月之前的数据？例如1：表示一个月")
	for input.Scan() {
		//控制循环退出
		if startTime == "" {
			startTime = input.Text()
		}

		if startTime !=""{
			break;
		}
	}
	return startTime
}

func scan1() (startTime string)  {
	startTime = ""
	input := bufio.NewScanner(os.Stdin)
	println("请输入token")
	for input.Scan() {
		//控制循环退出
		if startTime == "" {
			startTime = input.Text()
		}

		if startTime !=""{
			break;
		}
	}
	return
}
type Sell struct {
	Id int
	Name string
	Description string
	Name_with_namespace string
}
func curl(){
	url:="https://git.zmcms.cn/api/v4/projects?private_token=wjrgsq3aEtfDx1annWhb"
	header := map[string]string{
		"Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Encoding":"gzip, deflate",
		"Accept-Language":"zh-CN,zh;q=0.9",
		"Cache-Control":"max-age=0",
		"Connection":"keep-alive",
		"Cookie":"UM_distinctid=16682902e8c158-0d99123f8ac3c5-3c604504-100200-16682902e8d48d; _ga=GA1.2.1084651760.1539788780; _gid=GA1.2.89812774.1539788788; __guid=117854924.133888327568370320.1539789637405.4075; monitor_count=1",
		"Host":"t.weather.sojson.com",
		"Upgrade-Insecure-Requests":"1",
		"User-Agent":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	}
	spider := &Spider{url, header}
	res := spider.get_html_header()
	fmt.Printf("res=%+v\n", res)
}
//定义新的数据类型
type Spider struct {
	url    string
	header map[string]string
}



//定义 Spider get的方法
func (keyword Spider) get_html_header() interface{} {
	client := &http.Client{}
	println(keyword.url)
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
	}

	defer resp.Body.Close()

	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {

		}
	} else {
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
	}
	/*body = []byte(`{
			"price":"666.55",
			"subjects":[
				"go",
				"c++",
				"pathon",
				"Test"
			]
		}`)*/
		sells :=[]Sell{}
	err = json.Unmarshal(body,&sells)
	println(err)
	return sells

}
