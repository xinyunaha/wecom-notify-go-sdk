/**
 * @Author: xinyunaha
 * @File:  io
 * @Date: 2024/1/12 15:38
 * @Version: 1.0.0
 * @Description: 用户文件相关组件
 */

package utils

import "os"

// Exists 判断文件或文件夹是否存在
func Exists(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
