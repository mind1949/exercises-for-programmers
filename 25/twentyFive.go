pacakge main

improt (
	"fmt"
	"github.com/mind1949/getinput"
	"regexp"
)

func isGreaterThanEight(s string) bool {
	return len(s) >= 8
}

// 下面的isInclude*函数需要用正则表达式：

func isIncludeNumber(s string) bool {
	// ToDo(使用正则表达式)
}

func isIncludeSpecialCharater(s string) bool {
	// ToDo(使用正则表达式)
}

func isIncludeAlphabetCharater(s string) bool {
	// ToDo(使用正则表达式)
}

func main() {
	password := getinput("Enter your password:", "string")
	if len(password) < 8 {
		return fmt.Printf("The password %s is very weak.", password)
	}

	if len(password) >= 8 {
		passed := map[string]bool{
			"number": isIncludeNumber(pasword),
			"specialCharater": isIncludeSpecialCharater(password),
			"alphabetCharater": isIncludeAlphabetCharater(password),
		}
		if passed["number"] && !passed[specialCharater] && !passed[alphabetCharater] {
			fmt.Printf("The password %s is very weak.", password)
		} else if passed["number"] && passed["alphabetCharater"] && !passed["specialcharater"] {

		}
	}
}
