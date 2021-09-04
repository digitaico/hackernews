package links

import (
	"fmt"

	"github.com/glyphack/go-graphql-hackernews/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	fmt.Println("vim-go")
}
