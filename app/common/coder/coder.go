package coder

import (
	"fmt"
	"github.com/yanggelinux/cattle/common/errors"
	"net/http"
	"sync"
)

var (
	unknownCoder DefaultCoder = DefaultCoder{1, http.StatusInternalServerError, "服务端内部错误"}
)

type Coder interface {
	// HTTP status that should be used for the associated error code.
	HTTPStatus() int

	// External (user) facing error text.
	String() string

	// Code returns the code of the coder
	Code() int
}

type DefaultCoder struct {
	// C refers to the integer code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Msg string
}

func (coder DefaultCoder) Code() int {
	return coder.C

}

// String implements stringer. String returns the external error message,
// if any.
func (coder DefaultCoder) String() string {
	return coder.Msg
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder DefaultCoder) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}

	return coder.HTTP
}

func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}

	if v, ok := errors.ParseWithCoder(err); ok {
		if coder, ok := codes[v.Code()]; ok {
			return coder
		}
	}

	return unknownCoder
}

// IsCode reports whether any error in err's chain contains the given error code.
func IsCode(err error, code int) bool {
	if v, ok := errors.ParseWithCoder(err); ok {
		if v.Code() == code {
			return true
		}

		if v.Cause() != nil {
			return IsCode(v.Cause(), code)
		}
		return false
	}

	return false
}

var codes = map[int]Coder{}
var codeMux = &sync.Mutex{}

func Register(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved by `errors` as unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	codes[coder.Code()] = coder
}

func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("code '0' is reserved by 'github.com/marmotedu/errors' as ErrUnknown error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	codes[coder.Code()] = coder
}

func init() {
	codes[unknownCoder.Code()] = unknownCoder
}
