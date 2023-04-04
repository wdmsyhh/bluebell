
- 连接 mysql
```shell
docker exec -it bluebell_backend-mysql8019-1 bash

mysql -u root -p
```

- mysql数据库中查看当前使用的数据库是哪个数据库
```shell
select database();
# 查询有哪些表
show tables;
```