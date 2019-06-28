package models

import "fmt"

/*购买vip（更新会员信息）*/
func FuncVipShop(uid int) error {
	x := getDBEngine()
	vip := LoginUser{Ifvip: true}
	affected, err := x.Where("Uid=?", uid).Update(vip)
	if err == nil {
		fmt.Println(affected)
		fmt.Println("购买成功")
	} else {
		fmt.Println("购买失败")
	}
	return err
}

/*查看是否是vip*/
func Checkvip(uid int) (error, bool) {
	x := getDBEngine()
	vip := LoginUser{}
	affected, err := x.Where("Uid=?", uid).Get(&vip)
	if err == nil {
		fmt.Println(affected)
		fmt.Println("可以查到")
	} else {
		fmt.Println("查不到")
	}
	return err, vip.Ifvip
}

/*更新用户信息*/
func EditUserinfo(uid int, email string, address string, phone string) *LoginUser {
	x := getDBEngine()
	userinfo := new(LoginUser)
	has, err := x.Where("Uid=?", uid).Get(userinfo)
	if has && err == nil {
		fmt.Println(userinfo.Username)
		userinfo.Email = email
		userinfo.Address = address
		userinfo.Phone = phone
		affected, err := x.Where("Uid=?", uid).Update(userinfo)
		if err == nil {
			fmt.Print(affected)
			fmt.Println("更新成功")
		}
	}
	num, result := x.Where("Uid=?", uid).Get(userinfo)
	if result == nil {
		fmt.Println(num)
	}
	return userinfo
}
