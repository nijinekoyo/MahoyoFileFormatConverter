/*
 * @Author: nijineko
 * @Date: 2023-04-20 21:02:36
 * @LastEditTime: 2024-01-03 20:55:23
 * @LastEditors: nijineko
 * @Description: Folder related operations
 * @FilePath: \MahoyoFileFormatConverter\File\Path.go
 */
package File

import (
	"os"
	"path/filepath"
)

/**
 * @description: Traverse all file paths in a folder
 * @param {string} DirPth Folder path
 * @return {[]string} File paths
 * @return {error} Error
 */
func GetFilePaths(DirPth string) ([]string, error) {
	DirPth = filepath.Clean(DirPth)
	var Dirs []string
	Dir, err := os.ReadDir(DirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	var Files []string
	for _, fi := range Dir {
		if fi.IsDir() {
			Dirs = append(Dirs, filepath.Clean(DirPth+PthSep+fi.Name()))
		} else {
			Files = append(Files, filepath.Clean(DirPth+PthSep+fi.Name()))
		}
	}

	for _, Table := range Dirs {
		TempFiles, _ := GetFilePaths(Table)
		for _, TempFile := range TempFiles {
			Files = append(Files, filepath.Clean(TempFile))
		}
	}

	return Files, nil
}
