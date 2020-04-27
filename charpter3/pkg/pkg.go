package pkg

//var Id = 9527
var Id int

//var NameData [8]byte
var Name string // Go字符串是一种只读的引用类型

// Id变量的初始化在pkg_amd64.s中的汇编语言中
func AssemblyRun(){
	println(Id)
	println(Name)
}