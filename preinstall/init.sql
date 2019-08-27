CREATE TABLE articles (id SERIAL, title varchar(20), content varchar(200), author varchar(20));

INSERT INTO articles (title, content, author) values ('aaa-title', 'bbb-content', 'ccc-author');
INSERT INTO articles (title, content, author) values ('ddd-title', 'eee-content', 'fff-author');

SELECT * FROM articles