# Ascii-art-output
Welcome to the one and only **Ascii-art-output!**

## What is it?
This programme take an option and a sequence of character as an argument and record its ascii art form in a file.

## How does it work?

# Needed files
"main.go", "standard.txt", "shadow.txt" and "thinkertoy.txt" are needed in the "main" folder.

# Launching the proramme
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard

To launch the programme, you need to launch main.go with at least one argument as shawn in this example:
go run main.go thisisatest

# Extra features
This ascii art is universal which means you can also use the option --color, --align and also use a banner.
You can only use one option at a time (it was not tested with more than one option) but you can use a banner with any option.