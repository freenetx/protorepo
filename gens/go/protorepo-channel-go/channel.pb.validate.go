// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/channel/channel.proto

package channelpb

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

// Validate checks the field values on SetKey with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SetKey) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetKey with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SetKeyMultiError, or nil if none found.
func (m *SetKey) ValidateAll() error {
	return m.validate(true)
}

func (m *SetKey) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SelfKey

	// no validation rules for DestinationKey

	if len(errors) > 0 {
		return SetKeyMultiError(errors)
	}

	return nil
}

// SetKeyMultiError is an error wrapping multiple validation errors returned by
// SetKey.ValidateAll() if the designated constraints aren't met.
type SetKeyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetKeyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetKeyMultiError) AllErrors() []error { return m }

// SetKeyValidationError is the validation error returned by SetKey.Validate if
// the designated constraints aren't met.
type SetKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetKeyValidationError) ErrorName() string { return "SetKeyValidationError" }

// Error satisfies the builtin error interface
func (e SetKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetKeyValidationError{}

// Validate checks the field values on Data with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DataMultiError, or nil if none found.
func (m *Data) ValidateAll() error {
	return m.validate(true)
}

func (m *Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SrcKey

	// no validation rules for DstKey

	// no validation rules for Data

	if len(errors) > 0 {
		return DataMultiError(errors)
	}

	return nil
}

// DataMultiError is an error wrapping multiple validation errors returned by
// Data.ValidateAll() if the designated constraints aren't met.
type DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataMultiError) AllErrors() []error { return m }

// DataValidationError is the validation error returned by Data.Validate if the
// designated constraints aren't met.
type DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataValidationError) ErrorName() string { return "DataValidationError" }

// Error satisfies the builtin error interface
func (e DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataValidationError{}

// Validate checks the field values on ChannelRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChannelRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChannelRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChannelRequestMultiError,
// or nil if none found.
func (m *ChannelRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChannelRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Req.(type) {
	case *ChannelRequest_SetKey:
		if v == nil {
			err := ChannelRequestValidationError{
				field:  "Req",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSetKey()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ChannelRequestValidationError{
						field:  "SetKey",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ChannelRequestValidationError{
						field:  "SetKey",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSetKey()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ChannelRequestValidationError{
					field:  "SetKey",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ChannelRequest_Data:
		if v == nil {
			err := ChannelRequestValidationError{
				field:  "Req",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetData()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ChannelRequestValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ChannelRequestValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ChannelRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ChannelRequestMultiError(errors)
	}

	return nil
}

// ChannelRequestMultiError is an error wrapping multiple validation errors
// returned by ChannelRequest.ValidateAll() if the designated constraints
// aren't met.
type ChannelRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChannelRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChannelRequestMultiError) AllErrors() []error { return m }

// ChannelRequestValidationError is the validation error returned by
// ChannelRequest.Validate if the designated constraints aren't met.
type ChannelRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChannelRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChannelRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChannelRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChannelRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChannelRequestValidationError) ErrorName() string { return "ChannelRequestValidationError" }

// Error satisfies the builtin error interface
func (e ChannelRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChannelRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChannelRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChannelRequestValidationError{}

// Validate checks the field values on KeyAccepted with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KeyAccepted) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KeyAccepted with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KeyAcceptedMultiError, or
// nil if none found.
func (m *KeyAccepted) ValidateAll() error {
	return m.validate(true)
}

func (m *KeyAccepted) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return KeyAcceptedMultiError(errors)
	}

	return nil
}

// KeyAcceptedMultiError is an error wrapping multiple validation errors
// returned by KeyAccepted.ValidateAll() if the designated constraints aren't met.
type KeyAcceptedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KeyAcceptedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KeyAcceptedMultiError) AllErrors() []error { return m }

// KeyAcceptedValidationError is the validation error returned by
// KeyAccepted.Validate if the designated constraints aren't met.
type KeyAcceptedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyAcceptedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyAcceptedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyAcceptedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyAcceptedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyAcceptedValidationError) ErrorName() string { return "KeyAcceptedValidationError" }

// Error satisfies the builtin error interface
func (e KeyAcceptedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyAccepted.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyAcceptedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyAcceptedValidationError{}

// Validate checks the field values on ReceiveResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReceiveResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveResponseMultiError, or nil if none found.
func (m *ReceiveResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Resp.(type) {
	case *ReceiveResponse_KeyAccepted:
		if v == nil {
			err := ReceiveResponseValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetKeyAccepted()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReceiveResponseValidationError{
						field:  "KeyAccepted",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReceiveResponseValidationError{
						field:  "KeyAccepted",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetKeyAccepted()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiveResponseValidationError{
					field:  "KeyAccepted",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ReceiveResponse_Data:
		if v == nil {
			err := ReceiveResponseValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetData()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReceiveResponseValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReceiveResponseValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiveResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ReceiveResponseMultiError(errors)
	}

	return nil
}

// ReceiveResponseMultiError is an error wrapping multiple validation errors
// returned by ReceiveResponse.ValidateAll() if the designated constraints
// aren't met.
type ReceiveResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveResponseMultiError) AllErrors() []error { return m }

// ReceiveResponseValidationError is the validation error returned by
// ReceiveResponse.Validate if the designated constraints aren't met.
type ReceiveResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveResponseValidationError) ErrorName() string { return "ReceiveResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReceiveResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveResponseValidationError{}
