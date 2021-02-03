package Convert

func (*Convert)ToInterface(str []string) []interface{} {
	newInterface := make([]interface{}, len(str))
	for i, v := range str {
		newInterface[i] = v
	}

	return newInterface
}
