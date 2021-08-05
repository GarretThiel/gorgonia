package stdops

import (
	"context"
	"runtime/trace"

	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

// Sub is a tensor-tensor elementwise subtraction
type Sub struct{ binop }

// String implements fmt.Stringer.
func (op Sub) String() string { return "-" }

// Do performs elementwise subtraction.
func (op Sub) Do(ctx context.Context, vs ...values.Value) (values.Value, error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err := tensor.Sub(a, b, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise subtraction but with a preallocated return value.
// PreallocDo allows Sub to implement ops.PreallocOp
func (op Sub) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (values.Value, error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err := tensor.Sub(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// SubVS is a tensor-scalar elementwise subtraction
type SubVS struct{ binopVS }

// String implements fmt.Stringer.
func (op SubVS) String() string { return "-·" }

// Do performs elementwise subtraction.
func (op SubVS) Do(ctx context.Context, vs ...values.Value) (values.Value, error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err := tensor.Sub(a, b, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise subtraction but with a preallocated return value.
// PreallocDo allows SubVS to implement ops.PreallocOp
func (op SubVS) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (values.Value, error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err := tensor.Sub(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// SubSV is a scalar-tensor elementwise subtraction
type SubSV struct{ binopSV }

// String implements fmt.Stringer.
func (op SubSV) String() string { return "·-" }

// Do performs elementwise subtraction.
func (op SubSV) Do(ctx context.Context, vs ...values.Value) (values.Value, error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err := tensor.Sub(a, b, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise subtraction but with a preallocated return value.
// PreallocDo allows SubSV to implement ops.PreallocOp
func (op SubSV) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (values.Value, error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err := tensor.Sub(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}
