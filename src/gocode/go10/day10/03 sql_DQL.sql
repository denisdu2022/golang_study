#join查询

#查询三国演义的出版社的名称
select * from book,publisher where book.id = publisher.id;