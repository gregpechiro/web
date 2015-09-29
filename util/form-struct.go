package util

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func FormToStruct(ptr interface{}, vals url.Values, start string) {
	var strct reflect.Value
	if reflect.TypeOf(ptr) == reflect.TypeOf(reflect.Value{}) {
		strct = ptr.(reflect.Value)
	} else {
		strct = reflect.ValueOf(ptr).Elem()
	}
	strctType := strct.Type()
	for i := 0; i < strct.NumField(); i++ {
		fld := strct.Field(i)
		if vals.Get(start+strctType.Field(i).Name) != "" || fld.Kind() == reflect.Struct {
			switch fld.Kind() {
			case reflect.String:
				strct.Field(i).SetString(vals.Get(start + ToLowerFirst(strctType.Field(i).Name)))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				in, _ := strconv.ParseInt(vals.Get(start+ToLowerFirst(strctType.Field(i).Name)), 10, 64)
				strct.Field(i).SetInt(in)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				u, _ := strconv.ParseUint(vals.Get(start+ToLowerFirst(strctType.Field(i).Name)), 10, 64)
				strct.Field(i).SetUint(u)
			case reflect.Float32, reflect.Float64:
				f, _ := strconv.ParseFloat(vals.Get(start+ToLowerFirst(strctType.Field(i).Name)), 64)
				strct.Field(i).SetFloat(f)
			case reflect.Bool:
				b, _ := strconv.ParseBool(vals.Get(start + ToLowerFirst(strctType.Field(i).Name)))
				strct.Field(i).SetBool(b)
			case reflect.Map:
				strct.Field(i).Set(reflect.MakeMap(strct.Field(i).Type()))
			case reflect.Slice:
				ss := reflect.MakeSlice(strct.Field(i).Type(), 0, 0)
				strct.Field(i).Set(genSlice(ss, vals.Get(start+ToLowerFirst(strctType.Field(i).Name))))
			case reflect.Struct:
				st := reflect.Indirect(reflect.New(strct.Field(i).Type()))
				FormToStruct(st, vals, start+ToLowerFirst(strctType.Field(i).Name)+".")
				strct.Field(i).Set(st)
			}
		}
	}
}

func genSlice(sl reflect.Value, val string) reflect.Value {
	vs := strings.Split(val, ",")
	for _, v := range vs {
		switch sl.Type().String() {
		case "[]string":
			sl = reflect.Append(sl, reflect.ValueOf(v))
		case "[]int":
			in, _ := strconv.ParseInt(v, 10, 0)
			sl = reflect.Append(sl, reflect.ValueOf(int(in)))
		case "[]int8":
			in, _ := strconv.ParseInt(v, 10, 8)
			sl = reflect.Append(sl, reflect.ValueOf(int8(in)))
		case "[]int16":
			in, _ := strconv.ParseInt(v, 10, 16)
			sl = reflect.Append(sl, reflect.ValueOf(int16(in)))
		case "[]int32":
			in, _ := strconv.ParseInt(v, 10, 32)
			sl = reflect.Append(sl, reflect.ValueOf(int32(in)))
		case "[]int64":
			in, _ := strconv.ParseInt(v, 10, 64)
			sl = reflect.Append(sl, reflect.ValueOf(int64(in)))
		case "[]uint":
			in, _ := strconv.ParseUint(v, 10, 0)
			sl = reflect.Append(sl, reflect.ValueOf(uint(in)))
		case "[]uint8":
			in, _ := strconv.ParseUint(v, 10, 8)
			sl = reflect.Append(sl, reflect.ValueOf(uint8(in)))
		case "[]uint16":
			in, _ := strconv.ParseUint(v, 10, 16)
			sl = reflect.Append(sl, reflect.ValueOf(uint16(in)))
		case "[]uint32":
			in, _ := strconv.ParseUint(v, 10, 32)
			sl = reflect.Append(sl, reflect.ValueOf(uint32(in)))
		case "[]uint64":
			in, _ := strconv.ParseUint(v, 10, 64)
			sl = reflect.Append(sl, reflect.ValueOf(uint64(in)))
		case "[]float32":
			in, _ := strconv.ParseFloat(v, 32)
			sl = reflect.Append(sl, reflect.ValueOf(float32(in)))
		case "[]float64":
			in, _ := strconv.ParseFloat(v, 64)
			sl = reflect.Append(sl, reflect.ValueOf(float64(in)))
		case "[]bool":
			b, _ := strconv.ParseBool(v)
			sl = reflect.Append(sl, reflect.ValueOf(b))
		}
	}
	return sl
}
