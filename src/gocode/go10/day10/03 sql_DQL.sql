#join查询

#查询三国演义的出版社的名称
select publisher.name
from book,
     publisher
where book.id = publisher.id
  and title = "三国演义";
#inner join
select publisher.name
from book
         inner join publisher on book.pub_id = publisher.id
where title = "三国演义";

#left join

desc book;

#修改字段属性可以为空
alter table book
    modify pub_id int(11) null;

select *
from book
         left join publisher on book.pub_id = publisher.id;

#right join
select *
from publisher
         right join book on book.pub_id = publisher.id;

#查询yuan出版的书籍名称
desc book;
desc author;
desc publisher;

select book.title
from book
         inner join book2author on book.id = book2author.book_id
         inner join author on book2author.author_id = author.id
where author.name = "yuan";

#查询手机号等于100的作者出版过那些书籍名称
select book.title, author.NAME, authordetail.tel
from book
         inner join book2author on book.id = book2author.book_id
         inner join author on book2author.author_id = author.id
         inner join authordetail on author.id = authordetail.author_id
where authordetail.tel = 911;








