/**
三步搭建服务端
1.定义含有Pmap,Phost 的struct，Phost为服务端IP+PORT拼接的字符串
Pmap为自定义数据包类型与数据包名称的映射

2.实例化对象为字段赋值，实现对应已定义`包名称`的数据包处理方法，
方法名必为"P[包名称]"，如type包的处理方法为Ptype。方法中请定义数据处理逻辑，
输入输出皆为[]byte类型

3.stpro.New() 传入实例化对象，如无报错则服务端开始监听，
并按照你所定义的逻辑处理数据包，返回响应数据
**/

package main

import (
	"fmt"

	"github.com/by-zhang/stpro"
)

type Server struct {
	Phost string
	Pmap  map[uint8]string
}

func (m Server) Ptype(in []byte) (out []byte) {
	fmt.Printf("客户端发来type包：%s\n", in)
	/** process... **/
	return []byte("hello1")
}

func (m Server) Pname(in []byte) (out []byte) {
	fmt.Printf("客户端发来name包：%s\n", in)
	/** process... **/
	return []byte("hello2")
}

func main() {
	m := Server{
		Phost: ":9091",
		Pmap:  make(map[uint8]string),
	}
	m.Pmap[0x01] = "type"
	m.Pmap[0x02] = "name"

	if err := stpro.New(m); err != nil {
		fmt.Println(err)
	}

}
