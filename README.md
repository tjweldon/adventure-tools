## Usage


### Top Level Command

```
Usage: main <command> [<args>]

Options:
  --help, -h             display this help and exit

Commands:
  scale                  use this command to scale png sprites into a resulting png
```

### Scale Command

```
Usage: main scale [--source SOURCE] [--destination DESTINATION] [--x X] [--y Y]

Options:
  --source SOURCE, -s SOURCE
                         The input file path
  --destination DESTINATION, -d DESTINATION
                         The output file path. To write to stdout, do not provide this option
  --x X, -x X            Set the x scaling factor. [default: 1]
  --y Y, -y Y            Set the x scaling factor. [default: 1]
  --help, -h             display this help and exit
```
