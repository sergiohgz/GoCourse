package main

import (
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	wg.Add(len(sites)) // add as many threads as goroutes will be launch

	for _, site := range sites {
		go func(site string) {
			defer wg.Done() // defer thread sync to waitgroup if an error occurs before normal end

			res, err := http.Get(site)
			if err != nil {
			}

			io.WriteString(os.Stdout, res.Status+"\n")
		}(site)
	}

	wg.Wait()
}
