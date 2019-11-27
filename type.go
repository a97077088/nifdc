package nifdc

import "github.com/a97077088/appuser"

type Channel struct {
	Name string
	Type string
}
type Data_o struct {
	Pcname_dd         string
	Resource_org_name string
	Check_org_name    string
	Update_time       string
	New_sample_name   string
	Check_user_name   string
	Sample_code       string
	Sample_name       string
	Sp_d_38           string
	Sp_s_3            string
	Sp_s_20           string
}
type Download_Data_r struct {
	RecordsTotal int
	Draw         int
	Data         []*Data_o
}
type Viewnormalsample_o struct {
}

type Api_food_getFood_o struct {
	Id           int
	Org_type     int
	Sample_state int
	Sp_d_38      int64
	Sp_d_46      int64
	Sp_d_86      int64
	Sp_i_state   int
	Sp_s_1       string
	Sp_s_2_1     string
	Sp_s_3       string
	Sp_s_14      string
	Sp_s_16      string
	Sp_s_35      string
	Sp_s_71      string
	Sp_s_43      string
	Sp_s_44      string
	Created_at   int64
	Updated_at   int64
	User         appuser.AppUser
}
type Api_food_getFood_r struct {
	Total int
	Rows  []*Api_food_getFood_o
}

type UploadData struct {
	appuser.AppUser
}

func (this *UploadData) Subitem() []map[string]string {
	if this.Env_for_key("subitem") == nil {
		return nil
	}
	return this.Env_for_key("subitem").([]map[string]string)
}
func (this *UploadData) AddSubitem(mp map[string]string) {
	if this.Env_for_key("subitem") == nil {
		this.Set_env_value("subitem", make([]map[string]string, 0))
	}
	subitem := this.Subitem()
	subitem = append(subitem, mp)
	this.Set_env_value("subitem", subitem)
}

type Test_platform_api_food_getTestInfo_o struct {
	Id           int
	Jylx         string
	Sp_s_44      string
	Spdata_0     string
	Spdata_0_old string
	Spdata_1     string
	Spdata_2     string
	Spdata_3     string
	Spdata_4     string
	Spdata_5     string
	Spdata_6     string
	Spdata_7     string
	Spdata_8     string
	Spdata_9     string
	Spdata_10    string
	Spdata_11    string
	Spdata_12    string
	Spdata_13    string
	Spdata_14    string
	Spdata_15    string
	Spdata_16    string
	Spdata_17    string
	Spdata_18    string
	Spdata_19    string
	Spdata_20    string
	Spdata_21    string
}
type Test_platform_api_food_getTestInfo_r struct {
	Total int
	Rows  []*Test_platform_api_food_getTestInfo_o
}

type Test_platform_api_food_save_r struct {
	Msg     string
	Success bool
}
