[![Build Status](https://travis-ci.com/sukovanej/lang.svg?branch=master)](https://travis-ci.com/sukovanej/lang)

# Basics

## Compilation

```bash
$ cd lang-source-code-dir
$ go build -o bin/lang
```

## Repl

```
$ ./bin/lang
```

## Basic Expressions

I'm trying to follow functional and OOP paradigm at the same time. Evaluation is mostly inspired by 
Scheme/Lisp language. From the OOP point-of-view every expression evaluated is an object. And from the
functional side, (almost) every symbol (sign, identifier, number, etc.) is an expression, generaly if
possible I prefer expressions before statements. Evaluation is based on the binary expressions evaluation.

The syntax tree is generated from expressions `E op E` if possible. Of course there is syntax sugar for lists 
(`[1, 2, 3]`), if-else statement (`x if condition else y`), condition statement (`cond { cond1: e1 }`), and so on.

Evaluater takes binary expression of the form `<expression1> <operator> <expression2>` and every of these three
symbols is evaluated to an `*Object`. `<operator>` must be an object with a `__binary__` slot which must be 
callable (it is an object with `__call__` slot defined) then internally 
`<operator>.__binary__.__call__(<expression1>, <expression2>)` is called and the result is returned.

The important thing is that expressions like `+` or `-` are first-class citizen expression.

```
>>> +
<binary +>
>>> -
<binary ->
```

You can try basic expression with builtin binaries (`+`, `/`, `%`, `^`, etc.).

```
>>> 2 + 1
3
>>> 5 * (6 + 7)
65
```

The very basic feature of every programming language is the variable defition. Every scope contains a list
of symbols. One can use `scope` function to get the dictionary (key-value store realized by hashtable) of
symbols mapped to the objects.

```
>>> x = 1
1
>>> scope()
{x: 1}
```

Also it is useful to be able to define a *lambda expression*.

```
>>> (x) -> x + 1
<callable> @ 0xc0000928d0
```

Lambda can be assigned to a symbol. Also special syntax `<function-name>(<args>) -> <result>` can be used.

```
>>> f = (x) -> x + 1
<callable> @ 0xc0000928d0
>>> f(1)
2
>>> g(x) -> x - 1
<callable> @ 0xc000092a80
>>> g(1)
0
```

# TODO

 - [ ] decorators
 - [ ] iterators
 - [ ] unpacking
 ```
 f(x) -> (x, x + 1)
 x, y = f(1)
 ```
