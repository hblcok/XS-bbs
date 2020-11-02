package dao

import (
	"gorm.io/gorm"
	"xs.bbs/internal/app/post/model"
)

var _ IPostDao = (*postDao)(nil)

//var PostDaoSet = wire.NewSet(
//	new(postDao), "*",
//	wire.Bind(new(IPostDao), new(*postDao)),
//)

type (
	PostModel = model.Post
	postDao   struct {
		db *gorm.DB
	}

	IPostDao interface {
		Create(post *PostModel) error
	}
)

func NewPostDao(db *gorm.DB) IPostDao {
	return &postDao{db: db}
}