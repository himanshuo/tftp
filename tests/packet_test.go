package tests

import (
	"github.com/himanshuo/tftp/packet"
	"testing"
	"bytes"
	"flag"
	"os"
	"reflect"
	//"fmt"
)

var ALLBYTES  = []byte{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,
	19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,
	41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,
	63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,
	85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,
	105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,
	122,123,124,125,126,127,128,129,130,131,132,133,134,135,136,137,138,
	139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,
	156,157,158,159,160,161,162,163,164,165,166,167,0,168,169,170,171,172, //note, there is a \x00 in the middle of the file.
	173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,
	190,191,192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,
	207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,
	224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,
	241,242,243,244,245,246,247,248,249,250,251,252,253,254,255}

func TestMain(m *testing.M) {
	//var buffer bytes.Buffer
	//for i := 0; i <= 256; i++ {
		//buffer.WriteString(string(byte(i)))
	//}
	//ALLBYTES = buffer.Bytes()
	
	flag.Parse()
	os.Exit(m.Run())
}

func stringToBytes(str string) []byte{
 var buffer bytes.Buffer
 buffer.WriteString(str)
 return buffer.Bytes()
}

func nullTerminatedStringFromBytes(in []byte) string{
	n := bytes.IndexByte(in, byte(0))
	return string(in[:n])
}


//test ReadPacket conversion to bytes
var testReadPacket = []struct{
	packetVersion packet.ReadPacket
	bitVersion []byte
	err bool
}{
	{packet.NewReadPacket("HI"),stringToBytes("\x00"+string(packet.RRQ)+"HI"+"\x00"+packet.OCTECT+"\x00"), false}, //basic read
	{packet.NewReadPacket(""),stringToBytes("\x00"+string(packet.RRQ)+""+"\x00"+packet.OCTECT+"\x00"), false}, //empty read
	{packet.NewReadPacket(nullTerminatedStringFromBytes(ALLBYTES)),stringToBytes("\x00"+string(packet.RRQ)+nullTerminatedStringFromBytes(ALLBYTES)+"\x00"+packet.OCTECT+"\x00"), false}, //all bytes
}
func TestReadPacketToBytes(t *testing.T) {
	for i, test := range testReadPacket {
		ret := test.packetVersion.ToBytes()
		
		same := true
		for j,b := range test.bitVersion {
			if(ret[j] != b){
				same = false
			}
		}
		if !same {
			t.Errorf("Failed Test %d: ReadPacket.ToBytes(%v)=%v DESIRED: %v", i, test.packetVersion, ret, test.bitVersion)
		}
	}
}
func TestBytesToReadPacket(t *testing.T){
	for i, test := range testReadPacket {
		ret := packet.ToPacket(test.bitVersion)
		if(!reflect.DeepEqual(ret,test.packetVersion)){
			t.Errorf("Failed Test %d: ReadPacket: ToPacket(%v)=%v DESIRED: %v", i, test.bitVersion, ret, test.packetVersion)
		}	
	}
}

//test WritePacket conversion to bytes
var testWritePacket = []struct{
	packetVersion packet.WritePacket
	bitVersion []byte
	err bool
}{
	{packet.NewWritePacket("HI"),stringToBytes("\x00"+string(packet.WRQ)+"HI"+"\x00"+packet.OCTECT+"\x00"), false}, //basic read
	{packet.NewWritePacket(""),stringToBytes("\x00"+string(packet.WRQ)+""+"\x00"+packet.OCTECT+"\x00"), false}, //empty read
	{packet.NewWritePacket(nullTerminatedStringFromBytes(ALLBYTES)),stringToBytes("\x00"+string(packet.WRQ)+nullTerminatedStringFromBytes(ALLBYTES)+"\x00"+packet.OCTECT+"\x00"), false}, //all bytes
}
func TestWritePacketToBytes(t *testing.T) {
	for i, test := range testWritePacket {
		ret := test.packetVersion.ToBytes()
		
		same := true
		for j,b := range test.bitVersion {
			if(ret[j] != b){
				same = false
			}
		}
		if !same {
			t.Errorf("Failed Test %d: WritePacket.ToBytes(%v)=%v DESIRED: %v", i, test.packetVersion, ret, test.bitVersion)
		}
	}
}
func TestBytesToWritePacket(t *testing.T){
	for i, test := range testWritePacket {
		ret := packet.ToPacket(test.bitVersion)
		if(!reflect.DeepEqual(ret,test.packetVersion)){
			t.Errorf("Failed Test %d: WritePacket: ToPacket(%v)=%v DESIRED: %v", i, test.bitVersion, ret, test.packetVersion)
		}	
	}
}


//test DataPacket conversion to bytes
var testDataPacket = []struct{
	packetVersion packet.DataPacket
	bitVersion []byte
	err bool
}{
	{packet.NewDataPacket(uint16(0), []byte("a")),stringToBytes("\x00"+string(packet.DATA)+"\x00\x00"+"a"), false}, //basic read
	{packet.NewDataPacket(uint16(1), []byte("")),stringToBytes("\x00"+string(packet.DATA)+"\x00\x01"+""), false}, //empty read
	{packet.NewDataPacket(uint16(12289), ALLBYTES),stringToBytes("\x00"+string(packet.DATA)+"\x30\x01"+string(ALLBYTES)), false}, //all bytes
	{packet.NewDataPacket(uint16(0xffff), ALLBYTES),stringToBytes("\x00"+string(packet.DATA)+"\xff\xff"+string(ALLBYTES)), false}, //max blocknum
}
func TestDataPacketToBytes(t *testing.T) {
	for i, test := range testDataPacket {
		ret := test.packetVersion.ToBytes()
		
		same := true
		same = len(test.bitVersion) == len(ret)
		for j,b := range test.bitVersion {
			if(ret[j] != b){
				same = false
			}
			
			
		}
		
		if !same {
			t.Errorf("Failed Test %d: DataPacket: %v.ToBytes=%v DESIRED: %v", i, test.packetVersion, ret, test.bitVersion)
		}
	}
}
func TestBytesToDataPacket(t *testing.T){
	for i, test := range testDataPacket {
		ret := packet.ToPacket(test.bitVersion)
		if(!reflect.DeepEqual(ret,test.packetVersion)){
			t.Errorf("Failed Test %d: DataPacket: toPacket(%v)=%v DESIRED: %v", i, test.bitVersion, ret, test.packetVersion)
		}	
	}
}


//test AckPacket conversion to bytes
var testAckPacket = []struct{
	packetVersion packet.AckPacket
	bitVersion []byte
	err bool
}{
	{packet.NewAckPacket(uint16(0)),stringToBytes("\x00"+string(packet.ACK)+"\x00\x00"), false}, //0
	{packet.NewAckPacket(uint16(1)),stringToBytes("\x00"+string(packet.ACK)+"\x00\x01"), false}, //1
	{packet.NewAckPacket(uint16(12289)),stringToBytes("\x00"+string(packet.ACK)+"\x30\x01"), false}, //order of blocknum bytes is big endian
	{packet.NewAckPacket(uint16(0xffff)),stringToBytes("\x00"+string(packet.ACK)+"\xff\xff"), false}, //max blocknum
}
func TestAckPacketToBytes(t *testing.T) {
	for i, test := range testAckPacket {
		ret := test.packetVersion.ToBytes()
		
		same := true
		same = len(test.bitVersion) == len(ret)
		for j,b := range test.bitVersion {
			if(ret[j] != b){
				same = false
			}
			
			
		}
		
		if !same {
			t.Errorf("Failed Test %d: AckPacket: %v.ToBytes=%v DESIRED: %v", i, test.packetVersion, ret, test.bitVersion)
		}
	}
}
func TestBytesToAckPacket(t *testing.T){
	for i, test := range testAckPacket {
		ret := packet.ToPacket(test.bitVersion)
		if(!reflect.DeepEqual(ret,test.packetVersion)){
			t.Errorf("Failed Test %d: AckPacket: toPacket(%v)=%v DESIRED: %v", i, test.bitVersion, ret, test.packetVersion)
		}	
	}
}


//test ErrorPacket conversion to bytes
var testErrorPacket = []struct{
	packetVersion packet.ErrorPacket
	bitVersion []byte
	err bool
}{
	{packet.NewErrorPacket(uint16(0), "a"),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x00"+"a"+"\x00"), false}, //0
	{packet.NewErrorPacket(uint16(1), "a"),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x01"+"a"+"\x00"), false}, //1
	{packet.NewErrorPacket(uint16(2), "a"),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x02"+"a"+"\x00"), false}, //2
	{packet.NewErrorPacket(uint16(3), "a"),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x03"+"a"+"\x00"), false}, //3
	{packet.NewErrorPacket(uint16(4), "a"),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x04"+"a"+"\x00"), false}, //4
	{packet.NewErrorPacket(uint16(5), "a"),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x05"+"a"+"\x00"), false}, //5
	{packet.NewErrorPacket(uint16(6), "a"),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x06"+"a"+"\x00"), false}, //6
	{packet.NewErrorPacket(uint16(7), "a"),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x07"+"a"+"\x00"), false}, //7
	{packet.NewErrorPacket(uint16(2), string(ALLBYTES)),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x02"+string(ALLBYTES)+"\x00"), false}, //max bytes error string
	{packet.NewErrorPacket(uint16(2), ""),stringToBytes("\x00"+string(packet.ERROR)+"\x00\x02\x00"), false}, //empty error string
		
}
func TestErrorPacket(t *testing.T) {
	for i, test := range testErrorPacket {
		ret := test.packetVersion.ToBytes()
		
		same := true
		same = len(test.bitVersion) == len(ret)
		for j,b := range test.bitVersion {
			if(ret[j] != b){
				same = false
			}
			
			
		}
		
		if !same {
			t.Errorf("Failed Test %d: ErrorPacket: %v.ToBytes=%v DESIRED: %v", i, test.packetVersion, ret, test.bitVersion)
		}
	}
}
func TestBytesToErrorPacket(t *testing.T){
	for i, test := range testErrorPacket {
		ret := packet.ToPacket(test.bitVersion)
		if(!reflect.DeepEqual(ret,test.packetVersion)){
			t.Errorf("Failed Test %d: ErrorPacket: toPacket(%v)=%v DESIRED: %v", i, test.bitVersion, ret, test.packetVersion)
		}	
	}
}



