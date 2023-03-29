package work

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/require"
	"go/types"
	"log"
	"strconv"
	"strings"
	json2 "techkuze.com/bigdata/study/study-golang/json"
	"testing"
)

type TypeDecl struct {
	Name     string     `json:"name,omitempty"`
	Type     string     `json:"type"`
	Required bool       `json:"required,omitempty"`
	Fields   []TypeDecl `json:"fields,omitempty"`
	Items    *TypeDecl  `json:"items,omitempty"`
}

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
	Errors  []*ErrorMessage
	Deletes []string
}

func validateData(level string, typeDecl TypeDecl, data interface{}, r *Result) {
	typeValue := typeDecl.Type
	required := typeDecl.Required
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
					r.Deletes = append(r.Deletes, level)
				}
			} else {
				for _, fieldObj := range typeDecl.Fields {
					validateData(level+"/"+fieldObj.Name, fieldObj, objectData[fieldObj.Name], r)
				}
			}
		case "array":
			objectData, ok := data.([]interface{})
			if !ok {
				if required {
					r.Errors = append(r.Errors, &ErrorMessage{PropertyPath: level, Message: level + " 必须是 array 类型"})
				} else {
					r.Deletes = append(r.Deletes, level)
				}
			} else {
				for index, value := range objectData {
					validateData(level+"/["+strconv.Itoa(index)+"]", *typeDecl.Items, value, r)
				}
			}
		default:
			dataType := DataType(data)
			if dataType != typeValue {
				if required {
					r.Errors = append(r.Errors, &ErrorMessage{PropertyPath: level, Message: level + " 是预置属性，期望类型是 " + typeValue + "，实际类型是 " + dataType})
				} else {
					r.Deletes = append(r.Deletes, level)
				}
			}

		}
	}
}

func TestSchemaValidate(t *testing.T) {

	assertions := require.New(t)

	var schemaBytes = []byte(`{
		"name": "test",
		"type": "object",
		"required": true,
		"fields": [
			{
				"name": "#id",
				"type": "integer",
				"required": true
			},
			{
				"name": "#person",
				"type": "object",
				"required": true,
				"fields": [
					{
						"name": "#age",
						"type": "integer",
						"required": true
					},
					{
						"name": "#name",
						"type": "string",
						"required": true
					},
					{
						"name": "address",
						"type": "string",
						"required": false
					}
				]
			},
			{
				"name": "#friends",
				"type": "array",
				"required": true,
				"items": {
					"type": "object",
					"required": true,
					"fields": [
						{
							"name": "#age",
							"type": "integer",
							"required": true
						},
						{
							"name": "name",
							"type": "string",
							"required": false
						}
					]
				}
			},
			{
				"name": "#likes",
				"type": "array",
				"required": true,
				"items": {
					"type": "string",
					"required": false
				}
			}
		]
	}`)

	var typeDecl = TypeDecl{}
	err := json.Unmarshal(schemaBytes, &typeDecl)
	assertions.Nil(err)

	var dataBytes = []byte(`{
		"#id": 1,
		"#person": {
			"#age": 1,
			"#name": "1",
			"address": 1,
			"other": 1
		},
		"#friends": [
			{
				"#age": 1,
				"name": "1",
				"other": 1
			},
			{
				"#age": 2,
				"name": "2",
				"other": "2"
			}
		],
		"#likes": [
			"a",
			"b",
			1
		],
		"other": 1
	}`)

	value, err := json2.DecodeJsonBytes(dataBytes)
	assertions.Nil(err)

	r := Result{
		Errors:  make([]*ErrorMessage, 0, 1),
		Deletes: make([]string, 0, 1),
	}

	validateData("", typeDecl, value, &r)

	for _, massage := range r.Errors {
		log.Println(*massage)
	}

	result := dataBytes
	for _, del := range r.Deletes {
		log.Println(del)
		split := strings.Split(del, "/")[1:]
		result = jsonparser.Delete(result, split...)
	}

	log.Println(string(result))
}
