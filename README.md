# Goat - File watcher

## Abstract

Goat is a file watcher written in Golang. Goat watches files which have specific extensions and executes specific commands when one of these files is created, updated or removed.

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

## Configuration file

To run goat, you have to create a configuration file named `goat.json` or `goat.yml` in your project root directory.

The JSON file looks like the following:

```json
{
  "init_tasks": [
    {
      "command": "make stop"
    },
    {
      "command": "make run",
      "nowait": true
    }
  ],
  "watchers": [
    {
      "extension": "go",
      "tasks": [
        {
          "command": "make stop"
        },
        {
          "command": "make run",
          "nowait": true
        }
      ]
    },
    {
      "extension": "styl",
      "tasks": [
        {
          "command": "make stylus"
        }
      ]
    },
    {
      "extension": "css",
      "excludes": [{"pattern": "all.css", "algorithm": "exact"}, {"pattern": "min.css", "algorithm": "suffix"}],
      "tasks": [
        {
          "command": "make catcss"
        },
        {
          "command": "make uglifycss"
        }
      ]
    },
    {
      "extension": "js",
      "directory": "test",
      "excludes": [{"pattern": "all.js", "algorithm": "exact"}, {"pattern": "min.js", "algorithm": "suffix"}],
      "tasks": [
        {
          "command": "make catjs"
        },
        {
          "command": "make uglifyjs"
        }
      ]
    }
  ]
}
```


The equivalent YAML file looks like the following:
```yaml
init_tasks:
 - command: "make stop"
 - command: "make run"
   nowait: true


watchers:
 - extension: go
   tasks:
   - command: "make stop"
   - command: "make run"
     nowait: true

 - extension: styl
   tasks:
   - command: "make stylus"

 - extension: css
   excludes:
   - pattern: "all.css"
     algorithm: "exact"
   - pattern: "min.css"
     algorithm: "suffix"
  tasks:
   - command: "make catcss"
   - command: "make uglifycss"

 - extension: js
   directory: test
   excludes:
   - pattern: "all.js"
     algorithm: "exact"
   - pattern: "min.js"
     algorithm: "suffix"
  tasks:
   - command: "make catjs"
   - command: "make uglifyjs"
```

* `init_tasks` defines an array of initial tasks. This definition is optional. Each task definition has the following properties:
  * `command` (required)
  * `nowait` (optional)
* `command` defines a command which is executed when one of the target files is created, updated or removed.
* `nowait` defines whether Goat waits the completion of the command or not.
* `watchers` defines an array of file watchers. Each watcher definition has the following properties:
  * `extension` (required)
  * `tasks` (required)
  * `excludes` (optional)
    * `pattern` defile filename or filename pattern which matched with the respective algorithm.
    * `algorithm` (optional, default is `exact`)
      * `exact` exclude the file matched exactly
      * `regexp` exclude the file(s) matched with regexp pattern.
      * `suffix` exclude the file(s) ends with suffix.
      * `prefix` exclude the file(s) starts with prefix.
  * `directory` (optional)
* `extension` defines target file's extension. Goat watches all files which have this extension in and under your project root directory.
* `tasks` defines an array of tasks.
* `excludes` defines an array of file names which is out of watching range.
* `directory` defines the subdirectory. Goat watches all files which have the specified extension in and under this subdirectory under your project root directory. Defaults to your project root directory, if not specified.

## Execution

On the your project root directory which has `goat.json` or `goat.yml` file, execute the following command:

```sh
$ goat
2014/03/06 01:22:04 [Watcher for go files under project root] Watching...
2014/03/06 01:22:04 [Watcher for js files under test] Watching...
2014/03/06 01:22:04 [Watcher for css files under project root] Watching...
2014/03/06 01:22:04 [Watcher for styl files under project root] Watching...
```

Goat launches watcher processes defined on `goat.json` or `goat.yml` file.

Default interval time of each watcher's file check loop is 500 ms. You can change this interval time by specifying -i flag. The following example shows a command which sets the interval time to 1000 ms:

```sh
$ goat -i 1000
```
