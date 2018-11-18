package interpreter

import (
    "errors"
    "fmt"
)

func (o *Object) GetBool() (bool, error) {
    if value, ok := o.Value.(bool); ok {
        return value, nil
    } else {
        return false, errors.New(fmt.Sprintf("Error: cant convert %v (%T) to bool", o.Value, o.Value))
    }
}
