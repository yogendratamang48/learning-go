package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "DELETE"}
	switch r.Method {
	case "GET":
		println("Get Request")
	case "DELETE":
		println("DELETE Request")
	case "POST":
		println("POST Request")
	case "PUT":
		println("PUT Request")
	default:
		println("Unhandled..")
	}

}
