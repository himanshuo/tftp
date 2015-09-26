package packet
import (
	"bytes"
	"encoding/binary"
)

const (
	/*packet types*/
	/*no iota because opcode is defined*/
	RRQ = 1
	WRQ = 2
	DATA = 3
	ACK = 4
	ERROR = 5
	
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

var ErrorCodes map[uint16]string

func Init(){
	ErrorCodes = map[uint16]string{
		uint16(0) : "Not defined, see error message (if any).",
		uint16(1) : "File not found.",
		uint16(2) : "Access violation.",
		uint16(3) : "Disk full or allocation exceeded.",
		uint16(4) : "Illegal TFTP operation.",
		uint16(5) : "Unknown transfer ID.",
		uint16(6) : "File already exists.",
		uint16(7) : "No such user.",
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

//todo:
//1) init for all structs
//2) ToBytes for all structs
//3) fromBytes for all structs

type Packet interface{
	ToBytes() []byte
}

type AbstractPacket struct{
	PacketType int
}
func (p AbstractPacket) HeadersToBytes() []byte{
	//todo: implement!
	return []byte{}
}
func (p AbstractPacket) ToBytes() []byte{
	var buffer bytes.Buffer
	buffer.Write(p.HeadersToBytes())
	buffer.Write([]byte{0,byte(p.PacketType)})
	return buffer.Bytes()
}


type ReadWritePacket struct{
  AbstractPacket
  FileName string
  Mode string
}
func (p ReadWritePacket) ToBytes() []byte{
  var buffer bytes.Buffer
  buffer.Write(p.AbstractPacket.ToBytes())
  buffer.WriteString(p.FileName)
  buffer.WriteByte(byte(0))
  buffer.WriteString(p.Mode)
  buffer.WriteByte(byte(0))
  return buffer.Bytes()
}

type DataPacket struct{
  AbstractPacket
  BlockNum uint16
  Data []byte 
}
func (p DataPacket) ToBytes() []byte{
  var buffer bytes.Buffer
  buffer.Write(p.AbstractPacket.ToBytes())
  blockNumAsTwoBytes := make([]byte, 2)
  binary.LittleEndian.PutUint16(blockNumAsTwoBytes, uint16(p.BlockNum))
  buffer.Write(blockNumAsTwoBytes)
  buffer.Write(p.Data)
  return buffer.Bytes()
}

type AckPacket struct{
  AbstractPacket
  BlockNum uint16
}
func (p AckPacket) ToBytes() []byte{
  var buffer bytes.Buffer
  buffer.Write(p.AbstractPacket.ToBytes())
  blockNumAsTwoBytes := make([]byte, 2)
  binary.LittleEndian.PutUint16(blockNumAsTwoBytes, uint16(p.BlockNum))
  buffer.Write(blockNumAsTwoBytes)
  return buffer.Bytes()
}

type ErrorPacket struct{
  AbstractPacket
  ErrorCode uint16
  ErrMsg string
}
func (p ErrorPacket) ToBytes() []byte{
  var buffer bytes.Buffer
  buffer.Write(p.AbstractPacket.ToBytes())
  blockNumAsTwoBytes := make([]byte, 2)
  binary.LittleEndian.PutUint16(blockNumAsTwoBytes, uint16(p.ErrorCode))
  buffer.Write(blockNumAsTwoBytes)
  buffer.WriteString(ErrorCodes[p.ErrorCode])
  buffer.WriteByte(byte(0))
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
