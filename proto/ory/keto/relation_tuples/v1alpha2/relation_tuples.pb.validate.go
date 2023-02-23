// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ory/keto/relation_tuples/v1alpha2/relation_tuples.proto

package rts

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

// Validate checks the field values on RelationTuple with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RelationTuple) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RelationTuple with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RelationTupleMultiError, or
// nil if none found.
func (m *RelationTuple) ValidateAll() error {
	return m.validate(true)
}

func (m *RelationTuple) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Namespace

	// no validation rules for Object

	// no validation rules for Relation

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RelationTupleValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RelationTupleValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationTupleValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.RestApiSubject.(type) {
	case *RelationTuple_SubjectId:
		if v == nil {
			err := RelationTupleValidationError{
				field:  "RestApiSubject",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for SubjectId
	case *RelationTuple_SubjectSet:
		if v == nil {
			err := RelationTupleValidationError{
				field:  "RestApiSubject",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSubjectSet()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RelationTupleValidationError{
						field:  "SubjectSet",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RelationTupleValidationError{
						field:  "SubjectSet",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSubjectSet()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RelationTupleValidationError{
					field:  "SubjectSet",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return RelationTupleMultiError(errors)
	}

	return nil
}

// RelationTupleMultiError is an error wrapping multiple validation errors
// returned by RelationTuple.ValidateAll() if the designated constraints
// aren't met.
type RelationTupleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RelationTupleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RelationTupleMultiError) AllErrors() []error { return m }

// RelationTupleValidationError is the validation error returned by
// RelationTuple.Validate if the designated constraints aren't met.
type RelationTupleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationTupleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationTupleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationTupleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationTupleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationTupleValidationError) ErrorName() string { return "RelationTupleValidationError" }

// Error satisfies the builtin error interface
func (e RelationTupleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationTuple.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationTupleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationTupleValidationError{}

// Validate checks the field values on RelationQuery with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RelationQuery) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RelationQuery with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RelationQueryMultiError, or
// nil if none found.
func (m *RelationQuery) ValidateAll() error {
	return m.validate(true)
}

func (m *RelationQuery) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Namespace != nil {
		// no validation rules for Namespace
	}

	if m.Object != nil {
		// no validation rules for Object
	}

	if m.Relation != nil {
		// no validation rules for Relation
	}

	if m.Subject != nil {

		if all {
			switch v := interface{}(m.GetSubject()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RelationQueryValidationError{
						field:  "Subject",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RelationQueryValidationError{
						field:  "Subject",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RelationQueryValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RelationQueryMultiError(errors)
	}

	return nil
}

// RelationQueryMultiError is an error wrapping multiple validation errors
// returned by RelationQuery.ValidateAll() if the designated constraints
// aren't met.
type RelationQueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RelationQueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RelationQueryMultiError) AllErrors() []error { return m }

// RelationQueryValidationError is the validation error returned by
// RelationQuery.Validate if the designated constraints aren't met.
type RelationQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationQueryValidationError) ErrorName() string { return "RelationQueryValidationError" }

// Error satisfies the builtin error interface
func (e RelationQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationQueryValidationError{}

// Validate checks the field values on Subject with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Subject) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Subject with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SubjectMultiError, or nil if none found.
func (m *Subject) ValidateAll() error {
	return m.validate(true)
}

func (m *Subject) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Ref.(type) {
	case *Subject_Id:
		if v == nil {
			err := SubjectValidationError{
				field:  "Ref",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Id
	case *Subject_Set:
		if v == nil {
			err := SubjectValidationError{
				field:  "Ref",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSet()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubjectValidationError{
						field:  "Set",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubjectValidationError{
						field:  "Set",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSet()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubjectValidationError{
					field:  "Set",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return SubjectMultiError(errors)
	}

	return nil
}

// SubjectMultiError is an error wrapping multiple validation errors returned
// by Subject.ValidateAll() if the designated constraints aren't met.
type SubjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectMultiError) AllErrors() []error { return m }

// SubjectValidationError is the validation error returned by Subject.Validate
// if the designated constraints aren't met.
type SubjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectValidationError) ErrorName() string { return "SubjectValidationError" }

// Error satisfies the builtin error interface
func (e SubjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectValidationError{}

// Validate checks the field values on SubjectSet with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubjectSet) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubjectSet with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubjectSetMultiError, or
// nil if none found.
func (m *SubjectSet) ValidateAll() error {
	return m.validate(true)
}

func (m *SubjectSet) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Namespace

	// no validation rules for Object

	// no validation rules for Relation

	if len(errors) > 0 {
		return SubjectSetMultiError(errors)
	}

	return nil
}

// SubjectSetMultiError is an error wrapping multiple validation errors
// returned by SubjectSet.ValidateAll() if the designated constraints aren't met.
type SubjectSetMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectSetMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectSetMultiError) AllErrors() []error { return m }

// SubjectSetValidationError is the validation error returned by
// SubjectSet.Validate if the designated constraints aren't met.
type SubjectSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectSetValidationError) ErrorName() string { return "SubjectSetValidationError" }

// Error satisfies the builtin error interface
func (e SubjectSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectSetValidationError{}

// Validate checks the field values on SubjectQuery with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubjectQuery) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubjectQuery with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubjectQueryMultiError, or
// nil if none found.
func (m *SubjectQuery) ValidateAll() error {
	return m.validate(true)
}

func (m *SubjectQuery) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Ref.(type) {
	case *SubjectQuery_Id:
		if v == nil {
			err := SubjectQueryValidationError{
				field:  "Ref",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Id
	case *SubjectQuery_Set:
		if v == nil {
			err := SubjectQueryValidationError{
				field:  "Ref",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSet()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubjectQueryValidationError{
						field:  "Set",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubjectQueryValidationError{
						field:  "Set",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSet()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubjectQueryValidationError{
					field:  "Set",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return SubjectQueryMultiError(errors)
	}

	return nil
}

// SubjectQueryMultiError is an error wrapping multiple validation errors
// returned by SubjectQuery.ValidateAll() if the designated constraints aren't met.
type SubjectQueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectQueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectQueryMultiError) AllErrors() []error { return m }

// SubjectQueryValidationError is the validation error returned by
// SubjectQuery.Validate if the designated constraints aren't met.
type SubjectQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectQueryValidationError) ErrorName() string { return "SubjectQueryValidationError" }

// Error satisfies the builtin error interface
func (e SubjectQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectQueryValidationError{}

// Validate checks the field values on SubjectSetQuery with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SubjectSetQuery) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubjectSetQuery with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubjectSetQueryMultiError, or nil if none found.
func (m *SubjectSetQuery) ValidateAll() error {
	return m.validate(true)
}

func (m *SubjectSetQuery) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Namespace

	// no validation rules for Object

	// no validation rules for Relation

	if len(errors) > 0 {
		return SubjectSetQueryMultiError(errors)
	}

	return nil
}

// SubjectSetQueryMultiError is an error wrapping multiple validation errors
// returned by SubjectSetQuery.ValidateAll() if the designated constraints
// aren't met.
type SubjectSetQueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectSetQueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectSetQueryMultiError) AllErrors() []error { return m }

// SubjectSetQueryValidationError is the validation error returned by
// SubjectSetQuery.Validate if the designated constraints aren't met.
type SubjectSetQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectSetQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectSetQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectSetQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectSetQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectSetQueryValidationError) ErrorName() string { return "SubjectSetQueryValidationError" }

// Error satisfies the builtin error interface
func (e SubjectSetQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectSetQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectSetQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectSetQueryValidationError{}