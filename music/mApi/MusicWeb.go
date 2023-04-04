package mApi

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"WangYiYunDemo/music/mService"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Hello(context *gin.Context) {
	log.Println(">>>>> Hello gin Success! <<<<<")
	context.String(http.StatusOK, "Music Ok!")

}

// 为了swag识别接口， 通常需要加一些接口描述， 是格式化的， 需要阅读文档
// todo: complete swag description
func ListSongs(context *gin.Context) {
	log.Println("come into ListSongs function")
	req := new(mdDef.ListSongsReq)
	err := context.BindJSON(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	resp, err := mService.ListSongs(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	//todo : 完善http response结构体封装， 在data的基础上加上success， code， msg等
	context.JSON(http.StatusOK, resp)
}

func ListComments(context *gin.Context) {
	log.Println("come into ListComments function")
	req := new(mdDef.ReturnCommentReq)
	err := context.ShouldBind(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
	}

	resp, err := mService.ListComment(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
	}
	context.JSON(http.StatusOK, resp)
}

func UploadComment(context *gin.Context) {
	log.Println("come into UploadComment function")
	req := new(mdDef.UploadCommentReq)
	err := context.ShouldBind(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
	}
	resp, err := mService.WriteComment(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
	}
	context.JSON(http.StatusOK, resp)
}
