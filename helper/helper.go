package helper

var version = "1.0.0"      // tidak bisa di akses diluar package
var Application = "golang" // bisa  di akses diluar package

func SayHello(name string) string {
	return "Hello " + name
}