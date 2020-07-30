package nifdc

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/a97077088/errorv"
	. "github.com/a97077088/grequests"
	jsoniter "github.com/json-iterator/go"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var useragent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36"

//验证码获取,暂时用不到
func Code(ck string, session *Session) ([]byte, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcj.gsxt.gov.cn/code")
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"cookie": ck,
		},
	})
	if err != nil {
		return nil, err
	}
	img := r.Bytes()
	return img, nil
}

func Findval(s string, sre string) string {
	re := regexp.MustCompile(sre)
	fds := re.FindStringSubmatch(s)
	if len(fds) < 2 {
		return ""
	}
	return fds[1]
}

//登录准备
func InitLoginck(session *Session) (string, string, string, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcj.gsxt.gov.cn/login")
	r, err := cli.Get(surl, &RequestOptions{
		UserAgent:     useragent,
		RedirectLimit: -1,
	})
	if err != nil {
		return "", "", "", err
	}
	rck := r.Header.Get("Set-Cookie")
	sbd := r.String()
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(sbd))
	lt := rt.Find("input[name=lt]").AttrOr("value", "")
	execution := rt.Find("input[name=execution]").AttrOr("value", "")
	return lt, execution, rck, nil
}

//登录
func Login(username string, password string, lt string, execution string, ck string, session *Session) (string, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcj.gsxt.gov.cn/login")
	username = base64.StdEncoding.EncodeToString([]byte(username))
	password = base64.StdEncoding.EncodeToString([]byte(password))
	spd := fmt.Sprintf("username=%s&password=%s&validate=&certKey=&lt=%s&execution=%s&_eventId=submit&UserPwd=&UserSignedData=&UserCert=&ContainerName=&strRandom=", username, password, lt, execution)
	r, err := cli.Post(surl, &RequestOptions{
		RequestBody: strings.NewReader(spd),
		Headers: map[string]string{
			"Cookie":       ck,
			"Content-Type": "application/x-www-form-urlencoded",
			"Referer":      "http://spcj.gsxt.gov.cn/login",
		},
		RedirectLimit: 10,
		UserAgent:     useragent,
	})
	if err != nil {
		return "", errorv.NewNetError_e(err)
	}
	sbd := r.String()
	if strings.Index(sbd, "<i class=\"icon-user\"></i>") == -1 {
		serr := Findval(sbd, "<div id=\"msg_\" class=\"errors\">(.*?)</div>")
		if serr == "The credentials you provided cannot be determined to be authentic." {
			serr = "无法确定您提供的凭据是真实的。"
		}
		return "", errors.New(serr)
	}
	cks := r.RawResponse.Cookies()
	scks := fmt.Sprintf("%s", ck)
	for _, ck := range cks {
		scks = fmt.Sprintf("%s;%s", scks, ck.String())
	}
	if scks == "" {
		return "", errors.New("获取cookie失败")
	}
	return scks, nil
}

//首页
func Index(ck string, session *Session) error {
	cli := Cli(session)
	surl := "http://spcj.gsxt.gov.cn/ui/index"
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie":           ck,
			"x-requested-with": "XMLHttpRequest",
			"Content-Type":     "application/json;charset=UTF-8",
		},
		UserAgent: useragent,
	})
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if strings.Index(r.String(), "<a href=\"/ui/userInfor?userName=\">") != -1 {
		return errors.New("登录已过期")
	}
	return nil
}

//登录到任务大平台
func Sample_login(ck string, session *Session) (string, []*Channel, string, error) {
	cli := Cli(session)
	surl := "http://spcj.gsxt.gov.cn/login?service=http%3A%2F%2Fspcjsample.gsxt.gov.cn%2Findex.php%3Fm%3DAdmin%26c%3DSSO%26a%3Dindex"
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		UserAgent: useragent,
	})
	if err != nil {
		return "", nil, "", err
	}
	sbd := r.String()
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(sbd))
	uid := rt.Find("#user_id").AttrOr("value", "")
	chs := make([]*Channel, 0)
	rt.Find("#user_select option").Each(func(i int, selection *goquery.Selection) {
		chs = append(chs, &Channel{
			Name: selection.Text(),
			Type: selection.AttrOr("value", ""),
		})
	})
	//if uid == "" {
	//	return "", nil, errors.New("获取uid失败")
	//}
	ul, err := url.Parse("http://spcjsample.gsxt.gov.cn/index.php")
	if err != nil {
		return "", nil, "", err
	}
	cks := cli.HTTPClient.Jar.Cookies(ul)
	scks := ""
	for _, ck := range cks {
		scks = fmt.Sprintf("%s;%s", scks, ck.String())
	}
	return uid, chs, scks, nil
}

//检验检测平台
func Test_platform_login(ck string, session *Session) (string, error) {
	cli := Cli(session)
	surl := "http://spcj.gsxt.gov.cn/login?service=http%3A%2F%2Fspcjinsp.gsxt.gov.cn%2Ftest_platform%2F%3Ftoken%3D"
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		UserAgent: useragent,
	})
	if err != nil {
		return "", err
	}
	sbd := r.String()
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(sbd))
	if strings.Index(rt.Find("#btn-logout").Text(), "退出") == -1 {
		return "", errors.New("打开检验平台失败")
	}
	ul, err := url.Parse("http://spcjinsp.gsxt.gov.cn/test_platform/?token=")
	if err != nil {
		return "", err
	}
	cks := cli.HTTPClient.Jar.Cookies(ul)
	scks := ""
	for _, ck := range cks {
		scks = fmt.Sprintf("%s;%s", scks, ck.String())
	}
	return scks, nil
}

//任务大平台通道
func Sample_switchchannel(uuid string, _type string, ck string, session *Session) (string, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjsample.gsxt.gov.cn/index.php?m=Admin&c=SSO&a=logined&ca_uuid=%s&user_type=%s", uuid, _type)
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		RedirectLimit: -1,
		UserAgent:     useragent,
	})
	if err != nil {
		return "", errorv.NewNetError_e(err)
	}
	cks := r.RawResponse.Cookies()
	scks := ""
	for _, ck := range cks {
		scks = fmt.Sprintf("%s;%s", scks, ck.String())
	}
	return scks, nil
}

//已接受全字段导出
func Viewcheckedsample_full(sample_code string, ck string, session *Session) (jsoniter.Any, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjsample.gsxt.gov.cn/index.php?m=Admin&c=TaskList&a=viewcheckedsample&sample_code=%s", sample_code)
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return nil, errorv.NewNetError("http状态错误")
	}
	sbd := r.String()
	return jsoniter.Wrap(StoMap_sample(sbd)), nil
}

//数据查看,任务大平台
func DownData(resource_org_id string, sample_state int, cyTimeStart, cyTimeEnd string, smple_code string, ck string, session *Session) (*Download_Data_r, error) {
	cli := Cli(session)
	surl := "http://spcjsample.gsxt.gov.cn/index.php?m=Admin&c=TaskList&a=gettasklist"
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		Data: map[string]string{
			"sEcho":           "3",
			"iColumns":        "14",
			"sColumns":        ",,,,,,,,,,,,,",
			"iDisplayStart":   "0",
			"iDisplayLength":  "10000",
			"mDataProp_0":     "",
			"mDataProp_1":     "update_time",
			"mDataProp_2":     "sp_s_3",
			"mDataProp_3":     "sample_code",
			"mDataProp_4":     "new_sample_name",
			"mDataProp_5":     "sampling_unit_name",
			"mDataProp_6":     "check_unit_name",
			"mDataProp_7":     "create_time",
			"mDataProp_8":     "creator_id",
			"mDataProp_9":     "check_result",
			"mDataProp_10":    "scaname",
			"mDataProp_11":    "fcc_grade_one_name",
			"mDataProp_12":    "syncResult",
			"mDataProp_13":    "sample_code",
			"sample_state":    fmt.Sprintf("%d", sample_state),
			"cyTimeStart":     cyTimeStart,
			"cyTimeEnd":       cyTimeEnd,
			"sample_code":     smple_code,
			"resource_org_id": resource_org_id,
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, err
	}
	var rs Download_Data_r
	s := r.String()
	err = json.Unmarshal([]byte(s), &rs)
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}

	return &rs, nil
}

//搜索 普通食品
func Test_platform_api_food_getFood(taskfrom string, datatype int, startdate string, enddate string, offset int, limit int, sort string, order string, sampleNo string, ck string, session *Session) (*Test_platform_r, error) {
	cli := Cli(session)
	datatype = datatype
	sdatatype := ""
	if datatype != 0 {
		sdatatype = fmt.Sprintf("dataType=%d", datatype)
	} else {
		sdatatype = fmt.Sprintf("dataType=%d", 8)
	}
	if sort != "" {
		sort = fmt.Sprintf("sort=%s", sort)
	}
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/food/getFood?%s&order=%s&offset=%d&limit=%d&%s&startDate=%s&endDate=%s&taskFrom=%s&samplingUnit=&testUnit=&enterprise=&sampledUnit=&foodName=&province=&reportNo=&bsfla=&bsflb=&sampleNo=%s&foodType1=&foodType4=&_=%d", sort, order, offset, limit, sdatatype, startdate, enddate, taskfrom, sampleNo, time.Now().UnixNano())
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, err
	}
	s := r.String()
	if strings.Index(s, "安装证书助手") != -1 {
		return nil, errors.New("登录状态错误")
	}

	var rs Test_platform_r
	err = json.Unmarshal([]byte(s), &rs)
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	return &rs, nil
}

//搜索 农产品
func Test_platform_api_agriculture_getAgriculture(taskfrom string, datatype int, startdate string, enddate string, offset int, limit int, sort string, order string, sampleNo string, ck string, session *Session) (*Test_platform_r, error) {
	cli := Cli(session)
	datatype = datatype
	sdatatype := ""
	if datatype != 0 {
		sdatatype = fmt.Sprintf("dataType=%d", datatype)
	}
	if sort != "" {
		sort = fmt.Sprintf("sort=%s", sort)
	}
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/agriculture/getAgriculture?%s&order=%s&offset=%d&limit=%d&%s&startDate=%s&endDate=%s&taskFrom=%s&samplingUnit=&testUnit=&enterprise=&sampledUnit=&foodName=&province=&reportNo=&bsfla=&bsflb=&sampleNo=%s&foodType1=&foodType4=&_=%d", sort, order, offset, limit, sdatatype, startdate, enddate, taskfrom, sampleNo, time.Now().UnixNano())
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, err
	}
	s := r.String()
	if strings.Index(s, "安装证书助手") != -1 {
		return nil, errors.New("登录状态错误")
	}

	var rs Test_platform_r
	err = json.Unmarshal([]byte(s), &rs)
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	return &rs, nil
}

//农产品查看详情
func Test_platform_agricultureTest_agricultureDetail(id int, fetchitems bool, ck string, session *Session) (jsoniter.Any, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/agricultureTest/agricultureDetail/%d", id)
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie":                    ck,
			"Accept-Encoding":           "deflate",
			"Accept-Language":           "zh-CN,zh;q=0.9",
			"Referer":                   "http://spcjinsp.gsxt.gov.cn/test_platform/",
			"Upgrade-Insecure-Requests": "1",
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return nil, errorv.NewNetError("http状态错误")
	}
	sbd := r.String()
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(sbd))
	sd := rt.Find("#sd").AttrOr("value", "")
	if sd == "" {
		return nil, errorv.NewNetError_e(errors.New("获取sd失败"))
	}
	rmp := make(map[string]interface{}, 0)
	rmp["sample_code"] = rt.Find("#hid_sp_s_16").AttrOr("value", "")
	rmp["sd"] = sd
	rmp["type1"] = rt.Find("#hid_type1").AttrOr("value", "")
	rmp["type2"] = rt.Find("#hid_type2").AttrOr("value", "")
	rmp["type3"] = rt.Find("#hid_type3").AttrOr("value", "")
	rmp["type4"] = rt.Find("#hid_type4").AttrOr("value", "")
	rmp["bsfla"] = rt.Find("#hid_bsfla").AttrOr("value", "")
	rmp["bsflb"] = rt.Find("#hid_bsflb").AttrOr("value", "")
	rmp["test_unit"] = rt.Find("#test_unit").AttrOr("value", "")
	rmp["report_no"] = rt.Find("#report_no").AttrOr("value", "")
	rmp["test_date"] = rt.Find("#test_date").AttrOr("value", "")
	rmp["contact"] = rt.Find("#contact").AttrOr("value", "")
	rmp["contact_tel"] = rt.Find("#contact_tel").AttrOr("value", "")
	rmp["contact_email"] = rt.Find("#contact_email").AttrOr("value", "")
	rmp["fy_person"] = rt.Find("#fy_person").AttrOr("value", "")
	rmp["fy_tel"] = rt.Find("#fy_tel").AttrOr("value", "")
	rmp["fy_email"] = rt.Find("#fy_email").AttrOr("value", "")
	rmp["conclusion"] = rt.Find("#hid_conclusion").AttrOr("value", "")
	rmp["jd_bz"] = rt.Find("#jd_bz").Text()
	rmp["fx_bz"] = rt.Find("#fx_bz").Text()
	rmp["tb_date"] = rt.Find("#tb_date").AttrOr("value", "")
	rmp["report_type"] = rt.Find("select[name=report_type]").Find("option[selected=selected]").AttrOr("value", "")
	rmp["test_aims"] = rt.Find("select[name=test_aims]").Find("option[selected=selected]").AttrOr("value", "")
	rmp["test_conclusion"] = rt.Find("#test_conclusion").Text()
	rmp["sign_date"] = rt.Find("#sign_date").Find("option[selected=selected]").AttrOr("value", "")
	rmp["sd"] = sd

	detailmp := StoMap_test_platform(sbd)
	for k, v := range detailmp {
		rmp[k] = v
	}

	rmp["检验项目"] = make([]map[string]string, 0)
	if fetchitems == true {
		tr, err := Test_platform_api_food_getTestInfo(sd, ck, cli)
		if err != nil {
			return nil, err
		}
		for _, it := range tr.Rows {
			jyitem := make(map[string]string)
			jyitem["检验项目"] = it.Spdata_0
			jyitem["检验结果"] = it.Spdata_1
			jyitem["结果单位"] = it.Spdata_18
			jyitem["结果判定"] = it.Spdata_2
			jyitem["检验依据"] = it.Spdata_3
			jyitem["判定依据"] = it.Spdata_4
			jyitem["最小允许限"] = it.Spdata_11
			jyitem["最大允许限"] = it.Spdata_15
			jyitem["允许限单位"] = it.Spdata_16
			jyitem["方法检出限"] = it.Spdata_7
			jyitem["检出限单位"] = it.Spdata_8
			jyitem["备注"] = it.Spdata_20
			jyitem["说明"] = it.Spdata_17
			rmp["检验项目"] = append(rmp["检验项目"].([]map[string]string), jyitem)
		}

	}

	return jsoniter.Wrap(rmp), nil
}

//普通食品查看详情
func Test_platform_foodTest_foodDetail(id int, fetchitems bool, ck string, session *Session) (jsoniter.Any, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/foodTest/foodDetail/%d", id)
	r, err := cli.Get(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie":                    ck,
			"Accept-Encoding":           "deflate",
			"Accept-Language":           "zh-CN,zh;q=0.9",
			"Referer":                   "http://spcjinsp.gsxt.gov.cn/test_platform/",
			"Upgrade-Insecure-Requests": "1",
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return nil, errorv.NewNetError("http状态错误")
	}
	sbd := r.String()
	rt, _ := goquery.NewDocumentFromReader(strings.NewReader(sbd))
	sd := rt.Find("#sd").AttrOr("value", "")
	if sd == "" {
		return nil, errorv.NewNetError_e(errors.New("获取sd失败"))
	}
	rmp := make(map[string]interface{}, 0)
	rmp["sample_code"] = rt.Find("#hid_sp_s_16").AttrOr("value", "")
	rmp["sd"] = sd
	rmp["type1"] = rt.Find("#hid_type1").AttrOr("value", "")
	rmp["type2"] = rt.Find("#hid_type2").AttrOr("value", "")
	rmp["type3"] = rt.Find("#hid_type3").AttrOr("value", "")
	rmp["type4"] = rt.Find("#hid_type4").AttrOr("value", "")
	rmp["bsfla"] = rt.Find("#hid_bsfla").AttrOr("value", "")
	rmp["bsflb"] = rt.Find("#hid_bsflb").AttrOr("value", "")
	rmp["test_unit"] = rt.Find("#test_unit").AttrOr("value", "")
	rmp["report_no"] = rt.Find("#report_no").AttrOr("value", "")
	rmp["test_date"] = rt.Find("#test_date").AttrOr("value", "")
	rmp["contact"] = rt.Find("#contact").AttrOr("value", "")
	rmp["contact_tel"] = rt.Find("#contact_tel").AttrOr("value", "")
	rmp["contact_email"] = rt.Find("#contact_email").AttrOr("value", "")
	rmp["fy_person"] = rt.Find("#fy_person").AttrOr("value", "")
	rmp["fy_tel"] = rt.Find("#fy_tel").AttrOr("value", "")
	rmp["fy_email"] = rt.Find("#fy_email").AttrOr("value", "")
	rmp["conclusion"] = rt.Find("#hid_conclusion").AttrOr("value", "")
	rmp["jd_bz"] = rt.Find("#jd_bz").Text()
	rmp["fx_bz"] = rt.Find("#fx_bz").Text()
	rmp["tb_date"] = rt.Find("#tb_date").AttrOr("value", "")
	rmp["report_type"] = rt.Find("select[name=report_type]").Find("option[selected=selected]").AttrOr("value", "")
	rmp["test_aims"] = rt.Find("select[name=test_aims]").Find("option[selected=selected]").AttrOr("value", "")
	rmp["test_conclusion"] = rt.Find("#test_conclusion").Text()
	rmp["sign_date"] = rt.Find("#sign_date").Find("option[selected=selected]").AttrOr("value", "")
	rmp["sd"] = sd
	detailmp := StoMap_test_platform(sbd)
	for k, v := range detailmp {
		rmp[k] = v
	}
	rmp["检验项目"] = make([]map[string]string, 0)
	if fetchitems == true {
		tr, err := Test_platform_api_food_getTestInfo(sd, ck, cli)
		if err != nil {
			return nil, err
		}
		for _, it := range tr.Rows {
			jyitem := make(map[string]string)
			jyitem["检验项目"] = it.Spdata_0
			jyitem["检验结果"] = it.Spdata_1
			jyitem["结果单位"] = it.Spdata_18
			jyitem["结果判定"] = it.Spdata_2
			jyitem["检验依据"] = it.Spdata_3
			jyitem["判定依据"] = it.Spdata_4
			jyitem["最小允许限"] = it.Spdata_11
			jyitem["最大允许限"] = it.Spdata_15
			jyitem["允许限单位"] = it.Spdata_16
			jyitem["方法检出限"] = it.Spdata_7
			jyitem["检出限单位"] = it.Spdata_8
			jyitem["备注"] = it.Spdata_20
			jyitem["说明"] = it.Spdata_17
			rmp["检验项目"] = append(rmp["检验项目"].([]map[string]string), jyitem)
		}

	}
	return jsoniter.Wrap(rmp), nil
}

//获取agriculture_getTestItems
func Test_platform_api_agriculture_getTestItems(fddetail jsoniter.Any, ck string, session *Session) (*Test_platform_api_food_getTestItems_r, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/agriculture/getTestItems")
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		Data: map[string]string{
			"type1": fddetail.Get("抽样基础信息_食品大类").ToString(),
			"type2": fddetail.Get("抽样基础信息_食品亚类").ToString(),
			"type3": fddetail.Get("抽样基础信息_食品次亚类").ToString(),
			"type4": fddetail.Get("抽样基础信息_食品细类").ToString(),
			"bsflA": fddetail.Get("抽样基础信息_报送分类A").ToString(),
			"bsflB": fddetail.Get("抽样基础信息_报送分类B").ToString(),
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return nil, errorv.NewNetError("http状态错误")
	}
	var rs Test_platform_api_food_getTestItems_r
	err = r.JSON(&rs)
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	//if len(rs.Rows) == 0 {
	//	return nil, errors.New("至少需要一个检验项目")
	//}
	return &rs, nil
}

//获取agriculture_testinfo
func Test_platform_api_agriculture_getTestInfo(sd string, ck string, session *Session) (*Test_platform_api_food_getTestInfo_r, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/agriculture/getTestInfo")
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		Data: map[string]string{
			"sd": sd,
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return nil, errorv.NewNetError("http状态错误")
	}
	var rs Test_platform_api_food_getTestInfo_r
	err = r.JSON(&rs)
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	//if len(rs.Rows) == 0 {
	//	return nil, errors.New("至少需要一个检验项目")
	//}
	return &rs, nil
}

//保存testinfo
func Test_platform_api_food_init(fooddetail map[string]string, testinfos []*Test_platform_api_food_getTestItems_o, ck string, session *Session) error {
	items := make([]map[string]string, 0)
	for _, tinfo := range testinfos {
		itmap := make(map[string]string)
		itmap["id"] = fmt.Sprintf("")
		itmap["item_old"] = tinfo.Item
		itmap["item"] = tinfo.Item
		itmap["sp_data_1"] = "" //结果
		itmap["sp_data_2"] = "未检验"
		itmap["sp_data_3"] = tinfo.TestReason[0].Sm
		itmap["sp_data_4"] = tinfo.VerifyReason[0].Spdata_4
		//itmap["sp_data_5"] = tinfo.Spdata_5
		//itmap["sp_data_6"] = tinfo.Spdata_6
		//itmap["sp_data_7"] = tinfo.Spdata_7
		//itmap["sp_data_8"] = tinfo.Spdata_8
		//itmap["sp_data_9"] = tinfo.Spdata_9
		//itmap["sp_data_10"] = tinfo.Spdata_10
		//itmap["sp_data_11"] = tinfo.Spdata_11
		//itmap["sp_data_12"] = tinfo.Spdata_12
		//itmap["sp_data_13"] = tinfo.Spdata_13
		//itmap["sp_data_15"] = tinfo.Spdata_15
		//itmap["sp_data_16"] = tinfo.Spdata_16
		//itmap["sp_data_17"] = tinfo.Spdata_17
		//itmap["sp_data_18"] = tinfo.Spdata_18
		//itmap["bz"] = tinfo.Spdata_20
		//itmap["sm"] = tinfo.Spdata_19
		//itmap["sp_data_21"] = tinfo.Spdata_21
		//itmap["jylx"] = tinfo.Jylx
		items = append(items, itmap)
	}

	sitems, err := json.Marshal(items)
	if err != nil {
		return err
	}
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/food/save")
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie":       ck,
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		},
		Data: map[string]string{
			"type1":           fooddetail["type1"],
			"type2":           fooddetail["type2"],
			"type3":           fooddetail["type3"],
			"type4":           fooddetail["type4"],
			"bsfla":           fooddetail["bsfla"],
			"bsflb":           fooddetail["bsflb"],
			"test_unit":       fooddetail["test_unit"],
			"report_no":       fooddetail["report_no"],
			"test_date":       fooddetail["test_date"],
			"contact":         fooddetail["contact"],
			"contact_tel":     fooddetail["contact_tel"],
			"contact_email":   fooddetail["contact_email"],
			"fy_person":       fooddetail["fy_person"],
			"fy_tel":          fooddetail["fy_tel"],
			"fy_email":        fooddetail["fy_email"],
			"conclusion":      fooddetail["conclusion"],
			"jd_bz":           fooddetail["jd_bz"],
			"fx_bz":           fooddetail["fx_bz"],
			"tb_date":         fooddetail["tb_date"],
			"report_type":     fooddetail["report_type"],
			"test_aims":       fooddetail["test_aims"],
			"test_conclusion": fooddetail["test_conclusion"],
			"sign_date":       fooddetail["sign_date"],
			"sd":              fooddetail["sd"],
			"fdtoken12":       "201812FoodDetail",
			"fdtoken1201":     "20181201FoodDetail",
			"isSubmit":        "false",
			"items":           string(sitems),
		},
		UserAgent: useragent,
	})
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return errorv.NewNetError("http状态错误")
	}
	var rs Test_platform_api_food_save_r
	err = r.JSON(&rs)
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if rs.Success != true {
		return errors.New(rs.Msg)
	}
	return nil
}

//保存agriculture_init
func Test_platform_api_agriculture_init(fooddetail map[string]string, testinfos []*Test_platform_api_food_getTestItems_o, ck string, session *Session) error {
	items := make([]map[string]string, 0)
	for _, tinfo := range testinfos {
		itmap := make(map[string]string)
		itmap["id"] = fmt.Sprintf("")
		itmap["item_old"] = tinfo.Item
		itmap["item"] = tinfo.Item
		itmap["sp_data_1"] = "" //结果
		itmap["sp_data_2"] = "未检验"
		itmap["sp_data_3"] = tinfo.TestReason[0].Sm
		itmap["sp_data_4"] = tinfo.VerifyReason[0].Spdata_4
		//itmap["sp_data_5"] = tinfo.Spdata_5
		//itmap["sp_data_6"] = tinfo.Spdata_6
		//itmap["sp_data_7"] = tinfo.Spdata_7
		//itmap["sp_data_8"] = tinfo.Spdata_8
		//itmap["sp_data_9"] = tinfo.Spdata_9
		//itmap["sp_data_10"] = tinfo.Spdata_10
		//itmap["sp_data_11"] = tinfo.Spdata_11
		//itmap["sp_data_12"] = tinfo.Spdata_12
		//itmap["sp_data_13"] = tinfo.Spdata_13
		//itmap["sp_data_15"] = tinfo.Spdata_15
		//itmap["sp_data_16"] = tinfo.Spdata_16
		//itmap["sp_data_17"] = tinfo.Spdata_17
		//itmap["sp_data_18"] = tinfo.Spdata_18
		//itmap["bz"] = tinfo.Spdata_20
		//itmap["sm"] = tinfo.Spdata_19
		//itmap["sp_data_21"] = tinfo.Spdata_21
		//itmap["jylx"] = tinfo.Jylx
		items = append(items, itmap)
	}

	sitems, err := json.Marshal(items)
	if err != nil {
		return err
	}
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/agriculture/save")
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie":       ck,
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		},
		Data: map[string]string{
			"type1":           fooddetail["type1"],
			"type2":           fooddetail["type2"],
			"type3":           fooddetail["type3"],
			"type4":           fooddetail["type4"],
			"bsfla":           fooddetail["bsfla"],
			"bsflb":           fooddetail["bsflb"],
			"test_unit":       fooddetail["test_unit"],
			"report_no":       fooddetail["report_no"],
			"test_date":       fooddetail["test_date"],
			"contact":         fooddetail["contact"],
			"contact_tel":     fooddetail["contact_tel"],
			"contact_email":   fooddetail["contact_email"],
			"fy_person":       fooddetail["fy_person"],
			"fy_tel":          fooddetail["fy_tel"],
			"fy_email":        fooddetail["fy_email"],
			"conclusion":      fooddetail["conclusion"],
			"jd_bz":           fooddetail["jd_bz"],
			"fx_bz":           fooddetail["fx_bz"],
			"tb_date":         fooddetail["tb_date"],
			"report_type":     fooddetail["report_type"],
			"test_aims":       fooddetail["test_aims"],
			"test_conclusion": fooddetail["test_conclusion"],
			"sign_date":       fooddetail["sign_date"],
			"sd":              fooddetail["sd"],
			"fdtoken12":       "201812FoodDetail",
			"fdtoken1201":     "20181201FoodDetail",
			"isSubmit":        "false",
			"items":           string(sitems),
		},
		UserAgent: useragent,
	})
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return errorv.NewNetError("http状态错误")
	}
	var rs Test_platform_api_food_save_r
	err = r.JSON(&rs)
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if rs.Success != true {
		return errors.New(rs.Msg)
	}
	return nil
}

//保存agriculture_testinfo
func Test_platform_api_agriculture_save(fooddetail jsoniter.Any, updatas []map[string]string, ck string, session *Session) error {
	items := updatas
	sitems, err := json.Marshal(items)
	if err != nil {
		return err
	}
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/agriculture/save")
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie":       ck,
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		},
		Data: map[string]string{
			"type1":           fooddetail.Get("type1").ToString(),
			"type2":           fooddetail.Get("type2").ToString(),
			"type3":           fooddetail.Get("type3").ToString(),
			"type4":           fooddetail.Get("type4").ToString(),
			"bsfla":           fooddetail.Get("bsfla").ToString(),
			"bsflb":           fooddetail.Get("bsflb").ToString(),
			"test_unit":       fooddetail.Get("test_unit").ToString(),
			"report_no":       fooddetail.Get("report_no").ToString(),
			"test_date":       fooddetail.Get("test_date").ToString(),
			"contact":         fooddetail.Get("contact").ToString(),
			"contact_tel":     fooddetail.Get("contact_tel").ToString(),
			"contact_email":   fooddetail.Get("contact_email").ToString(),
			"fy_person":       fooddetail.Get("fy_person").ToString(),
			"fy_tel":          fooddetail.Get("fy_tel").ToString(),
			"fy_email":        fooddetail.Get("fy_email").ToString(),
			"conclusion":      fooddetail.Get("conclusion").ToString(),
			"jd_bz":           fooddetail.Get("jd_bz").ToString(),
			"fx_bz":           fooddetail.Get("fx_bz").ToString(),
			"tb_date":         fooddetail.Get("tb_date").ToString(),
			"report_type":     fooddetail.Get("report_type").ToString(),
			"test_aims":       fooddetail.Get("test_aims").ToString(),
			"test_conclusion": fooddetail.Get("test_conclusion").ToString(),
			"sign_date":       fooddetail.Get("sign_date").ToString(),
			"sd":              fooddetail.Get("sd").ToString(),
			"fdtoken12":       "201812FoodDetail",
			"fdtoken1201":     "20181201FoodDetail",
			"isSubmit":        "false",
			"items":           string(sitems),
		},
		UserAgent: useragent,
	})
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return errorv.NewNetError("http状态错误")
	}
	var rs Test_platform_api_food_save_r
	err = r.JSON(&rs)
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if rs.Success != true {
		return errors.New(rs.Msg)
	}
	return nil
}

//获取food_getTestItems
func Test_platform_api_food_getTestItems(fddetail map[string]string, ck string, session *Session) (*Test_platform_api_food_getTestItems_r, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/food/getTestItems")
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		Data: map[string]string{
			"type1": fddetail["抽样基础信息_食品大类"],
			"type2": fddetail["抽样基础信息_食品亚类"],
			"type3": fddetail["抽样基础信息_食品次亚类"],
			"type4": fddetail["抽样基础信息_食品细类"],
			"bsflA": fddetail["抽样基础信息_报送分类A"],
			"bsflB": fddetail["抽样基础信息_报送分类B"],
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return nil, errorv.NewNetError("http状态错误")
	}
	var rs Test_platform_api_food_getTestItems_r
	err = r.JSON(&rs)
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	//if len(rs.Rows) == 0 {
	//	return nil, errors.New("至少需要一个检验项目")
	//}
	return &rs, nil
}

//获取testinfo
func Test_platform_api_food_getTestInfo(sd string, ck string, session *Session) (*Test_platform_api_food_getTestInfo_r, error) {
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/food/getTestInfo")
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie": ck,
		},
		Data: map[string]string{
			"sd": sd,
		},
		UserAgent: useragent,
	})
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return nil, errorv.NewNetError("http状态错误")
	}
	var rs Test_platform_api_food_getTestInfo_r
	err = r.JSON(&rs)
	if err != nil {
		return nil, errorv.NewNetError_e(err)
	}
	//if len(rs.Rows) == 0 {
	//	return nil, errors.New("至少需要一个检验项目")
	//}
	return &rs, nil
}

//保存testinfo
func Test_platform_api_food_save(fooddetail map[string]string, updatas []map[string]string, ck string, session *Session) error {
	items := updatas

	sitems, err := json.Marshal(items)
	if err != nil {
		return err
	}
	cli := Cli(session)
	surl := fmt.Sprintf("http://spcjinsp.gsxt.gov.cn/test_platform/api/food/save")
	r, err := cli.Post(surl, &RequestOptions{
		Headers: map[string]string{
			"Cookie":       ck,
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		},
		Data: map[string]string{
			"type1":           fooddetail["type1"],
			"type2":           fooddetail["type2"],
			"type3":           fooddetail["type3"],
			"type4":           fooddetail["type4"],
			"bsfla":           fooddetail["bsfla"],
			"bsflb":           fooddetail["bsflb"],
			"test_unit":       fooddetail["test_unit"],
			"report_no":       fooddetail["report_no"],
			"test_date":       fooddetail["test_date"],
			"contact":         fooddetail["contact"],
			"contact_tel":     fooddetail["contact_tel"],
			"contact_email":   fooddetail["contact_email"],
			"fy_person":       fooddetail["fy_person"],
			"fy_tel":          fooddetail["fy_tel"],
			"fy_email":        fooddetail["fy_email"],
			"conclusion":      fooddetail["conclusion"],
			"jd_bz":           fooddetail["jd_bz"],
			"fx_bz":           fooddetail["fx_bz"],
			"tb_date":         fooddetail["tb_date"],
			"report_type":     fooddetail["report_type"],
			"test_aims":       fooddetail["test_aims"],
			"test_conclusion": fooddetail["test_conclusion"],
			"sign_date":       fooddetail["sign_date"],
			"sd":              fooddetail["sd"],
			"fdtoken12":       "201812FoodDetail",
			"fdtoken1201":     "20181201FoodDetail",
			"isSubmit":        "false",
			"items":           string(sitems),
		},
		UserAgent: useragent,
	})
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if r.StatusCode != 200 {
		return errorv.NewNetError("http状态错误")
	}
	var rs Test_platform_api_food_save_r
	err = r.JSON(&rs)
	if err != nil {
		return errorv.NewNetError_e(err)
	}
	if rs.Success != true {
		return errors.New(rs.Msg)
	}
	return nil
}
