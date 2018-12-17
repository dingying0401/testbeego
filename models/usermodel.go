package models

import "fmt"

func FuncVipShop(uid int) error{
	x:= getDBEngine()
	vip := LoginUser{Ifvip:"yes"}
	affected,err := x.Where("Uid=?", uid).Update(vip)
	if err == nil {
		fmt.Println(affected)
		fmt.Println("购买成功")
	}else{
		fmt.Println("购买失败")
	}
	return err
}


func Checkvip(uid int) (error,string){
	x:= getDBEngine()
	vip := LoginUser{}
	affected,err := x.Where("Uid=?", uid).Get(&vip)
	if err ==nil {
		fmt.Println(affected)
		fmt.Println("可以查到")
	}else{
		fmt.Println("查不到")
	}
	return err,vip.Ifvip
}

func EditUserinfo(uid int,email string,address string, phone string) *LoginUser {
		x := getDBEngine()
		userinfo := new(LoginUser)
		has, err := x.Where("Uid=?", uid).Get(userinfo)
		if (has && err == nil) {
		fmt.Println(userinfo.Username)
		userinfo.Email = email
		userinfo.Address = address
		userinfo.Phone = phone
		affected,err := x.Where("Uid=?", uid).Update(userinfo)
		if err == nil{
		fmt.Print(affected)
		fmt.Println("更新成功")
	}
	}
		num,result := x.Where("Uid=?", uid).Get(userinfo)
		if result == nil{
		fmt.Println(num)
	}
		return userinfo
	}
