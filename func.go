package nifdc

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func TestInfotoMap(tos []*Test_platform_api_food_getTestInfo_o) []map[string]string {
	r := make([]map[string]string, 0)
	for _, to := range tos {
		subr := make(map[string]string)
		subr["检验项目*"] = to.Spdata_0
		subr["检验结果*"] = to.Spdata_1
		subr["结果单位*"] = to.Spdata_18
		subr["结果判定*"] = to.Spdata_2
		subr["检验依据*"] = to.Spdata_3
		subr["判定依据*"] = to.Spdata_4
		subr["最小允许限*"] = to.Spdata_11
		subr["最大允许限*"] = to.Spdata_15
		subr["允许限单位*"] = to.Spdata_16
		subr["方法检出限*"] = to.Spdata_7
		subr["检出限单位*"] = to.Spdata_8
		subr["备注"] = to.Spdata_20
		subr["说明"] = to.Spdata_17
		r = append(r, subr)
	}
	return r
}

func StoMap_foodDetail(s string) map[string]string {
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	//抽样基础信息
	sel0 := rt.Find(".ibox.float-e-margins").Eq(0)
	//抽样单位信息
	sel1 := rt.Find(".ibox.float-e-margins").Eq(1)
	//抽检场所信息
	sel2 := rt.Find(".ibox.float-e-margins").Eq(2)
	//生产企业信息
	sel3 := rt.Find(".ibox.float-e-margins").Eq(3)
	//抽检样品信息
	sel4 := rt.Find(".ibox.float-e-margins").Eq(4)
	////照片信息
	//sel5 := rt.Find(".ibox.float-e-margins").Eq(5)
	//sel5html, _ := sel5.Html()
	//检验信息
	sel6 := rt.Find(".ibox.float-e-margins").Eq(6)
	//sel6html, _ := sel6.Html()
	mkr := make(map[string]string, 0)

	mkr["抽样基础信息_任务来源"] = FindNextNodeVal(sel0, "任务来源")
	mkr["抽样基础信息_报送分类A"] = FindNextNodeVal(sel0, "报送分类A")
	mkr["抽样基础信息_报送分类B"] = FindNextNodeVal(sel0, "报送分类B")
	mkr["抽样基础信息_食品大类"] = FindNextNodeVal(sel0, "食品大类")
	mkr["抽样基础信息_食品亚类"] = FindNextNodeVal(sel0, "食品亚类")
	mkr["抽样基础信息_食品次亚类"] = FindNextNodeVal(sel0, "食品次亚类")
	mkr["抽样基础信息_食品细类"] = FindNextNodeVal(sel0, "食品细类")
	mkr["抽样基础信息_抽样单编号"] = FindNextNodeVal(sel0, "抽样单编号")
	mkr["抽样基础信息_抽样类型"] = FindNextNodeVal(sel0, "抽样类型")

	mkr["抽样单位信息_抽样单位名称"] = FindNextNodeVal(sel1, "抽样单位名称")
	mkr["抽样单位信息_单位地址"] = FindNextNodeVal(sel1, "单位地址")
	mkr["抽样单位信息_所在省份"] = FindNextNodeVal(sel1, "所在省份")
	mkr["抽样单位信息_抽样人员"] = FindNextNodeVal(sel1, "抽样人员")
	mkr["抽样单位信息_抽样人员电话"] = FindNextNodeVal(sel1, "抽样人员电话")
	mkr["抽样单位信息_单位联系人"] = FindNextNodeVal(sel1, "单位联系人")
	mkr["抽样单位信息_传真"] = FindNextNodeVal(sel1, "传真")
	mkr["抽样单位信息_邮编"] = FindNextNodeVal(sel1, "邮编")
	mkr["抽样单位信息_联系人电话"] = FindNextNodeVal(sel1, "联系人电话")

	mkr["抽检场所信息_单位名称"] = FindNextNodeVal(sel2, "单位名称")
	mkr["抽检场所信息_单位地址"] = FindNextNodeVal(sel2, "单位地址")
	mkr["抽检场所信息_所在地"] = FindNextNodeVal(sel2, "所在地")
	mkr["抽检场所信息_区域类型"] = FindNextNodeVal(sel2, "区域类型")
	mkr["抽检场所信息_抽样环节"] = FindNextNodeVal(sel2, "抽样环节")
	mkr["抽检场所信息_抽样地点"] = FindNextNodeVal(sel2, "抽样地点")
	mkr["抽检场所信息_营业执照/社会信用代码"] = FindNextNodeVal(sel2, "营业执照\\/社会信用代码")
	mkr["抽检场所信息_许可证类型"] = FindNextNodeVal(sel2, "许可证类型")
	mkr["抽检场所信息_许可证号"] = FindNextNodeVal(sel2, "许可证号")
	mkr["抽检场所信息_单位法人"] = FindNextNodeVal(sel2, "单位法人")
	mkr["抽检场所信息_联系人"] = FindNextNodeVal(sel2, "联系人")
	mkr["抽检场所信息_联系人电话"] = FindNextNodeVal(sel2, "联系人电话")
	mkr["抽检场所信息_传真"] = FindNextNodeVal(sel2, "传真")
	mkr["抽检场所信息_邮编"] = FindNextNodeVal(sel2, "邮编")
	mkr["抽检场所信息_年销售额"] = FindNextNodeVal(sel2, "年销售额")

	mkr["生产企业信息_企业名称"] = FindNextNodeVal(sel3, "企业名称")
	mkr["生产企业信息_所在地"] = FindNextNodeVal(sel3, "所在地")
	mkr["生产企业信息_企业地址"] = FindNextNodeVal(sel3, "企业地址")
	mkr["生产企业信息_生产许可证编号"] = FindNextNodeVal(sel3, "生产许可证编号")
	mkr["生产企业信息_联系人"] = FindNextNodeVal(sel3, "联系人")
	mkr["生产企业信息_电话"] = FindNextNodeVal(sel3, "电话")
	mkr["生产企业信息_是否存在第三方企业信息"] = FindNextNodeVal(sel3, "是否存在第三方企业信息")

	mkr["抽检样品信息_样品条码"] = FindNextNodeVal(sel4, "样品条码")
	mkr["抽检样品信息_样品商标"] = FindNextNodeVal(sel4, "样品商标")
	mkr["抽检样品信息_样品类型"] = FindNextNodeVal(sel4, "样品类型")
	mkr["抽检样品信息_样品来源"] = FindNextNodeVal(sel4, "样品来源")
	mkr["抽检样品信息_样品属性"] = FindNextNodeVal(sel4, "样品属性")
	mkr["抽检样品信息_包装分类"] = FindNextNodeVal(sel4, "包装分类")
	mkr["抽检样品信息_样品名称"] = FindNextNodeVal(sel4, "样品名称")
	mkr["抽检样品信息_生产日期"] = FindNextNodeVal(sel4, "生产日期") + FindNextNodeVal(sel4, "抽样日期")
	mkr["抽检样品信息_保质期"] = FindNextNodeVal(sel4, "保质期")
	mkr["抽检样品信息_样品批号"] = FindNextNodeVal(sel4, "样品批号")
	mkr["抽检样品信息_规格型号"] = FindNextNodeVal(sel4, "规格型号")
	mkr["抽检样品信息_质量等级"] = FindNextNodeVal(sel4, "质量等级")
	mkr["抽检样品信息_单价"] = FindNextNodeVal(sel4, "单价")
	mkr["抽检样品信息_是否进口"] = FindNextNodeVal(sel4, "是否进口")
	mkr["抽检样品信息_原产地"] = FindNextNodeVal(sel4, "原产地")
	mkr["抽检样品信息_抽样日期"] = FindNextNodeVal(sel4, "抽样日期")
	mkr["抽检样品信息_抽样方式"] = FindNextNodeVal(sel4, "抽样方式")
	mkr["抽检样品信息_抽样时样品储存条件"] = FindNextNodeVal(sel4, "抽样时样品储存条件")
	mkr["抽检样品信息_抽样基数"] = FindNextNodeVal(sel4, "抽样基数")
	mkr["抽检样品信息_抽样数量"] = FindNextNodeVal(sel4, "抽样数量")
	mkr["抽检样品信息_备样数量"] = FindNextNodeVal(sel4, "备样数量")
	mkr["抽检样品信息_抽样数量单位"] = FindNextNodeVal(sel4, "抽样数量单位")
	mkr["抽检样品信息_执行标准/技术文件"] = FindNextNodeVal(sel4, "执行标准\\/技术文件")
	mkr["抽检样品信息_备注"] = FindNextNodeVal(sel4, "备注")

	mkr["检验信息_检验机构名称"] = FindNextNodeVal(sel6, "检验机构名称")
	mkr["检验信息_报告书编号"] = FindNextNodeVal(sel6, "报告书编号")
	mkr["检验信息_样品到达日期"] = FindNextNodeVal(sel6, "样品到达日期")
	mkr["检验信息_联系人"] = FindNextNodeVal(sel6, "联系人")
	mkr["检验信息_联系人电话"] = FindNextNodeVal(sel6, "联系人电话")
	mkr["检验信息_联系人邮箱"] = FindNextNodeVal(sel6, "联系人邮箱")
	mkr["检验信息_检查封样人员"] = FindNextNodeVal(sel6, "检查封样人员")
	mkr["检验信息_检查封样人电话"] = FindNextNodeVal(sel6, "检查封样人电话")
	mkr["检验信息_检查封样人邮箱"] = FindNextNodeVal(sel6, "检查封样人邮箱")
	mkr["检验信息_结论"] = FindNextNodeVal(sel6, "结论")
	mkr["检验信息_监督抽检报告备注"] = FindNextNodeVal(sel6, "监督抽检报告备注")
	mkr["检验信息_风险监测报告备注"] = FindNextNodeVal(sel6, "风险监测报告备注")
	mkr["检验信息_复检状态"] = FindNextNodeVal(sel6, "复检状态")

	mkr["检验结论"] = strings.TrimSpace(rt.Find("#testform").Find("h3:contains(检验结论)").Parent().Find("p").Text())
	mkr["状态"] = strings.TrimSpace(strings.ReplaceAll(rt.Find("h2:contains(状态\\:)").Text(), "状态:", ""))

	if mkr["状态"] != "" {
		mkr["检验信息_检验目的/任务类别"] = FindNextNodeVal(sel6, "检验目的\\/任务类别")
		mkr["检验信息_报告类别"] = FindNextNodeVal(sel6, "报告类别")
	} else {
		mkr["检验信息_检验目的/任务类别"] = sel6.Find("label:contains(检验目的\\/任务类别：)").Next().Find("option[selected=selected]").Text()
		mkr["检验信息_报告类别"] = sel6.Find("label:contains(报告类别：)").Next().Find("option").First().Text()
	}
	return mkr
}
func StoMap_chouyangwancheng_mode1(s string) map[string]string {
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	//抽样基础信息
	sel0 := rt.Find(".formArea1.formAreaCom").Eq(0)
	sel0html, _ := sel0.Html()
	//抽样单位信息
	sel1 := rt.Find(".formArea1.formAreaCom").Eq(1)
	sel1html, _ := sel1.Html()
	//抽样场所信息
	sel2 := rt.Find(".formArea1.formAreaCom").Eq(2)
	sel2html, _ := sel2.Html()
	//抽样生产企业信
	sel3 := rt.Find(".formArea1.formAreaCom").Eq(3)
	sel3html, _ := sel3.Html()
	//抽检样品信息
	sel4 := rt.Find(".formArea1.formAreaCom").Eq(4)
	sel4html, _ := sel4.Html()
	mkr := make(map[string]string, 0)

	mkr["委托单位"] = FindFdval(sel0html, "任务来源")
	mkr["抽样地点"] = FindFdval(sel0html, "抽样地点")
	mkr["抽样单号"] = FindFdval(sel0html, "抽样单编号")
	mkr["检验类型"] = FindFdval(sel0html, "检验目的/任务类别")
	mkr["抽送样人"] = FindFdval(sel1html, "抽样人员")
	mkr["受检单位"] = FindFdval(sel2html, "单位名称")
	mkr["地址"] = FindFdval(sel2html, "单位地址")
	mkr["联系人"] = FindFdval(sel2html, "联系人")
	mkr["电话"] = FindFdval(sel2html, "联系电话")
	mkr["生产单位地址"] = FindFdval(sel3html, "生产者地址")
	mkr["生产单位"] = FindFdval(sel3html, "生产者名称")
	mkr["生产单位联系人"] = FindFdval(sel3html, "联系人")
	mkr["生产单位电话"] = FindFdval(sel3html, "联系电话")
	mkr["商标"] = FindFdval(sel4html, "样品商标")
	mkr["样品名称br"] = FindFdval(sel4html, "样品名称")
	mkr["生产日期"] = FindFdval(sel4html, "购进日期") + FindFdval(sel4html, "生产日期") + FindFdval(sel4html, "加工日期") + FindFdval(sel4html, "抽样日期")
	mkr["保质期"] = FindFdval(sel4html, "保质期")
	mkr["生产批号"] = FindFdval(sel4html, "样品批号")
	mkr["规格型号"] = FindFdval(sel4html, "规格型号")
	mkr["样品等级"] = FindFdval(sel4html, "质量等级")
	mkr["抽到样日期"] = FindFdval(sel4html, "抽样日期")
	mkr["抽样方式"] = FindFdval(sel4html, "抽样方式")
	mkr["样品状态"] = FindFdval(sel4html, "样品状态")
	mkr["样品状态2"] = FindFdval(sel4html, "样品状态2")
	mkr["保存条件"] = FindFdval(sel4html, "储存条件")
	mkr["抽样基数"] = FindFdval(sel4html, "抽样基数")
	mkr["样品数"] = FindFdval(sel4html, "抽样数量")
	mkr["检验依据"] = FindFdval(sel4html, "执行标准/技术文件")
	mkr["备注"] = FindFdval(sel4html, "整个样品备注")
	return mkr
}
func StoMap_chouyangwancheng_full(s string) map[string]string {
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	//抽样基础信息
	sel0 := rt.Find(".formArea1.formAreaCom").Eq(0)
	sel0html, _ := sel0.Html()
	//抽样单位信息
	sel1 := rt.Find(".formArea1.formAreaCom").Eq(1)
	sel1html, _ := sel1.Html()
	//抽样场所信息
	sel2 := rt.Find(".formArea1.formAreaCom").Eq(2)
	sel2html, _ := sel2.Html()
	//抽样生产企业信
	sel3 := rt.Find(".formArea1.formAreaCom").Eq(3)
	sel3html, _ := sel3.Html()
	//抽检样品信息
	sel4 := rt.Find(".formArea1.formAreaCom").Eq(4)
	sel4html, _ := sel4.Html()

	mkr := make(map[string]string, 0)

	mkr["抽样基础信息_任务来源"] = FindFdval(sel0html, "任务来源")
	mkr["抽样基础信息_报送分类"] = FindFdval(sel0html, "报送分类")
	mkr["抽样基础信息_检验机构名称"] = FindFdval(sel0html, "检验机构名称")
	mkr["抽样基础信息_部署机构"] = FindFdval(sel0html, "部署机构")
	mkr["抽样基础信息_抽样类型"] = FindFdval(sel0html, "抽样类型")
	mkr["抽样基础信息_抽样环节"] = FindFdval(sel0html, "抽样环节")
	mkr["抽样基础信息_抽样地点"] = FindFdval(sel0html, "抽样地点")
	mkr["抽样基础信息_食品分类"] = strings.ReplaceAll(FindFdval(sel0html, "食品分类"), "                                    ", " ")
	mkr["抽样基础信息_抽样单编号"] = FindFdval(sel0html, "抽样单编号")
	mkr["抽样基础信息_检验目的/任务类别"] = FindFdval(sel0html, "检验目的/任务类别")
	mkr["抽样单位信息_单位名称"] = FindFdval(sel1html, "单位名称")
	mkr["抽样单位信息_单位地址"] = FindFdval(sel1html, "单位地址")
	mkr["抽样单位信息_所在省份"] = FindFdval(sel1html, "所在省份")
	mkr["抽样单位信息_抽样人员"] = FindFdval(sel1html, "抽样人员")
	mkr["抽样单位信息_联系人"] = FindFdval(sel1html, "联系人")
	mkr["抽样单位信息_电子邮箱"] = FindFdval(sel1html, "电子邮箱")
	mkr["抽样单位信息_电话"] = FindFdval(sel1html, "电话")
	mkr["抽样单位信息_传真"] = FindFdval(sel1html, "传真")
	mkr["抽样单位信息_邮编"] = FindFdval(sel1html, "邮编")
	mkr["抽检场所信息_所在地"] = strings.ReplaceAll(FindFdval(sel2html, "所在地"), "                                        ", " ")
	mkr["抽检场所信息_区域类型"] = FindFdval(sel2html, "区域类型")
	mkr["抽检场所信息_单位名称"] = FindFdval(sel2html, "单位名称")
	mkr["抽检场所信息_单位地址"] = FindFdval(sel2html, "单位地址")
	mkr["抽检场所信息_营业执照/社会信用代码"] = FindFdval(sel2html, "营业执照/社会信用代码")
	mkr["抽检场所信息_许可证类型"] = FindFdval(sel2html, "许可证类型")
	mkr["抽检场所信息_经营许可证号"] = FindFdval(sel2html, "经营许可证号") + FindFdval(sel2html, "许可证号")
	mkr["抽检场所信息_年销售额"] = FindFdval(sel2html, "年销售额")
	mkr["抽检场所信息_单位法人"] = FindFdval(sel2html, "单位法人") + FindFdval(sel2html, "法人代表")
	mkr["抽检场所信息_联系人"] = FindFdval(sel2html, "联系人")
	mkr["抽检场所信息_电话"] = FindFdval(sel2html, "联系电话")
	mkr["抽检场所信息_传真"] = FindFdval(sel2html, "传真")
	mkr["抽检场所信息_邮编"] = FindFdval(sel2html, "邮编")
	mkr["抽检场所信息_摊位号或姓名"] = FindFdval(sel2html, "摊位号或姓名")
	mkr["抽检场所信息_身份证号"] = FindFdval(sel2html, "身份证号")
	mkr["抽样生产企业信息_所在地"] = strings.ReplaceAll(FindFdval(sel3html, "所在地"), "                                        ", " ")
	mkr["抽样生产企业信息_企业地址"] = FindFdval(sel3html, "生产者地址")
	mkr["抽样生产企业信息_企业名称"] = FindFdval(sel3html, "生产者名称")
	mkr["抽样生产企业信息_生产许可证编号"] = FindFdval(sel3html, "生产许可证编号")
	mkr["抽样生产企业信息_生产单位联系人"] = FindFdval(sel3html, "生产单位联系人")
	mkr["抽样生产企业信息_生产单位电话"] = FindFdval(sel3html, "联系电话")
	mkr["抽样生产企业信息_是否存在第三方企业信息"] = FindFdval(sel3html, "是否存在第三方企业信息")
	if mkr["抽样生产企业信息_是否存在第三方企业信息"] == "是" {
		rhm, _ := sel3.Find(".wtsc.clearfix").Html()
		dsfdz := strings.ReplaceAll(FindFdval(rhm, "所在地"), "                                                ", " ")
		spdsfdz := strings.Split(dsfdz, " ")
		if len(spdsfdz) == 3 {
			mkr["抽样生产企业信息_第三方企业省份"] = spdsfdz[0]
			mkr["抽样生产企业信息_第三方企业市区"] = spdsfdz[1]
			mkr["抽样生产企业信息_第三方企业县区"] = spdsfdz[2]
		}
		mkr["抽样生产企业信息_第三方企业地址"] = FindFdval(sel3html, "第三方企业地址")
		mkr["抽样生产企业信息_第三方企业名称"] = FindFdval(sel3html, "第三方企业名称")
		mkr["抽样生产企业信息_第三方企业许可证编号"] = FindFdval(sel3html, "第三方企业许可证编号")
		mkr["抽样生产企业信息_第三方企业联系人"] = ""
		mkr["抽样生产企业信息_第三方企业电话"] = FindFdval(sel3html, "第三方企业电话")
		mkr["抽样生产企业信息_第三方企业性质"] = FindFdval(sel3html, "第三方企业性质")
	}
	mkr["抽检样品信息_样品条码"] = FindFdval(sel4html, "样品条码")
	mkr["抽检样品信息_样品商标"] = FindFdval(sel4html, "样品商标")
	mkr["抽检样品信息_样品类型"] = FindFdval(sel4html, "样品类型")
	mkr["抽检样品信息_样品来源"] = FindFdval(sel4html, "样品来源")
	mkr["抽检样品信息_样品属性"] = FindFdval(sel4html, "样品属性")
	mkr["抽检样品信息_包装分类"] = FindFdval(sel4html, "包装分类")
	mkr["抽检样品信息_样品名称"] = FindFdval(sel4html, "样品名称")
	mkr["抽检样品信息_购进日期"] = FindFdval(sel4html, "购进日期") + FindFdval(sel4html, "生产日期") + FindFdval(sel4html, "抽样日期")
	mkr["抽检样品信息_保质期"] = FindFdval(sel4html, "保质期")
	mkr["抽检样品信息_样品批号"] = FindFdval(sel4html, "样品批号")
	mkr["抽检样品信息_规格型号"] = FindFdval(sel4html, "规格型号")
	mkr["抽检样品信息_质量等级"] = FindFdval(sel4html, "质量等级")
	mkr["抽检样品信息_单价"] = FindFdval(sel4html, "单价")
	mkr["抽检样品信息_是否进口"] = FindFdval(sel4html, "是否进口")
	mkr["抽检样品信息_原产地"] = FindFdval(sel4html, "原产地")
	mkr["抽检样品信息_抽样日期"] = FindFdval(sel4html, "抽样日期")
	mkr["抽检样品信息_抽样方式"] = FindFdval(sel4html, "抽样方式")
	mkr["抽检样品信息_样品形态"] = FindFdval(sel4html, "样品形态")
	mkr["抽检样品信息_样品包装"] = FindFdval(sel4html, "样品包装")
	mkr["抽检样品信息_抽样工具"] = FindFdval(sel4html, "抽样工具")
	mkr["抽检样品信息_抽样时样品储存条件"] = FindFdval(sel4html, "储存条件")
	mkr["抽检样品信息_抽样基数"] = FindFdval(sel4html, "抽样基数")
	mkr["抽检样品信息_抽样数量"] = FindFdval(sel4html, "抽样数量")
	mkr["抽检样品信息_备样数量"] = FindFdval(sel4html, "备样数量")
	mkr["抽检样品信息_抽样数量单位"] = FindFdval(sel4html, "抽样数量单位")
	mkr["抽检样品信息_执行标准/技术文件"] = FindFdval(sel4html, "执行标准/技术文件")
	mkr["抽检样品信息_备注"] = FindFdval(sel4html, "整个样品备注")

	return mkr
}
func StoMap_yijieshou_full(s string) map[string]string {
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	//抽样基础信息
	sel0 := rt.Find(".formArea1.formAreaCom").Eq(0)
	sel0html, _ := sel0.Html()
	//抽样单位信息
	sel1 := rt.Find(".formArea1.formAreaCom").Eq(1)
	sel1html, _ := sel1.Html()
	//抽样场所信息
	sel2 := rt.Find(".formArea1.formAreaCom").Eq(2)
	sel2html, _ := sel2.Html()
	//抽样生产企业信
	sel3 := rt.Find(".formArea1.formAreaCom").Eq(3)
	sel3html, _ := sel3.Html()
	//抽检样品信息
	sel4 := rt.Find(".formArea1.formAreaCom").Eq(4)
	sel4html, _ := sel4.Html()

	mkr := make(map[string]string, 0)

	mkr["抽样基础信息_任务来源"] = FindFdval(sel0html, "任务来源")
	mkr["抽样基础信息_报送分类"] = FindFdval(sel0html, "报送分类")
	mkr["抽样基础信息_检验机构名称"] = FindFdval(sel0html, "检验机构名称")
	mkr["抽样基础信息_部署机构"] = FindFdval(sel0html, "部署机构")
	mkr["抽样基础信息_抽样类型"] = FindFdval(sel0html, "抽样类型")
	mkr["抽样基础信息_抽样环节"] = FindFdval(sel0html, "抽样环节")
	mkr["抽样基础信息_抽样地点"] = FindFdval(sel0html, "抽样地点")
	mkr["抽样基础信息_食品分类"] = strings.ReplaceAll(FindFdval(sel0html, "食品分类"), "                                    ", " ")
	mkr["抽样基础信息_抽样单编号"] = FindFdval(sel0html, "抽样单编号")
	mkr["抽样基础信息_检验目的/任务类别"] = FindFdval(sel0html, "检验目的/任务类别")
	mkr["抽样单位信息_单位名称"] = FindFdval(sel1html, "单位名称")
	mkr["抽样单位信息_单位地址"] = FindFdval(sel1html, "单位地址")
	mkr["抽样单位信息_所在省份"] = FindFdval(sel1html, "所在省份")
	mkr["抽样单位信息_抽样人员"] = FindFdval(sel1html, "抽样人员")
	mkr["抽样单位信息_联系人"] = FindFdval(sel1html, "联系人")
	mkr["抽样单位信息_电子邮箱"] = FindFdval(sel1html, "电子邮箱")
	mkr["抽样单位信息_电话"] = FindFdval(sel1html, "电话")
	mkr["抽样单位信息_传真"] = FindFdval(sel1html, "传真")
	mkr["抽样单位信息_邮编"] = FindFdval(sel1html, "邮编")
	mkr["抽检场所信息_所在地"] = strings.ReplaceAll(FindFdval(sel2html, "所在地"), "                                        ", " ")
	mkr["抽检场所信息_区域类型"] = FindFdval(sel2html, "区域类型")
	mkr["抽检场所信息_单位名称"] = FindFdval(sel2html, "单位名称")
	mkr["抽检场所信息_单位地址"] = FindFdval(sel2html, "单位地址")
	mkr["抽检场所信息_营业执照/社会信用代码"] = FindFdval(sel2html, "营业执照/社会信用代码")
	mkr["抽检场所信息_许可证类型"] = FindFdval(sel2html, "许可证类型")
	mkr["抽检场所信息_经营许可证号"] = FindFdval(sel2html, "经营许可证号") + FindFdval(sel2html, "许可证号")
	mkr["抽检场所信息_年销售额"] = FindFdval(sel2html, "年销售额")
	mkr["抽检场所信息_单位法人"] = FindFdval(sel2html, "单位法人") + FindFdval(sel2html, "法人代表")
	mkr["抽检场所信息_联系人"] = FindFdval(sel2html, "联系人")
	mkr["抽检场所信息_电话"] = FindFdval(sel2html, "联系电话")
	mkr["抽检场所信息_传真"] = FindFdval(sel2html, "传真")
	mkr["抽检场所信息_邮编"] = FindFdval(sel2html, "邮编")
	mkr["抽检场所信息_摊位号或姓名"] = FindFdval(sel2html, "摊位号或姓名")
	mkr["抽检场所信息_身份证号"] = FindFdval(sel2html, "身份证号")
	mkr["抽样生产企业信息_所在地"] = strings.ReplaceAll(FindFdval(sel3html, "所在地"), "                                        ", " ")
	mkr["抽样生产企业信息_企业地址"] = FindFdval(sel3html, "生产者地址")
	mkr["抽样生产企业信息_企业名称"] = FindFdval(sel3html, "生产者名称")
	mkr["抽样生产企业信息_生产许可证编号"] = FindFdval(sel3html, "生产许可证编号")
	mkr["抽样生产企业信息_生产单位联系人"] = FindFdval(sel3html, "生产单位联系人")
	mkr["抽样生产企业信息_生产单位电话"] = FindFdval(sel3html, "联系电话")
	mkr["抽样生产企业信息_是否存在第三方企业信息"] = FindFdval(sel3html, "是否存在第三方企业信息")
	if mkr["抽样生产企业信息_是否存在第三方企业信息"] == "是" {
		rhm, _ := sel3.Find(".wtsc.clearfix").Html()
		dsfdz := strings.ReplaceAll(FindFdval(rhm, "所在地"), "                                                ", " ")
		spdsfdz := strings.Split(dsfdz, " ")
		if len(spdsfdz) == 3 {
			mkr["抽样生产企业信息_第三方企业省份"] = spdsfdz[0]
			mkr["抽样生产企业信息_第三方企业市区"] = spdsfdz[1]
			mkr["抽样生产企业信息_第三方企业县区"] = spdsfdz[2]
		}
		mkr["抽样生产企业信息_第三方企业地址"] = FindFdval(sel3html, "第三方企业地址")
		mkr["抽样生产企业信息_第三方企业名称"] = FindFdval(sel3html, "第三方企业名称")
		mkr["抽样生产企业信息_第三方企业许可证编号"] = FindFdval(sel3html, "第三方企业许可证编号")
		mkr["抽样生产企业信息_第三方企业联系人"] = ""
		mkr["抽样生产企业信息_第三方企业电话"] = FindFdval(sel3html, "第三方企业电话")
		mkr["抽样生产企业信息_第三方企业性质"] = FindFdval(sel3html, "第三方企业性质")
	}
	mkr["抽检样品信息_样品条码"] = FindFdval(sel4html, "样品条码")
	mkr["抽检样品信息_样品商标"] = FindFdval(sel4html, "样品商标")
	mkr["抽检样品信息_样品类型"] = FindFdval(sel4html, "样品类型")
	mkr["抽检样品信息_样品来源"] = FindFdval(sel4html, "样品来源")
	mkr["抽检样品信息_样品属性"] = FindFdval(sel4html, "样品属性")
	mkr["抽检样品信息_包装分类"] = FindFdval(sel4html, "包装分类")
	mkr["抽检样品信息_样品名称"] = FindFdval(sel4html, "样品名称")
	mkr["抽检样品信息_购进日期"] = FindFdval(sel4html, "购进日期") + FindFdval(sel4html, "生产日期") + FindFdval(sel4html, "抽样日期")
	mkr["抽检样品信息_保质期"] = FindFdval(sel4html, "保质期")
	mkr["抽检样品信息_样品批号"] = FindFdval(sel4html, "样品批号")
	mkr["抽检样品信息_规格型号"] = FindFdval(sel4html, "规格型号")
	mkr["抽检样品信息_质量等级"] = FindFdval(sel4html, "质量等级")
	mkr["抽检样品信息_单价"] = FindFdval(sel4html, "单价")
	mkr["抽检样品信息_是否进口"] = FindFdval(sel4html, "是否进口")
	mkr["抽检样品信息_原产地"] = FindFdval(sel4html, "原产地")
	mkr["抽检样品信息_抽样日期"] = FindFdval(sel4html, "抽样日期")
	mkr["抽检样品信息_抽样方式"] = FindFdval(sel4html, "抽样方式")
	mkr["抽检样品信息_样品形态"] = FindFdval(sel4html, "样品形态")
	mkr["抽检样品信息_样品包装"] = FindFdval(sel4html, "样品包装")
	mkr["抽检样品信息_抽样工具"] = FindFdval(sel4html, "抽样工具")
	mkr["抽检样品信息_抽样时样品储存条件"] = FindFdval(sel4html, "储存条件")
	mkr["抽检样品信息_抽样基数"] = FindFdval(sel4html, "抽样基数")
	mkr["抽检样品信息_抽样数量"] = FindFdval(sel4html, "抽样数量")
	mkr["抽检样品信息_备样数量"] = FindFdval(sel4html, "备样数量")
	mkr["抽检样品信息_抽样数量单位"] = FindFdval(sel4html, "抽样数量单位")
	mkr["抽检样品信息_执行标准/技术文件"] = FindFdval(sel4html, "执行标准/技术文件")
	mkr["抽检样品信息_备注"] = FindFdval(sel4html, "整个样品备注")

	return mkr
}

//填充报告
func Fill_item(new map[string]string, fooddetail map[string]string) {
	fooddetail["report_no"] = new["报告书编号"]
	fooddetail["jd_bz"] = new["监督抽检报告备注"]
	fooddetail["fx_bz"] = new["风险监测报告备注"]
	fooddetail["conclusion"] = new["结论"]
	fooddetail["report_type"] = new["报告类别"]
	fooddetail["test_conclusion"] = new["检验结论"]
}
func GetTestInfo(k string, testinfos []*Test_platform_api_food_getTestInfo_o) *Test_platform_api_food_getTestInfo_o {
	for _, it := range testinfos {
		if it.Spdata_0 == k {
			return it
		}
	}
	return nil
}

//填充检测项目
func Fill_subitem(news []map[string]string, testinfos []*Test_platform_api_food_getTestInfo_o) {
	for _, new := range news {
		info := GetTestInfo(new["检验项目"], testinfos)
		if info == nil {
			continue
		}
		info.Spdata_1 = new["检验结果"]
		info.Spdata_18 = new["结果单位"]
		info.Spdata_2 = new["结果判定"]
		info.Spdata_19 = new["检验依据"]
		info.Spdata_4 = new["判定依据"]
		info.Spdata_11 = new["最小允许限"]
		info.Spdata_15 = new["最大允许限"]
		info.Spdata_16 = new["允许限单位"]
		info.Spdata_7 = new["方法检出限"]
		info.Spdata_8 = new["检出限单位"]
		info.Spdata_17 = new["说明"]
		fmt.Println(new)
	}
}

//获取所有不合格
func Getunqualified(testinfos []map[string]string) []map[string]string {
	r := make([]map[string]string, 0)
	for _, it := range testinfos {
		if it["结果判定"] == "不合格项" {
			r = append(r, it)
		}
	}
	return r
}

//获取所有判定依据
func GetAllPandingyiju(testinfos []map[string]string) []string {
	r := make([]string, 0)
	for _, it := range testinfos {
		pd := false
		for _, rit := range r {
			if it["判定依据"] == rit {
				pd = true
				break
			}
		}
		if pd == false {
			r = append(r, it["判定依据"])
		}
	}
	return r
}

//生成指定判断依据
func Buildpanduanyiju(yj string, testinfos []map[string]string) string {
	r := ""
	for _, it := range testinfos {
		if it["结果判定"] == "不合格项" && it["判定依据"] == yj {
			r = fmt.Sprintf("%s%s", r, it["检验项目"])
			r = fmt.Sprintf("%s,", r)
		}
	}
	if r == "" {
		return ""
	}
	if string(r[len(r)-1]) == "," {
		r = r[:len(r)-1]
	}
	r = strings.ReplaceAll(r, ",", "，")
	r = fmt.Sprintf("%s项目不符合%s要求", r, yj)
	return r
}

//生成报告
func Buildbaogao(testinfos []map[string]string) string {
	unqualifieds := Getunqualified(testinfos)
	allpdyiju := GetAllPandingyiju(testinfos)
	if len(unqualifieds) == 0 {
		rs := "经抽样检验，所检项目符合 "
		for i, yj := range allpdyiju {
			rs = fmt.Sprintf("%s%s", rs, yj)
			if len(allpdyiju)-1 != i {
				rs = fmt.Sprintf("%s，", rs)
			}
		}
		rs = fmt.Sprintf("%s要求。", rs)
		return rs
	}
	rs := "经抽样检验，"
	for _, unq := range allpdyiju {
		sps := Buildpanduanyiju(unq, unqualifieds)
		if sps != "" {
			rs = fmt.Sprintf("%s%s", rs, sps)
			rs = fmt.Sprintf("%s,", rs)
		}
	}
	if string(rs[len(rs)-1]) == "," {
		rs = rs[:len(rs)-1]
	}
	rs = strings.ReplaceAll(rs, ",", "，")
	rs = fmt.Sprintf("%s，检验结论为不合格。", rs)
	return rs
}
