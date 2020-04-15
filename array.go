package  main

import (
	"fmt"
	"math/rand"
	"time"
)

// func test(x [2]int){
// 	fmt.Printf("x: %p\n", &x)
// 	x[1] = 1000
// }

// func main(){
// 	a := [2]int{}
// 	fmt.Printf("a: %p\n", &a)
// 	test(a)
// 	fmt.Println(a)
// }



// func printArr(arr *[5]int){
// 	arr[0] = 10 
// 	for i,v := range arr {
// 		fmt.Println(i,v)
// 	}
// }


// func main(){
// 	var arr1 [5]int
// 	printArr(&arr1)
// 	fmt.Println(arr1)
// 	arr2 := [...]int{2,4,6,8,10}
// 	printArr(&arr2)
// 	fmt.Println(arr2)
// }



func sumArr(a [10]int) int{
	var sum int = 0 
	for i :=0; i< len(a); i++ {
		sum +=a[i]
	}
	return sum 
}



func main() {
	rand.Seed(time.Now().Unix())
	var b [10]int
	for i :=0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Printf("sum=%d\n",sum)
}