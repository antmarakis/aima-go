# aima-go

## How to Get Started

To download:

```
git clone https://github.com/MrDupin/aima-go.git
cd aima-go
git submodule init
git submodule update
```

To run the test suit execute the following command in the `aima-go` directory:

`go test`

This runs the `*_test.go` files in the directory. The functions called should be in the form `TestXxx(t *testing.T)`.

## Structure of the Project

The decision to abandon the usual Go project structure is a concious one (and up for discussion). I feel that since this project is very simple it is best to keep the structure as minimalistic as possible. Some points in favor of this choice:

* Nothing to build.
* There will not be many `go` files.
* Only one "package", `utils.go`.
* Apart from the `go` files, we only have test files.
* The student doesn't need to set $GOPATH.
* After cloning the repository, the student only needs to run `go test` from the directory.

All that means the student can easily download, edit and test their code without much hassle. The only downside I see is that the structure may help develop bad practises, but I feel that is unlikely since it is obvious this project is for educational purposes only.
