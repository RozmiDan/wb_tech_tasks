package main

import "fmt"

func valueType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("type: int")
		return
	case string:
		fmt.Println("type: string")
		return
	case bool:
		fmt.Println("type: bool")
		return
	case chan int:
		fmt.Println("type: chan int")
		return
	case chan string:
		fmt.Println("type: chan string")
		return
	case chan bool:
		fmt.Println("type: chan bool")
		return
	}

	fmt.Printf("type: unsupported (%T)\n", value)
}

func main() {
	valueType(42)
	valueType("hello")
	valueType(true)

	ch1 := make(chan int)
	ch2 := make(chan bool)
	valueType(ch1)
	valueType(ch2)

	valueType(3.14)
}
