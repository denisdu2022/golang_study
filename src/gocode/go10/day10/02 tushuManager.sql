#图书管理系统
create table tushuInfo(
    title varchar(32) unique not null, --unique唯一约束
    price decimal(6,2) --9999.99
    pub_date date,
    publish_name varchar(32),
    publish_addr varchar(32)
);

-- 书籍表
CREATE TABLE book(
                     id INT PRIMARY KEY AUTO_INCREMENT,
                     title VARCHAR(32),
                     price DOUBLE(5,2),
                     pub_id INT NOT NULL
)ENGINE=INNODB CHARSET=utf8;


-- 出版社表
CREATE TABLE publisher(
                          id INT PRIMARY KEY AUTO_INCREMENT,
                          name VARCHAR(32),
                          email VARCHAR(32),
                          addr VARCHAR(32)
)ENGINE=INNODB CHARSET=utf8;

-- 插入数据
INSERT INTO book(title,price,pub_id) VALUES
                                         ('西游记',15,1),
                                         ('三国演义',45,2),
                                         ('红楼梦',66,3),
                                         ('水浒传',21,2),
                                         ('红与黑',67,3),
                                         ('乱世佳人',44,6),
                                         ('飘',56,1),
                                         ('放风筝的人',78,3);

INSERT INTO publisher(id,name,email,addr) VALUES
                                              (1,'清华出版社',"123","bj"),
                                              (2,'北大出版社',"234","bj"),
                                              (3,'机械工业出版社',"345","nj"),
                                              (4,'邮电出版社',"456","nj"),
                                              (5,'电子工业出版社',"567","bj"),
                                              (6,'人民大学出版社',"678","bj");

-- 作者表
CREATE TABLE author(
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       NAME VARCHAR(32) NOT NULL
)ENGINE=INNODB CHARSET=utf8;

-- 作者表和书籍表的多对多关系表
CREATE TABLE book2author(
                            id INT NOT NULL UNIQUE AUTO_INCREMENT,
                            author_id INT NOT NULL,
                            book_id INT NOT NULL
)ENGINE=INNODB CHARSET=utf8;

-- 插入数据

INSERT INTO author(NAME) VALUES
                             ('yuan'),
                             ('rain'),
                             ('alvin'),
                             ('eric');

-- 插入关系数据
INSERT INTO book2author(author_id,book_id) VALUES
                                               (1,1),
                                               (1,2),
                                               (2,1),
                                               (3,3),
                                               (3,4),
                                               (1,3);


CREATE TABLE authorDetail(
                             id INT PRIMARY KEY AUTO_INCREMENT,
                             tel VARCHAR(32),
                             addr VARCHAR(32),
                             author_id INT NOT NULL unique -- 也可以给author添加一个关联字段：   alter table author add authorDetail_id INT NOT NULL
)ENGINE=INNODB CHARSET=utf8;

-- 插入数据
INSERT INTO authorDetail(tel,addr,author_id) VALUES
                                                 ("110","北京",1),
                                                 ("911","成都",2),
                                                 ("119","上海",3),
                                                 ("111","广州",4);

#子查询  查询三国演义的出版社的名称 查询一对多

#查看表结构
desc book;

select pub_id from book where title = "三国演义";

desc publisher;

select name as "出版社名字" from publisher where id =2;

#查询多对多
#查询手机号等于110的作者出版过那些书籍名称
desc authorDetail;

select tel as "手机号", author_id from authorDetail where tel = 110;

desc author;
#查询作者名
select name as "作者" from author where id =1;

desc book2author;
select book_id from book2author where author_id = 1;

select title as "书名" from book where id in(1,2,3);


