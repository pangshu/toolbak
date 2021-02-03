package Number

import "math"

// GeoDistance 获取地理距离/米.
// 参数分别为两点的经度和纬度.lat:-90~90,lng:-180~180.
func (*Number)GeoDistance(lng1, lat1, lng2, lat2 float64) float64 {
	//地球半径
	radius := 6371000.0
	rad := math.Pi / 180.0

	lng1 = lng1 * rad
	lat1 = lat1 * rad
	lng2 = lng2 * rad
	lat2 = lat2 * rad
	theta := lng2 - lng1

	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}
