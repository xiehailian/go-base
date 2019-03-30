package main

import (
	"fmt"
)

// 用结构体提供更多的错误信息
type AreaError struct {
    err    string 	//error description
    length float64 	//length which caused the error
    width  float64 	//width which caused the error
}

func NewAreaError(err string, length float64, width float64) *AreaError {
	return &AreaError{err, length, width}
}

func (e *AreaError) Error() string {
    return e.err
}

func (e *AreaError) lengthNegative() bool {
    return e.length < 0
}

func (e *AreaError) widthNegative() bool {
    return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
    err := ""
    if length < 0 {
        err += "length is less than zero"
    }
    if width < 0 {
        if err == "" {
            err = "width is less than zero"
        } else {
            err += ", width is less than zero"
        }
    }
    if err != "" {
        return 0, NewAreaError(err, length, width)
    }
    return length * width, nil
}


func main() {
    length, width := -5.0, -9.0
    area, err := rectArea(length, width)
    if err != nil {
        if err, ok := err.(*AreaError); ok {    // 错误类型断言，类似于python中try expect 后的那个异常类型
            if err.lengthNegative() {
                fmt.Printf("error: length %0.2f is less than zero\n", err.length)

            }
            if err.widthNegative() {
                fmt.Printf("error: width %0.2f is less than zero\n", err.width)

            }
            return
        }
        fmt.Println(err)
        return
    }
    fmt.Println("area of rect", area)
}


