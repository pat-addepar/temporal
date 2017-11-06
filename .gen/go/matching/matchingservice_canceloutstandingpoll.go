// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package matching

import (
	"errors"
	"fmt"
	"strings"

	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
)

type MatchingService_CancelOutstandingPoll_Args struct {
	Request *CancelOutstandingPollRequest `json:"request,omitempty"`
}

func (v *MatchingService_CancelOutstandingPoll_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _CancelOutstandingPollRequest_Read(w wire.Value) (*CancelOutstandingPollRequest, error) {
	var v CancelOutstandingPollRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *MatchingService_CancelOutstandingPoll_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _CancelOutstandingPollRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *MatchingService_CancelOutstandingPoll_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}
	return fmt.Sprintf("MatchingService_CancelOutstandingPoll_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *MatchingService_CancelOutstandingPoll_Args) Equals(rhs *MatchingService_CancelOutstandingPoll_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}
	return true
}

func (v *MatchingService_CancelOutstandingPoll_Args) MethodName() string {
	return "CancelOutstandingPoll"
}

func (v *MatchingService_CancelOutstandingPoll_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var MatchingService_CancelOutstandingPoll_Helper = struct {
	Args           func(request *CancelOutstandingPollRequest) *MatchingService_CancelOutstandingPoll_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*MatchingService_CancelOutstandingPoll_Result, error)
	UnwrapResponse func(*MatchingService_CancelOutstandingPoll_Result) error
}{}

func init() {
	MatchingService_CancelOutstandingPoll_Helper.Args = func(request *CancelOutstandingPollRequest) *MatchingService_CancelOutstandingPoll_Args {
		return &MatchingService_CancelOutstandingPoll_Args{Request: request}
	}
	MatchingService_CancelOutstandingPoll_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		default:
			return false
		}
	}
	MatchingService_CancelOutstandingPoll_Helper.WrapResponse = func(err error) (*MatchingService_CancelOutstandingPoll_Result, error) {
		if err == nil {
			return &MatchingService_CancelOutstandingPoll_Result{}, nil
		}
		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for MatchingService_CancelOutstandingPoll_Result.BadRequestError")
			}
			return &MatchingService_CancelOutstandingPoll_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for MatchingService_CancelOutstandingPoll_Result.InternalServiceError")
			}
			return &MatchingService_CancelOutstandingPoll_Result{InternalServiceError: e}, nil
		}
		return nil, err
	}
	MatchingService_CancelOutstandingPoll_Helper.UnwrapResponse = func(result *MatchingService_CancelOutstandingPoll_Result) (err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		return
	}
}

type MatchingService_CancelOutstandingPoll_Result struct {
	BadRequestError      *shared.BadRequestError      `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError `json:"internalServiceError,omitempty"`
}

func (v *MatchingService_CancelOutstandingPoll_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("MatchingService_CancelOutstandingPoll_Result should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *MatchingService_CancelOutstandingPoll_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("MatchingService_CancelOutstandingPoll_Result should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *MatchingService_CancelOutstandingPoll_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	return fmt.Sprintf("MatchingService_CancelOutstandingPoll_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *MatchingService_CancelOutstandingPoll_Result) Equals(rhs *MatchingService_CancelOutstandingPoll_Result) bool {
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	return true
}

func (v *MatchingService_CancelOutstandingPoll_Result) MethodName() string {
	return "CancelOutstandingPoll"
}

func (v *MatchingService_CancelOutstandingPoll_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}