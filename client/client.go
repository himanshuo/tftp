package main
 
import (
    "fmt"
    "net"
    //"time"
    //"strconv"
    "github.com/himanshuo/tftp/packet"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
 
    defer Conn.Close()
    
    //for {
    msg := "Himanshu"
    dataPacket := packet.DataPacket{packet.DATA, uint16(0xff), []byte("hi")}
        
    _,err = Conn.Write(packet.DataPacketToBytes(dataPacket))
    if err != nil {
		fmt.Println(msg, err)
    }
        
    //}
}
