package packager

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	if !strings.HasSuffix(dirname, string(os.PathSeparator)) {
		dirname += string(os.PathSeparator)
	}
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
	//fmt.Println(Content)
	os.Mkdir(path, 0755)
	filecontent := []byte(
		"package " + path + "\n\nvar Content = map[string][]byte {\n")
	for key, content := range Content {
		_ = content
		//fmt.Println(key)
		filecontent = append(filecontent,
			[]byte("    \""+key+"\": []byte(")...)

		content = []byte(strconv.Quote(string(content)))

		filecontent = append(filecontent, content...)
		filecontent = append(filecontent, []byte("),\n")...)
	}
	filecontent = append(filecontent, []byte("}")...)
	ioutil.WriteFile(path+string(os.PathSeparator)+path+".go", filecontent, 0644)
}
