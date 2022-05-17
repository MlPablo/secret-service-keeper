package handlers

const MAXMESSAGELENGTH = 1024
const MAXTTL = 86400

func validateMessageLength(msg string) bool {
	return MAXMESSAGELENGTH >= len(msg)
}

func validateTTLDuration(ttl int) bool {
	return MAXTTL >= ttl
}
