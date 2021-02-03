package Array

// ArrayFill 用给定的值value填充数组,num为插入元素的数量.
func (*Array)ArrayFill(value interface{}, num int) []interface{} {
	if num <= 0 {
		return nil
	}

	var res = make([]interface{}, num)
	for i := 0; i < num; i++ {
		res[i] = value
	}

	return res
}
