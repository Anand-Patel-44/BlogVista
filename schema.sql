CREATE TABLE users(
username VARCHAR(30) NOT NULL,
password VARCHAR(100) NOT NULL,
bio VARCHAR(125)NOT NULL ,
color VARCHAR(20) NOT NULL,
profession VARCHAR(20)NOT NULL,
PRIMARY KEY(username)
);

CREATE TABLE blogs(
blogID INT AUTO_INCREMENT NOT NULL,
title VARCHAR(50)NOT NULL,
blogText VARCHAR(500)NOT NULL,
username varchar(30) NOT NULL,
PRIMARY KEY(blogID),
FOREIGN KEY(username) references users(username)
);
