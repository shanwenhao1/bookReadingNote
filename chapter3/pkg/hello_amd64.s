TEXT ·Hello(SB), $16-0
    // MQVQ命令将helloWorld对应的字符串头部结构体的16个字节复制到栈指针SP对应的16字节的空间
    MOVQ ·helloWorld+0(SB), AX; MOVQ AX, 0(SP)
    MOVQ ·helloWorld+8(SB), BX; MOVQ BX, 8(SP)
    CALL runtime·printstring(SB)
    CALL runtime·printnl(SB)
    RET
