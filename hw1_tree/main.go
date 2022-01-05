package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func init1() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, files bool) error {
	return dirTreeNested(out, path, files, "")
}

func dirTreeNested(out io.Writer, path string, files bool, suffix string) error {
	fs, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprint(out, "error")
		return err
	}
	if !files {
		newFs := make([]os.DirEntry, 0)
		for _, elem := range fs {
			if elem.IsDir() {
				newFs = append(newFs, elem)
			}
		}
		fs = newFs
	}
	var newSuffix string
	for idx, elem := range fs {
		fmt.Fprint(out, suffix)
		if idx+1 == len(fs) {
			newSuffix = suffix + "\t"
			fmt.Fprint(out, "└───")
		} else {
			newSuffix = suffix + "│\t"
			fmt.Fprint(out, "├───")
		}
		fmt.Fprint(out, elem.Name())
		if elem.IsDir() {
			fmt.Fprint(out, "\n")
			err = dirTreeNested(out, path+"/"+elem.Name(), files, newSuffix)
			if err != nil {
				return err
			}
		} else {
			fi, err := os.Stat(path + "/" + elem.Name())
			if err != nil {
				return err
			}
			fileSize := strconv.Itoa(int(fi.Size()))
			if fileSize == "0" {
				fileSize = "empty"
			} else {
				fileSize += "b"
			}
			fmt.Fprint(out, " (", fileSize, ")\n")
		}
	}

	return nil
}

func main() {
	init1()
}
