package packet
import (
	"bytes"
	"encoding/binary"
	//"strconv"
	//"fmt"
)

const (
	/*packet types*/
	/*no iota because opcode is defined*/
	RRQ = byte(1)
	WRQ = byte(2)
	DATA = byte(3)
	ACK = byte(4)
	ERROR = byte(5)
		
	
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
 * INTERNET: I think the datagram header and internet header are combined in udp packet design
 * DATAGRAM: source port, dest port, length, checksum in order. 16 bits each. source/dest port's are tid's. 
 * TFTP: 2 byte opcode field which indicates type of packet
 */

//WHEN WORKING WITH HEADERS, can use anonymous field to do subclassing. 
//type Fields map[string]string
//headers map[int]Fields


type Packet interface{
	ToBytes() []byte
}

type AbstractPacket struct{
	PacketType byte
	Source uint16
	Dest uint16
	//Length uint16
	//Checksum uint16
}
func (p AbstractPacket) HeadersToBytes() []byte{
	var buffer bytes.Buffer
	buffer.Write(unit16ToBytes(p.Source))
	buffer.Write(unit16ToBytes(p.Dest))
	//buffer.Write(unit16ToBytes(p.Length))
	//buffer.Write(unit16ToBytes(p.Checksum))
	return buffer.Bytes()
}
func (p AbstractPacket) ToBytes() []byte{
	var buffer bytes.Buffer
	buffer.Write(p.HeadersToBytes())
	buffer.Write([]byte{0,p.PacketType})
	return buffer.Bytes()
}


type ReadPacket struct{
  AbstractPacket
  FileName string
  Mode string
}
func NewReadPacket(filename string, source uint16, dest uint16) ReadPacket {
    return ReadPacket{AbstractPacket{RRQ, source, dest}, filename, OCTECT}
}
func (p ReadPacket) ToBytes() []byte{
  var buffer bytes.Buffer
  buffer.Write(p.AbstractPacket.ToBytes())
  buffer.WriteString(p.FileName)
  buffer.WriteByte(byte(0))
  buffer.WriteString(p.Mode)
  buffer.WriteByte(byte(0))
  return buffer.Bytes()
}

type WritePacket struct{
  AbstractPacket
  FileName string
  Mode string
}
func NewWritePacket(filename string, source uint16, dest uint16) WritePacket {
    return WritePacket{AbstractPacket{WRQ, source, dest}, filename, OCTECT}
}
func (p WritePacket) ToBytes() []byte{
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
func NewDataPacket(blockNum uint16, data []byte, source uint16, dest uint16) DataPacket{
	/*Length : Number of bytes in UDP packet, including UDP header.*/
	//length := uint16(0)
	//temp := DataPacket{AbstractPacket{DATA, source, dest, length}, blockNum, data}
	//length = len(temp.ToBytes())
	return DataPacket{AbstractPacket{DATA, source, dest}, blockNum, data}
}
func (p DataPacket) ToBytes() []byte{
  var buffer bytes.Buffer
  buffer.Write(p.AbstractPacket.ToBytes())
  buffer.Write(unit16ToBytes(p.BlockNum))
  buffer.Write(p.Data)
  return buffer.Bytes()
}

type AckPacket struct{
  AbstractPacket
  BlockNum uint16
}
func NewAckPacket(blockNum uint16, source uint16, dest uint16) AckPacket{
	return AckPacket{AbstractPacket{ACK, source, dest}, blockNum}
}
func (p AckPacket) ToBytes() []byte{
  var buffer bytes.Buffer
  buffer.Write(p.AbstractPacket.ToBytes())
  buffer.Write(unit16ToBytes(p.BlockNum))
  return buffer.Bytes()
}

type ErrorPacket struct{
  AbstractPacket
  ErrorCode uint16
  ErrMsg string
}
func NewErrorPacket(errCode uint16, errMsg string, source uint16, dest uint16) ErrorPacket{
	return ErrorPacket{AbstractPacket{ERROR, source, dest}, errCode, errMsg}
}
func (p ErrorPacket) ToBytes() []byte{
  var buffer bytes.Buffer
  buffer.Write(p.AbstractPacket.ToBytes())
  buffer.Write(unit16ToBytes(p.ErrorCode))
  buffer.WriteString(p.ErrMsg)
  buffer.WriteByte(byte(0))
  return buffer.Bytes()
}






func unit16ToBytes(u uint16) []byte{
	arr := make([]byte, 2)
    binary.BigEndian.PutUint16(arr, u)
    return arr
}

func bytesToUint16(in []byte) uint16{
	return binary.LittleEndian.Uint16(in)
}
