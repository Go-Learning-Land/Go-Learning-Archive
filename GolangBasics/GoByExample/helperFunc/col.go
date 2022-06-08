package helperFunc

import "fmt"

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func Info(s string) {
	fmt.Printf(InfoColor, s)
}
func Notice(s string) {
	fmt.Printf(NoticeColor, s)
}
func Warning(s string) {
	fmt.Printf(WarningColor, s)
}
func Error(s string) {
	fmt.Printf(ErrorColor, s)
}
func Debug(s string) {
	fmt.Printf(DebugColor, s)
}
