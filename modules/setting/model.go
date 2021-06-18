package setting

import "ping/com/mysql"

type AccessCredentials struct {
	Id             int64  `sql:"id" form:"id"`                           //ID 自动生成
	Name           string `sql:"name" form:"name"`                       //给访问凭证取的名字
	UserName       string `sql:"username" form:"username"`               // 用于ssh的用户名
	Password       string `sql:"password" form:"password"`               // 用于ssh的密码
	BecomeMethod   string `sql:"become_method" form:"become_method"`     //当需要特权时是su  或者 enable
	BecomeUser     string `sql:"become_user" form:"become_user"`         ////当需要特权时 使用的用户名
	BecomePassword string `sql:"become_password" form:"become_password"` ////当需要特权时 使用的密码
	PublicKey      string `sql:"public_key" form:"public_key"`           //公钥认证
}

func (a *AccessCredentials) Clone() mysql.Object {
	return &AccessCredentials{}
}
