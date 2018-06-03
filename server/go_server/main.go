package main

import (
	"fmt"
	"log"
	"net/http"

	"cicada/server/go_server/router"
)

func main() {
	addr := ":8520"
	fmt.Println("listen addr", addr)

	if err := http.ListenAndServe(addr, router.NewMux()); err != nil {
		log.Fatal(err)
	}
}
