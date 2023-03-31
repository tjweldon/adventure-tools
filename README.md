## Usage

A tool for handling PNG based sprite sheets.
Built specifically to manage assets for [Wombatlord](https://github.com/Wombatlord)/[adventure-league](https://github.com/adventure-league). 
Go take a look, it's pretty neat.

### Top Level Command

```
Usage: adventure-tools <command> [<args>]

Options:
  --help, -h             display this help and exit

Commands:
  scale                  use this command to scale png sprites into a resulting png
```

### Scale Command

```
Usage: adventure-tools scale [--source SOURCE] [--destination DESTINATION] [--x X] [--y Y]

Options:
  --source SOURCE, -s SOURCE
                         The input file path
  --destination DESTINATION, -d DESTINATION
                         The output file path. To write to stdout, do not provide this option
  --x X, -x X            Set the x scaling factor. [default: 1]
  --y Y, -y Y            Set the y scaling factor. [default: 1]
  --help, -h             display this help and exit
```
