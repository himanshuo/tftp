package tests

import (
	"github.com/himanshuo/tftp"
	"testing"
)

var testFieldsToBytes = []struct{
	in tftp.Fields
	expected []byte
}{
	//simple
	{
		Fields{
			"a":"a",
			"f2":"f2"
		},
		[]byte{"a","f2"} 
		
	}
	
}


var testToBytes = []struct {
	packetType  int
	headers map[int]Fields
	payload map[string][]byte 
	option bool
	ok  bool
}{  

}







func TestSiteDataDir(t *testing.T) {
	for i, test := range testData {
		//option = multipath
		ret, err := appdirs.SiteDataDir(test.name, test.author, test.version, test.option)
		
		// ok
		// !ok
		// ret exist
		// err exists
		// ret no exist
		// error no exist
		
		if platform == appdirs.LINUX {
			// ok/!ok, ret exist, err exist
			if ret != "" && err != nil {
				t.Errorf("#%d: Both return value and error exist: ret value and error: %v and %v",i, ret, err)
			// ok/!ok, ret not exist, err not exist
			} else if ret != "" && err != nil {
				t.Errorf("#%d: Both return value and error are nil", i)
			// ok, ret exist, err not exist
			} else if  test.ok && ret != "" && err == nil {
				shouldBe := filepath.Join("/usr/local/share/",test.name)
				shouldBe2 := filepath.Join("/usr/share/", test.name)
				if test.version != "" {
					shouldBe = filepath.Join(shouldBe, test.version)
					shouldBe2 = filepath.Join(shouldBe2, test.version)
				}
				
				if test.option {
					shouldBe =  shouldBe + ":" + shouldBe2
				}
				if ret != shouldBe {
					t.Errorf("#%d: Incorrect result. Expected:%v Got:%v", i, shouldBe, ret)
				}
			// ok, ret not exist, err exist
			} else if test.ok && ret == "" && err != nil {
				t.Errorf("#%d: result was not ok. Expected ok.", i)
			// !ok, ret exist, err not exist
			} else if !test.ok && ret != "" && err == nil {
				t.Errorf("#%d: result was ok. Expected error. Input was: {%v, %v, %v, %v}. Result was:", i, test.name, test.author, test.version, test.option, ret)
			// !ok, ret not exist, err exist
			} else if !test.ok && ret == "" && err != nil{
				//expected error, got it. all good.
			} 
	
		}
		
	
	}
}


