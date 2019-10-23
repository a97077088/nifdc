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
func StoMap_full(s string)map[string]string{
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


	mkr["抽样基础信息_任务来源"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-12.pl0").Eq(0).Text(),"任务来源：",""),"\n")," ")
	mkr["抽样基础信息_报送分类"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-12.pl0").Eq(1).Text(),"报送分类：",""),"\n")," ")
	mkr["抽样基础信息_检验机构名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-12.pl0").Eq(2).Text(),"检验机构名称：",""),"\n")," ")
	mkr["抽样基础信息_部署机构"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(0).Text(),"部署机构：",""),"\n")," ")
	mkr["抽样基础信息_抽样类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(1).Text(),"抽样类型：",""),"\n")," ")
	mkr["抽样基础信息_抽样环节"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(2).Text(),"抽样环节：",""),"\n")," ")
	mkr["抽样基础信息_抽样地点"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(3).Text(),"抽样地点：",""),"\n")," ")
	mkr["抽样基础信息_食品分类"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(4).Text(),"食品分类：",""),"\n")," ")
	mkr["抽样基础信息_抽样单编号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(5).Text(),"抽样单编号：",""),"\n")," ")
	mkr["抽样基础信息_检验目的/任务类别"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(6).Text(),"检验目的/任务类别：",""),"\n")," ")
	mkr["抽样单位信息_单位名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widthtwo.pull-left").Eq(0).Text(),"单位名称：",""),"\n")," ")
	mkr["抽样单位信息_单位地址"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widthtwo.pull-left").Eq(1).Text(),"单位地址：",""),"\n")," ")
	mkr["抽样单位信息_所在省份"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(0).Text(),"所在省份：",""),"\n")," ")
	mkr["抽样单位信息_抽样人员"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(1).Text(),"抽样人员：",""),"\n")," ")
	mkr["抽样单位信息_联系人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(2).Text(),"联系人：",""),"\n")," ")
	mkr["抽样单位信息_电子邮箱"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(3).Text(),"电子邮箱：",""),"\n")," ")
	mkr["抽样单位信息_电话"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(4).Text(),"电话：",""),"\n",""),"\r\n")," ")
	mkr["抽样单位信息_传真"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(5).Text(),"传真：",""),"\n",""),"\r\n")," ")
	mkr["抽样单位信息_邮编"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(6).Text(),"邮编：",""),"\n",""),"\r\n")," ")
	mkr["抽检场所信息_所在地"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm").Eq(0).Text(),"所在地：",""),"\n")," ")
	mkr["抽检场所信息_区域类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm.ar").Eq(0).Text(),"区域类型：",""),"\n")," ")
	mkr["抽检场所信息_单位名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm").Eq(2).Text(),"单位名称：",""),"\n")," ")
	mkr["抽检场所信息_单位地址"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm.ar").Eq(1).Text(),"单位地址：",""),"\n")," ")
	mkr["抽检场所信息_营业执照/社会信用代码"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm").Eq(4).Text(),"营业执照/社会信用代码：",""),"\n")," ")
	mkr["抽检场所信息_许可证类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm").Eq(5).Text(),"许可证类型：",""),"\n")," ")
	mkr["抽检场所信息_经营许可证号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm.ar").Eq(2).Text(),"经营许可证号：",""),"\n")," ")
	mkr["抽检场所信息_年销售额"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm.ar").Eq(3).Text(),"年销售额：",""),"\n")," ")
	mkr["抽检场所信息_单位法人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(5).Text(),"单位法人：",""),"\n")," ")
	mkr["抽检场所信息_联系人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(6).Text(),"联系人：",""),"\n")," ")
	mkr["抽检场所信息_电话"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(7).Text(),"电话：",""),"\n")," ")
	mkr["抽检场所信息_传真"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(8).Text(),"传真：",""),"\n")," ")
	mkr["抽检场所信息_邮编"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(9).Text(),"邮编：",""),"\n")," ")
	mkr["抽检场所信息_摊位号或姓名"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(10).Text(),"摊位号或姓名：",""),"\n")," ")
	mkr["抽检场所信息_身份证号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(11).Text(),"身份证号：",""),"\n")," ")
	mkr["抽样生产企业信息_所在地"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-12.pl0").Eq(0).Text(),"所在地：",""),"\n")," ")
	mkr["抽样生产企业信息_企业地址"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-12.pl0").Eq(1).Text(),"企业地址：",""),"\n")," ")
	mkr["抽样生产企业信息_企业名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-6.pl0").Eq(0).Text(),"企业名称：",""),"\n")," ")
	mkr["抽样生产企业信息_生产许可证编号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-6").Eq(1).Text(),"生产许可证编号：",""),"\n")," ")
	mkr["抽样生产企业信息_生产单位联系人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".widTone.pull-left").Eq(0).Text(),"联系人：",""),"\n")," ")
	mkr["抽样生产企业信息_生产单位电话"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".widTone.pull-left").Eq(1).Text(),"电话：",""),"\n")," ")
	mkr["抽样生产企业信息_是否存在第三方企业信息"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".widTone.pull-left").Eq(2).Text(),"是否存在第三方企业信息：",""),"\n")," ")
	mkr["抽样生产企业信息_第三方企业省份"]=""
	mkr["抽样生产企业信息_第三方企业市区"]=""
	mkr["抽样生产企业信息_第三方企业县区"]=""
	mkr["抽样生产企业信息_第三方企业地址"]=""
	mkr["抽样生产企业信息_第三方企业名称"]=""
	mkr["抽样生产企业信息_第三方企业许可证编号"]=""
	mkr["抽样生产企业信息_第三方企业联系人"]=""
	mkr["抽样生产企业信息_第三方企业电话"]=""
	mkr["抽样生产企业信息_第三方企业性质"]=""
	mkr["抽检样品信息_样品条码"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-6.pl0").Eq(0).Text(),"样品条码：",""),"\n")," ")
	mkr["抽检样品信息_样品商标"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(0).Text(),"样品商标：",""),"\n")," ")
	mkr["抽检样品信息_样品类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(1).Text(),"样品类型：",""),"\n")," ")
	mkr["抽检样品信息_样品来源"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(2).Text(),"样品来源：",""),"\n")," ")
	mkr["抽检样品信息_样品属性"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(3).Text(),"样品属性：",""),"\n")," ")
	mkr["抽检样品信息_包装分类"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(4).Text(),"包装分类：",""),"\n")," ")
	mkr["抽检样品信息_样品名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(5).Text(),"样品名称：",""),"\n")," ")
	mkr["抽检样品信息_购进日期"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(sel4.Find(".col-xs-6.pl0").Eq(1).Text(),"购进日期：",""),"生产日期：",""),"\n")," ")
	mkr["抽检样品信息_保质期"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(6).Text(),"保质期：",""),"\n")," ")
	mkr["抽检样品信息_样品批号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(7).Text(),"样品批号：",""),"\n")," ")
	mkr["抽检样品信息_规格型号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(8).Text(),"规格型号：",""),"\n")," ")
	mkr["抽检样品信息_质量等级"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(9).Text(),"质量等级：",""),"\n")," ")
	mkr["抽检样品信息_单价"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(0).Text(),"单价：",""),"\n")," ")
	mkr["抽检样品信息_是否进口"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(1).Text(),"是否进口：",""),"\n")," ")
	mkr["抽检样品信息_原产地"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(2).Text(),"原产地：",""),"\n")," ")
	mkr["抽检样品信息_抽样日期"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(4).Text(),"抽样日期：",""),"\n")," ")
	mkr["抽检样品信息_抽样方式"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(11).Text(),"抽样方式：",""),"\n")," ")
	mkr["抽检样品信息_样品形态"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(3).Text(),"样品形态：",""),"\n")," ")
	mkr["抽检样品信息_样品包装"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(4).Text(),"样品包装：",""),"\n")," ")
	mkr["抽检样品信息_抽样工具"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(5).Text(),"抽样工具：",""),"\n")," ")
	mkr["抽检样品信息_抽样时样品储存条件"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(12).Text(),"抽样时样品储存条件：",""),"\n")," ")
	mkr["抽检样品信息_抽样基数"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(6).Text(),"抽样基数：",""),"\n")," ")
	mkr["抽检样品信息_抽样数量"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(7).Text(),"抽样数量：",""),"\n")," ")
	mkr["抽检样品信息_备样数量"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(8).Text(),"备样数量：",""),"\n")," ")
	mkr["抽检样品信息_抽样数量单位"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(13).Text(),"抽样数量单位：",""),"\n")," ")
	mkr["抽检样品信息_执行标准/技术文件"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-12.pl0").Eq(0).Text(),"执行标准/技术文件：",""),"\n")," ")
	mkr["抽检样品信息_备注"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-12.pl0").Eq(1).Text(),"备注：",""),"\n")," ")




	return mkr
}
func StoMap_yijieshou_full(s string)map[string]string{
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


	mkr["抽样基础信息_任务来源"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-12.pl0").Eq(0).Text(),"任务来源：",""),"\n")," ")
	mkr["抽样基础信息_报送分类"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-12.pl0").Eq(1).Text(),"报送分类：",""),"\n")," ")
	mkr["抽样基础信息_检验机构名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-12.pl0").Eq(2).Text(),"检验机构名称：",""),"\n")," ")
	mkr["抽样基础信息_部署机构"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(0).Text(),"部署机构：",""),"\n")," ")
	mkr["抽样基础信息_抽样类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(1).Text(),"抽样类型：",""),"\n")," ")
	mkr["抽样基础信息_抽样环节"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(2).Text(),"抽样环节：",""),"\n")," ")
	mkr["抽样基础信息_抽样地点"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(3).Text(),"抽样地点：",""),"\n")," ")
	mkr["抽样基础信息_食品分类"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(4).Text(),"食品分类：",""),"\n")," ")
	mkr["抽样基础信息_抽样单编号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(5).Text(),"抽样单编号：",""),"\n")," ")
	mkr["抽样基础信息_检验目的/任务类别"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel0.Find(".col-xs-6.pl0").Eq(6).Text(),"检验目的/任务类别：",""),"\n")," ")
	mkr["抽样单位信息_单位名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widthtwo.pull-left").Eq(0).Text(),"单位名称：",""),"\n")," ")
	mkr["抽样单位信息_单位地址"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widthtwo.pull-left").Eq(1).Text(),"单位地址：",""),"\n")," ")
	mkr["抽样单位信息_所在省份"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(0).Text(),"所在省份：",""),"\n")," ")
	mkr["抽样单位信息_抽样人员"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(1).Text(),"抽样人员：",""),"\n")," ")
	mkr["抽样单位信息_联系人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(2).Text(),"联系人：",""),"\n")," ")
	mkr["抽样单位信息_电子邮箱"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(3).Text(),"电子邮箱：",""),"\n")," ")
	mkr["抽样单位信息_电话"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(4).Text(),"电话：",""),"\n",""),"\r\n")," ")
	mkr["抽样单位信息_传真"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(5).Text(),"传真：",""),"\n",""),"\r\n")," ")
	mkr["抽样单位信息_邮编"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(sel1.Find(".widTone.pull-left").Eq(6).Text(),"邮编：",""),"\n",""),"\r\n")," ")
	mkr["抽检场所信息_所在地"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm").Eq(0).Text(),"所在地：",""),"\n")," ")
	mkr["抽检场所信息_区域类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm.ar").Eq(0).Text(),"区域类型：",""),"\n")," ")
	mkr["抽检场所信息_单位名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm").Eq(2).Text(),"单位名称：",""),"\n")," ")
	mkr["抽检场所信息_单位地址"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm.ar").Eq(1).Text(),"单位地址：",""),"\n")," ")
	mkr["抽检场所信息_营业执照/社会信用代码"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm").Eq(4).Text(),"营业执照/社会信用代码：",""),"\n")," ")
	mkr["抽检场所信息_许可证类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm").Eq(5).Text(),"许可证类型：",""),"\n")," ")
	mkr["抽检场所信息_经营许可证号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm.ar").Eq(2).Text(),"许可证号：",""),"\n")," ")
	mkr["抽检场所信息_年销售额"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".pull-left.insLabelForm.ar").Eq(3).Text(),"年销售额：",""),"\n")," ")
	mkr["抽检场所信息_单位法人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(5).Text(),"单位法人：",""),"\n")," ")
	mkr["抽检场所信息_联系人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(6).Text(),"联系人：",""),"\n")," ")
	mkr["抽检场所信息_电话"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(7).Text(),"电话：",""),"\n")," ")
	mkr["抽检场所信息_传真"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(8).Text(),"传真：",""),"\n")," ")
	mkr["抽检场所信息_邮编"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(9).Text(),"邮编：",""),"\n")," ")
	mkr["抽检场所信息_摊位号或姓名"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(10).Text(),"摊位号或姓名：",""),"\n")," ")
	mkr["抽检场所信息_身份证号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel2.Find(".widTone.pull-left").Eq(11).Text(),"身份证号：",""),"\n")," ")
	mkr["抽样生产企业信息_所在地"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-12.pl0").Eq(0).Text(),"所在地：",""),"\n")," ")
	mkr["抽样生产企业信息_企业地址"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-6.pl0").Eq(0).Text(),"企业地址：",""),"\n")," ")
	mkr["抽样生产企业信息_企业名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-6.pl0").Eq(1).Text(),"企业名称：",""),"\n")," ")
	mkr["抽样生产企业信息_生产许可证编号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".col-xs-6").Eq(2).Text(),"生产许可证编号：",""),"\n")," ")
	mkr["抽样生产企业信息_生产单位联系人"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".widTone.pull-left").Eq(0).Text(),"联系人：",""),"\n")," ")
	mkr["抽样生产企业信息_生产单位电话"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".widTone.pull-left").Eq(1).Text(),"电话：",""),"\n")," ")
	mkr["抽样生产企业信息_是否存在第三方企业信息"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel3.Find(".widTone.pull-left").Eq(2).Text(),"是否存在第三方企业信息：",""),"\n")," ")
	mkr["抽样生产企业信息_第三方企业省份"]=""
	mkr["抽样生产企业信息_第三方企业市区"]=""
	mkr["抽样生产企业信息_第三方企业县区"]=""
	mkr["抽样生产企业信息_第三方企业地址"]=""
	mkr["抽样生产企业信息_第三方企业名称"]=""
	mkr["抽样生产企业信息_第三方企业许可证编号"]=""
	mkr["抽样生产企业信息_第三方企业联系人"]=""
	mkr["抽样生产企业信息_第三方企业电话"]=""
	mkr["抽样生产企业信息_第三方企业性质"]=""
	mkr["抽检样品信息_样品条码"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-4.pl0").Eq(0).Text(),"样品条码：",""),"\n")," ")
	mkr["抽检样品信息_样品商标"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-4").Eq(1).Text(),"样品商标：",""),"\n")," ")
	mkr["抽检样品信息_样品类型"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-4").Eq(2).Text(),"样品类型：",""),"\n")," ")
	mkr["抽检样品信息_样品来源"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(0).Text(),"样品来源：",""),"\n")," ")
	mkr["抽检样品信息_样品属性"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(1).Text(),"样品属性：",""),"\n")," ")
	mkr["抽检样品信息_包装分类"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(2).Text(),"包装分类：",""),"\n")," ")
	mkr["抽检样品信息_样品名称"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(3).Text(),"样品名称：",""),"\n")," ")
	mkr["抽检样品信息_购进日期"]=strings.Trim(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(sel4.Find(".col-xs-6.pl0").Eq(0).Text(),"购进日期：",""),"生产日期：",""),"\n")," ")
	mkr["抽检样品信息_保质期"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(4).Text(),"保质期：",""),"\n")," ")
	mkr["抽检样品信息_样品批号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(5).Text(),"样品批号：",""),"\n")," ")
	mkr["抽检样品信息_规格型号"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(6).Text(),"规格型号：",""),"\n")," ")
	mkr["抽检样品信息_质量等级"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(7).Text(),"质量等级：",""),"\n")," ")
	mkr["抽检样品信息_单价"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(0).Text(),"单价：",""),"\n")," ")
	mkr["抽检样品信息_是否进口"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(1).Text(),"是否进口：",""),"\n")," ")
	mkr["抽检样品信息_原产地"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(2).Text(),"原产地：",""),"\n")," ")
	mkr["抽检样品信息_抽样日期"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3.pl0").Eq(4).Text(),"抽样日期：",""),"\n")," ")
	mkr["抽检样品信息_抽样方式"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(8).Text(),"抽样方式：",""),"\n")," ")
	mkr["抽检样品信息_样品形态"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(3).Text(),"样品形态：",""),"\n")," ")
	mkr["抽检样品信息_样品包装"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(4).Text(),"样品包装：",""),"\n")," ")
	mkr["抽检样品信息_抽样工具"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(5).Text(),"抽样工具：",""),"\n")," ")
	mkr["抽检样品信息_抽样时样品储存条件"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(10).Text(),"抽样时样品储存条件：",""),"\n")," ")
	mkr["抽检样品信息_抽样基数"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-3").Eq(11).Text(),"抽样基数：",""),"\n")," ")
	mkr["抽检样品信息_抽样数量"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(6).Text(),"抽样数量：",""),"\n")," ")
	mkr["抽检样品信息_备样数量"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(7).Text(),"备样数量：",""),"\n")," ")
	mkr["抽检样品信息_抽样数量单位"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-2").Eq(8).Text(),"抽样数量单位：",""),"\n")," ")
	mkr["抽检样品信息_执行标准/技术文件"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-15.pl0").Eq(0).Text(),"执行标准/技术文件：",""),"\n")," ")
	mkr["抽检样品信息_备注"]=strings.Trim(strings.Trim(strings.ReplaceAll(sel4.Find(".col-xs-15.pl0").Eq(1).Text(),"备注：",""),"\n")," ")




	return mkr
}