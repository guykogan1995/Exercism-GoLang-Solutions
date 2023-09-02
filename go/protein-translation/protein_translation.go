package protein

import "errors"

const (
	MET = iota + 1
	PHE
	LEU
	SER
	TYR
	CYS
	TRY
	STOP
)

var m = map[string]int{
	"AUG": MET,
	"UUU": PHE, "UUC": PHE,
	"UUA": LEU, "UUG": LEU,
	"UCU": SER, "UCC": SER, "UCA": SER, "UCG": SER,
	"UAU": TYR, "UAC": TYR,
	"UGU": CYS, "UGC": CYS,
	"UGG": TRY,
	"UAA": STOP, "UAG": STOP, "UGA": STOP,
}

var ErrStop = errors.New("stop on an empty string")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	var sArray []string
	var err error
	var s string
	for i := 0; i < len(rna); i = i + 3 {
		s, err = FromCodon(rna[i : i+3])
		if err == ErrInvalidBase {
			return []string{}, ErrInvalidBase
		} else if err == ErrStop {
			return sArray, nil
		}
		sArray = append(sArray, s)
	}
	return sArray, nil
}

func FromCodon(codon string) (string, error) {
	switch m[codon] {
	case MET:
		return "Methionine", nil
	case PHE:
		return "Phenylalanine", nil
	case LEU:
		return "Leucine", nil
	case SER:
		return "Serine", nil
	case TYR:
		return "Tyrosine", nil
	case CYS:
		return "Cysteine", nil
	case TRY:
		return "Tryptophan", nil
	case STOP:
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
