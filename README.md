# 线上交易系统
### 框架信息
- 项目搭建框架：beego
- beego docs:
``
https://beego.me/docs/intro/
``

### 系统信息
- 登陆访问：www.xiaobizai.com：33333/login
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
#### 登陆模块
登陆/注册逻辑

#### 主页
商品：增删改查？

#### 订单
增删改查订单？

##### 用户信息
用户信息管理？
 

