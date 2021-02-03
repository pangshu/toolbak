package Convert
//
//import (
//	"bytes"
//	"fmt"
//	"reflect"
//	"sort"
//	"strings"
//)
//
////// Serialize 对变量进行序列化.
////func Serialize(val interface{}) ([]byte, error) {
////	buf := bytes.Buffer{}
////	enc := gob.NewEncoder(&buf)
////	gob.Register(val)
////
////	err := enc.Encode(&val)
////	if err != nil {
////		return nil, err
////	}
////
////	return buf.Bytes(), nil
////}
//
//
//func (toolConvert *Convert)Serialize(value interface{}) ([]byte, error) {
//
//	if value == nil {
//		return MarshalNil(), nil
//	}
//
//	t := reflect.TypeOf(value)
//	switch t.Kind() {
//	case reflect.Bool:
//		return MarshalBool(value.(bool)), nil
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
//		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64,
//		reflect.Float32, reflect.Float64:
//		return MarshalNumber(value), nil
//	case reflect.String:
//		return MarshalString(value.(string)), nil
//	case reflect.Map:
//		return MarshalMap(value)
//	case reflect.Slice:
//		return MarshalSlice(value)
//	default:
//		return nil, fmt.Errorf("Marshal: Unknown type %T with value %#v", t, value)
//	}
//
//	return nil, nil
//}
//
//func MarshalNil() []byte {
//	return []byte("N;")
//}
//
//func MarshalBool(value bool) []byte {
//	if value {
//		return []byte("b:1;")
//	}
//
//	return []byte("b:0;")
//}
//
//func MarshalNumber(value interface{}) []byte {
//	var val string
//	var toolConvert Convert
//
//	isFloat := false
//
//	switch value.(type) {
//	default:
//		val = "0"
//	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
//		val = toolConvert.ToStr(value)
//	case float32, float64:
//		val = toolConvert.ToStr(value)
//		isFloat = true
//	}
//
//	if isFloat {
//		return []byte("d:" + val + ";")
//
//	} else {
//		return []byte("i:" + val + ";")
//	}
//}
//
//func MarshalString(value string) []byte {
//	return []byte(fmt.Sprintf("s:%d:\"%s\";", len(value), value))
//}
//
//func MarshalMap(value interface{}) ([]byte, error) {
//
//	s := reflect.ValueOf(value)
//
//	mapKeys := s.MapKeys()
//	sort.Slice(mapKeys, func(i, j int) bool {
//		return LessValue(mapKeys[i], mapKeys[j])
//	})
//
//	var buffer bytes.Buffer
//	for _, mapKey := range mapKeys {
//		m, err := Serialize(mapKey.Interface())
//		if err != nil {
//			return nil, err
//		}
//
//		buffer.Write(m)
//
//		m, err = Serialize(s.MapIndex(mapKey).Interface())
//		if err != nil {
//			return nil, err
//		}
//
//		buffer.Write(m)
//	}
//
//	return []byte(fmt.Sprintf("a:%d:{%s}", s.Len(), buffer.String())), nil
//}
//
//func MarshalSlice(value interface{}) ([]byte, error) {
//	s := reflect.ValueOf(value)
//
//	var buffer bytes.Buffer
//	for i := 0; i < s.Len(); i++ {
//		m, err := Serialize(i)
//		if err != nil {
//			return nil, err
//		}
//
//		buffer.Write(m)
//
//		m, err = Serialize(s.Index(i).Interface())
//		if err != nil {
//			return nil, err
//		}
//
//		buffer.Write(m)
//	}
//
//	return []byte(fmt.Sprintf("a:%d:{%s}", s.Len(), buffer.String())), nil
//}
//
//func LessValue(a, b reflect.Value) bool {
//	aValue, aNumerical := NumericalValue(a)
//	bValue, bNumerical := NumericalValue(b)
//
//	if aNumerical && bNumerical {
//		return aValue < bValue
//	}
//
//	if !aNumerical && !bNumerical {
//		// In theory this should mean they are both strings. In reality
//		// they could be any other type and the String() representation
//		// will be something like "<bool>" if it is not a string. Since
//		// distinct values of non-strings still return the same value
//		// here that's what makes the ordering undefined.
//		return strings.Compare(a.String(), b.String()) < 0
//	}
//
//	// Numerical values are always treated as less than other types
//	// (including strings that might represent numbers themselves). The
//	// inverse is also true.
//	return aNumerical && !bNumerical
//}
//
//func NumericalValue(value reflect.Value) (float64, bool) {
//	switch value.Type().Kind() {
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//		return float64(value.Int()), true
//
//	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//		return float64(value.Uint()), true
//
//	case reflect.Float32, reflect.Float64:
//		return value.Float(), true
//
//	default:
//		return 0, false
//	}
//}