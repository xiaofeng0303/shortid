package shortid_test

import (
	"huanghe/third_party/shortid"
	"testing"
)

func Test_GenerateID(t *testing.T) {
	prefix := "wan"
	id := shortid.GenerateID(prefix)
	t.Log(id)
}
