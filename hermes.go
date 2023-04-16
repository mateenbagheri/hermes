package hermes

import "fmt"

func NewHermesLogger(str string) hermesLogger {
	return hermesLogger{
		String: str,
	}
}

type hermesLogger struct {
	String string
}

func (h hermesLogger) Print() {
	fmt.Println("salam", h.String)
}
