package utils

import (
	"fmt"
	"testing"
)

func TestFindBySlice(t *testing.T) {
	numbers := []int{1, 2, 3, 5, 7, 10, 20, 11}

	result, isExist := FindBySlice[int](numbers, func(value int, _ int) bool {
		return value == 3
	})

	if !isExist || result != 3 {
		t.Errorf("FindBySlice fail, result: %d, isExist: %v\n", result, isExist)
	}

	t.Logf("FindBySlice success, result: %d, isExist: %v\n", result, isExist)
}

func TestMapBySlice(t *testing.T) {
	numbers1, numbers2 := []int{1, 2, 3, 4}, []int{10, 20, 30, 40}

	numbers3 := MapBySlice[int, int](numbers1, func(value int, _ int) int {
		return value * 10
	})

	for index, value := range numbers3 {
		if numbers2[index] != value {
			t.Errorf("MapBySlice fail, numbers2: %d, numbers3: %d\n", numbers2[index], value)
		}
	}

	t.Logf("MapBySlice success, result: %v", numbers3)
}

func TestForEachBySlice(t *testing.T) {
	numbers, result := []int{1, 2, 3, 4, 5}, 0

	ForEachBySlice[int](numbers, func(_ int, _ int) {
		result += 1
	})

	if result != 5 {
		t.Errorf("EachBySlice fail, result: %v", result)
		return
	}
	t.Logf("EachBySlice success, result: %v", result)
}

func TestFilterBySlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	results := FilterBySlice[int](numbers, func(value, index int) bool {
		return value > 4
	})

	if len(results) != 1 && results[0] != 5 {
		t.Errorf("FilterBySlice fail, results: %v", results)
		return
	}
	t.Logf("FilterBySlice success, results: %v", results)
}

func TestIndexOfBySlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	result := IndexOfBySlice[int](numbers, 3)

	if result != 2 {
		t.Errorf("IndexOfBySlice fail, result: %d", result)
		return
	}
	t.Logf("IndexOfBySlice success, result: %d", result)
}

func TestIsExistBySlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	isExist := IsExistBySlice[int](numbers, 3)

	if !isExist {
		t.Errorf("IsExistBySlice fail, result: %v", isExist)
		return
	}
	t.Logf("IsExistBySlice success, result: %v", isExist)
}

func TestAtBySlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	result, isExist := AtBySlice[int](numbers, 1)

	if !isExist || result != 2 {
		t.Errorf("AtBySlice fail, result: %d, isExist: %v\n", result, isExist)
		return
	}
	t.Logf("AtBySlice success, result: %d, isExist: %v\n", result, isExist)
}

func TestConcatBySlice(t *testing.T) {
	numbers1, numbers2 := []int{1, 2, 3, 4, 5}, []int{6, 7, 8}

	numbers3 := ConcatBySlice[int](numbers1, numbers2)

	if len(numbers3) != 8 {
		t.Errorf("ConcatBySlice fail, numbers3: %v\n", numbers3)
		return
	}
	t.Logf("ConcatBySlice success, numbers3: %v\n", numbers3)
}

type myString string

func (ms myString) ToString() string {
	return fmt.Sprintf("%v", ms)
}

func TestJoinBySlice(t *testing.T) {
	result := JoinBySlice[myString]([]myString{"hello", "world"}, ", ")

	if result != "hello, world" {
		t.Errorf("JoinBySlice fail, result: %s\n", result)
		return
	}
	t.Logf("JoinBySlice success, result: %s\n", result)
}
