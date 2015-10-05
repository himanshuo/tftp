package storage_engine
import (
	"crypto/md5"
	"fmt"
	"errors"
	
)

const BLOCKSIZE int = 512
/*
 * to deduplicate:
 * 
 *  File: filename, checksum, length, []512 data byte pointers in proper order
 *  Storage: hash_of_512_byte -> 512 byte
 * 
 *  to add a new file to storage, take a hash of each of its 512 byte portions.
 *  for each hash, check if it exists in storage
 * 		if hash in storage: do nothing.
 * 		if hash not in storage: add hash->512bytes into storage 
 *  
 *  to get file from storage:
 *  totalFile = []byte
 *  for hash in File.512bytepointers:
 *  	totalFile = append(totalFile, Storage[hash])
 *  
 */


type File struct {
	Filename string
	Data []byte
}

type Storage struct{
	blocks map[string][]byte // checksum_string to represent block_id -> bytes for that block 
	mapper map[string][]string // filename -> list of checksum_strings to represent block_ids
}

/*singleton design pattern in order to make sure we only use 1 Storage Engine */
var storage Storage

func init(){
	fmt.Println("creating new storage instance")
	a := map[string][]byte{}
	b := map[string][]string{}
	storage = Storage{a,b}
}

//func StorageInstance() *Storage{
	//fmt.Println("giving you already created storage instance")
	//return &storage
//}

func Put(f File){
	storage.mapper[f.Filename] = make([]string,0)
	for i:=0; i<len(f.Data); i += BLOCKSIZE{
		
		var curBlock []byte
		if(i+BLOCKSIZE < len(f.Data)){
			curBlock = f.Data[i:i+BLOCKSIZE]
		} else {
			curBlock = f.Data[i:]
		}
		
		
		b := md5.Sum(curBlock)
		curBlockId := string(b[:])
		if _, ok := storage.blocks[curBlockId]; !ok {
			storage.blocks[curBlockId] = curBlock
		}
		storage.mapper[f.Filename] = append(storage.mapper[f.Filename], curBlockId) 
	}
}

func Get(filename string) (File,error) {
	f := File{filename, make([]byte,0), }
	if blocks_ids, ok := storage.mapper[filename]; ok {	
		for _,block_id := range blocks_ids{
			f.Data = append(f.Data, storage.blocks[block_id]...)
		}
		return f, nil
	} else{
		return f, errors.New(fmt.Sprintf("%s not in storage", filename))
	}
}

func NumBlocks() int{
	return len(storage.blocks)
}

func Reset() {
	fmt.Println("resetting storage")
	a := map[string][]byte{}
	b := map[string][]string{}
	storage = Storage{a,b}
}
