package adapter

import "testing"

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	res := adapter.Request()
	if res != expect {
		t.Fatal("expect: %s, actual: %s", expect, res)
	}
}
