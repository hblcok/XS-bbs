package dao

import (
	"gorm.io/gorm"
	"xs.bbs/internal/app/community/model"
)

var _ ICommunityDao = (*CommunityDao)(nil)

// CommunityDaoSet CommunityDao依赖注入
//var CommunityDaoSet = wire.NewSet(
//	wire.Struct(new(CommunityDao), "*"),
//	wire.Bind(new(ICommunityDao), new(*CommunityDao)))

type (
	CommunityModel = model.Community
	CommunityDao   struct {
		db *gorm.DB
	}

	// ICommunityDao 文章标签接口
	ICommunityDao interface {
		GetCommunityList() ([]CommunityModel, error)
		GetCommunityDetailByID(int64) (*CommunityModel, error)
	}
)

func NewCommunityDao(db *gorm.DB) ICommunityDao {
	return &CommunityDao{db: db}
}
