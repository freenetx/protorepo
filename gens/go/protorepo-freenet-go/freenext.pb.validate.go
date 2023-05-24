// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/freenet/freenext.proto

package freenetpb

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

// Validate checks the field values on SyncMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SyncMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SyncMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SyncMessageMultiError, or
// nil if none found.
func (m *SyncMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *SyncMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for KnownClients

	if len(errors) > 0 {
		return SyncMessageMultiError(errors)
	}

	return nil
}

// SyncMessageMultiError is an error wrapping multiple validation errors
// returned by SyncMessage.ValidateAll() if the designated constraints aren't met.
type SyncMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SyncMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SyncMessageMultiError) AllErrors() []error { return m }

// SyncMessageValidationError is the validation error returned by
// SyncMessage.Validate if the designated constraints aren't met.
type SyncMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SyncMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SyncMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SyncMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SyncMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SyncMessageValidationError) ErrorName() string { return "SyncMessageValidationError" }

// Error satisfies the builtin error interface
func (e SyncMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSyncMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SyncMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SyncMessageValidationError{}

// Validate checks the field values on MessageAck with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MessageAck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageAck with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MessageAckMultiError, or
// nil if none found.
func (m *MessageAck) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageAck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return MessageAckMultiError(errors)
	}

	return nil
}

// MessageAckMultiError is an error wrapping multiple validation errors
// returned by MessageAck.ValidateAll() if the designated constraints aren't met.
type MessageAckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageAckMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageAckMultiError) AllErrors() []error { return m }

// MessageAckValidationError is the validation error returned by
// MessageAck.Validate if the designated constraints aren't met.
type MessageAckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageAckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageAckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageAckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageAckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageAckValidationError) ErrorName() string { return "MessageAckValidationError" }

// Error satisfies the builtin error interface
func (e MessageAckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageAck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageAckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageAckValidationError{}

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

// Validate checks the field values on ClientInitMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ClientInitMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientInitMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClientInitMessageMultiError, or nil if none found.
func (m *ClientInitMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientInitMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ClientAddress

	if len(errors) > 0 {
		return ClientInitMessageMultiError(errors)
	}

	return nil
}

// ClientInitMessageMultiError is an error wrapping multiple validation errors
// returned by ClientInitMessage.ValidateAll() if the designated constraints
// aren't met.
type ClientInitMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientInitMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientInitMessageMultiError) AllErrors() []error { return m }

// ClientInitMessageValidationError is the validation error returned by
// ClientInitMessage.Validate if the designated constraints aren't met.
type ClientInitMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientInitMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientInitMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientInitMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientInitMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientInitMessageValidationError) ErrorName() string {
	return "ClientInitMessageValidationError"
}

// Error satisfies the builtin error interface
func (e ClientInitMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientInitMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientInitMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientInitMessageValidationError{}

// Validate checks the field values on NeedClientInitMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NeedClientInitMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NeedClientInitMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NeedClientInitMessageMultiError, or nil if none found.
func (m *NeedClientInitMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *NeedClientInitMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return NeedClientInitMessageMultiError(errors)
	}

	return nil
}

// NeedClientInitMessageMultiError is an error wrapping multiple validation
// errors returned by NeedClientInitMessage.ValidateAll() if the designated
// constraints aren't met.
type NeedClientInitMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NeedClientInitMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NeedClientInitMessageMultiError) AllErrors() []error { return m }

// NeedClientInitMessageValidationError is the validation error returned by
// NeedClientInitMessage.Validate if the designated constraints aren't met.
type NeedClientInitMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NeedClientInitMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NeedClientInitMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NeedClientInitMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NeedClientInitMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NeedClientInitMessageValidationError) ErrorName() string {
	return "NeedClientInitMessageValidationError"
}

// Error satisfies the builtin error interface
func (e NeedClientInitMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNeedClientInitMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NeedClientInitMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NeedClientInitMessageValidationError{}

// Validate checks the field values on ClientMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClientMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClientMessageMultiError, or
// nil if none found.
func (m *ClientMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Request.(type) {
	case *ClientMessage_Data:
		if v == nil {
			err := ClientMessageValidationError{
				field:  "Request",
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
					errors = append(errors, ClientMessageValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClientMessageValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientMessageValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ClientMessage_ClientInitMessage:
		if v == nil {
			err := ClientMessageValidationError{
				field:  "Request",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetClientInitMessage()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClientMessageValidationError{
						field:  "ClientInitMessage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClientMessageValidationError{
						field:  "ClientInitMessage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetClientInitMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientMessageValidationError{
					field:  "ClientInitMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	switch v := m.Response.(type) {
	case *ClientMessage_MessageAck:
		if v == nil {
			err := ClientMessageValidationError{
				field:  "Response",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetMessageAck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClientMessageValidationError{
						field:  "MessageAck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClientMessageValidationError{
						field:  "MessageAck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMessageAck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientMessageValidationError{
					field:  "MessageAck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ClientMessageMultiError(errors)
	}

	return nil
}

// ClientMessageMultiError is an error wrapping multiple validation errors
// returned by ClientMessage.ValidateAll() if the designated constraints
// aren't met.
type ClientMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientMessageMultiError) AllErrors() []error { return m }

// ClientMessageValidationError is the validation error returned by
// ClientMessage.Validate if the designated constraints aren't met.
type ClientMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientMessageValidationError) ErrorName() string { return "ClientMessageValidationError" }

// Error satisfies the builtin error interface
func (e ClientMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientMessageValidationError{}

// Validate checks the field values on ServerMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServerMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServerMessageMultiError, or
// nil if none found.
func (m *ServerMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Request.(type) {
	case *ServerMessage_Data:
		if v == nil {
			err := ServerMessageValidationError{
				field:  "Request",
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
					errors = append(errors, ServerMessageValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerMessageValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerMessageValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ServerMessage_SyncMessage:
		if v == nil {
			err := ServerMessageValidationError{
				field:  "Request",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSyncMessage()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerMessageValidationError{
						field:  "SyncMessage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerMessageValidationError{
						field:  "SyncMessage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSyncMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerMessageValidationError{
					field:  "SyncMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ServerMessage_NeedClientInitMessage:
		if v == nil {
			err := ServerMessageValidationError{
				field:  "Request",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetNeedClientInitMessage()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerMessageValidationError{
						field:  "NeedClientInitMessage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerMessageValidationError{
						field:  "NeedClientInitMessage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetNeedClientInitMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerMessageValidationError{
					field:  "NeedClientInitMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	switch v := m.Response.(type) {
	case *ServerMessage_MessageAck:
		if v == nil {
			err := ServerMessageValidationError{
				field:  "Response",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetMessageAck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerMessageValidationError{
						field:  "MessageAck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerMessageValidationError{
						field:  "MessageAck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMessageAck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerMessageValidationError{
					field:  "MessageAck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ServerMessageMultiError(errors)
	}

	return nil
}

// ServerMessageMultiError is an error wrapping multiple validation errors
// returned by ServerMessage.ValidateAll() if the designated constraints
// aren't met.
type ServerMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerMessageMultiError) AllErrors() []error { return m }

// ServerMessageValidationError is the validation error returned by
// ServerMessage.Validate if the designated constraints aren't met.
type ServerMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerMessageValidationError) ErrorName() string { return "ServerMessageValidationError" }

// Error satisfies the builtin error interface
func (e ServerMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerMessageValidationError{}
