package controller

import (
	"bytes"
	"douyin/dao"
	"douyin/model"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	user, v := TokenIsValid(token)
	if !v {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "token is not valid."})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

	filePath := StrBulider("./public/", finalName)

	// 此处返回值为./public/image.png 若想直接存数据库就启用
	_, err = GetSnapshot(filePath, finalName)
	if err != nil {
		log.Println("--GetSnapshot--:", err)
	}
	finalImageName := StrBulider(finalName, ".png")

	video := model.Video{
		AuthorId:      user.Id,
		Title:         title,
		PlayUrl:       finalName,
		CoverUrl:      finalImageName,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	err = dao.Mgr.InsertToVideo(video)
	if err != nil {
		fmt.Println("InsertToVideo error:", err)
	}
}

func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")

	_, b := TokenIsValid(token)
	if !b {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Token is not valid."},
		})
		return
	}

	videos, err := dao.Mgr.GetVideoByUserId(userId)
	if err != nil {
		log.Println("ReturnVideoByUserId:", err)
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}

func GetSnapshot(videoPath, snapshotPath string) (snapshotName string, err error) {
	snapshotPath = StrBulider("./public/", snapshotPath)
	frameNum := 1
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Println("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Println("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Println("生成缩略图失败：", err)
		return "", err
	}

	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return snapshotName, nil
}
