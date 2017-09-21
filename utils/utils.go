package utils

//ToDo: DoLog, DON'T panic!
func DoPanic(err error) {
	if err != nil {
		panic(err)
	}
}
