package base_converter

type Base int

const (
	BASE_UNDEFINED   Base = 0
	BASE_BINARY      Base = 2
	BASE_DECIMAL     Base = 10
	BASE_HEXADECIMAL Base = 16
)

func ParseBase(val int) Base {
	switch val {
	case 2:
		return BASE_BINARY
	case 10:
		return BASE_DECIMAL
	case 16:
		return BASE_HEXADECIMAL
	default:
		return BASE_UNDEFINED
	}
}

func (b Base) String() string {
	switch b {
	case BASE_BINARY:
		return "2"
	case BASE_DECIMAL:
		return "10"
	case BASE_HEXADECIMAL:
		return "16"
	default:
		return ""
	}
}

func AllBasesInStrings() []string {
	return []string{
		BASE_BINARY.String(),
		BASE_DECIMAL.String(),
		BASE_HEXADECIMAL.String(),
	}
}
