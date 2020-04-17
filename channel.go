package main

import (
	"fmt"
)


import "time"

func main() {

	var c1,c2,c3 chan int 
	var i1,i2 int 

	select {
	case i1 = <-c1:
		fmt.Printf("received ",i1," from c1\n")

	case c2 <- i2:
		fmt.Printf("sent ",i2, " to c2\n")
	
	case i3, ok := <-c3:

		if ok {
			fmt.Printf("recived ",i3, " from c3\n")

		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}


var resChan = make(chan int)


func main() {
	

	select {
	case data := <-resChan:
		doData(data)
	
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}

}


func doData(data int) {
	fmt.Printf("adasd")

}


func main(){
	ch := make(chan int,5)
	data := 0
	ch <- data
	ch <-data
	ch <-data
	ch <-data
	ch <-data
	select {
	case ch <- data:
		fmt.Printf(" no man")
	default:
		fmt.Printf("man")
	}
}