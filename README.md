# Go CI/CD Demo

This is a simple Go project created to demonstrate how to set up basic CI/CD workflows using GitHub Actions.

# About the Project

The project contains minimal Go code, just enough to build and test.The main goal is to show how to automate CI and CD processes with GitHub Actions.

# CI Workflow

On every push to the `main` branch,the CI workflow automatically runs the following steps:

- Checkout the source code
- Set up Go environment
- Run tests
- Build the binary

# CD Workflow

When a new GitHub Release is created, the CD workflow is triggered. It performs the following steps:

- Checkout the source code and set up Go
- Run `golangci-lint` for linting
- Run tests and build the binary
- Upload the binary to the GitHub Release

# GitHub Actions Used

- `actions/checkout`
- `actions/setup-go`
- `golangci/golangci-lint-action`
- `softprops/action-gh-release`
