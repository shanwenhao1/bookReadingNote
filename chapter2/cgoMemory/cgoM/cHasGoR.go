package cgoM

/*
extern char* NewGoString(char* );
extern void FreeGoString(char* );
extern void PrintGoString(char* );

static void printString(char* s){
	char* gs = NewGoString(s);
	PrintGoString(gs);
	FreeGoString(gs);

}
*/
import "C"
import "unsafe"

/*
	C 长期持有Go指针对象(尽量不长期持有), 可通过将Go语言内存对象在Go语言空间映射为一个int类型的ID, 然后通过此ID来间接
访问和控制Go语言对象.
*/

//export NewGoString
func NewGoString(s *C.char) *C.char {
	gs := C.GoString(s)
	id := NewObjectId(gs)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export FreeGoString
func FreeGoString(p *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(p)))
	id.Free()
}

//export PrintGoString
func PrintGoString(s *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(s)))
	gs := id.Get().(string)
	print(gs)
}

func Run() {
	C.printString(C.CString("hello"))
}
