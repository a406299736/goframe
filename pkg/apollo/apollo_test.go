package apollo

import (
	"testing"
)

func TestApollo(t *testing.T) {
	if err := CheckStart(); err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}
