package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// 解析命令行参数
	var dir string
	flag.StringVar(&dir, "d", "", "directory path")
	var pattern string
	flag.StringVar(&pattern, "p", "", "name pattern, eg. newname%d")
	flag.Parse()
	fmt.Sprintln("dir", dir, "pattern", pattern)
	if dir == "" || pattern == "" {
		flag.Usage()
		return
	}
	// 遍历文件夹，获取文件路径
	paths := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	// 遍历文件路径，修改文件名
	for i, path := range paths {
		newPath := filepath.Join(filepath.Dir(path), fmt.Sprintf(pattern, i+1)+filepath.Ext(path))
		os.Rename(path, newPath)
	}
}
