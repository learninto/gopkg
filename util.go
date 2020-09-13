package gopkg

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

func RowsToStruct(rows *sql.Rows) (message json.RawMessage, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return
	}

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	var mapList []map[string]interface{}
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			continue
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			v := values[i]

			b, ok := v.([]byte)
			if ok {
				entry[col] = string(b)
			} else {
				entry[col] = v
			}
		}
		mapList = append(mapList, entry)
	}
	// 序列化数据
	message, err = json.Marshal(mapList)

	return
}

// RowToStruct sql.Row 转 结构体
func RowToStruct(bean interface{}, row *sql.Row) (err error) {
	var columns []string

	t := reflect.TypeOf(bean).Elem()
	for i := 0; i < t.NumField(); i++ {
		jStr := t.Field(i).Tag.Get("json")
		if jStr == "-" {
			continue
		}
		columns = append(columns, jStr)
	}

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	if err = row.Scan(scanArgs...); err != nil {
		return
	}

	//result := make(map[string][]byte)
	entry := make(map[string]interface{})
	for i, col := range columns {
		v := values[i]

		b, ok := v.([]byte)
		if ok {
			entry[col] = string(b)
		} else {
			entry[col] = v
		}

		//switch reflect.TypeOf(v).Kind() {
		//case :
		//
		//}

		//pp.Println(reflect.TypeOf(v), col, string(reflect.ValueOf(v).Bytes()))

	}

	//pp.Println(entry)

	message, err := json.Marshal(entry)
	return json.Unmarshal(message, bean)
}

// FindJsonTag 查询json标签
func FindJsonTag(param interface{}) (res []string) {
	t := reflect.TypeOf(param).Elem()

	for i := 0; i < t.NumField(); i++ {
		jStr := t.Field(i).Tag.Get("json")
		if jStr == "-" {
			continue
		}
		res = append(res, jStr)
	}

	return res
}

func JsonToStruct(bean interface{}, body []byte) (err error) {
	dataStruct := reflect.Indirect(reflect.ValueOf(bean))
	err = json.Unmarshal(body, &dataStruct)
	if err != nil {
		return
	}
	return json.Unmarshal(body, bean)
}
