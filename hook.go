//go:build amd64
// +build amd64

package flag

import (
	oflag "flag"

	"github.com/agiledragon/gomonkey"
)

/*
inline 了 flag 的函数， 这些hook 无效
go build 的时候需要加上 -gcflags='flag=-l'
*/

func init() {
	gomonkey.ApplyFunc(oflag.BoolVar, BoolVar)
	gomonkey.ApplyFunc(oflag.StringVar, StringVar)
	gomonkey.ApplyFunc(oflag.IntVar, IntVar)
	gomonkey.ApplyFunc(oflag.Int64Var, Int64Var)
	gomonkey.ApplyFunc(oflag.Uint64Var, Uint64Var)

	gomonkey.ApplyFunc(oflag.Bool, Bool)
	gomonkey.ApplyFunc(oflag.String, String)
	gomonkey.ApplyFunc(oflag.Int, Int)
	gomonkey.ApplyFunc(oflag.Int64, Int64)
	gomonkey.ApplyFunc(oflag.Uint64, Uint64)
	//gomonkey.ApplyFunc(oflag.Func, Func)
}
