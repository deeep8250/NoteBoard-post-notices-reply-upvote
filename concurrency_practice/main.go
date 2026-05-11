package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websites := []string{

		"https://google.com",
		"https://microsoft.com",
		"https://gmail.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(web string) {
	defer wg.Done()
	res, err := http.Get(web)
	if err != nil {
		fmt.Println("OOPS : ", err.Error())
	}
	fmt.Println("status code : ", res.StatusCode)
}
