package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

// RoundedFixed 小数点后 n 位 - 四舍五入
func RoundedFixed(val float64, n int) float64 {
	shift := math.Pow(10, float64(n))
	fv := 0.0000000001 + val //对浮点数产生.xxx999999999 计算不准进行处理
	return math.Floor(fv*shift+.5) / shift
}

// TruncRound 小数点后 n 位 - 舍去
func TruncRound(val float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n+1)+"f", val)
	temp := strings.Split(floatStr, ".")
	var newFloat string
	if len(temp) < 2 || n >= len(temp[1]) {
		newFloat = floatStr
	} else {
		newFloat = temp[0] + "." + temp[1][:n]
	}
	inst, _ := strconv.ParseFloat(newFloat, 64)
	return inst
}

// TimeTransDate 时间戳转换成年月日
func TimeTransDate() string {
	timeLayout := "2006-01-02 15:04:05"
	dateTime := time.Unix(time.Now().Unix(), 0).Format(timeLayout)
	return dateTime
}

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}

	switch value := i.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.Itoa(int(value))
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	default:
		// String conversion using JSON by default
		jsonContent, _ := json.Marshal(value)
		return string(jsonContent)
	}
}
