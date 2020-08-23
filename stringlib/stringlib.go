package stringlib

import (
	"strconv"
	"strings"
)

/**
匈牙利命名
开头字母用变量类型的缩写，其余部分用变量的英文或英文的缩写，要求单词第一个字母大写。

ex:
int iMyAge; “i”是int类型的缩写
char cMyName[10]; “c”是char类型的缩写
float fManHeight; “f”是flota类型的缩写
*/
func ToHungary() {

}

/**
驼峰式命名法
驼峰式命名又叫小驼峰命名法。第一个单词字母小写，后面其他单词首字母大写。

ex:
int myAge;
char myName[10];
float manHeight;
*/
func ToCamel() {

}

/**
帕斯卡命名法又叫大驼峰命名法。每个单词的第一个字母都大写。

ex:
int MyAge;
char MyName[10];
float ManHeight;
*/
func ToPascal(name string) (res string) {
	if len(name) == 0 {
		return
	}

	result := strings.Builder{}
	for _, value := range strings.Split(name, "_") {
		result.WriteString(Capitalize(value))
	}
	return result.String()
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				//fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// SnakeString     XxYy to xx_yy , XxYY to xx_yy
// 驼峰字符串 转换为 蛇形字符串，例如: XxYy to xx_yy , XxYY to xx_yy
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

// camel string, xx_yy to XxYy
// 蛇形字转换为驼峰字符串
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// ToInt
func ToInt(str string) (i int, e error) {

	i, e = strconv.Atoi(str)
	//#string到int64
	//int64, err := strconv.ParseInt(string, 10, 64)
	return
}

// ToInt64
func ToInt64(str string) (i int64, e error) {
	i, e = strconv.ParseInt(str, 10, 64)
	return
}

// ToInt64
func ToFloat64(str string) (i float64, e error) {
	i, e = strconv.ParseFloat(str, 64)
	return
}
