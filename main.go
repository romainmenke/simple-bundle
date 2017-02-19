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

	extensions := make(map[string]string)
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
		if content != "" {
			extensions[extension] = extensions[extension] + content + newLine
		}
	}

	for ext, content := range extensions {
		writeFile(content, ext, outDir)
	}

}

func writeFile(content string, extension string, out string) {
	if content == "" {
		return
	}

	err := ioutil.WriteFile(out+"bundle."+extension, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func readFile(name string) string {
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
	return string(buf.Bytes())
}

const newLine = `
`

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
