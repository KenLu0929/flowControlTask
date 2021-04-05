package migrations

import (
	"github.com/KenLu0929/flowControlTask/config/db"
	"github.com/KenLu0929/flowControlTask/pkg/models/post"
	"github.com/KenLu0929/flowControlTask/pkg/models/user"
)

func init() {
	db.DB.AutoMigrate(&user.User{})
	db.DB.AutoMigrate(&post.Post{})
}
