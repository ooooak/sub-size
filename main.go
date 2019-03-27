package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const root = "D:\\"
const sp = os.PathSeparator

func main() {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		dir := root + string(sp) + file.Name()
		size := int64(0)

		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				size = size + info.Size()
			}

			return nil
		})

		fmt.Println(file.Name(), size/(1024*1024), "mb")
	}
}
