package gosass

// #cgo LDFLAGS: -lsass
// #include <sass/context.h>
import "C"

func Render(scss string) ([]byte, error) {
	ctx := C.sass_make_data_context(C.CString(scss))
	compiler := C.sass_make_data_compiler(ctx)
	C.sass_compiler_parse(compiler)
	C.sass_compiler_execute(compiler)
	ctxOut := C.sass_data_context_get_context(ctx)
	css := C.GoString(C.sass_context_get_output_string(ctxOut))
	C.sass_delete_compiler(compiler)
	C.sass_delete_data_context(ctx)
	return []byte(css), nil
}
