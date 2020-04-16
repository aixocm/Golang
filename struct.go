package main

import (
	"fmt"
)

type NewInt int 

type MyInt = int 

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n",a)
	fmt.Printf("type of b:%T\n",b)
}

type person struct {
	name string
	city string
	age int8
}

func main(){
	var p1 person
	p1.name = "pprof.cn"
	p1.city = "北京"
	p1.age  = 18
	fmt.Printf("p1=%v\n",p1)
	fmt.Printf("p1=%#v\n",p1)
	var p2 = new(person)
	fmt.Printf("%T\n",p2)
	fmt.Printf("p2=%#v\n",p2)

	p3 := &person{}

	p5 := person {
		name: "pprof.cn",
		city: "北京",
		age: 18,
	}

	p6 := &person{
		name: "pprof.cn",
		city: "北京",
		age: 18,
	}

	p7 := &pserson{
		city: "北京",
	}

	p8 := &person{
		"pprof.cn",
		"北京",
		18
	}
}

func main(){
	var user struct{Name string;Age int}
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%#v\n",user)

}


type test struct {
	a int8
	b int8
	c int8
	d int8
}



func main(){
	n := test{
		1,2,3,4,
	} 
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n",&n.b)
	fmt.Printf("n.c %p\n",&n.c)
	fmt.Printf("n.d %p\n", &n.d)
	
}


type person struct {
	name string
	city string
	age int8
}

func newPerson(name,city string,age int8) *person {
	return &person{
		name: name,
		city: city,
		age: age,
	}
}


func (p person) Dream(){
	fmt.Printf("%s的梦想是学好Go语言！\n",p.name)
}

func (p *person) SetAge(newAge int8){
	p.age = newAge
}


func main(){
	p9 := newPerson("pprof.cn","测试",90)
	fmt.Printf("%v\n",p9)
	p1 := newPerson("册数","计数",25)
	p1.Dream()

	p2 := newPerson("测试","cc",25)
	fmt.Println(p2.age)
	p2.SetAge(30)
	fmt.Println(p2.age)
}

type MyInt int 


func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int.")
}

func main(){
	var m1 MyInt
	m1.SayHello()
	m1 = 100
	fmt.Printf("$#v   %T\n",m1,m1)
}

type person struct {
	string
	int 
}

func main(){
	p1 := person{
		"pprof.cn",
		18,
	}
	fmt.Printf("%#v\n",p1)
	fmt.Println(p1.string,p1.int)
}


type Address struct {
	Province string
	City string
}

type User struct {
	Name string
	Gender string
	Address Address
}

func main(){
	user1 := User{
		Name: "pprof",
		Gender: "女",
		Address: Address{
			Province: "黑龙江",
			City: "哈尔滨",
		},
	}
	fmt.Printf("user1=%#v\n",user1)
}


type Address struct {
	Province string
	City string
}


type User struct {
	Name 	string
	Gender  string
	Address
}


func main(){
	var user2 User
	user2.Name = "pprof"
	user2.Gender = "女"
	user2.Address.Province = "黑龙江"
	user2.City = "哈尔滨"
	fmt.Printf("user2=%#v\n",user2)
}

type Address struct {
	Province string
	City string
	CreateTime string
}

type Email struct {
	Account string
	CreateTime string
}

type User struct {
	Name 	string
	Gender  string
	Address
	Email 
}


func main(){
	var user3 User 
	user3.Name  = "pprof"
	user3.Gender = "女"
	user3.Address.CreateTime = "2000"
	user3.Email.CreateTime = "2000"
	fmt.Println(user3)
}

type Animal struct {
	name string
}

func (a *Animal) move(){
	fmt.Printf("%s惠东！\n", a.name)
}


type Dog struct {
	Feet int8
	*Animal
}


func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪\n",d.name)
}

func main(){

	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}

	d1.wang()
	d1.move()
	
}

import "encoding/json"

type Student struct {
	ID 		int 
	Gender  string 
	Name  	string 
}

type Class struct {
	Title string 
	Students []*Student
}

func main(){

	c := &Class{
		Title: "101",
		Students: make([]*Student,0,200),
	}

	for i:=0;i<10;i++{
		stu   :=  &Student{
			Name: fmt.Sprintf("stu%02d",i),
			Gender: "男",
			ID:  i,
		}
		c.Students = append(c.Students, stu)		
	}

	data,err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return 
	}

	fmt.Printf("json:%s\n",data)

	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1  := &Class{}
	err = json.Unmarshal([]byte(str),c1)

	if err != nil {
		fmt.Println("json unmarshal failed!")
		return 
	}
	fmt.Printf("%v\n",*c1)
}

import "encoding/json"
type Student struct {
	ID int `json:"id"`
	Gender string
	name string
}


func main(){
	s1 := Student{
		ID: 1,
		Gender: "女",
		name: "pprof",
	}
	data,err := json.Marshal(s1)

	if err != nil {
		fmt.Println("json marshal failed!")
		return 
	}
	fmt.Printf("json str:%s\n",data)
}