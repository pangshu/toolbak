package Convert

import "encoding/json"

func (*Convert)JsonToMap(jsonString string) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		return nil, err
	}
	return result, nil
}