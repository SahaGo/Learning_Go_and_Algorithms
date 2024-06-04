package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
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

func dirTree(out io.Writer, path string, printFiles bool) error {
	getDirTree(out, "", path, printFiles)
	return nil
}

func getDirTree(out io.Writer, prefix, path string, printFiles bool) error {
	fileObj, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open %s: %s", path, err.Error())
	}
	defer fileObj.Close()

	files, err := fileObj.Readdir(-1)
	if err != nil {
		log.Fatalf("Could not read dir names in %s: %s", path, err.Error())
	}

	// сортировка с приоритетом папок над файлами

	//sort.Slice(files, func(i, j int) bool {
	//	if files[i].IsDir() == files[j].IsDir() {
	//		return files[i].Name() < files[j].Name()
	//	}
	//	return files[i].IsDir() && !files[j].IsDir()
	//})

	// сортировка по заданию, все файлы подряд
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	directories := make([]os.FileInfo, 0)
	if !printFiles {
		for _, file := range files {
			if file.IsDir() {
				directories = append(directories, file)
			}
		}
		files = directories
	}

	l := len(files)
	for i, file := range files {
		if file.IsDir() {
			var prefixNew string
			if l > i+1 {
				fmt.Fprintf(out, prefix+"├───%s\n", file.Name())
				prefixNew = prefix + "│\t"
			} else {
				fmt.Fprintf(out, prefix+"└───%s\n", file.Name())
				prefixNew = prefix + "\t"
			}

			newDir := filepath.Join(path, file.Name())
			getDirTree(out, prefixNew, newDir, printFiles)
		} else if printFiles {
			var size string
			if file.Size() > 0 {
				size = fmt.Sprintf("%vb", file.Size())
			} else {
				size = "empty"
			}
			if l > i+1 {
				fmt.Fprintf(out, prefix+"├───%s (%s)\n", file.Name(), size)
			} else {
				fmt.Fprintf(out, prefix+"└───%s (%s)\n", file.Name(), size)
			}
		}
	}
	return nil
}
