package main 
import (
	"time" 
 	"github.com/kidoman/embd"
 	_ "github.com/kidoman/embd/host/rpi"
)

func main(){
for { 
	embd.LEDToggle("LED0")
	time.Sleep(250 * time.Millisecond)
}
}

