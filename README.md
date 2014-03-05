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

## Configuration file

To run goat, you have to create a configuration file named `goat.json` on in your project root directory. This file looks like the following:

```json
{
  "watchers": [
    {
      "extension": "go",
      "tasks": [
        {
          "command": "make rerun",
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
      "excludes": ["all.css", "all.min.css"],
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
      "excludes": ["all.js", "all.min.js"],
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

* `watchers` defines an array of file watchers. Each watcher definition has the following properties:
  * `extension` (required)
  * `tasks` (required)
  * `excludes` (optional)
* `extension` defines target file's extension. Goat watches all files which have this extension in and under your project root directory.
* `tasks` defines an array of tasks. Each task definition has the following properties:
  * `command` (required)
  * `nowait` (optional)
* `excludes` defines an array of file names which is out of watching range.
* `command` defines a command which is executed when one of the target files is created, updated or removed.
* `nowait` defines whether Goat waits the completion of the command or not.

## Execution

On the your project root directory which has `goat.json` file, execute the following command:

```sh
$ goat
2014/03/06 01:22:04 [go wathcer] Watching...
2014/03/06 01:22:04 [js wathcer] Watching...
2014/03/06 01:22:04 [css wathcer] Watching...
2014/03/06 01:22:04 [styl wathcer] Watching...
```

Goat launches watcher processes.
