package utils

import (
	"strconv"
	"strings"
)


// IntToStr converts integer to string
func IntToStr(num int) string {
	return strconv.Itoa(num)
}

// StrToInt64 converts string to int64
func StrToInt64(s string) int64 {
	int64Val, _ := strconv.ParseInt(s, 10, 64)
	return int64Val
}

// BytesToInt64 converts []bytes to int64
func BytesToInt64(s []byte) int64 {
	int64Val, _ := strconv.ParseInt(string(s), 10, 64)
	return int64Val
}

// StrToUint64 converts string to the unsinged int64
func StrToUint64(s string) uint64 {
	ret, _ := strconv.ParseUint(s, 10, 64)
	return ret
}

// Float64ToStr converts float64 to string
func Float64ToStr(f float64, prec ...string) string {
	return strconv.FormatFloat(f, 'f', 13, 64)
}

// StrToInt converts string to integer
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StringToIntDefault(str string, defVal int) int {
	if in, err := strconv.Atoi(str); err != nil {
		return defVal
	} else {
		return in
	}
}

// StrToFloat64 converts string to float64
func StrToFloat64(s string) float64 {
	Float64, _ := strconv.ParseFloat(s, 64)
	return Float64
}

// BytesToFloat64 converts []byte to float64
func BytesToFloat64(s []byte) float64 {
	Float64, _ := strconv.ParseFloat(string(s), 64)
	return Float64
}

// BytesToInt converts []byte to integer
func BytesToInt(s []byte) int {
	i, _ := strconv.Atoi(string(s))
	return i
}
func StrToBool(s string) bool {
	if s == "0" || s == "false" {
		return false
	}
	return true
}

// SnakeString XxYy to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// CamelCase  very_important_person => VeryImportantPerson
func CamelCase(in string) string {
	tokens := strings.Split(in, "_")
	for i := range tokens {
		tokens[i] = strings.Title(strings.Trim(tokens[i], " "))
	}
	return strings.Join(tokens, "")
}
