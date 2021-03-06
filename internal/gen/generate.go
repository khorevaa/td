package gen

// This file defines how to generate templates and example
// generated files.

// Templates should be first.
//go:generate go run github.com/go-bindata/go-bindata/go-bindata -pkg=internal -o=internal/bindata.go -mode=420 -modtime=1 ./_template/...

//go:generate go run github.com/gotd/td/cmd/gotdgen --doc "https://localhost:80/doc" --clean --package td --target example --schema _testdata/example.tl
