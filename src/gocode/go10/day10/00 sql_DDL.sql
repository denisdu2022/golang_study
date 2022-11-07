--显示数据库
show databases;
--查看某一个数据库的创建信息
show create database go10;
--创建数据库
create database if not exists go10;
--进入数据库
use go11;
--查看当前所在的数据库
select database();
--删除数据库
drop database go11;
--创建数据库并指定字符集为 gbk
create database go11 default character set gbk;

show create database go11;
--修改数据库字符集
alter database go11 character set utf8;

