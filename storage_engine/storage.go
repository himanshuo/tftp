package storage_engine
import (
	"crypto/md5"
	"fmt"
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

func NewStorage() *Storage{
	a := map[string][]byte{}
	b := map[string][]string{}
	return &Storage{a,b}
}

func (s Storage) Add(f File){
	s.mapper[f.Filename] = make([]string,0)
	for i:=0; i<len(f.Data); i += BLOCKSIZE{
		
		var curBlock []byte
		if(i+BLOCKSIZE < len(f.Data)){
			curBlock = f.Data[i:i+BLOCKSIZE]
		} else {
			curBlock = f.Data[i:]
		}
		
		
		b := md5.Sum(curBlock)
		curBlockId := string(b[:])
		if _, ok := s.blocks[curBlockId]; !ok {
			s.blocks[curBlockId] = curBlock
		}
		s.mapper[f.Filename] = append(s.mapper[f.Filename], curBlockId) 
	}
}

func (s Storage) Get(filename string) File {
	if blocks_ids, ok := s.mapper[filename]; ok {
		f := File{filename, make([]byte,0), }
		for _,block_id := range blocks_ids{
			f.Data = append(f.Data, s.blocks[block_id]...)
		}
		return f
	} else{
		panic(fmt.Sprintf("%s not in storage", filename))
	}
}

func (s Storage) NumBlocks() int{
	return len(s.blocks)
}
