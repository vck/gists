package main 

import "net/http"
import "log"
import "os"
import "fmt"

func main(){
	port, err := os.Getenv("PORT"); 
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	-, err := http.ListenAndServe(port, http.FileServer(http.Dir("/var/www")));
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
		fmt.Println("server running on localhost"+port)
}
