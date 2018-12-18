// copyright 2018 de-liKeR, CreatorQsF

package herb

import "testing"

func TestBuild(t *testing.T) {
	testquery := "github.com/golang/example/hello"
	repo, err := Get(testquery)
	b, err := repo.Build(nil)
	t.Log(b, err)
}
