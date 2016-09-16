package utils

// CheckErr is a silly wrapper for error checking and panicking
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
