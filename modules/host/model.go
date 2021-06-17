package host

import "ping/com/mysql"

type Device struct {
	Id         int64  `form:"id" sql:"id" json:"id"`
	Hostname   string `form:"hostname" sql:"hostname" json:"hostname" xlsx:"0"`
	Address    string `form:"hostip" sql:"hostip" json:"hostip" xlsx:"2"`
	Port       int32  `form:"port" sql:"port" json:"port" xlsx:"3"`
	Group      string `form:"hostgroup" sql:"group" json:"hostgroup" xlsx:"1" `
	Comment    string `form:"description" sql:"description" json:"description" xlsx:"5"`
	UserName   string `form:"username" sql:"username" xlsx:"4"`
	Password   string `form:"password" sq:"password" xlsx:"5"`
	StatusCode int32  `json:"statuscode"  sq:"statuscode"`
}

type Line struct {
	Id              int64  `sql:"id"`
	Operators       string `sql:"operators"`         //运营商，如电信、联通
	AEndAccess      string `sql:"a_end_access"`      //A端接入设备
	AEndInterface   string `sql:"a_end_interface"`   //A端设备接口
	AEndIpMask      string `sql:"a_end_ip_mask"`     // 	A端互联IP及掩码
	StandbyMode     string `sql:"standby_mode"`      // 	主备冗余
	LineType        string `sql:"line_type"`         // 	线路类别
	BandWidth       int    `sql:"band_width"`        // 	带宽
	BEndDescription string `sql:"b_end_description"` // 	B端简述
	LocalOrOther    string `sql:"local_or_other"`    // 	线路性质，本地或省或其它
	InerOrOuter     string `sql:"inner_or_outer"`    // 	内线外线
	LineNumber      string `sql:"line_number"`       // 	电路编号
	VlanId          string `sql:"vlan_id"`           // 	vlan id
	AEndCreateBy    string `sql:"a_end_create_by"`   // A端电路发起人
	BEndContact     string `sql:"b_end_contact"`     // 	B端电路联系人
	Remarks         string `sql:"remarks"`           // 	备注

}

type Switch Device
type Firewall Device
type Router Device
type Linux Device

// Host_Line 主机与线路关系表
type Host_Line struct {
	id      int64 `sql:"id"`
	Host_id int64 `sql:"host_id"`
	Line_id int64 `sql:"line_id"`
}

func (h *Device) Clone() (obj mysql.Object) {
	return &Device{}
}

func (h *Switch) Clone() (obj mysql.Object) {
	return &Switch{}
}

func (h *Firewall) Clone() (obj mysql.Object) {
	return &Firewall{}
}

func (h *Linux) Clone() (obj mysql.Object) {
	return &Linux{}
}

func (h *Router) Clone() (obj mysql.Object) {
	return &Router{}
}

func (h *Line) Clone() (obj mysql.Object) {
	return &Line{}
}

func (h *Host_Line) Clone() (obj mysql.Object) {
	return &Host_Line{}
}
