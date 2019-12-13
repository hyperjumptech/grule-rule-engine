package examples

import "fmt"

// MyPoGo serve as example plain Plai Old Go Object.
type MyPoGo struct {
}

// GetStringLength will return the length of provided string argument
func (p *MyPoGo) GetStringLength(sarg string) int {
	return len(sarg)
}

// Compare will compare the equality between the two string.
func (p *MyPoGo) Compare(t1, t2 string) bool {
	fmt.Println(t1, t2)
	return t1 == t2
}

// User is an example user struct.
type User struct {
	Name string
	Age  int
	Male bool
}
