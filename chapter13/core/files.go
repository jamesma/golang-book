package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFileDemo() {
	file, err := os.Open("../README.md")
	if err != nil {
		fmt.Println("error in os.Open()", err)
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("error in file.Stat()", err)
		return
	}

	// read the file
	bytes := make([]byte, stat.Size())
	_, err = file.Read(bytes)
	if err != nil {
		fmt.Println("error in stat.Size()", err)
		return
	}

	str := string(bytes)
	fmt.Println(str)

	// reading files is very common, so there's a shorter way to achieve above
	bs, err := ioutil.ReadFile("../README.md")
	if err != nil {
		fmt.Println("error in ioutil.ReadFile()", err)
		return
	}
	str = string(bs)
	fmt.Println(str)
}

func WriteFileDemo() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("error in os.Create()", err)
		return
	}
	defer file.Close()

	file.WriteString("test contents")

	// remove the file just created
	err = os.Remove("test.txt")
	if err != nil {
		fmt.Println("error in os.Remove()", err)
	}
}

func ReadDirDemo() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("error in os.Open()", err)
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("error in dir.Readdir()", err)
		return
	}
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}

func WalkDemo() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
