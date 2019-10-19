package nifdc

import (
	"fmt"
	."github.com/a97077088/grequests"
	"strings"
)


func Code(session *Session)(string,[]byte,error){
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://gc.nifdc.org.cn/code")
	r,err:=cli.Get(surl,&RequestOptions{
	})
	if err!=nil{
		return "",nil,err
	}
	ck:=r.Header.Get("Cookie")
	img:=r.Bytes()
	return ck,img,nil
}
func Login(username string,password string,ck string,validate string,session *Session)error{
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://gc.nifdc.org.cn/login")
	r,err:=cli.Post(surl,&RequestOptions{
		RequestBody:strings.NewReader("username=MTU3Mzg4ODk3MzA%3D&password=MTIzNDU2Nzg%3D&validate=z3k5&certKey=&lt=LT-2237647-hhamdL7DO7Yid2LCvdIniT032mqCWQ&execution=e3s1&_eventId=submit&UserPwd=&UserSignedData=&UserCert=&ContainerName=&strRandom="),
		Headers: map[string]string{
			"Cookie":ck,
			"Content-Type":"application/x-www-form-urlencoded",
			"Referer":"http://gc.nifdc.org.cn/login",
		},
		RedirectLimit:-1,
	})
	if err!=nil{
		return err
	}


	fmt.Println(r.Header)
	return nil
}