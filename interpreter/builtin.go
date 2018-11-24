package interpreter

import (
    "fmt"
    "os"
    "path"
    "path/filepath"
)

func createStringObject(value string) (*Object) {
    result, err := NewStringObject(value)
    if err != nil { panic("") }

    return result
}

func CopyObject(value *Object) *Object {
    return NewObject(value.Type, value.Value, value.Meta, CopySlots(value.Slots))
}

func CopyArguments(arguments [](*Object)) [](*Object) {
    newArguments := make([](*Object), len(arguments))

    for index, value := range arguments {
        newArguments[index] = CopyObject(value)
    }

    return newArguments
}

func CopySlots(slots map[string](*Object)) map[string](*Object) {
    newSlots := make(map[string](*Object), len(slots))

    for symbol, value := range slots {
        newSlots[symbol] = CopyObject(value)
    }

    return newSlots
}

func BuiltInNewInstance(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    newArguments := CopyArguments(arguments)

    if initCallable, ok := newArguments[0].Slots["__init__"]; ok {
        initCallable.Slots["__call__"].Value.(ObjectCallable)(
            newArguments,
            scope,
            ast,
        )
    }

    newArguments[0].Type = TYPE_OBJECT

    return newArguments[0], nil
}

var MetaObject = &Object{Type: TYPE_OBJECT, Slots: map[string](*Object) {
    "__string__": createStringObject("<type object>"),
    "__call__": NewObject(TYPE_OBJECT, ObjectCallable(BuiltInNewInstance), nil, nil),
}}

func (object *Object) GetMetaObject() *Object {
    if object.Meta == nil {
        return MetaObject
    }
    return object.Meta
}

func (object *Object) GetAttribute(name string, ast *AST) (*Object, *RuntimeError) {
    var result *Object
    var found bool

    nextObject := object

    for nextObject != nil {
        result, found = nextObject.Slots[name]
        if found { break }

        nextObject = nextObject.Parent
    }

    if !found { return nil, NewRuntimeError("symbol " + name + " not found", ast.Value) }

    return result, nil
}

func NewObject(objectType ObjectType, value interface{}, meta *Object, slots map[string](*Object)) *Object {
    return &Object{
        Meta: meta,
        Value: value,
        Type: objectType,
        Slots: slots,
        Parent: nil,
    }
}

var NumberMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type number>") })
var FloatMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type float>") })
var StringMetaObject = NewObject(TYPE_OBJECT, nil, nil, nil)
var ListMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type list>") })
var MapMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type map>") })
var TupleMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type tuple>") })
var BoolMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) {})

var NilObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("nil") })

var TrueObject = NewObject(TYPE_BOOL, true, BoolMetaObject, map[string](*Object) {})
var FalseObject = NewObject(TYPE_BOOL, false, BoolMetaObject, map[string](*Object) {})


func BuiltInPlus(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0].Slots["__plus__"].Slots["__call__"].Value.(ObjectCallable)(input, scope, ast)
}
func BuiltInMinus(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0].Slots["__minus__"].Slots["__call__"].Value.(ObjectCallable)(input, scope, ast)
}
func BuiltInAsterisk(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0].Slots["__asterisk__"].Slots["__call__"].Value.(ObjectCallable)(input, scope, ast)
}
func BuiltInSlash(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0].Slots["__slash__"].Slots["__call__"].Value.(ObjectCallable)(input, scope, ast)
}
func BuiltInModulo(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0].Slots["__modulo__"].Slots["__call__"].Value.(ObjectCallable)(input, scope, ast)
}
func BuiltInPower(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0].Slots["__power__"].Slots["__call__"].Value.(ObjectCallable)(input, scope, ast)
}
func BuiltInEqualCompare(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0].Slots["__equal__"].Slots["__call__"].Value.(ObjectCallable)(input, scope, ast)
}

func NewBinaryOperatorObject(name string, callable ObjectCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, nil, map[string](*Object) {
        "__binary__": NewObject(TYPE_OBJECT, callable, nil, nil),
        "__string__": createStringObject("<binary " + name + ">"),
    })
}

func NewBinaryFormObject(name string, callable ObjectFormCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, nil, map[string](*Object){
        "__binary__": NewObject(TYPE_OBJECT, callable, nil, map[string](*Object){ "__form__": TrueObject }),
        "__string__": createStringObject("<object " + name + ">"),
    })
}

func NewCallable(callable ObjectCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, nil, map[string](*Object) {
        "__call__": NewObject(TYPE_OBJECT, callable, nil, nil),
    })
}

func BuiltInDotForm(input [](*AST), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    object, err := evaluateAST(input[0], scope)
    if err != nil { return nil, err }

    return object.GetAttribute(input[1].Value.Value, ast)
}

func BuiltInDefineForm(input [](*AST), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    leftSide := input[0]

    value, err := evaluateAST(input[1], scope)
    if err != nil { return nil, err }

    if leftSide.Value.Type == SIGN && leftSide.Value.Value == "." {
        symbol := leftSide.Right.Value.Value
        object, err := evaluateAST(leftSide.Left, scope)
        if err != nil { return nil, err }

        object.Slots[symbol] = value
    } else if input[0].Left == nil && input[0].Right == nil {
        scope.Symbols[input[0].Value.Value] = value
    } else {
        return nil, NewRuntimeError("lhs must be symbol or object property", leftSide.Value)
    }

	return value, nil
}

func BuiltInPrint(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    for _, obj := range input {
        strObject, err := obj.GetStringRepresentation(scope, ast)
        if err != nil { return nil, err }

        str, err := strObject.GetString(ast)
        if err != nil { return nil, err }

        fmt.Println(str)
    }

	return NilObject, nil
}

func BuiltInStr(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    stringObject, err := input[0].GetStringRepresentation(scope, ast)
    if err != nil { return nil, err }

	return stringObject, nil
}

func BuiltInFunctionScope(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    scopeMap := make(MapObject)

    for name, valueObject := range scope.Symbols {
        keyObject, err := NewStringObject(name)
        if err != nil { return nil, err }

        hash, err := keyObject.GetHash(scope, ast)
        if err != nil { return nil, err }

        scopeMap[hash] = [2](*Object) { keyObject, valueObject }
    }

	return NewMapObject(scopeMap)
}

func BuiltInMeta(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0].GetMetaObject(), nil
}

func BuiltInAssert(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    if input[0] != TrueObject {
        return NilObject, NewRuntimeError("AssertError", ast.Value)
    }
    return NilObject, nil
}

func BuiltInImport(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    importPathListObject, err := scope.SearchSymbol("IMPORT_PATH", ast)
    if err != nil { return nil, err }

    importPathList, err := importPathListObject.GetList(ast)
    if err != nil { return nil, err }

    modulePath, err := input[0].GetString(ast)
    if err != nil { return nil, err }

    for _, importPathObject := range importPathList {
        importPath, err := importPathObject.GetString(ast)
        if err != nil { return nil, err }

        moduleFullPath := path.Join(importPath, modulePath + ".lang")

        if _, err := os.Stat(moduleFullPath); !os.IsNotExist(err) {
            moduleScope := NewScope(scope)
            file, e := os.Open(moduleFullPath)
            if e != nil { return nil, NewRuntimeError("unknown import error", ast.Value) }

            Evaluate(NewReaderWithPosition(file), moduleScope)

            return NewObject(TYPE_OBJECT, nil, nil, moduleScope.Symbols), nil
        }
    }

    return nil, NewRuntimeError("can't import module " + modulePath, ast.Value)
}

func BuiltInIf(input [](*AST), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    if input[1].Value.Value == "else" {
        cond, err := evaluateAST(input[1].Left, scope)
        if err != nil { return nil, err }

        if cond == TrueObject {
            left, err := evaluateAST(input[0], scope)
            if err != nil { return nil, err }

            return left, nil
        } else {
            right, err := evaluateAST(input[1].Right, scope)
            if err != nil { return nil, err }

            return right, nil
        }
    }

    cond, err := evaluateAST(input[0].Right, scope)
    if err != nil { return nil, err }

    if cond == TrueObject {
        left, err := evaluateAST(input[0], scope)
        if err != nil { return nil, err }
        return left, nil
    } else {
        return NilObject, nil
    }
}

func BuiltInSlots(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    scopeMap := make(MapObject)

    for name, valueObject := range input[0].Slots {
        keyObject, err := NewStringObject(name)
        if err != nil { return nil, err }

        hash, err := keyObject.GetHash(scope, ast)
        if err != nil { return nil, err }

        scopeMap[hash] = [2](*Object) { keyObject, valueObject }
    }

	return NewMapObject(scopeMap)
}

func GenerateImportPath() *Object {
    cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    cwdString, _ := NewStringObject(cwd)

    importPath, _ := NewListObject([](*Object){
        cwdString,
    })

    return importPath
}

var BuiltInScope = &Scope{
    Symbols: map[string](*Object){
        "+": NewBinaryOperatorObject("+", BuiltInPlus),
        "-": NewBinaryOperatorObject("-", BuiltInMinus),
        "*": NewBinaryOperatorObject("*", BuiltInAsterisk),
        "/": NewBinaryOperatorObject("/", BuiltInSlash),
        "%": NewBinaryOperatorObject("%", BuiltInModulo),
        "^": NewBinaryOperatorObject("^", BuiltInPower),
        "=": NewBinaryFormObject("=", BuiltInDefineForm),
        ".": NewBinaryFormObject(".", BuiltInDotForm),
        "==": NewBinaryOperatorObject("==", BuiltInEqualCompare),

        "if": NewBinaryFormObject("if", BuiltInIf),
        //"else": NewBinaryFormObject("else", BuiltInElse),

        "object": MetaObject,
        "num": NumberMetaObject,
        "float": FloatMetaObject,
        "string": StringMetaObject,
        "list": ListMetaObject,
        "map": MapMetaObject,
        "tuple": TupleMetaObject,

        "true": TrueObject,
        "false": FalseObject,
        "nil": NilObject,

        "meta": NewCallable(BuiltInMeta),
        "print": NewCallable(BuiltInPrint),
        "scope": NewCallable(BuiltInFunctionScope),
        "str": NewCallable(BuiltInStr),
        "assert": NewCallable(BuiltInAssert),
        "import": NewCallable(BuiltInImport),
        "slots": NewCallable(BuiltInSlots),

        "IMPORT_PATH": GenerateImportPath(),
    },
}
