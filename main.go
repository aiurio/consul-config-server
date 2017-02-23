package main

import (
	"github.com/hashicorp/consul/api"
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/davecgh/go-spew/spew"
	"compress/gzip"
	"bytes"
	"io"
)

func main() {



	if bytes, err1 := bindataRead([]bytes(""), "bootstrap.yml"); err1 != nil {

	}

	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	pair, _, err := kv.Get("config/application/data", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: \n%v \n\n\n\n", string(pair.Value))

	config := &Config{}
	yaml.Unmarshal(pair.Value, config)
	spew.Dump(config)
}





func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type (
	Config struct{
		Info InfoConfig
	}
	InfoConfig struct {
		G2c string
	}
)
