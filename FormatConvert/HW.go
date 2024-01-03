/*
 * @Author: nijineko
 * @Date: 2024-01-03 20:47:18
 * @LastEditTime: 2024-01-03 21:17:42
 * @LastEditors: nijineko
 * @Description: .hw file format converter
 * @FilePath: \MahoyoFileFormatConverter\FormatConvert\HW.go
 */
package FormatConvert

import "errors"

/**
 * @description: .hw to .ogg
 * @param {[]byte} Bytes .hw file bytes
 * @return {[]byte} .ogg file bytes
 */
func HWToOgg(Bytes []byte) ([]byte, error) {
	// Check file header
	if Bytes[4] != 0x68 || Bytes[5] != 0x77 {
		return nil, errors.New("invalid file header")
	}

	// Get offset
	var Offset int
	for i := 0; i < len(Bytes); i++ {
		if Bytes[i] == 0x4F && Bytes[i+1] == 0x67 && Bytes[i+2] == 0x67 && Bytes[i+3] == 0x53 {
			Offset = i
			break
		}
	}

	// remove header
	Bytes = Bytes[Offset:]

	return Bytes, nil
}
