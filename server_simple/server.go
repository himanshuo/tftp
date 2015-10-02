package main
import (
    "fmt"
    "net"
    "os"
    "bytes"
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

func generateTid() uint16{
	return uint16(rand.Uint32())	
}


func startReadProcess(in []byte,conn *net.UDPConn){
	i := bytes.IndexByte(in, byte(0))
	filename := string(in[0:i])
	fmt.Println(filename)
	//file := storage[p.FileName]
	//dataPacket := packet.NewDataPacket(uint16(0), file.data, p.Dest, p.Source) //0 for 0th blocknum
	//_, err := conn.WriteToUDP(dataPacket.ToBytes(), addr)
	//CheckError(err)
	
}
func routePacket(in []byte, conn *net.UDPConn){
	
	switch in[1] {
		case byte(1):
			fmt.Println("read packet for filename:", string(in))
			startReadProcess(in[2:], conn)
		case byte(2):
			fmt.Println("write packet for filename:",string(in))
			//startStorageProcess(cur, conn, addr)
		case byte(3):
			fmt.Println("data packet containing:",string(in))
			//handleDataPacket(cur, conn, addr)
		case byte(4):
			fmt.Println("acknowledge packet for some blocknum:",string(in))
			//handleAckPacket(cur, conn, addr)
		case byte(5):
			fmt.Println("error packet with ErrMsg:",string(in))
		//case packet.AbstractPacket:
			//fmt.Println("abstract packet")
		//case packet.Packet:
			//fmt.Println("interface")
		default:
			fmt.Println("didn't match anytype")
	}
}

func serve(){
	//initialize storage and inProcess
	storage = map[string]File{}
	inProcess = map[uint16]File{}
	
	
	/* Lets prepare a address at any address at port 10001*/   
    ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
    fmt.Println(ServerAddr.String())
    CheckError(err)
 
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close() /* at end of function, ServerConn is closed. */
	
	buf := make([]byte, 516)
 
    for {
        n,err := ServerConn.Read(buf)
        fmt.Println(buf)
        fmt.Println(err)
        //fmt.Println(addr)
        fmt.Println(n)
		go routePacket(buf, ServerConn)
        CheckError(err)
    }
}



 
func main() {
    serve()
}
