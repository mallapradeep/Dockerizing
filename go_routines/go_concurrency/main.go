package main


import (
	"fmt"
	"sync"
)

//we hav min of 3 go routines running
//one in the func main n two other
//use waitgroups to synchronize your code

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i:= 0; i<45; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done()

}

func bar() {
	for i:=0; i<45; i++ {
		fmt.Println("Bar", i)
	}
	wg.Done()
}
