package models

import (
"fmt"
_ "github.com/go-sql-driver/mysql"
)
func BoyInfo() string {
	x := getDBEngine()
	var productlist []ProductInfo
	results,err := x.Search(&productlist).Json()
	//affected := x.Cols().Find(&productlist)
	if err ==nil{
		fmt.Println("有商品列表")
	}else {
		fmt.Println("无商品列表")
	}
	return results
}
