package gojoin

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Options struct {
	Sep           string    // the separator between keys (or contains their values)
	KVSep         string    // the separator between key & its value
	IgnoreKey     bool      // whether ignore key , if yes, the key will be ignored, but the value will reserve
	IgnoreEmpty   bool      // whether ignore empty value , if yes, the key & its value will be ignored
	ExceptKeys    []string  // the keys & their values will be ignored
	Order         joinOrder // the join order
	DefinedOrders []string  // the keys order, using with Order == Defined
	StructTag     string    // struct tag, using when src struct type, if not set, will use struct filed name, only support export fields
	URLCoding     urlCoding // the value format, using when format value
	Unwrap        bool      // whether unwrap the internal map or struct
}

type joinOrder uint

const (
	ASCII joinOrder = iota
	ASCIIDesc
	Defined
)

type urlCoding uint

const (
	None urlCoding = iota
	Encoding
	Decoding
)

// Join joins src data to string with defined options
// src is the original data, supports map,struct,url.Values,encoded url string
// options contains join rules
// if join successfully, it will return the result string
// if the kind of src is not support or options contains unsupported rules, it will return error
func Join(src interface{}, options Options) (dst string, err error) {
	j := &join{
		src:     src,
		options: options,
		data:    make(map[string]interface{}),
	}
	return j.Join()
}

type join struct {
	src     interface{}
	options Options
	data    map[string]interface{} // src parsed from src
}

func (j *join) Join() (dst string, err error) {
	rv := reflect.ValueOf(j.src)
	if rv.Kind() == reflect.Ptr {
		rv = reflect.Indirect(rv)
	}
	switch rv.Kind() {
	case reflect.String:
		err = j.parseURLString()
	case reflect.Struct:
		err = j.parseStruct(rv)
	case reflect.Map:
		err = j.parseMap(rv)
	default:
		return "", errors.New(fmt.Sprintf("unsupported type :%s", rv.Type().Name()))
	}
	if err != nil {
		return "", err
	}
	var list []string
	switch j.options.Order {
	case ASCII, ASCIIDesc:
		list, err = j.joinInASCII()
	case Defined:
		if len(j.options.DefinedOrders) == 0 {
			return "", errors.New("need 'DefinedOrders' in Defied order mode")
		}
		list, err = j.joinInDefined()
	default:
		return "", errors.New("unsupported order")
	}
	return strings.Join(list, j.options.Sep), nil
}

func (j *join) parseURLString() (err error) {
	str, ok := j.src.(string)
	if !ok {
		return nil
	}
	index := strings.Index(str, "?")
	values, err := url.ParseQuery(str[index+1:])
	if err != nil {
		return err
	}
	for key := range values {
		j.data[key] = values.Get(key)
	}
	return
}

func (j *join) parseStruct(rv reflect.Value) (err error) {
	if rv.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("unsupported type :%s", rv.Type().Name()))
	}
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		key := rt.Field(i).Name
		if j.options.StructTag != "" {
			key = rt.Field(i).Tag.Get(j.options.StructTag)
		}
		value := rv.Field(i)
		if err = j.parseValue(key, value); err != nil {
			return
		}
	}
	return
}

func (j *join) parseMap(rv reflect.Value) (err error) {
	if rv.Kind() != reflect.Map {
		return errors.New(fmt.Sprintf("unsupported type :%s", rv.Type().Name()))
	}
	for _, key := range rv.MapKeys() {
		kv := ""
		switch key.Kind() {
		case reflect.Bool:
			kv = strconv.FormatBool(key.Bool())
		case reflect.String:
			kv = key.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			kv = strconv.FormatInt(key.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			kv = strconv.FormatUint(key.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			kv = strconv.FormatFloat(key.Float(), 'f', -1, 64)
		default:
			return errors.New(fmt.Sprintf("unsupported key type :%s", rv.Type().Name()))
		}
		value := rv.MapIndex(key)
		if err = j.parseValue(kv, value); err != nil {
			return
		}
	}
	return nil
}

func (j *join) parseValue(key string, rv reflect.Value) (err error) {
	if !rv.CanInterface() {
		return
	}
	if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		if rv.Len() == 0 {
			var i interface{}
			switch rv.Type().Elem().Kind() {
			case reflect.Bool:
				i = false
			case reflect.String:
				i = ""
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				i = 0
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				i = 0
			case reflect.Float32, reflect.Float64:
				i = 0
			default:
				return
			}
			rv = reflect.ValueOf(i)
		} else {
			rv = rv.Index(0)
		}
	}
	//fmt.Println(rv, rv.Type().Kind())
	switch rv.Kind() {
	case reflect.Struct:
		if j.options.Unwrap {
			err = j.parseStruct(rv)
		}
	case reflect.Map:
		if j.options.Unwrap {
			err = j.parseMap(rv)
		}
	case reflect.Ptr, reflect.Slice, reflect.Array, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return errors.New(fmt.Sprintf("unsupported type :%s", rv.Type().Name()))
	default:
		j.data[key] = rv.Interface()
	}
	return
}

func (j *join) joinInASCII() ([]string, error) {
	exceptKeys := make(map[string]int)
	if len(j.options.ExceptKeys) > 0 {
		for _, except := range j.options.ExceptKeys {
			exceptKeys[except] = 1
		}
	}
	keys := make([]string, 0, len(j.data))
	for key := range j.data {
		if _, ok := exceptKeys[key]; ok {
			continue
		}
		keys = append(keys, key)
	}
	switch j.options.Order {
	case ASCII:
		sort.Strings(keys)
	case ASCIIDesc:
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] > keys[j]
		})
	}
	list := make([]string, 0, len(keys))
	for _, key := range keys {
		if value, ignore := j.getValue(key); !ignore {
			list = append(list, value)
		}
	}
	return list, nil
}

func (j *join) joinInDefined() ([]string, error) {
	list := make([]string, 0, len(j.data))
	for _, key := range j.options.DefinedOrders {
		if _, ok := j.data[key]; !ok {
			continue
		}
		if value, ignore := j.getValue(key); !ignore {
			list = append(list, value)
		}
	}
	return list, nil
}

func (j *join) getValue(key string) (value string, ignore bool) {
	temp := j.data[key]
	if j.options.IgnoreEmpty && value == "" {
		return "", true
	}
	value, ok := temp.(string)
	if !ok {
		return value, false
	}
	switch j.options.URLCoding {
	case None:
	case Encoding:
		value = url.QueryEscape(value)
	case Decoding:
		value, _ = url.QueryUnescape(value)
	}
	if j.options.IgnoreKey {
		return value, false
	}
	return key + j.options.KVSep + value, false
}
