package Array

// ArrayPush 将一个或多个元素压入数组的末尾(入栈),返回处理之后数组的元素个数.
func (*Array)ArrayPush(s *[]interface{}, elements ...interface{}) int {
	*s = append(*s, elements...)
	return len(*s)
}
