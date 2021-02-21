package gopylists

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type PyList struct {
	objs  []string
	types []string
}

func NewPyList() PyList {
	return PyList{objs: []string{}, types: []string{}}
}

func NewPyListFromArgs(objs ...interface{}) PyList {
	plist := NewPyList()
	for _, obj := range objs {
		plist.Append(obj)
	}
	return plist
}

func (plist *PyList) Append(obj interface{}) {
	switch v := obj.(type) {
	case int:
		plist.objs = append(plist.objs, strconv.FormatInt(int64(v), 10))
		plist.types = append(plist.types, "int")
	case int8:
		plist.objs = append(plist.objs, strconv.FormatInt(int64(v), 10))
		plist.types = append(plist.types, "int8")
	case int16:
		plist.objs = append(plist.objs, strconv.FormatInt(int64(v), 10))
		plist.types = append(plist.types, "int16")
	case int32:
		plist.objs = append(plist.objs, strconv.FormatInt(int64(v), 10))
		plist.types = append(plist.types, "int32")
	case int64:
		plist.objs = append(plist.objs, strconv.FormatInt(int64(v), 10))
		plist.types = append(plist.types, "int64")
	case uint:
		plist.objs = append(plist.objs, strconv.FormatUint(uint64(v), 10))
		plist.types = append(plist.types, "uint")
	case uint8:
		plist.objs = append(plist.objs, strconv.FormatUint(uint64(v), 10))
		plist.types = append(plist.types, "uint8")
	case uint16:
		plist.objs = append(plist.objs, strconv.FormatUint(uint64(v), 10))
		plist.types = append(plist.types, "uint16")
	case uint32:
		plist.objs = append(plist.objs, strconv.FormatUint(uint64(v), 10))
		plist.types = append(plist.types, "uint32")
	case uint64:
		plist.objs = append(plist.objs, strconv.FormatUint(uint64(v), 10))
		plist.types = append(plist.types, "uint64")
	case float32:
		plist.objs = append(plist.objs, strconv.FormatFloat(float64(v), 'f', -1, 64))
		plist.types = append(plist.types, "float32")
	case float64:
		plist.objs = append(plist.objs, strconv.FormatFloat(float64(v), 'f', -1, 64))
		plist.types = append(plist.types, "float64")
	case complex64:
		plist.objs = append(plist.objs, strconv.FormatComplex(complex128(v), 'f', -1, 64))
		plist.types = append(plist.types, "complex64")
	case complex128:
		plist.objs = append(plist.objs, strconv.FormatComplex(complex128(v), 'f', -1, 64))
		plist.types = append(plist.types, "complex128")
	case string:
		plist.objs = append(plist.objs, v)
		plist.types = append(plist.types, "string")
	}
	// append(plist.objs, strconv.)
}

func (plist *PyList) Insert(obj interface{}, index uint) {
	switch v := obj.(type) {
	case int:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatInt(int64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "int"
	case int8:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatInt(int64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "int8"
	case int16:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatInt(int64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "int16"
	case int32:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatInt(int64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "int32"
	case int64:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatInt(int64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "int64"
	case uint:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatUint(uint64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "uint"
	case uint8:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatUint(uint64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "uint8"
	case uint16:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatUint(uint64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "uint16"
	case uint32:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatUint(uint64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "uint32"
	case uint64:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatUint(uint64(v), 10)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "uint64"
	case float32:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatFloat(float64(v), 'f', -1, 64)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "float32"
	case float64:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatFloat(float64(v), 'f', -1, 64)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "float64"
	case complex64:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatComplex(complex128(v), 'f', -1, 64)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "complex64"
	case complex128:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = strconv.FormatComplex(complex128(v), 'f', -1, 64)
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "complex128"
	case string:
		plist.objs = append(plist.objs, "0")
		copy(plist.objs[index+1:], plist.objs[index:])
		plist.objs[index] = v
		plist.types = append(plist.types, "")
		copy(plist.types[index+1:], plist.types[index:])
		plist.types[index] = "string"
	}
}

func (plist *PyList) Remove(index uint) {
	plist.objs = append(plist.objs[:index], plist.objs[index+1:]...)
	plist.types = append(plist.types[:index], plist.types[index+1:]...)
}

func (plist *PyList) Get(index uint) interface{} {
	if int(index) >= plist.GetLength() {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, plist.GetLength()))
	}
	switch plist.types[index] {
	case "int":
		out, _ := strconv.ParseInt(plist.objs[index], 10, 0)
		return int(out)
	case "int8":
		out, _ := strconv.ParseInt(plist.objs[index], 10, 8)
		return int8(out)
	case "int16":
		out, _ := strconv.ParseInt(plist.objs[index], 10, 16)
		return int16(out)
	case "int32":
		out, _ := strconv.ParseInt(plist.objs[index], 10, 32)
		return int32(out)
	case "int64":
		out, _ := strconv.ParseInt(plist.objs[index], 10, 64)
		return int64(out)
	case "uint":
		out, _ := strconv.ParseUint(plist.objs[index], 10, 0)
		return uint(out)
	case "uint8":
		out, _ := strconv.ParseUint(plist.objs[index], 10, 8)
		return uint8(out)
	case "uint16":
		out, _ := strconv.ParseUint(plist.objs[index], 10, 16)
		return uint16(out)
	case "uint32":
		out, _ := strconv.ParseUint(plist.objs[index], 10, 32)
		return uint32(out)
	case "uint64":
		out, _ := strconv.ParseUint(plist.objs[index], 10, 64)
		return uint64(out)
	case "float32":
		out, _ := strconv.ParseFloat(plist.objs[index], 32)
		return float32(out)
	case "float64":
		out, _ := strconv.ParseFloat(plist.objs[index], 64)
		return float64(out)
	case "complex64":
		out, _ := strconv.ParseComplex(plist.objs[index], 64)
		return complex64(out)
	case "complex128":
		out, _ := strconv.ParseComplex(plist.objs[index], 128)
		return complex128(out)
	case "string":
		return plist.objs[index]
	default:
		panic("error getting element")
		return -1
	}
}

func (plist *PyList) GetLength() int {
	return len(plist.objs)
}

func (plist *PyList) Clear() {
	plist.objs = []string{}
	plist.types = []string{}
}

func (plist *PyList) Concatenate(plist2 PyList) {
	for i := 0; i < plist2.GetLength(); i++ {
		plist.Append(plist2.Get(uint(i)))
	}
}

func (plist *PyList) Count(obj interface{}) int {
	times := 0
	for i := 0; i < plist.GetLength(); i++ {
		if plist.Get(uint(i)) == obj {
			times++
		}
	}

	return times
}

func In(array interface{}, obj interface{}) bool {
	isin := false
	// fmt.Println(reflect.TypeOf(array).String())
	if strings.Replace(reflect.TypeOf(array).String(), "[]", "", -1) != reflect.TypeOf(obj).String() && reflect.TypeOf(array).String() != "gopylists.PyList" {
		panic("array and object are not same type")
	}
	switch v := array.(type) {
	case []int:
		objr := obj.(int)
		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []int8:
		objr := obj.(int8)
		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []int16:
		objr := obj.(int16)

		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []int32:
		objr := obj.(int32)

		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []int64:
		objr := obj.(int64)

		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []uint:
		objr := obj.(uint)

		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []uint8:
		var objr uint8 = obj.(uint8)

		for _, o := range v {
			fmt.Println(reflect.TypeOf(v))
			fmt.Println(reflect.TypeOf(obj))
			if o == objr {
				isin = true
				break
			}
		}
		fmt.Println("uint8")
	case []uint16:
		objr := obj.(uint16)

		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []uint32:
		objr := obj.(uint32)

		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []uint64:
		objr := obj.(uint64)

		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []float32:
		objr := obj.(float32)
		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []float64:
		objr := obj.(float64)

		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []complex64:
		objr := obj.(complex64)
		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	case []complex128:
		objr := obj.(complex128)
		for _, o := range v {
			if o == objr {
				isin = true
				break
			}
		}
	// case []byte:
	// case []rune:
	case []string:
		for _, o := range v {
			if o == obj {
				isin = true
				break
			}
		}
	case PyList:
		for i := 0; i < v.GetLength(); i++ {
			if v.Get(uint(i)) == obj {
				isin = true
				break
			}
		}
	// case rune:
	default:
		panic("unrecongized type")

	}
	return isin
}

func (plist PyList) String() string {
	out := "["
	for i := 0; i < plist.GetLength(); i++ {
		if i == plist.GetLength()-1 {
			out += plist.objs[i]
		} else {
			out += plist.objs[i] + ", "
		}
	}
	out += "]"
	return out
}
