package main
import (
	"bytes"
)

const (
	/*packet types*/
	/*no iota because opcode is defined*/
	RRQ = "01"
	WRQ = "02"
	DATA = "03"
	ACK = "04"
	ERROR = "05"
	
	/*header types*/
	LOCAL = iota
	INTERNET
	DATAGRAM
	TFTP
	
	
	/*modes*/
	NETASCII = "NETASCII"
	OCTECT = "OCTECT"
	MAIL = "MAIL"
)
/* how do headers look?
 * LOCAL: you choose. 
 * INTERNET: 
 * DATAGRAM:source port, dest port, length, checksum in order. 16 bits each.
 * TFTP:
 * 
 * 
 */

//WHEN WORKING WITH HEADERS, can use anonymous field to do subclassing. 
//type Fields map[string]string
//headers map[int]Fields

type ReadWritePacket struct{
  PacketType string
  Filename string
  Mode string
}

type DataPacket struct{
  PacketType string
  BlockNum string
  Data []byte 
}

//type AckPacket struct{}
//type ErrorPacket struct{}

func (p Packet) ToBytes() []byte{
	var buffer bytes.Buffer
    buffer.WriteString(p.packetType)
	buffer.WriteString(headersToString(p.headers))
	
	return buffer.Bytes()
}

func fieldsToBytes(f* Fields)[]byte{
	var buffer bytes.Buffer
	
	for k,v := range f {
		
	}
	return out
}

func headersToString(map[int]Fields) []byte{
	var buffer bytes.Buffer
	
	if fields, ok := p.headers[LOCAL]; ok {
		buffer.WriteString(fieldsToString(fields))
	}
	
	buffer.WriteString(fieldsToString(p.headers[INTERNET]))
	buffer.WriteString(fieldsToString(p.headers[DATAGRAM]))
	buffer.WriteString(fieldsToString(p.headers[TFTP]))
	
	//if packetType == RRQ || packetType == WRQ {
		//buffer.WriteString(packetType,filename, byte(0), mode, byte(0))
	//}
}

//buf := new(bytes.Buffer)
	//var pi float64 = math.Pi
	//err := binary.Write(buf, binary.LittleEndian, pi)

