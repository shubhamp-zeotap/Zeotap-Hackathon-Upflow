package utils

import (
	"strconv"
	"os/exec"
)

func StrToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}


func RunCommand(cmd string) (out []byte, err error) {
	out, err = exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		panic(err)
	}
	return
}