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
    <title>查看抽样单和检验信息--已接受</title>
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
                                <div class="col-xs-12 pl0">任务来源：运城市盐湖区市场监督管理局                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    报送分类：市县级农产品专项抽检2019年山西运城盐湖抽检计划                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    检验机构名称：河南恒晟检测技术有限公司                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    部署机构：运城市盐湖区市场监督管理局                                </div>
                                <div class="col-xs-6 pl0">
                                    抽样类型：农产品抽样                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    抽样环节：流通                                </div>
                                <div class="col-xs-6 pl0">
                                    抽样地点：其他                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    食品分类：食用农产品                                    蔬菜                                    茄果类蔬菜                                    辣椒                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-6 pl0">
                                    抽样单编号： NCP19142734463231660                                </div>
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
                                    抽样人员：
                                    凌振峰、陆顺                                </div>
                                <div class="widTone pull-left pl10">
                                    联系人：王力                                                                        </div>
                                <div class="widTone pull-left pl10">
                                    电子邮箱：hengshengjiance23@163.com                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="widTone pull-left">
                                    电话：
                                    0371-55929768                                                                        </div>
                                <div class="widTone pull-left pl10">
                                    传真：
                                    0371-55929768                                                                        </div>
                                <div class="widTone pull-left pl10">
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
                                        <span class="pull-left insLabelForm">所在地：山西                                            运城                                        盐湖</span>
                                </div>
                                <div class="widTone pull-left">
                                    <span class="pull-left insLabelForm ar">区域类型：城市</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    <span class="pull-left insLabelForm">单位名称：山西中苑宜购商贸有限公司</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="col-xs-12 pl0">
                                    <span class="pull-left insLabelForm ar">单位地址：运城市盐湖区禹都大街南侧禹香苑对面（新世纪中心广场一层至二层）</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                                                    <div class="widTone pull-left">
                                        <span class="pull-left insLabelForm">营业执照/社会信用代码：91140802MA0JXB7H8L</span>
                                    </div>
                                    <div class="widTone pull-left xuz pl10">
                                        <span class="pull-left insLabelForm" id="sp_s_sfjk">许可证类型：经营许可证</span>
                                    </div>                                <div class="widTone pull-left pl10">
                                    <span class="pull-left insLabelForm ar">许可证号：JY11408020030619</span>
                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="widTone pull-left">
                                    <span class="pull-left insLabelForm ar">
                                                                            年销售额：                                    /</span>
                                </div>
                                <div class="widTone pull-left pl10">
                                    单位法人：景丽霞                                </div>
                                <div class="widTone pull-left pl10">
                                    联系人：张红晶                                </div>
                            </div>
                            <div class="form-group clearfix">
                                <div class="widTone pull-left">
                                    电话：13835969913                                </div>
                                <div class="widTone pull-left pl10">
                                    传真：/                                </div>
                                <div class="widTone pull-left pl10">
                                    邮编：/                                </div>
                            </div>
                            <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        摊位号或姓名：/                                    </div>
                                    <div class="widTone pull-left">
                                        身份证号：/                                    </div>
                                </div>                                                    </div>
                    </div>
                    <!--抽样场所信息end-->

                    <!--抽样生产企业信息start-->
                    <div class="formArea1 formAreaCom">
                                                    <h6 class="areaTitle">抽样生产企业信息</h6>
                            <div class="areaContent">
                                <div class="form-group clearfix">
                                    <div class="col-xs-12 pl0">
                                        所在地：山西                                        运城                                        盐湖                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        企业地址：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        企业名称：/                                    </div>
                                    <div class="col-xs-6">
                                        生产许可证编号：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="widTone pull-left">
                                        联系人：/                                    </div>
                                    <div class="widTone pull-left pl10">
                                        电话：/                                    </div>
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
                                    <div class="col-xs-4 pl0">
                                        样品条码：/                                    </div>
                                    <div class="col-xs-4">
                                        样品商标：/                                    </div>
                                    <div class="col-xs-4">
                                        样品类型：食用农产品                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        样品来源：外购                                    </div>
                                    <div class="col-xs-3 pl0">
                                        样品属性：普通食品                                    </div>
                                    <div class="col-xs-3">
                                        包装分类：散装                                    </div>
                                    <div class="col-xs-3">
                                        样品名称：线椒                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-6 pl0">
                                        购进日期：2019-10-14                                    </div>
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
                                    <div class="col-xs-2">
                                        单价：5元/kg                                    </div>
                                    <div class="col-xs-2">
                                        是否进口：否                                    </div>
                                    <div class="col-xs-2">
                                        原产地：中国                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-3 pl0">
                                        抽样日期：2019-10-17                                    </div>
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
                                    <div class="col-xs-3 pl0">
                                        抽样基数：10kg                                    </div>
                                    <div class="col-xs-2">
                                        抽样数量：2                                    </div>
                                    <div class="col-xs-2">
                                        备样数量：1kg                                    </div>
                                    <div class="col-xs-2">
                                        抽样数量单位：kg                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-15 pl0">
                                        执行标准/技术文件：/                                    </div>
                                </div>
                                <div class="form-group clearfix">
                                    <div class="col-xs-15 pl0">
                                        备注：/                                    </div>
                                </div>
                            </div>
                        </div>                    <!--抽样样品信息end-->
                    <!--抽样生产企业信息start-->
                    <div class="formArea1 formAreaCom">
                        <h6 class="areaTitle">现场照片</h6>
                        <div class="areaContent pb30">
                            <ul class="addList clearfix">
                                <li><a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972515107535.png" data-magnify="gallery" data-group="g1" data-caption=""><img src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972515107535.png"></a></li><li><a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972644487786.png" data-magnify="gallery" data-group="g1" data-caption=""><img src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972644487786.png"></a></li><li><a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972627759191.png" data-magnify="gallery" data-group="g1" data-caption=""><img src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972627759191.png"></a></li><li><a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972439076279.png" data-magnify="gallery" data-group="g1" data-caption=""><img src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972439076279.png"></a></li><li><a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972931583230.png" data-magnify="gallery" data-group="g1" data-caption=""><img src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972931583230.png"></a></li><li><a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972741940683.png" data-magnify="gallery" data-group="g1" data-caption=""><img src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972741940683.png"></a></li><li><a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972684495200.png" data-magnify="gallery" data-group="g1" data-caption=""><img src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972684495200.png"></a></li><li><a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972839863324.png" data-magnify="gallery" data-group="g1" data-caption=""><img src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972839863324.png"></a></li>                            </ul>
                        </div>
                    </div>
                    <!--抽样生产企业信息end-->
                    <!--告知书&抽样单start-->
                    <div class="formArea1 formAreaCom attention">
                        <h6 class="areaTitle">告知书&抽样单</h6>
                        <div class="areaContent pb30 ">
                            <span class="downResult btn-sm cydUpload" id="">
                                <i>抽样单:
                                                                        <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130974089559913.png" data-magnify="gallery" data-group="g2">
                                        157130974089559913.png</a>
                                                                    </i>
                            </span>
                            <span class="downResult btn-sm gzsUpload">
                                <i>告知书:
                                                                        <a href="javascript:void(0)" data-src="http://upload1.nifdc.org.cn/image/2019/10/17/157130972816133245.png" data-magnify="gallery" data-group="g2">
                                            157130972816133245.png</a>
                                                                    </i>
                            </span>
                        </div>
                    </div>
                    <!--抽告知书&抽样单end-->
                </div>
            </div>
            <div class="ibox clearfix mt30">
                <div class="ibox-title clearfix pl20 borderline">
                    检验信息
                </div>
                <div class="formAreaCom">
                    <div class="areaContent">
                        <div class="form-group clearfix">
                            <div class="col-xs-12 pl0">
                                检验机构名称：河南恒晟检测技术有限公司                            </div>
                        </div>
                        <div class="form-group clearfix">
                            <div class="widTone pull-left">
                                收检日期：2019-10-17                            </div>
                            <div class="widTone pull-left pl10">
                                联系人：王力                            </div>
                            <div class="widTone pull-left pl10">
                                联系人电话：0371-55929768                            </div>
                        </div>
                        <div class="form-group clearfix">
                            <div class="widTone pull-left">
                                联系人邮箱：hengshengjiance23@163.com                            </div>
                            <div class="widTone pull-left pl10">
                                检查封样人员：河南恒晟检测技术有限公司                            </div>
                            <div class="widTone pull-left pl10">
                                系统接样日期：2019-10-23                            </div>
                        </div>
                        <div class="form-group clearfix">
                            <div class="widTone pull-left">
                                检查封样人电话：15738889730                            </div>
                            <div class="widTone pull-left pl10">
                                检查封样人邮箱：rdh123@163.com                            </div>
                        </div>
                        
                                            </div>
                </div>
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
</html>`


	mkr:=StoMap_yijieshou_full(olds)
	for k,v:=range mkr{
		fmt.Println("k:",k)
		fmt.Println("v:",v)
	}
}