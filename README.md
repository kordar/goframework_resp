# goframework-resp

## 安装

```go
go get github.com/kordar/goframework_resp v0.0.2
```

## 初始化

- `gin`框架

```go
f := func(c interface{}, httpStatus int, code int, message string, data interface{}, count int64) {
    ctx := c.(*gin.Context)
    if data == nil {
        ctx.JSON(httpStatus, map[string]interface{}{"code": code, "message": message})
        return
    }

    if count < 0 {
        ctx.JSON(httpStatus, map[string]interface{}{"code": code, "message": message, "data": data})
        return
    }

    ctx.JSON(httpStatus, map[string]interface{}{"code": code, "message": message, "data": data, "count": count})
}
goframework_resp.RegResultCallFunc(f)
```
