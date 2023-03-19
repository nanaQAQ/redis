# redis study

## 环境

### redis：ubuntu docker 6379

go get -u github.com/go-redis/redis

在redis-cli中开启monitor模式

### mysql：ubuntu docker 3309

 docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql