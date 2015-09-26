package packet
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
var ErrorCodes map[int]string

func Init(){
	ErrorCodes = map[int]string{
	0 : "Not defined, see error message (if any).",
	1 : "File not found.",
    2 : "Access violation.",
    3 : "Disk full or allocation exceeded.",
    4 : "Illegal TFTP operation.",
    5 : "Unknown transfer ID.",
    6 : "File already exists.",
    7 : "No such user.",
}
	}
	
/* how do headers look?
 * LOCAL: you choose. 
 * INTERNET: 
 * DATAGRAM:source port, dest port, length, checksum in order. 16 bits each.
 * TFTP:
 */

//WHEN WORKING WITH HEADERS, can use anonymous field to do subclassing. 
//type Fields map[string]string
//headers map[int]Fields

type ReadWritePacket struct{
  PacketType string
  FileName string
  Mode string
}

type DataPacket struct{
  PacketType string
  BlockNum uint16
  Data []byte 
}

type AckPacket struct{
  PacketType string
  BlockNum uint16
}

type ErrorPacket struct{
  PacketType string
  ErrorCode string
  ErrMsg string
}

func ReadWritePacketToBytes(p ReadWritePacket) []byte{
	var buffer bytes.Buffer
    //buffer.WriteString(headersToString(p.headers))
    buffer.WriteString(p.PacketType)
	buffer.WriteString(p.FileName)
	buffer.WriteByte(byte(0))
	buffer.WriteString(p.Mode)
	buffer.WriteByte(byte(0))

	return buffer.Bytes()
}


func DataPacketToBytes(p DataPacket) []byte{
	var buffer bytes.Buffer
    //buffer.WriteString(headersToString(p.headers))
    buffer.WriteString(p.PacketType)
	buffer.WriteByte(byte(p.BlockNum))
	buffer.Write(p.Data)
	
	return buffer.Bytes()
}

func AckPacketToBytes(p AckPacket) []byte{
	var buffer bytes.Buffer
    //buffer.WriteString(headersToString(p.headers))
    buffer.WriteString(p.PacketType)
	buffer.WriteByte(byte(p.BlockNum))
	
	return buffer.Bytes()
}

func ErrorPacketToBytes(p ErrorPacket) []byte{
	var buffer bytes.Buffer
    //buffer.WriteString(headersToString(p.headers))
    buffer.WriteString(p.PacketType)
	buffer.WriteString(p.ErrorCode)
	buffer.WriteString(p.ErrMsg)
	buffer.Write([]byte{0})
	
	return buffer.Bytes()
}
//func fieldsToBytes(f* Fields)[]byte{
	//var buffer bytes.Buffer
	
	//for k,v := range f {
		
	//}
	//return out
//}

//func headersToString(map[int]Fields) []byte{
	//var buffer bytes.Buffer
	
	//if fields, ok := p.headers[LOCAL]; ok {
		//buffer.WriteString(fieldsToString(fields))
	//}
	
	//buffer.WriteString(fieldsToString(p.headers[INTERNET]))
	//buffer.WriteString(fieldsToString(p.headers[DATAGRAM]))
	//buffer.WriteString(fieldsToString(p.headers[TFTP]))
	
	////if packetType == RRQ || packetType == WRQ {
		////buffer.WriteString(packetType,filename, byte(0), mode, byte(0))
	////}
//}


