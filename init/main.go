package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: init <proc-attr-env_sample FullPath>\n")
		return
	}
	parentPath := os.Args[1]

	os.Clearenv()
	os.Setenv("HOGE", "hoge")
	os.Setenv("HIGE", "hige")

	var procAttr os.ProcAttr
	procAttr.Env = os.Environ()
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}

	path := parentPath + "/child1/child1"
	process1, e := os.StartProcess(path, []string{path, parentPath}, &procAttr)
	if e != nil {
		panic(e)
	}

	path = parentPath + "/child2/child2"
	process2, e := os.StartProcess(path, nil, &procAttr)
	if e != nil {
		panic(e)
	}

	for _, s := range os.Environ() {
		fmt.Printf("init:%s\n", s)
	}

	process1.Wait()
	process2.Wait()
}
