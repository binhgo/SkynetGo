package service

import "time"

type PostArgs struct {
	Page   int
	UserID int
}

type PostRPC struct {
}

type Post struct {
	Title      string
	Content    string
	UpdateTime time.Time
}

type PostResult struct {
	Posts []*Post
}

func (t *PostRPC) GetPosts(args PostArgs, result *PostResult) error {

	p1 := &Post{"Hello World!", "This is content...", time.Now()}
	p2 := &Post{"Hello Binh!", "This is content...", time.Now()}
	p3 := &Post{"Hello Candy!", "This is content...", time.Now()}

	pr := &PostResult{}
	pr.Posts = append(pr.Posts, p1, p2, p3)

	(*result).Posts = pr.Posts

	return nil
}
