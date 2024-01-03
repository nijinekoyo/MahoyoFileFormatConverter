/*
 * @Author: nijineko
 * @Date: 2024-01-03 20:46:29
 * @LastEditTime: 2024-01-03 21:21:45
 * @LastEditors: nijineko
 * @Description: main.go
 * @FilePath: \MahoyoFileFormatConverter\main.go
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nijinekoyo/MahoyoFileFormatConverter/File"
	"github.com/nijinekoyo/MahoyoFileFormatConverter/FormatConvert"
)

func main() {
	folderPath := flag.String("i", "", "input folder path")
	DeleteOriginalFiles := flag.Bool("delete_original", false, "delete original files")
	flag.Parse()

	if *folderPath != "" {
		err := convert(*folderPath, *DeleteOriginalFiles)
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	flag.Usage()
}

/**
 * @description: Convert all files in a folder
 * @param {string} folderPath Folder path
 */
func convert(folderPath string, DeleteOriginalFiles bool) error {
	// Get all file paths
	FilePaths, err := File.GetFilePaths(folderPath)
	if err != nil {
		return err
	}

	for _, FilePath := range FilePaths {
		File, err := os.ReadFile(FilePath)
		if err != nil {
			return err
		}

		// Judge file suffix
		switch filepath.Ext(FilePath) {
		case ".hw":
			// Convert .hw to .ogg
			OggFile, err := FormatConvert.HWToOgg(File)
			if err != nil {
				fmt.Println("Convert", FilePath, "failed:", err)
			} else {
				// Write .ogg file
				oggSavePath := FilePath[:len(FilePath)-2] + "ogg"

				err = os.WriteFile(oggSavePath, OggFile, 0666)
				if err != nil {
					fmt.Println("Write", oggSavePath, "failed:", err)
				}

				// Delete original file
				if DeleteOriginalFiles {
					err = os.Remove(FilePath)
					if err != nil {
						fmt.Println("Delete", FilePath, "failed:", err)
					}
				}

				continue
			}
		}
	}

	return nil
}
