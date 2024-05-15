// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: lannet/channel.proto

package lannetpb

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

// Validate checks the field values on HelloServer with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelloServer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelloServer with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelloServerMultiError, or
// nil if none found.
func (m *HelloServer) ValidateAll() error {
	return m.validate(true)
}

func (m *HelloServer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for MyEncodedAddress

	if len(errors) > 0 {
		return HelloServerMultiError(errors)
	}

	return nil
}

// HelloServerMultiError is an error wrapping multiple validation errors
// returned by HelloServer.ValidateAll() if the designated constraints aren't met.
type HelloServerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloServerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloServerMultiError) AllErrors() []error { return m }

// HelloServerValidationError is the validation error returned by
// HelloServer.Validate if the designated constraints aren't met.
type HelloServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloServerValidationError) ErrorName() string { return "HelloServerValidationError" }

// Error satisfies the builtin error interface
func (e HelloServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloServerValidationError{}

// Validate checks the field values on CenterRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CenterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CenterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CenterRequestMultiError, or
// nil if none found.
func (m *CenterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CenterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Req.(type) {
	case *CenterRequest_Hello:
		if v == nil {
			err := CenterRequestValidationError{
				field:  "Req",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetHello()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CenterRequestValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CenterRequestValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHello()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CenterRequestValidationError{
					field:  "Hello",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CenterRequest_Data:
		if v == nil {
			err := CenterRequestValidationError{
				field:  "Req",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Data
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return CenterRequestMultiError(errors)
	}

	return nil
}

// CenterRequestMultiError is an error wrapping multiple validation errors
// returned by CenterRequest.ValidateAll() if the designated constraints
// aren't met.
type CenterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CenterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CenterRequestMultiError) AllErrors() []error { return m }

// CenterRequestValidationError is the validation error returned by
// CenterRequest.Validate if the designated constraints aren't met.
type CenterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CenterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CenterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CenterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CenterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CenterRequestValidationError) ErrorName() string { return "CenterRequestValidationError" }

// Error satisfies the builtin error interface
func (e CenterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCenterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CenterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CenterRequestValidationError{}

// Validate checks the field values on StringList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StringList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StringList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StringListMultiError, or
// nil if none found.
func (m *StringList) ValidateAll() error {
	return m.validate(true)
}

func (m *StringList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StringListMultiError(errors)
	}

	return nil
}

// StringListMultiError is an error wrapping multiple validation errors
// returned by StringList.ValidateAll() if the designated constraints aren't met.
type StringListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StringListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StringListMultiError) AllErrors() []error { return m }

// StringListValidationError is the validation error returned by
// StringList.Validate if the designated constraints aren't met.
type StringListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringListValidationError) ErrorName() string { return "StringListValidationError" }

// Error satisfies the builtin error interface
func (e StringListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringListValidationError{}

// Validate checks the field values on HelloClient with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelloClient) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelloClient with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelloClientMultiError, or
// nil if none found.
func (m *HelloClient) ValidateAll() error {
	return m.validate(true)
}

func (m *HelloClient) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Clients

	{
		sorted_keys := make([]string, len(m.GetClientIps()))
		i := 0
		for key := range m.GetClientIps() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetClientIps()[key]
			_ = val

			// no validation rules for ClientIps[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, HelloClientValidationError{
							field:  fmt.Sprintf("ClientIps[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, HelloClientValidationError{
							field:  fmt.Sprintf("ClientIps[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return HelloClientValidationError{
						field:  fmt.Sprintf("ClientIps[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for Vpn

	if len(errors) > 0 {
		return HelloClientMultiError(errors)
	}

	return nil
}

// HelloClientMultiError is an error wrapping multiple validation errors
// returned by HelloClient.ValidateAll() if the designated constraints aren't met.
type HelloClientMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloClientMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloClientMultiError) AllErrors() []error { return m }

// HelloClientValidationError is the validation error returned by
// HelloClient.Validate if the designated constraints aren't met.
type HelloClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloClientValidationError) ErrorName() string { return "HelloClientValidationError" }

// Error satisfies the builtin error interface
func (e HelloClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloClient.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloClientValidationError{}

// Validate checks the field values on CenterResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CenterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CenterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CenterResponseMultiError,
// or nil if none found.
func (m *CenterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CenterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Resp.(type) {
	case *CenterResponse_Hello:
		if v == nil {
			err := CenterResponseValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetHello()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CenterResponseValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CenterResponseValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHello()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CenterResponseValidationError{
					field:  "Hello",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CenterResponse_Data:
		if v == nil {
			err := CenterResponseValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Data
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return CenterResponseMultiError(errors)
	}

	return nil
}

// CenterResponseMultiError is an error wrapping multiple validation errors
// returned by CenterResponse.ValidateAll() if the designated constraints
// aren't met.
type CenterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CenterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CenterResponseMultiError) AllErrors() []error { return m }

// CenterResponseValidationError is the validation error returned by
// CenterResponse.Validate if the designated constraints aren't met.
type CenterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CenterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CenterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CenterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CenterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CenterResponseValidationError) ErrorName() string { return "CenterResponseValidationError" }

// Error satisfies the builtin error interface
func (e CenterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCenterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CenterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CenterResponseValidationError{}

// Validate checks the field values on DirectNetHelloServer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DirectNetHelloServer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DirectNetHelloServer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DirectNetHelloServerMultiError, or nil if none found.
func (m *DirectNetHelloServer) ValidateAll() error {
	return m.validate(true)
}

func (m *DirectNetHelloServer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return DirectNetHelloServerMultiError(errors)
	}

	return nil
}

// DirectNetHelloServerMultiError is an error wrapping multiple validation
// errors returned by DirectNetHelloServer.ValidateAll() if the designated
// constraints aren't met.
type DirectNetHelloServerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DirectNetHelloServerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DirectNetHelloServerMultiError) AllErrors() []error { return m }

// DirectNetHelloServerValidationError is the validation error returned by
// DirectNetHelloServer.Validate if the designated constraints aren't met.
type DirectNetHelloServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DirectNetHelloServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DirectNetHelloServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DirectNetHelloServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DirectNetHelloServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DirectNetHelloServerValidationError) ErrorName() string {
	return "DirectNetHelloServerValidationError"
}

// Error satisfies the builtin error interface
func (e DirectNetHelloServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDirectNetHelloServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DirectNetHelloServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DirectNetHelloServerValidationError{}

// Validate checks the field values on DirectNetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DirectNetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DirectNetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DirectNetRequestMultiError, or nil if none found.
func (m *DirectNetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DirectNetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Req.(type) {
	case *DirectNetRequest_Hello:
		if v == nil {
			err := DirectNetRequestValidationError{
				field:  "Req",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetHello()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DirectNetRequestValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DirectNetRequestValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHello()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DirectNetRequestValidationError{
					field:  "Hello",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DirectNetRequest_Data:
		if v == nil {
			err := DirectNetRequestValidationError{
				field:  "Req",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Data
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return DirectNetRequestMultiError(errors)
	}

	return nil
}

// DirectNetRequestMultiError is an error wrapping multiple validation errors
// returned by DirectNetRequest.ValidateAll() if the designated constraints
// aren't met.
type DirectNetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DirectNetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DirectNetRequestMultiError) AllErrors() []error { return m }

// DirectNetRequestValidationError is the validation error returned by
// DirectNetRequest.Validate if the designated constraints aren't met.
type DirectNetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DirectNetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DirectNetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DirectNetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DirectNetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DirectNetRequestValidationError) ErrorName() string { return "DirectNetRequestValidationError" }

// Error satisfies the builtin error interface
func (e DirectNetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDirectNetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DirectNetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DirectNetRequestValidationError{}

// Validate checks the field values on DirectNetHelloClient with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DirectNetHelloClient) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DirectNetHelloClient with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DirectNetHelloClientMultiError, or nil if none found.
func (m *DirectNetHelloClient) ValidateAll() error {
	return m.validate(true)
}

func (m *DirectNetHelloClient) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DirectNetHelloClientMultiError(errors)
	}

	return nil
}

// DirectNetHelloClientMultiError is an error wrapping multiple validation
// errors returned by DirectNetHelloClient.ValidateAll() if the designated
// constraints aren't met.
type DirectNetHelloClientMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DirectNetHelloClientMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DirectNetHelloClientMultiError) AllErrors() []error { return m }

// DirectNetHelloClientValidationError is the validation error returned by
// DirectNetHelloClient.Validate if the designated constraints aren't met.
type DirectNetHelloClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DirectNetHelloClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DirectNetHelloClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DirectNetHelloClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DirectNetHelloClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DirectNetHelloClientValidationError) ErrorName() string {
	return "DirectNetHelloClientValidationError"
}

// Error satisfies the builtin error interface
func (e DirectNetHelloClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDirectNetHelloClient.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DirectNetHelloClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DirectNetHelloClientValidationError{}

// Validate checks the field values on DirectNetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DirectNetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DirectNetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DirectNetResponseMultiError, or nil if none found.
func (m *DirectNetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DirectNetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Resp.(type) {
	case *DirectNetResponse_Hello:
		if v == nil {
			err := DirectNetResponseValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetHello()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DirectNetResponseValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DirectNetResponseValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHello()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DirectNetResponseValidationError{
					field:  "Hello",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DirectNetResponse_Data:
		if v == nil {
			err := DirectNetResponseValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Data
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return DirectNetResponseMultiError(errors)
	}

	return nil
}

// DirectNetResponseMultiError is an error wrapping multiple validation errors
// returned by DirectNetResponse.ValidateAll() if the designated constraints
// aren't met.
type DirectNetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DirectNetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DirectNetResponseMultiError) AllErrors() []error { return m }

// DirectNetResponseValidationError is the validation error returned by
// DirectNetResponse.Validate if the designated constraints aren't met.
type DirectNetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DirectNetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DirectNetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DirectNetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DirectNetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DirectNetResponseValidationError) ErrorName() string {
	return "DirectNetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DirectNetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDirectNetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DirectNetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DirectNetResponseValidationError{}
