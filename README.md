[查看中文文档](https://github.com/yurencloud/cmdp/blob/master/README-zh.md)
## cmdp command prompt tool

## Usage Scenes

It is hard to remember too many command parameters. It might also be troublesome
to create a note to record maybe just for a single command and look it up. Though
you guys can look up a command by Ctrl+R in linux, those who always build linux
system, and use different servers or computers, will find the command prompt
actually cannot sync to all the computers.

## Function
Cmdp can quickly record a command, code, account, passcode, text, etc., notes are
also allowed. Relevant contents can be quickly found through key words, and even
sync remotely to other computers. A user can acquire all the reminders once he or
she logs in with the same user name, even in different computers with different
systems and servers.

## 1. Installation
#### 1.1 If you have installed `go`
You can directly install it according to the command listed below.
~~~
go get github.com/yurencloud/cmdp
~~~

#### 1.2 Download & Installation

[windows download](https://github.com/yurencloud/cmdp/raw/master/windows/cmdp.exe)

[mac download](https://github.com/yurencloud/cmdp/raw/master/mac/cmdp)

[linux download](https://github.com/yurencloud/cmdp/raw/master/linux/cmdp)

You can directly use the tool, but are suggested adding the directory of the command tool to the system PATH.

If `Permission denied` occurs in mac or linux, please execute permission by the following method:
~~~
chmod +x cmdp
~~~
#### 1.3 Use node to simply package on the basis of cmdp, while no color shows after output.
~~~
npm install -g yu.cmdp
~~~

## 2. Usage
All the examples are listed under linux. You can change it into cmdp.exe so that it can work in windows.

#### 2.1 Register
~~~
cmdp -register -u USERNAME -p PASSWORD
~~~
token file will be created to GOROOT after register, do not delete token file, token will expired after one year.
#### 2.2 Login
If you have registered or the token has expired, or was deleted, please use the function of log in, then the new token will turn up.
~~~
cmdp -login -u USERNAME -p PASSWORD
~~~
#### 2.3 Add command prompt
~~~
cmdp -c "docker start mysql" -m "use docker to create a mysql container"
~~~
#### 2.4 Search command prompt
~~~
cmdp mysql
~~~
The result appears, colored
~~~
docker start mysql  use docker to create a mysql container id:2
mysql -uroot -p     login mysql id:31
results: 2
~~~

#### 2.5 Delete command
Look it up first, then delete it according to the ID that shows in the end.
~~~
cmdp -d ID
~~~

#### 2.6 Change the users’ login password
~~~
cmdp -reset -p NEW_PASSWORD
~~~
