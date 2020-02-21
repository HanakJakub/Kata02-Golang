package main

import "testing"

func TestChop(t *testing.T) {
	testData := []int{1, 5, 6}

	res := chop(1, testData)
	if res != 0 {
		t.Error("Chop function should found 1 on position 0")
	}

	res = chop(5, testData)
	if res != 1 {
		t.Error("Chop function should found 5 on position 1")
	}

	res = chop(6, testData)
	if res != 2 {
		t.Error("Chop function should found 6 on position 2")
	}

	res = chop(3, testData)
	if res != -1 {
		t.Error("Chop function should return -1 because 3 is not in testData")
	}

	res = chop(10, testData)
	if res != -1 {
		t.Error("Chop function should return -1 because 10 is not in testData")
	}
}
