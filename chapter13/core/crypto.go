package core

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error in ioutil.ReadFile()", err)
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func Crc32Demo() {
	h1, err := getHash("resources/test1.txt")
	if err != nil {
		fmt.Println("error in getHash()", err)
		return
	}
	h2, err := getHash("resources/test2.txt")
	if err != nil {
		fmt.Println("error in getHash()", err)
		return
	}
	fmt.Println("test1.txt:", h1, "test2.txt:", h2, h1 == h2)
}

func Sha1Demo() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println("sha1:", bs)
}
