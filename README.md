快海报 SDK for Golang
===

## Install

```bash
go get github.com/googolmo/khb-sdk-go
```


## Import

```golang
import "github.com/googolmo/khb-sdk-go"
```

### Use

```golang
ss := khb.Screenshot{
    Template: "jdp3yn4s",
    Data: map[string]interface{}{
        "backgroundColor": "#c8f1fd",
        "user": map[string]interface{}{
            "avatar":   "https://khb-sample.oss-cn-shanghai.aliyuncs.com/sample/girl_2.jpg",
            "nickname": "筱墨",
        },
        "tip":    "快海报专业设计师提供的模板设计. ",
        "cover":  "https://khb-sample.oss-cn-shanghai.aliyuncs.com/sample/watermelon.jpg",
        "qrcode": "https://khb-sample.oss-cn-shanghai.aliyuncs.com/sample/sample_qr_0.png",
        "brand":  "快海报",
        "slogan": "小程序分享海报生成服务",
        },
    Type: khb.TypeJpg,
}
result, err := ss.invoke("YOUR TOKEN")
```