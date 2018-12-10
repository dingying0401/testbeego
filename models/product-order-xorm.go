package models

import (
"fmt"
_ "github.com/go-sql-driver/mysql"
)
type ProductInfo struct{
	Pid int `xorm:"pk"`//`orm:"column(pid);pk"` // 设置主键*/
	Name string
	Length float64
	Weight float64
	Briefinfo string
	Singleprice float64
}
func BoyInfo() map[int64]ProductInfo {
	x := getDBEngine()

	//sql_2_1 := "select * from product_login"
	//results, err := x.QueryString(sql_2_1)
	//var productlist ProductInfo
	users := make(map[int64]ProductInfo)
	//results,err := x.Where("id=?", 0).Search(&productlist).Json()
	//sql := "select * from product_login"
	//results,err := x.SQL(sql).Query().Json()
	err := x.Find(&users)
	if err ==nil{
		fmt.Println("有商品列表")
	}else {
		fmt.Println("无商品列表")
	}
	return users
}
