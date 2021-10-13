package pkg

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func GetFromDoubleValue(v *wrapperspb.DoubleValue) *float64 {
	if v == nil {
		return nil
	}
	return &v.Value
}

func GetFromFloatValue(v *wrapperspb.FloatValue) *float32 {
	if v == nil {
		return nil
	}
	return &v.Value
}

func GetFromInt64Value(v *wrapperspb.Int64Value) *int64 {
	if v == nil {
		return nil
	}
	return &v.Value
}

func GetFromUInt64Value(v *wrapperspb.UInt64Value) *uint64 {
	if v == nil {
		return nil
	}
	return &v.Value
}

func GetFromInt32Value(v *wrapperspb.Int32Value) *int32 {
	if v == nil {
		return nil
	}
	return &v.Value
}

func GetFromUInt32Value(v *wrapperspb.UInt32Value) *uint32 {
	if v == nil {
		return nil
	}
	return &v.Value
}

func GetFromBoolValue(v *wrapperspb.BoolValue) *bool {
	if v == nil {
		return nil
	}
	return &v.Value
}

func GetFromStringValue(v *wrapperspb.StringValue) *string {
	if v == nil {
		return nil
	}
	return &v.Value
}

func GetFromBytesValue(v *wrapperspb.BytesValue) *[]byte {
	if v == nil {
		return nil
	}
	return &v.Value
}

func CreateDoubleValue(v float64) *wrapperspb.DoubleValue {
	return &wrapperspb.DoubleValue{Value: v}
}

func CreateFloatValue(v float32) *wrapperspb.FloatValue {
	return &wrapperspb.FloatValue{Value: v}
}

func CreateInt64Value(v int64) *wrapperspb.Int64Value {
	return &wrapperspb.Int64Value{Value: v}
}

func CreateUInt64Value(v uint64) *wrapperspb.UInt64Value {
	return &wrapperspb.UInt64Value{Value: v}
}

func CreateInt32Value(v int32) *wrapperspb.Int32Value {
	return &wrapperspb.Int32Value{Value: v}
}

func CreateUInt32Value(v uint32) *wrapperspb.UInt32Value {
	return &wrapperspb.UInt32Value{Value: v}
}

func CreateBoolValue(v bool) *wrapperspb.BoolValue {
	return &wrapperspb.BoolValue{Value: v}
}

func CreateStringValue(v string) *wrapperspb.StringValue {
	return &wrapperspb.StringValue{Value: v}
}

func CreateBytesValue(v *[]byte) *wrapperspb.BytesValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.BytesValue{Value: *v}
}

func CreateBytesValueStr(s string) *wrapperspb.BytesValue {
	return &wrapperspb.BytesValue{Value: []byte(s)}
}

func CreateBytesValueStrPtr(s *string) *wrapperspb.BytesValue {
	if s == nil {
		return nil
	}
	return &wrapperspb.BytesValue{Value: []byte(*s)}
}

func CreateDoubleValuePtr(v *float64) *wrapperspb.DoubleValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.DoubleValue{Value: *v}
}

func CreateFloatValuePtr(v *float32) *wrapperspb.FloatValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.FloatValue{Value: *v}
}

func CreateInt64ValuePtr(v *int64) *wrapperspb.Int64Value {
	if v == nil {
		return nil
	}
	return &wrapperspb.Int64Value{Value: *v}
}

func CreateUInt64ValuePtr(v *uint64) *wrapperspb.UInt64Value {
	if v == nil {
		return nil
	}
	return &wrapperspb.UInt64Value{Value: *v}
}

func CreateInt32ValuePtr(v *int32) *wrapperspb.Int32Value {
	if v == nil {
		return nil
	}
	return &wrapperspb.Int32Value{Value: *v}
}

func CreateUInt32ValuePtr(v *uint32) *wrapperspb.UInt32Value {
	if v == nil {
		return nil
	}
	return &wrapperspb.UInt32Value{Value: *v}
}

func CreateBoolValuePtr(v *bool) *wrapperspb.BoolValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.BoolValue{Value: *v}
}

func CreateStringValuePtr(v *string) *wrapperspb.StringValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.StringValue{Value: *v}
}
