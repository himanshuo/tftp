package tests

import (
	"github.com/himanshuo/tftp/storage_engine"
	"testing"
	"fmt"
	"math/rand"
	//"flag"
	//"os"
)


func sameFile(a, b storage_engine.File) bool {
  if &a == &b {
    return true
  }
  
  if a.Filename != b.Filename {
    return false
  }
  
  if len(a.Data) != len(b.Data){
    return false
  }

	for i := range a.Data {		
		if (a.Data[i] != b.Data[i]){
			return false
		}
	
	}  
  return true
}





func TestBasicPutAndGetSameFile(t *testing.T) {
	storage_engine.Reset()
	a := storage_engine.File{"myfile.txt", []byte{0,1,2,3,4,5}}
	storage_engine.Put(a)
	if !sameFile(storage_engine.Get("myfile.txt"), a){
		t.Errorf("TestBasicAddAndGetSameFile %v != %v", storage_engine.Get("myfile.txt"), a)
	} 
}


func TestLargePutAndGetSameFile(t *testing.T) {
	storage_engine.Reset()
	a := storage_engine.File{"myfile.txt", make([]byte,0)}
	for i:=0;i<99999999;i++{
		a.Data = append(a.Data, byte(i))
	}
	storage_engine.Put(a)
	if !sameFile(storage_engine.Get("myfile.txt"), a){
		t.Errorf("TestLargeAddAndGetSameFile %v != %v", storage_engine.Get("myfile.txt"), a)
	} 
}

func TestDeduplicateBlocks(t *testing.T) {
	storage_engine.Reset()
	a := storage_engine.File{"a.txt", make([]byte,0)}
	b := storage_engine.File{"b.txt", make([]byte,0)}
	
	for i:=0;i<99999999;i++{
		a.Data = append(a.Data, byte(i))
		b.Data = append(b.Data, byte(i))
	}
	
	storage_engine.Put(a)
	previous := storage_engine.NumBlocks();
	storage_engine.Put(b)
	
	if !sameFile(storage_engine.Get("a.txt"), a){
		t.Errorf("TestDeduplicateBlocks %v != %v", storage_engine.Get("a.txt"), a)
	} 
	if !sameFile(storage_engine.Get("b.txt"), b){
		t.Errorf("TestDeduplicateBlocks %v != %v", storage_engine.Get("b.txt"), b)
	} 
	if previous != storage_engine.NumBlocks(){
		t.Errorf("TestDeduplicateBlocks %v != %v", previous, storage_engine.NumBlocks())
	}
	
	 
}


func TestPartialDeduplicateBlocks(t *testing.T) {
	storage_engine.Reset()
	
	a := storage_engine.File{"a.txt", make([]byte,0)}
	b := storage_engine.File{"b.txt", make([]byte,0)}
	temp := make([]byte, 0)
	for i:=0;i < storage_engine.BLOCKSIZE*5; i++{
		temp = append(temp, byte(rand.Int()))
	}
	a.Data = temp
	for i:=0; i < storage_engine.BLOCKSIZE*5; i++{
		temp = append(temp, byte(rand.Int()))
	}
	b.Data = temp
	fmt.Println(len(a.Data))
	fmt.Println(len(b.Data))


	storage_engine.Put(a)
	previous := storage_engine.NumBlocks();
	storage_engine.Put(b)
	
	if !sameFile(storage_engine.Get("a.txt"), a){
		t.Errorf("TestPartialDeduplicateBlocks %v != %v", storage_engine.Get("a.txt"), a)
	} 
	if !sameFile(storage_engine.Get("b.txt"), b){
		t.Errorf("TestPartialDeduplicateBlocks %v != %v", storage_engine.Get("b.txt"), b)
	} 
	if previous*2 != storage_engine.NumBlocks(){
		
		t.Errorf("TestPartialDeduplicateBlocks %v != %v", previous*2, storage_engine.NumBlocks())
	}
	
}

func TestOverride(t *testing.T){}


//func TestMain(m *testing.M) {
	//storage_engine.Reset()
	//flag.Parse()
	//os.Exit(m.Run())
//}

