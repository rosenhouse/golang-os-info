package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	properties := map[string]string{
		"path list separator": string(os.PathListSeparator),
		"path separator":      string(os.PathSeparator),
		"goos":                runtime.GOOS,
		"goarch":              runtime.GOARCH,
		"compiler":            runtime.Compiler,
		"num cpus":            fmt.Sprintf("%d", runtime.NumCPU()),
		"go version":          runtime.Version(),
	}

	for k, v := range properties {
		fmt.Printf("%25s = %s\n", k, v)
	}
}
