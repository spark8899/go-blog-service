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
$ curl -X POST 'http://127.0.0.1:8000/auth' -d 'app_key=test&app_secret=test001'
{"token":"xxx-xxx-xxxx-xxxxx"}
$ curl -X GET http://127.0.0.1:8000/api/v1/tags -H 'token: xxx-xxx-xxxx-xxxxx'
```
