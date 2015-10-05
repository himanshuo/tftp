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

//func generateTid() uint16{
	//return uint16(rand.Uint32())	
//}

//func handleDataPacket(p packet.DataPacket, conn *net.UDPConn, addr *net.UDPAddr){
	
	///*
	  //determine which file this packet is for
	  //process this packet into the total bytes for the file
	  //if file is done:
	      //send done ack
	      //move from inProcess to storage
	  //if file is not done: 
	  	  //send recieved cur datapacket ack
	//*/

	////determine which file this packet is for
	////this is done via header. not done with header code. thus assume for inProcess[tid=uint16(0)] for now.
	//clientTid := p.Source
	//sourceTid := p.Dest
	
	//file := inProcess[clientTid]
	////process this packet into the total bytes for the file
	//file.data = append(file.data, p.Data...)
	
	////create ackpacket response
	//ackPacket := packet.NewAckPacket(p.BlockNum, sourceTid, clientTid)
	
	////if last packet 
	//if len(p.Data) < MAXDATASIZE{
		////move completed file to storage
		//storage[file.filename] = file
		////remove completed file from inprocess
		//delete(inProcess, clientTid)
	//} 
	
	////send ackpacket response
	//_, err := conn.WriteToUDP(ackPacket.ToBytes(), addr)
	//CheckError(err)
	
//}

//func handleAckPacket(p packet.AckPacket, conn *net.UDPConn, addr *net.UDPAddr){
	
	///*determine which file this packet is for via tid
	  //blocknum = ackpacket.blocknum+1
	  //get the appropriate block for the file based on new blocknum
	  //create a datapacket
	  //send datapacket 
	//*/

	////determine which file this packet is for via tid
	////this is done via header. not done with header code. thus assume for inProcess[tid=uint16(0)] for now.
	//clientTid := p.Source
	//serverTid := p.Dest
	//file := inProcess[clientTid]
	//blockNum := p.BlockNum + uint16(1)
	
	////get the appropriate block for the file based on new blocknum
	//data := file.data[blockNum*MAXDATASIZE:blockNum*MAXDATASIZE+MAXDATASIZE]
	
	
	////create and send datapacket
	//dataPacket := packet.NewDataPacket(blockNum, data, serverTid, clientTid)  
	//_, err := conn.WriteToUDP(dataPacket.ToBytes(), addr)
	//CheckError(err)
//}


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

//func startStorageProcess(p packet.WritePacket, conn *net.UDPConn, addr *net.UDPAddr){
	//tid := generateTid()
	//fmt.Println(tid, "to be used soon once we have headers...")
	//ackPacket := packet.NewAckPacket(uint16(0), tid, p.Source) //0 for 0th blocknum
	//_, err := conn.WriteToUDP(ackPacket.ToBytes(), addr)
	//CheckError(err)
	
	//inProcess[uint16(0)] = File{filename:p.FileName}
//}

//func startReadProcess(p packet.ReadPacket, conn *net.UDPConn, addr *net.UDPAddr){
	////supposed to get tid from packet. for now, assume, tid=uint16(0)
	//file := storage[p.FileName]
	//dataPacket := packet.NewDataPacket(uint16(0), file.data, p.Dest, p.Source) //0 for 0th blocknum
	//_, err := conn.WriteToUDP(dataPacket.ToBytes(), addr)
	//CheckError(err)
//}


const PORT int = 10001 
func serve(){
	//TODO: initialize storage and inProcess
	
	
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
