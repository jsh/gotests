package main

import (

	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"bytes"
	)

func main() {
	cmd := exec.Command("/bin/echo", "hello, world")
	out,err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Whoops")
		os.Exit(1)
	}
	expect, err := ioutil.ReadFile("./bmk")
	if err != nil {
		os.Exit(1)
	}
	if bytes.Compare(expect,out) == 0 {
		fmt.Println("Same")
	} else {
		fmt.Println("Differ")
	}
}
