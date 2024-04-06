package cmd

import (
	"os"
	"strings"

	"github.com/gvcgo/goutils/pkgs/gutils"
)

var descardingList []string = []string{
	`?\u001b\\\u001b[6n`,
}

func verify(line string) bool {
	for _, s := range descardingList {
		if strings.Contains(line, s) {
			return false
		}
	}
	return true
}

func FixCast(fPath string) {
	content, _ := os.ReadFile(fPath)
	if len(content) > 0 {
		sList := strings.Split(string(content), "\n")
		data := []string{}
		for _, line := range sList {
			if verify(line) {
				data = append(data, line)
			}
		}
		if len(data) > 0 {
			s := strings.Join(data, "\n")
			os.WriteFile(fPath, []byte(s), os.ModePerm)
		}
	}
}

func FixHeaderForEditOperations(inputFile, outputFile string) {
	ok1, _ := gutils.PathIsExist(inputFile)
	ok2, _ := gutils.PathIsExist(outputFile)
	if ok1 && ok2 {
		content1, _ := os.ReadFile(inputFile)
		sList := strings.SplitN(string(content1), "\n", 2)
		header := sList[0]
		content2, _ := os.ReadFile(outputFile)
		sList = strings.SplitN(string(content2), "\n", 2)
		sList[0] = header
		data := strings.Join(sList, "\n")
		os.WriteFile(outputFile, []byte(data), os.ModePerm)
	}
}
