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

func TestRe1(t *testing.T) {
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
            <small class="text-danger">&nbsp;&nbsp;&nbsp;&nbsp; 抽样编号:XC19410223463231108</small>
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
                                    <div class="col-sm-8">尉氏县市场监督管理局</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">报送分类A：</label>
                                    <div class="col-sm-8" id="bsfla">抽检监测（县级本级）
                                        <input type="hidden" id="hid_bsfla" value="抽检监测（县级本级）"/>
                                    </div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">报送分类B：</label>
                                    <div class="col-sm-8" id="bsflb">2019年河南开封尉氏抽检计划（河南恒晟检测）
                                        <input type="hidden" id="hid_bsflb" value="2019年河南开封尉氏抽检计划（河南恒晟检测）"/>
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
                                    <div class="col-sm-8" id="type2">调味料
                                        <input type="hidden" id="hid_type2" value="调味料"/>
                                    </div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">食品次亚类：</label>
                                    <div class="col-sm-8" id="type3">固体复合调味料
                                        <input type="hidden" id="hid_type3" value="固体复合调味料"/>
                                    </div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">食品细类：</label>
                                    <div class="col-sm-8" id="type4">鸡粉、鸡精调味料
                                        <input type="hidden" id="hid_type4" value="鸡粉、鸡精调味料"/>
                                    </div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样单编号：</label>
                                    <div class="col-sm-8">XC19410223463231108</div>
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
                                    <div class="col-sm-8">宋自钟、陆顺</div>
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
                                    <div class="col-sm-8">尉氏县水坡镇新安购物广场</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单位地址：</label>
                                    <div class="col-sm-8">尉氏县水坡镇张寨村</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">所在地：</label>
                                    <div class="col-sm-8">河南/开封/尉氏</div>
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
                                    <div class="col-sm-8">超市</div>
                                </div>

                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">营业执照/社会信用代码：</label>
                                    <div class="col-sm-8">92410223MA421E923G</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">许可证类型：</label>
                                    <div class="col-sm-8">经营许可证</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">许可证号：</label>
                                    <div class="col-sm-8">JY14102230000381</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单位法人：</label>
                                    <div class="col-sm-8">王永辉</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">联系人：</label>
                                    <div class="col-sm-8">王永辉</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">联系人电话：</label>
                                    <div class="col-sm-8">13569526491</div>
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
                                    <div class="col-sm-8">广东佳隆食品股份有限公司英格山分公司</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">所在地：</label>
                                    <div class="col-sm-8">广东/揭阳/普宁</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">企业地址：</label>
                                    <div class="col-sm-8">普宁市大坝镇陂乌村英歌山片区</div>
                                </div>
                            </div>
                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">生产许可证编号：</label>
                                    <div class="col-sm-8">SC10344528101740</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">联系人：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">电话：</label>
                                    <div class="col-sm-8">800-888-6101</div>
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
                                    <div class="col-sm-8">6909605001853</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品商标：</label>
                                    <div class="col-sm-8">佳隆</div>
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
                                    <div class="col-sm-8">鸡精调味料</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">生产日期:</label>
                                    <div class="col-sm-8">2019-03-10</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">保质期：</label>
                                    <div class="col-sm-8">18个月</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">样品批号：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">规格型号：</label>
                                    <div class="col-sm-8">180g/袋</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">质量等级：</label>
                                    <div class="col-sm-8">/</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">单价：</label>
                                    <div class="col-sm-8">6元/袋/袋</div>
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
                                    <div class="col-sm-8">2019-08-06</div>
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
                                    <div class="col-sm-8">10袋</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样数量：</label>
                                    <div class="col-sm-8">4.0</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">备样数量：</label>
                                    <div class="col-sm-8">2袋</div>
                                </div>
                            </div>

                            <div class="row form-group">
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">抽样数量单位：</label>
                                    <div class="col-sm-8">袋</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">执行标准/技术文件：</label>
                                    <div class="col-sm-8">SB/T 10371</div>
                                </div>
                                <div class="col-sm-4">
                                    <label class="control-label col-sm-4">备注：</label>
                                    <div class="col-sm-8">样品信息由被抽样单位提供</div>
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
                                                     src="http://upload.nifdc.org.cn/image/2019/08/06/156509704269792141.png"    data-original="http://upload.nifdc.org.cn/image/2019/08/06/156509704269792141.png"
                                                     alt="现场抽样图片">
                                                <p>图片1</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload.nifdc.org.cn/image/2019/08/06/156509704095765178.png"    data-original="http://upload.nifdc.org.cn/image/2019/08/06/156509704095765178.png"
                                                     alt="现场抽样图片">
                                                <p>图片2</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload.nifdc.org.cn/image/2019/08/06/156509704549669133.png"    data-original="http://upload.nifdc.org.cn/image/2019/08/06/156509704549669133.png"
                                                     alt="现场抽样图片">
                                                <p>图片3</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload.nifdc.org.cn/image/2019/08/06/156509704398317231.png"    data-original="http://upload.nifdc.org.cn/image/2019/08/06/156509704398317231.png"
                                                     alt="现场抽样图片">
                                                <p>图片4</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload.nifdc.org.cn/image/2019/08/06/15650970423332.851.png"    data-original="http://upload.nifdc.org.cn/image/2019/08/06/15650970423332.851.png"
                                                     alt="现场抽样图片">
                                                <p>图片5</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload.nifdc.org.cn/image/2019/08/06/156509704271814263.png"    data-original="http://upload.nifdc.org.cn/image/2019/08/06/156509704271814263.png"
                                                     alt="现场抽样图片">
                                                <p>图片6</p>
                                            </div>
                                        
                                            <div style="display: inline-block;  overflow: hidden">
                                                <img style="width:150px;height:150px;margin-right: 15px;"
                                                     src="http://upload.nifdc.org.cn/image/2019/08/06/156509704463378681.png"    data-original="http://upload.nifdc.org.cn/image/2019/08/06/156509704463378681.png"
                                                     alt="现场抽样图片">
                                                <p>图片7</p>
                                            </div>
                                        
                                    </ul>
                                </div>
                            </div>
                            <div class="row form-group">
                                <div class="col-sm-12">
                                    
                                        
                                        
                                            <a href="http://upload.nifdc.org.cn/image/2019/08/06/156509709493965972.png" target="_blank"
                                               class="btn btn-danger btn-xs"><i
                                                    class="fa fa-search"></i> 抽样单电子版</a>
                                        
                                    
                                    
                                        
                                        
                                            <a href="http://upload.nifdc.org.cn/image/2019/08/06/156509705113449865.png" target="_blank"
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
                                    <label class="control-label">2019（食）A2096</label>

                                </div>
                            </div>
                            <div class="col-sm-4 form-group">
                                <label class="control-label col-sm-5">样品到达日期：</label>
                                <div class="col-sm-7">
                                    <label class="control-label">
                                        2019-08-06</label>

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
                                        
                                            2019-08-12
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
                                    经抽样检验，所检项目符合 GB 2760-2014《食品安全国家标准 食品添加剂使用标准》，SB/T 10371-2003《鸡精调味料》，GB 2762-2017《食品安全国家标准 食品中污染物限量》 要求。
                                </p>
                            </div>
                        </div>
                        <input type="hidden" id="sd" name="sd" value="+834QSQYyS6UYPKbNvh03CLyxhjeremyM99HrtJo">
                        <input type="hidden" id="st" name="st" value="4">
                        <input type="hidden" id="sis" name="sis" value="9">
                        <input type="hidden" id="rd" name="rd" value="4">
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
                                                            onclick="detailManager.viewReport('http://111.204.20.233:39054/report/jd_report/4907/20190831/XC19410223463231108_JDCJ_33c3b2b6-d0e7-449a-a6d0-51d877360825.pdf','查看监督抽检报告')"
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
<script src="/test_platform/js/plugins/fancybox/jquery.fancybox.js?v=1.2.0.42"></script>
<script src="/test_platform/easyui/jquery.easyui.min.js?v=1.2.0.42"></script>
<script src="/test_platform/js/plugins/fancybox/viewer.min.js"></script>
<script src="/test_platform/easyui/locale/easyui-lang-zh_CN.js?v=1.2.0.42"></script>
<script src="/test_platform/js/uploadTool.js"></script>
<script src="/test_platform/js/food/foodDetailReadOnly.js?v=1.2.0.42"></script>
</body>
</html>
`

	mkr := StoMap_foodDetail(olds)
	for k, v := range mkr {
		fmt.Printf("%s:%s\n", k, v)
	}
}
