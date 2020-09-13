package floatlib

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 保留指定小数
func Decimal(f float64, d int) float64 {

	format := strings.Builder{}
	format.WriteString("%.")
	format.WriteString(strconv.Itoa(d))
	format.WriteString("f")

	f, _ = strconv.ParseFloat(fmt.Sprintf(format.String(), f), 64)
	return f
}

// 保留两位小数
func Decimal2(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

//Float64ToByte Float64转byte
func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func IsEqual(f1, f2 float64) bool {
	const MIN = 0.000001
	return math.Dim(f1, f2) < MIN
}
