package fakers

import (
	"confluence_fake/importserv"
	"confluence_fake/internal/consts"
	"confluence_fake/internal/service"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/hashicorp/go-memdb"
)

type sFakers struct {
	// Create the DB schema
	Schema      *memdb.DBSchema
	SchemaGroup *memdb.DBSchema
	SchemaUser  *memdb.DBSchema
	// Create a new data base
	DBSpace *memdb.MemDB
	DBGroup *memdb.MemDB
	DBUser  *memdb.MemDB
	Cache   *gcache.Cache
}

func init() {
	service.RegisterFakers(New())
}

func New() *sFakers {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"space": {
				Name: "space",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"name": {
						Name:    "name",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"key": {
						Name:    "key",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Key"},
					},
					"type": {
						Name:    "type",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Type"},
					},
				},
			},
		},
	}

	schemaGroup := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"group": {
				Name: "group",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"type": {
						Name:    "type",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Type"},
					},
					"index": {
						Name:    "index",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Index"},
					},
				},
			},
		},
	}

	schemaUser := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": {
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "UserKey"},
					},
					"type": {
						Name:    "type",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Type"},
					},
					"status": {
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Status"},
					},
					"username": {
						Name:    "username",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Username"},
					},
					"displayName": {
						Name:    "displayName",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "DisplayName"},
					},
					"fullName": {
						Name:    "fullName",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "FullName"},
					},
					"email": {
						Name:    "email",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"unknownUser": {
						Name:    "unknownUser",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "UnknownUser"},
					},
					"anonymous": {
						Name:    "anonymous",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Anonymous"},
					},
					"gpName": {
						Name:    "gpName",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "GpName"},
					},
					"index": {
						Name:    "index",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Index"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	dbGroup, err := memdb.NewMemDB(schemaGroup)
	if err != nil {
		panic(err)
	}

	dbUser, err := memdb.NewMemDB(schemaUser)
	if err != nil {
		panic(err)
	}
	// 缓存数据
	res := &sFakers{
		Schema:      schema,
		SchemaGroup: schemaGroup,
		SchemaUser:  schemaUser,
		DBSpace:     db,
		DBGroup:     dbGroup,
		DBUser:      dbUser,
		Cache:       gcache.New(),
	}
	err = res.ConfigSpaces(10)
	if err != nil {
		panic(err)
	}
	err = res.ConfigGroup([]string{"confluence-administrators", "confluence-users", "wwj-test-gp-1", "wwj-test-gp-2", "wwj-test-gp-3"})
	if err != nil {
		panic(err)
	}
	err = res.ConfigUser(10, "confluence-users")
	if err != nil {
		panic(err)
	}
	res.SetCache(context.Background(), "success", true)
	return res
}

func (s *sFakers) ConfigSpaces(total int) error {
	for i := 0; i < total; i++ {
		space := &importserv.ConfluenceSpaceData{
			Name: fmt.Sprintf("SpaceName-%v-%v", time.Now().Unix(), i),
			Key:  fmt.Sprintf("SPACE%d", i),
			ID:   i,
			Type: "global",
		}
		// Create a write transaction
		txn := s.DBSpace.Txn(true)
		if err := txn.Insert("space", space); err != nil {
			fmt.Println("insert space err:", space)
			return err
		}
		// Commit the transaction
		txn.Commit()
		fmt.Println("insert space success:", space)
	}
	return nil
}

func (s *sFakers) ConfigGroup(gpName []string) error {

	for _, v := range gpName {
		group := &importserv.UserGroup{
			Type: "group",
			Name: v,
		}
		// Create a write transaction
		txn := s.DBGroup.Txn(true)
		if err := txn.Insert("group", group); err != nil {
			fmt.Println("insert group err:", group)
			return err
		}
		// Commit the transaction
		txn.Commit()
		fmt.Println("insert group success:", group)

	}
	return nil
}

func (s *sFakers) ConfigUser(total int, inGroup string) error {

	for i := 0; i < total; i++ {
		user := &importserv.UserConfluence{
			Type:        "known",
			Status:      "current",
			Username:    fmt.Sprintf("fake-name%d", i),
			UserKey:     fmt.Sprintf("fake-key%d", i),
			DisplayName: fmt.Sprintf("fake-displayName%d", i),
			FullName:    fmt.Sprintf("fake-displayName%d", i),
			Email:       fmt.Sprintf("fake-email%d@%dem.test", i, i),
			UnknownUser: false,
			Anonymous:   false,
			GpName:      inGroup,
		}
		// Create a write transaction
		txn := s.DBUser.Txn(true)
		if err := txn.Insert("user", user); err != nil {
			fmt.Println("insert user err:", user)
			return err
		}
		// Commit the transaction
		txn.Commit()
		fmt.Println("insert user success:", user)
	}

	return nil
}

// FakeSpaceListBySpaceKeyList 查询Elements列表
func (s *sFakers) GetAllFakeSpace(ctx context.Context, limit int, start int) (out *importserv.FakeConfluenceSpaceResult, err error) {

	out = &importserv.FakeConfluenceSpaceResult{}
	spaceData := make([]*importserv.ConfluenceSpaceData, 0)
	// List all the people
	txn := s.DBSpace.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("space", "id")
	if err != nil {
		panic(err)
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*importserv.ConfluenceSpaceData)
		if p.ID >= start {
			p.Links.Self = "http://59.110.32.216:8000/rest/api/space/" + p.Key
			p.Links.Webui = "/display/" + p.Key
			p.Expandable.Homepage = fmt.Sprintf("/rest/api/content/20231212%d", p.ID)
			spaceData = append(spaceData, p)
		}
	}
	out.Results = spaceData
	out.Limit = limit
	out.Start = start
	out.Size = len(out.Results)
	out.Links.Base = "http://59.110.32.216:8000"
	out.Links.Self = "http://59.110.32.216:8000/rest/api/space"
	return out, err
}

func (s *sFakers) GetAllGrout(ctx context.Context, limit int, start int) (out *importserv.FakeGroupResult, err error) {

	out = &importserv.FakeGroupResult{}
	groupData := make([]*importserv.UserGroup, 0)
	// List all the people
	txn := s.DBGroup.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("group", "id")
	if err != nil {
		panic(err)
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*importserv.UserGroup)
		p.Links.Self = "http://59.110.32.216:8000/rest/api/group/" + p.Name

		if p.Index >= start {
			groupData = append(groupData, p)
		}
	}
	out.Results = groupData
	out.Limit = limit
	out.Start = start
	out.Size = len(out.Results)
	out.Links.Base = "http://59.110.32.216:8000"
	out.Links.Self = "ttp://59.110.32.216:8084/rest/api/group/"
	return out, err
}

func (s *sFakers) GetAllConfulenceUserByGroupName(ctx context.Context, gpName string, limit int, start int) (out *importserv.FakeGroupUserListResp, err error) {

	out = &importserv.FakeGroupUserListResp{}
	userData := make([]*importserv.UserConfluence, 0)
	txn := s.DBUser.Txn(false)
	defer txn.Abort()
	it, err := txn.Get("user", "gpName", gpName)
	if err != nil {
		panic(err)
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*importserv.UserConfluence)
		p.Links.Self = "http://59.110.32.216:8000/rest/api/user?key=" + p.UserKey
		p.ProfilePicture.IsDefault = true

		if p.Index >= start {
			userData = append(userData, p)
		}

	}
	out.Results = userData
	out.Limit = limit
	out.Start = start
	out.Size = len(out.Results)
	out.Links.Base = "http://59.110.32.216:8000"
	out.Links.Self = "http://59.110.32.216:8000/rest/api/group/" + gpName + "/member?expand=status"
	return out, err
}

func (s *sFakers) GetConfulenceUserUserName(ctx context.Context, userName string) (out *importserv.UserConfluence, err error) {
	out = &importserv.UserConfluence{}
	txn := s.DBUser.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("user", "username", userName)
	if err != nil {
		return nil, err
	}
	if raw != nil {
		out = raw.(*importserv.UserConfluence)
	}
	return out, err
}

func (s *sFakers) FakeCheckConfig(ctx context.Context, result consts.SetConfigFakeResult) {
	str := ""
	status := 200
	switch result {
	case consts.SetConfigSuccess:
		str = `{
			"results":[
		
			],
			"start":0,
			"limit":2,
			"size":0,
			"_links":{
				"self":"http://59.110.32.216:8000/rest/api/space",
				"base":"http://59.110.32.216:8000",
				"context":""
			}
		}`
		status = 200
	case consts.SetConfigUnSuccess:
		str = `{"SetConfigUnSuccess":"error"}`
		status = 401
	default:
		status = 500
		str = "<error>"
	}

	g.RequestFromCtx(ctx).Response.WriteStatusExit(status, str)
}

func (s *sFakers) SetCache(ctx context.Context, k interface{}, v interface{}) {
	s.Cache.Set(ctx, k, v, 0)
}

func (s *sFakers) GetCache(ctx context.Context) *gcache.Cache {
	return s.Cache
}

// http://59.110.32.216:8000/rest/api/group/?limit=200&start=0

// FakeSpaceListByTotal 查询Elements列表
func (s *sFakers) FakeSpaceListByTotal(ctx context.Context, total int) (out []*importserv.ConfluenceSpaceListResp, err error) {
	out = make([]*importserv.ConfluenceSpaceListResp, 0)
	for i := 0; i < total; i++ {
		out = append(out, &importserv.ConfluenceSpaceListResp{
			SpaceName: fmt.Sprintf("SpaceName%d", i),
			SpaceKey:  fmt.Sprintf("SPACE%d", i),
			SpaceId:   1,
			Type:      "global",
			Status:    0, //0 未迁移 1迁移过
		})
	}
	return out, nil
}
