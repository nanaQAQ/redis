# redis study

## 环境

虚拟机挂起后docker镜像会失去连接

1. 重启docker服务,`systemctl restart docker`
2. 重新启动容器，`docker restart daca2576c021(容器ID)`
3. 重新连接成功

### redis

go get -u github.com/go-redis/redis

在redis-cli中开启monitor模式

 docker run -p 6379:6379 --name redis -v /redis/data:/data -d redis

### mysql

 docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql

# 黑马点评

## 短信登陆

### 发送短信验证码

提交手机号->校验手机号->生成验证码->保存验证码到session->发送验证码

### 短信验证码登录、注册

提交手机号和验证码->校验验证码->根据手机号查询用户->存在则保存到session，不存在则创建新用户->保存到数据库

### 校验登录状态

请求并携带cookie->从session中获取用户->不存在则拦截，存在则保存在ThreadLocal（线程独立存储空间）中

## 商户查询缓存

## 达人探店

## 优惠券秒杀

## 好友关注

## 附近的商户

## 用户签到

## UV统计