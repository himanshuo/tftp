package main
import (
    "fmt"
    "net"
    //"time"
    //"strconv"
    "github.com/himanshuo/tftp/packet"
    "flag"
    "io/ioutil"
)
 
 
var filepath = flag.String("filepath", "", "the path (linux) to the file where you want to upload from")

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func SendToServer(filepath string){
	ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
    defer Conn.Close()
    
    
    msg,err := ioutil.ReadFile(filepath)
    if err!=nil{
		fmt.Println(filepath)
		fmt.Println(err)
	} else {
		dataPacket := packet.NewDataPacket(uint16(1), []byte(msg))
			
		_,err = Conn.Write(dataPacket.ToBytes())
		if err != nil {
			fmt.Println(msg, err)
		}
	}   
}




func main() {
	flag.Parse()
	if(*filepath == ""){
		*filepath = flag.Arg(0)
	}
	
	SendToServer(*filepath)
 
}
