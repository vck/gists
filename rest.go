package	main

import "encoding/json"
import "net/http"

type Data struct{
	Title string `json:"title"`
	Data string `json:"data"`
}

func main(){
	http.HandleFunc("/", Model)
	http.ListenAndServe(":8080", nil)
}

func Model(w http.ResponseWriter, r	*http.Request)	{
	model	:=	Data{"title", "data"}
	js,	err	:=	json.Marshal(model)
		if	err	!=	nil	{
			http.Error(w,	err.Error(),	http.StatusInternalServerError)
			
			return
		}
		w.Header().Set("Content-Type",	"application/json")
		w.Write(js)
}