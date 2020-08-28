package nifdc

import (
	"encoding/gob"
	"fmt"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {

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
            <small class="text-danger">&nbsp;&nbsp;&nbsp;&nbsp; 抽样编号:DC20131100120631183</small>
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
                                        <div class="col-sm-8">衡水市市场监督管理局</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">报送分类A：</label>
                                        <div class="col-sm-8" id="bsfla">抽检监测（市级本级）
                                            <input type="hidden" id="hid_bsfla" value="抽检监测（市级本级）"/>
                                        </div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">报送分类B：</label>
                                        <div class="col-sm-8" id="bsflb">2020年河北衡水常规抽检监测计划（流通环节）
                                            <input type="hidden" id="hid_bsflb" value="2020年河北衡水常规抽检监测计划（流通环节）"/>
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
                                        <div class="col-sm-8" id="type2">生干坚果与籽类食品
                                            <input type="hidden" id="hid_type2" value="生干坚果与籽类食品"/>
                                        </div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">食品次亚类：</label>
                                        <div class="col-sm-8" id="type3">生干坚果与籽类食品
                                            <input type="hidden" id="hid_type3" value="生干坚果与籽类食品"/>
                                        </div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">食品细类：</label>
                                        <div class="col-sm-8" id="type4">生干坚果
                                            <input type="hidden" id="hid_type4" value="生干坚果"/>
                                        </div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样单编号：</label>
                                        <div class="col-sm-8">DC20131100120631183</div>
                                        <input type="hidden" id="hid_sp_s_16" value="DC20131100120631183"/>
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
                                        <div class="col-sm-8">衡水市食品药品检验检测中心</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">单位地址：</label>
                                        <div class="col-sm-8">衡水市永兴西路2488号</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">所在省份：</label>
                                        <div class="col-sm-8">河北</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样人员：</label>
                                        <div class="col-sm-8">雷倩、蔡勇</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样人员电话：</label>
                                        <div class="col-sm-8">03182361893</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">单位联系人：</label>
                                        <div class="col-sm-8">李春花</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    
                                    
                                    
                                    
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">传真：</label>
                                        <div class="col-sm-8">03182813793</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">邮编：</label>
                                        <div class="col-sm-8">053000</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">联系人电话：</label>
                                        <div class="col-sm-8">03182361893</div>
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
                                        <div class="col-sm-8">衡水悦华超市有限公司</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">单位地址：</label>
                                        <div class="col-sm-8">衡水市高新区人民路北侧、顺兴街西侧尚品林溪综合楼负一层</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">所在地：</label>
                                        <div class="col-sm-8">河北/衡水/工业新区</div>
                                    </div>
                                </div>


                                
                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">区域类型：</label>
                                        <div class="col-sm-8">城市</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样环节：</label>
                                        <div class="col-sm-8">流通</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样地点：</label>
                                        <div class="col-sm-8">超市</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">营业执照/社会信用代码：</label>
                                        <div class="col-sm-8">91131101MA0CYJGE1F</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">许可证类型：</label>
                                        <div class="col-sm-8">经营许可证</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">许可证号：</label>
                                        <div class="col-sm-8">JY11311000002991</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">单位法人：</label>
                                        <div class="col-sm-8">王晓东</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">联系人：</label>
                                        <div class="col-sm-8">武文华</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">联系人电话：</label>
                                        <div class="col-sm-8">15633543111</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    
                                        
                                        
                                    
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
                                        <div class="col-sm-8">河北/衡水/工业新区</div>
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
                                        <div class="col-sm-8">无包装</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">样品名称：</label>
                                        <div class="col-sm-8">核桃</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">购进日期:</label>
                                        <div class="col-sm-8">2020-08-03</div>
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
                                        <div class="col-sm-8">计量称重</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">质量等级：</label>
                                        <div class="col-sm-8">/</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">单价：</label>
                                        <div class="col-sm-8">25.8元/kg</div>
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
                                        <div class="col-sm-8">2020-08-03</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样方式：</label>
                                        <div class="col-sm-8">非无菌采样</div>
                                    </div>
                                    
                                        
                                        
                                    
                                </div>

                                <div class="row form-group">
                                    
                                        
                                        
                                    
                                    
                                        
                                        
                                    
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">储存条件：</label>
                                        <div class="col-sm-8">常温</div>
                                    </div>
                                </div>

                                <div class="row form-group">
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样基数：</label>
                                        <div class="col-sm-8">6kg</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">抽样数量：</label>
                                        <div class="col-sm-8">3.236</div>
                                    </div>
                                    <div class="col-sm-4">
                                        <label class="control-label col-sm-4">备样数量：</label>
                                        <div class="col-sm-8">1kg</div>
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
                                                         src="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644279570657974.jpeg"    data-original="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644279570657974.jpeg"
                                                         alt="现场抽样图片">
                                                    <p>图片1</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280879772737.jpeg"    data-original="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280879772737.jpeg"
                                                         alt="现场抽样图片">
                                                    <p>图片2</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280878596638.jpeg"    data-original="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280878596638.jpeg"
                                                         alt="现场抽样图片">
                                                    <p>图片3</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280884836142.jpeg"    data-original="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280884836142.jpeg"
                                                         alt="现场抽样图片">
                                                    <p>图片4</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280886186155.jpeg"    data-original="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280886186155.jpeg"
                                                         alt="现场抽样图片">
                                                    <p>图片5</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280878899410.jpeg"    data-original="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644280878899410.jpeg"
                                                         alt="现场抽样图片">
                                                    <p>图片6</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/15964512646101.525.jpeg"    data-original="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/15964512646101.525.jpeg"
                                                         alt="现场抽样图片">
                                                    <p>图片7</p>
                                                </div>
                                            
                                                <div style="display: inline-block;  overflow: hidden">
                                                    <img style="width:150px;height:150px;margin-right: 15px;"
                                                         src="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159645126427121176.jpeg"    data-original="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159645126427121176.jpeg"
                                                         alt="现场抽样图片">
                                                    <p>图片8</p>
                                                </div>
                                            
                                        </ul>
                                    </div>

                                </div>
                                <div class="row form-group">
                                    <div class="col-sm-12">
                                        
                                            
                                            
                                                <a href="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159645127886180204.png" target="_blank"
                                                   class="btn btn-danger btn-xs"><i
                                                        class="fa fa-search"></i> 抽样单电子版</a>
                                            
                                        
                                        
                                            
                                            
                                                <a href="http://spcjupload2.gsxt.gov.cn/image/2020/08/03/159644281237030686.png" target="_blank"
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
                                    <input type="text" disabled="disabled" value="衡水市食品药品检验检测中心"
                                           class="form-control">
                                    <input type="hidden" id="test_unit" name="test_unit" value="衡水市食品药品检验检测中心"/>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">报告书编号：</label>
                                <div class="col-sm-7">
                                    <input type="text" id="report_no" name="report_no" value="HSSP202008007"
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
                                               value="2020-08-04"
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
                                           value="李春花"
                                           class="form-control">
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">联系人电话：</label>
                                <div class="col-sm-7">
                                    <input type="text" disabled id="contact_tel"  name="contact_tel"
                                           value="03182361893"
                                           class="form-control"
                                           title="请输入有效的座机号或手机号">
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">联系人邮箱：</label>
                                <div class="col-sm-7">
                                    <input type="text" disabled title="请输入有效的邮箱"
                                            id="contact_email"
                                           name="contact_email" value="hssspyp@sina.com"
                                           class="form-control">
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人员：</label>
                                <div class="col-sm-7">
                                    <input type="text" id="fy_person"   name="fy_person" value="李聪"
                                           class="form-control" required>
                                </div>
                            </div>

                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人电话：</label>
                                <div class="col-sm-7">
                                    <input type="text"  title="请输入有效的座机号或手机号"
                                           required id="fy_tel" name="fy_tel" value="15333388077"
                                           class="form-control">
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">检查封样人邮箱：</label>
                                <div class="col-sm-7">
                                    <input type="text"  title="请输入有效的邮箱"
                                           required
                                           id="fy_email" name="fy_email" value="/"
                                           class="form-control">
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">结论：</label>
                                <div class="col-sm-7">
                                    <input type="text" disabled id="conclusion" value="纯抽检合格样品"
                                           class="form-control">
                                    <input type="hidden" id="hid_conclusion" name="conclusion"
                                           value="纯抽检合格样品"/>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">报告类别：</label>
                                <div class="col-sm-7">
                                    <select class="form-control" id="sp_bsb_bgfl" name="report_type">
                                        
                                            <option  selected="selected"
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
                                               value="2020-08-04"
                                               placeholder="请选择日期"
                                               data-provide="datepicker" id="startDate">
                                    </div>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">监督抽检报告备注：</label>
                                <div class="col-sm-7">
                                    <textarea id="jd_bz" rows="4" name="jd_bz"
                                              class="form-control">123</textarea>
                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">风险监测报告备注：</label>
                                <div class="col-sm-7">
                                    <textarea id="fx_bz" rows="4" name="fx_bz"
                                              class="form-control">444</textarea>
                                </div>
                            </div>
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
                                      rows="5">经抽样检验，所检项目符合 GB 19300-2014《食品安全国家标准 坚果与籽类食品》，GB 2762-2017《食品安全国家标准 食品中污染物限量》 要求。</textarea>
                        </div>
                    </div>
                    <input type="hidden" id="sd" name="sd" value="ISwaBUeyxhjeremy4b9Bn7Y4w4CyxhjeremyioLh4LmShC4yxhjeremykVyxhjeremyq2hBlV68=">
                    <input type="hidden" id="st" name="sd" value="-1">
                    <input type="hidden" id="userId" value="R0rN23mmF5kjyxhjeremyfsZQiHBPrTlBBKtLup5">
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
<script type="text/javascript" src="/test_platform/js/common/globaltools.js?v=1.2.0.44"></script>


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
<script src="/test_platform/js/food/foodDetail.js?v=1.2.0.44"></script>

</body>
</html>


`

	mkr := StoMap_test_platform(olds)
	fmt.Println(mkr)
	//fmt.Println(mkr["抽检样品信息_生产日期"])
	//for k, v := range mkr {
	//	fmt.Printf("%s:%s\n", k, v)
	//}
	//tmj := template.New("tmj")
	//tmj.Funcs(map[string]interface{}{
	//	"replace": strings.ReplaceAll,
	//})
	//tmj.Parse("aaa{{ .抽样单位信息_传真 }} {{ .抽样基础信息_食品次亚类 }}")
	//fmt.Println(tmj.Execute(os.Stdout, mkr))
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
	fddetail, err := Test_platform_agricultureTest_agricultureDetail(12940542, true, test_platform_ck, nil)
	if err != nil {
		t.Fatal(err)
	}
	sd := fddetail.Get("sd").ToString()
	itemsr, err := Test_platform_api_agriculture_getTestItems(fddetail, test_platform_ck, nil)
	if err != nil {
		t.Fatal(err)
	}
	testinfor, err := Test_platform_api_agriculture_getTestInfo(sd, test_platform_ck, nil)
	if err != nil {
		t.Fatal(err)
	}

	mps := Build_agriculture_updata(itemsr.Rows, testinfor.Rows, []map[string]string{
		{
			"检验项目": "五氯酚酸钠(以五氯酚计)",
			"检验结果": "0.1",
			"结果判定": "合格项",
			"检验方法": "/",
			"判定依据": "/1",
			"说明":   "aaa",
		},
	})

	err = Test_platform_api_agriculture_save(fddetail, mps, test_platform_ck, nil)
	fmt.Println(err)
	//for _,it:=range mps{
	//	for k,v:=range it{
	//		fmt.Printf("%s:%s\n",k,v)
	//	}
	//}

}

func TestBuildbaogao(t *testing.T) {
	var updates []map[string]string
	redf, err := os.Open("./test_baogaodata")
	if err != nil {
		panic(err)
	}
	dc := gob.NewDecoder(redf)
	err = dc.Decode(&updates)
	if err != nil {
		panic(err)
	}
	fmt.Println(updates[0]["item"])
	updates[0]["sp_data_2"] = "不合格项"
	updates[1]["sp_data_2"] = "不合格项"
	updates[2]["sp_data_2"] = "不合格项"
	updates[3]["sp_data_2"] = "不合格项"
	//for _, item := range updates {
	//	fmt.Println(item)
	//}
	cvdata := Convbaotaodata(updates)
	fmt.Println(Buildbaogao(cvdata))
}

func TestMerge_subitem(t *testing.T) {
	updatas := make([]map[string]string, 0)
	subitem := make([]map[string]string, 0)
	fc, _ := os.Open("./jobtmp")
	e := gob.NewDecoder(fc)
	e.Decode(&updatas)

	fc1, _ := os.Open("./jobtmp1")
	e1 := gob.NewDecoder(fc1)
	e1.Decode(&subitem)
	updatas = MergeUpdates(updatas, subitem)

	for _, update := range subitem {
		fmt.Println(update["检验项目"], "=>", update["检验结果"])
	}

	for _, update := range updatas {
		fmt.Println(update["item"], "=>", update["sp_data_1"])
	}
}

func TestTest_platform_foodTest_foodDetail(t *testing.T) {
	ck := "JSESSIONID=C248ACBED0C85FC725157B9ECE50066C-n3;sod=leYDFuSnkFivrosGM5PIfUP972YJATho+2y9TVZBqWDT6EZLBtjURt+FfyxhjeremyFLs5gYY4yfC2qsdjk="
	r, err := Test_platform_foodTest_foodDetail(15042534, true, ck, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("抽样基础信息_抽样单编号:", r.Get("抽样基础信息_抽样单编号").ToString())
	t.Log(r.ToString())
}
