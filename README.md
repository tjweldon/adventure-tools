# Adventure Tools

A tool for handling PNG based sprite sheets.
Built specifically to manage assets for [Wombatlord](https://github.com/Wombatlord)/[adventure-league](https://github.com/Wombatlord/adventure-league). 
Go take a look, it's pretty neat.

## Top Level Command

```
Usage: adventure-tools <command> [<args>]

Options:
  --help, -h             display this help and exit

Commands:
  scale                  use this command to scale png sprites into a resulting png
  hue-variants           creates a spritesheet with hue variants of the source subsheet
```

## Scale Command

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

## Hue Variants Command

```
Usage: adventure-tools hue-variants [--source SOURCE] [--destination DESTINATION] [--width WIDTH] [--height HEIGHT] [--begin BEGIN] [--end END] [--variants VARIANTS]

Options:
  --source SOURCE, -s SOURCE
                         The source png containing the original sprites
  --destination DESTINATION, -d DESTINATION
                         The destination path to write the resulting png to. Leave blank to write to stdout
  --width WIDTH, -w WIDTH
                         The width of a column in pixels [default: 16]
  --height HEIGHT, -h HEIGHT
                         The height of a row in pixels [default: 16]
  --begin BEGIN, -b BEGIN
                         The index of the sprite to begin at. The begin index is inclusive. [default: 0]
  --end END, -e END      The index of the sprite to end at. The end index is exclusive and a value of 0 denotes the whole span [default: 0]
  --variants VARIANTS, -v VARIANTS
                         The number of variants to produce in the output file [default: 3]
  --help, -h             display this help and exit
```

### Example

This example creates an output spritesheet with 4 colour variants based on the first 8 sprites in the source sheet:

```bash
adventure-tools hue-variants \
	--source assets/sprites/IsometricTRPGAssetPack_OutlinedEntities.png \
	--destination assets/sprites/dude_one_variants.png \
	--width 16 \
	--height 16 \
	--begin 0 \
	--end 8 \
	--variants 4
```

The resulting sheet will have the same sprite dimensions, but will contain 8 sprites per row (corresponding to the eight selected sprites)
and 4 rows, each containing one hue-rotated variant, with the first row always having a hue rotation of 0. The subsequent variants have hue rotations that are evenly distributed around the colour wheel.
