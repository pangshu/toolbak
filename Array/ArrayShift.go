package Array

// ArrayShift 将数组开头的元素移出数组,并返回该元素.
func (*Array)ArrayShift(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	e := (*s)[0]
	*s = (*s)[1:]
	return e
}
