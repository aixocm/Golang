package  main 

import "fmt"

// func main(){
// 	var s1 []int
	
// 	if s1 == nil {
// 		fmt.Println("是空")
// 	} else {
// 		fmt.Println("不是空")
// 	}

// 	s2 			:= []int{}

// 	var s3 []int = make([]int,0)
	
// 	fmt.Println(s1,s2,s3)

// 	var s4 []int = make([]int,0,0)

// 	fmt.Println(s4)
	
// 	s5 := []int{1,2,3}
	
// 	fmt.Println(s5)

// 	arr := [5]int{1,2,3,4,5}
// 	var s6 []int
	
// 	s6 = arr[1:4]
// 	fmt.Println(s6)

// }


// func main() {
// 	s1 := []int{0,1,2,3,8: 100}
// 	fmt.Println(s1,len(s1),cap(s1))

// 	s2 := make([]int,6,8)
// 	fmt.Println(s2,len(s2),cap(s2))

// 	s3 := make([]int,6)
// 	fmt.Println(s3,len(s3),cap(s3))

// }

// func main(){

// 	s := []int{0,1,2,3}

// 	p := &s[2]

// 	*p += 100

// 	fmt.Println(s)
// }

// func main(){
// 	d := [5]struct{
// 		x int
// 	}{}

// 	s := d[:]

// 	d[1].x = 10 
// 	s[2].x = 20

// 	fmt.Println(d)
// 	fmt.Printf("%p, %p\n", &d, &d[0])
	
// }


// func main(){
// 	var a = []int{1,2,3}
// 	fmt.Printf("slice a : %v\n", a)

// 	var b = []int{4,5,6}
// 	fmt.Printf("slice b : %v\n",b)

// 	c := append(a, b...)
// 	fmt.Printf("slice c : %v\n", c)

// 	d := append(c,7)
// 	fmt.Printf("slice d : %v\n", d)

// 	e := append(d,8,9,10)
// 	fmt.Printf("slice e : %v\n", e)
// }

// func main(){
// 	s1 := make([]int,0,5)
// 	fmt.Printf("%p\n", &s1)

// 	s2 := append(s1,1)

// 	fmt.Printf("%p\n",&s2)

// 	fmt.Println(s1,s2)
// }


// func main(){

// 	s := make([]int, 0,1)
// 	c := cap(s)


// 	for i := 0; i< 50; i++ {
// 		s = append(s,i)
// 		if n := cap(s); n > c {
// 			fmt.Printf("cap: %d -> %d\n",c,n)
// 			c = n
// 		}
// 	}
// }


// func main(){

// 	s1 := []int{1,2,3,4,5}
// 	fmt.Printf("slice s1 : %v\n",s1)
// 	s2 := make([]int,10)
// 	fmt.Printf("slice s2 : %v\n", s2)
// 	copy(s2,s1)
// 	fmt.Printf("copied slice s1 : %v\n",s1)
// 	fmt.Printf("copied slice s2 : %v\n",s2)
// 	s3 := []int{1,2,3}
// 	fmt.Printf("slice s3 : %v\n",s3)
// 	s3 = append(s3,s2...)
// 	fmt.Printf("appended slice s3 : %v\n",s3)
// 	s3 = append(s3,4,5,6)
// 	fmt.Printf("last slice s3 : %v\n",s3)
// }


// func main() {
// 	data := [...]int{0,1,2,3,4,5,6,7,8,9}
// 	fmt.Println("array data : ",data)
// 	s1 := data[8:]
// 	s2 := data[:5]
// 	fmt.Printf("slice s1 : %v\n",s1)
// 	fmt.Printf("slice s2 : %v\n",s2)
// 	copy(s2,s1)
// 	fmt.Printf("copied slice s1 : %v\n",s1)
// 	fmt.Printf("copied slice s2 : %v\n",s2)
// 	fmt.Println("last array data : ",data)
// }

// func main(){
// 	data  := [...]int{0,1,2,3,4,5,6,7,8,9}
// 	slice := data[:]
// 	for index, value := range slice {
// 		fmt.Printf("index : %v , value : %v\n",index,value)
// 	}
// }

// func main(){
// 	var a= []int{1,3,4,5}
// 	fmt.Printf("slice a :%v, len(a): %v\n",a,len(a))
	
// 	b := a[1:2]
// 	fmt.Printf("slice b : %v, len(b) : %v\n",b,len(b))
	
// 	c := b[0:3]
// 	fmt.Printf("slice c : %v, len(c) : %v\n",c,len(c))

// }

// func main(){
// 	str := "hello world"
// 	s1 := str[0:5]
// 	fmt.Println(s1)

// 	s2 := str[6:]
// 	fmt.Println(s2)

// }

// func main(){

// 	str := "Hello world"
// 	s 	:= []byte(str)
// 	s[6] = 'G'
// 	s = s[:8]
// 	s = append(s,'!')
// 	str = string(s)
// 	fmt.Println(str) 
// }

// func main(){
// 	str := "你好，世界！hello world! "
// 	s   :=  []rune(str)
// 	s[3] = '够'
// 	s[4] = '浪'
// 	s[12] = 'g'

// 	s = s[:14]

// 	str = string(s)

// 	fmt.Println(str)
// }

// func main(){
// 	slice := []int{0,1,2,3,4,5,6,7,8,9}
// 	d1 := slice[6:8]
// 	fmt.Println(d1,len(d1),cap(d1))
// 	d2 := slice[:6:8]
// 	fmt.Println(d2,len(d2),cap(d2))

// }