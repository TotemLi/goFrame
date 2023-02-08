package stringx

import (
	"errors"
	"reflect"
	"strconv"
)

// StrCopyToPtr 用string的内容赋值给指针
func StrCopyToPtr(valStr string, ptr any) error {
	var err error
	v := reflect.ValueOf(ptr)
	switch v.Elem().Kind() {
	case reflect.Bool:
		*(v.Interface().(*bool)), err = strconv.ParseBool(valStr)
	case reflect.String:
		*(v.Interface().(*string)) = valStr
	case reflect.Int:
		*(v.Interface().(*int)), err = strconv.Atoi(valStr)
	case reflect.Int8:
		var i int64
		i, err = strconv.ParseInt(valStr, 10, 8)
		*(v.Interface().(*int8)) = int8(i)
	case reflect.Int16:
		var i int64
		i, err = strconv.ParseInt(valStr, 10, 16)
		*(v.Interface().(*int16)) = int16(i)
	case reflect.Int32:
		var i int64
		i, err = strconv.ParseInt(valStr, 10, 32)
		*(v.Interface().(*int32)) = int32(i)
	case reflect.Int64:
		*(v.Interface().(*int64)), err = strconv.ParseInt(valStr, 10, 64)
	case reflect.Uint:
		var i uint64
		i, err = strconv.ParseUint(valStr, 10, 64)
		*(v.Interface().(*uint)) = uint(i)
	case reflect.Uint8:
		var i uint64
		i, err = strconv.ParseUint(valStr, 10, 8)
		*(v.Interface().(*uint8)) = uint8(i)
	case reflect.Uint16:
		var i uint64
		i, err = strconv.ParseUint(valStr, 10, 16)
		*(v.Interface().(*uint16)) = uint16(i)
	case reflect.Uint32:
		var i uint64
		i, err = strconv.ParseUint(valStr, 10, 32)
		*(v.Interface().(*uint32)) = uint32(i)
	case reflect.Uint64:
		*(v.Interface().(*uint64)), err = strconv.ParseUint(valStr, 10, 64)
	case reflect.Float32:
		var i float64
		i, err = strconv.ParseFloat(valStr, 32)
		*(v.Interface().(*float32)) = float32(i)
	case reflect.Float64:
		*(v.Interface().(*float64)), err = strconv.ParseFloat(valStr, 64)
	default:
		err = errors.New("stringx.StrCopyToPtr can't support type: " + v.Type().String())
	}
	return err
}
