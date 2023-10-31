package rkw

import (
	"strings"
)

// Lang represents a programming language
type Lang string

const (
	CPP Lang = "cpp" // C++
	PY  Lang = "py"  // Python
	TS  Lang = "ts"  // TypeScript
	JS  Lang = "js"  // JavaScript
	GO  Lang = "go"  // Go
	UE  Lang = "ue"  // Unreal Engine C++
	QT  Lang = "qt"  // Qt C++
)

// DisplayName returns the display name of the language
func (l Lang) DisplayName() string {
	switch l {
	case CPP:
		return "C++"
	case PY:
		return "Python"
	case TS:
		return "TypeScript"
	case JS:
		return "JavaScript"
	case GO:
		return "Go"
	case UE:
		return "Unreal Engine C++"
	case QT:
		return "Qt C++"
	default:
		return string(l)
	}
}

func cppReservedKeywords() []string {
	return []string{
		"alignas", "alignof", "and", "and_eq",
		"asm", "auto", "bitand", "bitor",
		"bool", "break", "case", "catch",
		"char", "char8_t", "char16_t", "char32_t",
		"class", "compl", "concept", "const",
		"consteval", "constexpr", "const_cast", "continue",
		"co_await", "co_return", "co_yield", "decltype",
		"default", "delete", "do", "double",
		"dynamic_cast", "else", "enum", "explicit",
		"export", "extern", "false", "float",
		"for", "friend", "goto", "if",
		"inline", "int", "long", "mutable",
		"namespace", "new", "noexcept", "not",
		"not_eq", "nullptr", "operator", "or",
		"or_eq", "private", "protected", "public",
		"register", "reinterpret_cast", "requires", "return",
		"short", "signed", "sizeof", "static",
		"static_assert", "static_cast", "struct", "switch",
		"synchronized", "template", "this", "thread_local",
		"throw", "true", "try", "typedef",
		"typeid", "typename", "union", "unsigned",
		"using", "virtual", "void", "volatile",
		"wchar_t", "while", "xor", "xor_eq",
	}
}

func pyReservedKeywords() []string {
	return []string{
		"false", "none", "true", "and",
		"as", "assert", "async", "await",
		"break", "class", "continue", "def",
		"del", "elif", "else", "except",
		"finally", "for", "from", "global",
		"if", "import", "in", "is",
		"lambda", "nonlocal", "not", "or",
		"pass", "raise", "return", "try",
		"while", "with", "yield",
	}
}

func tsReservedKeywords() []string {
	return []string{
		"abstract", "any", "as", "assert",
		"async", "await", "boolean", "break",
		"byte", "case", "catch", "char",
		"class", "const", "continue", "debugger",
		"default", "delete", "do", "double",
		"else", "enum", "export", "extends",
		"false", "final", "finally", "float",
		"for", "function", "goto", "if",
		"implements", "import", "in", "instanceof",
		"int", "interface", "is", "long",
		"module", "namespace", "native", "new",
		"null", "number", "package", "private",
		"protected", "public", "return", "short",
		"static", "super", "switch", "synchronized",
		"this", "throw", "throws", "transient",
		"true", "try", "typeof", "var",
		"void", "volatile", "while", "with",
		"yield",
	}
}

func jsReservedKeywords() []string {
	return []string{
		"abstract", "arguments", "await", "boolean",
		"break", "byte", "case", "catch",
		"char", "class", "const", "continue",
		"debugger", "default", "delete", "do",
		"double", "else", "enum", "eval",
		"export", "extends", "false", "final",
		"finally", "float", "for", "function",
		"goto", "if", "implements", "import",
		"in", "instanceof", "int", "interface",
		"let", "long", "native", "new",
		"null", "package", "private", "protected",
		"public", "return", "short", "static",
		"super", "switch", "synchronized", "this",
		"throw", "throws", "transient", "true",
		"try", "typeof", "var", "void",
		"volatile", "while", "with", "yield",
	}
}

func goReservedKeywords() []string {
	return []string{
		"break", "case", "chan", "const",
		"continue", "default", "defer", "else",
		"fallthrough", "for", "func", "go",
		"goto", "if", "import", "interface",
		"map", "package", "range", "return",
		"select", "struct", "switch", "type",
		"var",
	}
}

func unrealCPlusPlusKeywords() []string {
	return []string{
		"uclass", "ufunction", "uproperty",
		"aactor", "uobject", "uactorcomponent",
		// Additional Unreal Engine-specific keywords/macros go here
	}
}

func qtReservedKeywords() []string {
	return []string{
		"q_object", "signals", "slots", "q_property",
		"q_declare_interface", "q_interfaces", "q_enum", "q_flag",
		// Additional Qt Framework-specific keywords/macros go here
	}
}

var (
	// map[lang][]keywords
	reservedKeywordsPerLang = makeReservedKeywordsPerLang()
	// map[keyword][]lang
	reservedKeywords = makeReservedKeywords()
)

func makeReservedKeywordsPerLang() map[Lang][]string {
	m := make(map[Lang][]string)
	m[CPP] = cppReservedKeywords()
	m[PY] = pyReservedKeywords()
	m[TS] = tsReservedKeywords()
	m[JS] = jsReservedKeywords()
	m[GO] = goReservedKeywords()
	m[UE] = append(cppReservedKeywords(), unrealCPlusPlusKeywords()...)
	m[QT] = append(cppReservedKeywords(), qtReservedKeywords()...)
	return m
}

// makes a map of with keywords as keys and languages as values
func makeReservedKeywords() map[string][]Lang {
	m := make(map[string][]Lang)
	for lang, keywords := range reservedKeywordsPerLang {
		for _, keyword := range keywords {
			m[keyword] = append(m[keyword], lang)
		}
	}
	return m
}

// IsKeywordReservedInLang returns true if the word is a reserved keyword in the given language
func IsKeywordReservedInLang(lang Lang, word string) bool {
	word = strings.ToLower(word)
	keywords, ok := reservedKeywordsPerLang[lang]
	if !ok {
		return false
	}
	for _, keyword := range keywords {
		if keyword == word {
			return true
		}
	}
	return false
}

// IsKeywordReserved returns true if the word is a reserved keyword in any language
// and returns the list of languages in which the word is a reserved keyword
func IsKeywordReserved(word string) ([]Lang, bool) {
	word = strings.ToLower(word)
	langs, ok := reservedKeywords[word]
	if !ok {
		return nil, false
	}
	return langs, true
}

// EscapeKeyword returns the escaped version of the given word
func EscapeKeyword(word string) string {
	return word + "_"
}

// CheckName checks if the given name is a reserved keyword in any language
// and logs a warning if it is
func CheckName(name string, scope string) {
	langs, ok := IsKeywordReserved(name)
	if ok {
		log.Warn().Msgf("%s: name %s is a reserved keyword in %s", scope, name, langs)
	}
}

// CheckAndEscapeName checks if the given name is a reserved keyword in any language
// and returns the escaped version of the name if it is.
func CheckAndEscapeName(name string, scope string) string {
	langs, ok := IsKeywordReserved(name)
	if ok {
		log.Warn().Msgf("%s: name %s is a reserved keyword in %s", scope, name, langs)
		return EscapeKeyword(name)
	}
	return name
}
