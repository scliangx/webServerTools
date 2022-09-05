package utils

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

func UUID() string {
	return uuid.NewV4().String()
}

// RandString random string
func RandString(n int) string {
	b := make([]rune, n)
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// PathExists 文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// StringsJoin string array join
func StringsJoin(strs ...string) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		b.WriteString(strs[i])
	}
	str = b.String()
	return str

}

// GetAllFile 根据前缀获取文件
func GetAllFile(pathname string, suffix string) (fileSlice []string) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return
	}

	for _, fi := range rd {
		if fi.IsDir() {
			continue
			//GetAllFile(path.Join(pathname, fi.Name()))
		} else {
			if suffix != "" {
				if strings.HasSuffix(fi.Name(), suffix) {
					fileSlice = append(fileSlice, fi.Name())
				}
			} else {
				fileSlice = append(fileSlice, fi.Name())

			}
		}
	}
	return
}
