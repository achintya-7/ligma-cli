# ligma-cli 
A cli made in Golang which can show you random jokes in your cli.

## Use this command to install
You need to have Go installed on your hardware
```
go install github.com/achintya-7/ligma-cli@latest
```

## To get a random joke
```
ligma-cli random
```

## To get a random joke with a term
```
ligma-cli random --term sea
```

## Packages and Tech
* golang
* [cobra](https://github.com/spf13/cobra)
* [cobra-cli](https://github.com/spf13/cobra-cli)


## Contribute
If you are using Cobra and Cobra-Cli.
You can add a new command using 
```
cobra-cli add [command--name]
```
This will make a new go file in your cmd and you can start writing your functions there.


