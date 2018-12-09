// copyright 2018 de-liKeR, CreatorQsF

package herb

// For application configs, use this struct.
type Config struct {
	BinCachePath string // every binary is stored here with versioning
	LinkPath     string // the latest version of binary is linked to here
}

const (
	_BINCACHEPATH = "~/.herb/"       // user home directory's hidden
	_LINKPATH     = "/usr/local/bin" // path executable dir
)
