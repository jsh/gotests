package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

var args []string
	// args := make([]string, 2)

const (
	bmk     = "./bmk"
	inpath  = "./true.in"
	outpath = "./true.out"
)


func main() {

	expect, err := ioutil.ReadFile(bmk)
	if err != nil {
		fmt.Println("cannot read file ", bmk)
		os.Exit(1)
	}

	program, err := ioutil.ReadFile(inpath)
	if err != nil {
		fmt.Println("cannot read file ", inpath)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(outpath, program, 0777); err != nil {
		fmt.Println("cannot write file ", outpath)
		os.Exit(1)
	}

	cmd := exec.Command(outpath, args...)

	c1 := make(chan string, 1)
	go func() {
		out, err := cmd.CombinedOutput()
		switch {
		case (bytes.Compare(expect, out) != 0) && (err != nil):
			{
				fmt.Printf("bad-out-bad-exit\n")
			}
		case (bytes.Compare(expect, out) != 0):
			{
				fmt.Printf("bad-out-good-exit\n")
			}
		case err != nil:
			{
				fmt.Printf("good-out-bad-exit\n")
			}
		}
		c1 <- "done" // dummy
	}()
	select {
	case <-c1:

	case <-time.After(time.Second * 1):
		{
			fmt.Printf("time-out-no-exit\n")
			cmd2 := exec.Command("/usr/bin/killall", "true.out")
			if err := cmd2.Run(); err != nil {
				fmt.Println("cannot killall ", outpath, " : ", err)
				os.Exit(1)
			}
			cmd2 = exec.Command("rm", "true.out")
			if err := cmd2.Run(); err != nil {
				fmt.Println("cannot remove ", outpath, " : ", err)
				os.Exit(1)
			}
		}
	}
}
