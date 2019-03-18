
//line aa_match.ragel:1
package ragel





//line aa_match.ragel:7

//line aa_match.go:12
const aamatch_start int = 1
const aamatch_first_final int = 5
const aamatch_error int = 0

const aamatch_en_main int = 1


//line aa_match.ragel:8


func (r *Stuff) AAMatch(data []byte) bool {
	cs, p, pe := 0, 0, len(data)

	
//line aa_match.go:27
	{
	cs = aamatch_start
	}

//line aa_match.go:32
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	}
	goto st_out
	st_case_1:
		if data[p] == 97 {
			goto st2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
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
		if data[p] == 97 {
			goto tr4
		}
		goto st0
tr4:
//line aa_match.ragel:14
 return true 
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line aa_match.go:91
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line aa_match.ragel:20


	return false
}

