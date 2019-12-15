HTML/CSS Lookout Analyzer
=========================

This is a [lookout](https://github.com/src-d/lookout) analyzer that checks if your PR has validated HTML/CSS when submitting it.

## Usage

Install the analyzer, and run it:

```shell
$ go install github.com/meyskens/lookout-html-analyzer/cmd/lookout-html-analyzer
$ lookout-html-analyzer
```

The analyzer will start listening for pull review requests from lookout.


### test it

To test it, `cd` to the desired repository, and proceed as it follows:

Get the latest `lookout-sdk` from [lookout/releases](https://github.com/src-d/lookout/releases) or just run:

```shell
$ wget -O - https://raw.githubusercontent.com/src-d/lookout-sdk/master/_tools/install-lookout-latest.sh | bash
```

And then run:

```shell
$ ./lookout-sdk review
```

It will mock a Pull Request containing the changes made by `HEAD` over `HEAD~1`, and it will send it to `lookout-html-analyzer` that you ran in the [previous step](#usage).
