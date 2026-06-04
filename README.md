# 短连接项目

## 搭建项目的骨架
1.建库建表
新建发号器
对应 sequence.sql

新建长链接短连接映射表
对应short_url.sql

2.搭建go-zero框架的骨架
编写api文件  使用goctl 命令生成api文件
写好文件后，根据api文件生成代码
goctl api go -api .\shortener.api -dir .


3.根据数据表生成model 层代码
goctl model mysql datasource -url="root:jinyubo@tcp(127.0.0.1:3306)/test" -table="short_url_map" -dir="./model"
goctl model mysql datasource -url="root:jinyubo@tcp(127.0.0.1:3306)/test" -table="sequence" -dir="./model"     


4 go mod tidy

5 运行
看是否可以运行

6 修改配置文件

## 查看短链接
### 缓存版
有两种方式
1。使用自己实现的缓存     surl->lurl 节省缓存数据量
2. 使用go-zero自带的缓存  surl->数据行  不需要自己实现，开发量小
这里使用第二种
1.添加缓存配置  -配置文件   -配置config结构体
2.删除旧的model层 -删除shorturlmodel文件
3. 重新生成model层代码 goctl model mysql datasource -url="root:jinyubo@tcp(127.0.0.1:3306)/test" -table="short_url_map" -dir="./model" -c