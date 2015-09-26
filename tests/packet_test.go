package tests

import (
	"github.com/himanshuo/tftp/packet"
	"testing"
	"bytes"
)

var ALLBYTES string
func init(){
	var buffer bytes.Buffer
	for i := 1; i <= 256; i++ {
		buffer.WriteString(string(byte(i)))
	}
    
	ALLBYTES = buffer.String()

}


func stringToBytes(str string) []byte{
 var buffer bytes.Buffer
 buffer.WriteString(str)
 return buffer.Bytes()
 	
}

var testReadWritePacketToBytes = []struct{
	in packet.ReadWritePacket
	expected []byte
}{
	{packet.ReadWritePacket{packet.RRQ, "hi", packet.OCTECT},stringToBytes(packet.RRQ+"hi"+"\x00"+packet.OCTECT+"\x00")}, //basic read
	{packet.ReadWritePacket{packet.RRQ, "", packet.OCTECT},stringToBytes(packet.RRQ+""+"\x00"+packet.OCTECT+"\x00")}, //empty read
	{packet.ReadWritePacket{packet.RRQ, ALLBYTES, packet.OCTECT},stringToBytes(packet.RRQ+ALLBYTES+"\x00"+packet.OCTECT+"\x00")}, //all bytes
	{packet.ReadWritePacket{packet.WRQ, "hi", packet.OCTECT},stringToBytes(packet.WRQ+"hi"+"\x00"+packet.OCTECT+"\x00")}, //basic write
	{packet.ReadWritePacket{packet.WRQ, "", packet.OCTECT},stringToBytes(packet.WRQ+""+"\x00"+packet.OCTECT+"\x00")}, //empty write
	{packet.ReadWritePacket{packet.WRQ, ALLBYTES, packet.OCTECT},stringToBytes(packet.WRQ+ALLBYTES+"\x00"+packet.OCTECT+"\x00")}, //all write			
}

func TestReadWritePacketToBytes(t *testing.T) {
	for i, test := range testReadWritePacketToBytes {
		ret := packet.ReadWritePacketToBytes(test.in)
		
		same := true
		for j,b := range test.expected {
			if(ret[j] != b){
				same = false
			}
		}
		if !same {
			t.Errorf("Failed Test %d: readWritePacketToBytes(%v)=%v DESIRED: %v", i, test.in, ret, test.expected)
		}
	}
}

var testDataPacketToBytes = []struct{
	in packet.DataPacket
	expected []byte
}{
	{packet.DataPacket{packet.DATA, uint16(0), []byte("a")},stringToBytes(packet.DATA+"0"+"a")}, //basic data
	{packet.DataPacket{packet.DATA, uint16(1), []byte("")},stringToBytes(packet.DATA+"1"+"")}, //empty data
	{packet.DataPacket{packet.DATA, uint16(3), []byte(ALLBYTES)},stringToBytes(packet.DATA+"3"+ALLBYTES)}, //all bytes
	{packet.DataPacket{packet.DATA, uint16(0xff), []byte(ALLBYTES)},stringToBytes(packet.WRQ+"\xff"+ALLBYTES)}, //max block
}

func TestDataPacketToBytes(t *testing.T) {
	for i, test := range testDataPacketToBytes {
		var ret []byte
		ret = packet.DataPacketToBytes(test.in)
	
		
		same := true
		for j,b := range test.expected {
			if(ret[j] != b){
				same = false
			}
		}
		if !same {
			t.Errorf("Failed Test %d: readWritePacketToBytes(%v)=%v DESIRED: %v", i, test.in, ret, test.expected)
		}
	}
}






	


//var testFieldsToBytes = []struct{
	//in tftp.Fields
	//expected []byte
//}{
	////simple
	//{
		//Fields{
			//"a":"a",
			//"f2":"f2"
		//},
		//[]byte{"a","f2"} 
		
	//}
	
//}


//var testToBytes = []struct {
	//packetType  int
	//headers map[int]Fields
	//payload map[string][]byte 
	//option bool
	//ok  bool
//}{  

//}







