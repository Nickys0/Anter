package Anter

import "fmt"

const(
	NO_COLOR_ASCII_CODE = 	"\033[0m"
	BLACK_ASCCI_CODE =		"\033[0;30m"
	RED_ASCII_CODE = 		"\033[0;31m"
	GREEN_ASCII_CODE =		"\033[0;32m"
	YELLOW_ASCII_CODE =		"\033[0;33m"
	BLUE_ASCII_CODE =		"\033[0;34m"
	MAGENTA_ASCII_CODE =	"\033[0;35m"
	CYAN_ASCCI_CODE = 		"\033[0;36m"
	WHITE_ASCCI_CODE = 		"\033[0;37m"
)

func make_color(c, txt string) string{
	return fmt.Sprintf("%s%s%s", c, txt, NO_COLOR_ASCII_CODE)	
}

func BlackTxt(txt string) string{
	return make_color(BLACK_ASCCI_CODE, txt)
}

func RedTxt(txt string) string{
	return make_color(RED_ASCII_CODE, txt)
}

func GreenTxt(txt string) string{
	return make_color(GREEN_ASCII_CODE, txt)
}

func YellowTxt(txt string) string{
	return make_color(YELLOW_ASCII_CODE, txt)
}

func BlueTxt(txt string) string{
	return make_color(BLUE_ASCII_CODE, txt)
}

func MagentaTxt(txt string) string{
	return make_color(MAGENTA_ASCII_CODE, txt)
}

func CyanTxt(txt string) string{
	return make_color(CYAN_ASCCI_CODE, txt)
}

func WhiteTxt(txt string) string{
	return make_color(WHITE_ASCCI_CODE, txt)
}
