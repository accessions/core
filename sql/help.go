package sql

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func BuildBatchUpdateSqlArray(tableName string, dataList interface{}) ([]string) {
	fieldValue := reflect.ValueOf(dataList)
	fieldType := reflect.TypeOf(dataList).Elem().Elem()
	sliceLength := fieldValue.Len()
	fieldNum := fieldType.NumField()
	// 检验结构体标签是否为空和重复
	verifyTagDuplicate := make(map[string]string)
	for i := 0; i < fieldNum; i++ {
		fieldTag := fieldType.Field(i).Tag.Get("gorm")
		fieldName := _getFieldName(fieldTag)
		if len(strings.TrimSpace(fieldName)) == 0 {
			return nil
		}
		/*if !strings.HasPrefix(fieldName, "id;") {
			return nil
		}*/
		_, ok := verifyTagDuplicate[fieldName]
		if !ok {
			verifyTagDuplicate[fieldName] = fieldName
		} else {
			return nil
		}
	}
	var IDList []string
	updateMap := make(map[string][]string)
	for i := 0; i < sliceLength; i++ {
		// 得到某一个具体的结构体的
		structValue := fieldValue.Index(i).Elem()
		for j := 0; j < fieldNum; j++ {
			elem := structValue.Field(j)
			var temp string
			switch elem.Kind() {
			case reflect.Int64:
				temp = strconv.FormatInt(elem.Int(), 10)
			case reflect.String:
				if strings.Contains(elem.String(), "'") {
					temp = fmt.Sprintf("'%v'", strings.ReplaceAll(elem.String(), "'", "\\'"))
				} else {
					temp = fmt.Sprintf("'%v'", elem.String())
				}
			case reflect.Float64:
				temp = strconv.FormatFloat(elem.Float(), 'f', -1, 64)
			case reflect.Bool:
				temp = strconv.FormatBool(elem.Bool())
			default:
				return nil
			}
			gormTag := fieldType.Field(j).Tag.Get("gorm")
			fieldTag := _getFieldName(gormTag)
			if strings.HasPrefix(fieldTag, "id;") {
				id, err := strconv.ParseInt(temp, 10, 64)
				if err != nil {
					return nil
				}
				if id < 1 {
					return nil
				}
				IDList = append(IDList, temp)
				continue
			}
			valueList := append(updateMap[fieldTag], temp)
			updateMap[fieldTag] = valueList
		}
	}
	length := len(IDList)
	size := 200
	SQLQuantity := _getSqlQuantity(length, size)
	var SQLArray []string
	k := 0
	for i := 0; i < SQLQuantity; i++ {
		count := 0
		var record bytes.Buffer
		record.WriteString("UPDATE " + tableName + " SET ")
		for fieldName, fieldValueList := range updateMap {
			record.WriteString(fieldName)
			record.WriteString(" = CASE " + "id")
			for j := k; j < len(IDList) && j < len(fieldValueList) && j < size+k; j++ {
				record.WriteString(" WHEN " + IDList[j] + " THEN " + fieldValueList[j])
			}
			count++
			if count != fieldNum-1 {
				record.WriteString(" END, ")
			}
		}
		record.WriteString(" END WHERE ")
		record.WriteString("id" + " IN (")
		min := size + k
		if len(IDList) < min {
			min = len(IDList)
		}
		record.WriteString(strings.Join(IDList[k:min], ","))
		record.WriteString(");")
		k += size
		SQLArray = append(SQLArray, record.String())
	}
	return SQLArray
}

func _getSqlQuantity(length, size int) int {
	SQLQuantity := int(math.Ceil(float64(length) / float64(size)))
	return SQLQuantity
}

func _getFieldName(fieldTag string) string {
	fieldTagArr := strings.Split(fieldTag, ":")
	if len(fieldTagArr) == 0 {
		return ""
	}
	fieldName := fieldTagArr[len(fieldTagArr)-1]
	return fieldName
}


func CombinSql(mainSql string, whereSql []string) string {
	if len(whereSql) < 1 {
		return strings.ReplaceAll(mainSql, "$WHERE", "")
	}
	str := " WHERE "
	for _, i := range whereSql {
		str += fmt.Sprintf(" %s ~", i)
	}
	str = strings.TrimRight(str, "~")
	str = strings.ReplaceAll(str, "~", "AND ")
	mainSql = strings.ReplaceAll(mainSql, "$WHERE", str)
	return mainSql
}


// AddRetailDeals batch execute
/*func AddRetailDeals(installables []rsp.AiotRetailDeal) {
	fmt.Println(len(installables))
	db := mysqlconn.GetDB(setting.AiotDb)
	defer db.Close()
	inserts := utils.ToInterfaceArr(installables)
	length := len(installables) / 1000
	for i := 0; i < length; i++ {
		insertItems := inserts[i*1000 : (i+1)*1000-1]
		err := gormbulk.BulkInsert(db, insertItems, len(insertItems))
		if err != nil {
			fmt.Println(err)
		}
	}
	insertItems := inserts[length*1000:]
	err := gormbulk.BulkInsert(db, insertItems, len(insertItems))
	if err != nil {
		fmt.Println(err)
	}
}*/
