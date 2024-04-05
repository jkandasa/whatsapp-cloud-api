package types

// used in the context as key type
type ContextKey string

const (
	LoggerContextKey = ContextKey("logger_key")
)
