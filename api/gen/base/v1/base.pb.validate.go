// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: base/v1/base.proto

package basev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on BaseOrder with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BaseOrder) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BaseOrder with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BaseOrderMultiError, or nil
// if none found.
func (m *BaseOrder) ValidateAll() error {
	return m.validate(true)
}

func (m *BaseOrder) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sort

	// no validation rules for Sidx

	if len(errors) > 0 {
		return BaseOrderMultiError(errors)
	}

	return nil
}

// BaseOrderMultiError is an error wrapping multiple validation errors returned
// by BaseOrder.ValidateAll() if the designated constraints aren't met.
type BaseOrderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BaseOrderMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BaseOrderMultiError) AllErrors() []error { return m }

// BaseOrderValidationError is the validation error returned by
// BaseOrder.Validate if the designated constraints aren't met.
type BaseOrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BaseOrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BaseOrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BaseOrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BaseOrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BaseOrderValidationError) ErrorName() string { return "BaseOrderValidationError" }

// Error satisfies the builtin error interface
func (e BaseOrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBaseOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BaseOrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BaseOrderValidationError{}

// Validate checks the field values on WhereExt with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WhereExt) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WhereExt with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WhereExtMultiError, or nil
// if none found.
func (m *WhereExt) ValidateAll() error {
	return m.validate(true)
}

func (m *WhereExt) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Column

	// no validation rules for Symbol

	switch v := m.Val.(type) {
	case *WhereExt_StrVal:
		if v == nil {
			err := WhereExtValidationError{
				field:  "Val",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for StrVal
	case *WhereExt_IntVal:
		if v == nil {
			err := WhereExtValidationError{
				field:  "Val",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for IntVal
	case *WhereExt_BoolVal:
		if v == nil {
			err := WhereExtValidationError{
				field:  "Val",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for BoolVal
	case *WhereExt_DoubleVal:
		if v == nil {
			err := WhereExtValidationError{
				field:  "Val",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for DoubleVal
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return WhereExtMultiError(errors)
	}

	return nil
}

// WhereExtMultiError is an error wrapping multiple validation errors returned
// by WhereExt.ValidateAll() if the designated constraints aren't met.
type WhereExtMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WhereExtMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WhereExtMultiError) AllErrors() []error { return m }

// WhereExtValidationError is the validation error returned by
// WhereExt.Validate if the designated constraints aren't met.
type WhereExtValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WhereExtValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WhereExtValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WhereExtValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WhereExtValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WhereExtValidationError) ErrorName() string { return "WhereExtValidationError" }

// Error satisfies the builtin error interface
func (e WhereExtValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWhereExt.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WhereExtValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WhereExtValidationError{}

// Validate checks the field values on BaseList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BaseList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BaseList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BaseListMultiError, or nil
// if none found.
func (m *BaseList) ValidateAll() error {
	return m.validate(true)
}

func (m *BaseList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Size

	// no validation rules for Sort

	// no validation rules for Sidx

	for idx, item := range m.GetOrder() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BaseListValidationError{
						field:  fmt.Sprintf("Order[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BaseListValidationError{
						field:  fmt.Sprintf("Order[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BaseListValidationError{
					field:  fmt.Sprintf("Order[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetWhereExt() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BaseListValidationError{
						field:  fmt.Sprintf("WhereExt[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BaseListValidationError{
						field:  fmt.Sprintf("WhereExt[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BaseListValidationError{
					field:  fmt.Sprintf("WhereExt[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetWhereLike() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BaseListValidationError{
						field:  fmt.Sprintf("WhereLike[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BaseListValidationError{
						field:  fmt.Sprintf("WhereLike[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BaseListValidationError{
					field:  fmt.Sprintf("WhereLike[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BaseListMultiError(errors)
	}

	return nil
}

// BaseListMultiError is an error wrapping multiple validation errors returned
// by BaseList.ValidateAll() if the designated constraints aren't met.
type BaseListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BaseListMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BaseListMultiError) AllErrors() []error { return m }

// BaseListValidationError is the validation error returned by
// BaseList.Validate if the designated constraints aren't met.
type BaseListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BaseListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BaseListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BaseListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BaseListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BaseListValidationError) ErrorName() string { return "BaseListValidationError" }

// Error satisfies the builtin error interface
func (e BaseListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBaseList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BaseListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BaseListValidationError{}
