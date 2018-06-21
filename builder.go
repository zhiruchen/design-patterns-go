package design

type Options struct {
	op1 string
	op2 int32
	op3 float64
	op4 map[string]interface{}
	op5 interface{}
}

type ComplexObject struct {
	opts *Options
}

func NewComplexObject(opts ...ObjectOption) *ComplexObject {
	var objOpts = &Options{}
	for _, opt := range opts {
		opt(objOpts)
	}

	return &ComplexObject{opts: objOpts}
}

type ObjectOption func(o *Options)

func WithOp1(op1 string) ObjectOption {
	return func(opts *Options) {
		opts.op1 = op1
	}
}

func WithOp2(op2 int32) ObjectOption {
	return func(opts *Options) {
		opts.op2 = op2
	}
}

func WithOp3(op3 float64) ObjectOption {
	return func(opts *Options) {
		opts.op3 = op3
	}
}

func WithOp4(op4 map[string]interface{}) ObjectOption {
	return func(opts *Options) {
		opts.op4 = op4
	}
}

func WithOp5(op5 interface{}) ObjectOption {
	return func(opts *Options) {
		opts.op5 = op5
	}
}
