package main

import (
	"github.com/himanshuo/tftp/packet"
	//"github.com/himanshuo/tftp/server"
	//"github.com/himanshuo/tftp/client"
	"fmt"
)

func main(){
   p := packet.NewWritePacket("myfile.txt")
   fmt.Println(p.ToBytes())
}


