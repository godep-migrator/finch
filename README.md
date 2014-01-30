# Finch

[![GoDoc](https://godoc.org/github.com/BrianHicks/finch?status.png)](https://godoc.org/github.com/BrianHicks/finch)
[![Build Status](https://travis-ci.org/BrianHicks/finch.png?branch=master)](https://travis-ci.org/BrianHicks/finch)

A task manager, written in Go. Finch implements the [Final Version][fv]
algorithm by [Mark Forster][mf].

## Installation

Grab the binaries from GitHub. Or to build specifically for your platform:

```sh
go get github.com/BrianHicks/finch
cd $GOPATH/src/github.com/BrianHicks/finch
make finch
```

This will change in the future to be `go install`able, but it's gotta ship for
now! If it's a big deal for you, leave me a +1 on
[#1](https://github.com/BrianHicks/finch/issues/1).

## Usage

To sum up how this works: add tasks to a list, mark some as "selected" (always
starting with the oldest), then work through your selected list
newest-to-oldest. In Finch, the sequence would look like this:

```
$ finch add email Jane about the meeting
Added "email Jane about the meeting"

$ finch add fix the holodeck
Added "fix the holodeck"

$ finch add eat lunch
Added "eat lunch"

$ finch select
0: email Jane about the meeting
1: fix the holodeck
2: eat lunch

$ finch select 0 2
Selecting "0"... selected "email Jane about the meeting"
Selecting "2"... selected "eat lunch"
Wrote 2 tasks to DB

$ finch next
eat lunch

# after lunch...

$ finch done
Marked "eat lunch" done

$ finch next
email Jane about the meeting

# but you decide that "email Jane about the meeting" can't be done now, or you
# do a little bit of work on it and are interrupted. FV says you should
# re-enter the task at the end of the list. So we have "delay" as well.

$ finch delay
Delayed "email Jane about the meeting"

$ finch select
0: fix the holodeck
1: email Jane about the meeting
```

To get descriptions of the commands, run `finch help` or `finch help [command]`.

## Contributing/Development/Hacking

To be painfully honest I haven't thought through the licensing of this tool
yet. It will be some form of open source, probably MIT or BSD. If you want to
contribute now (as if anyone would) please understand that I may not accept
pull requests until I figure this out and you agree with how your work is being
licensed.

That said, `go get github.com/BrianHicks/finch` (or `git clone`, of course)
should do the trick. You can run `make deps` to get everything. Before you
commit, please run `make test`, `make lint`, and `gofmt`. It'll make everyone's
lives easier.

If you have any thoughts about licensing or build tools please tell me! My
email address is in my Github profile if you don't want to just open an issue.
This is my first "real" go project and I'm open to changing things to fit with
what the community usually expects.

[mf]: http://markforster.squarespace.com/ "Get Everything Done - Mark Forster"
[fv]: http://archive.constantcontact.com/fs004/1100358239599/archive/1109980854493.html "Final Version"
