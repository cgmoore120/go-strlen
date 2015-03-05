# GO-STRLEN

## Purpose:

Docker image find the length of a string using Go, unicode compliant.

## Usage

### clone the repo

`git clone git@github.com:cgmoore120/go-strlen.git`

### Build the container

`docker build -t cgmoore120/go-strlen ./`

### Run it

`docker run cgmoore120/go-strlen ./go-strlen "this is X characters long" `
