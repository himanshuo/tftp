package transport

import (
    "fmt"
    "net"
    //"os"
    "github.com/himanshuo/tftp/packet"
)

const MAXPACKETSIZE = 516

type Tid uint16

//internally takes care of cur byte we are sending/recieving for file
//if some type of error occurs, we close the connection ??????? WHEN DO WE KNOW THIS OCCURED????
type Transport struct {
	file File
	serverTid Tid
	clientTid Tid
	conn *net.UDPConn
}

func NewReadTransport(p packet.ReadPacket, Tid clientPort) Transport{
	//returns a transport that sets init vals
	
    ServerAddr,err := net.ResolveUDPAddr("udp",":0")
    CheckError(err)
	
	//I am allowing the Golang implementation of UDP to handle choosing a random port that is available 
    //serverTid := ServerAddr.Port     
    
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    //defer ServerConn.Close() /* at end of function, ServerConn is closed. */
	
    return Transport{file:}
}

func (t Transport) Start()
func (t Transport) sendAckPacket()
func (t Transport) sendDataPacket()
func (t Transport) sendErrorPacket()
func (t Transport) storeFile()//when transport is done, we store the data into our storageEngine

