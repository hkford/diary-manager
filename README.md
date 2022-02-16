# diary-manager
[![test status](https://github.com/hkford/diary-manager/actions/workflows/unittest.yml/badge.svg)](https://github.com/hkford/diary-manager/actions/workflows/unittest.yml)

Diary manager written in Golang

## Usage

### init
Generate diary templates. They are composed of files such as `2020/202001.txt`. Each file has the following format.

```text
2020,January,01,Wed

2020,January,02,Thur

```

```shell
# Generate template of 2020
$ mydiary init --year 2020
```

### show
Show diary of specified date.

```shell
# Show diary of 2020/01/01
$ mydiary show --date 20200101
```

## Developing

### Build
This repo contains `Dockerfile` so it is easy to build the diary-manager using `Docker`, especially the VSCode Remote Containers extension.

### Unit test
Run the following command. This run all tests in current directory and all of its subdirectories.
```shell
$ go test -v ./...
```