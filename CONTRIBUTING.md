# Contributing to goreturn

First off, thank you for considering contributing to `goreturn`. It's people like you that make `goreturn` such a great tool.

## Getting Started

- Make sure you have a [GitHub account](https://github.com/signup/free)
- Submit a ticket for your issue, assuming one does not already exist.
  - Clearly describe the issue including steps to reproduce when it is a bug.
  - Make sure you fill in the earliest version that you know has the issue.

## Making Changes

- Fork the repository on GitHub.
- Create a topic branch from where you want to base your work.
  - This is usually the master branch.
  - Only target release branches if you are certain your fix must be on that branch.
  - To quickly create a topic branch based on master; `git branch fix/master/my_contribution master` then checkout the new branch with `git checkout fix/master/my_contribution`. Please avoid working directly on the `master` branch.
- Make commits of logical units.
- Check for unnecessary whitespace with `git diff --check` before committing.
- Make sure your commit messages are in the proper format.

## Submitting Changes

- Push your changes to a topic branch in your fork of the repository.
- Submit a pull request to the repository in the `Eyal-Shalev` organization.
- The core team looks at Pull Requests on a regular basis.
- After feedback has been given we expect responses within two weeks. After two weeks we may close the pull request if it isn't showing any activity.

## Code Style

Please follow the existing coding style. This project uses `gofmt`, so you can format your code correctly by running `gofmt -s -w .` in your project directory.

## Running Tests

Please make sure all tests pass before submitting a Pull Request. You can run tests with the `go test` command.

## License

By contributing your code, you agree to license your contribution under the terms of the MIT: https://opensource.org/licenses/MIT

## Thank you

Your contributions to open source, large or small, make great projects like this possible. Thank you for taking the time to contribute.