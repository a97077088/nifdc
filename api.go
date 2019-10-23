package nifdc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	. "github.com/a97077088/grequests"
	"github.com/a97077088/nettool"
	"regexp"
	"strings"
)





func Code(ck string,session *Session)([]byte,error){
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://gc.nifdc.org.cn/code")
	r,err:=cli.Get(surl,&RequestOptions{
		Headers: map[string]string{
			"cookie":ck,
		},
	})
	if err!=nil{
		return nil,err
	}
	img:=r.Bytes()
	return img,nil
}
func Findtextval(s string,name string)string{
	re:=regexp.MustCompile(fmt.Sprintf("<input type=\"hidden\" name=\"%s\" value=\"(.*)?\"/>",name))
	fds:=re.FindStringSubmatch(s)
	if len(fds)<2{
		return ""
	}
	return fds[1]
}
func Findval(s string,sre string)string{
	re:=regexp.MustCompile(sre)
	fds:=re.FindStringSubmatch(s)
	if len(fds)<2{
		return ""
	}
	return fds[1]
}
func Findvaln(s string,sre string,idx int)string{
	re:=regexp.MustCompile(sre)
	fds:=re.FindStringSubmatch(s)
	return fds[idx]
}
func Initck(session *Session)(string,string,string,error){
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://gc.nifdc.org.cn/login")
	r,err:=cli.Get(surl,&RequestOptions{
		RedirectLimit:-1,
	})
	if err!=nil{
		return "","","",err
	}
	rck:=r.Header.Get("Set-Cookie")
	sbd:=r.String()
	rt,_:=goquery.NewDocumentFromReader(strings.NewReader(sbd))
	lt:=rt.Find("input[name=lt]").AttrOr("value","")
	execution:=rt.Find("input[name=execution]").AttrOr("value","")
	return rck,lt,execution,nil
}
//登录
func Login(username string,password string,ck string,lt string,execution string,session *Session)(string,error){
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://gc.nifdc.org.cn/login")
	username=base64.StdEncoding.EncodeToString([]byte(username))
	password=base64.StdEncoding.EncodeToString([]byte(password))
	spd:=fmt.Sprintf("username=%s&password=%s&validate=&certKey=&lt=%s&execution=%s&_eventId=submit&UserPwd=&UserSignedData=&UserCert=&ContainerName=&strRandom=",username,password,lt,execution)
	r,err:=cli.Post(surl,&RequestOptions{
		RequestBody:strings.NewReader(spd),
		Headers: map[string]string{
			"Cookie":ck,
			"Content-Type":"application/x-www-form-urlencoded",
			"Referer":"http://gc.nifdc.org.cn/login",
		},
		RedirectLimit:10,
	})
	if err!=nil{
		return "",err
	}
	sbd:=r.String()
	if strings.Index(sbd,"<i class=\"icon-user\"></i>")==-1{
		return "",errors.New(Findval(sbd,"<div id=\"msg_\" class=\"errors\">(.*?)</div>"))
	}
	cks:=r.RawResponse.Cookies()
	scks:=""
	for _,ck:=range cks{
		scks=fmt.Sprintf("%s;%s",scks,ck.String())
	}

	return scks,nil
}
//任务大平台
func TaskIndex(ck string,session *Session)(string,[]*Channel,error){
	cli:=Cli(session)
	surl:="http://gc.nifdc.org.cn/login?service=http%3A%2F%2Fsample.nifdc.org.cn%2Findex.php%3Fm%3DAdmin%26c%3DSSO%26a%3Dindex"
	r,err:=cli.Get(surl,&RequestOptions{
		Headers: map[string]string{
			"Cookie":ck,
		},
	})
	if err!=nil{
		return "",nil,err
	}
	sbd:=r.String()
	rt,_:=goquery.NewDocumentFromReader(strings.NewReader(sbd))
	uid:=rt.Find("#user_id").AttrOr("value","")
	chs:=make([]*Channel,0)
	rt.Find("#user_select option").Each(func(i int, selection *goquery.Selection) {
		chs=append(chs,&Channel{
			Name: selection.Text(),
			Type: selection.AttrOr("value",""),
		})
	})
	if uid==""{
		return "",nil,errors.New("获取uid失败")
	}
	if len(chs)==0{
		return "",nil,errors.New("获取通道失败")
	}
	return uid,chs,nil
}

func Switchchannel(uuid string,_type string,ck string,session *Session)(string,error){
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://sample.nifdc.org.cn/index.php?m=Admin&c=SSO&a=logined&ca_uuid=%s&user_type=%s",uuid,_type)
	fmt.Println(surl)
	r,err:=cli.Get(surl,&RequestOptions{
		Headers: map[string]string{
			"Cookie":ck,
		},
		RedirectLimit:-1,
	})
	if err!=nil{
		return "",err
	}
	cks:=r.RawResponse.Cookies()
	scks:=""
	for _,ck:=range cks{
		scks=fmt.Sprintf("%s;%s",scks,ck.String())
	}
	return scks,nil
}



//已接受全字段导出
func Viewcheckedsample_full(sample_code string,ck string,session *Session)(map[string]string,error){
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://sample.nifdc.org.cn/index.php?m=Admin&c=TaskList&a=viewcheckedsample&sample_code=%s",sample_code)
	r,err:=cli.Get(surl,&RequestOptions{
		Headers: map[string]string{
			"Cookie":ck,
		},
	})
	if err!=nil{
		return nil,nettool.New_neterror_with_e(err)
	}
	if r.StatusCode!=200{
		return nil,nettool.New_neterror_with_s("http状态错误")
	}
	sbd:=r.String()

	return StoMap_yijieshou_full(sbd),nil
}

//抽样完成全字段导出
func Viewnormalsample_full(sample_code string,ck string,session *Session)(map[string]string,error){
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://sample.nifdc.org.cn/index.php?m=Admin&c=TaskList&a=viewnormalsample&sample_code=%s",sample_code)
	r,err:=cli.Get(surl,&RequestOptions{
		Headers: map[string]string{
			"Cookie":ck,
		},
	})
	if err!=nil{
		return nil,nettool.New_neterror_with_e(err)
	}
	if r.StatusCode!=200{
		return nil,nettool.New_neterror_with_s("http状态错误")
	}
	sbd:=r.String()

	return StoMap_full(sbd),nil
}
//抽样完成半字段导出
func Viewnormalsample(sample_code string,ck string,session *Session)(map[string]string,error){
	cli:=Cli(session)
	surl:=fmt.Sprintf("http://sample.nifdc.org.cn/index.php?m=Admin&c=TaskList&a=viewnormalsample&sample_code=%s",sample_code)
	r,err:=cli.Get(surl,&RequestOptions{
		Headers: map[string]string{
			"Cookie":ck,
		},
	})
	if err!=nil{
		return nil,nettool.New_neterror_with_e(err)
	}
	if r.StatusCode!=200{
		return nil,nettool.New_neterror_with_s("http状态错误")
	}
	sbd:=r.String()

	return StoMap(sbd),nil
}

//数据查看
func DownData(sample_state int,ck string,session *Session)(*Download_Data_r,error){
	cli:=Cli(session)
	surl:="http://sample.nifdc.org.cn/index.php?m=Admin&c=TaskList&a=gettasklist"
	r,err:=cli.Post(surl,&RequestOptions{
		Headers: map[string]string{
			"Cookie":ck,
		},
		Data: map[string]string{
			"sEcho":"2",
			"iDisplayStart":"0",
			"iDisplayLength":"10000",
			"sample_state":fmt.Sprintf("%d",sample_state),
		},
	})
	if err!=nil{
		return nil,err
	}
	var rs Download_Data_r
	err=r.JSON(&rs)
	if err!=nil{
		return nil,err
	}

	return &rs,nil
}
