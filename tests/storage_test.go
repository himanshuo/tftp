package tests

import (
	"github.com/himanshuo/tftp/storage_engine"
	"testing"
	"fmt"
	"math/rand"
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





func TestBasicAddAndGetSameFile(t *testing.T) {
	a := storage_engine.File{"myfile.txt", []byte{0,1,2,3,4,5}}
	store := storage_engine.NewStorage()
	store.Add(a)
	if !sameFile(store.Get("myfile.txt"), a){
		t.Errorf("TestBasicAddAndGetSameFile %v != %v", store.Get("myfile.txt"), a)
	} 
}


func TestLargeAddAndGetSameFile(t *testing.T) {
	a := storage_engine.File{"myfile.txt", make([]byte,0)}
	for i:=0;i<99999999;i++{
		a.Data = append(a.Data, byte(i))
	}
	store := storage_engine.NewStorage()
	store.Add(a)
	if !sameFile(store.Get("myfile.txt"), a){
		t.Errorf("TestLargeAddAndGetSameFile %v != %v", store.Get("myfile.txt"), a)
	} 
}

func TestDeduplicateBlocks(t *testing.T) {
	a := storage_engine.File{"a.txt", make([]byte,0)}
	b := storage_engine.File{"b.txt", make([]byte,0)}
	
	for i:=0;i<99999999;i++{
		a.Data = append(a.Data, byte(i))
		b.Data = append(b.Data, byte(i))
	}
	store := storage_engine.NewStorage()
	store.Add(a)
	previous := store.NumBlocks();
	store.Add(b)
	
	if !sameFile(store.Get("a.txt"), a){
		t.Errorf("TestDeduplicateBlocks %v != %v", store.Get("a.txt"), a)
	} 
	if !sameFile(store.Get("b.txt"), b){
		t.Errorf("TestDeduplicateBlocks %v != %v", store.Get("b.txt"), b)
	} 
	if previous != store.NumBlocks(){
		t.Errorf("TestDeduplicateBlocks %v != %v", previous, store.NumBlocks())
	}
	
	 
}


func TestPartialDeduplicateBlocks(t *testing.T) {
	a := storage_engine.File{"a.txt", make([]byte,0)}
	b := storage_engine.File{"b.txt", make([]byte,0)}
	temp := make([]byte, 0)
	for i:=0;i<storage_engine.BLOCKSIZE*5;i++{
		temp = append(temp, byte(rand.Int()))
	}
	a.Data = temp
	for i:=0;i<storage_engine.BLOCKSIZE*5;i++{
		temp = append(temp, byte(rand.Int()))
	}
	b.Data = temp
	fmt.Println(len(a.Data))
	fmt.Println(len(b.Data))

	store := storage_engine.NewStorage()
	store.Add(a)
	previous := store.NumBlocks();
	store.Add(b)
	
	if !sameFile(store.Get("a.txt"), a){
		t.Errorf("TestPartialDeduplicateBlocks %v != %v", store.Get("a.txt"), a)
	} 
	if !sameFile(store.Get("b.txt"), b){
		t.Errorf("TestPartialDeduplicateBlocks %v != %v", store.Get("b.txt"), b)
	} 
	if previous*2 != store.NumBlocks(){
		
		t.Errorf("TestPartialDeduplicateBlocks %v != %v", previous*2, store.NumBlocks())
	}
	
}

func TestOverride(t *testing.T){}

