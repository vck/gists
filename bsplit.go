package main 

import "fmt"
import "bytes"

func main(){
	target := [] byte("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z")
	splitter := [] byte(",")
	splitted_string := bytes.Split(target, splitter)
	for i := 0; i < len(splitted_string); i++ {
		fmt.Println(string(splitted_string[i]))
	}
	fmt.Println(splitted_string, len(splitted_string))
}