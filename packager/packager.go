package packager

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var Content map[string][]byte = make(map[string][]byte)

func Create(path, output string) {
	Traverse(path)
	Write()
}

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
					fmt.Println("Error while trying to open " + dirname + f.Name() + " (" + err.Error() + ").")
				} else {
					Content[dirname+f.Name()] = content
					fmt.Println("File " + dirname + f.Name() + " successfully packaged.")
				}
			}
		}
	}
}

func Write() {
	fmt.Println(Content)
}
