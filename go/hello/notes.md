# How to Write Go Code

## Notes

  - `$GOPATH`: resolves import statements
  - `$GOBIN`: Where to place executables
  - In the `go.mod` file:
    * `module <module-name>`: The path id of the module (used for imports)
    * `go <version>`: the version number
  - Importing a local module?
    * `import <module>/<path-to-package>/<package-name>` (unless the path and package name are the same, in that case, just use the path)

### Commands
  - `go mod init <module-path>`
    * Ex. `go mod init example/user/hello`: creates a module (and `go.mod` file) with that path
  - `go install <path>` compiles the program and saves it to the system
    * Puts the exe at `$GOBIN` or `$GOPATH/bin`
    * If `$` are not set, then `$HOME/go/bin/<module>`
    * Can also use `go install`
  - `go build`: compiles the code to the local cache/dir
  - `go env <var>`: Used to read/update/unset environment variables
    * Ex. `go env -w GOBIN=/somewhere/else/bin`: Changes the installation directory
  - 

