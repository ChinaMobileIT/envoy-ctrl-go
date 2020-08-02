// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/local_ratelimit/v3/local_rate_limit.proto

package envoy_extensions_filters_network_local_ratelimit_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _local_rate_limit_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on LocalRateLimit with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LocalRateLimit) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetStatPrefix()) < 1 {
		return LocalRateLimitValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	if m.GetTokenBucket() == nil {
		return LocalRateLimitValidationError{
			field:  "TokenBucket",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetTokenBucket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "TokenBucket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRuntimeEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "RuntimeEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LocalRateLimitValidationError is the validation error returned by
// LocalRateLimit.Validate if the designated constraints aren't met.
type LocalRateLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalRateLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalRateLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalRateLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalRateLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalRateLimitValidationError) ErrorName() string { return "LocalRateLimitValidationError" }

// Error satisfies the builtin error interface
func (e LocalRateLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalRateLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalRateLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalRateLimitValidationError{}
