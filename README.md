# constrict

-----

## about

Constrict is a custom build system I wrote for myself [think of cmake, gmake, that sort of thing]  and isn't really meant for the general public, but it has about the most k.i.s.s syntax as possible.

## build

For me, this is the last makefile I will ever write, so you can run make to build the project.

## syntax

Syntax documentation. Basically the only docs.

### variables

Variables go up top, cannot double apply variables unless you pass the ! into its name. 
```
x = 1
!x = y
y = main.c
obj = main.o
```

### functions

Functions are declared through an all caps word. Like `CLEAN` or `INSTALL`.
There is one reserved function, called `DEFAULT`. It is the default function run when constrict is run.

```
DEFAULT
echo "constrict, is named because it sounds close to construct, and constructing stuff is kind of like compiling."

CLEAN
rm $(obj)
```

## todo

I will try to add incremental building, and a couple other features including printing to stdout every single command the program runs.

#### naming

Constrict, is named because it sounds close to construct, and constructing stuff is kind of like compiling a program. Also it sounded snakey. I would've said it reminds me of golang but I think the Go language logo is like a blue beaver or something. Thanks for reading, this wasn't a big project just something for myself since I had enough of writing makefiles.


