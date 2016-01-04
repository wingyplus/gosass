package gosass

// #cgo CFLAGS: -I./libs/include
// #cgo LDFLAGS: -L./libs/lib -lsass
// #include <sass/context.h>
import "C"

func Render(scss string) ([]byte, error) {
	ctx := C.sass_make_data_context(C.CString(scss))
	compiler := C.sass_make_data_compiler(ctx)
	C.sass_compiler_parse(compiler)
	C.sass_compiler_execute(compiler)
	css := C.GoString(C.sass_context_get_output_string(ctx))
	C.sass_delete_compiler(compiler)
	return []byte(css), nil
}
