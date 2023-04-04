package msDef

import "WangYiYunDemo/music/mDAO/mdDef"

type ListCommentResp struct {
	Length           int                   `json:"length"`
	CommentMusicData []*mdDef.CommentBasic `json:"commentMusic_data"`
}

type WriteCommentResp struct {
	IsSuccess string `json:"is_success"`
}
