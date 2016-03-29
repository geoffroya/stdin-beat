# Stdin-beat

Welcome to Stdin-beat.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/geoffroya/stdin-beat`

## Getting Started with Stdin-beat

### Init Project
To get running with Stdin-beat, run the following command:

```
make init
```

To commit the first version before you modify it, run:

```
make commit
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Stdin-beat in the git repository, run the following commands:

```
git remote set-url origin https://src/github.com/geoffroya/stdin-beat/stdin-beat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Stdin-beat run the command below. This will generate a binary
in the same directory with the name stdin-beat.

```
make
```


### Run

To run Stdin-beat with debugging output enabled, run:

```
./stdin-beat -c stdin-beat.yml -e -d "*"
```


### Test

To test Stdin-beat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`


### Package

To cross-compile and package Stdin-beat for all supported platforms, run the following commands:

```
cd dev-tools/packer
make deps
make images
make
```

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/stdin-beat.template.json and etc/stdin-beat.asciidoc

```
make update
```


### Cleanup

To clean  Stdin-beat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Stdin-beat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/geoffroya/stdin-beat
cd ${GOPATH}/src/github.com/geoffroya/stdin-beat
git clone https://src/github.com/geoffroya/stdin-beat/stdin-beat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).
