package types_test

import (
	"testing"

	"github.com/rsdate/utils/types"
)

func TestCast(t *testing.T) {
	var value any = "1"
	var typ string = "int"
	val := types.Cast(value, true, typ)
	if val != 1 {
		t.Errorf("Cast() failed, expected 1, got %v", val)
	}
	value = 1
	typ = "string"
	val = types.Cast(value, true, typ)
	if val != "1" {
		t.Errorf("Cast() failed, expected '1', got %v", val)
	}
	value = 1
	typ = "float32"
	val = types.Cast(value, true, typ)
	if val != float32(1) {
		t.Errorf("Cast() failed, expected 1.0, got %v", val)
	}
	value = 1
	typ = "float64"
	val = types.Cast(value, true, typ)
	if val != 1.0 {
		t.Errorf("Cast() failed, expected 1.0, got %v", val)
	}
	typ = "unknown"
	val = types.Cast(value, true, typ)
	if val != "casting error: unknown type" {
		t.Errorf("Cast() failed, expected 'casting error: unknown type', got %v", val)
	}
}
