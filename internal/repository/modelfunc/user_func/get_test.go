package user_func

import (
	"testing"
)

func TestOne(t *testing.T) {
	user, err := One(1)
	t.Log(user.Id, user.Name, err)
}
