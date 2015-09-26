package main

import (
	"github.com/himanshuo/tftp/packet"
	//"github.com/himanshuo/tftp/server"
	//"github.com/himanshuo/tftp/client"
	"fmt"
)

func main(){
   p := packet.ReadWritePacket{packet.AbstractPacket{packet.RRQ}, "myfile.txt", packet.OCTECT}
   fmt.Println(p.ToBytes())
}


