package nifdc

import (
	"encoding/gob"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ahmetb/go-linq/v3"
	"github.com/json-iterator/go"
	"os"
	"strings"
)
type JudgmentbasisItem struct {
	Name string
	Items []map[string]interface{}
}

func Get_mapsitem(key string, items []map[string]string) map[string]string {
	for _, it := range items {
		//fmt.Println(it["items"],"=>",key)
		if it["item"] == key {

			return it
		}
	}
	return nil
}
func MergeUpdates(updates []map[string]string, remotes []map[string]string) []map[string]string {
	for _, remote := range remotes {
		xm := remote["检验项目"]
		xm = strings.ReplaceAll(xm, "（", "(")
		xm = strings.ReplaceAll(xm, "）", ")")
		update := Get_mapsitem(xm, updates)
		if update != nil {
			update["sp_data_1"] = remote["检验结果"]
			update["sp_data_2"] = remote["结果判定"]
			update["sp_data_17"] = remote["说明"]
		}
	}
	return updates
}

//合并检验项目
func Merge_subitem(items []map[string]string, subitems []map[string]string) {
	for _, subit := range subitems {
		tmpitm := Get_item(subit["检验项目"], items)
		tmpitm["检验结果*"] = subit["检验结果"]
		tmpitm["结果单位*"] = subit["结果单位"]
		tmpitm["结果判定*"] = subit["结果判定"]
		tmpitm["检验依据*"] = subit["检验依据"]
		tmpitm["判定依据*"] = subit["判定依据"]
		tmpitm["最小允许限*"] = subit["最小允许限"]
		tmpitm["最大允许限*"] = subit["最大允许限"]
		tmpitm["允许限单位*"] = subit["允许限单位"]
		tmpitm["方法检出限*"] = subit["方法检出限"]
		tmpitm["检出限单位*"] = subit["检出限单位"]
		tmpitm["备注"] = subit["备注"]
		tmpitm["说明"] = subit["说明"]
	}
}
func Get_item(k string, items []map[string]string) map[string]string {
	for _, it := range items {
		if it["检验项目*"] == k {
			return it
		}
	}
	return nil
}

//转换检测项目
func TestInfotoMap(tos []*Test_platform_api_food_getTestInfo_o, ios []*Test_platform_api_food_getTestItems_o) []map[string]string {
	r := make([]map[string]string, 0)
	for _, io := range ios {
		subr := make(map[string]string)
		subr["检验项目*"] = io.Item
		subr["检验结果*"] = ""
		subr["结果单位*"] = io.TestReason[0].Spdata_18
		subr["结果判定*"] = "未检验"
		subr["检验依据*"] = io.TestReason[0].Sm
		subr["判定依据*"] = io.VerifyReason[0].Spdata_4
		subr["最小允许限*"] = io.VerifyReason[0].Spdata_9
		subr["最大允许限*"] = io.VerifyReason[0].Spdata_13
		subr["允许限单位*"] = io.VerifyReason[0].Spdata_10
		subr["方法检出限*"] = io.TestReason[0].Spdata_5
		subr["检出限单位*"] = io.VerifyReason[0].Spdata_14
		subr["备注"] = io.TestReason[0].Bz
		subr["说明"] = ""
		r = append(r, subr)
	}

	for i, to := range tos {
		if len(r) <= i {
			subr := make(map[string]string)
			r = append(r, subr)
		}
		subr := r[i]
		subr["检验项目*"] = to.Spdata_0
		subr["检验结果*"] = to.Spdata_1
		subr["结果单位*"] = to.Spdata_18
		subr["结果判定*"] = to.Spdata_2
		subr["检验依据*"] = to.Spdata_19
		subr["判定依据*"] = to.Spdata_4
		subr["最小允许限*"] = to.Spdata_11
		subr["最大允许限*"] = to.Spdata_15
		subr["允许限单位*"] = to.Spdata_16
		subr["方法检出限*"] = to.Spdata_7
		subr["检出限单位*"] = to.Spdata_8
		subr["备注"] = to.Spdata_20
		subr["说明"] = to.Spdata_17
	}
	return r
}

//转换检测模块
func StoMap_test_platform(s string) map[string]string {
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	mkr := make(map[string]string, 0)

	rt.Find(".ibox").Each(func(i int, sel *goquery.Selection) {
		title := sel.Find("h2").First().Text()
		sel.Find(".col-sm-4:contains(：)").Each(func(i int, selection *goquery.Selection) {
			sptxt := strings.TrimSpace(strings.ReplaceAll(selection.Text(), "\n", ""))
			spsel := strings.Split(sptxt, "：")
			if len(spsel) != 2 {
				return
			}
			k := fmt.Sprintf("%s_%s", title, strings.TrimSpace(spsel[0]))
			//fmt.Println(k)
			v := strings.TrimSpace(spsel[1])
			//v = strings.ReplaceAll(v, "                                        ", " ")
			if mkr[k] == "" {
				mkr[k] = v
			}
		})
		sel.Find(".col-sm-4:contains(\\:)").Each(func(i int, selection *goquery.Selection) {
			sptxt := strings.TrimSpace(strings.ReplaceAll(selection.Text(), "\n", ""))
			spsel := strings.Split(sptxt, ":")
			if len(spsel) != 2 {
				return
			}
			k := fmt.Sprintf("%s_%s", title, strings.TrimSpace(spsel[0]))
			//fmt.Println(k)
			v := strings.TrimSpace(spsel[1])
			//v = strings.ReplaceAll(v, "                                        ", " ")
			if mkr[k] == "" {
				mkr[k] = v
			}
		})
	})

	mkr["检验结论"] = strings.TrimSpace(rt.Find("#testform").Find("h3:contains(检验结论)").Parent().Find("p").Text())
	mkr["状态"] = strings.TrimSpace(strings.ReplaceAll(rt.Find("h2:contains(状态\\:)").Text(), "状态:", ""))
	if mkr["状态"] == "" && mkr["检验结论"] == "" {
		rt.Find(".ibox").Each(func(i int, sel *goquery.Selection) {
			title := sel.Find("h2").First().Text()
			if title == "检验信息" {
				sel.Find(".col-sm-4:contains(：)").Each(func(i int, selection *goquery.Selection) {

					sptxt := strings.TrimSpace(strings.ReplaceAll(selection.Text(), "\n", ""))
					spsel := strings.Split(sptxt, "：")
					selval := selection.Find("*[value],textarea")
					k := fmt.Sprintf("%s_%s", title, strings.TrimSpace(spsel[0]))
					//fmt.Println(k)
					v := strings.TrimSpace(selval.AttrOr("value", ""))
					if v == "" {
						v = strings.TrimSpace(selval.Text())
					}
					mkr[k] = v
					if k=="检验信息_结论"&&v==""{
						mkr[k]="未检验样品"
					}
					//fmt.Printf("%s=>%s\n", k, v)
				})
			}
		})
		mkr["检验结论"] = strings.TrimSpace(rt.Find("#test_conclusion").Text())
	}
	////抽样基础信息
	//sel0 := rt.Find(".ibox.float-e-margins").Eq(0)
	////抽样单位信息
	//sel1 := rt.Find(".ibox.float-e-margins").Eq(1)
	////抽检场所信息
	//sel2 := rt.Find(".ibox.float-e-margins").Eq(2)
	////生产企业信息
	//sel3 := rt.Find(".ibox.float-e-margins").Eq(3)
	////抽检样品信息
	//sel4 := rt.Find(".ibox.float-e-margins").Eq(4)
	//////照片信息
	////sel5 := rt.Find(".ibox.float-e-margins").Eq(5)
	////sel5html, _ := sel5.Html()
	////检验信息
	//sel6 := rt.Find(".ibox.float-e-margins").Eq(6)
	////sel6html, _ := sel6.Html()
	//mkr := make(map[string]string, 0)
	//
	//mkr["抽样基础信息_任务来源"] = FindNextNodeVal(sel0, "任务来源")
	//mkr["抽样基础信息_报送分类A"] = FindNextNodeVal(sel0, "报送分类A")
	//mkr["抽样基础信息_报送分类B"] = FindNextNodeVal(sel0, "报送分类B")
	//mkr["抽样基础信息_食品大类"] = FindNextNodeVal(sel0, "食品大类")
	//mkr["抽样基础信息_食品亚类"] = FindNextNodeVal(sel0, "食品亚类")
	//mkr["抽样基础信息_食品次亚类"] = FindNextNodeVal(sel0, "食品次亚类")
	//mkr["抽样基础信息_食品细类"] = FindNextNodeVal(sel0, "食品细类")
	//mkr["抽样基础信息_抽样单编号"] = FindNextNodeVal(sel0, "抽样单编号")
	//mkr["抽样基础信息_检验目的/任务类别"] = FindNextNodeVal(sel0, "检验目的\\/任务类别")
	//mkr["抽样基础信息_抽样类型"] = FindNextNodeVal(sel0, "抽样类型")
	//
	//mkr["抽样单位信息_抽样单位名称"] = FindNextNodeVal(sel1, "抽样单位名称")
	//mkr["抽样单位信息_单位地址"] = FindNextNodeVal(sel1, "单位地址")
	//mkr["抽样单位信息_所在省份"] = FindNextNodeVal(sel1, "所在省份")
	//mkr["抽样单位信息_抽样人员"] = FindNextNodeVal(sel1, "抽样人员")
	//mkr["抽样单位信息_抽样人员电话"] = FindNextNodeVal(sel1, "抽样人员电话")
	//mkr["抽样单位信息_单位联系人"] = FindNextNodeVal(sel1, "单位联系人")
	//mkr["抽样单位信息_传真"] = FindNextNodeVal(sel1, "传真")
	//mkr["抽样单位信息_邮编"] = FindNextNodeVal(sel1, "邮编")
	//mkr["抽样单位信息_联系人电话"] = FindNextNodeVal(sel1, "联系人电话")
	//
	//mkr["抽检场所信息_单位名称"] = FindNextNodeVal(sel2, "单位名称")
	//mkr["抽检场所信息_单位地址"] = FindNextNodeVal(sel2, "单位地址")
	//mkr["抽检场所信息_所在地"] = FindNextNodeVal(sel2, "所在地")
	//mkr["抽检场所信息_区域类型"] = FindNextNodeVal(sel2, "区域类型")
	//mkr["抽检场所信息_抽样环节"] = FindNextNodeVal(sel2, "抽样环节")
	//mkr["抽检场所信息_抽样地点"] = FindNextNodeVal(sel2, "抽样地点")
	//mkr["抽检场所信息_营业执照/社会信用代码"] = FindNextNodeVal(sel2, "营业执照\\/社会信用代码")
	//mkr["抽检场所信息_许可证类型"] = FindNextNodeVal(sel2, "许可证类型")
	//mkr["抽检场所信息_许可证号"] = FindNextNodeVal(sel2, "许可证号")
	//mkr["抽检场所信息_单位法人"] = FindNextNodeVal(sel2, "单位法人")
	//mkr["抽检场所信息_联系人"] = FindNextNodeVal(sel2, "联系人")
	//mkr["抽检场所信息_联系人电话"] = FindNextNodeVal(sel2, "联系人电话")
	//mkr["抽检场所信息_传真"] = FindNextNodeVal(sel2, "传真")
	//mkr["抽检场所信息_邮编"] = FindNextNodeVal(sel2, "邮编")
	//mkr["抽检场所信息_年销售额"] = FindNextNodeVal(sel2, "年销售额")
	//
	//mkr["生产企业信息_企业名称"] = FindNextNodeVal(sel3, "企业名称")
	//mkr["生产企业信息_所在地"] = FindNextNodeVal(sel3, "所在地")
	//mkr["生产企业信息_企业地址"] = FindNextNodeVal(sel3, "企业地址")
	//mkr["生产企业信息_生产许可证编号"] = FindNextNodeVal(sel3, "生产许可证编号")
	//mkr["生产企业信息_联系人"] = FindNextNodeVal(sel3, "联系人")
	//mkr["生产企业信息_电话"] = FindNextNodeVal(sel3, "电话")
	//mkr["生产企业信息_是否存在第三方企业信息"] = FindNextNodeVal(sel3, "是否存在第三方企业信息")
	//mkr["生产企业信息_第三方联系人"] = FindNextNodeVal(sel3, "第三方联系人")
	//mkr["生产企业信息_第三方电话"] = FindNextNodeVal(sel3, "第三方电话")
	//mkr["生产企业信息_第三方所在省(市、自治区)*"] = FindNextNodeVal(sel3, "第三方所在省(市\\、自治区)\\*")
	//mkr["生产企业信息_第三方所在地区(市、州、盟)"] = FindNextNodeVal(sel3, "第三方所在地区(市\\、州\\、盟)")
	//mkr["生产企业信息_第三方所在县(市、区)"] = FindNextNodeVal(sel3, "第三方所在县(市\\、区)")
	//mkr["生产企业信息_第三方企业名称"] = FindNextNodeVal(sel3, "第三方企业名称")
	//mkr["生产企业信息_第三方企业地址"] = FindNextNodeVal(sel3, "第三方企业地址")
	//mkr["生产企业信息_第三方生产许可证编号"] = FindNextNodeVal(sel3, "第三方生产许可证编号")
	//mkr["生产企业信息_第三方性质"] = FindNextNodeVal(sel3, "第三方性质")
	//
	//mkr["抽检样品信息_样品条码"] = FindNextNodeVal(sel4, "样品条码")
	//mkr["抽检样品信息_是否进口"] = FindNextNodeVal(sel4, "是否进口")
	//mkr["抽检样品信息_样品商标"] = FindNextNodeVal(sel4, "样品商标")
	//mkr["抽检样品信息_样品类型"] = FindNextNodeVal(sel4, "样品类型")
	//mkr["抽检样品信息_样品来源"] = FindNextNodeVal(sel4, "样品来源")
	//mkr["抽检样品信息_样品属性"] = FindNextNodeVal(sel4, "样品属性")
	//mkr["抽检样品信息_包装规格"] = FindNextNodeVal(sel4, "包装规格")
	//mkr["抽检样品信息_样品名称"] = FindNextNodeVal(sel4, "样品名称")
	//mkr["抽检样品信息_生产日期"] = FindNextNodeVal(sel4, "生产日期")
	//mkr["抽检样品信息_购进日期"] = FindNextNodeVal(sel4, "购进日期")
	//mkr["抽检样品信息_加工日期"] = FindNextNodeVal(sel4, "加工日期")
	//
	//mkr["抽检样品信息_保质期"] = FindNextNodeVal(sel4, "保质期")
	//mkr["抽检样品信息_样品批号"] = FindNextNodeVal(sel4, "样品批号")
	//mkr["抽检样品信息_样品规格"] = FindNextNodeVal(sel4, "样品规格")
	//mkr["抽检样品信息_质量等级"] = FindNextNodeVal(sel4, "质量等级")
	//mkr["抽检样品信息_单价"] = FindNextNodeVal(sel4, "单价")
	//mkr["抽检样品信息_是否进口"] = FindNextNodeVal(sel4, "是否进口")
	//mkr["抽检样品信息_原产地"] = FindNextNodeVal(sel4, "原产地")
	//mkr["抽检样品信息_抽样日期"] = FindNextNodeVal(sel4, "抽样日期")
	//mkr["抽检样品信息_抽样方式"] = FindNextNodeVal(sel4, "抽样方式")
	//mkr["抽检样品信息_抽样时样品储存条件"] = FindNextNodeVal(sel4, "抽样时样品储存条件")
	//mkr["抽检样品信息_抽样基数"] = FindNextNodeVal(sel4, "抽样基数")
	//mkr["抽检样品信息_抽样数量"] = FindNextNodeVal(sel4, "抽样数量")
	//mkr["抽检样品信息_备样数量"] = FindNextNodeVal(sel4, "备样数量")
	//mkr["抽检样品信息_抽样数量单位"] = FindNextNodeVal(sel4, "抽样数量单位")
	//mkr["抽检样品信息_执行标准/技术文件"] = FindNextNodeVal(sel4, "执行标准\\/技术文件")
	//mkr["抽检样品信息_备注"] = FindNextNodeVal(sel4, "备注")
	//mkr["抽检样品信息_批准文号"] = FindNextNodeVal(sel4, "批准文号")
	//
	//mkr["检验信息_检验机构名称"] = FindNextNodeVal(sel6, "检验机构名称")
	//mkr["检验信息_报告书编号"] = FindNextNodeVal(sel6, "报告书编号")
	//mkr["检验信息_样品到达日期"] = FindNextNodeVal(sel6, "样品到达日期")
	//mkr["检验信息_联系人"] = FindNextNodeVal(sel6, "联系人")
	//mkr["检验信息_联系人电话"] = FindNextNodeVal(sel6, "联系人电话")
	//mkr["检验信息_联系人邮箱"] = FindNextNodeVal(sel6, "联系人邮箱")
	//mkr["检验信息_检查封样人员"] = FindNextNodeVal(sel6, "检查封样人员")
	//mkr["检验信息_检查封样人电话"] = FindNextNodeVal(sel6, "检查封样人电话")
	//mkr["检验信息_检查封样人邮箱"] = FindNextNodeVal(sel6, "检查封样人邮箱")
	//mkr["检验信息_结论"] = FindNextNodeVal(sel6, "结论")
	//mkr["检验信息_监督抽检报告备注"] = FindNextNodeVal(sel6, "监督抽检报告备注")
	//mkr["检验信息_风险监测报告备注"] = FindNextNodeVal(sel6, "风险监测报告备注")
	//mkr["检验信息_复检状态"] = FindNextNodeVal(sel6, "复检状态")
	//
	//mkr["检验结论"] = strings.TrimSpace(rt.Find("#testform").Find("h3:contains(检验结论)").Parent().Find("p").Text())
	//mkr["状态"] = strings.TrimSpace(strings.ReplaceAll(rt.Find("h2:contains(状态\\:)").Text(), "状态:", ""))
	//
	//if mkr["状态"] != "" {
	//	mkr["检验信息_检验目的/任务类别"] = FindNextNodeVal(sel6, "检验目的\\/任务类别")
	//	mkr["检验信息_报告类别"] = FindNextNodeVal(sel6, "报告类别")
	//} else {
	//	mkr["检验信息_检验目的/任务类别"] = sel6.Find("label:contains(检验目的\\/任务类别：)").Next().Find("option[selected=selected]").Text()
	//	mkr["检验信息_报告类别"] = sel6.Find("label:contains(报告类别：)").Next().Find("option").First().Text()
	//}
	return mkr
}

//转换任务大平台
func StoMap_sample(s string) map[string]string {
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	//抽样基础信息
	//sel0 := rt.Find(".formAreaCom").Eq(0)
	//sel0html, _ := sel0.Html()
	////抽样单位信息
	//sel1 := rt.Find(".formAreaCom").Eq(1)
	//sel1html, _ := sel1.Html()
	////抽样场所信息
	//sel2 := rt.Find(".formAreaCom").Eq(2)
	//sel2html, _ := sel2.Html()
	////抽样生产企业信
	//sel3 := rt.Find(".formAreaCom").Eq(3)
	//sel3html, _ := sel3.Html()
	////抽检样品信息
	//sel4 := rt.Find(".formAreaCom").Eq(4)
	//sel4html, _ := sel4.Html()
	////检验信息
	//sel7 := rt.Find(".formAreaCom").Eq(7)
	//sel7html, _ := sel7.Html()

	mkr := make(map[string]string, 0)

	rt.Find(".formAreaCom").Each(func(i int, sel *goquery.Selection) {
		title := sel.Find(".areaTitle").First().Text()
		if i == 7 && title == "" {
			title = "检验信息"
		}
		sel.Find("div:contains(：)").Each(func(i int, selection *goquery.Selection) {
			sptxt := strings.TrimSpace(strings.ReplaceAll(selection.Text(), "\n", ""))
			spsel := strings.Split(sptxt, "：")

			if len(spsel) != 2 {
				return
			}
			k := fmt.Sprintf("%s_%s", title, strings.TrimSpace(spsel[0]))
			//fmt.Println(k)
			v := strings.TrimSpace(spsel[1])
			v = strings.ReplaceAll(v, "                                        ", " ")
			if mkr[k] == "" {
				mkr[k] = v
			}
		})
		sel.Find("div:contains(\\:)").Each(func(i int, selection *goquery.Selection) {
			sptxt := strings.TrimSpace(strings.ReplaceAll(selection.Text(), "\n", ""))
			spsel := strings.Split(sptxt, ":")

			if len(spsel) != 2 {
				return
			}
			k := fmt.Sprintf("%s_%s", title, strings.TrimSpace(spsel[0]))
			//fmt.Println(k)
			v := strings.TrimSpace(spsel[1])
			v = strings.ReplaceAll(v, "                                        ", " ")
			if mkr[k] == "" {
				mkr[k] = v
			}
		})
	})

	//mkr["抽样基础信息_任务来源"] = FindFdval(sel0html, "任务来源")
	//mkr["抽样基础信息_报送分类"] = FindFdval(sel0html, "报送分类")
	//mkr["抽样基础信息_检验机构名称"] = FindFdval(sel0html, "检验机构名称")
	//mkr["抽样基础信息_部署机构"] = FindFdval(sel0html, "部署机构")
	//mkr["抽样基础信息_抽样类型"] = FindFdval(sel0html, "抽样类型")
	//mkr["抽样基础信息_抽样环节"] = FindFdval(sel0html, "抽样环节")
	//mkr["抽样基础信息_抽样地点"] = FindFdval(sel0html, "抽样地点")
	//mkr["抽样基础信息_食品分类"] = strings.ReplaceAll(FindFdval(sel0html, "食品分类"), "                                    ", " ")
	//mkr["抽样基础信息_抽样单编号"] = FindFdval(sel0html, "抽样单编号")
	//mkr["抽样基础信息_检验目的/任务类别"] = FindFdval(sel0html, "检验目的\\/任务类别")
	//mkr["抽样单位信息_单位名称"] = FindFdval(sel1html, "单位名称")
	//mkr["抽样单位信息_单位地址"] = FindFdval(sel1html, "单位地址")
	//mkr["抽样单位信息_所在省份"] = FindFdval(sel1html, "所在省份")
	//mkr["抽样单位信息_抽样人员"] = FindFdval(sel1html, "抽样人员")
	//mkr["抽样单位信息_联系人"] = FindFdval(sel1html, "联系人")
	//mkr["抽样单位信息_电子邮箱"] = FindFdval(sel1html, "电子邮箱")
	//mkr["抽样单位信息_电话"] = FindFdval(sel1html, "电话")
	//mkr["抽样单位信息_传真"] = FindFdval(sel1html, "传真")
	//mkr["抽检场所信息_所在地"] = strings.ReplaceAll(FindFdval(sel2html, "所在地"), "                                        ", " ")
	//mkr["抽检场所信息_区域类型"] = FindFdval(sel2html, "区域类型")
	//mkr["抽检场所信息_单位名称"] = FindFdval(sel2html, "单位名称")
	//mkr["抽检场所信息_单位地址"] = FindFdval(sel2html, "单位地址")
	//mkr["抽检场所信息_营业执照/社会信用代码"] = FindFdval(sel2html, "营业执照\\/社会信用代码")
	//mkr["抽检场所信息_许可证类型"] = FindFdval(sel2html, "许可证类型")
	//mkr["抽检场所信息_许可证号"] = FindFdval(sel2html, "许可证号")
	//mkr["抽检场所信息_法人代表"] = FindFdval(sel2html, "法人代表")
	//mkr["抽检场所信息_联系人"] = FindFdval(sel2html, "联系人")
	//mkr["抽检场所信息_联系电话"] = FindFdval(sel2html, "联系电话")
	//mkr["抽检场所信息_摊位号或姓名"] = FindFdval(sel2html, "摊位号或姓名")
	//mkr["抽检场所信息_身份证号"] = FindFdval(sel2html, "身份证号")
	//mkr["抽样生产企业信息_所在地"] = strings.ReplaceAll(FindFdval(sel3html, "所在地"), "                                        ", " ")
	//mkr["抽样生产企业信息_生产者地址"] = FindFdval(sel3html, "生产者地址")
	//mkr["抽样生产企业信息_生产者名称"] = FindFdval(sel3html, "生产者名称")
	//mkr["抽样生产企业信息_生产许可证编号"] = FindFdval(sel3html, "生产许可证编号")
	//mkr["抽样生产企业信息_联系电话"] = FindFdval(sel3html, "联系电话")
	//mkr["抽样生产企业信息_是否存在第三方企业信息"] = FindFdval(sel3html, "是否存在第三方企业信息")
	//if mkr["抽样生产企业信息_是否存在第三方企业信息"] == "是" {
	//	rhm, _ := sel3.Find(".wtsc.clearfix").Html()
	//	dsfdz := strings.ReplaceAll(FindFdval(rhm, "所在地"), "                                                ", " ")
	//	spdsfdz := strings.Split(dsfdz, " ")
	//	if len(spdsfdz) == 3 {
	//		mkr["抽样生产企业信息_第三方企业省份"] = spdsfdz[0]
	//		mkr["抽样生产企业信息_第三方企业市区"] = spdsfdz[1]
	//		mkr["抽样生产企业信息_第三方企业县区"] = spdsfdz[2]
	//	}
	//	mkr["抽样生产企业信息_第三方企业地址"] = FindFdval(sel3html, "第三方企业地址")
	//	mkr["抽样生产企业信息_第三方企业名称"] = FindFdval(sel3html, "第三方企业名称")
	//	mkr["抽样生产企业信息_第三方企业许可证编号"] = FindFdval(sel3html, "第三方企业许可证编号")
	//	mkr["抽样生产企业信息_第三方企业联系人"] = ""
	//	mkr["抽样生产企业信息_第三方企业电话"] = FindFdval(sel3html, "第三方企业电话")
	//	mkr["抽样生产企业信息_第三方企业性质"] = FindFdval(sel3html, "第三方企业性质")
	//}
	//mkr["抽检样品信息_样品条码"] = FindFdval(sel4html, "样品条码")
	//mkr["抽检样品信息_样品商标"] = FindFdval(sel4html, "样品商标")
	//mkr["抽检样品信息_样品类型"] = FindFdval(sel4html, "样品类型")
	//mkr["抽检样品信息_样品来源"] = FindFdval(sel4html, "样品来源")
	//mkr["抽检样品信息_样品属性"] = FindFdval(sel4html, "样品属性")
	//mkr["抽检样品信息_包装分类"] = FindFdval(sel4html, "包装分类")
	//mkr["抽检样品信息_样品名称"] = FindFdval(sel4html, "样品名称")
	//mkr["抽检样品信息_生产日期"] = FindFdval(sel4html, "生产日期")
	//mkr["抽检样品信息_购进日期"] = FindFdval(sel4html, "购进日期")
	//mkr["抽检样品信息_加工日期"] = FindFdval(sel4html, "加工日期")
	//mkr["抽检样品信息_保质期"] = FindFdval(sel4html, "保质期")
	//mkr["抽检样品信息_样品批号"] = FindFdval(sel4html, "样品批号")
	//mkr["抽检样品信息_规格型号"] = FindFdval(sel4html, "规格型号")
	//mkr["抽检样品信息_质量等级"] = FindFdval(sel4html, "质量等级")
	//mkr["抽检样品信息_单价"] = FindFdval(sel4html, "单价")
	//mkr["抽检样品信息_是否进口"] = FindFdval(sel4html, "是否进口")
	//mkr["抽检样品信息_抽样日期"] = FindFdval(sel4html, "抽样日期")
	//mkr["抽检样品信息_抽样方式"] = FindFdval(sel4html, "抽样方式")
	//mkr["抽检样品信息_原产地"] = FindFdval(sel4html, "原产地")
	//mkr["抽检样品信息_储存条件"] = FindFdval(sel4html, "储存条件")
	//mkr["抽检样品信息_抽样基数"] = FindFdval(sel4html, "抽样基数")
	//mkr["抽检样品信息_抽样数量"] = FindFdval(sel4html, "抽样数量")
	//mkr["抽检样品信息_备样数量"] = FindFdval(sel4html, "备样数量")
	//mkr["抽检样品信息_抽样数量单位"] = FindFdval(sel4html, "抽样数量单位")
	//mkr["抽检样品信息_执行标准/技术文件"] = FindFdval(sel4html, "执行标准\\/技术文件")
	//mkr["抽检样品信息_整个样品备注"] = FindFdval(sel4html, "整个样品备注")
	//
	//mkr["检验信息_检验机构名称"] = FindFdval(sel7html, "检验机构名称")
	//mkr["检验信息_样品送达日期"] = FindFdval(sel7html, "样品送达日期")
	//mkr["检验信息_联系人"] = FindFdval(sel7html, "联系人")
	//mkr["检验信息_联系人电话"] = FindFdval(sel7html, "联系人电话")
	//mkr["检验信息_联系人邮箱"] = FindFdval(sel7html, "联系人邮箱")
	//mkr["检验信息_检查封样人员"] = FindFdval(sel7html, "检查封样人员")
	//mkr["检验信息_系统接样日期"] = FindFdval(sel7html, "系统接样日期")
	//mkr["检验信息_检查封样人电话"] = FindFdval(sel7html, "检查封样人电话")
	//mkr["检验信息_检查封样人邮箱"] = FindFdval(sel7html, "检查封样人邮箱")

	return mkr
}

//填充报告
//func Fill_item(new map[string]string, fooddetail map[string]string) {
//	fooddetail["report_no"] = new["报告书编号"]
//	fooddetail["jd_bz"] = new["监督抽检报告备注"]
//	fooddetail["fx_bz"] = new["风险监测报告备注"]
//	fooddetail["conclusion"] = new["结论"]
//	fooddetail["report_type"] = new["报告类别"]
//	fooddetail["test_conclusion"] = new["检验结论"]
//}
//func GetTestInfo(k string, testinfos []*Test_platform_api_food_getTestInfo_o) *Test_platform_api_food_getTestInfo_o {
//	for _, it := range testinfos {
//		if it.Spdata_0 == k {
//			return it
//		}
//	}
//	return nil
//}

//填充检测项目
//func Fill_subitem(news []map[string]string, testinfos []*Test_platform_api_food_getTestInfo_o) {
//	for _, new := range news {
//		info := GetTestInfo(new["检验项目"], testinfos)
//		if info == nil {
//			continue
//		}
//		info.Spdata_1 = new["检验结果"]
//		info.Spdata_18 = new["结果单位"]
//		info.Spdata_2 = new["结果判定"]
//		info.Spdata_19 = new["检验依据"]
//		info.Spdata_4 = new["判定依据"]
//		info.Spdata_11 = new["最小允许限"]
//		info.Spdata_15 = new["最大允许限"]
//		info.Spdata_16 = new["允许限单位"]
//		info.Spdata_7 = new["方法检出限"]
//		info.Spdata_8 = new["检出限单位"]
//		info.Spdata_17 = new["说明"]
//		fmt.Println(new)
//	}
//}

//获取报告类别
func GetReport_type(items jsoniter.Any)string{
	report_type:="合格报告"
	for i := 0; i < items.Size(); i++ {
		it:=items.Get(i)
		if it.Get("结果判定").ToString()=="不合格项"{
			report_type="一般不合格报告"
			break
		}
	}
	return report_type
}
//获取结论
func GetConclusion(items jsoniter.Any)string{
	jielun:="纯抽检合格样品"
	for i := 0; i < items.Size(); i++ {
		it:=items.Get(i)
		if it.Get("结果判定").ToString()=="不合格项"{
			jielun="纯抽检不合格样品"
			break
		}
	}
	return jielun
}
//获取所有不合格
//func Getunqualified(testinfos []map[string]string) []map[string]string {
//	r := make([]map[string]string, 0)
//	for _, it := range testinfos {
//		if it["结果判定"] == "不合格项" {
//			r = append(r, it)
//		}
//	}
//	return r
//}
//func Panding(yj string, testinfos []map[string]string) bool {
//	for _, it := range testinfos {
//		if it["sp_data_2"] == "不合格项" {
//			return false
//		}
//	}
//	return true
//}

//获取所有判定依据
//func GetAllPandingyiju(testinfos []map[string]string) ([]string, []string) {
//	allyiju := make(map[string]string, 0)
//	for _, it := range testinfos {
//		if it["sp_data_2"] != "未检验" {
//			allyiju[it["sp_data_4"]] = ""
//		}
//	}
//	okyiju := []string{}
//	erryiju := []string{}
//	for yiju, _ := range allyiju {
//		if Panding(yiju, testinfos) == true {
//			okyiju = append(okyiju, yiju)
//		} else {
//			erryiju = append(erryiju, yiju)
//		}
//	}
//	return okyiju, erryiju
//}

////生成指定判断依据
//func Buildpanduanyiju(yj string, testinfos []map[string]string) string {
//	r := ""
//	for _, it := range testinfos {
//		if it["sp_data_2"] == "不合格项" && it["sp_data_4"] == yj {
//			r = fmt.Sprintf("%s%s", r, it["item"])
//			r = fmt.Sprintf("%s,", r)
//		}
//	}
//	if r == "" {
//		return ""
//	}
//	if string(r[len(r)-1]) == "," {
//		r = r[:len(r)-1]
//	}
//	r = strings.ReplaceAll(r, ",", "，")
//	r = fmt.Sprintf("%s项目不符合%s要求", r, yj)
//	return r
//}

//func Convbaotaodata(testinfos []map[string]string) []map[string]string {
//	r := make([]map[string]string, 0)
//	for _, item := range testinfos {
//		ditem := map[string]string{}
//		for k, v := range item {
//			ditem[k] = v
//			if k == "sp_data_2" {
//				ditem["结果判定"] = v
//			} else if k == "sp_data_4" {
//				ditem["判定依据"] = v
//			}
//		}
//		r = append(r, ditem)
//	}
//	return r
//}


//获取所有判定依据
func QueryJudgmentbasisItems(testinfos jsoniter.Any)[]string{
	allJudgmentbasis:=make([]string,0)
	for i := 0; i < testinfos.Size(); i++ {
		it:=testinfos.Get(i)
		jieguo:=it.Get("sp_data_2").ToString()
		pandingyiju:=it.Get("sp_data_4").ToString()
		if jieguo!="未检验"{
			n:=linq.From(allJudgmentbasis).Where(func(i interface{}) bool {
				lit:=i.(string)
				if pandingyiju==lit{
					return true
				}
				return false
			}).Count()
			if n==0{
				allJudgmentbasis=append(allJudgmentbasis,pandingyiju)
			}
		}
	}
	return allJudgmentbasis
}
//获取所有合格依据
func QueryQualifiedJudgmentbasisItems(testinfos jsoniter.Any)[]*JudgmentbasisItem{
	allyiju:=QueryJudgmentbasisItems(testinfos)
	jits:=make([]*JudgmentbasisItem,0)
	mptestinfos:=testinfos.GetInterface().([]map[string]interface{})
	for _, yijuit := range allyiju {
		qrs:=make([]map[string]interface{},0)
		linq.From(mptestinfos).Where(func(i interface{}) bool {
			it:=i.(map[string]interface{})
			jieguo:=it["sp_data_2"].(string)
			pandingyiju:=it["sp_data_4"].(string)
			if pandingyiju==yijuit{
				if jieguo=="合格项"{
					return true
				}
			}
			return false
		}).ToSlice(&qrs)
		jits=append(jits,&JudgmentbasisItem{
			Name: yijuit,
			Items: qrs,
		})
	}
	rjits:=make([]*JudgmentbasisItem,0)
	for _, it := range jits {
		if len(it.Items)!=0{
			rjits=append(rjits,it)
		}
	}
	return rjits
}
//获取所有不合格依据
func QueryUnqualifiedJudgmentbasisItems(testinfos jsoniter.Any)[]*JudgmentbasisItem{
	allyiju:=QueryJudgmentbasisItems(testinfos)
	jits:=make([]*JudgmentbasisItem,0)
	mptestinfos:=testinfos.GetInterface().([]map[string]interface{})
	for _, yijuit := range allyiju {
		qrs:=make([]map[string]interface{},0)
		linq.From(mptestinfos).Where(func(i interface{}) bool {
			it:=i.(map[string]interface{})
			jieguo:=it["sp_data_2"].(string)
			pandingyiju:=it["sp_data_4"].(string)
			if pandingyiju==yijuit{
				if jieguo=="不合格项"{
					return true
				}
			}
			return false
		}).ToSlice(&qrs)
		jits=append(jits,&JudgmentbasisItem{
			Name: yijuit,
			Items: qrs,
		})
	}

	rjits:=make([]*JudgmentbasisItem,0)
	for _, it := range jits {
		if len(it.Items)!=0{
			rjits=append(rjits,it)
		}
	}
	return rjits
}

//构建报告
func BuildTest_conclusion(testinfos jsoniter.Any) string {
	qualifieds:=QueryQualifiedJudgmentbasisItems(testinfos)
	unqualifieds:=QueryUnqualifiedJudgmentbasisItems(testinfos)
	results:=strings.Builder{}
	if len(unqualifieds)>0{
		results.WriteString("经抽样检验，")
		for i, unqualified := range unqualifieds {
			for idx, it := range unqualified.Items {
				results.WriteString(fmt.Sprintf("%s",it["item"].(string)))
				if idx!=len(unqualified.Items)-1{
					results.WriteString("，")
				}
			}
			results.WriteString(fmt.Sprintf("项目不符合%s要求",unqualified.Name))
			if i!=len(unqualifieds)-1{
				results.WriteString("，")
			}
		}
		results.WriteString("，检验结论为不合格。")
		fmt.Println(results.String())
	}else{
		results.WriteString("经抽样检验，所检项目符合 ")
		for i, qualified := range qualifieds {
			results.WriteString(fmt.Sprintf("%s",qualified.Name))
			if i!=len(qualifieds)-1{
				results.WriteString("，")
			}
		}
		results.WriteString(" 要求。")
	}
	return results.String()
}
//构建报告
//func Buildbaogao(testinfos []map[string]string) string {
//	okyiju, erryiju := GetAllPandingyiju(testinfos)
//	if len(okyiju) == 0 && len(erryiju) == 0 {
//		return ""
//	}
//	if len(erryiju) != 0 {
//		rs := "经抽样检验，"
//		for _, unq := range erryiju {
//			sps := Buildpanduanyiju(unq, testinfos)
//			if sps != "" {
//				rs = fmt.Sprintf("%s%s", rs, sps)
//				rs = fmt.Sprintf("%s,", rs)
//			}
//		}
//		if string(rs[len(rs)-1]) == "," {
//			rs = rs[:len(rs)-1]
//		}
//		rs = strings.ReplaceAll(rs, ",", "，")
//		rs = fmt.Sprintf("%s，检验结论为不合格。", rs)
//		return rs
//	} else {
//		rs := "经抽样检验，所检项目符合 "
//		for i, yj := range okyiju {
//			rs = fmt.Sprintf("%s%s", rs, yj)
//			if len(okyiju)-1 != i {
//				rs = fmt.Sprintf("%s，", rs)
//			}
//		}
//		rs = fmt.Sprintf("%s要求。", rs)
//		return rs
//	}
//
//	return ""
//}

func loop_testinfo(k string, testinfos []*Test_platform_api_food_getTestInfo_o) *Test_platform_api_food_getTestInfo_o {
	for _, it := range testinfos {
		if it.Spdata_0 == k {
			return it
		}
	}
	return nil
}
func loop_userdata(k string, userdata jsoniter.Any) jsoniter.Any {
	for i := 0; i < userdata.Size(); i++ {
		it:=userdata.Get(i)
		if it.Get("检验项目").ToString()==k{
			return it
		}
	}
	return nil
}
func loop_TestReason(k string, testreasons []*TestReason_o) *TestReason_o {
	for _, it := range testreasons {
		if strings.Contains(it.Spdata_3, k) {
			return it
		}
	}
	return nil
}
func loop_VerifyReason(k string, testreasons []*VerifyReason_o) *VerifyReason_o {
	for _, it := range testreasons {
		if strings.Contains(it.Spdata_4, k) {
			return it
		}
	}
	return nil
}

//合并上传数据
func Build_agriculture_updata(userdatas jsoniter.Any,testitems []*Test_platform_api_food_getTestItems_o, testinfos []*Test_platform_api_food_getTestInfo_o) jsoniter.Any {
	r := make([]map[string]interface{}, 0)
	for _, it := range testitems {
		itmap := make(map[string]interface{})
		itmap["id"] = ""
		itmap["item_old"] = it.Item
		itmap["item"] = it.Item
		itmap["sp_data_1"] = ""                            //结果
		itmap["sp_data_2"] = "未检验"                         //结果判定
		itmap["sp_data_3"] = it.TestReason[0].Spdata_3     //检验依据
		itmap["sp_data_4"] = it.VerifyReason[0].Spdata_4   //判定依据
		itmap["sp_data_5"] = it.TestReason[0].Spdata_5     //方法检出限
		itmap["sp_data_6"] = it.TestReason[0].Spdata_6     //结果单位
		itmap["sp_data_7"] = it.TestReason[0].Spdata_5     //方法检出限
		itmap["sp_data_8"] = it.TestReason[0].Spdata_18    //检出限单位
		itmap["sp_data_9"] = it.VerifyReason[0].Spdata_9   //最小允许限
		itmap["sp_data_10"] = it.VerifyReason[0].Spdata_10 //允许限单位
		itmap["sp_data_11"] = it.VerifyReason[0].Spdata_9  //最小允许限
		itmap["sp_data_12"] = it.TestReason[0].Spdata_6    //结果单位
		itmap["sp_data_13"] = it.VerifyReason[0].Spdata_13 //最大允许限
		itmap["sp_data_15"] = it.VerifyReason[0].Spdata_13 //最大允许限
		itmap["sp_data_16"] = it.VerifyReason[0].Spdata_10 //允许限单位
		itmap["sp_data_17"] = ""                           //说明
		itmap["sp_data_18"] = it.VerifyReason[0].Spdata_10 //结果单位
		itmap["bz"] = it.TestReason[0].Bz                  //备注
		itmap["sm"] = it.TestReason[0].Sm                  //检测依据简写
		itmap["sp_data_21"] = it.TestReason[0].Spdata_21   //spdata_21提示
		itmap["jylx"] = it.ItemType                        //抽检类型

		ittestinfo := loop_testinfo(it.Item, testinfos)
		if ittestinfo != nil {
			itmap["id"] = fmt.Sprintf("%d", ittestinfo.Id)
			itmap["sp_data_1"] = ittestinfo.Spdata_1 //结果
			itmap["sp_data_2"] = ittestinfo.Spdata_2 //结果判定
			itmap["sp_data_3"] = ittestinfo.Spdata_3 //检验依据
			itmap["sp_data_4"] = ittestinfo.Spdata_4 //判定依据
			itmap["sp_data_5"] = ittestinfo.Spdata_5 //方法检出限
			itmap["sp_data_6"] = ittestinfo.Spdata_6
			itmap["sp_data_7"] = ittestinfo.Spdata_7
			itmap["sp_data_8"] = ittestinfo.Spdata_8
			itmap["sp_data_9"] = ittestinfo.Spdata_9
			itmap["sp_data_10"] = ittestinfo.Spdata_10
			itmap["sp_data_11"] = ittestinfo.Spdata_11
			itmap["sp_data_12"] = ittestinfo.Spdata_12
			itmap["sp_data_13"] = ittestinfo.Spdata_13
			itmap["sp_data_15"] = ittestinfo.Spdata_15
			itmap["sp_data_16"] = ittestinfo.Spdata_16
			itmap["sp_data_17"] = ittestinfo.Spdata_17 //说明
			itmap["sp_data_18"] = ittestinfo.Spdata_18
			itmap["bz"] = ittestinfo.Spdata_20
			itmap["sm"] = ittestinfo.Spdata_19
			itmap["sp_data_21"] = ittestinfo.Spdata_21
			itmap["jylx"] = ittestinfo.Jylx
		}
		userdata := loop_userdata(it.Item, userdatas)
		if userdata != nil {
			jyff := userdata.Get("检验依据").ToString()
			pdyj := userdata.Get("判定依据").ToString()
			if jyff != "/" && jyff != "" {
				jyffo := loop_TestReason(jyff, it.TestReason)
				if jyffo != nil {
					itmap["sp_data_3"] = jyffo.Spdata_3   //检验依据
					itmap["sm"] = jyffo.Sm                //检测依据简写
					itmap["bz"] = jyffo.Bz                //检测依据简写
					itmap["sp_data_21"] = jyffo.Spdata_21 //检测依据简写
					itmap["sp_data_12"] = jyffo.Spdata_6  //结果单位
					itmap["sp_data_5"] = jyffo.Spdata_5   //方法检出限
					itmap["sp_data_6"] = jyffo.Spdata_6   //结果单位
					itmap["sp_data_7"] = jyffo.Spdata_5   //方法检出限
					itmap["sp_data_8"] = jyffo.Spdata_18  //检出限单位
					itmap["sp_data_18"] = jyffo.Spdata_18 //允许限单位
					itmap["sp_data_16"] = jyffo.Spdata_18 //允许限单位
					itmap["sp_data_10"] = jyffo.Spdata_18 //允许限单位
				}
			}
			if pdyj != "/" && jyff != "" {
				pdyjo := loop_VerifyReason(pdyj, it.VerifyReason)
				if pdyjo != nil {
					itmap["sp_data_4"] = pdyjo.Spdata_4   //判定依据
					itmap["sp_data_13"] = pdyjo.Spdata_13 //最大允许限
					itmap["sp_data_15"] = pdyjo.Spdata_13 //最大允许限
					itmap["sp_data_11"] = pdyjo.Spdata_9  //最小允许限
					itmap["sp_data_9"] = pdyjo.Spdata_9   //最小允许限
				}
			}

			itmap["sp_data_1"] = userdata.Get("检验结果").ToString() //结果
			itmap["sp_data_2"] = userdata.Get("结果判定").ToString() //结果判定
			itmap["sp_data_17"] = userdata.Get("说明").ToString()  //结果
			if userdata.Get("结果单位").ToString()!=""||userdata.Get("结果单位").ToString()!="/"{
				itmap["sp_data_18"]=userdata.Get("结果单位").ToString()
			}
		}
		r = append(r, itmap)
	}
	return jsoniter.Wrap(r)
}

func Savegob(i interface{}, fname string) error {
	of, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer of.Close()
	en := gob.NewEncoder(of)
	en.Encode(i)
	return nil
}
func Lavegob(i interface{}, fname string) error {
	of, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer of.Close()
	en := gob.NewDecoder(of)
	en.Decode(i)
	return nil
}
