package main

import "fmt"
import "os"
import "lib"
func main(){
	var func_name string 
	func_name = os.Getenv("FUNC_NAME")
	fmt.Print(func_name+"---")
	fmt.Println(lib.Hello())
}