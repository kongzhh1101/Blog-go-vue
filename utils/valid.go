package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// 返回结构体中的msg参数
func GetValidMsg(err error, obj any) string {
	// 使用时传入object指针
	getObj := reflect.TypeOf(obj)

	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功，循环每一个错误信息，根据报错字段名，获取结构体具体字段
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}
