package main 
import "fmt"

func main(){
	data := [] byte("halo")
	data[0] = "j"
	fmt.Println(string(data))
}