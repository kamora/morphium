package morphium

import (
	"reflect"
	"testing"
)

func TestNumbers(t *testing.T) {

	value1 := "0b100"
	value2 := "0o010"
	value3 := "0x0A1"
	value4 := "01000"
	value5 := "-50000"
	value6 := "5.100"
	value7 := "999.11"
	value8 := "TRUE"
	value9 := "F"
	value10 := "3.14+2.71i"
	value11 := "3.14+2.71i"

	var (
		err error
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64
		f32 float32
		f64 float64
		b1  bool
		b2  bool
		c1  complex64
		c2  complex128
	)

	i, err = Morph[int](value1)
	if err != nil {
		t.Error(err)
	}

	if i != 4 {
		t.Errorf("got %d, want 4", i)
	}

	if reflect.TypeOf(i).Kind() != reflect.Int {
		t.Errorf("got %s, want int", reflect.TypeOf(i).Name())
	}

	i8, err = Morph[int8](value2)
	if err != nil {
		t.Error(err)
	}

	if i8 != 8 {
		t.Errorf("got %d, want 8", i)
	}

	if reflect.TypeOf(i8).Kind() != reflect.Int8 {
		t.Errorf("got %s, want int8", reflect.TypeOf(i8).Name())
	}

	i16, err = Morph[int16](value3)
	if err != nil {
		t.Error(err)
	}

	if i16 != 161 {
		t.Errorf("got %d, want 161", i16)
	}

	if reflect.TypeOf(i16).Kind() != reflect.Int16 {
		t.Errorf("got %s, want int16", reflect.TypeOf(i16).Name())
	}

	i32, err = Morph[int32](value4)
	if err != nil {
		t.Error(err)
	}

	if i32 != 1000 {
		t.Errorf("got %d, want 1000", i32)
	}

	if reflect.TypeOf(i32).Kind() != reflect.Int32 {
		t.Errorf("got %s, want int32", reflect.TypeOf(i32).Name())
	}

	i64, err = Morph[int64](value5)
	if err != nil {
		t.Error(err)
	}

	if i64 != -50000 {
		t.Errorf("got %d, want -50000", i64)
	}

	if reflect.TypeOf(i64).Kind() != reflect.Int64 {
		t.Errorf("got %s, want int64", reflect.TypeOf(i64).Name())
	}

	f32, err = Morph[float32](value6)
	if err != nil {
		t.Error(err)
	}

	if f32 != 5.1 {
		t.Errorf("got %f, want 5.1", f32)
	}

	if reflect.TypeOf(f32).Kind() != reflect.Float32 {
		t.Errorf("got %s, want float32", reflect.TypeOf(f32).Name())
	}

	f64, err = Morph[float64](value7)
	if err != nil {
		t.Error(err)
	}

	if f64 != 999.11 {
		t.Errorf("got %f, want 999.11", f64)
	}

	if reflect.TypeOf(f64).Kind() != reflect.Float64 {
		t.Errorf("got %s, want float64", reflect.TypeOf(f64).Name())
	}

	b1, err = Morph[bool](value8)
	if err != nil {
		t.Error(err)
	}

	if b1 != true {
		t.Errorf("got %v, want true", b1)
	}

	if reflect.TypeOf(b1).Kind() != reflect.Bool {
		t.Errorf("got %s, want bool", reflect.TypeOf(b1).Name())
	}

	b2, err = Morph[bool](value9)
	if err != nil {
		t.Error(err)
	}

	if b2 != false {
		t.Errorf("got %v, want false", b2)
	}

	if reflect.TypeOf(b2).Kind() != reflect.Bool {
		t.Errorf("got %s, want bool", reflect.TypeOf(b2).Name())
	}

	c1, err = Morph[complex64](value10)
	if err != nil {
		t.Error(err)
	}

	if c1 != complex64(3.14+2.71i) {
		t.Errorf("got %v, want 3.14+2.71i", c1)
	}

	if reflect.TypeOf(c1).Kind() != reflect.Complex64 {
		t.Errorf("got %s, want bool", reflect.TypeOf(c1).Name())
	}

	c2, err = Morph[complex128](value11)
	if err != nil {
		t.Error(err)
	}

	if c2 != 3.14+2.71i {
		t.Errorf("got %v, want 3.14+2.71i", c2)
	}

	if reflect.TypeOf(c2).Kind() != reflect.Complex128 {
		t.Errorf("got %s, want bool", reflect.TypeOf(c2).Name())
	}
}
