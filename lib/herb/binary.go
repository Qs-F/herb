// copyright 2018 de-liKeR, CreatorQsF

package herb

import (
	"hash"
)

// Binary struct provides binary info
type Binary struct {
	Name    string    // file name
	Repo    string    // repository info
	Version string    // binary version
	Hash    hash.Hash // hash of binary file
}

func NewBinary(name, repo string) {
}
