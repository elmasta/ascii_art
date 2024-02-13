# Ascii-art-color
Welcome to the one and only **Ascii-art-color!**

## What is it?
This programme take an option and a sequence of character as an argument and print its ascii art form with some colors

## How does it work?

# Needed files
"main.go", "standard.txt", "shadow.txt" and "thinkertoy.txt" are needed in the "main" folder.

# Launching the program
Usage: go run . [OPTION] [STRING]

EX: go run . --color=yellow "ing" "something"

To launch the program, you need to launch main.go with at least one argument as shawn in this example:
go run main.go thisisatest

Color choices are: red, green, yellow, blue, purple, cyan, orange, white (even though default color is white).

# Extra features
This ascii art is universal which means you can also use the option --output, --align and also use a banner.
You can only use one option at a time (it was not tested with more than one option) but you can use a banner with any option.
