package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func ParseVersion(i interface{}) int64 {
	rVal := reflect.ValueOf(i)
	fmt.Println(rVal.Kind())
	switch rVal.Kind() {
	case reflect.String:
		if ret, err := strconv.ParseInt(rVal.String(), 10, 64); err == nil {
			return ret
		}
	case reflect.Float64:
		return int64(rVal.Float())
	case reflect.Int64:
		return rVal.Int()
	}
	return 0
}

func main()  {

	fmt.Println(ParseVersion(float64(2323)))

}