
# douyin

## 功能描述

极简版抖音，可以实现登录、注册、刷视频、点赞、评论、关注等功能



## 项目结构

```sh
│  go.mod
│  go.sum
│  main.exe
│  main.go    # main 函数
│  README.md
│  router.go    #路由
│
├─controller    #逻辑处理函数
│      comment.go
│      common.go
│      demo_data.go
│      favorite.go
│      feed.go
│      publish.go
│      relation.go
│      user.go
│
├─dao    #数据库相关，gorm相关接口
│      comment.go
│      common.go
│      db_init.go
│      favorite.go
│      feed.go
│      publish.go
│      relation.go
│      user.go
│
├─model    #生成数据表相关的结构体
│      model.go
│
├─pkg    
│  └─constrant    #存放常量
│          constrant.go
│
└─public    #存放视频、封面
        14_8e847570a9eecb15b4226c0590d0eb4b.mp4
        14_8e847570a9eecb15b4226c0590d0eb4b.mp4.png
        16_669fe936e68e3d735fddbeea916058bc.mp4
        16_669fe936e68e3d735fddbeea916058bc.mp4.png
        bear.mp4
        data

```



## 运行环境

由于本项目使用 FFmpeg 第三方工具来完成视频截取第一帧作为封面并存储到本机，

所以运行之前需要预先安装 FFmpeg 到本机并配置好环境变量。

参考链接：[FFmpeg](http://ffmpeg.org/)