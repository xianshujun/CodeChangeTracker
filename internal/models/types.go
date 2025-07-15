package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// StringArray 自定义字符串数组类型，用于存储JSON数组到数据库
type StringArray []string

// Value 实现driver.Valuer接口
func (s StringArray) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	return json.Marshal(s)
}

// Scan 实现sql.Scanner接口
func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = StringArray{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	default:
		return fmt.Errorf("cannot scan %T into StringArray", value)
	}
}