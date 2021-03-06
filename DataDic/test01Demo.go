package main

import (
	"fmt"
	"io/ioutil"
)

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("读取文件失败:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}


func main() {
	//遍历打印所有的文件名
	var s []string
	s, _ = GetAllFile("F:/IdeaProjects/b2c-server", s)
	fmt.Printf("slice: %v", s)
}