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
                                    任务来源：晋中市市场监督管理局                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    报送分类：抽检监测（市级本级）2019年山西晋中抽检计划                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    检验机构名称：河南恒晟检测技术有限公司                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    部署机构：晋中市市场监督管理局                                </div>
                                <div class="col-xs-6 pl0">
                                    抽样类型：常规抽样                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    抽样环节：流通                                </div>
                                <div class="col-xs-6 pl0">
                                    抽样地点：超市                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    食品分类：食用油、油脂及其制品                                    食用植物油(含煎炸用油)                                    食用植物油(半精炼、全精炼)                                    其他食用植物油(半精炼、全精炼)                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    抽样单编号： DC19142400463230158                                </div>
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
                                    抽样人员：曹永昌、谢天赐                                </div>
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
                                    <span class="pull-left insLabelForm">所在地：山西                                        晋中                                        榆次</span>
                                </div>
                                <div class="widTone pull-left">
                                    <span class="pull-left insLabelForm ar">区域类型：城市</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    <span class="pull-left insLabelForm">单位名称：晋中市榆次区熊杰商行</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    <span class="pull-left insLabelForm ar">单位地址：晋中市榆次区蕴华街218号兴达商务楼</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                                                    <div class="widTone pull-left">
                                        <span class="pull-left insLabelForm">营业执照/社会信用代码：92140702MA0K2AX62M</span>
                                    </div>
                                    <div class="widTone pull-left xuz pl10">
                                        <span class="pull-left insLabelForm" id="sp_s_sfjk">许可证类型：经营许可证</span>
                                    </div>                                <div class="widTone pull-left pl10">
                                    <span class="pull-left insLabelForm ar">经营许可证号：JY11407020031943</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="widTone pull-left">
                                    法人代表：武志强                                </div>
                                <div class="widTone pull-left  pl10">
                                    联系人：武志强                                </div>
                                <div class="widTone pull-left">
                                    联系电话：13038032121                                </div>

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
                                        所在地：天津                                        宝坻                                        宝坻                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        生产者地址：天津自贸试验区（天津港保税区）津滨大道95号                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        生产者名称：嘉里粮油（天津）有限公司（代码TJN）                                    </div>
                                    <div class="col-xs-6">
                                        生产许可证编号：SC10212011600523                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        联系电话：400-616-5757                                    </div>

                                    <div class="widTone pull-left pl10">
                                        是否存在第三方企业信息：是                                    </div>
                                </div>
                                <div class="wtsc clearfix">
                                        <h6 class="title" style="color: #6079d4;font-size: 14px">第三方企业相关信息</h6>
                                        <div class="form-group clearfix">
                                            <div class="col-xs-12 pl0">
                                                所在地：上海                                                浦东                                                浦东                                            </div>
                                        </div>
                                        <div class="form-group clearfix">
                                            <div class="col-xs-12 pl0">
                                                第三方企业地址：上海市浦东新区光明路718号715室                                            </div>
                                        </div>
                                        <div class="form-group clearfix">
                                            <div class="col-xs-6 pl0">
                                                第三方企业名称：益海嘉里食品营销有限公司                .                                            </div>
                                            <div class="col-xs-6">
                                                第三方企业许可证编号：/                                            </div>
                                        </div>
                                        <div class="form-group clearfix">
                                            <div class="widTone pull-left">
                                                第三方企业电话：400-616-5757                                            </div>
                                            <div class="widTone pull-left">
                                                第三方企业性质：委托                                            </div>
                                        </div>
                                    </div>                            </div>                    </div>
                    <!--抽样生产企业信息end-->
                    <!--抽样样品信息start-->
                                            <div class="formArea1 formAreaCom">
                            <h6 class="areaTitle">抽检样品信息</h6>
                            <div class="areaContent">
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        样品条码：6948195800323                                    </div>
                                    <div class="col-xs-3">
                                        样品商标：金龙鱼                                    </div>
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
                                        样品名称：金龙鱼精炼一级大豆油                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        生产日期：2019-08-04                                    </div>
                                    <div class="col-xs-3">
                                        保质期：18个月                                    </div>
                                    <div class="col-xs-3">
                                        样品批号：20190804                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        规格型号：1.8升/瓶                                    </div>
                                    <div class="col-xs-3 pl0">
                                        质量等级：一级                                    </div>
                                    <div class="col-xs-3 pl0">
                                        单价：20.00/瓶/瓶                                    </div>
                                    <div class="col-xs-3 pl0">
                                        是否进口：否                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样日期：2019-11-19                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样方式：非无菌采样                                    </div>
                                    <div class="col-xs-3 pl0">
                                        原产地：中国                                    </div>
                                    <div class="col-xs-3 pl0">
                                        储存条件：常温                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样基数：12瓶瓶                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样数量：2                                    </div>
                                    <div class="col-xs-3 pl0">
                                        备样数量：1瓶瓶                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样数量单位：瓶                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        执行标准/技术文件：Q/BBAH0019S                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        整个样品备注：/                                    </div>
                                </div>
                            </div>
                        </div>                    <!--抽样样品信息end-->
                    <!--抽样生产企业信息start-->
                    <div class="formArea1 formAreaCom">
                        <h6 class="areaTitle">现场照片</h6>
                        <div class="areaContent pb30">
                            <ul class="addList clearfix">
                                <li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413663735814364.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/11/19/157413663735814364.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413665925017832.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/11/19/157413665925017832.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666989840531.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666989840531.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666969971749.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666969971749.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666972029773.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666972029773.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666938232343.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666938232343.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666936059190.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/11/19/157413666936059190.png">
                                            </a>
                                        </li><li>
                                            <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413667013119615.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                <img src="http://upload1.nifdc.org.cn/image/2019/11/19/157413667013119615.png">
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
                                                                        <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413669159015367.png" data-magnify="gallery" data-group="g2">
                                        157413669159015367.png</a>
                                                                    </i>
                                    </span>
                            <span class="downResult btn-sm gzsUpload">
                                <i>告知书:
                                                                        <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/19/157413667436911127.png" data-magnify="gallery" data-group="g2">
                                            157413667436911127.png</a>
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
         if("2" != 33){
             printTitle = [
                 "第一联 组织抽样检验的市场监管部门存",
                 "第二联 承检机构存",
                 "第三联 （标称）食品生产者或境内代理商存",
                 "第四联 抽样单位存",
                 "第五联 被抽样单位存"
             ];
         }
        var AjaxUrl = '/index.php?m=Admin&c=Sample&a=printSample';
        function printSample (){
            $("#printModal").modal("show");
        }
        function submit(type) {
            var code = "DC19142400463230158";
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

	//StoMap_chouyangwancheng_full(olds)
	mkr := StoMap_chouyangwancheng_full(olds)
	for k, v := range mkr {
		fmt.Printf("%s:%s\n", k, v)
	}
}
