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

--查看所有数据库
show databases;
--创建数据库
create database fastapi;
--创建数据库时确认是否存在
create database if not exists fastapi;
--进入数据库
use go10;
--查询当前所在数据库
select database();

use fastapi;
select database();
--查询某个数据库的创建信息
show create database go10;
--创建数据库并指定字符集
create database go12 character set gbk;

show create database go11;
--删除数据库
drop database go12;
--修改数据库字符集
alter database go11 default character set gbk;

use go11;

use go10;

--查看所有表
show tables;

--查看某张表的创建信息
show create table userinfo;


--创建表结构
    --decimal 8是8位99999999 2是小数点后的两位999999.99
    --not null 不能为空
create table userinfo(
    name varchar(5) not null,
    age int,
    birth date,
    salary decimal(8,2)
);

--查看某张表的结构
desc userinfo;
--查看表记录
select * from userinfo;
--删除表
drop table userinfo;

--修改表结构
--添加一个字段
alter table userinfo add gender varchar(32) not null after age;
--修改字段名字
alter table userinfo change gender sex varchar(20) not null first;
--修改表名
alter table userinfo rename to userinfos;

desc userinfos;

alter table userinfos modify  name varchar(32)  not null first;

alter table userinfos modify sex varchar(32) null;

alter table userinfos change sex gender varchar(32);
















