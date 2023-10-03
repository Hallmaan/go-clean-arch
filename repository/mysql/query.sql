CREATE TABLE products (
   id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   name varchar(255) DEFAULT NULL COMMENT 'product name')

   
INSERT into products (name) VALUES ('RD1652')


CREATE TABLE transactions (
   id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   transaction_name VARCHAR(100) NOT NULL,
   product_id int(10) NOT NULL
)