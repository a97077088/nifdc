package nifdc

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	a, b, c, err := InitLoginck(nil)
	if err != nil {
		t.Fatal(err)
	}
	ck, err := Login("15738889730", "12345678", a, b, c, nil)
	if err != nil {
		t.Fatal(err)
	}
	test_platform_ck, err := Test_platform_login(ck, nil)
	if err != nil {
		t.Fatal(err)
	}
	fddetail, err := Test_platform_foodTest_foodDetail(10268179, test_platform_ck, nil)
	if err != nil {
		t.Fatal(err)
	}
	testinfo, err := Test_platform_api_food_getTestInfo(fddetail["sd"], test_platform_ck, nil)
	if err != nil {
		t.Fatal(err)
	}
	Fill_item(map[string]string{
		"报告书编号":    "w3c",
		"监督抽检报告备注": "监督抽检报告备注",
		"风险监测报告备注": "风险监测报告备注",
	}, fddetail)
	Fill_subitem([]map[string]string{
		{
			"检验项目":  "蛋白质",
			"结果单位":  "jkm",
			"检验结果":  "0.87",
			"结果判定":  "不合格项",
			"判定依据":  "GB 19645-2010《食品安全国家标准 巴氏杀菌乳》",
			"检验依据":  "GB 5009.5-2016(第一法)",
			"最小允许限": "0.1",
			"最大允许限": "1.0",
			"允许限单位": "km",
			"方法检出限": "wkm",
			"检出限单位": "jdm",
			"说明":    "测试",
		},
	}, testinfo.Rows)
	err = Test_platform_api_food_save(fddetail, testinfo.Rows, test_platform_ck, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(testinfo.Rows[0])
}
func TestRe(t *testing.T) {
	olds := `






<html>
<head>
    <title>普通食品检验填报</title>
     
 <link href="/test_platform/css/bootstrap.min.css?v=1.0.1" rel="stylesheet">

 <link href="/test_platform/css/font-awesome.min.css?v=4.4.0" rel="stylesheet">
 <link href="/test_platform/css/animate.min.css" rel="stylesheet">
 <link href="/test_platform/css/style.min.css?v=4.0.2" rel="stylesheet">
 <link href="/test_platform/css/plugins/toastr/toastr.min.css?v=4.0.2" rel="stylesheet">
 <link href="/test_platform/css/plugins/datapicker/datepicker3.css?v=4.0.2" rel="stylesheet">
 <link href="/test_platform/css/plugins/chosen/chosen.css?v=4.0.2" rel="stylesheet">
 <link rel="shortcut icon" href="/test_platform/favicon.ico">
 <link href="/test_platform/js/plugins/fancybox/jquery.fancybox.css" rel="stylesheet">
 <link href="/test_platform/layer/theme/default/layer.css" rel="stylesheet">

 <style>
  .fixed-table-body{
   height: auto !important;
  }
  .row input[type=text],.row select,.row textarea{
   font-size: 13px !important;
  }

 </style>

    <link rel="stylesheet" type="text/css" href="/test_platform/easyui/themes/metro/easyui.css">
    <link rel="stylesheet" type="text/css" href="/test_platform/easyui/themes/icon.css">
    <link rel="stylesheet" type="text/css" href="/test_platform/easyui/themes/color.css">
    <link rel="stylesheet" type="text/css" href="/test_platform/css/webuploader.css">
    <link rel="stylesheet" type="text/css" href="/test_platform/js/plugins/fancybox/viewer.min.css">
    <style>
        .tb_input {
            border: none;
            text-align: center;
            background: transparent;
            margin: 0 auto;
        }

        /*bootstrap兼容问题和easyui的bug*/
        .panel-header, .panel-body {

        }

        .datagrid, .combo-p {
            border: solid 1px #D4D4D4;
        }

        .datagrid * {
            -webkit-box-sizing: content-box;
            -moz-box-sizing: content-box;
            box-sizing: content-box;
        }
    </style>
</head>
<body>
<div class="wrapper wrapper-content   fadeInDown">
    <div class="container-fluid">
        <h1>普通食品检验数据填报
            <small class="text-danger">&nbsp;&nbsp;&nbsp;&nbsp; 抽样编号:NCP19131128463231006</small>
        </h1>
        <div class="panel panel-success">
            <div class="panel-heading">
                <i class="fa fa-info-circle"></i> 抽样信息内容
            </div>

            <div class="panel-body" style="padding: 15px;">
                
                <div class="row">
                    <div class="col-sm-12">
                        <div class="ibox float-e-margins">
                            <h2>抽样基础信息</h2>
                            <div class="hr-line-dashed"></div>
                            <div class="row form-group">
                                <div class="col-sm-4 ">
                                    <label class="control-label col-sm-4">任务来源：</label>
                                    <div class="col-sm-8">阜城县市场监督管理局</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">报送分类A：</label>
                                    <div class="col-sm-8" id="bsfla">抽检监测（县级本级）
                                        <input type="hidden" id="hid_bsfla" value="抽检监测（县级本级）"/>
                                    </div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">报送分类B：</label>
                                    <div class="col-sm-8" id="bsflb">2019年河北衡水阜城抽检计划1
                                        <input type="hidden" id="hid_bsflb" value="2019年河北衡水阜城抽检计划1"/>
                                    </div>
                                </div>
                            </div>
                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">食品大类：</label>
                                    <div class="col-sm-8" id="type1">食用农产品
                                        <input type="hidden" id="hid_type1" value="食用农产品"/>
                                    </div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">食品亚类：</label>
                                    <div class="col-sm-8" id="type2">蔬菜
                                        <input type="hidden" id="hid_type2" value="蔬菜"/>
                                    </div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">食品次亚类：</label>
                                    <div class="col-sm-8" id="type3">叶菜类蔬菜
                                        <input type="hidden" id="hid_type3" value="叶菜类蔬菜"/>
                                    </div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">食品细类：</label>
                                    <div class="col-sm-8" id="type4">油麦菜
                                        <input type="hidden" id="hid_type4" value="油麦菜"/>
                                    </div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样单编号：</label>
                                    <div class="col-sm-8">NCP19131128463231006</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样类型：</label>
                                    <div class="col-sm-8">农产品抽样</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                
                <div class="row">
                    <div class="col-sm-12 ">
                        <div class="ibox float-e-margins">
                            <h2>抽样单位信息</h2>
                            <div class="hr-line-dashed"></div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样单位名称：</label>
                                    <div class="col-sm-8">河南恒晟检测技术有限公司</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单位地址：</label>
                                    <div class="col-sm-8">新郑市薛店镇中德产业园2-3号</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">所在省份：</label>
                                    <div class="col-sm-8">河南</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样人员：</label>
                                    <div class="col-sm-8">河南恒晟检测技术有限公司</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样人员电话：</label>
                                    <div class="col-sm-8">0371-55929768</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单位联系人：</label>
                                    <div class="col-sm-8">王力</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                
                                
                                
                                
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">传真：</label>
                                    <div class="col-sm-8">0371-55929768</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">邮编：</label>
                                    <div class="col-sm-8">451162</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">联系人电话：</label>
                                    <div class="col-sm-8">0371-55929768</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                
                <div class="row">
                    <div class="col-sm-12 ">
                        <div class="ibox float-e-margins">
                            <h2>抽检场所信息</h2>
                            <div class="hr-line-dashed"></div>
                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单位名称：</label>
                                    <div class="col-sm-8">阜城县便利综合商店</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单位地址：</label>
                                    <div class="col-sm-8">蒋坊乡前蒋</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">所在地：</label>
                                    <div class="col-sm-8">河北/衡水/阜城</div>
                                </div>
                            </div>


                            
                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">区域类型：</label>
                                    <div class="col-sm-8">乡村</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样环节：</label>
                                    <div class="col-sm-8">流通</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样地点：</label>
                                    <div class="col-sm-8">小食杂店</div>
                                </div>

                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">营业执照/社会信用代码：</label>
                                    <div class="col-sm-8">131128600082059</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">许可证类型：</label>
                                    <div class="col-sm-8">经营许可证</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">许可证号：</label>
                                    <div class="col-sm-8">JY11311280001111</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单位法人：</label>
                                    <div class="col-sm-8">蒋国贤</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">联系人：</label>
                                    <div class="col-sm-8">蒋国贤</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">联系人电话：</label>
                                    <div class="col-sm-8">18333809141</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">传真：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">邮编：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">年销售额：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                
                <div class="row">
                    <div class="col-sm-12 ">
                        <div class="ibox float-e-margins">
                            <h2>生产企业信息</h2>
                            <div class="hr-line-dashed"></div>
                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">企业名称：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">所在地：</label>
                                    <div class="col-sm-8">河北/衡水/阜城</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">企业地址：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                            </div>
                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">生产许可证编号：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">联系人：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">电话：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                            </div>
                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">是否存在第三方企业信息：</label>
                                    <div class="col-sm-8">否</div>
                                </div>
                                <div class="col-sm-4">
                                    
                                </div>
                                <div class="col-sm-4">
                                    
                                </div>
                            </div>
                            
                        </div>
                    </div>
                </div>
                
                <div class="row">
                    <div class="col-sm-12 ">
                        <div class="ibox float-e-margins">
                            <h2>抽检样品信息</h2>
                            <div class="hr-line-dashed"></div>
                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品条码：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品商标：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品类型：</label>
                                    <div class="col-sm-8">食用农产品</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品来源：</label>
                                    <div class="col-sm-8">外购</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品属性：</label>
                                    <div class="col-sm-8">普通食品</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">包装规格：</label>
                                    <div class="col-sm-8">散装</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品名称：</label>
                                    <div class="col-sm-8">油麦菜</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">购进日期:</label>
                                    <div class="col-sm-8">2019-08-12</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">保质期：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品批号：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品规格：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">质量等级：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单价：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">是否进口：</label>
                                    <div class="col-sm-8">否</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">原产地：</label>
                                    <div class="col-sm-8">中国</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样日期：</label>
                                    <div class="col-sm-8">2019-08-12</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样方式：</label>
                                    <div class="col-sm-8">非无菌采样</div>
                                </div>
                                
                                    
                                    
                                
                            </div>

                            <div class="row form-group">
                                
                                    
                                    
                                
                                
                                    
                                    
                                
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样时样品储存条件：</label>
                                    <div class="col-sm-8">常温</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样基数：</label>
                                    <div class="col-sm-8">5kg</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样数量：</label>
                                    <div class="col-sm-8">1.5</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">备样数量：</label>
                                    <div class="col-sm-8">0.5kg</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样数量单位：</label>
                                    <div class="col-sm-8">kg</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">执行标准/技术文件：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">备注：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>


                
                <div class="row">
                    <div class="col-sm-12">
                        <div class="ibox float-e-margins">
                            <h2>照片信息</h2>
                            <div class="row form-group">
                                <div class="col-sm-12">
                                    <ul id="dowebok">
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload1.nifdc.org.cn/image/2019/10/23/157179720856246718.jpg"    data-original="http://upload1.nifdc.org.cn/image/2019/10/23/157179720856246718.jpg"
                                                     alt="现场抽样图片">
                                                <p>图片1</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload1.nifdc.org.cn/image/2019/10/23/157179707133890105.jpg"    data-original="http://upload1.nifdc.org.cn/image/2019/10/23/157179707133890105.jpg"
                                                     alt="现场抽样图片">
                                                <p>图片2</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload1.nifdc.org.cn/image/2019/10/23/157179707127242612.jpg"    data-original="http://upload1.nifdc.org.cn/image/2019/10/23/157179707127242612.jpg"
                                                     alt="现场抽样图片">
                                                <p>图片3</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload1.nifdc.org.cn/image/2019/10/23/15717970712929.323.jpg"    data-original="http://upload1.nifdc.org.cn/image/2019/10/23/15717970712929.323.jpg"
                                                     alt="现场抽样图片">
                                                <p>图片4</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload1.nifdc.org.cn/image/2019/10/23/157179707076136982.jpg"    data-original="http://upload1.nifdc.org.cn/image/2019/10/23/157179707076136982.jpg"
                                                     alt="现场抽样图片">
                                                <p>图片5</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload1.nifdc.org.cn/image/2019/10/23/157179707047563317.jpg"    data-original="http://upload1.nifdc.org.cn/image/2019/10/23/157179707047563317.jpg"
                                                     alt="现场抽样图片">
                                                <p>图片6</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload1.nifdc.org.cn/image/2019/10/23/157179707020553549.jpg"    data-original="http://upload1.nifdc.org.cn/image/2019/10/23/157179707020553549.jpg"
                                                     alt="现场抽样图片">
                                                <p>图片7</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload1.nifdc.org.cn/image/2019/10/23/157179706993669134.jpg"    data-original="http://upload1.nifdc.org.cn/image/2019/10/23/157179706993669134.jpg"
                                                     alt="现场抽样图片">
                                                <p>图片8</p>
                                            </div>
                                        
                                    </ul>
                                </div>
                            </div>
                            <div class="row form-group">
                                <div class="col-sm-12">
                                    
                                        
                                        
                                            <a href="http://upload1.nifdc.org.cn/image/2019/10/23/157179719824911257.png" target="_blank"
                                               class="btn btn-danger btn-xs"><i
                                                    class="fa fa-search"></i> 抽样单电子版</a>
                                        
                                    
                                    
                                        
                                        
                                            <a href="http://upload1.nifdc.org.cn/image/2019/10/23/157179707914408531.png" target="_blank"
                                               class="btn btn-danger btn-xs"><i
                                                    class="fa fa-search"></i> 抽样检验告知书电子版</a>
                                        
                                    

                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        
        <div class="row">

            <div class="col-sm-12 form-horizontal">
                <div class="ibox float-e-margins">
                    <h2>检验信息</h2>
                    <div class="hr-line-dashed"></div>
                    <form id="testform">
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检验机构名称：</label>
                                <div class="col-sm-7  text-left">
                                    <label class="control-label">河南恒晟检测技术有限公司</label>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">报告书编号：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">2019（食）A2562</label>

                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">样品到达日期：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">
                                        2019-08-12</label>

                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">联系人：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">王力</label>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">联系人电话：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">0371-55929768</label>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">联系人邮箱：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">hengshengjiance23@163.com</label>

                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人员：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">河南恒晟检测技术有限公司</label>
                                </div>
                            </div>

                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人电话：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">15738889730</label>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人邮箱：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">rdh123@163.com</label>

                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">结论：</label>
                                <div class="col-sm-7">
                                    <label id="hid_conclusion" class="control-label">纯抽检合格样品</label>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">监督抽检报告备注：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">
                                        
                                            
                                                /
                                            
                                            
                                        
                                    </label>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">风险监测报告备注：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">
                                        
                                            
                                                /
                                            
                                            
                                        
                                    </label>
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group hidden">
                                <label class="control-label col-sm-5">接样日期：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">
                                        
                                            2019-10-23
                                    </label>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">报告类别：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">
                                        
                                            
                                            
                                                合格报告
                                            
                                        
                                    </label>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检验目的/任务类别：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">监督抽检</label>
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5" disabled>历史退回原因：</label>
                                <div class="col-sm-7">
                                    <a id="showBackReason" class="btn btn-success btn-xs">点击查看</a>
                                </div>
                            </div>
                            
                                
                                
                                
                            
                            
                                

                                
                                
                            
                        </div>

                        

                            <div class="row">
                                <div class="col-sm-4 form-group">
                                    <label class="control-label col-sm-5">复检状态：</label>
                                    <div class="col-sm-7">
                                        
                                            

                                            
                                            
                                            
                                            
                                            
                                        
                                    </div>
                                </div>
                            </div>
                        

                        <div class="row hidden" id="unqualified-upload1">
                            <div class="col-sm-6">
                                <div class="alert alert-danger">
                                    <div id="uploader2" class="wu-example">
                                        <!--用来存放文件信息-->
                                        <div id="thelist2" class="uploader-list"></div>
                                        <span id="picker2">选择文件</span>
                                        <p class="help-block">只能上传 word 和 pdf 格式的文件</p>
                                        <label class="help-block m-b-none text-danger">上传不合格结果通知书</label>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-12">
                                <div class="alert alert-info alert-dismissable">
                                    <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×
                                    </button>
                                    <p>（1）标准最小允许限、标准最大允许限是根据判定依据确定或根据常见食品品种核定，承检机构应根据检验的具体产品确定最小允许限、最大允许限。</p>
                                    <p>（2）标准方法检出限是根据检验方法标准确定或核定（标准中未规定检出限或定量限时），承检机构应根据具体情况确定。</p>
                                </div>
                            </div>
                        </div>

                        <div class="row">
                            <div class="col-sm-12">
                                
                                <table id="dg" title="检验结果" class="easyui-datagrid" style="width:auto;height: 460px;"
                                       rownumbers="true" fitColumns="true" singleSelect="true">
                                </table>
                            </div>
                        </div>
                        <div class="hr-line-dashed"></div>

                        <div class="row">
                            <div class="col-sm-12">
                                <a id="showFujianInfo" class="btn btn-success btn-xs">复检信息</a>
                                <div class="hr-line-dashed"></div>
                            </div>
                        </div>


                        <div class="row">
                            <div class="col-sm-12">
                                <h3>检验结论</h3>
                                <div class="hr-line-dashed"></div>
                                <p>
                                    经抽样检验，所检项目符合 GB 2763-2016《食品安全国家标准 食品中农药最大残留限量》 要求。
                                </p>
                            </div>
                        </div>
                        <input type="hidden" id="sd" name="sd" value="69bg7TjNNdEeByxhjeremyauXgOKgmxMbkUD67zXj0wm2lPtmCc=">
                        <input type="hidden" id="st" name="st" value="4">
                        <input type="hidden" id="sis" name="sis" value="9">
                        <input type="hidden" id="rd" name="rd" value="1">
                        <input type="hidden" id="unqualifed" value="">
                    </form>
                    <div class="hr-line-dashed"></div>
                    <div class="row hidden" id="unqualified-upload">
                        <div class="col-sm-6">
                            <div class="alert alert-danger">
                                <div id="uploader1" class="wu-example">
                                    <!--用来存放文件信息-->
                                    <div id="thelist1" class="uploader-list"></div>
                                    <span id="picker1">选择文件</span>
                                    <p class="help-block">只能上传 word 和 pdf 格式的文件</p>
                                    <label class="help-block m-b-none text-danger">上传不合格结果通知书</label>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-sm-12">
                            
                                
                                
                                
                                
                                
                                
                                    <div class="row">
                                        <div class="col-sm-12">
                                            
                                                
                                                    <button data-toggle="tooltip" data-placement="top" title="查看"
                                                            onclick="detailManager.viewReport('http://111.204.20.233:39054/report/jd_report/4907/20191024/NCP19131128463231006_JDCJ_1fca6db8-a6d6-4032-aff5-112d73ecbccf.pdf','查看监督抽检报告')"
                                                            class="btn btn-primary btn-xs btn-outline"
                                                            id="btn-previewJd" type="button">
                                                        <i class="fa fa-eye"></i> 监督抽检报告预览
                                                    </button>
                                                
                                                
                                                
                                                
                                            
                                        </div>
                                    </div>
                                    <h2 class="text-navy">状态: 已完全提交</h2>
                                
                                
                                
                                
                                
                                
                                

                                
                                
                                
                                
                                
                                

                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                                
                            
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="row hidden">
                        <div class="col-sm-12">
                            <h3>历史报告</h3>
                            <button class="btn btn-default" type="button">无</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="hidden" id="back-div">
    <textarea class="form-control " style="width: 99%; margin: 5px auto;overflow-x: hidden" rows="5"
              placeholder="请输入退回原因..." id="backReason"></textarea>
    <div class="text-right">
        <button id="btn-back" type="button" style="margin: 5px 5px" class="btn btn-danger">退回</button>
        <button id="btn-cancel" type="button" style="margin: 5px 5px" class="btn btn-default">取消</button>
    </div>
</div>

<input type="hidden" id="test_aims" value="监督抽检"/>


<script type="text/javascript" src="/test_platform/js/jquery.min.js"></script>
<script type="text/javascript" src="/test_platform/js/bootstrap.min.js"></script>
<script type="text/javascript" src="/test_platform/js/plugins/toastr/toastr.min.js"></script>
<script type="text/javascript" src="/test_platform/js/plugins/datapicker/bootstrap-datepicker.js"></script>
<script type="text/javascript" src="/test_platform/js/plugins/chosen/chosen.jquery.js"></script>
<script type="text/javascript" src="/test_platform/js/plugins/layer/layer.min.js"></script>
<script type="text/javascript" src="/test_platform/js/jquery.cookie.js"></script>
<script type="text/javascript" src="/test_platform/js/common/globaltools.js?v=1.2.0.43"></script>


<script>
    $(function () {
        $("[data-toggle='tooltip']").tooltip();
        var viewer = new Viewer(document.getElementById('dowebok'), {
            url: 'data-original'
        });
    });
</script>
<script src="/test_platform/js/webuploader.js"></script>
<script src="/test_platform/js/plugins/fancybox/jquery.fancybox.js?v=1.2.0.43"></script>
<script src="/test_platform/easyui/jquery.easyui.min.js?v=1.2.0.43"></script>
<script src="/test_platform/js/plugins/fancybox/viewer.min.js"></script>
<script src="/test_platform/easyui/locale/easyui-lang-zh_CN.js?v=1.2.0.43"></script>
<script src="/test_platform/js/uploadTool.js"></script>
<script src="/test_platform/js/food/foodDetailReadOnly.js?v=1.2.0.43"></script>
</body>
</html>

`

	mkr := StoMap_foodDetail(olds)
	for k, v := range mkr {
		fmt.Printf("%s:%s\n", k, v)
	}
	tmj := template.New("tmj")
	tmj.Funcs(map[string]interface{}{
		"replace": strings.ReplaceAll,
	})
	tmj.Parse("aaa{{ .抽样单位信息_传真 }} {{ .抽样基础信息_食品次亚类 }}")
	fmt.Println(tmj.Execute(os.Stdout, mkr))
}
func TestTest_platform_api_food_getTestItems(t *testing.T) {
	a, b, c, err := InitLoginck(nil)
	if err != nil {
		t.Fatal(err)
	}
	ck, err := Login("17761660651", "12345678", a, b, c, nil)
	if err != nil {
		t.Fatal(err)
	}
	test_platform_ck, err := Test_platform_login(ck, nil)
	if err != nil {
		t.Fatal(err)
	}
	fddetail, err := Test_platform_foodTest_foodDetail(12156977, test_platform_ck, nil)
	if err != nil {
		t.Fatal(err)
	}
	sd:=fddetail["sd"]
	itemsr,err:=Test_platform_api_food_getTestItems(sd,ck,nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(itemsr)
}