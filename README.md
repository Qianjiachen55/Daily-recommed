# Daily-recommed
1.0：每天通过api获取一些新闻推荐给邮箱，不经过算法

1.xxxxx 修复bug并测试

2.0: 前端埋点获取用户点击行为

2.xxxxx 增加feature

3.0: 整合docker,docker-compose,k8s

4.0: 尝试推荐系统
- content-base
- user base cf
- item base cf


### 记录开发过程
```shell
21.10.22.12.11
修改 utils/error.go :添加处理错误和panic的方法,defer PanicFun() 记录panic到日志中 
```