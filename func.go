package nifdc

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func StoMap(s string)map[string]string{
	rt,_:=goquery.NewDocumentFromReader(strings.NewReader(s))
	//抽样基础信息
	sel0:=rt.Find(".formArea1.formAreaCom").Eq(0)
	//抽样单位信息
	sel1:=rt.Find(".formArea1.formAreaCom").Eq(1)
	//抽样场所信息
	sel2:=rt.Find(".formArea1.formAreaCom").Eq(2)
	//抽样生产企业信
	sel3:=rt.Find(".formArea1.formAreaCom").Eq(3)
	//抽检样品信息
	sel4:=rt.Find(".formArea1.formAreaCom").Eq(4)
	mkr:=make(map[string]string,0)


	mkr["委托单位"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-12.pl0").Eq(0).Text(),"任务来源：",""),"\n")," ")
	mkr["抽样地点"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(3).Text(),"抽样地点：",""),"\n")," ")
	mkr["抽样单号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(5).Text(),"抽样单编号：",""),"\n")," ")
	mkr["检验类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(6).Text(),"检验目的/任务类别：",""),"\n")," ")
	mkr["抽送样人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(1).Text(),"抽样人员：",""),"\n")," ")
	mkr["受检单位"]=strings.ReplaceAll(sel2.Find("span").Eq(2).Text(),"单位名称：","")
	mkr["地址"]=strings.ReplaceAll(sel2.Find("span").Eq(3).Text(),"单位地址：","")
	mkr["联系人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left.pl10").Eq(3).Text(),"联系人：",""),"\n")," ")
	mkr["电话"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(7).Text(),"电话：",""),"\n")," ")
	mkr["生产单位地址"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-12.pl0").Eq(1).Text(),"企业地址：",""),"\n")," ")
	mkr["生产单位"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-6.pl0").Eq(0).Text(),"企业名称：",""),"\n")," ")
	mkr["生产单位联系人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".widTone.pull-left").Eq(0).Text(),"联系人：",""),"\n")," ")
	mkr["生产单位电话"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".widTone.pull-left.pl10").Eq(0).Text(),"电话：",""),"\n")," ")
	mkr["商标"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(0).Text(),"样品商标：",""),"\n")," ")
	mkr["样品名称br"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(5).Text(),"样品名称：",""),"\n")," ")
	mkr["生产日期"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(sel4.Find(".col-xs-6.pl0").Eq(1).Text(),"购进日期：",""),"生产日期：",""),"加工日期：",""),"\n")," ")
	mkr["保质期"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(6).Text(),"保质期：",""),"\n")," ")
	mkr["生产批号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(7).Text(),"样品批号：",""),"\n")," ")
	mkr["规格型号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(2).Text(),"规格型号：",""),"\n")," ")
	mkr["样品等级"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(3).Text(),"质量等级：",""),"\n")," ")
	mkr["抽到样日期"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(4).Text(),"抽样日期：",""),"\n")," ")
	mkr["抽样方式"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(5).Text(),"抽样方式：",""),"\n")," ")
	mkr["样品状态"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(3).Text(),"样品形态：",""),"\n")," ")
	mkr["样品状态2"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(4).Text(),"样品包装：",""),"\n")," ")
	mkr["保存条件"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(6).Text(),"抽样时样品储存条件：",""),"\n")," ")
	mkr["抽样基数"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2.pl0").Eq(0).Text(),"抽样基数：",""),"\n")," ")
	mkr["样品数"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(7).Text(),"抽样数量：",""),"\n")," ")
	mkr["检验依据"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-12.pl0").Eq(0).Text(),"执行标准/技术文件：",""),"\n")," ")
	mkr["备注"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-12.pl0").Eq(1).Text(),"备注：",""),"\n")," ")
	return mkr
}
