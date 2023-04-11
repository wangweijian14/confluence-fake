package controller

import (
	"context"
	"fmt"

	v1 "confluence_fake/api/v1"
	"confluence_fake/internal/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) FakeGetSpaceList(ctx context.Context, req *v1.FakeGetSpaceListReq) (res *v1.FakeGetSpaceListRes, err error) {
	// service.Fakers().FakeCheckConfig(ctx, consts.SetConfigSuccess)
	out, err := service.Fakers().GetAllFakeSpace(ctx, req.Limit, req.Start)
	if err != nil {
		return nil, err
	}
	v, err := service.Fakers().GetCache(ctx).Get(ctx, "success")
	g.Log().Info(ctx, "v ", v.Bool())

	if !v.Bool() {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(401, "error:401")
	} else {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(200, out)
	}
	return
}

func (c *cHello) FakeGetGroupList(ctx context.Context, req *v1.FakeGetGroupListReq) (res *v1.FakeGetGroupListRes, err error) {
	out, err := service.Fakers().GetAllGrout(ctx, req.Limit, req.Start)
	if err != nil {
		return nil, err
	}
	v, err := service.Fakers().GetCache(ctx).Get(ctx, "success")
	g.Log().Info(ctx, "v ", v.Bool())

	if !v.Bool() {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(401, "error:401")
	} else {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(200, out)
	}
	return
}

func (c *cHello) FakeGetConfluenceUserList(ctx context.Context, req *v1.FakeGetConfluenceUserListReq) (res *v1.FakeGetConfluenceUserListRes, err error) {
	out, err := service.Fakers().GetAllConfulenceUserByGroupName(ctx, req.GpName, req.Limit, req.Start)
	if err != nil {
		return nil, err
	}
	v, err := service.Fakers().GetCache(ctx).Get(ctx, "success")
	g.Log().Info(ctx, "v ", v.Bool())

	if !v.Bool() {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(401, "error:401")
	} else {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(200, out)
	}
	return
}

func (c *cHello) FakeGetConfluenceUserByUserName(ctx context.Context, req *v1.FakeGetConfluenceUserByUserNameReq) (res *v1.FakeGetConfluenceUserListRes, err error) {
	out, err := service.Fakers().GetConfulenceUserUserName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	v, err := service.Fakers().GetCache(ctx).Get(ctx, "success")
	g.Log().Info(ctx, "v ", v.Bool())

	if !v.Bool() {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(401, "error:401")
	} else {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(200, out)
	}
	return
}

func (c *cHello) FakeConfig(ctx context.Context, req *v1.FakeConfigReq) (res *v1.FakeConfigRes, err error) {
	// service.Fakers().FakeCheckConfig(ctx, consts.SetConfigSuccess)

	switch req.Entity {
	case "group-count":
		err = service.Fakers().ConfigGroup(req.Gps)
	case "space-count":
		err = service.Fakers().ConfigSpaces(req.Total)
	case "user-count":
		err = service.Fakers().ConfigUser(req.Total, req.InGroup)
	case "success":
		service.Fakers().SetCache(ctx, "success", req.Success)
	}

	if err != nil {
		return nil, err
	}

	return &v1.FakeConfigRes{
		Msg: "ok",
	}, nil
}

func (c *cHello) ToIndex(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	v, err := service.Fakers().GetCache(ctx).Get(ctx, "success")
	spaces, err := service.Fakers().GetAllFakeSpace(ctx, 0, 0)
	groups, err := service.Fakers().GetAllGrout(ctx, 0, 0)
	users, err := service.Fakers().GetAllConfulenceUserByGroupName(ctx, "confluence-users", 0, 0)

	params := map[string]interface{}{
		"success":   v.Bool(),
		"spaces":    spaces,
		"groups":    groups,
		"users":     users,
		"jsonSpace": gjson.New(spaces.Results).MustToJsonString(),
		"jsonGroup": gjson.New(groups.Results).MustToJsonString(),
		"jsonUser":  gjson.New(users.Results).MustToJsonString(),
	}
	// 渲染模板
	g.RequestFromCtx(ctx).Response.WriteTpl("index/index.html", params)
	return
}

func (c *cHello) FakeZipPath(ctx context.Context, req *v1.FakeZipPathReq) (res *v1.IndexRes, err error) {
	v, err := service.Fakers().GetCache(ctx).Get(ctx, "success")
	g.Log().Info(ctx, "v ", v.Bool())

	if !v.Bool() {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(401, "error:401")
	} else {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(200, fmt.Sprintf("http://59.110.32.216:8000/download/temp/%d.xml.zip", gtime.Now().UnixMilli()))
	}
	return
}

func (c *cHello) FakeFileDownload(ctx context.Context, req *v1.DownloadReq) (res *v1.IndexRes, err error) {
	v, err := service.Fakers().GetCache(ctx).Get(ctx, "success")
	g.Log().Info(ctx, "v ", v.Bool())

	if !v.Bool() {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(401, "error:401")
	} else {
		gfile.CopyFile("resource/test.xml.zip", "resource/cache/"+req.FileName)
		g.RequestFromCtx(ctx).Response.ServeFileDownload("resource/cache" + req.FileName)
	}
	return
}
