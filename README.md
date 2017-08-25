# Oiko
[![GitHub release](https://img.shields.io/github/release/qubyte/rubidium.svg)](https://github.com/MatteoJoliveau/oiko/releases/latest)
## A simple and powerful build tool for Golang

Oiko (from the Greek word *oikodómos*, "the builder") is a simple
build tool for [Golang](golang.org) that leverage the standard
Go tools to compile and manage projects.

## Rationale
Why a build tool that simply wraps the standard `go` command?
Because I was tired of the constraints Go tools impose on developers,
like having to have all the project code inside GOPATH which makes
moving projects around a bit of a pain.
I just want to `git clone` a project, run `build` and voilà, without
having to worry about paths, dependencies or anything else.

### Why not something else?
The only interesting (and maintained) choice other of the standard
goo tools is [GB](getgb.io), but its inability to automatically download dependencies,
requiring the developer to manually copy them or use git submodules
didn't quiet please me. Since Oiko uses the standard `go` tools,
it leverages `go get` to install all the dependencies in the standard
GOPATH. In addition, running a build automatically adds the source code to
GOPATH so that it can be built and linked to pre-existent projects or dependencies.

## Usage

#### Help:
`oiko` or `oiko help`

#### Initialize a new project:
`oiko init`

#### Build project (compile executables):
`oiko build`

#### Install packages and executables inside standard GOPATH:
`oiko install` ***NOT YET IMPLEMENTED***

#### Download missing dependencies:
`oiko update` or `oiko update` ***NOT YET IMPLEMENTED***

#### Clean build output:
`oiko clean` ***NOT YET IMPLEMENTED***

#### See `oiko help` for more information


## Directory Structure
Oiko uses a strict but simple dir structure.
Here is an example using Oiko's own source code:
```
.
├── build
│   └── oiko.exe
├── LICENSE
├── Oikofile
├── README.md
└── src
    └── github.com
        └── matteojoliveau
            └── oiko
                ├── cmd
                │   ├── build.go
                │   ├── init.go
                │   └── root.go
                ├── core
                │   ├── builder.go
                │   ├── config
                │   │   └── config.go
                │   ├── licenses
                │   │   └── licenses.go
                │   └── structures
                │       └── oikofile.go
                ├── oiko.go
                └── util
                    └── fileUtil.go


```

### Important directories and files:
 - **Oikofile** : main project definition file. It's a Yaml file containing
 informations about the projects. Some of them are vitals, such as the Project Name, the Owner or the Namespace
 - **build/** : output folder where executables are stored after running `oiko build`
 - **src/** : main source code repository. Follow the structure domain/user/project
 (eg: `github.com/matteojoliveau/oiko`) for namespaces

## Oikofiles
Oikofiles are the build definition files, similar to npm's `package.json`
or Maven's `pom.xml`.
Here's Oiko's own file:
```yaml
project_name: Oiko
namespace: github.com/matteojoliveau/oiko
version: 0.1.0
owner: Matteo Joliveau
email: matteojoliveau@gmail.com
license: apache2
executable_name: oiko
vcs:
  name: git
  url: https://github.com/MatteoJoliveau/oiko.git
dependencies:
    - github.com/one/depencency
    - github.com/two/depencency
    - github.com/three/depencency
```

The fields used by Oiko to describe and build your projects are:
 - project_name
 - namespace
 - executable_name
 - dependendencies

## Features
### Available
 - **Project generation**: `oiko init`
 - **Build executables**: `oiko build`
### Missing
 - **Project installation**: `oiko install`
 - **Dependency management**: `oiko update`
 - **Cleaning**: `oiko clean`

## How to build

- Git clone this repository in a folder of your choice (better if outside
the standard GOPATH, to avoid problems).
- Run `oiko build`
- Copy `build/oiko` (or `build\oiko.exe` if on Windows) to your PATH,
so that the `oiko` command is accessible from anywhere

In future releases you will be able to simply run `oiko install` to have
it available in PATH.
