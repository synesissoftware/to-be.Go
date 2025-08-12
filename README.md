# to-be.Go <!-- omit in toc -->

[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/to-be.Go.svg)](https://pkg.go.dev/github.com/synesissoftware/to-be.Go)

Simple Go library determining the truthyness of strings, that is whether they indicate *truey* or *falsy* values.


## Introduction

**to-be** is a library providing facilities for determine whether the truthyness of strings. It implemented in several languages: **to-be.Go** is the **Go** implementation.


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Terminology](#terminology)
- [Installation](#installation)
- [Components](#components)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development/Testing Dependencies](#developmenttesting-dependencies)
	- [Related projects](#related-projects)
	- [License](#license)


## Terminology

The term "*truthy*" is an unhelpfully overloaded term in the programming world, insofar as it is used to refer to the notion of "truthyness" - whether something can be _deemed to be_ interpretable as truth - and also the true side of that interpretation. In this library, the former interpretation is used, leaving us with the following terms:

* "*truthy*" - whether something can be can be _deemed to be_ interpretable as having truth;
* "*falsey*" - whether an object can be _deemed to be_ interpretable as being false;
* "*truey*" - whether an object can be _deemed to be_ interpretable as being true;

For example, consider the following **Go** program:

```Go
package main

import (
	to_be "github.com/synesissoftware/to-be.Go"
)

s1 := "no"
s2 := "True"
s3 := "orange"

// "no" is validly truthy, and is falsey
to_be.StringIsFalsey(s1)  // true
to_be.StringIsTruey(s1)   // false
to_be.StringIsTruthy(s1)  // true

// "True" is validly truthy, and is truey
to_be.StringIsFalsey(s2)  // false
to_be.StringIsTruey(s2)   // true
to_be.StringIsTruthy(s2)  // true

// "orange" is not validly truthy, and is neither falsey nor truey
to_be.StringIsFalsey(s3)  // false
to_be.StringIsTruey(s3)   // false
to_be.StringIsTruthy(s3)  // false
```


## Installation

```Go

import to-be "github.com/synesissoftware/to-be.Go"
```

## Components

TBC

## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).

## Project Information

### Where to get help

[GitHub Page](https://github.com/synesissoftware/to-be.Go "GitHub Page")

### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/to-be.Go.


### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go/);


#### Development/Testing Dependencies

* [**STEGoL**](https://github.com/synesissoftware/STEGoL/);
* [**testify**](https://github.com/stretchr/testify);


### Related projects

* [**to-be**](https://github.com/synesissoftware/to-be) (**C**);
* [**py2be**](https://github.com/synesissoftware/py2be) (**Python**);
* [**to-be.Ruby**](https://github.com/synesissoftware/to-be.Ruby);
* [**to-be.Rust**](https://github.com/synesissoftware/to-be.Rust);


### License

**to-be.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.
