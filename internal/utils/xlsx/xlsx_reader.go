package xlsx

import (
	"reflect"

	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

var (
	TagName       = "xlsx"
	SkipHeaderNum = 2 // 表头需要跳过的行数
	KeysIndex     = 3 // 索引字符所在的行索引
)

// 结构体中的名称
func SetTagName(s string) {
	TagName = s
}

// 跳过的行
func SetSkipHeaderNum(n int) {
	SkipHeaderNum = n
}

// excel表中与结构体tag中对应的名称所在的行
func SetFiledNameLine(n int) {
	KeysIndex = n
}

func GetSheetList(filename string) ([]string, error) {
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	xlsx.Close()
	return xlsx.GetSheetList(), nil
}

func LoadXlsxFile(filename string, out any, sheet string, rowCallback func(v any)) (err error) {

	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}
	defer xlsx.Close()

	if sheet == "" {
		sheet = "Sheet1"
	}
	rows, err := xlsx.GetRows(sheet)
	if err != nil {
		return err
	}
	indexMap := make(map[int]string)
	for i, row := range rows {

		if i == KeysIndex-1 {
			for i, key := range row {
				indexMap[i] = key
			}
			continue
		}
		if i < SkipHeaderNum {
			continue
		}

		isP := false
		val := reflect.ValueOf(out)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
			isP = true
		}
		innerTy := val.Type()
		outVal := reflect.New(innerTy)
		innerVal := outVal.Elem()
		outTagMap := getFieldInfos(innerTy)
		var loadSub = false
		for i, val := range row {
			if fieldIndex, ok := outTagMap[indexMap[i]]; ok { // 第一层
				err = SetField(innerVal.Field(fieldIndex), val, true)
				if err != nil {
					panic(err)
					return
				}
			} else if !loadSub {
				for j := 0; j < innerVal.NumField(); j++ {
					if innerVal.Field(j).Kind() == reflect.Struct || innerVal.Field(j).Kind() == reflect.Ptr {
						out := innerVal.Field(j)
						if out.Kind() == reflect.Ptr {
							out = reflect.New(out.Type())
						}
						v, err := getSubValue(out.Interface(), indexMap, row)
						if err != nil {
							zap.Error(err)
							return err
						}
						loadSub = true
						innerVal.Field(j).Set(v)
						break
					}
				}
			}
		}
		if isP {
			rowCallback(outVal.Interface())
		} else {
			rowCallback(innerVal.Interface())
		}
	}
	return nil
}

func getSubValue(out any, tagMap map[int]string, orig []string) (reflect.Value, error) {
	value := reflect.ValueOf(out)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	isP := false
	innerTy := value.Type()
	if innerTy.Kind() == reflect.Ptr {
		isP = true
		innerTy = innerTy.Elem()
	}
	subVal := reflect.New(innerTy)
	subInner := subVal.Elem()

	subTagMap := getFieldInfos(innerTy)
	for k, v := range orig {
		if subFieldIndex, ok := subTagMap[tagMap[k]]; ok {
			err := SetField(subInner.Field(subFieldIndex), v, true)
			if err != nil {
				return reflect.Value{}, err
			}
		}
	}
	if isP {
		return subVal, nil
	}
	return subInner, nil
}

func getFieldInfos(rType reflect.Type) map[string]int {
	fieldInfo := make(map[string]int)
	fieldsCount := rType.NumField()
	for i := 0; i < fieldsCount; i++ {
		field := rType.Field(i)
		if field.PkgPath != "" {
			continue
		}

		filteredTags := field.Tag.Get(TagName)
		if filteredTags == "" || filteredTags == "-" {
			continue
		}
		fieldInfo[filteredTags] = i
	}
	return fieldInfo
}
