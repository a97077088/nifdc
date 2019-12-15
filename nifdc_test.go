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

<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>检验完成</title>
        
    <link rel="shortcut icon" href="/Public/cfda.ico">
    <link href="/Public/AdminNew/css/bootstrap.min.css" rel="stylesheet">
    <link href="/Public/AdminNew/css/font-awesome.min.css" rel="stylesheet">
    <link href="/Public/AdminNew/css/animate.min.css" rel="stylesheet">
    <link href="/Public/AdminNew/css/datatables/dataTables.bootstrap.css" rel="stylesheet">
    <link href="/Public/AdminLTE/css/style.css?v=1.3.6" rel="stylesheet">
    <link href="/Public/AdminLTE/css/maintain.css?v=1.3.6" rel="stylesheet">
</head>
<body class="gray-bg">
	<div class="wrapper wrapper-content">
		<div class="row">
			<div class="col-xs-12 animated fadeInUp">
                <div class="mb20">

                        
                </div>
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
                                        任务来源：邯郸市市场监督管理局                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        报送分类：抽检监测（市级本级）2019年河北邯郸抽检计划E包                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        检验机构名称：河南恒晟检测技术有限公司                                    </div>
                                    <div class="col-xs-6 pl0">
                                        抽样类型：常规抽样                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        部署机构：邯郸市市场监督管理局                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        抽样环节：餐饮                                    </div>
                                    <div class="col-xs-6 pl0">
                                        抽样地点：中型餐馆                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        食品分类：调味品                                        调味料                                        液体复合调味料                                        其他液体调味料                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        抽样单编号： DC19130400463233304                                    </div>
                                    <div class="col-xs-6 pl0">
                                        检验目的/任务类别：   监督抽检                                    </div>
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
                                        单位名称：河南恒晟检测技术有限公司                                                                                </div>
                                    <!-- <div class="widTone pull-left">
                                        单位级别：省（区）级                                    </div> -->
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widthtwo pull-left">
                                        单位地址：新郑市薛店镇中德产业园2-3号                                                                                </div>
                                    <div class="widTone pull-left">
                                       所在省份：河南                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        抽样人员：周春伟、徐明欣                                    </div>
                                    <div class="widTone pull-left">
                                        联系人：王力                                                                                </div>
                                    <div class="widTone pull-left">
                                        电子邮箱：hengshengjiance23@163.com                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        电话：
                                        0371-55929768                                                                                </div>
                                    <div class="widTone pull-left">
                                        传真：
                                        0371-55929768                                                                                </div>
                                </div>
                            </div>
                        </div>
                        <!--抽样单位信息end-->
                        <!--抽样场所信息start-->
                        <div class="formArea1 formAreaCom">
                            <h6 class="areaTitle">抽检场所信息</h6>
                            <div class="areaContent">
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        <span class="pull-left insLabelForm">所在地：河北 邯郸                                        魏县</span>
                                    </div>
                                    <div class="col-xs-6 pl0">
                                        <span class="pull-left insLabelForm ar">区域类型：乡村</span>
                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        <span class="pull-left insLabelForm">单位名称：邯郸市马头生态工业城天外天饭庄</span>
                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        <span class="pull-left insLabelForm ar">单位地址：马头镇西村107国道东</span>
                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                                                            <div class="widTone pull-left">
                                            营业执照/社会信用代码：92130492MA09EEP68X                                        </div>
                                        <div class="widTone pull-left">
                                            许可证类型：生产许可证                                        </div>                                    <div class="widTone pull-left">
                                        许可证号：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        法人代表：晋保                                    </div>
                                    <div class="widTone pull-left">
                                        联系人：花荣                                    </div>
                                    <div class="widTone pull-left">
                                        联系电话：15027989926                                    </div>
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
                                            所在地：安徽                                            马鞍山                                            和县                                        </div>
                                    </div>
                                    <div class="form-group clearfix">
                                        <div class="col-xs-12 pl0">
                                            生产者地址：和县西埠镇盛家口工业园                                        </div>
                                    </div>
                                    <div class="form-group clearfix">
                                        <div class="col-xs-6 pl0">
                                            生产者名称：和县兴盛麻油有限责任公司                                        </div>
                                        <div class="col-xs-6">
                                            生产许可证编号：SC10234052305797                                        </div>
                                    </div>
                                    <div class="form-group clearfix">
                                        <div class="widTone pull-left">
                                            联系电话：400-634-8858                                        </div>
                                        <div class="widTone pull-left pl10">
                                            是否存在第三方企业信息：是                                        </div>
                                    </div>
                                    <div class="wtsc clearfix">
                                            <h6 class="title" style="color: #6079d4;font-size: 14px">第三方企业相关信息</h6>
                                            <div class="form-group clearfix">
                                                <div class="col-xs-12 pl0">
                                                    第三方企业所在地：四川                                                    成都                                                    高新区                                                </div>
                                            </div>
                                            <div class="form-group clearfix">
                                                <div class="widthhalf">
                                                    第三方企业地址：成都高新区富华南路1606号7栋1单元4层406号                                                </div>
                                            </div>
                                            <div class="form-group clearfix">
                                                <div class="col-xs-6 pl0">
                                                    第三方企业名称：成都汉原藜红商贸有限公司                                                </div>
                                                <div class="col-xs-6">
                                                    第三方企业许可证编号：/                                                </div>
                                            </div>
                                            <div class="form-group clearfix">
                                                <div class="widTone pull-left">
                                                    第三方企业电话：/                                                </div>
                                                <div class="widTone pull-left">
                                                    第三方企业性质：委托                                                </div>
                                            </div>
                                        </div>                                </div>                        </div>
                        <!--抽样生产企业信息end-->
                        <!--抽样样品信息start-->
                                                    <div class="formArea1 formAreaCom">
                            <h6 class="areaTitle">抽检样品信息</h6>
                            <div class="areaContent">
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        样品条码：6931547995969                                    </div>
                                    <div class="col-xs-3">
                                        样品商标：汉榞簗红                                    </div>
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
                                    <div class="col-xs-3 ">
                                        样品名称：花椒油                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        生产日期：2019-06-02                                    </div>
                                    <div class="col-xs-3">
                                        保质期：18个月                                    </div>
                                    <div class="col-xs-3">
                                        样品批号：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        规格型号：265ml/瓶                                    </div>
                                    <div class="col-xs-3 pl0">
                                        质量等级：合格品                                    </div>
                                    <div class="col-xs-3 pl0">
                                        单价：//kg                                    </div>
                                    <div class="col-xs-3 pl0">
                                        是否进口：否                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样日期：2019-11-27                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样方式：非无菌采样                                    </div>
                                    <div class="col-xs-3 pl0">
                                        原产地：中国                                    </div>
                                    <div class="col-xs-3 pl0">
                                        储存条件：常温                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样基数：2.65kg                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样数量：1.5                                    </div>
                                    <div class="col-xs-3 pl0">
                                        备样数量：0.5kg                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样数量单位：kg                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-15 pl0">
                                        执行标准/技术文件：Q/HXS0003S                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-15 pl0">
                                        整个样品备注：样品现场分装                                    </div>
                                </div>
                            </div>
                        </div>                        <!--抽样样品信息end-->
                        <!--抽样生产企业信息start-->
                        <div class="formArea1 formAreaCom">
                            <h6 class="areaTitle">现场照片</h6>
                            <div class="areaContent pb30">
                                <ul class="addList clearfix">
                                    <li>
                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/15748335846929.726.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://upload1.nifdc.org.cn/image/2019/11/27/15748335846929.726.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358413664840.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358413664840.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358399135620.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358399135620.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358411125172.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358411125172.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/1574833584936194.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://upload1.nifdc.org.cn/image/2019/11/27/1574833584936194.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358440734448.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358440734448.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358456283919.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358456283919.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358463687649.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://upload1.nifdc.org.cn/image/2019/11/27/157483358463687649.png">
                                                </a>
                                            </li>                                </ul>
                            </div>
                        </div>
                        <!--抽样生产企业信息end-->
                        <!--告知书&抽样单start-->
                        <div class="formArea1 formAreaCom attention">
                            <h6 class="areaTitle">告知书&抽样单</h6>
                            <div class="areaContent pb30 ">
                                <span class="downResult btn-sm cydUpload" id="">
                                <i>抽样单:
                                                                        <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/157483355917104470.png" data-magnify="gallery" data-group="g2">
                                        157483355917104470.png</a>
                                                                    </i>
                                    </span>
                                <span class="downResult btn-sm gzsUpload">
                                    <i>告知书:
                                                                                <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/11/27/157483356493634216.png" data-magnify="gallery" data-group="g2">
                                            157483356493634216.png</a>
                                                                            </i>
                                </span>
                            </div>
                        </div>
                        <!--抽告知书&抽样单end-->
                        <div class="ibox-title clearfix pl20 borderline">
                            检验信息
                        </div>
                        <div class="formAreaCom">
                            <div class="areaContent">
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        检验机构名称：河南恒晟检测技术有限公司                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        样品送达日期：2019-11-27                                    </div>
                                    <div class="widTone pull-left pl10">
                                        联系人：王力                                    </div>
                                    <div class="widTone pull-left pl10">
                                        联系人电话：0371-55929768                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        联系人邮箱：hengshengjiance23@163.com                                    </div>
                                    <div class="widTone pull-left pl10">
                                        检查封样人员：河南恒晟检测技术有限公司                                    </div>
                                    <div class="widTone pull-left pl10">
                                        系统接样日期：2019-12-02                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        检查封样人电话：15738889730                                    </div>
                                    <div class="widTone pull-left pl10">
                                        检查封样人邮箱：rdh123@163.com                                    </div>
                                </div>
                                
                                                            </div>
                        </div>					</div>
				</div>
            </div>
		</div>	
	</div>
	<script src="/Public/AdminNew/js/jquery.min.js"></script>
	<script src="/Public/AdminNew/js/bootstrap.min.js"></script>
	<script type="text/javascript" src="/Public/AdminNew/js/plugins/dataTables/jquery.dataTables.js"></script>
	<script type="text/javascript" src="/Public/AdminNew/js/plugins/dataTables/dataTables.bootstrap.js"></script>
    <script src="/Public/AdminNew/plugins/laydate/laydate.js?v=2.1.4"></script>
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
</body>
</html>

`

	mkr := StoMap_yijieshou_full(olds)
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
	itemsr,err:=Test_platform_api_food_getTestItems(fddetail,test_platform_ck,nil)
	if err != nil {
		t.Fatal(err)
	}
	testinfor,err:=Test_platform_api_food_getTestInfo(sd,test_platform_ck,nil)
	if err != nil {
		t.Fatal(err)
	}

	rmp:=TestInfotoMap(testinfor.Rows,itemsr.Rows)
	fmt.Println(rmp)
	for k,v:=range rmp[0]{
		fmt.Printf("%s:%s\n",k,v)
	}
}