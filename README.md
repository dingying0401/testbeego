# 线上交易系统
### 框架信息
- 项目搭建框架：beego
- beego docs:
``
https://beego.me/docs/intro/
``

### 系统信息

- 登陆访问：www.xiaobizai.com:33333/login
- 本地配置域名：hosts: 127.0.0.1 www.xiaobizai.com
- xorm安装：
``
go get github.com/xormplus/xorm
``
- 配置xorm示例：
```go
xorm.NewMySQL("mysql", "root:password@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai")
//用户名：密码@tcp(mysql服务器：端口/库名)
```
- 本地mysql导入sql文件：
``
/test.sql
``
###系统简介
本系统是一个模拟在线交易系统。主要角色有购买者（user）和商品管理者(admin)两种。
- 购买者和商品管理者共享一套用户登陆体系
- 购买者在注册账号后，可登陆交易页面
- 登陆验证采用jwt验证
- 登陆后可以跳转到查看所有商品的列表的页面
- 可单独选择某个商品进行购买（当前不支持购物车功能）
- 购买后会生成一笔订单
- 用户可查看自己的所有订单以及某条历史订单，也可以对订单进行删除操作
- 用户也可进入自己的信息管理页面，更新或者修改自己的个人账号信息
- 用户可购买vip会员
- 管理员可更新商品信息
- 管理员可以删除某个商品信息


#### 登陆模块

#### 主页

#### 订单

##### 用户信息


