// copyright 2018 de-liKeR, CreatorQsF

package herb

import "go/build"

func getGOPATH() string {
	return build.Default.GOPATH
}
