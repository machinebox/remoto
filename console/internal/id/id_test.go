package id_test

import (
	"testing"

	"github.com/machinebox/remoto/console/internal/id"
	"github.com/matryer/is"
)

func Test(t *testing.T) {
	is := is.New(t)
	var i int64 = 5687539843203072
	vid := id.Itoa(i)
	is.Equal(vid, "q32kmp33y")
	is.Equal(id.Atoi(vid), i)
	is.Equal(id.Itoa(0), "")
	is.Equal(id.Itoa(1), "1")
	is.Equal(id.Itoa(10), "a")
	is.Equal(id.Itoa(35), "z")
	is.Equal(id.Itoa(36), "A")
}
