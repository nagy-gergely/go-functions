package functions

import (
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("Reduce", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		add := func(acc, value int) int {
			return acc + value
		}
		initialValue := 0
		expected := 10
		got := Reduce(array, add, initialValue)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("Filter", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := []int{2, 4}
		got := Filter(array, isEven)
		if len(got) != len(expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("expected %v but got %v", expected, got)
			}
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("Map", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		double := func(value int) int {
			return value * 2
		}
		expected := []int{2, 4, 6, 8}
		got := Map(array, double)
		if len(got) != len(expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("expected %v but got %v", expected, got)
			}
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("Find", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := 2
		got := Find(array, isEven)
		if got == nil {
			t.Errorf("expected %v but got %v", expected, got)
		}
		if *got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Find", func(t *testing.T) {
		array := []int{1, 3, 5, 7}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := (*int)(nil)
		got := Find(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestFindLast(t *testing.T) {
	t.Run("FindLast", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := 4
		got := FindLast(array, isEven)
		if got == nil {
			t.Errorf("expected %v but got %v", expected, got)
		}
		if *got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("FindLast", func(t *testing.T) {
		array := []int{1, 3, 5, 7}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := (*int)(nil)
		got := FindLast(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestFindIndex(t *testing.T) {
	t.Run("FindIndex", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := 1
		got := FindIndex(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("FindIndex", func(t *testing.T) {
		array := []int{1, 3, 5, 7}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := -1
		got := FindIndex(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestFindLastIndex(t *testing.T) {
	t.Run("FindLastIndex", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := 3
		got := FindLastIndex(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("FindLastIndex", func(t *testing.T) {
		array := []int{1, 3, 5, 7}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := -1
		got := FindLastIndex(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestEvery(t *testing.T) {
	t.Run("Every", func(t *testing.T) {
		array := []int{2, 4, 6, 8}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := true
		got := Every(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Every", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := false
		got := Every(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestSome(t *testing.T) {
	t.Run("Some", func(t *testing.T) {
		array := []int{1, 3, 5, 7}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := false
		got := Some(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Some", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		isEven := func(value int) bool {
			return value%2 == 0
		}
		expected := true
		got := Some(array, isEven)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestIncludes(t *testing.T) {
	t.Run("Includes", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		value := 2
		expected := true
		got := Includes(array, value)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Includes", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		value := 5
		expected := false
		got := Includes(array, value)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestConcat(t *testing.T) {
	t.Run("Concat", func(t *testing.T) {
		array1 := []int{1, 2}
		array2 := []int{3, 4}
		expected := []int{1, 2, 3, 4}
		got := Concat(array1, array2)
		if len(got) != len(expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("expected %v but got %v", expected, got)
			}
		}
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("IndexOf", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		value := 3
		expected := 2
		got := IndexOf(array, value)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("IndexOf", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		value := 5
		expected := -1
		got := IndexOf(array, value)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestLastIndexOf(t *testing.T) {
	t.Run("LastIndexOf", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 3}
		value := 3
		expected := 4
		got := LastIndexOf(array, value)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("LastIndexOf", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		value := 5
		expected := -1
		got := LastIndexOf(array, value)
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		expected := []int{4, 3, 2, 1}
		got := Reverse(array)
		if len(got) != len(expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("expected %v but got %v", expected, got)
			}
		}
	})
}
