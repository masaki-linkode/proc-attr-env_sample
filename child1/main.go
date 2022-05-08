package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: child1 <proc-attr-env_sample FullPath>\n")
		return
	}
	parentPath := os.Args[1]

	os.Setenv("HOGE", "egoh")
	os.Setenv("HUGE", "huge")

	var procAttr os.ProcAttr
	procAttr.Env = os.Environ()
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}

	path := parentPath + "/child1_1/child1_1"

	process, e := os.StartProcess(path, nil, &procAttr)
	if e != nil {
		panic(e)
	}

	os.Setenv("HIGE", "egih")

	for _, s := range os.Environ() {
		fmt.Printf("child1:%s\n", s)
	}

	process.Wait()
}
