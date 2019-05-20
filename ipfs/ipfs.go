package ipfs

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mattn/go-shellwords"
)

const ipfs = "ipfs"

// Add will call the ipfs Add command
func Add(file string) (string, string) {
	return execCommand(ipfs, "add "+file)
}

// AddDir will call the ipfs add command with the correct arguments for a directory
func AddDir(dir string) (cid string, stderr string) {
	out, err := execCommand(ipfs, "add -r "+dir)
	if err != "" {
		return "", err
	}
	arr := strings.Split(out, "\n")
	arr2 := strings.Split(arr[len(arr)-2], " ")
	cid = arr2[len(arr2)-2]
	return cid, err
}

// Pin will pin the file in IPFS
func Pin(cid string) (string, string) {
	return execCommand(ipfs, "pin add "+cid)
}

// CheckForDir will check ipfs for the dir hash
func CheckForDir(hash string) (string, string) {
	return execCommand(ipfs, "ls "+hash)
}

// CheckForMaster will return the dir name of the hash
func CheckForMaster(hash string) (bool, string) {
	out, err := execCommand(ipfs, "ls "+hash)
	if err != "" {
		return false, err
	}
	masterExists := strings.Contains(out, "master.m3u8")
	return masterExists, ""
}

// CheckForPin will check to see if the current movie is pinned and will retrieve the hash
func CheckForPin(file string) (dirName string, stderr string) {
	out, err := execCommand(ipfs, "pin ls --type='recursive'")
	if err != "" {
		return out, err
	}
	pins := strings.Split(out, "\n")

	if len(pins) <= 1 {
		return "", err
	}
	for _, element := range pins {
		cid := strings.Split(element, " ")[0]
		valid, errDirName := CheckForMaster(cid)
		if errDirName != "" {
			return "", errDirName
		}

		if valid {
			return cid, ""
		}
	}

	return "", err
}

func execCommand(name string, arg string) (stdout string, stderr string) {
	var cmd *exec.Cmd
	var argS []string
	var err error

	argS, err = shellwords.Parse(arg)
	if err != nil {
		return "", fmt.Sprintf("there was an err with args: %s\r\n%s", arg, err.Error())
	}

	cmd = exec.Command(name, argS...)
	out, err := cmd.CombinedOutput()
	stdout = string(out)
	stderr = ""
	if err != nil {
		stderr = err.Error()
	}
	return stdout, stderr
}
