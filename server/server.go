package main
 
import (
    "fmt"
    "net"
    "os"
    "github.com/himanshuo/tftp/packet"
    "math/rand"
)
 
var storage map[string][]byte // filename -> file contents as []byte
var inProcess map[string][]byte // recieved tid -> file contents

/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}


func handleDataPacket(p packet.DataPacket, conn *net.UDPConn, addr *net.UDPAddr){
	
	/*determine which file this packet is for
	  process this packet into the total bytes for the file
	  if file is done:
	      send done ack
	      move from inProcess to storage
	  if file is not done: 
	  	  send recieved cur datapacket ack
	*/

	//determine which file this packet is for
	//this is done via header. not done with header code. thus assume for inProcess[0] for now.
	
	//process this packet into the total bytes for the file
	//if file is done:
	//    send done ack
	//    move from inProcess to storage
	//if file is not done: 
	//	  send recieved cur datapacket ack
}


func routePacket(p packet.Packet, conn *net.UDPConn, addr *net.UDPAddr){
	
	switch cur := p.(type) {
		case packet.ReadPacket:
			fmt.Println("read packet for filename:",string(cur.FileName))
		case packet.WritePacket:
			fmt.Println("write packet for filename:",string(cur.FileName))
			startStorage(cur, conn, addr)
		case packet.DataPacket:
			fmt.Println("data packet containing:",string(cur.Data))
			handleDataPacket(cur)
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

func startStorage(p packet.WritePacket, conn *net.UDPConn, addr *net.UDPAddr){
	tid := rand.Uint32()
	fmt.Println(tid, "to be used soon once we have headers...")
	ackPacket := packet.NewAckPacket(uint16(0)) //0 for 0th blocknum
	_, err := conn.WriteToUDP(ackPacket.ToBytes(), addr)
	CheckError(err)
	
	inProcess = append()
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
		go routePacket(p, ServerConn, addr)
        CheckError(err)
    }
}



 
func main() {
    serve()
 
}
