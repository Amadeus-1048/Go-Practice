package error_learn

import "fmt"

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("math: square root of a negative number")
		// 通常构造error时，要求传入的字符串首字母小写，结尾不带标点符号
		// fmt.Errorf 函数 先将字符串格式化，再调用 errors.New 函数来创建错误
		return 0, fmt.Errorf("math: square root of a negative number %g", f)
	}
	return f, nil
}
