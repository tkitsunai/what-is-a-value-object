package vo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tkitsunai/what-is-a-value-object/vo"
	"testing"
)

func TestName_二つのオブジェクトは同一である(t *testing.T) {
	name1 := vo.NewName("tkitsunai")
	name2 := vo.NewName("tkitsunai")

	if name1 != name2 {
		t.Fatal(t)
	}

	assert.Equal(t, name1, name2)
}

func TestName_二つのオブジェクトは同一ではない(t *testing.T) {
	name1 := vo.NewName("tkitsunai")
	name2 := vo.NewName("kitsunait")
	assert.NotEqual(t, name1, name2)
}
