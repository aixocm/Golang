package main

import (
	"fmt"
)

func main(){
	scoreMap := make(map[string]int,8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100 
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n",scoreMap)
}


func main(){
	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo)
}

func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90 
	scoreMap["小明"] = 100
	v,ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("查无此人")
	}
}


func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	
	for k,v := range scoreMap {
		fmt.Println(k,v)
	}
}

func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	
	for k := range scoreMap {
		fmt.Println(k)
	}
}

func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	
	for k := range scoreMap {
		fmt.Println(k)
	}
}

func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	delete(scoreMap,"小明")

	for k,v := range scoreMap{
		fmt.Println(k,v)
	}
}

import "time"
import "math/rand"
import "sort"
func main(){
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int,200)
	
	for i :=0;i<100;i++ {
		key 	:= fmt.Sprintf("stu%02d",i)
		value 	:= rand.Intn(100)
		scoreMap[key] = value
	 }

	 var keys = make([]string,0,200)

	 for key  := range scoreMap {
		 keys  = append(keys,key)
	 }

	 sort.Strings(keys)
	 for _, key := range keys {
		 fmt.Println(key,scoreMap[key])
	 }
}

func main(){
	var mapSlice = make([]map[string]string,3)
	for index ,value := range mapSlice {
		fmt.Printf("index:%d value:%v\n",index,value)
	}

	fmt.Println("after init")
	mapSlice[0] = make(map[string]string,10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index,value := range mapSlice {
		fmt.Printf("index:%d value:%v\n",index,value)
	}
}


func main(){
	var sliceMap = make(map[string][]string,3)

	fmt.Println(sliceMap)
	fmt.Println("after init")

	key := "中国"
	value,ok := sliceMap[key]

	if !ok {
		value = make([]string,0,2)
	}

	value = append(value,"北京","上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}