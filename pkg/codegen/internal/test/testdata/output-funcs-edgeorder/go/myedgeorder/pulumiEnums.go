// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package myedgeorder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Type of product filter.
type SupportedFilterTypes string

const (
	// Ship to country
	SupportedFilterTypesShipToCountries = SupportedFilterTypes("ShipToCountries")
	// Double encryption status
	SupportedFilterTypesDoubleEncryptionStatus = SupportedFilterTypes("DoubleEncryptionStatus")
)

func (SupportedFilterTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedFilterTypes)(nil)).Elem()
}

func (e SupportedFilterTypes) ToSupportedFilterTypesOutput() SupportedFilterTypesOutput {
	return pulumi.ToOutput(e).(SupportedFilterTypesOutput)
}

func (e SupportedFilterTypes) ToSupportedFilterTypesOutputWithContext(ctx context.Context) SupportedFilterTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SupportedFilterTypesOutput)
}

func (e SupportedFilterTypes) ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput {
	return e.ToSupportedFilterTypesPtrOutputWithContext(context.Background())
}

func (e SupportedFilterTypes) ToSupportedFilterTypesPtrOutputWithContext(ctx context.Context) SupportedFilterTypesPtrOutput {
	return SupportedFilterTypes(e).ToSupportedFilterTypesOutputWithContext(ctx).ToSupportedFilterTypesPtrOutputWithContext(ctx)
}

func (e SupportedFilterTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedFilterTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedFilterTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SupportedFilterTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SupportedFilterTypesOutput struct{ *pulumi.OutputState }

func (SupportedFilterTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedFilterTypes)(nil)).Elem()
}

func (o SupportedFilterTypesOutput) ToSupportedFilterTypesOutput() SupportedFilterTypesOutput {
	return o
}

func (o SupportedFilterTypesOutput) ToSupportedFilterTypesOutputWithContext(ctx context.Context) SupportedFilterTypesOutput {
	return o
}

func (o SupportedFilterTypesOutput) ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput {
	return o.ToSupportedFilterTypesPtrOutputWithContext(context.Background())
}

func (o SupportedFilterTypesOutput) ToSupportedFilterTypesPtrOutputWithContext(ctx context.Context) SupportedFilterTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportedFilterTypes) *SupportedFilterTypes {
		return &v
	}).(SupportedFilterTypesPtrOutput)
}

func (o SupportedFilterTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SupportedFilterTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedFilterTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SupportedFilterTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedFilterTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedFilterTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SupportedFilterTypesPtrOutput struct{ *pulumi.OutputState }

func (SupportedFilterTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedFilterTypes)(nil)).Elem()
}

func (o SupportedFilterTypesPtrOutput) ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput {
	return o
}

func (o SupportedFilterTypesPtrOutput) ToSupportedFilterTypesPtrOutputWithContext(ctx context.Context) SupportedFilterTypesPtrOutput {
	return o
}

func (o SupportedFilterTypesPtrOutput) Elem() SupportedFilterTypesOutput {
	return o.ApplyT(func(v *SupportedFilterTypes) SupportedFilterTypes {
		if v != nil {
			return *v
		}
		var ret SupportedFilterTypes
		return ret
	}).(SupportedFilterTypesOutput)
}

func (o SupportedFilterTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedFilterTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SupportedFilterTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SupportedFilterTypesInput is an input type that accepts SupportedFilterTypesArgs and SupportedFilterTypesOutput values.
// You can construct a concrete instance of `SupportedFilterTypesInput` via:
//
//          SupportedFilterTypesArgs{...}
type SupportedFilterTypesInput interface {
	pulumi.Input

	ToSupportedFilterTypesOutput() SupportedFilterTypesOutput
	ToSupportedFilterTypesOutputWithContext(context.Context) SupportedFilterTypesOutput
}

var supportedFilterTypesPtrType = reflect.TypeOf((**SupportedFilterTypes)(nil)).Elem()

type SupportedFilterTypesPtrInput interface {
	pulumi.Input

	ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput
	ToSupportedFilterTypesPtrOutputWithContext(context.Context) SupportedFilterTypesPtrOutput
}

type supportedFilterTypesPtr string

func SupportedFilterTypesPtr(v string) SupportedFilterTypesPtrInput {
	return (*supportedFilterTypesPtr)(&v)
}

func (*supportedFilterTypesPtr) ElementType() reflect.Type {
	return supportedFilterTypesPtrType
}

func (in *supportedFilterTypesPtr) ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput {
	return pulumi.ToOutput(in).(SupportedFilterTypesPtrOutput)
}

func (in *supportedFilterTypesPtr) ToSupportedFilterTypesPtrOutputWithContext(ctx context.Context) SupportedFilterTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SupportedFilterTypesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SupportedFilterTypesInput)(nil)).Elem(), SupportedFilterTypes("ShipToCountries"))
	pulumi.RegisterInputType(reflect.TypeOf((*SupportedFilterTypesPtrInput)(nil)).Elem(), SupportedFilterTypes("ShipToCountries"))
	pulumi.RegisterOutputType(SupportedFilterTypesOutput{})
	pulumi.RegisterOutputType(SupportedFilterTypesPtrOutput{})
}
