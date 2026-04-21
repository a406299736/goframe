package user_func

import (
	"testing"

	"github.com/a406299736/goframe/internal/repository/dbrepo/user"
)

func TestCreate(t *testing.T) {
	u := user.NewModel()
	u.Name = "晓诗"
	id, err := Create(u)
	t.Log(id, err)
}

func TestUpdateById(t *testing.T) {
	_ = UpdateById(1, map[string]any{
		"name": "爱晓诗",
	})
}
