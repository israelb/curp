package curp

import (
	"testing"
)

func TestWord(t *testing.T) {
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
