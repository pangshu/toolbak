package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGeoDistance Number/Global.go Number/GeoDistance.go Number/GeoDistance_test.go
func TestGeoDistance(t *testing.T) {
	var toolNumber Number
	var lat1, lng1, lat2, lng2 float64
	lat1, lng1 = 30.0, 45.0
	lat2, lng2 = 40.0, 90.0
	res1 := toolNumber.GeoDistance(lng1,lat1,lng2,lat2)
	fmt.Println(res1)
}

// go test -v -run TestGeoDistance -bench=BenchmarkGeoDistance -count=5 Number/Global.go Number/GeoDistance.go Number/GeoDistance_test.go
func BenchmarkGeoDistance(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	var lat1, lng1, lat2, lng2 float64
	lat1, lng1 = 30.0, 45.0
	lat2, lng2 = 40.0, 90.0
	for i:=0; i< t.N; i++ {
		_ = toolNumber.GeoDistance(lng1,lat1,lng2,lat2)
	}
}