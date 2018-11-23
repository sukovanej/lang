package interpreter

import (
    "fmt"
)

func (o *Object) GetBool(ast *AST) (bool, *RuntimeError) {
    if value, ok := o.Value.(bool); ok {
        return value, nil
    } else {
        return false, NewRuntimeError(fmt.Sprintf("Cant convert %v (%T) to bool", o.Value, o.Value), ast.Value)
    }
}
