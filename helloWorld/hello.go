package helloWorld

const (
	spanish = "Spanish"
	french  = "French"
)

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

var null = ""

func Hello(name string, lang string) string {

	if name == null {
		name = "World"
	}

	return greetingPrefix(lang) + name
}
func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
