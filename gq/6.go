package gq

import (
	"fmt"
	"strings"
)

/**
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
*/
func main() {
	fmt.Println(convert("LEETCODEISHIRING", 3))
}

/**
按行排序
思路

通过从左向右迭代字符串，我们可以轻松地确定字符位于 Z 字形图案中的哪一行。

算法

我们可以使用 \text{min}( \text{numRows}, \text{len}(s))min(numRows,len(s)) 个列表来表示 Z 字形图案中的非空行。

从左到右迭代 ss，将每个字符添加到合适的行。可以使用当前行和当前方向这两个变量对合适的行进行跟踪。

只有当我们向上移动到最上面的行或向下移动到最下面的行时，当前方向才会发生改变。
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	chars := strings.Split(s, "")
	l := numRows
	if len(s) < numRows {
		l = len(s)
	}
	var rows = make([]string, l)
	curRow := 0
	goDown := false

	for _, char := range chars {
		rows[curRow] = rows[curRow] + char
		if curRow == 0 || curRow == l-1 {
			goDown = !goDown
		}
		if goDown {
			curRow += 1
		} else {
			curRow += -1
		}
	}

	return strings.Join(rows, "")
}
