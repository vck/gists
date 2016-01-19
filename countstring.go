package main 

import "fmt"
import "unicode/utf8"

func main(){
	str := "Hello, World!"
	fmt.Println(str, len(str), utf8.RuneCountInString(str))
}