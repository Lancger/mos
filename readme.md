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
