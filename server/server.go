package main
 
import (
    "fmt"
    "net"
    "os"
    "github.com/himanshuo/tftp/packet"
    "math/rand"
)
const MAXDATASIZE = 512
type File struct {
	filename string
	data []byte
	checksum string
	length int
}
var storage map[string]File // filename -> File
var inProcess map[uint16]File // recieved tid -> File

/*
 * to deduplicate:
 * 
 *  File: filename, checksum, length, []512 data byte pointers in proper order
 *  Storage: hash_of_512_byte -> 512 byte
 * 
 *  to add a new file to storage, take a hash of each of its 512 byte portions.
 *  for each hash, check if it exists in storage
 * 		if hash in storage: do nothing.
 * 		if hash not in storage: add hash->512bytes into storage 
 *  
 *  to get file from storage:
 *  totalFile = []byte
 *  for hash in File.512bytepointers:
 *  	totalFile = append(totalFile, Storage[hash])
 *  
 */

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
	//this is done via header. not done with header code. thus assume for inProcess[tid=uint16(0)] for now.
	file := inProcess[uint16(0)]
	//process this packet into the total bytes for the file
	file.data = append(file.data, p.Data)
	
	//create ackpacket response
	ackPacket := packet.NewAckPacket(p.BlockNum)
	
	//if last packet 
	if p.Data < MAXDATASIZE{
		//move completed file to storage
		storage[file.filename] = file
		//remove completed file from inprocess
		delete(inProcess, file.filename)
	} 
	
	//send ackpacket response
	_, err := conn.WriteToUDP(ackPacket.ToBytes(), addr)
	CheckError(err)
	
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
	tid := rand.Uint16()
	fmt.Println(tid, "to be used soon once we have headers...")
	ackPacket := packet.NewAckPacket(uint16(0)) //0 for 0th blocknum
	_, err := conn.WriteToUDP(ackPacket.ToBytes(), addr)
	CheckError(err)
	
	inProcess[uint16(0)] = File{filename:p.FileName}
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
