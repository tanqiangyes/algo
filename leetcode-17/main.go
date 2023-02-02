// Package leetcode_17
/******************************************************************************
 * Author         : jack
 * Mail           : jack.tan@arsyun.com
 * Last modified  : 2023/2/2 15:05
 * Description    :
*******************************************************************************/
package main

func main() {
	letterCombinations("2456")
}

var result []string
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	result = []string{}
	// 从0开始，同时当前得结果为空
	backtrace(digits, 0, "")
	return result
}

func backtrace(digits string, index int, res string) {
	// 终止条件
	if len(digits) == index {
		result = append(result, res)
		return
	}

	// 获取当前index对应的字母
	// 这里没考虑digits不在范围内的情况，因为题目约定
	digit := phoneMap[string(digits[index])]
	// 获取字母个数
	lens := len(digit)
	// 遍历各个字母
	for i := 0; i < lens; i++ {
		backtrace(digits, index+1, res+string(digit[i]))
	}
}
