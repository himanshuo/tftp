package main
 
import (
    "fmt"
    "net"
    "os"
    "github.com/himanshuo/tftp/packet"
    "github.com/himanshuo/tftp/transport"
    
    //"math/rand"
)
const MAXDATASIZE = 512


/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}


func routePacket(p packet.Packet, addr *net.UDPAddr){
	var t transport.Transport
	switch cur := p.(type) {
		case packet.ReadPacket:
			fmt.Println("read packet for filename:",string(cur.FileName))
			t = transport.NewReadTransport(cur, addr)

		case packet.WritePacket:
			fmt.Println("write packet for filename:",string(cur.FileName))
			t = transport.NewWriteTransport(cur, addr)
			
		case packet.ErrorPacket:
			fmt.Println("error packet with ErrMsg:",string(cur.ErrMsg))
		default:
			fmt.Println("got weird packet sent to port", PORT)
	}
	
	t.Start()
}

const PORT int = 10001 
func serve(){
	
	
	/* prepare a address at port 10001*/
	port := fmt.Sprintf(":%d", PORT)   
	fmt.Println("server up and listening at port",port)
    ServerAddr,err := net.ResolveUDPAddr("udp",port)
    fmt.Println(ServerAddr.String())
    CheckError(err)
 
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close() /* at end of function, ServerConn is closed. */
	
	buf := make([]byte, 516)
 
    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println(buf[:n])
        fmt.Println(err)
        fmt.Println(addr)
        fmt.Println(n)
        p := packet.ToPacket(buf[0:n])
        fmt.Println("Received ", p , " from ",addr)
		go routePacket(p, addr)
        CheckError(err)
    }
}



 
func main() {
    serve()
}
