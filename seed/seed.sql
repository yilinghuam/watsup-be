DROP DATABASE IF EXISTS watsup;

CREATE DATABASE watsup;

CREATE TABLE watsup.users (
    user_id  VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    PRIMARY KEY(user_id)
);

CREATE TABLE watsup.groupbuy(
	groupbuy_id INT AUTO_INCREMENT NOT NULL,
    user_id  VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    order_date VARCHAR(255) NOT NULL,
    closing_date VARCHAR(255) NOT NULL,
    delivery_options BOOLEAN NOT NULL,
    delivery_price INT DEFAULT 0,
    status VARCHAR(255) DEFAULT 'open',
    PRIMARY KEY (groupbuy_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE watsup.item(
	id INT AUTO_INCREMENT NOT NULL,
    user_id  VARCHAR(255) NOT NULL,
    item VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    quantity INT NOT NULL,
    groupbuy_id INT NOT NULL,
	order_id INT DEFAULT 0,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (groupbuy_id) REFERENCES groupbuy(groupbuy_id)
);

CREATE TABLE watsup.order(
	order_id INT AUTO_INCREMENT NOT NULL,
    groupbuy_id INT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    status VARCHAR(255) DEFAULT "awaiting payment",
    PRIMARY KEY (order_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (groupbuy_id) REFERENCES groupbuy(groupbuy_id)
);
ALTER TABLE watsup.order AUTO_INCREMENT=1;

USE watsup;
-- Create a new account with name and address
INSERT INTO watsup.users (
    user_id,
    email) VALUES ('ling','lingy93@gmail.com'),
('host','huamyiling@gmail.com');

INSERT INTO watsup.groupbuy(
    groupbuy_id,
    user_id,
    name,
    description,
    order_date,
    closing_date,
    delivery_options
    ) VALUES (1,'ling','pineapples','','12-11-2021','01-11-2021',false);

INSERT INTO watsup.groupbuy(
    groupbuy_id,
    user_id,
    name,
    description,
    order_date,
    closing_date,
    delivery_options,
    delivery_price
    )VALUES(2,'host','mooncakes','from sheraton','10-10-2021','04-10-2021',true,5);
-- groupbuy host items
INSERT INTO watsup.item(
    id,
    user_id,
    item,
    price,
    quantity,
    groupbuy_id
    ) VALUES (1,'ling','3 pineapples',5,20,1),
    (2,'ling','5 pineapples',8,22,1),
(3,'host','box of 6',1.5,5,2);
-- groupbuy user item 
INSERT INTO watsup.item(
    id,
    user_id,
    item,
    price,
    quantity,
    groupbuy_id,
    order_id
    ) VALUES (4,'host','3 pineapples',25,5,1,1),
    (5,'host','5 pineapples',64,8,1,1),
(6,'ling','box of 6',7.5,5,2,2),
(7,'ling','3 pineapples',25,5,1,3),
    (8,'ling','5 pineapples',64,8,1,3);
INSERT INTO watsup.order(
    order_id,
    groupbuy_id,
    user_id,
    address,
    status
    ) VALUES (1,1,'host','null','awaiting payment'),
    (2,2,'host','333b anchorvale','order successful'),
(3, 1,'ling','crowhurst drive','payment failed');