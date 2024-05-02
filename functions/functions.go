package functions

func Reduce[T any, R any](array []T, callback func(R, T) R, initialValue R) R {
	result := initialValue
	for _, value := range array {
		result = callback(result, value)
	}
	return result
}

func Filter[T any](array []T, callback func(T) bool) []T {
	var result []T
	for _, value := range array {
		if callback(value) {
			result = append(result, value)
		}
	}
	return result
}

func Map[T any, R any](array []T, callback func(T) R) []R {
	var result []R
	for _, value := range array {
		result = append(result, callback(value))
	}
	return result
}

func Find[T any](array []T, callback func(T) bool) *T {
	for _, value := range array {
		if callback(value) {
			return &value
		}
	}
	return nil
}

func FindLast[T any](array []T, callback func(T) bool) *T {
	for i := len(array) - 1; i >= 0; i-- {
		if callback(array[i]) {
			return &array[i]
		}
	}
	return nil
}

func FindIndex[T any](array []T, callback func(T) bool) int {
	for i, value := range array {
		if callback(value) {
			return i
		}
	}
	return -1
}

func FindLastIndex[T any](array []T, callback func(T) bool) int {
	for i := len(array) - 1; i >= 0; i-- {
		if callback(array[i]) {
			return i
		}
	}
	return -1
}

func Every[T any](array []T, callback func(T) bool) bool {
	for _, value := range array {
		if !callback(value) {
			return false
		}
	}
	return true
}

func Some[T any](array []T, callback func(T) bool) bool {
	for _, value := range array {
		if callback(value) {
			return true
		}
	}
	return false
}

func Includes[T comparable](array []T, searchElement T) bool {
	for _, value := range array {
		if value == searchElement {
			return true
		}
	}
	return false
}

func Concat[T any](arrays ...[]T) []T {
	var result []T
	for _, array := range arrays {
		result = append(result, array...)
	}
	return result
}

func IndexOf[T comparable](array []T, searchElement T) int {
	for i, value := range array {
		if value == searchElement {
			return i
		}
	}
	return -1
}

func LastIndexOf[T comparable](array []T, searchElement T) int {
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == searchElement {
			return i
		}
	}
	return -1
}

func Reverse[T any](array []T) []T {
	result := make([]T, len(array))
	for i, value := range array {
		result[len(array)-1-i] = value
	}
	return result
}
