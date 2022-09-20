DROP DATABASE IF EXISTS test_db;
CREATE DATABASE test_db;
USE test_db;

CREATE TABLE hotVar(
    uid bigint primary key auto_increment,
    name char(255),
    value bigint,
    createtime int NOT NULL DEFAULT 0
)DEFAULT CHAR SET = utf8mb4;
ALTER TABLE hotVar ADD INDEX(name);
