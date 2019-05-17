package social

import (
	"crypto/dsa"
	"time"
)

type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CoverImage  string    `json:"cover_image"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	Username    string    `json:"username"`
}

type Comment struct {
	ID          string    `json:"id"`
	Content     string    `json:"content"`
	PostID      string    `json:"post_id"`
	Username    string    `json:"username"`
	CreatedTime time.Time `json:"created_time"`
}

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	PubKey     dsa.PublicKey
	PrivateKey dsa.PrivateKey
}
