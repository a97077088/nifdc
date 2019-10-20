package nifdc


type Channel struct{
	Name string
	Type string
}
type Data_o struct {
	Pcname_dd string
	Resource_org_name string
	Check_org_name string
	Check_user_name string
	Sample_code string
	Sample_name string
	Sp_d_38 string
	Sp_s_3 string
}
type Download_Data_r struct {
	RecordsTotal int
	Draw int
	Data []*Data_o
}
type Viewnormalsample_o struct{

}