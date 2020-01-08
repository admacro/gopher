package salut

func Salut(gender int) string {
	switch gender {
	case 0:
		return "Bonjour Madame!"
	case 1:
		return "Bonjour Monsieur!"
	default:
		return "Bonjour!"
	}
}
