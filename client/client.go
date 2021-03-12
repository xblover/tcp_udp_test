package main

import (
	"fmt"

	"github.com/by-zhang/stpro"
)

type Client struct {
	Phost string
	Pmap  map[byte]string
}

func (c Client) Ptype(in []byte) {
	fmt.Printf("收到type包的回复：%s\n", in)
}

func (c Client) Pname(in []byte) {
	fmt.Printf("收到name包的回复：%s\n", in)
}

func main() {
	m := Client{
		Phost: "192.168.80.254:9091",
		Pmap: map[byte]string{
			0x01: "type",
			0x02: "name",
		},
	}
	client, err := stpro.NewClient(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	// tk1 := time.NewTicker(3 * time.Second)
	// for {
	// 	select {
	// 	case <-tk1.C:
	if err = client.Send(0x02, []byte("今天真好2222222")); err != nil {
		fmt.Println(err)
		return
	}
	if err = client.Send(0x01, []byte("今天真好1111111")); err != nil {
		fmt.Println(err)
		return
	}
	// }
	// }

}
