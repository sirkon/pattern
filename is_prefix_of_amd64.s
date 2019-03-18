TEXT ·isIndexOf(SB), $0-40
    MOVQ    pattern + 0(FP), SI
    MOVQ    mask + 0(FP), DI
    MOVQ    source + 0(FP), DX
    MOVQ    length + 0(FP), CX
    XOR     AX, AX
    MOVQ    AX, length + 32(FP) // set up default result as false
    XORPS   xmm1, xm11          // need to have it zero

    // We are comparing using SSE 128 bit registers here. docts means "double octs"
docts:
    CMP     CX, 16
    JL      octs
    MOVDQU  (DI), xmm0
    ANDPS   (DX), xmm0
    XORPS   (SI), xmm0
    PTEST   xmm0, xmm0
    JNZ     exit
    ADDQ    DI, 16
    ADDQ    DX, 16
    ADDQ    SI, 16
    SUBQ    CX, 16
    JMP     docts

    // less than 16 bytes left, using 8 byte words to deal with
    OR      CX, CX
octs:
    JZ      success
    MOVQ    (DI), AX
    ANDQ    (DX), AX
    XORQ    (SI), AX
    JNZ     exit
    ADDQ    DI, 8
    ADDQ    DX, 8
    ADDQ    SI, 8
    SUBQ    CX, 8
    JMP     octs

success:
// making result true
    MOVQ    AX, 1
    MOVQ    AX, length + 32(FP)

exit:
    RET




