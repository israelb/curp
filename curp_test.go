package curp

import (
	"strings"
	"testing"
)

func TestFilterInappropriateWord(t *testing.T) {
	var RightWords = [...]string{"BXCA", "LXCO", "BXEI", "BXEY", "MXME", "CXCA", "MXMO",
		"CXCO", "MXAR", "CXGA", "MXAS", "CXGO", "MXON", "CXKA", "MXAR", "CXKO", "MXON",
		"CXGE", "MXCO", "CXGI", "MXKO", "CXJA", "MXLA", "CXJE", "MXLO", "CXJI", "NXCA",
		"CXJO", "NXCO", "CXLA", "PXDA", "CXLO", "PXDO", "FXLO", "PXNE", "FXTO", "PXPI",
		"GXTA", "PXTO", "GXEI", "PXPO", "GXEY", "PXTA", "JXTA", "PXTO", "JXTO", "QXLO",
		"KXCA", "RXTA", "KXCO", "RXBA", "KXGA", "RXBE", "KXGO", "RXBO", "KXKA", "RXIN",
		"KXKO", "SXNO", "KXGE", "TXTA", "KXGI", "VXCA", "KXJA", "VXGA", "KXJE", "VXGO",
		"KXJI", "VXKA", "KXJO", "VXEI", "KXLA", "VXEY", "KXLO", "WXEI", "LXLO", "WXEY",
		"LXCA"}

	for i, w := range inappropriateWords {
		if word := filterInappropriateWord(w); word != RightWords[i] {
			t.Errorf("invalid word %s", word)
		}
	}
}

func TestAddVerifiedDigit(t *testing.T) {
	if digit := addVerifiedDigit("BAAI810809HJCRCS0"); digit != "2" {
		t.Errorf("invalid digit %s", digit)
	}
	if digit := addVerifiedDigit("PICD861230HSLCRV0"); digit != "2" {
		t.Errorf("invalid digit %s", digit)
	}
	if digit := addVerifiedDigit("GECV901105HDFRHL0"); digit != "0" {
		t.Errorf("invalid digit %s", digit)
	}
	if digit := addVerifiedDigit("BARA860113MBSRMD0"); digit != "2" {
		t.Errorf("invalid digit %s", digit)
	}
	if digit := addVerifiedDigit("ROCR790513HZSBSG0"); digit != "1" {
		t.Errorf("invalid digit %s", digit)
	}
	if digit := addVerifiedDigit("OIRA901010HDFRZR0"); digit != "3" {
		t.Errorf("invalid digit %s", digit)
	}
}

func TestStripCtlAndExtFromUnicode(t *testing.T) {
	var origen = [...]string{"Ã", "À", "Á", "Ä", "Â", "È", "É", "Ë", "Ê", "Ì", "Í", "Ï", "Î",
		"Ò", "Ó", "Ö", "Ô", "Ù", "Ú", "Ü", "Û", "ã", "à", "á", "ä", "â",
		"è", "é", "ë", "ê", "ì", "í", "ï", "î", "ò", "ó", "ö", "ô", "ù",
		"ú", "ü", "û"}

	var destino = [...]string{"A", "A", "A", "A", "A", "E", "E", "E", "E", "I", "I", "I", "I",
		"O", "O", "O", "O", "U", "U", "U", "U", "a", "a", "a", "a", "a",
		"e", "e", "e", "e", "i", "i", "i", "i", "o", "o", "o", "o", "u",
		"u", "u", "u"}

	for i, w := range origen {
		if word := stripCtlAndExtFromUnicode(w); word != destino[i] {
			t.Errorf("invalid %s", word)
		}
	}
}

func TestIsValidSex(t *testing.T) {
	t.Log("validating sexs")

	if _, errSex := isValidSex("M"); errSex != nil {
		t.Errorf("invalid sex")
	}
	if _, errSex := isValidSex("H"); errSex != nil {
		t.Errorf("invalid sex")
	}
	if _, errSex := isValidSex("X"); errSex.Error() != "Sex initial is invalid, you have to use M or H" {
		t.Error(errSex)
	}
}

func TestValidState(t *testing.T) {
	t.Log("validating states")

	if _, errState := validState("JC"); errState != nil {
		t.Errorf("invalid state")
	}
	if _, errState := validState("X"); errState.Error() != "State is invalid" {
		t.Error(errState)
	}
}

func TestGetInitial(t *testing.T) {
	if initial := getInitial("maria"); initial != "M" {
		t.Errorf("error %s", initial)
	}
	if initial := getInitial("maria isabel"); initial != "I" {
		t.Errorf("error %s", initial)
	}
	if initial := getInitial("maria luis"); initial != "L" {
		t.Errorf("error %s", initial)
	}
	if initial := getInitial("maria jose"); initial != "J" {
		t.Errorf("error %s", initial)
	}
	if initial := getInitial("maria fernanda"); initial != "F" {
		t.Errorf("error %s", initial)
	}
	if initial := getInitial("maria angelica"); initial != "A" {
		t.Errorf("error %s", initial)
	}
	if initial := getInitial("ÑANDO"); initial != "X" {
		t.Errorf("error %s", initial)
	}
	if initial := getInitial("Maria ÑANDO"); initial != "X" {
		t.Errorf("error %s", initial)
	}
}

//BAAI810809HJCRCS02
func TestValidFirstLastName(t *testing.T) {
	if firstLastName := validFirstLastName("lopez"); firstLastName != "lopez" {
		t.Error("error!")
	}
	if firstLastName := validFirstLastName("RIVA PALACIO"); firstLastName != "RIVA" {
		t.Errorf("error! %s", firstLastName)
	}
}
func TestIsConpoundNameInvalid(t *testing.T) {
	if word := isConpoundNameInvalid("MC GREGOR"); word != true {
		t.Errorf("error %t", word)
	}
	if word := isConpoundNameInvalid("del romo"); word != true {
		t.Errorf("error %t", word)
	}
}
func TestGetBirthDate(t *testing.T) {
	year, birthDateFormatted := getBirthDate("1981-08-09")

	if year != 1981 && birthDateFormatted != "810809" {
		t.Error("error!")
	}
}
func TestGetHomonimia(t *testing.T) {
	if h := getHomonimia(1981); h != "0" {
		t.Error("error!")
	}

	if h := getHomonimia(2000); h != "A" {
		t.Error("error!")
	}
}
func TestGetFirstVowel(t *testing.T) {
	if h := getFirstVowel("LOPEZ"); h != "O" {
		t.Error("error!")
	}
	if h := getFirstVowel("ROMERO"); h != "O" {
		t.Error("error!")
	}
	if h := getFirstVowel("ICH"); h != "X" {
		t.Errorf("error! %s", h)
	}
	if h := getFirstVowel("ILL"); h != "X" {
		t.Errorf("error! %s", h)
	}
	if h := getFirstVowel("ACEVES"); h != "E" {
		t.Errorf("error! %s", h)
	}
	if h := getFirstVowel("BARBA"); h != "A" {
		t.Errorf("error! %s", h)
	}
	if h := getFirstVowel("HALL"); h != "A" {
		t.Errorf("error! %s", h)
	}
	if h := getFirstVowel("D/AMICO"); h != "X" {
		t.Errorf("error! %s", h)
	}
	if h := getFirstVowel("D-AMICO"); h != "X" {
		t.Errorf("error! %s", h)
	}
	if h := getFirstVowel("D.AMICO"); h != "X" {
		t.Errorf("error! %s", h)
	}
}

// func BenchmarkFirstVowel1(b *testing.B) { BenchmarkFirstVowel(b) }
func TestGetFirstConsonant(t *testing.T) {
	if consonant := getFirstConsonant("ALBERTO"); consonant != "L" {
		t.Errorf("error! %s", consonant)
	}
	if consonant := getFirstConsonant("RODRIGUEZ"); consonant != "D" {
		t.Errorf("error! %s", consonant)
	}
	if consonant := getFirstConsonant("OÑATE"); consonant != "X" {
		t.Errorf("error! %s", consonant)
	}
	if consonant := getFirstConsonant("PO"); consonant != "X" {
		t.Errorf("error! %s", consonant)
	}
	if consonant := getFirstConsonant("D/NAMO"); consonant != "X" {
		t.Errorf("error! %s", consonant)
	}
}

func TestNewCurp(t *testing.T) {
	if curp, _ := NewCurp("Israel", "barba", "Aceves", "H", "JC", "1981-08-09"); curp != "BAAI810809HJCRCS02" {
		t.Errorf("error curp %s", curp)
	}

	if _, err := NewCurp("Israel", "barba", "Aceves", "NOVALID", "JC", "1981-08-09"); err.Error() != "Sex initial is invalid, you have to use M or H" {
		t.Errorf("error curp %s", err.Error())
	}

	// _, err := NewCurp("Israel", "barba", "Aceves", "W", "JC", "1981-08-09")

	// if err != nil {
	// 	t.Errorf("error curp %s", err.Error())
	// }

	// if _, errState := validState("X"); errState.Error() != "State is invalid" {
	// 	t.Error(errState)
	// }
}

func TestGetFirstFourInitials(t *testing.T) {
	curp := &curp{
		name:           strings.ToUpper("carlos"),
		firstLastName:  strings.ToUpper("mc gregor"),
		secondLastName: strings.ToUpper("lopez"),
		sex:            strings.ToUpper("H"),
		stateCode:      strings.ToUpper("jc"),
		birthDate:      strings.ToUpper("1981-08-09"),
	}

	if result := curp.getFirstFourInitials(); result != "GELC" {
		t.Errorf("error 4 initials, expected: %s, and get: %s", "GELC", result)
	}
}

func BenchmarkCurp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCurp("Israel", "barba", "Aceves", "H", "JC", "1981-08-09")
	}
}
