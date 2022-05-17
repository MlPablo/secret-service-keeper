package handlers

const MAXMESSAGELENGTH = 1024

func validateMessageLength(msg string) bool {
	return MAXMESSAGELENGTH >= len(msg)
}
