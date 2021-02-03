package Array

// ArrayUnshift 在数组开头插入一个或多个元素,返回处理之后数组的元素个数.
func (*Array)ArrayUnshift(s *[]interface{}, elements ...interface{}) int {
	*s = append(elements, *s...)
	return len(*s)
}
