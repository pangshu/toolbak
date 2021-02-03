package Array

// ArrayPop 弹出数组最后一个元素(出栈),并返回该元素.
func (*Array)ArrayPop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}
