package nifdc

import (
	"fmt"
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
            <small class="text-danger">&nbsp;&nbsp;&nbsp;&nbsp; 抽样编号:DC19130400463233437</small>
        </h1>
        <div class="panel panel-success">
            <div class="panel-heading">
                <h5 class="panel-title" style="color: #ffffff;">
                    <a data-toggle="collapse" title="点击展开或收起" data-parent="#accordion" href="#collapseOne"
                       aria-expanded="false"
                       class="collapsed"><i class="fa fa-info-circle"></i> 抽样信息内容
                        <small style="color: #ffffff;">(点击展开或收起)</small>
                    </a>
                </h5>
            </div>
            <div id="collapseOne" class="panel-collapse collapse" aria-expanded="false" style="height: 0px;">
                <div class="panel-body" style="padding: 15px;">
                    
                    <div class="row">
                        <div class="col-sm-12">
                            <div class="ibox float-e-margins">
                                <h2>抽样基础信息</h2>
                                <div class="hr-line-dashed"></div>
                                <div class="row form-group">
                                    <div class="col-sm-4 ">
                                        <label class="control-label col-sm-4">任务来源：</label>
                                        <div class="col-sm-8">邯郸市市场监督管理局</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">报送分类A：</label>
                                        <div class="col-sm-8" id="bsfla">抽检监测（市级本级）
                                            <input type="hidden" id="hid_bsfla" value="抽检监测（市级本级）"/>
                                        </div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">报送分类B：</label>
                                        <div class="col-sm-8" id="bsflb">2019年河北邯郸抽检计划E包
                                            <input type="hidden" id="hid_bsflb" value="2019年河北邯郸抽检计划E包"/>
                                        </div>
                                    </div>
                                </div>
                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">食品大类：</label>
                                        <div class="col-sm-8" id="type1">调味品
                                            <input type="hidden" id="hid_type1" value="调味品"/>
                                        </div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">食品亚类：</label>
                                        <div class="col-sm-8" id="type2">调味料酒
                                            <input type="hidden" id="hid_type2" value="调味料酒"/>
                                        </div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">食品次亚类：</label>
                                        <div class="col-sm-8" id="type3">调味料酒
                                            <input type="hidden" id="hid_type3" value="调味料酒"/>
                                        </div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">食品细类：</label>
                                        <div class="col-sm-8" id="type4">料酒
                                            <input type="hidden" id="hid_type4" value="料酒"/>
                                        </div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样单编号：</label>
                                        <div class="col-sm-8">DC19130400463233437</div>
                                        <input type="hidden" id="hid_sp_s_16" value="DC19130400463233437"/>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样类型：</label>
                                        <div class="col-sm-8">常规抽样</div>
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
                                        <div class="col-sm-8">周春伟、徐明欣</div>
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
                                        <div class="col-sm-8">冀南新区郝鲜生餐饮店</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">单位地址：</label>
                                        <div class="col-sm-8">河北省邯郸市冀南新区城南办事处辛庄营村</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">所在地：</label>
                                        <div class="col-sm-8">河北/邯郸/冀南新区</div>
                                    </div>
                                </div>


                                
                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">区域类型：</label>
                                        <div class="col-sm-8">乡村</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样环节：</label>
                                        <div class="col-sm-8">餐饮</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样地点：</label>
                                        <div class="col-sm-8">小型餐馆</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">营业执照/社会信用代码：</label>
                                        <div class="col-sm-8">92130492MA0CTHKF28</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">许可证类型：</label>
                                        <div class="col-sm-8">经营许可证</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">许可证号：</label>
                                        <div class="col-sm-8">JY21304010003620</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">单位法人：</label>
                                        <div class="col-sm-8">王亚杰</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">联系人：</label>
                                        <div class="col-sm-8">王亚杰</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">联系人电话：</label>
                                        <div class="col-sm-8">15100097097</div>
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
                                        <div class="col-sm-8">山西紫林醋业股份有限公司</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">所在地：</label>
                                        <div class="col-sm-8">山西/太原/清徐</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">企业地址：</label>
                                        <div class="col-sm-8">山西省清徐县太茅路高花段550号</div>
                                    </div>
                                </div>
                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">生产许可证编号：</label>
                                        <div class="col-sm-8">SC10314012101087</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">联系人：</label>
                                        <div class="col-sm-8">/</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">电话：</label>
                                        <div class="col-sm-8">0351-5737056</div>
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
                                        <div class="col-sm-8">6922165902996</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">样品商标：</label>
                                        <div class="col-sm-8">紫林+图形</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">样品类型：</label>
                                        <div class="col-sm-8">工业加工食品</div>
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
                                        <label class="control-label col-sm-4">包装分类：</label>
                                        <div class="col-sm-8">预包装</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">样品名称：</label>
                                        <div class="col-sm-8">料酒  调味料酒</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">生产日期:</label>
                                        <div class="col-sm-8">2019-02-25</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">保质期：</label>
                                        <div class="col-sm-8">二年</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">样品批号：</label>
                                        <div class="col-sm-8">/</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">规格型号：</label>
                                        <div class="col-sm-8">500mL/瓶</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">质量等级：</label>
                                        <div class="col-sm-8">/</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">单价：</label>
                                        <div class="col-sm-8">3元/瓶</div>
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
                                        <div class="col-sm-8">2019-11-29</div>
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
                                        <div class="col-sm-8">4瓶</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样数量：</label>
                                        <div class="col-sm-8">1.0</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">备样数量：</label>
                                        <div class="col-sm-8">0.5瓶</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样数量单位：</label>
                                        <div class="col-sm-8">瓶</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">执行标准/技术文件：</label>
                                        <div class="col-sm-8">SB/T10416</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">备注：</label>
                                        <div class="col-sm-8">抽样信息由被抽样单位提供   样品现场分装</div>
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
                                                         src="http://upload1.nifdc.org.cn/image/2019/11/29/157500921826794911.png"    data-original="http://upload1.nifdc.org.cn/image/2019/11/29/157500921826794911.png"
                                                         alt="现场抽样图片">
                                                    <p>图片1</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://upload1.nifdc.org.cn/image/2019/11/29/157500921882637848.png"    data-original="http://upload1.nifdc.org.cn/image/2019/11/29/157500921882637848.png"
                                                         alt="现场抽样图片">
                                                    <p>图片2</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://upload1.nifdc.org.cn/image/2019/11/29/157500921871015135.png"    data-original="http://upload1.nifdc.org.cn/image/2019/11/29/157500921871015135.png"
                                                         alt="现场抽样图片">
                                                    <p>图片3</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://upload1.nifdc.org.cn/image/2019/11/29/157500921859756265.png"    data-original="http://upload1.nifdc.org.cn/image/2019/11/29/157500921859756265.png"
                                                         alt="现场抽样图片">
                                                    <p>图片4</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://upload1.nifdc.org.cn/image/2019/11/29/157500921860707360.png"    data-original="http://upload1.nifdc.org.cn/image/2019/11/29/157500921860707360.png"
                                                         alt="现场抽样图片">
                                                    <p>图片5</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://upload1.nifdc.org.cn/image/2019/11/29/157500921935917444.png"    data-original="http://upload1.nifdc.org.cn/image/2019/11/29/157500921935917444.png"
                                                         alt="现场抽样图片">
                                                    <p>图片6</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://upload1.nifdc.org.cn/image/2019/11/29/157500921982044914.png"    data-original="http://upload1.nifdc.org.cn/image/2019/11/29/157500921982044914.png"
                                                         alt="现场抽样图片">
                                                    <p>图片7</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://upload1.nifdc.org.cn/image/2019/11/29/157500921990963163.png"    data-original="http://upload1.nifdc.org.cn/image/2019/11/29/157500921990963163.png"
                                                         alt="现场抽样图片">
                                                    <p>图片8</p>
                                                </div>
                                            
                                        </ul>
                                    </div>

                                </div>
                                <div class="row form-group">
                                    <div class="col-sm-12">
                                        
                                            
                                            
                                                <a href="http://upload1.nifdc.org.cn/image/2019/11/29/157500916111694385.png" target="_blank"
                                                   class="btn btn-danger btn-xs"><i
                                                        class="fa fa-search"></i> 抽样单电子版</a>
                                            
                                        
                                        
                                            
                                            
                                                <a href="http://upload1.nifdc.org.cn/image/2019/11/29/157500916656939604.png" target="_blank"
                                                   class="btn btn-danger btn-xs"><i
                                                        class="fa fa-search"></i> 抽样检验告知书电子版</a>
                                            
                                        
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        
        <form>
            <div class="row">
                <div class="col-sm-12 form-horizontal">
                    <div class="ibox float-e-margins">
                        <h2>检验信息</h2>
                        <div class="hr-line-dashed"></div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检验机构名称：</label>
                                <div class="col-sm-7">
                                    <input type="text" disabled="disabled" value="河南恒晟检测技术有限公司"
                                           class="form-control">
                                    <input type="hidden" id="test_unit" name="test_unit" value="河南恒晟检测技术有限公司"/>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">报告书编号：</label>
                                <div class="col-sm-7">
                                    <input type="text" id="report_no" name="report_no" value=""
                                           class="form-control"
                                           title="请输入报告书编号" required>

                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">样品到达日期：</label>
                                <div class="col-sm-7">
                                    <div class="input-group date">
                                        <span class="input-group-addon"><i class="fa fa-calendar"></i></span>
                                        <input type="text" disabled id="test_date" name="test_date" class="form-control"
                                               value="2019-11-29"
                                               placeholder="请选择日期"
                                               data-provide="datepicker">
                                    </div>

                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">联系人：</label>
                                <div class="col-sm-7">
                                    <input type="text" disabled id="contact" name="contact"
                                           value="王力"
                                           class="form-control">
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">联系人电话：</label>
                                <div class="col-sm-7">
                                    <input type="text" disabled id="contact_tel"  name="contact_tel"
                                           value="0371-55929768"
                                           class="form-control"
                                           title="请输入有效的座机号或手机号">
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">联系人邮箱：</label>
                                <div class="col-sm-7">
                                    <input type="text" disabled title="请输入有效的邮箱"
                                            id="contact_email"
                                           name="contact_email" value="hengshengjiance23@163.com"
                                           class="form-control">
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人员：</label>
                                <div class="col-sm-7">
                                    <input type="text" id="fy_person"   name="fy_person" value="河南恒晟检测技术有限公司"
                                           class="form-control" required>
                                </div>
                            </div>

                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人电话：</label>
                                <div class="col-sm-7">
                                    <input type="text"  title="请输入有效的座机号或手机号"
                                           required id="fy_tel" name="fy_tel" value="15738889730"
                                           class="form-control">
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人邮箱：</label>
                                <div class="col-sm-7">
                                    <input type="text"  title="请输入有效的邮箱"
                                           required
                                           id="fy_email" name="fy_email" value="rdh123@163.com"
                                           class="form-control">
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">结论：</label>
                                <div class="col-sm-7">
                                    <input type="text" disabled id="conclusion" value=""
                                           class="form-control">
                                    <input type="hidden" id="hid_conclusion" name="conclusion"
                                           value=""/>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">报告类别：</label>
                                <div class="col-sm-7">
                                    <select class="form-control" id="sp_bsb_bgfl" name="report_type">
                                        
                                            <option 
                                                    value="合格报告">合格报告</option>
                                        
                                            <option 
                                                    value="一般不合格报告">一般不合格报告</option>
                                        
                                            <option 
                                                    value="一般问题报告">一般问题报告</option>
                                        
                                            <option 
                                                    value="一般不合格（问题）报告">一般不合格（问题）报告</option>
                                        
                                            <option 
                                                    value="24小时限时报告">24小时限时报告</option>
                                        
                                    </select>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检验目的/任务类别：</label>
                                <div class="col-sm-7">
                                    <select class="form-control" disabled id="test_aims" name="test_aims">
                                        
                                            <option  selected="selected"
                                                    value="监督抽检">监督抽检</option>
                                        
                                            <option 
                                                    value="抽检监测">抽检监测</option>
                                        
                                            <option 
                                                    value="风险监测">风险监测</option>
                                        
                                            <option 
                                                    value="评价性抽检">评价性抽检</option>
                                        
                                    </select>
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group hidden">
                                <label class="control-label col-sm-5">接样日期：</label>
                                <div class="col-sm-7">
                                    <div class="input-group date">
                                        <span class="input-group-addon"><i class="fa fa-calendar"></i></span>
                                        <input type="text" disabled="disabled" id="tb_date" name="tb_date"
                                               class="form-control"
                                               value="2019-12-03"
                                               placeholder="请选择日期"
                                               data-provide="datepicker" id="startDate">
                                    </div>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">监督抽检报告备注：</label>
                                <div class="col-sm-7">
                                    <textarea id="jd_bz" rows="4" name="jd_bz"
                                              class="form-control"></textarea>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">风险监测报告备注：</label>
                                <div class="col-sm-7">
                                    <textarea id="fx_bz" rows="4" name="fx_bz"
                                              class="form-control"></textarea>
                                </div>
                            </div>
                        </div>
                        <div class="col-sm-4 form-group hidden" id="signdate">
                            <label class="control-label col-sm-5">选择批准日期：</label>
                            <div class="col-sm-7">
                                <select class="form-control"  id="sign_date" name="sign_date">
                                    <option value="">请选择</option>
                                    <option  value="2019-09-20">2019-09-20</option>
                                    <option  value="2019-09-19">2019-09-19</option>
                                </select>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">历史退回原因：</label>
                                <div class="col-sm-7">
                                    <a id="showBackReason" class="btn btn-success btn-xs">点击查看</a>
                                </div>
                            </div>
                            
                                
                                
                                
                            
                            
                                

                                
                                
                            
                        </div>
                    </div>
                    <div class="row hidden" id="div-upload1">
                        <div class="col-sm-6">
                            <div class="alert alert-danger">
                                <div id="uploader1" class="wu-example">
                                    <!--用来存放文件信息-->
                                    <div id="thelist1" class="uploader-list"></div>
                                    <span id="picker1">选择文件</span>
                                    <label class="help-block m-b-none text-danger">上传24小时限时报告附件</label>
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
                        <div class="col-sm-6">
                            <div class="alert alert-danger">
                                <label>上下标提示: 1.输入上标：Ctrl + Alt + (数字或+-号) ， 2.输入下标：Alt + 数字(数字或+-号)。</label>
                            </div>
                        </div>
                        <div class="col-sm-6"></div>
                    </div>

                    <div class="row">
                        <div class="col-sm-4">
                            <input type="text" id="search" class="form-control" placeholder="输入项目名称进行快速定位"/>
                            <label class="help-block m-b-none text-danger">提示：在文本框中输入检验项目名称的关键字可以快速定位检验项目</label>
                        </div>
                        <div class="col-sm-4">
                        </div>
                        <div class="col-sm-4 text-right form-inline">
                            <input class="form-control" placeholder="请输入抽样单号" id="txt_sampleNO" autocomplete="on">
                            <button type="button" id="load_his" class="btn btn-primary" style="margin-bottom: 0;">
                                加载历史数据
                            </button>
                            <label class="help-block m-b-none text-danger">提示: 自动加载相同食品类别和报送分类的最后一次提交的检验数据</label>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-sm-12">
                            
                            <table id="dg" title="检验结果" class="easyui-datagrid" style="width:auto;height: 380px;"
                                   rownumbers="true" fitColumns="true" singleSelect="true">
                            </table>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-sm-12">
                            <p>
                            <h3>检验结论</h3></p>
                            <div class="hr-line-dashed"></div>
                            <textarea id="test_conclusion" name="test_conclusion"
                                      style="width: 100%;font-size: 16px;"
                                      rows="5"></textarea>
                        </div>
                    </div>
                    <input type="hidden" id="sd" name="sd" value="feZNnZmOeM4PkyckmjUJBk7NHIsXBZkMm6+ADHwKsYA=">
                    <input type="hidden" id="st" name="sd" value="-1">
                    <input type="hidden" id="userId" value="V5yxhjeremyLS3+3OeuhDX3skYbsrrXcp3Wlh8oA">
                    <input type="hidden" id="xs" value="">
                    <input type="hidden" id="unqualifed" value="">

                    <div class="row hidden" id="div-upload" style="margin-top: 10px">
                        <div class="col-sm-6">
                            <div class="alert alert-danger">
                                <div id="uploader" class="wu-example">
                                    <!--用来存放文件信息-->
                                    <div id="thelist" class="uploader-list"></div>
                                    <span id="picker">选择文件</span>
                                    <label class="help-block m-b-none text-danger">上传24小时限时报告附件</label>
                                </div>
                            </div>
                        </div>
                    </div>

                    
                    
                    

                    <div class="row" style="margin-top: 10px">
                        <div class="col-sm-12">
                            
                                
                                    <button class="btn btn-success" id="save" type="button">临时保存</button>
                                    <button class="btn btn-primary" id="submit" type="submit">检测数据提交入库</button>
                                    <button class="btn btn-danger" id="back" type="button">退修</button>
                                    <button class="btn btn-info" id="item-reset" type="button">检验项目重置</button>
                                    <span class="help-block m-b-none text-danger">提示:点击检验项目重置按钮可以重新加载最新的基础表</span>


                                
                                
                            
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
    </form>
</div>
</div>
<div class="hidden" id="back-div">

       
    <textarea class="form-control" style="width: 99%; margin: 5px auto;overflow-x: hidden" rows="5"
              placeholder="请输入退回原因..." id="backReason"></textarea>
    <div class="text-right">
        <button id="btn-back" type="button" style="margin: 5px 5px" class="btn btn-danger">退回</button>
        <button id="btn-cancel" type="button" style="margin: 5px 5px" class="btn btn-default">取消</button>
    </div>
</div>
<div class="hidden" id="save-div">
    <div class="input-group">
        <input type="text" class="form-control" id="verifyCode"name="verifyCode" placeholder="请输入验证码">
    </div>
    <span  style="text-align: center;">
        <img src="/test_platform/verifyCode" id="vcImg" onclick="myRefersh(this)">
    </span>

    <div class="text-right">
        <button id="btn-save" type="button" style="margin: 5px 5px" class="btn btn-danger">保存</button>
        <button id="btn-cance2" type="button" style="margin: 5px 5px" class="btn btn-default">取消</button>
    </div>
</div>



<script type="text/javascript" src="/test_platform/js/jquery.min.js"></script>
<script type="text/javascript" src="/test_platform/js/bootstrap.min.js"></script>
<script type="text/javascript" src="/test_platform/js/plugins/toastr/toastr.min.js"></script>
<script type="text/javascript" src="/test_platform/js/plugins/datapicker/bootstrap-datepicker.js"></script>
<script type="text/javascript" src="/test_platform/js/plugins/chosen/chosen.jquery.js"></script>
<script type="text/javascript" src="/test_platform/js/plugins/layer/layer.min.js"></script>
<script type="text/javascript" src="/test_platform/js/jquery.cookie.js"></script>
<script type="text/javascript" src="/test_platform/js/common/globaltools.js?v=1.2.0.42"></script>


<script>
    $(function () {
        $("[data-toggle='tooltip']").tooltip();

        var viewer = new Viewer(document.getElementById('dowebok'), {
            url: 'data-original'
        });
    });
</script>
<script src="/test_platform/js/webuploader.js"></script>
<script src="/test_platform/js/plugins/layer/layer.min.js"></script>
<script src="/test_platform/js/plugins/fancybox/jquery.fancybox.js"></script>
<script src="/test_platform/js/plugins/fancybox/viewer.min.js"></script>
<script src="/test_platform/easyui/jquery.easyui.min.js"></script>
<script src="/test_platform/easyui/locale/easyui-lang-zh_CN.js"></script>
<script src="/test_platform/js/uploadTool.js"></script>
<script src="/test_platform/js/food/foodDetail.js?v=1.2.0.42"></script>

</body>
</html>

`

	mkr := StoMap_foodDetail(olds)
	for k, v := range mkr {
		fmt.Printf("%s:%s\n", k, v)
	}
}
