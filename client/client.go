package main
import (
    "fmt"
    "net"
    //"time"
    //"strconv"
    "github.com/himanshuo/tftp/packet"
    "flag"
    //"io/ioutil"
    "bufio"
    "path/filepath"
    "os"
)
 
 
var pathToFile = flag.String("pathToFile", "", "the path (linux) to the file where you want to upload from")
const MAXDATASIZE = 512

func CheckError(err error) {
    if err  != nil {
        //fmt.Println("Error: " , err)
		panic(err)
    }
}

func SendDataPacketsToServer(pathToFile string, Conn *net.UDPConn){
	//open file for reading
	f, err := os.Open(pathToFile)
	CheckError(err)
	defer f.Close()
	
	//create buffer for storing MAXDATASIZE bytes of file at a time
	buffer := make([]byte, MAXDATASIZE)
	//reader for reading file into buffer
	reader := bufio.NewReader(f)
	
	tid := uint16(0)
	//keep track of block num
	blockNum := uint16(1)
	for{
		//read 
		actualBytesRead, err := reader.Read(buffer)
		CheckError(err)
		
		//store into data packet
		dataPacket := packet.NewDataPacket(blockNum, buffer, tid, uint16(69))
		
		//send data packet
		_,err = Conn.Write(dataPacket.ToBytes())
		CheckError(err)
		
		p := make([]byte, 2048)
		_, err = bufio.NewReader(Conn).Read(p)
		retPacket := packet.ToPacket(p)
		//make sure response is good with correctly returned ack packet
		switch cur := retPacket.(type){
			case packet.AckPacket:
				fmt.Println("all good for blocknum:", cur.BlockNum)
			case packet.ErrorPacket:
				fmt.Println("error!:", cur.ErrMsg)
			default:
				fmt.Println("ummm wrong type of packet recieved to confirm data packet sent")
		}
		
		//if done reading file AND sending it, then  quit loop
		if actualBytesRead < MAXDATASIZE {
			break
		}
		//increase blockNum for next portion of packet to be sent
		blockNum = blockNum + uint16(1)
	}
}


func SendToServer(pathToFile string){
	ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
    defer Conn.Close()
    
    
    
    if err!=nil{
		fmt.Println(pathToFile)
		fmt.Println(err)
	} else {
		//send write request
		tid := uint16(0)
		filename := filepath.Base(pathToFile)
		writePacket := packet.NewWritePacket(filename, tid, uint16(69))
		_,err = Conn.Write(writePacket.ToBytes())
		CheckError(err)
		
		p := make([]byte, 2048)
		_, err = bufio.NewReader(Conn).Read(p)
		CheckError(err)
		retPacket := packet.ToPacket(p)
		
		switch resp := retPacket.(type) {
			case packet.AckPacket:
				fmt.Println("acknowledge packet for some blocknum:",resp.BlockNum)
				//all is good. start sending data packets like normal
				SendDataPacketsToServer(pathToFile,Conn)	
			case packet.ErrorPacket:
				fmt.Println("error packet with ErrMsg:",string(resp.ErrMsg))
			default:
				fmt.Println("returned packet for write request is erronous")
		} 	

		
		
			
		
	}   
}




func main() {
	flag.Parse()
	if(*pathToFile == ""){
		*pathToFile = flag.Arg(0)
	}
	
	SendToServer(*pathToFile)
 
}
