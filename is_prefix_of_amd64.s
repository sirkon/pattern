TEXT ·isPrefixOf(SB), $0-40
    MOVQ    pattern + 0(FP), SI   // SI = pointer to pattern
    MOVQ    pattern + 8(FP), DI   // DI = pointer to mask
    MOVQ    pattern + 16(FP), DX  // DX = pointer to a text to match
    MOVQ    pattern + 24(FP), CX
    XORQ    AX, AX
    MOVQ    AX, pattern + 32(FP)  // set up default result as false
//    XORPS   X1, X1              // need to have it zero

    // We are compare using SSE 128 bit registers here. "docts" means "double octs"
docts:
    CMPQ    CX, $16
    JL      octs
    MOVOU   (DX), X0
    PAND    (DI), X0
    PXOR    (SI), X0
    PTEST   X0, X0
    JNZ     exit
    ADDQ    $16, DI
    ADDQ    $16, DX
    ADDQ    $16, SI
    SUBQ    $16, CX
    JMP     docts

octs:
    // less than 16 bytes left. I mean 8 bytes left, using 8 byte (64bit) general purpose register to deal with
    ORQ     CX, CX
    JZ      success

    MOVQ    (DI), AX
    ANDQ    (DX), AX
    XORQ    (SI), AX
    JNZ     exit

success:
// making result true
    MOVQ    $1, AX
    MOVQ    AX, length + 32(FP)

exit:
    RET




