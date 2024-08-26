package errs

import "fmt"

var NotFound = fmt.Errorf("not found")
var AlreadyClaimed = fmt.Errorf("daily reward already claimed")
