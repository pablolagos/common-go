package binaryutils

func Reverse(data []byte) []byte {

	result := []byte{}

	// Add bytes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	// Return new binary data.
	return result
}
