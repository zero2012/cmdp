## cmdp 命令行提示工具（远程同步）

![demo](http://cloud.yurencloud.com/index.php/s/S0j84GRI7Xkr41H/download)

#### 使用场景：

命令行参数太多，太难记，平常可能会为了1条命令，新建一个笔记，来记录，查询麻烦。
虽然linux有ctrl+R查询，但对于经常创建linux系统，使用不同的服务器，不同的电脑的人来说，命令提示不能同步到所有
电脑。

#### 作用：

快速记录一条命令，代码，账号，密码，文字等内容，并允许添加注释。可根据关键词快速查找到相关的内容，并且
可以远程同步。
不同电脑，系统，服务器，只要登录相同用户，即可同步获取所有提示。

### 1.安装

#### 1.1 若你已经安装go
可直接使用下面命令安装
~~~
go get github.com/yurencloud/cmdp
~~~

### 2.使用
~~~
以下均以linux下操作为例，windows下改为cmdp.exe命令就可以
~~~
#### 2.1注册
~~~
cmdp -register -u USERNAME -p PASSWORD
~~~
注册后自动在GOROOT目录下生成token文件，*请不要删除*，token有效期1年

#### 2.2登录
若已经注册，或token过期，或token删除，请使用登录功能，会生成新的token
~~~
cmdp -login -u USERNAME -p PASSWORD
~~~

#### 2.3添加命令行提示
~~~
cmdp -c "docker start mysql" -m "使用docker启动mysql容器"
~~~

#### 2.4搜索命令行
~~~
cmdp mysql
~~~
显示结果，彩色
~~~
docker start mysql  使用docker启动mysql容器 id:2
mysql -uroot -p     登入mysql id:31
results: 2
~~~

#### 2.5删除命令行
先查询，后根据结尾显示的id进行删除
~~~
cmdp -d ID
~~~

#### 2.6修改用户登录密码
~~~
cmdp -reset -p NEW_PASSWORD
~~~
