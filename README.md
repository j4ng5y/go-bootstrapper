# Go-Bootstrapper

This little application will set up one or more project directories for you per the standards set at the [Golang Standard Project Layout Repository](https://github.com/golang-standards/project-layout).

## Usage

`go-bootstrapper new-project <project name> [<project2-name> ...]`

## After Bootstrapping

Until I decide to add git initialization to this tool, you will need to navigate to your newly created project and run `git init`.

Also, the `go.mod` file that is generated has a junk module name, so you will want to change that so something relevant:

* e.g. `module github.com/j4ng5y/go-bootstrapper` rather than `module changeme.org/awesomeUsername/<project name>`

## Reason for this tool

I got tired of building the same directory structures time and time again :D