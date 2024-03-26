package main

import (
	"fmt"
	"reflect"
)

func getType(v interface{}) string {
	switch v := v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "channel"
	default:
		return fmt.Sprintf("Unknown type: %v", reflect.TypeOf(v))
	}
}

func main() {
	var a interface{} = 10
	var b interface{} = "hello"
	var c interface{} = true
	var d interface{} = make(chan int)

	fmt.Println("Type of a:", getType(a))
	fmt.Println("Type of b:", getType(b))
	fmt.Println("Type of c:", getType(c))
	fmt.Println("Type of d:", getType(d))
}
