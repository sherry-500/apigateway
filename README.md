# cloudwego Apigateway

- [cloudwego Apigateway](#cloudwego-apigateway)
  - [一、部署步骤](#一部署步骤)
    - [1. 前置条件](#1-前置条件)
    - [2. 开始部署](#2-开始部署)
      - [2.1 下载源码](#21-下载源码)
      - [2.2 数据库部署](#22-数据库部署)
      - [2.3  启动etcd](#23--启动etcd)
      - [2.4 启动HTTP服务器](#24-启动http服务器)
      - [2.5 启动IDL管理平台](#25-启动idl管理平台)
      - [2.6 启动RPC服务](#26-启动rpc服务)
  - [二、接口说明](#二接口说明)
    - [api网关（http-server）](#api网关http-server)
      - [1. 服务调用](#1-服务调用)
    - [idl管理平台（idl-manage）](#idl管理平台idl-manage)
      - [1. 增加idl](#1-增加idl)
      - [2. 删除idl](#2-删除idl)
      - [3. 更新idl](#3-更新idl)
      - [4. 根据服务名称获取idl](#4-根据服务名称获取idl)
      - [5. 列出idl](#5-列出idl)
  - [三、使用说明](#三使用说明)
  - [四、实现细节](#四实现细节)
  - [五、测试方案](#五测试方案)
    - [1. 单元测试](#1-单元测试)
    - [2. 性能测试](#2-性能测试)
  - [六、优化方案](#六优化方案)


## 一、部署步骤

### 1. 前置条件
1. [Golang开发环境](https://golang.google.cn/)
2. [etcd](https://etcd.io/)
3. [PostgreSQL](https://www.postgresql.org/)

### 2. 开始部署

#### 2.1 下载源码

1. 克隆源码到本地

#### 2.2 数据库部署

1. 本地新建一个空的数据库（如果本地安装了migrate，可以修改Makefile中的migratedown和migrateup，并执行）
2. 对该数据库执行啊apigateway/idl-manage/backend/db/idlmap.sql中的sql语句

3. 按需修改`apigateway/http-server/app.env`与`apigateway/idl-manage/backend/app.env`两个文件的内容

```shell
DB_DRIVER=postgres
DB_SOURCE=postgresql://[your-postgres-user-name]:[your-postgresql-password]@localhost:5432/idlmap?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8891
```

#### 2.3  启动etcd

1. 新建一个终端标签页，输入`etcd --log-level debug`

```shell
username@computername ~ % etcd --log-level debug
```

2. 验证是否正确启动，新建一个终端标签页

```shell
username@computername ~ % etcdctl endpoint health

127.0.0.1:2379 is healthy: successfully committed proposal: took = 5.028375ms
```

#### 2.4 启动HTTP服务器

1. 新建一个终端标签页，进入文件夹并编译

```shell
username@computername apigateway % cd http-server
username@computername http-server % go build
```

2. 启动服务器

```shell
username@computername http-server % ./apigateway 
```

#### 2.5 启动IDL管理平台

1. 新建一个终端标签页，进入文件夹并编译

```shell
username@computername http-server % cd ../idl-manage/backend

username@computername backend % go build
```

2. 启动服务器

```shell
pucci@M45kongzhitai backend % ./idlManage 
```

#### 2.6 启动RPC服务

1. 对每一个服务，新建一个终端标签页，进入文件夹、编译、启动服务器

```shell
username@computername backend % cd ../../rpc-server 

username@computername rpc-server % cd student 

username@computername student % sh build.sh 

username@computername student % sh output/bootstrap.sh 
```

## 二、接口说明

### api网关（http-server）

#### 1. 服务调用

请求路径：/apigateway/:svcName/:methodName

请求方法：POST

请求参数：取决希望调用的rpc服务的请求参数

请求示例：

```shell
 curl -X POST 
 -H "Content-Type:application/json" 
 -d '{"name":"sherry", "id":1}' 
 YOUR_SERVER_ADDRESS/apigateway/studentservice/Register
```

返回参数：与调用的rpc服务的返回参数一致

### idl管理平台（idl-manage）

#### 1. 增加idl

请求路径：/idl-manage

请求方法：POST

请求参数：

| 字段    | 说明                       | 类型   | 备注                                                         | 是否必填 |
| ------- | -------------------------- | ------ | ------------------------------------------------------------ | -------- |
| svcName | rpc服务在etcd注册的服务名  | string |                                                              | 是       |
| idl     | rpc服务的idl文件路径或内容 | string | 取决于type, 如果type == "PATH", idl的取值的idl文件相对于http-server/main.go的相对路径 | 是       |
| type    | idl的类型                  | string | PATH或CONTENT                                                | 是       |

请求示例：

```shell
curl -X POST 
-H "Content-Type:application/json"
-d '{"svcName":"teacherservice", "idl":"../idl/teacher.thrift", "type":"PATH"}'
YOUR_SERVER_ADDRESS/idl-manage
```

返回参数：

与请求参数相同

#### 2. 删除idl

请求路径：/idl-manage/:svcname

请求方法：DELETE

请求参数：

请求示例：

```shell
curl -X DELETE
-H "Content-Type:application/json"
YOUR_SERVER_ADDRESS/idl-manage/teacherservice
```

返回参数：

| 字段    | 说明         | 类型   | 备注 |
| ------- | ------------ | ------ | ---- |
| success | 删除是否成功 | bool   |      |
| message | 返回消息     | string |      |

#### 3. 更新idl

请求路径：/idl-manage/:svcname

请求方法：PATCH

请求参数：与增加idl的请求参数相同

请求示例：

```shell
 curl -X PATCH 
 -H "Content-Type:application/json" 
 -d '{"svcName":"teacher", "idl":"../teacher.thrift", "type":"PATH"}' 
 YOUR_SERVER_ADDRESS/idl-manage/teacherservice
```

返回参数：与请求参数相同

#### 4. 根据服务名称获取idl

请求路径：/idl-manage/:svcname

请求方法：GET

请求参数：

请求示例：

```shell
curl -X GET 
YOUR_SERVER_ADDRESS/idl-manage/teacherservice
```

返回参数：与增加idl的请求参数相同

#### 5. 列出idl

请求路径：/idl-manage/?page_id=&page_size=

请求方法：GET

请求参数：

请求示例：

```shell
curl -X GET
"YOUR_SERVER_ADDRESS/idl-manage?page_id=1&page_size=5"
```

## 三、使用说明

以rpc/teacher为例

1. 通过idl管理平台添加

   svcName是teacher在服务注册中心注册时的服务名

   如果选择type为PATH，则要确保“idl"字段中的地址相对于http-server/main.go是可以访问的

   如果选择type为Content，则”idl”字段应该是thrift文件的内容

```shell
curl -X POST 
-H "Content-Type:application/json"
-d '{"svcName":"teacherservice", "idl":"../idl/teacher.thrift", "type":"PATH"}'
IDL_MANAGE_ADDRESS/idl-manage
```

2. 确保该rpc服务为运行状态

3. 通过http-sever调用服务

   请求参数和返回参数均与直接向rpc服务发送请求时一致

```
curl -X POST 
-H "Content-Type:application/json" 
-d '{"name":"sherry", "id":5}' 
HTTP_SERVER_ADDRESS/apigateway/teacherservice/Register
```

## 四、实现细节

1. 使用sqlc生成CRUD的代码
2. 使用json泛化调用
3. 使用etcd进行服务注册和发现
4. 增加了对kitex客户端的缓存以提高性能
5. idl管理平台仅实现了后端部分
6. 在idl有变化时，会生成新的客户端
7. 为了将网关与idl管理平台独立开来，通过直接查找数据库的方式获取idl，而不是向idl管理平台发送请求

![](img/屏幕截图 2023-07-30 201826.png)

## 五、测试方案

### 1. 单元测试

对http-server和idl-manage的重要模块均进行了单元测试，详情可查看各xx_test.go文件

### 2. 性能测试

在http-server/main_test.go中使用go自带的benchmark进行了并行测试，通过设置并发度，模拟多个客户端同时发送请求的情况

![](img/屏幕截图 2023-07-28 093958.png)

-总共用了 16 个 CPU 核心

-执行次数：100 次

-每次耗时：28340315 ns

## 六、优化方案

增加了clients sync.Map作为泛化调用客户端的缓存，只有当初次向该服务发送请求和idl更新时，会触发kitex客户端的初始化

优化后的测试数据

![](img/屏幕截图 2023-07-28 095821.png)

-总共用了 16 个 CPU 核心

-执行次数：66 次

-每次耗时：18994985 ns