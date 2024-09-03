package color

const (
	RESET = "\033[0m"
	RED = "\033[31m"
	GREEN = "\033[32;1m"
)


func ColorWrapper(color string, body string) string {
	return color + body + RESET
}