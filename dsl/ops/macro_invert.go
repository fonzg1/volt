package ops

import (
	"context"

	"github.com/vim-volt/volt/dsl/ops/util"
	"github.com/vim-volt/volt/dsl/types"
)

func init() {
	opsMap[InvertOp.String()] = InvertOp
}

type invertOp struct {
	macroBase
}

// InvertOp is "$invert" operator
var InvertOp = &invertOp{macroBase("$invert")}

func (op *invertOp) InvertExpr(args []types.Value) (types.Value, error) {
	return op.macroInvertExpr(op.EvalExpr(context.Background(), args))
}

func (*invertOp) Bind(args ...types.Value) (types.Expr, error) {
	expr := types.NewExpr(ArrayOp, args, types.NewArrayType(types.AnyValue))
	return expr, nil
}

func (*invertOp) EvalExpr(ctx context.Context, args []types.Value) (types.Value, func(), error) {
	if err := util.Signature(types.AnyValue).Check(args); err != nil {
		return nil, NoRollback, err
	}
	val, err := args[0].Invert()
	return val, NoRollback, err
}
