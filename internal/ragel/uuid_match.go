
//line uuid_match.ragel:1
package ragel





//line uuid_match.ragel:7

//line uuid_match.go:12
const uuid_match_start int = 1
const uuid_match_first_final int = 33
const uuid_match_error int = 0

const uuid_match_en_main int = 1


//line uuid_match.ragel:8



func (r *Stuff) UUIDMatch(data []byte) bool {
	cs, p, pe := 0, 0, len(data)

    
//line uuid_match.go:28
	{
	cs = uuid_match_start
	}

//line uuid_match.go:33
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 0:
		goto st_case_0
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	}
	goto st_out
	st_case_1:
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 45 {
			goto st10
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 45 {
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 45 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 45 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		goto tr32
tr32:
//line uuid_match.ragel:15
 return true 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line uuid_match.go:322
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line uuid_match.ragel:21


	return false
}