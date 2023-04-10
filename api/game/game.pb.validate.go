// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: game-proto/game.proto

package game

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

// Validate checks the field values on Page with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Page) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Page with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PageMultiError, or nil if none found.
func (m *Page) ValidateAll() error {
	return m.validate(true)
}

func (m *Page) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Num

	// no validation rules for Size

	// no validation rules for Total

	// no validation rules for Disable

	if len(errors) > 0 {
		return PageMultiError(errors)
	}

	return nil
}

// PageMultiError is an error wrapping multiple validation errors returned by
// Page.ValidateAll() if the designated constraints aren't met.
type PageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageMultiError) AllErrors() []error { return m }

// PageValidationError is the validation error returned by Page.Validate if the
// designated constraints aren't met.
type PageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageValidationError) ErrorName() string { return "PageValidationError" }

// Error satisfies the builtin error interface
func (e PageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageValidationError{}

// Validate checks the field values on IdsRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdsRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IdsRequestMultiError, or
// nil if none found.
func (m *IdsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IdsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ids

	if len(errors) > 0 {
		return IdsRequestMultiError(errors)
	}

	return nil
}

// IdsRequestMultiError is an error wrapping multiple validation errors
// returned by IdsRequest.ValidateAll() if the designated constraints aren't met.
type IdsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdsRequestMultiError) AllErrors() []error { return m }

// IdsRequestValidationError is the validation error returned by
// IdsRequest.Validate if the designated constraints aren't met.
type IdsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdsRequestValidationError) ErrorName() string { return "IdsRequestValidationError" }

// Error satisfies the builtin error interface
func (e IdsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdsRequestValidationError{}

// Validate checks the field values on IdempotentReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IdempotentReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdempotentReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdempotentReplyMultiError, or nil if none found.
func (m *IdempotentReply) ValidateAll() error {
	return m.validate(true)
}

func (m *IdempotentReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return IdempotentReplyMultiError(errors)
	}

	return nil
}

// IdempotentReplyMultiError is an error wrapping multiple validation errors
// returned by IdempotentReply.ValidateAll() if the designated constraints
// aren't met.
type IdempotentReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdempotentReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdempotentReplyMultiError) AllErrors() []error { return m }

// IdempotentReplyValidationError is the validation error returned by
// IdempotentReply.Validate if the designated constraints aren't met.
type IdempotentReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdempotentReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdempotentReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdempotentReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdempotentReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdempotentReplyValidationError) ErrorName() string { return "IdempotentReplyValidationError" }

// Error satisfies the builtin error interface
func (e IdempotentReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdempotentReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdempotentReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdempotentReplyValidationError{}

// Validate checks the field values on GameReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GameReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GameReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GameReplyMultiError, or nil
// if none found.
func (m *GameReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GameReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Age

	if len(errors) > 0 {
		return GameReplyMultiError(errors)
	}

	return nil
}

// GameReplyMultiError is an error wrapping multiple validation errors returned
// by GameReply.ValidateAll() if the designated constraints aren't met.
type GameReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GameReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GameReplyMultiError) AllErrors() []error { return m }

// GameReplyValidationError is the validation error returned by
// GameReply.Validate if the designated constraints aren't met.
type GameReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GameReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GameReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GameReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GameReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GameReplyValidationError) ErrorName() string { return "GameReplyValidationError" }

// Error satisfies the builtin error interface
func (e GameReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGameReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GameReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GameReplyValidationError{}

// Validate checks the field values on CreateGameRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateGameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateGameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateGameRequestMultiError, or nil if none found.
func (m *CreateGameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateGameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 50 {
		err := CreateGameRequestValidationError{
			field:  "Name",
			reason: "value length must be between 2 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAge() < 0 {
		err := CreateGameRequestValidationError{
			field:  "Age",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateGameRequestMultiError(errors)
	}

	return nil
}

// CreateGameRequestMultiError is an error wrapping multiple validation errors
// returned by CreateGameRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateGameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateGameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateGameRequestMultiError) AllErrors() []error { return m }

// CreateGameRequestValidationError is the validation error returned by
// CreateGameRequest.Validate if the designated constraints aren't met.
type CreateGameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateGameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateGameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateGameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateGameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateGameRequestValidationError) ErrorName() string {
	return "CreateGameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateGameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateGameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateGameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateGameRequestValidationError{}

// Validate checks the field values on GetGameRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetGameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetGameRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetGameRequestMultiError,
// or nil if none found.
func (m *GetGameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetGameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetGameRequestMultiError(errors)
	}

	return nil
}

// GetGameRequestMultiError is an error wrapping multiple validation errors
// returned by GetGameRequest.ValidateAll() if the designated constraints
// aren't met.
type GetGameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetGameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetGameRequestMultiError) AllErrors() []error { return m }

// GetGameRequestValidationError is the validation error returned by
// GetGameRequest.Validate if the designated constraints aren't met.
type GetGameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGameRequestValidationError) ErrorName() string { return "GetGameRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetGameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGameRequestValidationError{}

// Validate checks the field values on GetGameReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetGameReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetGameReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetGameReplyMultiError, or
// nil if none found.
func (m *GetGameReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetGameReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Age

	if len(errors) > 0 {
		return GetGameReplyMultiError(errors)
	}

	return nil
}

// GetGameReplyMultiError is an error wrapping multiple validation errors
// returned by GetGameReply.ValidateAll() if the designated constraints aren't met.
type GetGameReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetGameReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetGameReplyMultiError) AllErrors() []error { return m }

// GetGameReplyValidationError is the validation error returned by
// GetGameReply.Validate if the designated constraints aren't met.
type GetGameReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGameReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGameReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGameReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGameReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGameReplyValidationError) ErrorName() string { return "GetGameReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetGameReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGameReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGameReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGameReplyValidationError{}

// Validate checks the field values on FindGameRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FindGameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindGameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindGameRequestMultiError, or nil if none found.
func (m *FindGameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FindGameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FindGameRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FindGameRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FindGameRequestValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.Age != nil {
		// no validation rules for Age
	}

	if len(errors) > 0 {
		return FindGameRequestMultiError(errors)
	}

	return nil
}

// FindGameRequestMultiError is an error wrapping multiple validation errors
// returned by FindGameRequest.ValidateAll() if the designated constraints
// aren't met.
type FindGameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindGameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindGameRequestMultiError) AllErrors() []error { return m }

// FindGameRequestValidationError is the validation error returned by
// FindGameRequest.Validate if the designated constraints aren't met.
type FindGameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindGameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindGameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindGameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindGameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindGameRequestValidationError) ErrorName() string { return "FindGameRequestValidationError" }

// Error satisfies the builtin error interface
func (e FindGameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindGameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindGameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindGameRequestValidationError{}

// Validate checks the field values on FindGameReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FindGameReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindGameReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FindGameReplyMultiError, or
// nil if none found.
func (m *FindGameReply) ValidateAll() error {
	return m.validate(true)
}

func (m *FindGameReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FindGameReplyValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FindGameReplyValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FindGameReplyValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FindGameReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FindGameReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FindGameReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FindGameReplyMultiError(errors)
	}

	return nil
}

// FindGameReplyMultiError is an error wrapping multiple validation errors
// returned by FindGameReply.ValidateAll() if the designated constraints
// aren't met.
type FindGameReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindGameReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindGameReplyMultiError) AllErrors() []error { return m }

// FindGameReplyValidationError is the validation error returned by
// FindGameReply.Validate if the designated constraints aren't met.
type FindGameReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindGameReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindGameReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindGameReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindGameReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindGameReplyValidationError) ErrorName() string { return "FindGameReplyValidationError" }

// Error satisfies the builtin error interface
func (e FindGameReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindGameReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindGameReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindGameReplyValidationError{}

// Validate checks the field values on UpdateGameRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateGameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateGameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateGameRequestMultiError, or nil if none found.
func (m *UpdateGameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateGameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Name != nil {

		if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 50 {
			err := UpdateGameRequestValidationError{
				field:  "Name",
				reason: "value length must be between 2 and 50 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Age != nil {

		if m.GetAge() < 0 {
			err := UpdateGameRequestValidationError{
				field:  "Age",
				reason: "value must be greater than or equal to 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UpdateGameRequestMultiError(errors)
	}

	return nil
}

// UpdateGameRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateGameRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateGameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateGameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateGameRequestMultiError) AllErrors() []error { return m }

// UpdateGameRequestValidationError is the validation error returned by
// UpdateGameRequest.Validate if the designated constraints aren't met.
type UpdateGameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateGameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateGameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateGameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateGameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateGameRequestValidationError) ErrorName() string {
	return "UpdateGameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateGameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateGameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateGameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateGameRequestValidationError{}
