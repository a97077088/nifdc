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
                                        任务来源：鹤壁市市场监督管理局鹤山分局                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        报送分类：市县级农产品专项抽检2020年河南鹤壁鹤山区抽检计划                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        检验机构名称：河南恒晟检测技术有限公司                                    </div>
                                    <div class="col-xs-6 pl0">
                                        抽样类型：农产品抽样                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        部署机构：鹤壁市市场监督管理局鹤山分局                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        抽样环节：流通                                    </div>
                                    <div class="col-xs-6 pl0">
                                        抽样地点：农贸市场                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        食品分类：食用农产品                                        畜禽肉及副产品                                        畜肉                                        猪肉                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        抽样单编号： NCP20410602463230031                                    </div>
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
                                        抽样人员：凌振峰、杜鹏举、杜冰                                    </div>
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
                                        <span class="pull-left insLabelForm">所在地：河南 鹤壁                                        鹤山区</span>
                                    </div>
                                    <div class="col-xs-6 pl0">
                                        <span class="pull-left insLabelForm ar">区域类型：城市</span>
                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        <span class="pull-left insLabelForm">单位名称：鹤壁市鹤山区秀成生肉店</span>
                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        <span class="pull-left insLabelForm ar">单位地址：鹤山区鹤壁集惠恩贸易中心内</span>
                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                                                            <div class="widTone pull-left">
                                            营业执照/社会信用代码：92410602MA457NW89C                                        </div>
                                        <div class="widTone pull-left">
                                            许可证类型：经营许可证                                        </div>                                    <div class="widTone pull-left">
                                        许可证号：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        法人代表：高雅                                    </div>
                                    <div class="widTone pull-left">
                                        联系人：高雅                                    </div>
                                    <div class="widTone pull-left">
                                        联系电话：18539233111                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                        <div class="widTone pull-left">
                                            摊位号或姓名：高雅                                        </div>
                                        <div class="widTone pull-left">
                                            身份证号：410602199603161026                                        </div>
                                    </div>                                                            </div>
                        </div>
                        <!--抽样场所信息end-->
                        
                        <!--抽样生产企业信息start-->
                        <div class="formArea1 formAreaCom">
                                                            <h6 class="areaTitle">抽样生产企业信息</h6>
                                <div class="areaContent">
                                    <div class="form-group clearfix">
                                        <div class="col-xs-12 pl0">
                                            所在地：河南                                            鹤壁                                            鹤山区                                        </div>
                                    </div>
                                    <div class="form-group clearfix">
                                        <div class="col-xs-12 pl0">
                                            生产者地址：/                                        </div>
                                    </div>
                                    <div class="form-group clearfix">
                                        <div class="col-xs-6 pl0">
                                            生产者名称：/                                        </div>
                                        <div class="col-xs-6">
                                            生产许可证编号：/                                        </div>
                                    </div>
                                    <div class="form-group clearfix">
                                        <div class="widTone pull-left">
                                            联系电话：/                                        </div>
                                        <div class="widTone pull-left pl10">
                                            是否存在第三方企业信息：否                                        </div>
                                    </div>
                                                                    </div>                        </div>
                        <!--抽样生产企业信息end-->
                        <!--抽样样品信息start-->
                                                    <div class="formArea1 formAreaCom">
                            <h6 class="areaTitle">抽检样品信息</h6>
                            <div class="areaContent">
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        样品条码：/                                    </div>
                                    <div class="col-xs-3">
                                        样品商标：/                                    </div>
                                    <div class="col-xs-3">
                                        样品类型：食用农产品                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        样品来源：外购                                    </div>
                                    <div class="col-xs-3 pl0">
                                        样品属性：普通食品                                    </div>
                                    <div class="col-xs-3">
                                        包装分类：无包装                                    </div>
                                    <div class="col-xs-3 ">
                                        样品名称：猪五花肉                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        购进日期：2020-01-20                                    </div>
                                    <div class="col-xs-3">
                                        保质期：/                                    </div>
                                    <div class="col-xs-3">
                                        样品批号：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        规格型号：/                                    </div>
                                    <div class="col-xs-3 pl0">
                                        质量等级：/                                    </div>
                                    <div class="col-xs-3 pl0">
                                        单价：54元/kg                                    </div>
                                    <div class="col-xs-3 pl0">
                                        是否进口：否                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样日期：2020-01-20                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样方式：非无菌采样                                    </div>
                                    <div class="col-xs-3 pl0">
                                        原产地：中国                                    </div>
                                    <div class="col-xs-3 pl0">
                                        储存条件：常温                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样基数：20kg                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样数量：1.5                                    </div>
                                    <div class="col-xs-3 pl0">
                                        备样数量：0.5kg                                    </div>
                                    <div class="col-xs-3 pl0">
                                        抽样数量单位：Kg                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-15 pl0">
                                        执行标准/技术文件：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-15 pl0">
                                        整个样品备注：以上信息由被抽样单位提供                                    </div>
                                </div>
                            </div>
                        </div>                        <!--抽样样品信息end-->
                        <!--抽样生产企业信息start-->
                        <div class="formArea1 formAreaCom">
                            <h6 class="areaTitle">现场照片</h6>
                            <div class="areaContent pb30">
                                <ul class="addList clearfix">
                                    <li>
                                                <a href="javascript:void(0)" data-src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/15794988411266.835.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/15794988411266.835.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884015823752.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884015823752.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884045359761.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884045359761.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884024170830.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884024170830.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884045945427.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884045945427.png">
                                                </a>
                                            </li><li>
                                                <a href="javascript:void(0)" data-src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884119255399.png" data-magnify="gallery" data-group="g1" data-caption="">
                                                    <img src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949884119255399.png">
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
                                                                        <a href="javascript:void(0)" data-src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949885021321168.png" data-magnify="gallery" data-group="g2">
                                        157949885021321168.png</a>
                                                                    </i>
                                    </span>
                                <span class="downResult btn-sm gzsUpload">
                                    <i>告知书:
                                                                                <a href="javascript:void(0)" data-src="http://spcjupload2.gsxt.gov.cn/image/2020/01/20/157949885666614315.png" data-magnify="gallery" data-group="g2">
                                            157949885666614315.png</a>
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
                                        样品送达日期：2020-01-20                                    </div>
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
                                        系统接样日期：2020-02-01                                    </div>
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
	fmt.Println(mkr["抽检样品信息_生产日期"])
	for k, v := range mkr {
		fmt.Printf("%s:%s\n", k, v)
	}
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