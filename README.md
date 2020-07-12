# :seedling: Lawn Mowers
A simulation of lawn mowers

## How to run the project
```
go get -u github.com/bariabbassi/lawn-mowers
cd $GOPATH/src/github.com/bariabbassi/lawn-mowers
go install
lawn-mowers
```

## What the project does
- Read a test file
```
5 5
1 2 N
LFLFLFLFF
3 3 E
FFRFFRFRRF
```
- Parse lawn dimensions from the first line
- For every 2 lines parse the mower (position and instructions) and add it to the lawn
- Wait till all mowers are parsed and added to the lawn
- Move the mowers on the lawn by executing their instructions without anyone running into another
- Wait till all mowers' instructions have been executed
- Display the positions of all mowers (one-indexed)
```
5 1 E
2 3 N
```

## Testing
All functions have been unit tested. To run all tests, use the command.
```
go test
```
