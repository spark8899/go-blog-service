# go-blog-service
tech golang blog-service

# add auth
vim internal/routers/router.go
```
- apiv1.Use() //middleware.JWT()
+ apiv1.Use(middleware.JWT()) //middleware.JWT()
```

# get token
```
$ curl -v -X POST 'http://127.0.0.1:8000/auth' -d 'app_key=test&app_secret=test001'
{"token":"xxx-xxx-xxxx-xxxxx"}
$ curl -v -X GET http://127.0.0.1:8000/api/v1/tags -H 'token: xxx-xxx-xxxx-xxxxx'
```

# tag op
```
$ curl -v -XPOST 'http://127.0.0.1:8000/api/v1/tags?name=11&state=1&created_by=test'
$ curl -v -XPOST 'http://127.0.0.1:8000/api/v1/tags?name=22&state=1&created_by=test'

$ curl -v '127.0.0.1:8000/api/v1/tags' 
```
