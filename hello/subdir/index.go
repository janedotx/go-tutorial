// this is for a demo of importing from subdirectories
// in the same module. It doesn't do anything, but it shows how to import from a subdirectory
package subdir

func SubdirFunc() string {
	return "this is a function in the subdir package"
}

// you can never pass nil to this,
// bc that only works when you're passing by
// reference, but in this case we pass by value
func TakeNilExperiment(d string) string {
	if d == "" {
		return "d is empty string"
	}
	return "d is a non-empty string"
}