# Rif Interpreter

This project is an interpreter for a custom programming language, Rif, built in Go. The interpreter was designed to parse, evaluate, and interpret code written in Rif using a tokenized, abstract syntax tree (AST) approach and Pratt Parsing to parse expressions.

## Overview

Rif is a dynamically typed language created to demonstrate core concepts in language design, parsing, and interpretation. The interpreter is designed with a modular structure to cover the following stages:

1. **Lexing** - Tokenizes the input source code into meaningful units.
2. **Parsing** - Uses Pratt Parsing to transform tokens into an AST, capturing the program structure.
3. **Evaluation** - Walks through the AST to produce results, implementing control flow, functions, and expressions.

## Features

The Rif interpreter supports various features often found in modern programming languages.

### 1. Lexing

The lexer tokenizes the source code and recognizes:
- Keywords and identifiers
- Numeric literals
- Arithmetic and logical operators
- Delimiters and structural tokens

### 2. Parsing (Pratt Parsing)

A Pratt parser is used for expression parsing, allowing for handling complex operator precedence and associativity. Parsing is supported for:
- **Variable Declarations**: Using `let` and `return` statements.
- **Expressions**: Supports prefix, infix, and grouped expressions, including arithmetic and boolean expressions.
- **Conditional Statements**: Parses `if` and `if-else` constructs.
- **Function Literals and Calls**: Supports defining functions, calling functions, closures, and higher-order functions.

### 3. Evaluation

The interpreter evaluates the parsed AST, allowing for real-time interpretation of Rif code. Supported evaluations include:

#### Basic Types

- **Integers**: Supports integer arithmetic.
- **Booleans**: Allows evaluation of logical expressions and conditionals.
- **Null**: Represents the absence of value.

#### Prefix and Infix Operators

- **Prefix Operators**: Supports `-` and `!`.
- **Infix Operators**: Supports standard arithmetic (`+`, `-`, `*`, `/`) and logical operators (`<`, `>`, `==`, `!=`).

#### Conditional Statements

- **If and If-Else Statements**: Evaluates conditional logic with or without an else branch.

#### Function Evaluation

- **Function Literals and Calls**: Supports anonymous functions, closures, and higher-order functions.
- **Return Statements**: Allows functions to return values.

### 4. Extended Data Types

- **Strings**: Allows string literals and concatenation.
- **Arrays**: Supports array literals, indexing, and built-in functions like `len`, `first`, `last`, and `rest`.
- **Hashes**: Enables hash literal creation, indexing, and evaluation of hash expressions.

### 5. Built-In Functions

Rif includes several built-in functions:
- **len**: Returns the length of arrays or strings.
- **first**: Retrieves the first element of an array.
- **last**: Retrieves the last element of an array.
- **rest**: Returns all elements of an array except the first.
- **puts**: Prints values to the console.

## Testing

Unit tests have been implemented for each core component of the interpreter. The tests cover:
- Token generation and lexical analysis
- Parsing expressions and statements
- AST structure validation
- Evaluation of expressions and functions
- Error handling during parsing and evaluation

Run the tests with:

```bash
go test ./...
```

## Running the Interpreter

To launch the REPL (Read-Eval-Print Loop) and start interpreting Rif code, use:

```bash
go run main.go
```

## Example Usage

Within the REPL, you can experiment with Rif code:

```rif
let x = 5;
let s = "hello";

let add = fn(a, b) { a + b };
add(x, 10);

if (x > 2) { puts("x is greater than 2"); }

let arr = [1, 2, "hello"];
puts(arr[0]);

let hashmap = { "hello": "there", 1: "one"};
puts(hashmap[s]);
puts(hashmap[2-1]);
```

## Project Structure

- **Lexer**: Tokenizes the input.
- **Parser**: Converts tokens into an AST using Pratt Parsing.
- **AST**: Represents the structure of Rif code.
- **Evaluator**: Interprets the AST, performing operations and resolving values.
- **REPL**: Runs the interpreter interactively in a loop.
