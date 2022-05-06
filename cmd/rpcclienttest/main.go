package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := rpc.Dial("http://localhost:8089")
	if err != nil {
		panic(err)
	}

	err = client.Call(nil, "rpcdatastore_put", "1", []byte(`101`))
	if err != nil {
		panic(err)
	}

	err = client.Call(nil, "rpcdatastore_put", "2", []byte(`bcd`))
	if err != nil {
		panic(err)
	}

	var res2 []byte
	err = client.Call(&res2, "rpcdatastore_get", "2")
	if err != nil {
		panic(err)
	}

	var res1 []byte
	err = client.Call(&res1, "rpcdatastore_get", "1")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res2))
	fmt.Println(string(res1))
}
