# CI Tool CLI App
A simple CLI App that can be used to perform Continuous Integration (CI) into a go project

## How to Install
### Clone the repository

    git clone https://github.com/dummyheaad/go-cli-goci

### Build the executable

    go build .

## Functionalities
Perform several CI tasks including formatting, linting, cycle complexity analysis, testing, building, and ended with pushing to github repository

    ./goci.exe -p PROJECT -b BRANCH
    OUTPUT:
    Gofmt: SUCCESS
    Golint: SUCCESS
    Gocyclo: SUCCESS
    Go Test: SUCCESS
    Go Build: SUCCESS
    Git Push: SUCCESS


Note:
- Make sure that there's already a commit on local repo before using the goci tool
- Make sure that the BRANCH should exist prior to using goci tool (create a branch using "git checkout -b BRANCH")
- If BRANCH is not specified then the master branch will be selected