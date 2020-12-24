package utils

import "encoding/base64"

/*
 *对msg数据进行
 */
func Base64Str(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}
