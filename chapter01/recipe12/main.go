package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	consulIP string
	connStr  string
}

func (c *Client) String() string {
	return fmt.Sprintf("ConsulIP: %s, Connection String: %s\n",
		c.consulIP, c.connStr)
}

var defaultClient = Client{
	consulIP: "localhost:9000",
	connStr:  "postgres://localhost:5432",
}

// Serves as a type to be used in options
type ConfigFunc func(opt *Client)

// Returns the ConfigFunc type
func FromFile(path string) ConfigFunc {
	return func(opt *Client) {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		decoder := json.NewDecoder(f)

		fop := struct {
			ConsulIP string `json:"consul_ip"`
		}{}
		err = decoder.Decode(&fop)
		if err != nil {
			panic(err)
		}
		opt.consulIP = fop.ConsulIP
	}
}

// Reads config from env vars
func FromEnv() ConfigFunc {
	return func(opt *Client) {
		connStr, exist := os.LookupEnv("CONN_DB")
		if exist {
			opt.connStr = connStr
		}
	}
}

func NewClient(opts ...ConfigFunc) *Client {
	client := defaultClient
	for _, val := range opts {
		val(&client)
	}
	return &client
}

func main() {
	client := NewClient(FromFile("config.json"), FromEnv())
	fmt.Println(client.String())
}
