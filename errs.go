package refer

// import (
// 	"errors"
// 	"fmt"
// )

// var errCallEmptyMethods = errors.New("call the struct is empty methods")

// func IsErrCallEmptyMethods(err error) bool {
// 	return errors.Is(err, errCallEmptyMethods)
// }

// type errCallNotHaveMethod struct {
// 	tip string
// }

// func (e errCallNotHaveMethod) Error() string {
// 	return fmt.Sprintf("call the struct is not have method: %+v", e.tip)
// }

// func IsErrCallNotHaveMethod(err error) bool {
// 	_, ok := err.(errCallNotHaveMethod)
// 	return ok
// }
