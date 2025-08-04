package morphium

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Morphable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~bool | ~string | ~complex64 | ~complex128
}

func Morph[T Morphable](value string) (T, error) {
	var def T
	switch reflect.TypeFor[T]().Kind() {

	case reflect.Int:
		converted, err := strconv.ParseInt(value, base(&value), 0)
		if err != nil {
			return def, err
		}

		return any(int(converted)).(T), nil

	case reflect.Int8:
		converted, err := strconv.ParseInt(value, base(&value), 8)

		if err != nil {
			return def, err
		}

		return any(int8(converted)).(T), nil

	case reflect.Int16:
		converted, err := strconv.ParseInt(value, base(&value), 16)

		if err != nil {
			return def, err
		}

		return any(int16(converted)).(T), nil

	case reflect.Int32:
		converted, err := strconv.ParseInt(value, base(&value), 32)

		if err != nil {
			return def, err
		}

		return any(int32(converted)).(T), nil

	case reflect.Int64:
		converted, err := strconv.ParseInt(value, base(&value), 64)

		if err != nil {
			return def, err
		}

		return any(converted).(T), nil

	case reflect.Uint:
		converted, err := strconv.ParseUint(value, base(&value), 0)
		if err != nil {
			return def, err
		}

		return any(uint(converted)).(T), nil

	case reflect.Uint8:
		converted, err := strconv.ParseUint(value, base(&value), 0)
		if err != nil {
			return def, err
		}

		return any(uint8(converted)).(T), nil

	case reflect.Uint16:
		converted, err := strconv.ParseUint(value, base(&value), 0)
		if err != nil {
			return def, err
		}

		return any(uint16(converted)).(T), nil

	case reflect.Uint32:
		converted, err := strconv.ParseUint(value, base(&value), 0)
		if err != nil {
			return def, err
		}

		return any(uint32(converted)).(T), nil

	case reflect.Uint64:
		converted, err := strconv.ParseUint(value, base(&value), 0)
		if err != nil {
			return def, err
		}

		return any(converted).(T), nil

	case reflect.Float32:
		converted, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return def, err
		}

		return any(float32(converted)).(T), nil

	case reflect.Float64:
		converted, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return def, err
		}

		return any(converted).(T), nil

	case reflect.Complex64:
		converted, err := strconv.ParseComplex(value, 64)
		if err != nil {
			return def, err
		}

		return any(complex64(converted)).(T), nil

	case reflect.Complex128:
		converted, err := strconv.ParseComplex(value, 128)
		if err != nil {
			return def, err
		}

		return any(converted).(T), nil

	case reflect.Bool:
		converted, err := strconv.ParseBool(value)
		if err != nil {
			return def, err
		}

		return any(converted).(T), nil

	case reflect.String:
		return any(value).(T), nil

	default:
		return def, fmt.Errorf("unsupported conversion for %s", reflect.TypeFor[T]().Name())
	}
}

func base(value *string) int {
	var ok bool

	if *value, ok = strings.CutPrefix(*value, "0x"); ok {
		return 16
	}

	if *value, ok = strings.CutPrefix(*value, "0X"); ok {
		return 16
	}

	if *value, ok = strings.CutPrefix(*value, "0o"); ok {
		return 8
	}

	if *value, ok = strings.CutPrefix(*value, "0O"); ok {
		return 8
	}

	if *value, ok = strings.CutPrefix(*value, "0b"); ok {
		return 2
	}

	if *value, ok = strings.CutPrefix(*value, "0B"); ok {
		return 2
	}

	return 10
}
