errors
===
[![Build Status](https://travis-ci.org/x-mod/errors.svg?branch=master)](https://travis-ci.org/x-mod/errors) [![Go Report Card](https://goreportcard.com/badge/github.com/x-mod/errors)](https://goreportcard.com/report/github.com/x-mod/errors) [![Coverage Status](https://coveralls.io/repos/github/x-mod/errors/badge.svg?branch=master)](https://coveralls.io/github/x-mod/errors?branch=master) [![GoDoc](https://godoc.org/github.com/x-mod/errors?status.svg)](https://godoc.org/github.com/x-mod/errors) 

extension of errors for the following features:

- annotation error
- error with code, support grpc error code

## Quick Start

````go

import "github.com/x-mod/errors"

//annotation error
e1 := errors.Annotate(err, "annotations")
e11 := errors.Annotatef(err, "annotations %s", "format")

//code error
e2 := errors.WithCode(err, errors.ErrNo(100))
//get error's code value
v2 := errors.ValueFrom(e2)

````
