package reflectfun

import (
	"fmt"
	"reflect"
)

/*
在不清楚传入数据类型和字段信息时可以基于reflect获取相关信息包括字段名/字段类型/值类型/值字面量等
reflect.TypeOf():获取类型信息
	: 获取传入参数类型 reflect.TypeOf(data).Kind() 可以实现类型断言
	: 只能获取struct字段个数 reflect.TypeOf(data).NumField()
	: 可以获取slice、数组、chan元素的类型 reflect.TypeOf(data).Elem()
	: 获取结构体每个字段信息包括字段名/字段类型/标签 reflect.TypeOf(data).NumField()
	: 获取结构体字段名也就是属性名 reflect.TypeOf(data).Field(i).Name
	: 获取结构体标签信息 get(key)获取每一个标签的信息 reflect.TypeOf(data).Field(i).Tip



reflect.ValueOf():获取值信息
	: 基于索引获取slice/map/chan/string的元素值 reflect.ValueOf(data).Index(i)
	: 获取所有key的切片 reflect.ValueOf(data).MapKeys()
	: 基于key获取value reflect.ValueOf(data).MapIndex(key)
	: 基于key获取value的数据类型 reflect.ValueOf(data).MapIndex(key).Kind()
	: 获取结构体每个字段的值 reflect.ValueOf(data).Field(i)
	: 获取结构体每个字段的值类型 reflect.ValueOf(data).Field(i).Kind()


*/

func ReflectGetFun[T any](data T) {
	// 获取传入变量类型信息
	types := reflect.TypeOf(data)
	// 获取传入变量对应类型
	pType := types.Kind()
	// 获取值信息
	values := reflect.ValueOf(data)

	// 获取字段或元素个数

	switch {
	case pType == reflect.Array || pType == reflect.Slice:
		// 数组和切片的操作
		// 获取每一个元素类型和字面值
		oneType := types.Elem()
		fmt.Println("每个元素类型是", oneType)

		fmt.Println("值是", values)
		for i := 0; i < values.Len(); i++ {
			fmt.Printf("传入参数名是%v, 元素类型是%v, 元素字面量是%v\n", types.Name(), oneType, values.Index(i))
		}

	case pType == reflect.Map:
		// map(映射)操作
		// 获取key信息
		fmt.Println("value 是", values)
		for _, key := range values.MapKeys() {
			fmt.Println("keys 是", key)
			fmt.Println("key 的类型是", key.Kind())
			fmt.Println("value 是", values.MapIndex(key))
			fmt.Println("value 的类型是", values.MapIndex(key).Kind())
		}
	case pType == reflect.String:
		// 字符串操作
		// 获取字符串字面量
		fmt.Println("字符串字面量是", values)

		// 获取字符串每一个元素
		for i := 0; i < values.Len(); i++ {
			fmt.Println("字符串每一个元素是", values.Index(i))
		}
	case pType == reflect.Struct:
		// 结构体操作
		nums := types.NumField()
		fmt.Println("传入的结构体字段数量是", nums)
		for i := 0; i < nums; i++ {
			value := values.Field(i)
			fmt.Println("value ", value)
			field := types.Field(i)
			fmt.Println("字段和标签信息是 ", field)
			keyName := field.Name
			fmt.Println("字段名 ", keyName)
			valueType := value.Kind()
			fmt.Println("值类型是 ", valueType)
			tips := field.Tag
			fmt.Println("标签是 ", tips)
			jsonName := tips.Get("json")
			fmt.Println("json名是", jsonName)
		}
	case pType == reflect.Chan:
		// chan(通道)操作
	case pType == reflect.Func:
		// 函数操作
	default:
		// 其他操作
	}
}
