package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dzhcool/genid/memcachep"
)

func main() {
	port := 7075
	ls, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if e != nil {
		log.Fatalf("Got an error:  %s", e)
	}
	memcachep.Listen(ls)
}
