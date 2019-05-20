package main

import (
	"fmt"
	"time"

	"github.com/dzhcool/genid/snow"
)

func main() {
	fmt.Println("run.")

	stime := time.Now()
	for i := 0; i < 10000000; i++ {
		snow.New().GenId()
		// fmt.Println(fmt.Sprintf("[%d] %d", i, id))
	}
	etime := time.Now()
	uptime := etime.Sub(stime)
	fmt.Println("use:", uptime)
}
