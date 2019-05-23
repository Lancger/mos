0、切换分支

```
#方式一
git clone https://github.com/Lancger/mos.git
git checkout -b iaas origin/iaas

git branch

#方式二
git clone -b iaas https://github.com/Lancger/mos.git
```

1. jenkins 构建历史上报接口
```
curl -X POST -H "'Content-type':'application/json'" -d {"title": "构建标题", "project": "所属项目", "module": "构建模块", "hosts": "'$hosts'", "build_number": "'$BUILD_NUMBER'", "tag": "'$tag'", "build_user": "'$BUILD_USER'", "status": "'$BUILD_STATUS'", "console_url": "'$BUILD_URL'"}

```

```
初始化数据库表结构
mysql -uroot -p123456 -e "create database db_mos;"

POST    http://127.0.0.1:8000/InitTable

{
    "code": 20000,
    "message": "初始化表成功"
}
```

# 二、运行
```
方式一：

 go build
 ./src -c config.yaml
 
 方式二：
 go run main.go -c config.yaml
```
