package nifdc

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	ck,img,err:=Code(nil)
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(ck)
	fmt.Println(img)

	f := bufio.NewReader(os.Stdin)
	s,err:=f.ReadString('\n')

	err=Login("a","33333333",ck,s,nil)
	fmt.Println(err)
}
