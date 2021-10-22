create database if not exists DailyRecommend;
USE mysql;
update user set host = '%' where user = 'root';
flush privileges;
