show tables;

desc userinfos;

#插入表记录
#insert table_name 字段 values 值
#插入所有字段的值
insert userinfos value ("daerwen","男",23,"1995-03-12",3000);
#按照字段插入数据
insert userinfos (name,gender,age,birth,salary) value ("dafenqi","男",22,"1996-03-12",5000);
#查询表中的所有记录
select * from userinfos;

#批量插入
insert userinfos values
                     ("bijiasuo","男",23,"1995-03-12",2000),
                     ("bijia1","男",23,"1995-03-12",3000),
                     ("bijia2","男",23,"1995-03-12",2000),
                     ("bijias3","男",23,"1995-03-12",4000),
                     ("bijiasu4","男",23,"1995-03-12",5000),
                     ("bijiasuo5","男",23,"1995-03-12",4000);

#set写法
insert userinfos set name = "bijialuo",age=22;

#查询表记录
#select 字段,字段2...* from table_name
#查询所有表记录  不建议
select * from userinfos;
#按照字段查询  建议这样写
select name,gender,age,birth,salary from userinfos;
#按条件查询 年龄=23
select name,gender,age,birth,salary from userinfos where age = 23;
#按条件查询 年龄!=23
select name,gender,age,birth,salary from userinfos where age != 23;
#按条件查询 薪资大于3000
select name,gender,age,birth,salary from userinfos where salary > 3000;
#按条件查询 年龄小于23并且薪资大于5000的
select name,gender,age,birth,salary from userinfos where age < 23 and salary > 5000;
#按条件查询 薪资在5000到10000之间的
select name,gender,age,birth,salary from userinfos where salary between 5000 and 10000;
#按条件查询 薪资在2000,3000,4000,5000的in操作
select name,gender,age,birth,salary from userinfos where salary in (2000,3000,4000,5000);
#按条件查询 名字以a开头
select name,gender,age,birth,salary from userinfos where name like "a%";
#按条件查询 名字以a开头或者薪资大于3000的
select name,gender,age,birth,salary from userinfos where name like "a%" or salary > 3000;

#order by 排序 薪资排序
select name,gender,age,birth,salary from userinfos where salary > 2000 order by salary;
#降序排序
select name,gender,age,birth,salary from userinfos where salary > 2000 order by salary desc;
#降序排序 按年龄 排序
select name,gender,age,birth,salary from userinfos where salary > 2000 order by salary desc,age;

#limmit限制条数 前两条
select name,gender,age,birth,salary from userinfos limit 2;
#limmit限制条数 偏移量 offset number
select name,gender,age,birth,salary from userinfos limit 2,4;

#disyinct去重  取出重复的name
select distinct name,gender,age,birth,salary from userinfos ;
#有主键的时候去重没有意义
select distinct * from userinfos;

#分组查询: max min avg count
#查询男生和女生的平均薪资
#select * from userinfos group by gender; 这样写是错误的
#as是别名
select gender,avg(salary) as "平均薪资" from userinfos group by gender;
#查询统计男女的人数
select gender as "性别",count(gender) as "人数" from userinfos group by gender;
#或者
select gender as "性别",count(*) as "人数" from userinfos group by gender;
#增加一个部门
alter table userinfos add dep varchar(32);
#查询每个部门的最高薪资
select dep,max(salary) from userinfos group by dep;

#having 筛选条件语句 分组之后再过滤  where是分组之前做过滤
#查询平均薪资大于5000的每个部门以及平均薪资
select dep as "部门",avg(salary) as avg_salary from userinfos group by dep having avg_salary > 5000;
#查询年龄大于22的的所有记录的每个部门以及平均薪资
select dep as "部门",avg(salary) as "avg_salary" from userinfos where age > 22 group by dep ;


#修改记录: update
update userinfos set birth="1991-01-12" where name = "daerwen" or name = "dafenqi";

select name,gender,age,birth,salary from userinfos where name ="daerwen" or name = "dafenqi";

#删除记录 delete
delete  from userinfos where name="daerwen";
















