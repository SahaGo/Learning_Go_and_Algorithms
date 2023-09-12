package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func mywalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Printf("%d %s; \n", len(info.Name())+1, info.Name())
	return nil

}

func main() {
	const root = "/Users/kisats/projects/stepik_3.5.12/main"
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
}
