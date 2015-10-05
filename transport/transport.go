package transport

import (
    "fmt"
    "net"
    "os"
    "github.com/himanshuo/tftp/packet"
    "github.com/himanshuo/tftp/storage_engine"
)

const MAXDATASIZE = 516

type Tid uint16

//internally takes care of cur byte we are sending/recieving for file
//if some type of error occurs, we close the connection ??????? WHEN DO WE KNOW THIS OCCURED????

type Transport interface{
	Start()
}

type AbstractTransport struct {
	file storage_engine.File
	//serverTid Tid
	//clientTid Tid
	clientAddr *net.UDPAddr
	conn *net.UDPConn
	blocknum uint16
}

type ReadTransport struct{
	AbstractTransport
}
type WriteTransport struct{
	AbstractTransport
}

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

//todo: can remove all the extra stuff built on top of UDPAddr to keep client and server TID's
func NewReadTransport(p packet.ReadPacket, clientAddr *net.UDPAddr) Transport{
	//returns a transport that sets init vals
	
    ServerAddr,err := net.ResolveUDPAddr("udp",":0")
    CheckError(err)
	
	//I am allowing the Golang implementation of UDP to handle choosing a random port that is available 
    //serverTid := Tid(ServerAddr.Port)     
    
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)

	file := storage_engine.Get(p.FileName)
    return ReadTransport{
		AbstractTransport{
			file: file, 
			clientAddr: clientAddr,  
			conn: ServerConn,
			blocknum: uint16(0),
			},
		}
}

func NewWriteTransport(p packet.WritePacket, clientAddr *net.UDPAddr) Transport{
	//returns a transport that sets init vals
	
    ServerAddr,err := net.ResolveUDPAddr("udp",":0")
    CheckError(err)
	
	//I am allowing the Golang implementation of UDP to handle choosing a random port that is available 
    //serverTid := Tid(ServerAddr.Port)     
    
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    

	file := storage_engine.File{p.FileName, make([]byte,0)}
    return WriteTransport{ 
		AbstractTransport{
			file:file, 
			clientAddr:clientAddr, 
			conn:ServerConn,
			blocknum: uint16(0),
			}, 
	}
}

func (t ReadTransport) Start(){
	buf := make([]byte, 516)
	defer t.conn.Close()
	for{
		//send data packet
        done := t.sendDataPacket()
        //if err {
			//fmt.Println("send data packet failed. do something.... ALSO, send proper error response. NOT BOOL.")
			//break
		//}
		
		//verify
        n,_,err := t.conn.ReadFromUDP(buf)
		CheckError(err)
		p := packet.ToPacket(buf[0:n])
		fmt.Printf("recieved packet from client for readtransport: %v\n",p)
		
		
		if done{
			break
		}
	}
	
}

func (t WriteTransport) Start(){
	buf := make([]byte, 516)
	defer t.conn.Close()
	for{
		//let client know that we previous write/data packet was all good
		
		t.sendAckPacket()
		
		
		//get data pack
		
		n,_,err := t.conn.ReadFromUDP(buf)
		
		
		CheckError(err)
		p:=packet.ToPacket(buf[0:n])
		switch cur := p.(type) {
			case packet.DataPacket:
				t.file.Data = append(t.file.Data, cur.Data...)
				
			default:
				fmt.Println("ERROR. should have gotten data packet from user. got something else.")
		}
		
		
		fmt.Println("n is", n)
		if n < MAXDATASIZE {
			fmt.Println("right before calling sendackpacket again, t.blocknum is", t.blocknum)
			t.sendAckPacket()
			break
		}
		
	}
	
	t.storeFile()
	
	
}

func (t *WriteTransport) sendAckPacket(){
	
	fmt.Printf( "sending ack packet (blocknum = %d) back to client with addr %v\n",t.blocknum, t.clientAddr )
	ackPacket := packet.NewAckPacket(t.blocknum)
	_, err := t.conn.WriteToUDP(ackPacket.ToBytes(), t.clientAddr)
	CheckError(err)
	
	
	t.blocknum = t.blocknum + uint16(1)
	fmt.Println("just increased t.blocknum to be", t.blocknum)
		
}

func (t *ReadTransport) sendDataPacket() bool{
	done := false
	//error := false
	
	//determine which file this packet is for via tid
	//this is done via header. not done with header code. thus assume for inProcess[tid=uint16(0)] for now.
	t.blocknum = t.blocknum + uint16(1)
	
	//get the appropriate block for the file based on new blocknum
	var data []byte
	start := (int(t.blocknum)-1) * MAXDATASIZE
	end := start + MAXDATASIZE
	if end < len(t.file.Data){
		data = t.file.Data[start:end]
	} else {

		data = t.file.Data[start:]
		done = true
	} 
		
	//create and send datapacket
	dataPacket := packet.NewDataPacket(t.blocknum, data)  
	_, err := t.conn.WriteToUDP(dataPacket.ToBytes(), t.clientAddr)
	CheckError(err)
	
	return done
}

//func (t Transport) sendErrorPacket()
func (t *WriteTransport) storeFile(){
	//when transport is done, we store the data into our storageEngine
	storage_engine.Put(t.file)
}

