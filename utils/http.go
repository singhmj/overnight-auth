package utils

func ErrorResponse(err error) map[interface{}]interface{} {
	return map[interface{}]interface{}{
		"error": err,
	}
}
