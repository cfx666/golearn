package main

import (
	"fmt"
	"io"
	"reflect"
)

type Person struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p Person) Print() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {

	var i int16 = 10
	iv := reflect.ValueOf(i)     // 获取i的值
	i2 := iv.Interface().(int16) // 将reflect.Value转换为int16
	fmt.Println(i2)

	// 通过反射修改值
	iv = reflect.ValueOf(&i) // 获取i的指针
	iv.Elem().SetInt(20)     // 因为获取的是指针，所以需要使用elem得到值，然后修改i的值
	fmt.Println(i)

	// 通过反射修改结构体的值
	person := Person{Name: "乐破晓", Age: 20}
	pv := reflect.ValueOf(&person)
	pv.Elem().Field(0).SetString("张三")
	pv.Elem().FieldByName("Name").SetString("张三")
	fmt.Println(person)

	// 获取底层类型
	fmt.Println("底层类型：", reflect.TypeOf(Person{}).Kind())
	fmt.Println("底层类型：", reflect.TypeOf(&Person{}).Kind())

	// 遍历结构体的字段和方法
	var p2 = Person{Name: "李四", Age: 20}
	pi := reflect.TypeOf(p2)
	for i := 0; i < pi.NumField(); i++ {
		fmt.Println("字段：", pi.Field(i).Name)
	}
	for i := 0; i < pi.NumMethod(); i++ {
		fmt.Println("方法：", pi.Method(i).Name)
	}

	for i := 0; i < pi.NumField(); i++ {
		fmt.Println("字段：", pi.Field(i).Tag.Get("json"))
	}

	// 判断是否实现了某个接口
	p := Person{Name: "乐破晓", Age: 20}
	pt := reflect.TypeOf(p)
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：", pt.Implements(stringerType))
	fmt.Println("是否实现了io.Writer：", pt.Implements(writerType))

	// 根据名字调用方法
	pv.MethodByName("Print").Call([]reflect.Value{})
}
