// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by "modeldecoder/generator". DO NOT EDIT.

package rumv3

import (
	"encoding/json"
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/pkg/errors"
)

var (
	patternAlphaNumericExtRegexp    = regexp.MustCompile(patternAlphaNumericExt)
	patternNoDotAsteriskQuoteRegexp = regexp.MustCompile(patternNoDotAsteriskQuote)
)

func (val *metadataRoot) IsSet() bool {
	return val.Metadata.IsSet()
}

func (val *metadataRoot) Reset() {
	val.Metadata.Reset()
}

func (val *metadataRoot) validate() error {
	if err := val.Metadata.validate(); err != nil {
		return errors.Wrapf(err, "m")
	}
	if !val.Metadata.IsSet() {
		return fmt.Errorf("'m' required")
	}
	return nil
}

func (val *metadata) IsSet() bool {
	return len(val.Labels) > 0 || val.Service.IsSet() || val.User.IsSet()
}

func (val *metadata) Reset() {
	for k := range val.Labels {
		delete(val.Labels, k)
	}
	val.Service.Reset()
	val.User.Reset()
}

func (val *metadata) validate() error {
	if !val.IsSet() {
		return nil
	}
	for k, v := range val.Labels {
		if k != "" && !patternNoDotAsteriskQuoteRegexp.MatchString(k) {
			return fmt.Errorf("'l': validation rule 'patternKeys(patternNoDotAsteriskQuote)' violated")
		}
		switch t := v.(type) {
		case nil:
		case string:
			if utf8.RuneCountInString(t) > 1024 {
				return fmt.Errorf("'l': validation rule 'maxLengthVals(1024)' violated")
			}
		case bool:
		case json.Number:
		default:
			return fmt.Errorf("'l': validation rule 'inputTypesVals(string;bool;number)' violated for key %s", k)
		}
	}
	if err := val.Service.validate(); err != nil {
		return errors.Wrapf(err, "se")
	}
	if !val.Service.IsSet() {
		return fmt.Errorf("'se' required")
	}
	if err := val.User.validate(); err != nil {
		return errors.Wrapf(err, "u")
	}
	return nil
}

func (val *metadataService) IsSet() bool {
	return val.Agent.IsSet() || val.Environment.IsSet() || val.Framework.IsSet() || val.Language.IsSet() || val.Name.IsSet() || val.Runtime.IsSet() || val.Version.IsSet()
}

func (val *metadataService) Reset() {
	val.Agent.Reset()
	val.Environment.Reset()
	val.Framework.Reset()
	val.Language.Reset()
	val.Name.Reset()
	val.Runtime.Reset()
	val.Version.Reset()
}

func (val *metadataService) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Agent.validate(); err != nil {
		return errors.Wrapf(err, "a")
	}
	if !val.Agent.IsSet() {
		return fmt.Errorf("'a' required")
	}
	if utf8.RuneCountInString(val.Environment.Val) > 1024 {
		return fmt.Errorf("'en': validation rule 'maxLength(1024)' violated")
	}
	if err := val.Framework.validate(); err != nil {
		return errors.Wrapf(err, "fw")
	}
	if err := val.Language.validate(); err != nil {
		return errors.Wrapf(err, "la")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Name.Val) < 1 {
		return fmt.Errorf("'n': validation rule 'minLength(1)' violated")
	}
	if val.Name.Val != "" && !patternAlphaNumericExtRegexp.MatchString(val.Name.Val) {
		return fmt.Errorf("'n': validation rule 'pattern(patternAlphaNumericExt)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if err := val.Runtime.validate(); err != nil {
		return errors.Wrapf(err, "ru")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *metadataServiceAgent) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *metadataServiceAgent) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *metadataServiceAgent) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Name.Val) < 1 {
		return fmt.Errorf("'n': validation rule 'minLength(1)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	if !val.Version.IsSet() {
		return fmt.Errorf("'ve' required")
	}
	return nil
}

func (val *metadataServiceFramework) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *metadataServiceFramework) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *metadataServiceFramework) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *metadataServiceLanguage) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *metadataServiceLanguage) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *metadataServiceLanguage) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *metadataServiceRuntime) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *metadataServiceRuntime) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *metadataServiceRuntime) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	if !val.Version.IsSet() {
		return fmt.Errorf("'ve' required")
	}
	return nil
}

func (val *user) IsSet() bool {
	return val.ID.IsSet() || val.Email.IsSet() || val.Name.IsSet()
}

func (val *user) Reset() {
	val.ID.Reset()
	val.Email.Reset()
	val.Name.Reset()
}

func (val *user) validate() error {
	if !val.IsSet() {
		return nil
	}
	switch t := val.ID.Val.(type) {
	case string:
		if utf8.RuneCountInString(t) > 1024 {
			return fmt.Errorf("'id': validation rule 'maxLength(1024)' violated")
		}
	case int:
	case json.Number:
		if _, err := t.Int64(); err != nil {
			return fmt.Errorf("'id': validation rule 'inputTypes(string;int)' violated")
		}
	case nil:
	default:
		return fmt.Errorf("'id': validation rule 'inputTypes(string;int)' violated ")
	}
	if utf8.RuneCountInString(val.Email.Val) > 1024 {
		return fmt.Errorf("'em': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'un': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *errorRoot) IsSet() bool {
	return val.Error.IsSet()
}

func (val *errorRoot) Reset() {
	val.Error.Reset()
}

func (val *errorRoot) validate() error {
	if err := val.Error.validate(); err != nil {
		return errors.Wrapf(err, "e")
	}
	if !val.Error.IsSet() {
		return fmt.Errorf("'e' required")
	}
	return nil
}

func (val *errorEvent) IsSet() bool {
	return val.Context.IsSet() || val.Culprit.IsSet() || val.Exception.IsSet() || val.ID.IsSet() || val.Log.IsSet() || val.ParentID.IsSet() || val.Timestamp.IsSet() || val.TraceID.IsSet() || val.Transaction.IsSet() || val.TransactionID.IsSet()
}

func (val *errorEvent) Reset() {
	val.Context.Reset()
	val.Culprit.Reset()
	val.Exception.Reset()
	val.ID.Reset()
	val.Log.Reset()
	val.ParentID.Reset()
	val.Timestamp.Reset()
	val.TraceID.Reset()
	val.Transaction.Reset()
	val.TransactionID.Reset()
}

func (val *errorEvent) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Context.validate(); err != nil {
		return errors.Wrapf(err, "c")
	}
	if utf8.RuneCountInString(val.Culprit.Val) > 1024 {
		return fmt.Errorf("'cl': validation rule 'maxLength(1024)' violated")
	}
	if err := val.Exception.validate(); err != nil {
		return errors.Wrapf(err, "ex")
	}
	if utf8.RuneCountInString(val.ID.Val) > 1024 {
		return fmt.Errorf("'id': validation rule 'maxLength(1024)' violated")
	}
	if !val.ID.IsSet() {
		return fmt.Errorf("'id' required")
	}
	if err := val.Log.validate(); err != nil {
		return errors.Wrapf(err, "log")
	}
	if utf8.RuneCountInString(val.ParentID.Val) > 1024 {
		return fmt.Errorf("'pid': validation rule 'maxLength(1024)' violated")
	}
	if !val.ParentID.IsSet() {
		if val.TraceID.IsSet() {
			return fmt.Errorf("'pid' required when 'tid' is set")
		}
		if val.TransactionID.IsSet() {
			return fmt.Errorf("'pid' required when 'xid' is set")
		}
	}
	if utf8.RuneCountInString(val.TraceID.Val) > 1024 {
		return fmt.Errorf("'tid': validation rule 'maxLength(1024)' violated")
	}
	if !val.TraceID.IsSet() {
		if val.ParentID.IsSet() {
			return fmt.Errorf("'tid' required when 'pid' is set")
		}
		if val.TransactionID.IsSet() {
			return fmt.Errorf("'tid' required when 'xid' is set")
		}
	}
	if err := val.Transaction.validate(); err != nil {
		return errors.Wrapf(err, "x")
	}
	if utf8.RuneCountInString(val.TransactionID.Val) > 1024 {
		return fmt.Errorf("'xid': validation rule 'maxLength(1024)' violated")
	}
	if !val.Exception.IsSet() && !val.Log.IsSet() {
		return fmt.Errorf("requires at least one of the fields 'ex;log'")
	}
	return nil
}

func (val *context) IsSet() bool {
	return len(val.Custom) > 0 || val.Page.IsSet() || val.Response.IsSet() || val.Request.IsSet() || val.Service.IsSet() || len(val.Tags) > 0 || val.User.IsSet()
}

func (val *context) Reset() {
	for k := range val.Custom {
		delete(val.Custom, k)
	}
	val.Page.Reset()
	val.Response.Reset()
	val.Request.Reset()
	val.Service.Reset()
	for k := range val.Tags {
		delete(val.Tags, k)
	}
	val.User.Reset()
}

func (val *context) validate() error {
	if !val.IsSet() {
		return nil
	}
	for k := range val.Custom {
		if k != "" && !patternNoDotAsteriskQuoteRegexp.MatchString(k) {
			return fmt.Errorf("'cu': validation rule 'patternKeys(patternNoDotAsteriskQuote)' violated")
		}
	}
	if err := val.Page.validate(); err != nil {
		return errors.Wrapf(err, "p")
	}
	if err := val.Response.validate(); err != nil {
		return errors.Wrapf(err, "r")
	}
	if err := val.Request.validate(); err != nil {
		return errors.Wrapf(err, "q")
	}
	if err := val.Service.validate(); err != nil {
		return errors.Wrapf(err, "se")
	}
	for k, v := range val.Tags {
		if k != "" && !patternNoDotAsteriskQuoteRegexp.MatchString(k) {
			return fmt.Errorf("'g': validation rule 'patternKeys(patternNoDotAsteriskQuote)' violated")
		}
		switch t := v.(type) {
		case nil:
		case string:
			if utf8.RuneCountInString(t) > 1024 {
				return fmt.Errorf("'g': validation rule 'maxLengthVals(1024)' violated")
			}
		case bool:
		case json.Number:
		default:
			return fmt.Errorf("'g': validation rule 'inputTypesVals(string;bool;number)' violated for key %s", k)
		}
	}
	if err := val.User.validate(); err != nil {
		return errors.Wrapf(err, "u")
	}
	return nil
}

func (val *contextPage) IsSet() bool {
	return val.Referer.IsSet() || val.URL.IsSet()
}

func (val *contextPage) Reset() {
	val.Referer.Reset()
	val.URL.Reset()
}

func (val *contextPage) validate() error {
	if !val.IsSet() {
		return nil
	}
	return nil
}

func (val *contextResponse) IsSet() bool {
	return val.DecodedBodySize.IsSet() || val.EncodedBodySize.IsSet() || val.Headers.IsSet() || val.StatusCode.IsSet() || val.TransferSize.IsSet()
}

func (val *contextResponse) Reset() {
	val.DecodedBodySize.Reset()
	val.EncodedBodySize.Reset()
	val.Headers.Reset()
	val.StatusCode.Reset()
	val.TransferSize.Reset()
}

func (val *contextResponse) validate() error {
	if !val.IsSet() {
		return nil
	}
	return nil
}

func (val *contextRequest) IsSet() bool {
	return len(val.Env) > 0 || val.Headers.IsSet() || val.HTTPVersion.IsSet() || val.Method.IsSet()
}

func (val *contextRequest) Reset() {
	for k := range val.Env {
		delete(val.Env, k)
	}
	val.Headers.Reset()
	val.HTTPVersion.Reset()
	val.Method.Reset()
}

func (val *contextRequest) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.HTTPVersion.Val) > 1024 {
		return fmt.Errorf("'hve': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Method.Val) > 1024 {
		return fmt.Errorf("'mt': validation rule 'maxLength(1024)' violated")
	}
	if !val.Method.IsSet() {
		return fmt.Errorf("'mt' required")
	}
	return nil
}

func (val *contextService) IsSet() bool {
	return val.Agent.IsSet() || val.Environment.IsSet() || val.Framework.IsSet() || val.Language.IsSet() || val.Name.IsSet() || val.Runtime.IsSet() || val.Version.IsSet()
}

func (val *contextService) Reset() {
	val.Agent.Reset()
	val.Environment.Reset()
	val.Framework.Reset()
	val.Language.Reset()
	val.Name.Reset()
	val.Runtime.Reset()
	val.Version.Reset()
}

func (val *contextService) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Agent.validate(); err != nil {
		return errors.Wrapf(err, "a")
	}
	if utf8.RuneCountInString(val.Environment.Val) > 1024 {
		return fmt.Errorf("'en': validation rule 'maxLength(1024)' violated")
	}
	if err := val.Framework.validate(); err != nil {
		return errors.Wrapf(err, "fw")
	}
	if err := val.Language.validate(); err != nil {
		return errors.Wrapf(err, "la")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if val.Name.Val != "" && !patternAlphaNumericExtRegexp.MatchString(val.Name.Val) {
		return fmt.Errorf("'n': validation rule 'pattern(patternAlphaNumericExt)' violated")
	}
	if err := val.Runtime.validate(); err != nil {
		return errors.Wrapf(err, "ru")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *contextServiceAgent) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *contextServiceAgent) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *contextServiceAgent) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *contextServiceFramework) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *contextServiceFramework) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *contextServiceFramework) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *contextServiceLanguage) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *contextServiceLanguage) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *contextServiceLanguage) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *contextServiceRuntime) IsSet() bool {
	return val.Name.IsSet() || val.Version.IsSet()
}

func (val *contextServiceRuntime) Reset() {
	val.Name.Reset()
	val.Version.Reset()
}

func (val *contextServiceRuntime) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Version.Val) > 1024 {
		return fmt.Errorf("'ve': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *errorException) IsSet() bool {
	return len(val.Attributes) > 0 || val.Code.IsSet() || len(val.Cause) > 0 || val.Handled.IsSet() || val.Message.IsSet() || val.Module.IsSet() || len(val.Stacktrace) > 0 || val.Type.IsSet()
}

func (val *errorException) Reset() {
	for k := range val.Attributes {
		delete(val.Attributes, k)
	}
	val.Code.Reset()
	for i := range val.Cause {
		val.Cause[i].Reset()
	}
	val.Cause = val.Cause[:0]
	val.Handled.Reset()
	val.Message.Reset()
	val.Module.Reset()
	for i := range val.Stacktrace {
		val.Stacktrace[i].Reset()
	}
	val.Stacktrace = val.Stacktrace[:0]
	val.Type.Reset()
}

func (val *errorException) validate() error {
	if !val.IsSet() {
		return nil
	}
	switch t := val.Code.Val.(type) {
	case string:
		if utf8.RuneCountInString(t) > 1024 {
			return fmt.Errorf("'cd': validation rule 'maxLength(1024)' violated")
		}
	case int:
	case json.Number:
		if _, err := t.Int64(); err != nil {
			return fmt.Errorf("'cd': validation rule 'inputTypes(string;int)' violated")
		}
	case nil:
	default:
		return fmt.Errorf("'cd': validation rule 'inputTypes(string;int)' violated ")
	}
	for _, elem := range val.Cause {
		if err := elem.validate(); err != nil {
			return errors.Wrapf(err, "ca")
		}
	}
	if utf8.RuneCountInString(val.Module.Val) > 1024 {
		return fmt.Errorf("'mo': validation rule 'maxLength(1024)' violated")
	}
	for _, elem := range val.Stacktrace {
		if err := elem.validate(); err != nil {
			return errors.Wrapf(err, "st")
		}
	}
	if utf8.RuneCountInString(val.Type.Val) > 1024 {
		return fmt.Errorf("'t': validation rule 'maxLength(1024)' violated")
	}
	if !val.Message.IsSet() && !val.Type.IsSet() {
		return fmt.Errorf("requires at least one of the fields 'mg;t'")
	}
	return nil
}

func (val *stacktraceFrame) IsSet() bool {
	return val.AbsPath.IsSet() || val.Classname.IsSet() || val.ColumnNumber.IsSet() || val.ContextLine.IsSet() || val.Filename.IsSet() || val.Function.IsSet() || val.LineNumber.IsSet() || val.Module.IsSet() || len(val.PostContext) > 0 || len(val.PreContext) > 0
}

func (val *stacktraceFrame) Reset() {
	val.AbsPath.Reset()
	val.Classname.Reset()
	val.ColumnNumber.Reset()
	val.ContextLine.Reset()
	val.Filename.Reset()
	val.Function.Reset()
	val.LineNumber.Reset()
	val.Module.Reset()
	val.PostContext = val.PostContext[:0]
	val.PreContext = val.PreContext[:0]
}

func (val *stacktraceFrame) validate() error {
	if !val.IsSet() {
		return nil
	}
	if !val.Filename.IsSet() {
		return fmt.Errorf("'f' required")
	}
	return nil
}

func (val *errorLog) IsSet() bool {
	return val.Level.IsSet() || val.LoggerName.IsSet() || val.Message.IsSet() || val.ParamMessage.IsSet() || len(val.Stacktrace) > 0
}

func (val *errorLog) Reset() {
	val.Level.Reset()
	val.LoggerName.Reset()
	val.Message.Reset()
	val.ParamMessage.Reset()
	for i := range val.Stacktrace {
		val.Stacktrace[i].Reset()
	}
	val.Stacktrace = val.Stacktrace[:0]
}

func (val *errorLog) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Level.Val) > 1024 {
		return fmt.Errorf("'lv': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.LoggerName.Val) > 1024 {
		return fmt.Errorf("'ln': validation rule 'maxLength(1024)' violated")
	}
	if !val.Message.IsSet() {
		return fmt.Errorf("'mg' required")
	}
	if utf8.RuneCountInString(val.ParamMessage.Val) > 1024 {
		return fmt.Errorf("'pmg': validation rule 'maxLength(1024)' violated")
	}
	for _, elem := range val.Stacktrace {
		if err := elem.validate(); err != nil {
			return errors.Wrapf(err, "st")
		}
	}
	return nil
}

func (val *errorTransactionRef) IsSet() bool {
	return val.Sampled.IsSet() || val.Type.IsSet()
}

func (val *errorTransactionRef) Reset() {
	val.Sampled.Reset()
	val.Type.Reset()
}

func (val *errorTransactionRef) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Type.Val) > 1024 {
		return fmt.Errorf("'t': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *metricsetRoot) IsSet() bool {
	return val.Metricset.IsSet()
}

func (val *metricsetRoot) Reset() {
	val.Metricset.Reset()
}

func (val *metricsetRoot) validate() error {
	if err := val.Metricset.validate(); err != nil {
		return errors.Wrapf(err, "me")
	}
	if !val.Metricset.IsSet() {
		return fmt.Errorf("'me' required")
	}
	return nil
}

func (val *metricset) IsSet() bool {
	return val.Samples.IsSet() || val.Span.IsSet() || len(val.Tags) > 0
}

func (val *metricset) Reset() {
	val.Samples.Reset()
	val.Span.Reset()
	for k := range val.Tags {
		delete(val.Tags, k)
	}
}

func (val *metricset) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Samples.validate(); err != nil {
		return errors.Wrapf(err, "sa")
	}
	if !val.Samples.IsSet() {
		return fmt.Errorf("'sa' required")
	}
	if err := val.Span.validate(); err != nil {
		return errors.Wrapf(err, "y")
	}
	for k, v := range val.Tags {
		if k != "" && !patternNoDotAsteriskQuoteRegexp.MatchString(k) {
			return fmt.Errorf("'g': validation rule 'patternKeys(patternNoDotAsteriskQuote)' violated")
		}
		switch t := v.(type) {
		case nil:
		case string:
			if utf8.RuneCountInString(t) > 1024 {
				return fmt.Errorf("'g': validation rule 'maxLengthVals(1024)' violated")
			}
		case bool:
		case json.Number:
		default:
			return fmt.Errorf("'g': validation rule 'inputTypesVals(string;bool;number)' violated for key %s", k)
		}
	}
	return nil
}

func (val *metricsetSamples) IsSet() bool {
	return val.TransactionDurationCount.IsSet() || val.TransactionDurationSum.IsSet() || val.TransactionBreakdownCount.IsSet() || val.SpanSelfTimeCount.IsSet() || val.SpanSelfTimeSum.IsSet()
}

func (val *metricsetSamples) Reset() {
	val.TransactionDurationCount.Reset()
	val.TransactionDurationSum.Reset()
	val.TransactionBreakdownCount.Reset()
	val.SpanSelfTimeCount.Reset()
	val.SpanSelfTimeSum.Reset()
}

func (val *metricsetSamples) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.TransactionDurationCount.validate(); err != nil {
		return errors.Wrapf(err, "xdc")
	}
	if err := val.TransactionDurationSum.validate(); err != nil {
		return errors.Wrapf(err, "xds")
	}
	if err := val.TransactionBreakdownCount.validate(); err != nil {
		return errors.Wrapf(err, "xbc")
	}
	if err := val.SpanSelfTimeCount.validate(); err != nil {
		return errors.Wrapf(err, "ysc")
	}
	if err := val.SpanSelfTimeSum.validate(); err != nil {
		return errors.Wrapf(err, "yss")
	}
	return nil
}

func (val *metricsetSampleValue) IsSet() bool {
	return val.Value.IsSet()
}

func (val *metricsetSampleValue) Reset() {
	val.Value.Reset()
}

func (val *metricsetSampleValue) validate() error {
	if !val.IsSet() {
		return nil
	}
	if !val.Value.IsSet() {
		return fmt.Errorf("'v' required")
	}
	return nil
}

func (val *metricsetSpanRef) IsSet() bool {
	return val.Subtype.IsSet() || val.Type.IsSet()
}

func (val *metricsetSpanRef) Reset() {
	val.Subtype.Reset()
	val.Type.Reset()
}

func (val *metricsetSpanRef) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Subtype.Val) > 1024 {
		return fmt.Errorf("'su': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Type.Val) > 1024 {
		return fmt.Errorf("'t': validation rule 'maxLength(1024)' violated")
	}
	return nil
}

func (val *transactionRoot) IsSet() bool {
	return val.Transaction.IsSet()
}

func (val *transactionRoot) Reset() {
	val.Transaction.Reset()
}

func (val *transactionRoot) validate() error {
	if err := val.Transaction.validate(); err != nil {
		return errors.Wrapf(err, "x")
	}
	if !val.Transaction.IsSet() {
		return fmt.Errorf("'x' required")
	}
	return nil
}

func (val *transaction) IsSet() bool {
	return val.Context.IsSet() || val.Duration.IsSet() || val.ID.IsSet() || val.Marks.IsSet() || len(val.Metricsets) > 0 || val.Name.IsSet() || val.Outcome.IsSet() || val.ParentID.IsSet() || val.Result.IsSet() || val.Sampled.IsSet() || val.SampleRate.IsSet() || val.SpanCount.IsSet() || len(val.Spans) > 0 || val.TraceID.IsSet() || val.Type.IsSet() || val.UserExperience.IsSet()
}

func (val *transaction) Reset() {
	val.Context.Reset()
	val.Duration.Reset()
	val.ID.Reset()
	val.Marks.Reset()
	for i := range val.Metricsets {
		val.Metricsets[i].Reset()
	}
	val.Metricsets = val.Metricsets[:0]
	val.Name.Reset()
	val.Outcome.Reset()
	val.ParentID.Reset()
	val.Result.Reset()
	val.Sampled.Reset()
	val.SampleRate.Reset()
	val.SpanCount.Reset()
	for i := range val.Spans {
		val.Spans[i].Reset()
	}
	val.Spans = val.Spans[:0]
	val.TraceID.Reset()
	val.Type.Reset()
	val.UserExperience.Reset()
}

func (val *transaction) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Context.validate(); err != nil {
		return errors.Wrapf(err, "c")
	}
	if val.Duration.Val < 0 {
		return fmt.Errorf("'d': validation rule 'min(0)' violated")
	}
	if !val.Duration.IsSet() {
		return fmt.Errorf("'d' required")
	}
	if utf8.RuneCountInString(val.ID.Val) > 1024 {
		return fmt.Errorf("'id': validation rule 'maxLength(1024)' violated")
	}
	if !val.ID.IsSet() {
		return fmt.Errorf("'id' required")
	}
	if err := val.Marks.validate(); err != nil {
		return errors.Wrapf(err, "k")
	}
	for _, elem := range val.Metricsets {
		if err := elem.validate(); err != nil {
			return errors.Wrapf(err, "me")
		}
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if val.Outcome.Val != "" {
		var matchEnum bool
		for _, s := range enumOutcome {
			if val.Outcome.Val == s {
				matchEnum = true
				break
			}
		}
		if !matchEnum {
			return fmt.Errorf("'o': validation rule 'enum(enumOutcome)' violated")
		}
	}
	if utf8.RuneCountInString(val.ParentID.Val) > 1024 {
		return fmt.Errorf("'pid': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Result.Val) > 1024 {
		return fmt.Errorf("'rt': validation rule 'maxLength(1024)' violated")
	}
	if err := val.SpanCount.validate(); err != nil {
		return errors.Wrapf(err, "yc")
	}
	if !val.SpanCount.IsSet() {
		return fmt.Errorf("'yc' required")
	}
	for _, elem := range val.Spans {
		if err := elem.validate(); err != nil {
			return errors.Wrapf(err, "y")
		}
	}
	if utf8.RuneCountInString(val.TraceID.Val) > 1024 {
		return fmt.Errorf("'tid': validation rule 'maxLength(1024)' violated")
	}
	if !val.TraceID.IsSet() {
		return fmt.Errorf("'tid' required")
	}
	if utf8.RuneCountInString(val.Type.Val) > 1024 {
		return fmt.Errorf("'t': validation rule 'maxLength(1024)' violated")
	}
	if !val.Type.IsSet() {
		return fmt.Errorf("'t' required")
	}
	if err := val.UserExperience.validate(); err != nil {
		return errors.Wrapf(err, "exp")
	}
	return nil
}

func (val *transactionMarks) IsSet() bool {
	return len(val.Events) > 0
}

func (val *transactionMarks) Reset() {
	for k := range val.Events {
		delete(val.Events, k)
	}
}

func (val *transactionMarks) validate() error {
	if !val.IsSet() {
		return nil
	}
	for k, v := range val.Events {
		if err := v.validate(); err != nil {
			return errors.Wrapf(err, "events")
		}
		if k != "" && !patternNoDotAsteriskQuoteRegexp.MatchString(k) {
			return fmt.Errorf("'events': validation rule 'patternKeys(patternNoDotAsteriskQuote)' violated")
		}
	}
	return nil
}

func (val *transactionMarkEvents) IsSet() bool {
	return len(val.Measurements) > 0
}

func (val *transactionMarkEvents) Reset() {
	for k := range val.Measurements {
		delete(val.Measurements, k)
	}
}

func (val *transactionMarkEvents) validate() error {
	if !val.IsSet() {
		return nil
	}
	for k := range val.Measurements {
		if k != "" && !patternNoDotAsteriskQuoteRegexp.MatchString(k) {
			return fmt.Errorf("'measurements': validation rule 'patternKeys(patternNoDotAsteriskQuote)' violated")
		}
	}
	return nil
}

func (val *transactionSpanCount) IsSet() bool {
	return val.Dropped.IsSet() || val.Started.IsSet()
}

func (val *transactionSpanCount) Reset() {
	val.Dropped.Reset()
	val.Started.Reset()
}

func (val *transactionSpanCount) validate() error {
	if !val.IsSet() {
		return nil
	}
	if !val.Started.IsSet() {
		return fmt.Errorf("'sd' required")
	}
	return nil
}

func (val *span) IsSet() bool {
	return val.Action.IsSet() || val.Context.IsSet() || val.Duration.IsSet() || val.ID.IsSet() || val.Name.IsSet() || val.Outcome.IsSet() || val.ParentIndex.IsSet() || val.SampleRate.IsSet() || len(val.Stacktrace) > 0 || val.Start.IsSet() || val.Subtype.IsSet() || val.Sync.IsSet() || val.Type.IsSet()
}

func (val *span) Reset() {
	val.Action.Reset()
	val.Context.Reset()
	val.Duration.Reset()
	val.ID.Reset()
	val.Name.Reset()
	val.Outcome.Reset()
	val.ParentIndex.Reset()
	val.SampleRate.Reset()
	for i := range val.Stacktrace {
		val.Stacktrace[i].Reset()
	}
	val.Stacktrace = val.Stacktrace[:0]
	val.Start.Reset()
	val.Subtype.Reset()
	val.Sync.Reset()
	val.Type.Reset()
}

func (val *span) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Action.Val) > 1024 {
		return fmt.Errorf("'ac': validation rule 'maxLength(1024)' violated")
	}
	if err := val.Context.validate(); err != nil {
		return errors.Wrapf(err, "c")
	}
	if val.Duration.Val < 0 {
		return fmt.Errorf("'d': validation rule 'min(0)' violated")
	}
	if !val.Duration.IsSet() {
		return fmt.Errorf("'d' required")
	}
	if utf8.RuneCountInString(val.ID.Val) > 1024 {
		return fmt.Errorf("'id': validation rule 'maxLength(1024)' violated")
	}
	if !val.ID.IsSet() {
		return fmt.Errorf("'id' required")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if val.Outcome.Val != "" {
		var matchEnum bool
		for _, s := range enumOutcome {
			if val.Outcome.Val == s {
				matchEnum = true
				break
			}
		}
		if !matchEnum {
			return fmt.Errorf("'o': validation rule 'enum(enumOutcome)' violated")
		}
	}
	for _, elem := range val.Stacktrace {
		if err := elem.validate(); err != nil {
			return errors.Wrapf(err, "st")
		}
	}
	if !val.Start.IsSet() {
		return fmt.Errorf("'s' required")
	}
	if utf8.RuneCountInString(val.Subtype.Val) > 1024 {
		return fmt.Errorf("'su': validation rule 'maxLength(1024)' violated")
	}
	if utf8.RuneCountInString(val.Type.Val) > 1024 {
		return fmt.Errorf("'t': validation rule 'maxLength(1024)' violated")
	}
	if !val.Type.IsSet() {
		return fmt.Errorf("'t' required")
	}
	return nil
}

func (val *spanContext) IsSet() bool {
	return val.Destination.IsSet() || val.HTTP.IsSet() || val.Service.IsSet() || len(val.Tags) > 0
}

func (val *spanContext) Reset() {
	val.Destination.Reset()
	val.HTTP.Reset()
	val.Service.Reset()
	for k := range val.Tags {
		delete(val.Tags, k)
	}
}

func (val *spanContext) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Destination.validate(); err != nil {
		return errors.Wrapf(err, "dt")
	}
	if err := val.HTTP.validate(); err != nil {
		return errors.Wrapf(err, "h")
	}
	if err := val.Service.validate(); err != nil {
		return errors.Wrapf(err, "se")
	}
	for k, v := range val.Tags {
		if k != "" && !patternNoDotAsteriskQuoteRegexp.MatchString(k) {
			return fmt.Errorf("'g': validation rule 'patternKeys(patternNoDotAsteriskQuote)' violated")
		}
		switch t := v.(type) {
		case nil:
		case string:
			if utf8.RuneCountInString(t) > 1024 {
				return fmt.Errorf("'g': validation rule 'maxLengthVals(1024)' violated")
			}
		case bool:
		case json.Number:
		default:
			return fmt.Errorf("'g': validation rule 'inputTypesVals(string;bool;number)' violated for key %s", k)
		}
	}
	return nil
}

func (val *spanContextDestination) IsSet() bool {
	return val.Address.IsSet() || val.Port.IsSet() || val.Service.IsSet()
}

func (val *spanContextDestination) Reset() {
	val.Address.Reset()
	val.Port.Reset()
	val.Service.Reset()
}

func (val *spanContextDestination) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Address.Val) > 1024 {
		return fmt.Errorf("'ad': validation rule 'maxLength(1024)' violated")
	}
	if err := val.Service.validate(); err != nil {
		return errors.Wrapf(err, "se")
	}
	return nil
}

func (val *spanContextDestinationService) IsSet() bool {
	return val.Name.IsSet() || val.Resource.IsSet() || val.Type.IsSet()
}

func (val *spanContextDestinationService) Reset() {
	val.Name.Reset()
	val.Resource.Reset()
	val.Type.Reset()
}

func (val *spanContextDestinationService) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if !val.Name.IsSet() {
		return fmt.Errorf("'n' required")
	}
	if utf8.RuneCountInString(val.Resource.Val) > 1024 {
		return fmt.Errorf("'rc': validation rule 'maxLength(1024)' violated")
	}
	if !val.Resource.IsSet() {
		return fmt.Errorf("'rc' required")
	}
	if utf8.RuneCountInString(val.Type.Val) > 1024 {
		return fmt.Errorf("'t': validation rule 'maxLength(1024)' violated")
	}
	if !val.Type.IsSet() {
		return fmt.Errorf("'t' required")
	}
	return nil
}

func (val *spanContextHTTP) IsSet() bool {
	return val.Method.IsSet() || val.Response.IsSet() || val.StatusCode.IsSet() || val.URL.IsSet()
}

func (val *spanContextHTTP) Reset() {
	val.Method.Reset()
	val.Response.Reset()
	val.StatusCode.Reset()
	val.URL.Reset()
}

func (val *spanContextHTTP) validate() error {
	if !val.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(val.Method.Val) > 1024 {
		return fmt.Errorf("'mt': validation rule 'maxLength(1024)' violated")
	}
	if err := val.Response.validate(); err != nil {
		return errors.Wrapf(err, "r")
	}
	return nil
}

func (val *spanContextHTTPResponse) IsSet() bool {
	return val.DecodedBodySize.IsSet() || val.EncodedBodySize.IsSet() || val.TransferSize.IsSet()
}

func (val *spanContextHTTPResponse) Reset() {
	val.DecodedBodySize.Reset()
	val.EncodedBodySize.Reset()
	val.TransferSize.Reset()
}

func (val *spanContextHTTPResponse) validate() error {
	if !val.IsSet() {
		return nil
	}
	return nil
}

func (val *spanContextService) IsSet() bool {
	return val.Agent.IsSet() || val.Name.IsSet()
}

func (val *spanContextService) Reset() {
	val.Agent.Reset()
	val.Name.Reset()
}

func (val *spanContextService) validate() error {
	if !val.IsSet() {
		return nil
	}
	if err := val.Agent.validate(); err != nil {
		return errors.Wrapf(err, "a")
	}
	if utf8.RuneCountInString(val.Name.Val) > 1024 {
		return fmt.Errorf("'n': validation rule 'maxLength(1024)' violated")
	}
	if val.Name.Val != "" && !patternAlphaNumericExtRegexp.MatchString(val.Name.Val) {
		return fmt.Errorf("'n': validation rule 'pattern(patternAlphaNumericExt)' violated")
	}
	return nil
}

func (val *transactionUserExperience) IsSet() bool {
	return val.CumulativeLayoutShift.IsSet() || val.FirstInputDelay.IsSet() || val.TotalBlockingTime.IsSet() || val.Longtask.IsSet()
}

func (val *transactionUserExperience) Reset() {
	val.CumulativeLayoutShift.Reset()
	val.FirstInputDelay.Reset()
	val.TotalBlockingTime.Reset()
	val.Longtask.Reset()
}

func (val *transactionUserExperience) validate() error {
	if !val.IsSet() {
		return nil
	}
	if val.CumulativeLayoutShift.Val < 0 {
		return fmt.Errorf("'cls': validation rule 'min(0)' violated")
	}
	if val.FirstInputDelay.Val < 0 {
		return fmt.Errorf("'fid': validation rule 'min(0)' violated")
	}
	if val.TotalBlockingTime.Val < 0 {
		return fmt.Errorf("'tbt': validation rule 'min(0)' violated")
	}
	if err := val.Longtask.validate(); err != nil {
		return errors.Wrapf(err, "lt")
	}
	return nil
}

func (val *longtaskMetrics) IsSet() bool {
	return val.Count.IsSet() || val.Max.IsSet() || val.Sum.IsSet()
}

func (val *longtaskMetrics) Reset() {
	val.Count.Reset()
	val.Max.Reset()
	val.Sum.Reset()
}

func (val *longtaskMetrics) validate() error {
	if !val.IsSet() {
		return nil
	}
	if val.Count.Val < 0 {
		return fmt.Errorf("'count': validation rule 'min(0)' violated")
	}
	if !val.Count.IsSet() {
		return fmt.Errorf("'count' required")
	}
	if val.Max.Val < 0 {
		return fmt.Errorf("'max': validation rule 'min(0)' violated")
	}
	if !val.Max.IsSet() {
		return fmt.Errorf("'max' required")
	}
	if val.Sum.Val < 0 {
		return fmt.Errorf("'sum': validation rule 'min(0)' violated")
	}
	if !val.Sum.IsSet() {
		return fmt.Errorf("'sum' required")
	}
	return nil
}
