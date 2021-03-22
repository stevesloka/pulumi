// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen/internal/test/testdata/simple-enum-schema/go/plant"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RubberTree struct {
	pulumi.CustomResourceState

	Container plant.ContainerPtrOutput `pulumi:"container"`
	Diameter  pulumi.Float64Output     `pulumi:"diameter"`
	Farm      pulumi.StringPtrOutput   `pulumi:"farm"`
	Size      pulumi.StringPtrOutput   `pulumi:"size"`
	Type      pulumi.StringOutput      `pulumi:"type"`
}

// NewRubberTree registers a new resource with the given unique name, arguments, and options.
func NewRubberTree(ctx *pulumi.Context,
	name string, args *RubberTreeArgs, opts ...pulumi.ResourceOption) (*RubberTree, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Diameter == 0 {
		args.Diameter = Diameter(6)
	}
	if args.Farm == nil {
		args.Farm = pulumi.StringPtr("(unknown)")
	}
	if args.Size == nil {
		e := TreeSize("medium")
		args.Size = &e
	}
	if args.Type == "" {
		args.Type = RubberTreeVariety("Burgundy")
	}
	var resource RubberTree
	err := ctx.RegisterResource("plant:tree/v1:RubberTree", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRubberTree gets an existing RubberTree resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRubberTree(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RubberTreeState, opts ...pulumi.ResourceOption) (*RubberTree, error) {
	var resource RubberTree
	err := ctx.ReadResource("plant:tree/v1:RubberTree", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RubberTree resources.
type rubberTreeState struct {
	Container *plant.Container `pulumi:"container"`
	Diameter  *float64         `pulumi:"diameter"`
	Farm      *string          `pulumi:"farm"`
	Size      *string          `pulumi:"size"`
	Type      *string          `pulumi:"type"`
}

type RubberTreeState struct {
	Container plant.ContainerPtrInput
	Diameter  *Diameter
	Farm      pulumi.StringPtrInput
	Size      *TreeSize
	Type      *RubberTreeVariety
}

func (RubberTreeState) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeState)(nil)).Elem()
}

type rubberTreeArgs struct {
	Container *plant.Container `pulumi:"container"`
	Diameter  float64          `pulumi:"diameter"`
	Farm      *string          `pulumi:"farm"`
	Size      *string          `pulumi:"size"`
	Type      string           `pulumi:"type"`
}

// The set of arguments for constructing a RubberTree resource.
type RubberTreeArgs struct {
	Container plant.ContainerPtrInput
	Diameter  Diameter
	Farm      pulumi.StringPtrInput
	Size      *TreeSize
	Type      RubberTreeVariety
}

func (RubberTreeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeArgs)(nil)).Elem()
}

type RubberTreeInput interface {
	pulumi.Input

	ToRubberTreeOutput() RubberTreeOutput
	ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput
}

func (*RubberTree) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTree)(nil))
}

func (i *RubberTree) ToRubberTreeOutput() RubberTreeOutput {
	return i.ToRubberTreeOutputWithContext(context.Background())
}

func (i *RubberTree) ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RubberTreeOutput)
}

type RubberTreeOutput struct {
	*pulumi.OutputState
}

func (RubberTreeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTree)(nil))
}

func (o RubberTreeOutput) ToRubberTreeOutput() RubberTreeOutput {
	return o
}

func (o RubberTreeOutput) ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RubberTreeOutput{})
}
