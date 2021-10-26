package main

const prefixBr = "Oi, "
const prefixEu = "Hello, "
const prefixEs = "Ei, "
const prefixFr = "Hi, "

// func main() {
// 	fmt.Println(Hello("World", ""))
// }

func Hello(str string, lang string) string {
	if str == "" {
		str = "World"
	}

	return prefix(lang) + str
}

func prefix(lang string) (prefix string) {
	switch lang {
	case "inglex":
		prefix = prefixEu
	case "espanhol":
		prefix = prefixEs
	case "frances":
		prefix = prefixFr
	case "brasileiro":
		prefix = prefixBr
	default:
		prefix = prefixEu
	}
	return prefix
}
