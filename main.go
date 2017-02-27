package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	source := flag.String("source", "./", "source directory")
	out := flag.String("out", "./", "output directory")
	flag.Parse()

	sourceDir := strings.TrimSuffix(*source, "/") + "/"
	outDir := strings.TrimSuffix(*out, "/") + "/"
	createIfMissing(outDir)

	exclude := flag.Args()

	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	extensions := make(map[string][]byte)
FILE_ITERATOR:
	for _, f := range files {
		if !isFile(sourceDir + f.Name()) {
			continue FILE_ITERATOR
		}

		for _, exc := range exclude {
			if strings.Contains(f.Name(), exc) {
				continue FILE_ITERATOR
			}
		}

		if strings.Contains(f.Name(), "bundle") {
			continue FILE_ITERATOR
		}

		nameComponents := strings.Split(f.Name(), ".")
		extension := nameComponents[len(nameComponents)-1]

		content := readFile(sourceDir + f.Name())
		if len(content) != 0 {
			content = append(content, 10)
			extensions[extension] = append(extensions[extension], content...)
		}
	}

	for ext, content := range extensions {
		writeFile(content, ext, outDir)
	}

}

func writeFile(content []byte, extension string, out string) {
	if len(content) == 0 {
		return
	}

	err := ioutil.WriteFile(out+"bundle."+extension, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func readFile(name string) []byte {
	buf := bytes.NewBuffer(nil)
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(buf, file)
	if err != nil {
		panic(err)
	}
	file.Close()
	return buf.Bytes()
}

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return !fileInfo.IsDir()
}

func createIfMissing(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
