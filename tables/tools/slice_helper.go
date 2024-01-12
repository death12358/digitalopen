package tools

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
)

// FindSlice elem在slice裡的index (-1不存在)
func FindSlice(slice, elem interface{}) (int, bool) {
	sV := reflect.ValueOf(slice)
	if sV.Kind() != reflect.Slice {
		log.Panic("FindSlice: not a slice")
	}

	for i := 0; i < sV.Len(); i++ {
		if reflect.DeepEqual(sV.Index(i).Interface(), elem) {
			return i, true
		}
	}
	return -1, false
}

// CalcSliceElementCount slice裡有多少elem
func CalcSliceElementCount(slice, elem interface{}) int {
	sV := reflect.ValueOf(slice)
	count := 0
	if sV.Kind() != reflect.Slice {
		return count
	}
	for i := 0; i < sV.Len(); i++ {
		if reflect.DeepEqual(sV.Index(i).Interface(), elem) {
			count++
		}
	}
	return count
}

func RemoveDuplicates(targetSlice interface{}) interface{} {
	sV := reflect.ValueOf(targetSlice)

	//check if the input is a slice
	if sV.Kind() != reflect.Slice {
		log.Panic("RemoveDuplicates: not a slice.")
	}

	//check if the input slice contains elements that can be sorted
	if !sV.Type().Elem().Comparable() {
		log.Panic("RemoveDuplicates: elements are not comparable.")
	}

	//sort
	var less func(i, j int) bool
	switch sV.Type().Elem().Kind() {
	case reflect.String:
		less = func(i, j int) bool {
			return sV.Index(i).String() < sV.Index(j).String()
		}
	case reflect.Int, reflect.Int32, reflect.Int64:
		less = func(i, j int) bool {
			return sV.Index(i).Int() < sV.Index(j).Int()
		}
	default:
		log.Panic("RemoveDuplicates: unsupported slice element type.")
	}
	sort.Slice(targetSlice, less)

	//remove duplicate elements from the slice
	rslt := reflect.MakeSlice(sV.Type(), 0, sV.Len())
	for i := 0; i < sV.Len(); i++ {
		if i == 0 || sV.Index(i).Interface() != sV.Index(i-1).Interface() {
			rslt = reflect.Append(rslt, sV.Index(i))
		}
	}
	return rslt.Interface()
}

func GetSliceElem[T int | int32 | int64 | float32 | float64 | string](slice []T, index int) (error, T) {
	if len(slice) >= index {
		return fmt.Errorf("index:[%d]out of range", index), slice[len(slice)-1]
	}
	return nil, slice[index]
}

// ConvertToSliceOfString 接受一个任意类型的切片并返回 []string{}
func ConvertToSliceOfString(input interface{}) ([]string, error) {
	// 检查输入是否为切片
	inputValue := reflect.ValueOf(input)
	if inputValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("输入必须是一个切片")
	}

	// 创建一个字符串切片来存储结果
	result := make([]string, inputValue.Len())

	// 遍历输入切片并将每个元素转换为字符串
	for i := 0; i < inputValue.Len(); i++ {
		elem := inputValue.Index(i)
		str, err := toString(elem)
		if err != nil {
			return nil, err
		}
		result[i] = str
	}

	return result, nil
}

// toString 将任意类型的值转换为字符串
func toString(value reflect.Value) (string, error) {
	switch value.Kind() {
	case reflect.String:
		return value.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("不支持的类型：%s", value.Kind())
	}
}
