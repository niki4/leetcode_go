package main

import "testing"

func createMinStackSample() MinStack {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	return obj
}

func TestMinStack_Push(t *testing.T) {
	ms := createMinStackSample()
	if len(ms.values) != 3 {
		t.Errorf("Expected len 3, got %d: %v", len(ms.values), ms.values)
	}
	if ms.minIdx != 2 {
		t.Errorf("Expected idx 2 of min val -3, got %d", ms.minIdx)
	}
}

func TestMinStack_Pop(t *testing.T) {
	ms := createMinStackSample()
	ms.Pop() // discard 0 from top
	if len(ms.values) != 2 {
		t.Errorf("Expected len 2, got %d: %v", len(ms.values), ms.values)
	}
	if ms.minIdx != 0 {
		t.Errorf("Expected idx 0 of min val -3, got %d", ms.minIdx)
	}
}

func TestMinStack_Top(t *testing.T) {
	ms := createMinStackSample()
	if x := ms.Top(); x != -3 { // last added num was -3
		t.Errorf("Expected %d, got %d", 0, x)
	}
}

func TestMinStack_GetMin(t *testing.T) {
	ms := createMinStackSample()
	if x := ms.GetMin(); x != -3 {
		t.Errorf("Expected %d, got %d", -3, x)
	}
}
