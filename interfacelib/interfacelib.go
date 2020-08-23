package interfacelib

import (
	"strconv"
)

func ToFloat64(params interface{}) (res float64) {
	switch params.(type) {

	case int64:
		t, _ := params.(int64)
		res = float64(t)

	case int32:
		t, _ := params.(int32)
		res = float64(t)

	case int16:
		t, _ := params.(int16)
		res = float64(t)

	case int8:
		t, _ := params.(int8)
		res = float64(t)

	case int:
		t, _ := params.(int)
		res = float64(t)

	case uint:
		t, _ := params.(uint)
		res = float64(t)

	case float64:
		res, _ = params.(float64)

	case float32:
		t, _ := params.(float32)
		res = float64(t)

	case string:
		str, _ := params.(string)
		res, _ = strconv.ParseFloat(str, 64)

	default:
		res = 0
	}

	return res
}

func ToInt64(params interface{}) (res int64) {
	switch params.(type) {
	case int64:
		res, _ = params.(int64)

	case int32:
		t, _ := params.(int32)
		res = int64(t)

	case int16:
		t, _ := params.(int16)
		res = int64(t)

	case int8:
		t, _ := params.(int8)
		res = int64(t)

	case int:
		t, _ := params.(int)
		res = int64(t)

	case uint:
		t, _ := params.(uint)
		res = int64(t)

	case float64:
		t, _ := params.(float64)
		res = int64(t)

	case float32:
		t, _ := params.(float32)
		res = int64(t)

	case string:
		t, _ := params.(string)
		res, _ = strconv.ParseInt(t, 10, 64)

	default:
		res = 0
	}

	return res
}

func ToString(params interface{}) (str string) {
	switch params.(type) {
	case int64:
		t, _ := params.(int64)
		str = strconv.FormatInt(int64(t), 10)

	case int32:
		t, _ := params.(int32)
		str = strconv.FormatInt(int64(t), 10)

	case int16:
		t, _ := params.(int16)
		str = strconv.FormatInt(int64(t), 10)

	case int8:
		t, _ := params.(int8)
		str = strconv.FormatInt(int64(t), 10)

	case int: // S
		t, _ := params.(int)
		str = strconv.FormatInt(int64(t), 10)

	case uint:
		t, _ := params.(uint)
		str = strconv.FormatInt(int64(t), 10)

	case []uint8:
		t, _ := params.([]uint8)
		var ba []byte
		for _, b := range t {
			ba = append(ba, byte(b))
		}
		str = string(ba)

	case float64:
		t, _ := params.(float64)
		str = strconv.FormatFloat(t, 'E', -1, 64)

	case float32:
		t, _ := params.(float32)
		str = strconv.FormatFloat(float64(t), 'E', -1, 64)

	case string:
		str, _ = params.(string)
	default:
		str = ""
	}

	return str
}

func StringToFloat64(params interface{}) float64 {
	fv, _ := params.(float64)
	return fv
}

func StringToInt64(params interface{}) int64 {
	fv, _ := strconv.ParseInt(ToString(params), 10, 64)
	return fv
}
