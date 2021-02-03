package Array

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayColumn Array/Global.go Array/ArrayColumn.go Array/ArrayColumn_test.go
func TestArrayColumn(t *testing.T) {
	var toolArray Array
	jsonStr := `[{"name":"zhang3","age":23,"sex":1},{"name":"li4","age":30,"sex":1},{"name":"wang5","age":25,"sex":0},{"name":"zhao6","age":50,"sex":0}]`
	var arr interface{}
	_ = json.Unmarshal([]byte(jsonStr), &arr)
	res1 := toolArray.ArrayColumn(arr, "name")
	fmt.Println(res1)
}

// go test -v -run TestArrayColumn -bench=BenchmarkArrayColumn -count=5 Array/Global.go Array/ArrayColumn.go Array/ArrayColumn_test.go
func BenchmarkArrayColumn(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	jsonStr := `[{"name":"zhang3","age":23,"sex":1},{"name":"li4","age":30,"sex":1},{"name":"wang5","age":25,"sex":0},{"name":"zhao6","age":50,"sex":0}]`
	var arr interface{}
	_ = json.Unmarshal([]byte(jsonStr), arr)
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayColumn(arr, "name")
	}
}