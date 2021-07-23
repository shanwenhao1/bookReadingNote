//GOLOBL symbol(SB), width   中GOLOBL命令用于将符号导出, symbol对应汇编符号名字, width为符号对应内存大小
GLOBL ·Id(SB),$8

//DATA symbol+offset(SB)/width, value   DATA命令用于初始化包变量, symbol为变量对应标识符, offset是符号开始地址的偏移量, width是初始化内存的宽度大小, value是初始化的值
DATA ·Id+0(SB)/1,$0x37
DATA ·Id+1(SB)/1,$0x25  //Id变量初始化为十六进制的0x2537, 对应十进制的9527

DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00



//GLOBL ·NameData(SB),$8
//DATA ·NameData(SB)/8,$"gopher"

// 定义·Name符号内存大小为16字节
//GLOBL ·Name(SB),$16
// 8字节用于·NameData符号对应的地址初始化
//DATA ·Name+0(SB)/8,$·NameData(SB)
// 8字节为常量6, 表示字符串长度
//DATA ·Name+8(SB)/8,$6


// .Name符号前16个字节对应reflect.StringHeader结构体, Data部分对应·Name+16(SB)  (表示数据的地址为Name符号往后偏移16个字节位置)
GLOBL ·Name(SB),$24

DATA ·Name+0(SB)/8,$·Name+16(SB)
DATA ·Name+8(SB)/8,$6
DATA ·Name+16(SB)/8,$"gopher"
