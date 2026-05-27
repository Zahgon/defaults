package defaults

import (
	"errors"
	"reflect"
)

var (
	errInvalidType = errors.New("not a struct pointer")
)

const (
	fieldName = "default"
)

// Set initializes members in a struct referenced by a pointer.
// Maps and slices are initialized by `make` and other primitive types are set with default values.
// `ptr` should be a struct pointer
func Set(ptr interface{}) error { _ = "STUB: not implemented"; return nil }

// MustSet function is a wrapper of Set function
// It will call Set and panic if err not equals nil.
func MustSet(ptr interface{}) { _ = "STUB: not implemented"; return }

func setField(field reflect.Value, defaultVal string) error { _ = "STUB: not implemented"; return nil }

func unmarshalByInterface(field reflect.Value, defaultVal string) bool {
	_ = "STUB: not implemented"
	return false
}

// if field implements encode.TextUnmarshaler, try to use it before decode by kind

// if field implements json.Unmarshaler, try to use it before decode by kind

func isInitialValue(field reflect.Value) bool { _ = "STUB: not implemented"; return false }

func shouldInitializeField(field reflect.Value, tag string) bool {
	_ = "STUB: not implemented"
	return false
}

// CanUpdate returns true when the given value is an initial value of its type
func CanUpdate(v interface{}) bool { _ = "STUB: not implemented"; return false }
