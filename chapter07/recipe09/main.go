package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const addr = "localhost:7070"

type City struct {
	ID       string
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (c City) toJSON() string {
	jsonStr, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(jsonStr)
}

func saveCity(city City) (City, error) {
	r, err := http.Post("http://"+addr+"/cities",
		"application/json",
		strings.NewReader(city.toJSON()))
	if err != nil {
		return City{}, err
	}
	defer r.Body.Close()
	return decodeCity(r.Body)
}

func getCities() ([]City, error) {
	r, err := http.Get("http://" + addr + "/cities")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return decodeCities(r.Body)
}

func decodeCity(r io.Reader) (City, error) {
	city := City{}
	err := json.NewDecoder(r).Decode(&city)
	return city, err
}

func decodeCities(r io.Reader) ([]City, error) {
	cities := make([]City, 0)
	err := json.NewDecoder(r).Decode(&cities)
	return cities, err
}

func createServer(addr string) http.Server {
	cities := []City{
		City{
			"0",
			"Prague",
			"Czechia",
		},
		City{
			"1",
			"Bratislava",
			"Slovakia",
		},
		City{
			"2",
			"Ridgefield",
			"USA",
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/cities", func(w http.ResponseWriter, req *http.Request) {
		enc := json.NewEncoder(w)
		if req.Method == http.MethodGet {
			enc.Encode(cities)
		} else if req.Method == http.MethodPost {
			data, err := ioutil.ReadAll(req.Body)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			req.Body.Close()
			city := City{}
			json.Unmarshal(data, &city)
			city.ID = strconv.Itoa(len(cities))
			cities = append(cities, city)
			enc.Encode(city)
		}
	})

	return http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func main() {

	s := createServer(addr)
	go s.ListenAndServe()

	cities, err := getCities()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Retrieved cities: %v\n", cities)

	city, err := saveCity(City{"", "Paris", "France"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saved city: %v\n", city)

}
