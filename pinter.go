package main

import (
	"fmt"
)


func main(){
	a := 10
	b := &a

	fmt.Printf("a:%d ptr:%p\n",a,&a)
	fmt.Printf("b:%p type:%T\n",b,b)
	fmt.Println(&b)
	
}

func modify1(x int){
	x = 100
}

func modify2(x *int){
	*x = 100 
}

func main(){
	a := 10
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}

func main(){
	var p *string
	fmt.Printf("p的值是%s\n",p)

	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

func main(){
	a	:= new(int)
	b	:= new(bool)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Println(*a)
	fmt.Println(*b)
}

func main(){
	var a *int
	a = new(int)
	*a = 10 
	fmt.Println(*a)
}

func main(){
	var b map[string]int
	b = make(map[string]int,10)
	b["测试"] = 100
	fmt.Println(b)
}

func main(){
	 var a int 
	 fmt.Println(&a)
	 var p *int 
	 p = &a
	 *p = 20 
	 fmt.Println(a) 
}