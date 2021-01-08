package main

import (
	"fmt"
	"os/exec"
)

func main() {
	t, err := exec.Command("ascii-art\\run.bat").CombinedOutput()
	fmt.Println(string(t))
	fmt.Println(err)
}
