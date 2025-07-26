#  Monkey Interpreter in Go

This repository contains a complete implementation of the "Monkey" programming language, written in Go, based on Thorsten Ball's book, "**Writing An Interpreter In Go**".

The goal of this project is to explore and understand the inner workings of an interpreter by building every component from scratch, without using any third-party libraries or tools

## The Monkey Language

Monkey is an educational language with a C-like syntax, designed to be simple yet powerful enough to include advanced features. Its key features include:
- C-like syntax
- Variable bindings with  `let`
- Data types: Integers, Booleans
- Arithmetic and logical expressions
- First-class and higher-order functions
- Closures
- A built-in function system

## How It Works: The Interpreter's Architecture

The interpreter is built following a classic "tree-walking" architecture, divided into three main stages:

```
+----------------+      +----------+      +---------+      +-----------+      +----------+
| Code Source    |----->|  Lexer   |----->|  Tokens |----->|  Parser   |----->|    AST   |
+----------------+      +----------+      +---------+      +-----------+      +----------+
                                                                                  |
                                                                                  |
                                                                                  v
                                                                            +----------+
                                                                            | Result   |
                                                                            +----------+
                                                                                  ^
                                                                                  |
                                                                                  |
                                                                            +-----------+
                                                                            | Evaluator |
                                                                            +-----------+
```

### 1. Lexer (Lexical Analysis)
The **Lexer**  is the first component that analyzes the code. It reads the source code character by character and transforms it into a sequence of "tokens." A token is the smallest logical unit of the language, such as a keyword (`let`, `fn`), an identifier (`x`), a number (`123`) or an operator (`+`).

`let five = 5;`  =>  `[LET, IDENT("five"), ASSIGN, INT("5"), SEMICOLON]`

### 2. Parser (Parsing)
The **Parser**  receives the sequence of tokens from the Lexer and checks if their arrangement follows the language's grammatical rules. If the syntax is correct, it builds a tree-like data structure called an **Abstract Syntax Tree (AST)**. The AST represents the hierarchical and logical structure of the program, ignoring details like whitespace or semicolons

To handle complex expressions and operator precedence  (es. `*` before of `+`), the parser uses the  **Pratt Parsing** algorithm

![AST](ast.png)





### 3. Evaluator (Evaluation)
The **Evaluator** is the heart of the interpreter. It "walks" the AST (tree-walking) node by node and gives meaning (semantics) to the program. It uses a recursive function, `Eval`, to perform the actions corresponding to each node::
-   **Computations**: Executes arithmetic and logical operations
-   **Variables**: Saves and retrieves variable values using a structure called an **Environment**, which acts as a "memory" for scopes
-   **Flow Control**: Handles `if/else`  conditions and  `return` statements
-   **Functions**: Creates function objects, handles calls, and, thanks to the Environment, supports closures

The final result of the evaluation is an internal "object" that represents the computed value.

## Repository Structure

The code is organized into packages, each with a specific responsibility:
-   `main.go`: The entry point of the program that starts the REPL.
-   `/ast`: Contains the data structure definitions for the Abstract Syntax Tree nodes.
-   `/lexer`: The tokenizer that transforms source code into tokens.
-   `/parser`: The parser that builds the AST from tokens.
-   `/evaluator`: The evaluator that executes the code by walking the AST.
-   `/object`: Defines the internal object system to represent values (integers, booleans, functions, etc.) during evaluation
-   `/token`: Defines the token types used by the Lexer and Parser.
-   `/repl`: Implements the Read-Eval-Print Loop, the interactive command-line interface.

## How to Run It

To start the interpreter in interactive mode (REPL), you just need to have Go installed and run:
```sh
go run main.go
```
A `>>` prompt will appear where you can write Monkey code.

## Monkey Code Examples

Here are some examples of what the Monkey language can do:

#### Variables and Operations
```monkey
>> let x = 10;
>> let y = x * 2;
>> y / 4;
5
```

#### Functions and Closures
```monkey
>> let newAdder = fn(x) { fn(y) { x + y }; };
>> let addTwo = newAdder(2);
>> addTwo(3);
5
```

#### Higher-Order Functions
```monkey
>> let twice = fn(f, x) { f(f(x)); };
>> let addThree = fn(x) { x + 3; };
>> twice(addThree, 2);
8
```



