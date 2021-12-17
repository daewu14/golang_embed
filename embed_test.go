package belajar_embed

import (
	"embed"
	//_ "embed" // karena embed dipanggil maka tanda _ tidak diperlukan
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte
func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS
func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS
func TestPathMatcher(t *testing.T) {
	dir, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}
	for _, entry := range dir {
		entryName := entry.Name()
		file, err := path.ReadFile("files/"+entryName)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(file))
	}
}