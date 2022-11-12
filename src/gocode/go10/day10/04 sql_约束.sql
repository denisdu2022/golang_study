#约束 约束就是对字段的限制
#default 约束 默认值
#not null 非空约束
#unique  唯一约束
#主键约束 primary key auto_increment自增
#如果没有主键从上到下找到第一个非空且唯一的设置为主键,找不到就设置默认主键

create table c1
(
    name   varchar(32) [约束],
    gender varchar(32) default "男",
    age    int not null,
    desc   varchar(32)
);

create table s1
(
    name   varchar(32),
    gender varchar(32) default "男",
    age    int not null,
    des    varchar(32)
);

show create table s1;

create table s2
(
    id      int primary key auto_increment,
    name    varchar(32) not null unique,
    gender  varchar(32) default "男",
    age     int         not null,
    des     varchar(32),
    clas_id int
);

create table clas
(
    id   int primary key auto_increment,
    name varchar(32) not null unique
);

create table s3
(
    id      int primary key auto_increment,
    name    varchar(32) not null unique,
    gender  varchar(32) default "男",
    age     int         not null,
    des     varchar(32),
    clas_id int,
    FOREIGN KEY (clas_id) REFERENCES clas3 (id) ON DELETE CASCADE ON UPDATE CASCADE -- 建立外键约束
);
create table clas3
(
    id   int primary key auto_increment,
    name varchar(32) not null unique
);

create table s4
(
    id      int primary key auto_increment,
    name    varchar(32) not null unique,
    gender  varchar(32) default "男",
    age     int         not null,
    des     varchar(32),
    clas_id int,
    FOREIGN KEY (clas_id) REFERENCES clas4 (id) -- 建立外键约束
);
create table clas4
(
    id   int primary key auto_increment,
    name varchar(32) not null unique
);