package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type FakeGetSpaceListReq struct {
	// g.Meta `path:"rest/api/space" tags:"fake" method:"get" summary:"checkconfig"`
	g.Meta `path:"/rest/api/space" tags:"fake" method:"get" summary:"checkconfig"`
	Start  int `json:"start"`
	Limit  int `json:"limit"`
}

type FakeGetGroupListReq struct {
	g.Meta `path:"/rest/api/group/" tags:"fake" method:"get" summary:""`
	Start  int `json:"start"`
	Limit  int `json:"limit"`
}

type FakeGetConfluenceUserListReq struct {
	g.Meta `path:"/rest/api/group/{GpName}/member" tags:"fake" method:"get" summary:""`
	GpName string `in:"path"  dc:"gpName"`
	Start  int    `json:"start"`
	Limit  int    `json:"limit"`
}

type FakeGetConfluenceUserByUserNameReq struct {
	g.Meta   `path:"/rest/mobile/1.0/profile/{UserName}" tags:"fake" method:"get" summary:""`
	UserName string `in:"path"  dc:"gpName"`
	Start    int    `json:"start"`
	Limit    int    `json:"limit"`
}

type FakeConfigReq struct {
	g.Meta  `path:"/fake/config/{Entity}" tags:"fake" method:"get" summary:"FakeConfigReq"`
	Entity  string   `in:"path"  dc:"Entity"`
	Success bool     `json:"success"` // success、unsuccess
	Total   int      `json:"total"`
	InGroup string   `json:"inGroup"`
	Gps     []string `json:"gps"`
}

type FakeConfigRes struct {
	Msg string `json:"msg"`
}

type FakeGetSpaceListRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type FakeGetGroupListRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type FakeGetConfluenceUserListRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type IndexReq struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页"`
}
type IndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
