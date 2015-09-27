package main
 
import (
    "fmt"
    "net"
    "os"
    "github.com/himanshuo/tftp/packet"
)
 
/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func storeFile(p packet.Packet){
	
	switch cur := p.(type) {
		case packet.ReadPacket:
			fmt.Println("read packet for filename:",string(cur.FileName))
		case packet.WritePacket:
			fmt.Println("write packet for filename:",string(cur.FileName))
		case packet.DataPacket:
			fmt.Println("data packet containing:",string(cur.Data))
		case packet.AckPacket:
			fmt.Println("acknowledge packet for some blocknum:",cur.BlockNum)
		case packet.ErrorPacket:
			fmt.Println("error packet with ErrMsg:",string(cur.ErrMsg))
		//case packet.AbstractPacket:
			//fmt.Println("abstract packet")
		//case packet.Packet:
			//fmt.Println("interface")
		default:
			fmt.Println("didn't match anytype")
	}
	
	
	
}

func serve(){
	/* Lets prepare a address at any address at port 10001*/   
    ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
    CheckError(err)
 
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close() /* at end of function, ServerConn is closed. */
	
	buf := make([]byte, 1024)
 
    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        p := packet.ToPacket(buf[0:n])
        fmt.Println("Received ", p , " from ",addr)
		storeFile(p)
		
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}



 
func main() {
    serve()
 
}
