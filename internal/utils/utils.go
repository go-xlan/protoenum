// Package utils: Private utilities handling pointers and zero values
// Provides generic functions to handle pointers and value extraction
//
// utils: 指针和零值操作的内部工具包
// 提供用于处理指针创建和值提取的泛型函数
package utils

// GetValuePointer returns a pointer to the given value
// Creates a new pointer pointing to a clone of the input value
//
// 返回给定值的指针
// 创建一个指向输入值副本的新指针
func GetValuePointer[T any](v T) *T {
	return &v
}

// GetPointerValue returns the value pointed to, returns zero value if nil
// Safe dereference that handles nil pointers without panic
//
// 返回指针指向的值，如果为 nil 则返回零值
// 安全的解引用操作，处理 nil 指针时不会 panic
func GetPointerValue[T any](v *T) T {
	if v != nil {
		return *v
	}
	return Zero[T]()
}

// Zero returns the zero value of type T
// Uses Go's default zero value initialization mechanism
//
// 返回类型 T 的零值
// 使用 Go 的默认零值初始化机制
func Zero[T any]() T {
	var zero T
	return zero
}
