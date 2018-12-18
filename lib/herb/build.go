// copyright 2018 de-liKeR, CreatorQsF

package herb

import "go/build"

type BuildOption struct{}

func getGOPATH() string {
	return build.Default.GOPATH
}

func (r *Repository) Build(bo *BuildOption) ([]byte, error) {
	return []byte{}, nil
}
