package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	ch := make(chan int)
	for range 15 {
		go func() {
			resp, err := client.Get("http://localhost:8080/request")
			ch <- 1
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(resp)
			}
		}()
	}
	for range 15 {
		<-ch
	}
}
