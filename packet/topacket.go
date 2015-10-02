package packet
import (
	"bytes"
	"encoding/binary"
	//"fmt"
)
const HeaderOffset = 8 // in bytes

func ToPacket(recievedBytes []byte) Packet{
	var p Packet
	
	source := bytesToUint16(recievedBytes[0:2])
	dest := bytesToUint16(recievedBytes[2:4])
	//length := bytesToUint16(recievedBytes[4:6])
	//checksum := bytesToUint16(recievedBytes[6:8])
	
	
	i := HeaderOffset+1 //i is cur byte we are looking at
	opcode := recievedBytes[i]
	i = i + 1
	
	
	switch opcode{
		case byte(1):
			n := bytes.IndexByte(recievedBytes[i:], byte(0)) + i
			fileName := string(recievedBytes[i:n])
			//fmt.Println(recievedBytes)
			//fmt.Println(n)
			//fmt.Println(fileName)
			//fmt.Println(i)
			return NewReadPacket(fileName, source, dest)
		case byte(2):
			n := bytes.IndexByte(recievedBytes[i:], byte(0)) + i
			fileName := string(recievedBytes[i:n])
			return NewWritePacket(fileName, source, dest)
		case byte(3):
			blockNum := binary.BigEndian.Uint16(recievedBytes[i:i+2])
			i = i+2
			data := recievedBytes[i:]
		    return NewDataPacket(blockNum, data, source, dest)
		case byte(4):
			blockNum := binary.BigEndian.Uint16(recievedBytes[i:i+2])
			return NewAckPacket(blockNum, source, dest)
		case byte(5):
			errorCode := binary.BigEndian.Uint16(recievedBytes[i:i+2])
		    i = i+2
		    //n := bytes.IndexByte(recievedBytes[i:], byte(0)) + i
		    n := len(recievedBytes)-1
		    errMsg := string(recievedBytes[i:n])
		    //fmt.Printf("%d:%d -> %v\n", i,n,errMsg)
		    return NewErrorPacket(errorCode, errMsg, source, dest)
	}
	return p
	
}
