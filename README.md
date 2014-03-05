# Goat - File watcher

## Abstract

Goat is a file watcher made by Golang. Goat watches files which have specific extensions and executes specific commands when one of these files is created, updated or removed.

## Use cases

You can use Goat to:

* Restart a Golang web server process when one of the Golang source files is updated.
* Compile Stylus, Sass/SCSS and LESS source files when one of these files is updated.
* Concatenate and compress JS and CSS source files when one of these files is updated.

## Installation

### From source codes

```sh
$ go get github.com/yosssi/goat/...
```

### By deploying a binary file

* [Linux(64bit)](https://s3-ap-northeast-1.amazonaws.com/yosssi/goat/linux_amd64/goat)
* [Linux(32bit)](https://s3-ap-northeast-1.amazonaws.com/yosssi/goat/linux_386/goat)
* [Mac OS X(64bit)](https://s3-ap-northeast-1.amazonaws.com/yosssi/goat/darwin_amd64/goat)
* [Mac OS X(32bit)](https://s3-ap-northeast-1.amazonaws.com/yosssi/goat/darwin_386/goat)
* [Windows(64bit)](https://s3-ap-northeast-1.amazonaws.com/yosssi/goat/windows_amd64/goat.exe)
* [Windows(32bit)](https://s3-ap-northeast-1.amazonaws.com/yosssi/goat/windows_386/goat.exe)
