package nifdc

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
}

func TestRe(t *testing.T){
	olds:=`

<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>查看正常抽样单--已完成</title>
    <link rel="shortcut icon" href="/Public/cfda.ico">
    <link href="/Public/AdminNew/css/bootstrap.min.css" rel="stylesheet">
    <link href="/Public/AdminNew/css/font-awesome.min.css" rel="stylesheet">
    <link href="/Public/AdminNew/css/animate.min.css" rel="stylesheet">
    <link href="/Public/AdminNew/css/datatables/dataTables.bootstrap.css" rel="stylesheet">
    <link href="/Public/AdminNew/css/plugins/sweetalert/sweetalert2.min.css" rel="stylesheet">
    <link href="/Public/AdminLTE/css/style.css?v=1.3.6" rel="stylesheet">
    <link href="/Public/AdminLTE/css/maintain.css?v=1.3.6" rel="stylesheet">

</head>
<body class="gray-bg">
<div class="wrapper wrapper-content">
    <div class="row">
        <div class="col-xs-12 animated fadeInUp">
            <div class="ibox clearfix">
                <div class="ibox-title clearfix pl20 borderline">
                    <h5>国家食品安全抽样检验抽样单</h5>
                </div>
                <div class="area">
                    <!--抽样基础信息start-->
                    <div class="formArea1 formAreaCom">
                        <h6 class="areaTitle">抽样基础信息</h6>
                        <div class="areaContent">
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    任务来源：郑州市市场监督管理局                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    报送分类：抽检监测（市级专项）2019年河南郑州（百荣市场）专项抽检计划                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    检验机构名称：河南恒晟检测技术有限公司                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    部署机构：郑州市市场监督管理局                                </div>
                                <div class="col-xs-6 pl0">
                                    抽样类型：常规抽样                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    抽样环节：流通                                </div>
                                <div class="col-xs-6 pl0">
                                    抽样地点：批发市场                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    食品分类：肉制品                                    预制肉制品                                    调理肉制品                                    调理肉制品(非速冻)                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    抽样单编号： DC19410100463233772                                </div>
                                <div class="col-xs-6 pl0">
                                    检验目的/任务类别： 监督抽检                                </div>
                            </div>
                        </div>
                    </div>
                    <!--抽样基础信息end-->
                    <!--抽样单位信息start-->
                    <div class="formArea1 formAreaCom">
                        <h6 class="areaTitle">抽样单位信息</h6>
                        <div class="areaContent">
                            <div class="form-group clearfix">
                                <div class="widthtwo pull-left">
                                    单位名称：河南恒晟检测技术有限公司                                                                        </div>
                                <!-- <div class="widTone pull-left">
                                    单位级别：省（区）级                                </div> -->
                            </div>
                            <div class="form-group clearfix">
                                <div class="widthtwo pull-left">
                                    单位地址：新郑市薛店镇中德产业园2-3号                                                                        </div>
                                <div class="widTone pull-left">
                                    所在省份：河南                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="widTone pull-left">
                                    抽样人员：周春伟、王可行                                </div>
                                <div class="widTone pull-left">
                                    联系人：王力                                                                        </div>
                                <div class="widTone pull-left">
                                    电子邮箱：hengshengjiance23@163.com                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="widTone pull-left">
                                    电话：
                                    0371-55929768                                                                        </div>
                                <div class="widTone pull-left">
                                    传真：
                                    0371-55929768                                                                        </div>
                                <div class="widTone pull-left">
                                    邮编：
                                    451162                                                                        </div>
                            </div>
                        </div>
                    </div>
                    <!--抽样单位信息end-->
                    <!--抽样场所信息start-->
                    <div class="formArea1 formAreaCom">
                        <h6 class="areaTitle">抽检场所信息</h6>
                        <div class="areaContent">
                            <div class="form-group clearfix">
                                <div class="widthtwo pull-left">
                                    <span class="pull-left insLabelForm">所在地：河南                                        郑州                                        管城回族区</span>
                                </div>
                                <div class="widTone pull-left">
                                    <span class="pull-left insLabelForm ar">区域类型：城市</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    <span class="pull-left insLabelForm">单位名称：郑州市管城区亿兴食品商行</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    <span class="pull-left insLabelForm ar">单位地址：郑州市管城区百荣世贸商城D座2层D2-042号</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                                                    <div class="widTone pull-left">
                                        <span class="pull-left insLabelForm">营业执照/社会信用代码：92410104MA45J3NM1A</span>
                                    </div>
                                    <div class="widTone pull-left xuz pl10">
                                        <span class="pull-left insLabelForm" id="sp_s_sfjk">许可证类型：经营许可证</span>
                                    </div>                                <div class="widTone pull-left pl10">
                                    <span class="pull-left insLabelForm ar">经营许可证号：JY14101040103251</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="widTone pull-left">
                                    <span class="pull-left insLabelForm ar">
                                                                            年销售额：                                     /   
                                    </span>
                                </div>
                                <div class="widTone pull-left pl10">
                                    单位法人：王宁霞                                </div>
                                <div class="widTone pull-left  pl10">
                                    联系人：王宁霞                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="widTone pull-left">
                                    电话：13783717969                                </div>
                                <div class="widTone pull-left pl10">
                                    传真：/                                </div>
                                <div class="widTone pull-left pl10">
                                    邮编：/                                </div>
                            </div>
                                                                                </div>
                    </div>
                    <!--抽样场所信息end-->

                    <!--抽样生产企业信息start-->
                    <div class="formArea1 formAreaCom">
                                                    <h6 class="areaTitle">抽样生产企业信息</h6>
                            <div class="areaContent">
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        所在地：河南                                        郑州                                        经济技术开发区                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        企业地址：郑州经济技术开发区航海东路1897号                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        企业名称：双汇集团.郑州双汇食品有限责任公司                                    </div>
                                    <div class="col-xs-6">
                                        生产许可证编号：SC10441019300064                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        联系人：/                                    </div>
                                    <div class="widTone pull-left pl10">
                                        电话：400-6168218                                    </div>

                                    <div class="widTone pull-left pl10">
                                        是否存在第三方企业信息：否                                    </div>
                                </div>
                                                            </div>                    </div>
                    <!--抽样生产企业信息end-->
                    <!--抽样样品信息start-->
                                            <div class="formArea1 formAreaCom">
                            <h6 class="areaTitle">抽检样品信息</h6>
                            <div class="areaContent">
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        样品条码：/                                    </div>
                                    <div class="col-xs-3">
                                        样品商标：双汇                                    </div>
                                    <div class="col-xs-3">
                                        样品类型：工业加工食品                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        样品来源：外购                                    </div>
                                    <div class="col-xs-3 pl0">
                                        样品属性：普通食品                                    </div>
                                    <div class="col-xs-3">
                                        包装分类：预包装                                    </div>
                                    <div class="col-xs-3">
                                        样品名称：香辣风味肠                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        生产日期：2019-09-03                                    </div>
                                    <div class="col-xs-3">
                                        保质期：90天                                    </div>
                                    <div class="col-xs-3">
                                        样品批号：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        规格型号：180g/袋                                    </div>
                                    <div class="col-xs-3 pl0">
                                        质量等级：/                                    </div>
                                    <div class="col-xs-2">
                                        单价：10.00/袋                                    </div>
                                    <div class="col-xs-2">
                                        是否进口：否                                    </div>
                                    <div class="col-xs-2">
                                        原产地：中国                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样日期：2019-10-20                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样方式：非无菌采样                                    </div>
                                    <div class="col-xs-2">
                                        样品形态：固体                                    </div>
                                    <div class="col-xs-2">
                                        样品包装：塑料袋                                    </div>
                                    <div class="col-xs-2">
                                        抽样工具：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样时样品储存条件：常温                                    </div>
                                    <div class="col-xs-2 pl0">
                                        抽样基数：20袋                                    </div>
                                    <div class="col-xs-2">
                                        抽样数量：8                                    </div>
                                    <div class="col-xs-2">
                                        备样数量：2袋                                    </div>
                                    <div class="col-xs-3">
                                        抽样数量单位：袋                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        执行标准/技术文件：SB/T10279                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        备注：/                                    </div>
                                </div>
                            </div>
                        </div>                    <!--抽样样品信息end-->
                    <!--抽样生产企业信息start-->
                    <div class="formArea1 formAreaCom">
                        <h6 class="areaTitle">现场照片</h6>
                        <div class="areaContent pb30">
                            <ul class="addList clearfix">
                                <li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978932020948.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978932020948.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978943507133.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978943507133.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978935878596.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978935878596.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978938349157.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978938349157.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978933234538.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978933234538.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978960804224.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978960804224.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978963454435.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/10/20/157156978963454435.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156980820714909.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/10/20/157156980820714909.png">
                                            </a>
                                        </li>                            </ul>
                        </div>
                    </div>
                    <!--抽样生产企业信息end-->
                    <!--告知书&抽样单start-->
                    <div class="formArea1 formAreaCom attention">
                        <h6 class="areaTitle">告知书&抽样单</h6>
                        <div class="areaContent pb30 ">
                            <span class="downResult btn-sm cydUpload" id="">
                                <i>抽样单:
                                                                        <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156473551219965.png" data-magnify="gallery" data-group="g2">
                                        157156473551219965.png</a>
                                                                    </i>
                                    </span>
                            <span class="downResult btn-sm gzsUpload">
                                <i>告知书:
                                                                        <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/20/157156503312079739.png" data-magnify="gallery" data-group="g2">
                                            157156503312079739.png</a>
                                                                    </i>
                            </span>
                        </div>
                    </div>
                    <!--抽告知书&抽样单end-->
                </div>
            </div>
            <div class="ibox-content mt20">
                    <span class="btn btn-warning btn-sm pull-right" onclick="printSample()">
                        打印抽样单
                    </span>
            </div>
        </div>
    </div>
    <div class="modal inmodal " id="printModal" tabindex="-1" role="dialog"  aria-hidden="true" data-backdrop="static">
        <div class="modal-dialog  quoteDia" >
            <div class="modal-content animated fadeIn">
                <div class="modal-header bsflHead">
                    <div class="closeBg">
                        <button type="button" class="close mt0" data-dismiss="modal">
                        <span aria-hidden="true">&times;</span><span class="sr-only">Close</span>
                        </button>
                    </div>
                    <h4 class="modal-title text-center">签章打印</h4>
                </div>
                <div class="modal-body pb0">
                   <div class="area" >
                        <form method="post" ID="LoginForm" name="LoginForm" action="verifyLoginServlet">
                        <!--抽样基础信息start-->
                        <div class="formArea1 formAreaCom">
                            <div class="quoteContent pt0 pb0" id="verbjca" style="display: none">
                                <div class="form-group clearfix mb0">
                                    <div class="clearfix pb10">
                                        <select class="form-control" id="UserList" name="UserList">
                                            <option>请选择证书</option>
                                        </select>
                                    </div>
                                    <div class="clearfix pb10">
                                        <input type="password" class="form-control" placeholder="请输入证书密码！" id="UserPwd">
                                    </div>
                                    <div class="clearfix pb10">
                                        <input type="text" class="form-control" placeholder="请输入组织机构代码" id="sample_org">
                                    </div>
                                    
                                </div>
                            </div>
                            <div class="quoteMsg  pt0 pl20" id="pdfTitle" style="line-height: 25px;font-size: 14px;">
                                        
                            </div>

                        </div>
                        </form>
                    </div>
                </div>
                <div class="bjcaDowntip pt0 pl20">
                    如未安装以下证书助手和驱动，请先点击下载安装（两个均要安装）
                    <br/>
                    <a href="http://upload.nifdc.org.cn/image/drivers/cahelper.zip" target="_blank" class="mr25">1.安装证书助手</a>
                    <br/>
                    <a href="http://upload.nifdc.org.cn/image/drivers/ESealCoreSrv_Setup.zip" target="_blank">2.驱动</a>
                </div>
                <div class="modal-footer ac bsfl mt0">
                    <button type="button" class="btn btn-white mr25" data-dismiss="modal">取消</button>
                    <button type="button" class="btn btn-blueLight" id="creatPdf" onclick="submit()">生成pdf</button>               
                </div>
            </div>
        </div>
    </div>
    <script src="/Public/AdminNew/js/jquery.min.js"></script>
    <script src="/Public/AdminNew/js/bootstrap.min.js"></script>
    <script type="text/javascript" src="/Public/AdminNew/js/plugins/dataTables/jquery.dataTables.js"></script>
    <script type="text/javascript" src="/Public/AdminNew/js/plugins/dataTables/dataTables.bootstrap.js"></script>
    <script src="/Public/AdminNew/plugins/laydate/laydate.js?v=2.1.4"></script>
    <script src="/Public/AdminNew/js/plugins/sweetalert/sweetalert2.min.js"></script>
    <script src="/Public/AdminNew/js/base.js"></script>
    <script src="/Public/AdminNew/js/XTXSAB.min.js?v=1.0.3"></script>
    <script src="/Public/AdminNew/js/bjca/bjca.min.js?v=1.0.1"></script>
    <script src="/Public/AdminNew/js/jquery.rotate.index.min.js?v=1"></script>
    <script src="/Public/AdminNew/js/jquery.rotate.min.js?v=1"></script>
    <script>
        $(function () {
            $('[data-magnify]').Magnify({
                Toolbar: [
                    'prev',
                    'next',
                    'rotateLeft',
                    'rotateRight',
                    'zoomIn',
                    'actualSize',
                    'zoomOut'
                ],
                keyboard:true,
                draggable:true,
                movable:true,
                modalSize:[800,600],
                beforeOpen:function (obj,data) {
                    //console.log('beforeOpen')
                },
                opened:function (obj,data) {
                    //console.log('opened')
                },
                beforeClose:function (obj,data) {
                    //console.log('beforeClose')
                },
                closed:function (obj,data) {
                    // console.log('closed')
                },
                beforeChange:function (obj,data) {
                    //console.log('beforeChange')
                },
                changed:function (obj,data) {
                    //console.log('changed')
                }
            });

        })

    </script>
    <script>
         var printTitle = [
            "第一联 组织抽样检验的市场监管部门存",
            "第二联 负责核查处置工作的市场监管部门存",
            "第三联 标称食品生产者存",
            "第四联 抽样单位存",
            "第五联 被抽样单位存" 
        ];
        var AjaxUrl = '/index.php?m=Admin&c=Sample&a=printSample';
        function printSample (){
            $("#printModal").modal("show");
        }
        function submit(type) {
            var code = "DC19410100463233772";
            var msg = {
                "sample_code": code, 
                "source": "view"
            }
            if($("#UserList").val()!='' && $("#verbjca").is(':visible')){
                if($("#sample_org").val() ==''|| $("#sample_org").val() ==undefined){
                    swal({
                        title: '组织机构代码不能为空',
                        text: '',
                        type: 'error',
                        showConfirmButton: false,
                        timer: 1500,
                    });
                    return;
                };
                var org = $("#sample_org").val();
                var data = {
                    callFun:function(loginStatus){
                        if(loginStatus.status==0){
                            msg["org"]=org;
                            printSampleCall(msg);
                        }else{
                            swal({
                                title: loginStatus.msg,
                                text: '',
                                type: 'error',
                                timer: 1500,
                                showConfirmButton: false
                            });
                        }
                    }
                }
                login(data);
            }else printSampleCall(msg);
        }
    </script>
</body>
</html>`


	mkr:=StoMap(olds)
	for k,v:=range mkr{
		fmt.Println("k:",k)
		fmt.Println("v:",v)
	}
}