package stdops

import (
	"context"
	"runtime/trace"

	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

// sign is a elementwise sign.
type signOp struct{ unop }

// String implements fmt.Stringer.
func (op signOp) String() string { return "Sign" }

// Do performs elementwise sign.
func (op signOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Sign(a, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise sign but with a preallocated return value.
// PreallocDo allows add to implement ops.PreallocOp.
func (op signOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Sign(a, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}                                           // DiffWRT returns {false, false} for sign
func (op signOp) DiffWRT(inputs int) []bool { return onefalse }
