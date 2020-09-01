# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Released]

### [1.0.0] - 2019-12-13

#### Added

- Grool is donated to hyperjump.tech with new name `grule-rule-engine`
- GRL Lexer, Parser using ANTLR
- Rule Builder
- Knowledge Base container
- Reflection Tooling to work with fact objects
- DataContext to store facts
- Cycle counter to ensure rule voting are not looped infinitely
- Grule Rule Engine
- Basic Built-In Functions
- RuleEngine are working on facts that based on pointer so it can modify facts struct instances. Thus adding variable into DataContext will be checked to ensure `ptr` to a `struct`.
- Added `Retract` function so rule can temporary retract a rule from knowledge base so it will not get voted any more in the next cycle.
- Method invocation and tracer bug
- Reflectools are now able to detect the object under reflection for its `value` vs `ptr to struct`
- Function invocation now check if the argument is an Interface, it should accept any type of argument type values.
 
### [1.0.1] - 2019-12-16

#### Added

- Added Pub-Sub mechanism for rule execution events.

#### Fix

- Better DRL load and compilation
- Better DRL error handling when compiling DRL

### [1.1.0] - 2019-12-27

#### Added

- Initial RETE algorithm were added into Grule with only optimization in the ExpressionAtom level.
- Naming and Versioning of knowledge base


### [1.2.0] - 2020-01-16

#### Added

- Newly revamped ANTLR4 Grammar for Grule, syntax and structure not changed but parsing get more efficient.
- Support for modulus % operator
- Support for bitwise OR and AND operator
- Operator precedence support
- RETE optimization to ensure reset of ExpressionAtom only happen if a known variable were changed

### [1.2.3] - 2020-02-14

#### Added

- Resource bundling, to load multiple DRL files by file path pattern
- Load DRL resources from GIT
- Resource bundling, to load multiple DRL files from GIT by the file path patteern

### [1.2.4] - 2020-02-24

#### Added

- EventBus implementation for Grule's internal event messaging now replaces the previous simple subscriber approach.
- Added documentation regarding this EventBus implementation

### [1.3.0] - 2020-06-11

#### Added

- Variadic function calling

### [1.4.0] - 2020-06-14

#### Added

- Support for `escape` character in string literal
- `RuleBuilder` is now to build rules in GRLs into `KnowledgeLibrary`
- Now you should obtain a `KnowledgeBase` instance from `KnowledgeLibrary`. This enable concurrency model in Grule. See `examples/Concurrency_test.go` to know how it works. 

### [1.5.0] - 2020-08-02

#### Added

- Support to build rule from JSON.
- Engine support for `context.Context` using `ExecuteWithContext` function.

### [1.6.0] - 2020-09-01
 
#### Added

- Enhancing in variable traversal, from previously using string tracing to struct-field lookup in reflect.Value
- Support for Array/Slice and Map handling.
- Support for Function chaining.
- Support for Constant functions.
- Grule engine optimization for selecting from conflict set. Instead of sorting salience in descending, simply look for the biggest value.

#### Removed

- Grule Event Bus is removed from Grule as it seems too complicated and no one use them. They just expect grule to just works. 