
1. jenkins 构建历史上报接口
```
curl -X POST -H "'Content-type':'application/json'" -d "{\"title\": \"构建标题\", \"project\": \"所属项目\", \"module\": \"构建模块\", \"hosts\": \"${hosts}\", \"build_number\": \"${BUILD_NUMBER}\", \"tag\": \"${tag}\", \"build_user\": \"${BUILD_USER}\", \"build_status\": \"${BUILD_STATUS}\", \"console_url\": \"${BUILD_URL}\"}" 'http://180.101.141.202:335/JenkinsPost'
```