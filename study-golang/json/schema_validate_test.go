package json

import (
	"go/types"
	"log"
	"strconv"
	"testing"
)

func DataType(data interface{}) string {
	switch data.(type) {
	case uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64, uintptr:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "boolean"
	case types.Array, types.Slice:
		return "array"
	case types.Map, types.Struct:
		return "object"
	default:
		return "unknown"
	}
}

type ErrorMessage struct {
	PropertyPath string
	Message      string
}

type Result struct {
	Errors []*ErrorMessage
}

func deleteItem(parentData *interface{}, dataFieldName string, dataArrayIndex int) {
	if dataFieldName == "" {
		pd := (*parentData).([]interface{})
		log.Println(pd)
	} else {
		pd := (*parentData).(map[string]interface{})
		delete(pd, dataFieldName)
	}
}

func validateData(level string, schema map[string]interface{}, data interface{}, parentData *interface{}, dataFieldName string, dataArrayIndex int, r *Result) *Result {
	typeValue := schema["type"].(string)
	required := schema["required"].(bool)
	// log.Printf("当前 level 为 %s ，typeValue 为 %s ，required 为 %v \n", level, typeValue, required)
	if required && data == nil {
		r.Errors = append(r.Errors, &ErrorMessage{PropertyPath: level, Message: level + " 是预置属性，不能为空"})
	}
	if data != nil {
		switch typeValue {
		case "object":
			objectData, ok := data.(map[string]interface{})
			if !ok {
				if required {
					r.Errors = append(r.Errors, &ErrorMessage{PropertyPath: level, Message: level + " 必须是 object 类型"})
				} else {
					deleteItem(parentData, dataFieldName, dataArrayIndex)
				}
			} else {
				propertiesMap := schema["properties"].(map[string]interface{})
				for filedName, settings := range propertiesMap {
					validateData(level+"/"+filedName, settings.(map[string]interface{}), objectData[filedName], &data, filedName, 0, r)
				}
			}
		case "array":
			objectData, ok := data.([]interface{})
			if !ok {
				if required {
					r.Errors = append(r.Errors, &ErrorMessage{PropertyPath: level, Message: level + " 必须是 array 类型"})
				} else {
					deleteItem(parentData, dataFieldName, dataArrayIndex)
				}
			} else {
				itemsMap := schema["items"].(map[string]interface{})
				for index, value := range objectData {
					validateData(level+"/"+strconv.Itoa(index), itemsMap, value, &data, "", index, r)
				}
			}
		default:
			dataType := DataType(data)
			if dataType != typeValue {
				if required {
					r.Errors = append(r.Errors, &ErrorMessage{PropertyPath: level, Message: level + " 是预置属性，期望类型是 " + typeValue + "，实际类型是 " + dataType})
				} else {
					deleteItem(parentData, dataFieldName, dataArrayIndex)
				}
			}
		}
	}
	return r
}

func TestSchemaValidate(t *testing.T) {

	schema := H{
		"type":     "object",
		"required": true,
		"properties": H{
			"#id": H{
				"type":     "integer",
				"required": true,
			},
			"#person": H{
				"type":     "object",
				"required": true,
				"properties": H{
					"#age": H{
						"type":     "integer",
						"required": true,
					},
					"#name": H{
						"type":     "string",
						"required": true,
					},
				},
			},
			"#friends": H{
				"type":     "array",
				"required": true,
				"items": H{
					"type":     "object",
					"required": true,
					"properties": H{
						"#age": H{
							"type":     "integer",
							"required": true,
						},
						"name": H{
							"type":            "string",
							"required":        false,
							"discard_invalid": true,
						},
					},
				},
			},
			"#likes": H{
				"type":     "array",
				"required": true,
				"items": H{
					"type":     "string",
					"required": false,
				},
			},
		},
	}

	value := map[string]interface{}{
		"#id": 1,
		"#person": H{
			"#age":  1,
			"#name": "1",
			"other": 1,
		},
		"#friends": A{
			H{"#age": 1, "name": "1", "other": 1},
			H{"#age": 2, "name": "2", "other": "2"},
		},
		"#likes": A{
			"a",
			"b",
			1,
		},
		"other": 1,
	}

	r := Result{
		Errors: make([]*ErrorMessage, 0, 1),
	}
	validateData("", schema, value, nil, "", 0, &r)

	for _, massage := range r.Errors {
		log.Println(massage)
	}

}
