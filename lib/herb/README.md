# pkg `herb`

This is the library package for Qs-F/herb command.  

## Instllationh

`go get github.com/Qs-F/herb/lib/herb`

## Design

### Phase 1

Simply, implementing `go get` command and get the binary. However, this will affect your working produts in $GOPATH, thus if you use any of these, it's possible that it can lost the changes. Also, go modules is another problem. If the user use `$GO111MODULE=on`, then `go get` doesn't work if the current directory is not the repository managed by go modules. And if it is, there is another problem because any of these `go get` will be automatically added on current projects' go.mod. This must be avoided.

### Phase 2

Implementing `go get` command inside chroot and get the binary. Inside chroot, all the other repositories (or other required packages, if possible) are introduced. However, if you cannot install commands with `go get` (e.g. private repo), this command cannot provide the way to manage the binary. 
