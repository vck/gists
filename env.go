package main 

import ("fmt"
		"os"
		"net/http"
)

func main(){
	port := os.Getenv("PORT")
	path := "/var/www"
	fmt.Println("server running on localhost:"+port)
	http.ListenAndServe(":"+port, http.FileServer(http.Dir(path)))
	
}