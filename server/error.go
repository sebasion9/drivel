package main
import "fmt"

type DoRequestError struct {
    Err error
}
func (e *DoRequestError) Error() string {
    return fmt.Sprintf("[ERR] failed to send a request through http.Client: \n%s\n", e.Err)
}
type BuildRequestError struct {
    Method string
    Err error
}
func (e* BuildRequestError) Error() string {
    return fmt.Sprintf("[ERR] failed to build [%s] request: \n%s\n", e.Method, e.Err)
}
type NotOkResponseError struct {
    Body string
    Code int
}
func (e* NotOkResponseError) Error() string {
    return fmt.Sprintf("[ERR] response not ok: \n%d\nbody:\n%s\n", e.Code, e.Body)
}

func pretty_error(message string, err error) {
    fmt.Println("<!ERR")
    fmt.Printf("\t")
    fmt.Println(message)
    fmt.Printf("\t")
    fmt.Println(err)
    fmt.Println("ERR!>")
}
func res_ok(code int) bool {
    if (code > 199 && code < 300) {
        return true
	}
    return false
}
