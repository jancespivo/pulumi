package codegentest
import (
        "context"
        "reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
// Codegen demo with const inputs
func FuncWithConstInput(ctx *pulumi.Context, args *FuncWithConstInputArgs, opts ...pulumi.InvokeOption) error {
	var rv struct{}
	err := ctx.Invoke("madeup-package:x:funcWithConstInput", args, &rv, opts...)
	return err
}

type FuncWithConstInputArgs struct {
	PlainInput *string `pulumi:"plainInput"`
}


func FuncWithConstInputApply(ctx *pulumi.Context, args FuncWithConstInputApplyInput, opts ...pulumi.InvokeOption) pulumi.Output {
	return args.ToFuncWithConstInputApplyOutput().ApplyT(func (v FuncWithConstInputArgs) (interface{}, error) {
	return nil, FuncWithConstInput(ctx, &v, opts...)

})}

// FuncWithConstInputApplyInput is an input type that accepts FuncWithConstInputApplyArgs and FuncWithConstInputApplyOutput values.
// You can construct a concrete instance of `FuncWithConstInputApplyInput` via:
//
//          FuncWithConstInputApplyArgs{...}
type FuncWithConstInputApplyInput interface {
	pulumi.Input

	ToFuncWithConstInputApplyOutput() FuncWithConstInputApplyOutput
	ToFuncWithConstInputApplyOutputWithContext(context.Context) FuncWithConstInputApplyOutput
}

type FuncWithConstInputApplyArgs struct {
	PlainInput pulumi.StringPtrInput `pulumi:"plainInput"`
}

func (FuncWithConstInputApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithConstInputArgs)(nil)).Elem()
}

func (i FuncWithConstInputApplyArgs) ToFuncWithConstInputApplyOutput() FuncWithConstInputApplyOutput {
	return i.ToFuncWithConstInputApplyOutputWithContext(context.Background())
}

func (i FuncWithConstInputApplyArgs) ToFuncWithConstInputApplyOutputWithContext(ctx context.Context) FuncWithConstInputApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FuncWithConstInputApplyOutput)
}

type FuncWithConstInputApplyOutput struct { *pulumi.OutputState }

func (FuncWithConstInputApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithConstInputArgs)(nil)).Elem()
}

func (o FuncWithConstInputApplyOutput) ToFuncWithConstInputApplyOutput() FuncWithConstInputApplyOutput {
	return o
}

func (o FuncWithConstInputApplyOutput) ToFuncWithConstInputApplyOutputWithContext(ctx context.Context) FuncWithConstInputApplyOutput {
	return o
}

func (o FuncWithConstInputApplyOutput) PlainInput() pulumi.StringPtrOutput {
	return o.ApplyT(func (v FuncWithConstInputArgs) *string { return v.PlainInput }).(pulumi.StringPtrOutput)
}

