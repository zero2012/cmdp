## cmdp 命令行提示工具（远程同步）

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

#### 1.2 下载安装


[windows下载](https://github.com/yurencloud/cmdp/raw/master/windows/cmdp.exe)

[mac下载](https://github.com/yurencloud/cmdp/raw/master/mac/cmdp)

[linux下载](https://github.com/yurencloud/cmdp/raw/master/linux/cmdp)

工具可以直接使用，但建议将命令工具所在目录添加到系统路径`PATH`中

在mac或linux若出现`Permission denied`问题，请用下面方法添加可执行权限
~~~
chmod +x cmdp
~~~

#### 1.3 npm 安装
~~~
// 使用node在cmdp基础上简单封装，缺点，输出没有色彩
npm install -g yu.cmdp
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