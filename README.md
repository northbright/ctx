# ctx

Update: this repo is no longer maintained. As of Go 1.7 this package is available in the standard library under the name context. <https://golang.org/pkg/context>. [ctxcopy](ctxcopy) and [ctxdownload](ctxdownload) can be replaced with [iocopy](https://github.com/northbright/iocopy).

[![Build Status](https://travis-ci.org/northbright/ctx.svg?branch=master)](https://travis-ci.org/northbright/ctx)
[![Go Report Card](https://goreportcard.com/badge/github.com/northbright/ctx)](https://goreportcard.com/report/github.com/northbright/ctx)

This repository holds [Golang](http://golang.org) packages to provide Google's Context-aware helper functions.

## [Google's Context](https://godoc.org/golang.org/x/net/context)
*  Context type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.
*  References:
  * <https://godoc.org/golang.org/x/net/context>
  * <http://blog.golang.org/context>

## Packages
* [ctxcopy](ctxcopy)  
  ctxcopy is a [Golang](http://golang.org) package which provides helper functions for performing context-aware copy task.

* [ctxdownload](ctxdownload)  
  ctxdownload is a [Golang](http://golang.org) package which provides helper functions for performing context-aware download task.
