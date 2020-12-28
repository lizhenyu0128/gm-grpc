# grpc
国密版grpc库，支持双证书国密tls和单证书国密tls。
##使用说明 
##### 1. 单证书tls模式  
编译项目的过程中，添加-tags single_cert标志，表示使用单证书模式，编写代码参考本项目test目录下的cli和server两个文件的代码。

##### 2. 双证书tls模式  
默认编译使用的是双证书tls模式，不需要额外添加标志，编写代码参考本项目test目录下的double-cli和double-server两个文件的代码。

安装
------------

安装此包最简单的方式就是运行以下命令:

```
$ go get -u github.com/Hyperledger-TWGC/grpc
```
