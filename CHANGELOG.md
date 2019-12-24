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

## [Planned]

### [1.0.1] - 2019-12-16

#### Added

- Versioning of knowledge base