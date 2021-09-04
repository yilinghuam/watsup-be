DROP DATABASE IF EXISTS watsup;

CREATE DATABASE watsup;

CREATE TABLE watsup.lists (
    id  INT AUTO_INCREMENT NOT NULL,
    category VARCHAR(255) NOT NULL,
    main_goal VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    frequency   INT NOT NULL,
	cycle       VARCHAR(255) NOT NULL,
	stars       INT NOT NULL,
    user   VARCHAR(255) NOT NULL,
	state	VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE watsup.active(
	id INT NOT NULL,
    current_frequency INT DEFAULT 0,
    completed VARCHAR(255) DEFAULT 'not',
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (id) REFERENCES lists(id)
);
USE watsup;
-- Create a new account with name and address
INSERT INTO watsup.lists (category, main_goal, name, description, frequency,cycle,stars,user,state) VALUES ('health','cycle','cycle', 'null', 2, 'week',13,'ling','active'),
('lifestyle','beautiful face', 'Facial' ,'put mask 5 times a month', 5, 'month',2,'ben','active');

INSERT INTO watsup.active(id,completed,start_date,end_date) VALUES (1, 'incomplete',CAST('2021-09-6' AS datetime),ADDDATE(CAST('2021-09-6' AS datetime),INTERVAL 1 WEEK)),
(2,'incomplete',CAST('2021-09-5' AS datetime),ADDDATE(CAST('2021-09-5' AS datetime),INTERVAL 1 MONTH));

