package gpslib

import (
	"go-common/pkg/floatlib"
	"math"
)

/**
高德坐标（火星坐标）：gcj02
         百度坐标 ：bd09
*/

const pi = 3.1415926535897932384626
const xPi = 3.14159265358979324 * 3000.0 / 180.0
const a = 6378245.0
const ee = 0.00669342162296594323

// https://blog.csdn.net/zxt94/article/details/79657780

func transformLat(x, y float64) float64 {
	ret := -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*pi) + 20.0*math.Sin(2.0*x*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*pi) + 40.0*math.Sin(y/3.0*pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*pi) + 320*math.Sin(y*pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLon(x, y float64) float64 {
	ret := 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*pi) + 20.0*math.Sin(2.0*x*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*pi) + 40.0*math.Sin(x/3.0*pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*pi) + 300.0*math.Sin(x/30.0*pi)) * 2.0 / 3.0
	return ret
}

func outOfChina(lat, lon float64) bool {
	if lon < 72.004 || lon > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}

/**
 * 84 to 火星坐标系 (GCJ-02) World Geodetic System ==> Mars Geodetic System
 *
 * @param lat
 * @param lon
 * @return
 */
func Gps84ToGcj02(lat, lon float64) map[int64]float64 {
	latLon := make(map[int64]float64)
	if outOfChina(lat, lon) {
		latLon[0] = lat
		latLon[1] = lon
		return latLon
	}
	dLat := transformLat(lon-105.0, lat-35.0)
	dLon := transformLon(lon-105.0, lat-35.0)
	radLat := lat / 180.0 * pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi)
	dLon = (dLon * 180.0) / (a / sqrtMagic * math.Cos(radLat) * pi)
	mgLat := lat + dLat
	mgLon := lon + dLon
	latLon[0] = mgLat
	latLon[1] = mgLon
	return latLon
}

/**
 * * 火星坐标系 (GCJ-02) to 84 * * @param lon * @param lat * @return
 * */
func Gcj02ToGps84(lat, lon float64) map[int64]float64 {
	latLon := make(map[int64]float64)
	gps := transform(lat, lon)
	lontitude := lon*2 - gps[1]
	latitude := lat*2 - gps[0]

	latLon[0] = latitude
	latLon[1] = lontitude
	return latLon
}

/**
 * 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 将 GCJ-02 坐标转换成 BD-09 坐标
 *
 * @param lat 纬度
 * @param lon 经度
 */
func Gcj02ToBd09(lat, lon float64) map[int64]float64 {
	latLon := make(map[int64]float64)
	x := lat
	y := lon
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*xPi)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*xPi)
	tempLon := z*math.Cos(theta) + 0.0065
	tempLat := z*math.Sin(theta) + 0.006

	latLon[0] = tempLat
	latLon[1] = tempLon
	return latLon
}

/**
 * * 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 * * 将 BD-09 坐标转换成GCJ-02 坐标 * * @param
 * bd_lat * @param bd_lon * @return
 */
func Bd09ToGcj02(lat, lon float64) map[int64]float64 {
	latLon := make(map[int64]float64)
	x := lon - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xPi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*xPi)
	tempLon := z * math.Cos(theta)
	tempLat := z * math.Sin(theta)
	latLon[0] = tempLat
	latLon[1] = tempLon
	return latLon
}

/**将gps84转为bd09
 * @param lat
 * @param lon
 * @return
 */
func Gps84Tobd09(lat, lon float64) map[int64]float64 {
	gcj02 := Gps84ToGcj02(lat, lon)
	bd09 := Gcj02ToBd09(gcj02[0], gcj02[1])
	return bd09
}
func Gd09Togps84(lat, lon float64) map[int64]float64 {
	gcj02 := Bd09ToGcj02(lat, lon)
	gps84 := Gcj02ToGps84(gcj02[0], gcj02[1])
	//保留小数点后六位
	gps84[0] = floatlib.Decimal(gps84[0], 6)
	gps84[1] = floatlib.Decimal(gps84[1], 6)
	return gps84
}
