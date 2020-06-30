package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // defer cancel execution safely if something goes wrong

	var wg sync.WaitGroup

	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		// "https://maps.googl.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	wg.Add(len(sites)) // add as many threads as goroutes will be launch

	for _, site := range sites {
		go func(site string) {
			defer wg.Done() // defer thread sync to waitgroup if an error occurs before normal end

			select {
			case <-ctx.Done():
				return
			default:
				res, err := http.Get(site)
				if err != nil {
					io.WriteString(os.Stderr, err.Error()+"\n")
					cancel()
				}

				io.WriteString(os.Stdout, res.Status+"\n")
			}
		}(site)
	}

	wg.Wait()
}
