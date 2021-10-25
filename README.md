To set up the database container:
```
docker run --name sample-api-mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=college -p 3306:3306 -d mysql:latest
```
```
docker exec -it sample-api-mysql bash
```
SQL commands:
```
mysql -u root -p

use college;

CREATE TABLE student (
id int NOT NULL AUTO_INCREMENT,
name varchar(255) NOT NULL,
major varchar(255) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (id)
); 
```
