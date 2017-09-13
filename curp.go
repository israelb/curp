package curp

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	ordinaryNames = [...]string{"MARIA", "MA", "MA.", "JOSE", "J", "J."}

	inappropriateWords = [...]string{"BACA", "LOCO", "BUEI", "BUEY", "MAME", "CACA", "MAMO",
		"CACO", "MEAR", "CAGA", "MEAS", "CAGO", "MEON", "CAKA", "MIAR", "CAKO", "MION",
		"COGE", "MOCO", "COGI", "MOKO", "COJA", "MULA", "COJE", "MULO", "COJI", "NACA",
		"COJO", "NACO", "COLA", "PEDA", "CULO", "PEDO", "FALO", "PENE", "FETO", "PIPI",
		"GETA", "PITO", "GUEI", "POPO", "GUEY", "PUTA", "JETA", "PUTO", "JOTO", "QULO",
		"KACA", "RATA", "KACO", "ROBA", "KAGA", "ROBE", "KAGO", "ROBO", "KAKA", "RUIN",
		"KAKO", "SENO", "KOGE", "TETA", "KOGI", "VACA", "KOJA", "VAGA", "KOJE", "VAGO",
		"KOJI", "VAKA", "KOJO", "VUEI", "KOLA", "VUEY", "KULO", "WUEI", "LILO", "WUEY",
		"LOCA"}

	// Official codes: https://es.wikipedia.org/wiki/ISO_3166-2:MX
	codeStates = [...]string{"AS", "BC", "BS", "CC", "CS", "CH", "CL", "CM", "DF", "DG",
		"GT", "GR", "HG", "JC", "MC", "MN", "MS", "NT", "NL", "OC", "PL", "QT",
		"QR", "SP", "SL", "SR", "TC", "TS", "TL", "VZ", "YN", "ZS", "NE"}

	characteresDigitVerified = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E",
		"F", "G", "H", "I", "J", "K", "L", "M", "N", "Ñ", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	conpoundNames = [...]string{"DA", "DAS", "DE", "DEL", "DER", "DI", "DIE", "DD", "EL", "LA", "LOS", "LAS", "LE", "LES", "MAC", "MC", "VAN", "VON", "Y"}
)

type curp struct {
	name           string
	firstLastName  string
	secondLastName string
	sex            string
	stateCode      string
	// year, month, day (xxxx-xx-xx)
	birthDate string

	// Optional, It values is used in order to avoid duplicates, it is assign by the Goverment.
	// By default is 0 if the datebirth is smaller or equal to 1999, or the value is A if the value is greatest than 2000
	homonimia string
}

// NewCurp generates a new curp
func NewCurp(name, firstLastName, secondLastName, sex, stateCode, birthDate string) string {
	curp := &curp{
		name:           name,
		firstLastName:  firstLastName,
		secondLastName: secondLastName,
		sex:            sex,
		stateCode:      stateCode,
		birthDate:      birthDate,
	}

	return curp.generate()
}

func (c curp) generate() string {
	// call all the methods here
	var curp bytes.Buffer

	firstLastName := validFirstLastName(c.firstLastName)

	p01 := getInitial(firstLastName)
	p02 := getFirstVowel(firstLastName)
	p03 := getInitial(c.secondLastName)
	p04 := getInitial(c.name)
	_, birthDate := getBirthDate(c.birthDate)
	// homonimia := getHomonimia(year)

	pos1_4 := filterInappropriateWord(p01 + p02 + p03 + p04)
	curp.WriteString(pos1_4)
	curp.WriteString(birthDate)

	if isValidSex(c.sex) {
		curp.WriteString(c.sex)
	}

	if validState(c.stateCode) {
		curp.WriteString(c.stateCode)
	}

	// posicion_14_16 = [
	// 	primerConsonante(param.apellido_paterno),
	// 	primerConsonante(param.apellido_materno),
	// 	primerConsonante(nombre_a_usar)
	//   ].join('');

	return curp.String()
}

func validFirstLastName(firstLastName string) string {
	firstLastName = strings.ToUpper(firstLastName)
	firstLastNames := strings.SplitAfter(firstLastName, " ")

	return strings.Replace(firstLastNames[0], " ", "", -1)
}

func filterInappropriateWord(word string) string {
	word = strings.ToUpper(word)
	var re = regexp.MustCompile(`^(\w)\w`)

	for _, w := range inappropriateWords {
		w = strings.ToUpper(w)
		if w == word {
			word = re.ReplaceAllString(word, "${1}X")
			break
		}
	}
	return word
}

// Example:
// B   A   A  I   8  1  0   8  0  9  H   J   C   R   C   S  0    =  2
// 11, 10, 10, 18, 8, 1, 0,  8, 0, 9, 17, 19, 12, 28, 12, 29, 0

// prev + (value * (18 -  index));

// Example first word
// 11 + (10 *(18-1))
// 11 + (10*17)
// 11 + 170
// result word B: 181

// sumarize:
// total =   1708
// digit = (10 - (total % 10));
func addVerifiedDigit(curp string) string {
	splitCurp := strings.SplitAfter(curp, "")
	values := [17]int{}
	var total int

	for indexCurp, letter := range splitCurp {
		for index, digit := range characteresDigitVerified {
			if letter == digit {
				values[indexCurp] = index
			}
		}
	}

	for idx, val := range values {
		if idx == 0 {
			total = (val * (18 - (1)))
		} else {
			total += values[idx-1] + (val * (18 - (idx + 1)))
		}
	}

	digit := (10 - (total % 10))

	if digit == 10 {
		digit = 0
	}

	return strconv.Itoa(digit)
}

func validState(state string) bool {
	state = strings.ToUpper(state)

	isValid := false

	for _, num := range codeStates {
		if num == state {
			isValid = true
		}
	}

	return isValid
}

func isValidSex(sex string) bool {
	if sex == "M" || sex == "H" {
		return true
	}
	return false
}

// Funcion que extrae la inicial del primer nombre, o, si tiene mas de 1 nombre Y el primer
// nombre es uno de la lista de nombres comunes, la inicial del segundo nombre
// @param {string} nombre - String que representa todos los nombres (excepto los apellidos) separados por espacio
func getInitial(fullName string) string {
	fullName = strings.ToUpper(fullName)
	names := strings.SplitAfter(fullName, " ")
	var word string

	if len(names) > 1 {
		for _, name := range ordinaryNames {
			if name == strings.TrimSpace(names[0]) {
				word = validInitial(names[1])
				return word[:1]
			}
		}
	}

	word = validInitial(names[0])

	return word[:1]
}

func validInitial(initial string) string {
	myInitial := strings.ToUpper(initial)
	var re = regexp.MustCompile(`(^Ñ)`)
	word := re.ReplaceAllString(myInitial, "${2}X")
	return word
}

func getFirstVowel(word string) string {
	word = strings.ToUpper(word)
	word = word[1:len(word)]
	var reConstant = regexp.MustCompile(`[BCDFGHJKLMNÑPQRSTVWXYZ]`)
	var reVowel = regexp.MustCompile(`[AEIOU]`)

	vowels := reConstant.ReplaceAllString(word, "${1}")

	if reVowel.FindStringIndex(word) == nil {
		return "X"
	}

	return string(vowels[0])
}

func getFirstConsonant(word string) string {
	word = strings.ToUpper(word)
	word = word[1:len(word)]
	var reVowel = regexp.MustCompile(`[AEIOU]`)
	var reConstant = regexp.MustCompile(`[BCDFGHJKLMNÑPQRSTVWXYZ]`)
	var eneEr = regexp.MustCompile(`(^Ñ)`)

	consonant := reVowel.ReplaceAllString(word, "${1}")

	if reConstant.FindStringIndex(consonant) == nil {
		return "X"
	}

	consonant = eneEr.ReplaceAllString(consonant, "${2}X")

	return string(consonant[0])
}

func getBirthDate(birthDate string) (int, string) {
	var date bytes.Buffer
	splited := strings.Split(birthDate, "-")
	year, _ := strconv.Atoi(splited[0])

	date.WriteString(splited[0][2:])
	date.WriteString(splited[1])
	date.WriteString(splited[2])

	return year, date.String()
}

func getHomonimia(year int) string {
	if year > 1999 {
		return "A"
	}
	return "0"
}

// Advanced Unicode normalization and filtering,
// see http://blog.golang.org/normalization and
// http://godoc.org/golang.org/x/text/unicode/norm for more
// details.
func stripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	// This Transformer could also trivially be applied as an io.Reader
	// or io.Writer filter to automatically do such filtering when reading
	// or writing data anywhere.
	str, _, _ = transform.String(t, str)
	return str
}
