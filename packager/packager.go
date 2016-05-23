package packager

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

// The content that will be packaged
var Content map[string][]byte = make(map[string][]byte)

func Create(path, output string) {
	//Test := []byte("Install gentoo")
	Traverse(path)
	Write(output)
}

// Traverse the content of dirname recursively and add every file in the
// Content map
func Traverse(dirname string) {
	content, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("Error opening " + dirname + " ...")
	}
	for _, f := range content {
		if !strings.HasPrefix(f.Name(), ".") {
			if f.IsDir() {
				Traverse(dirname + f.Name())
			} else {
				content, err := ioutil.ReadFile(dirname + f.Name())
				if err != nil {
					fmt.Println("Error while trying to open " +
						dirname + f.Name() + " (" + err.Error() + ").")
				} else {
					Content[dirname+f.Name()] = content
					fmt.Println("File " + dirname + f.Name() +
						" successfully packaged.")
				}
			}
		}
	}
}

// Write the Content map to a file
func Write(path string) {
	fmt.Println(Content)
	filecontent := []byte(
		"package main\n\nvar SpacContent = map[string][]byte = {\n")
	for key, content := range Content {
		_ = content
		fmt.Println(key)
		filecontent = append(filecontent,
			[]byte("    \""+key+"\": []byte(\"")...)

		content = bytes.Replace(content, []byte("\\"), []byte("\\\\"), -1)
		content = bytes.Replace(content, []byte("\""), []byte("\\\""), -1)

		filecontent = append(filecontent, content...)

		filecontent = append(filecontent, []byte("\"),")...)
	}
	filecontent = append(filecontent, []byte("]")...)
	ioutil.WriteFile(path, filecontent, 0644)
}
