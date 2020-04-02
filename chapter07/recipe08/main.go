package main

import (
	"fmt"
	"net/http"
)

const addr = "localhost:7070"

type RedirectServer struct {
	redirectCount int
}

func (r *RedirectServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.redirectCount++
	fmt.Println("Received header: " + req.Header.Get("Known-redirects"))
	http.Redirect(rw, req, fmt.Sprintf("/redirect%d", r.redirectCount), http.StatusTemporaryRedirect)
}

func main() {

	s := http.Server{
		Addr:    addr,
		Handler: &RedirectServer{0},
	}
	go s.ListenAndServe()

	client := http.Client{}
	redirectCount := 0

	// follow redirects until limit reached
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirected")
		if redirectCount > 2 {
			return fmt.Errorf("too many redirects")
		}
		redirectCount++
		for _, prReq := range via {
			fmt.Printf("Previous request: %v\n", prReq.URL)
		}
		return nil
	}

	_, err := client.Get("http://" + addr)
	if err != nil {
		panic(err)
	}

}
