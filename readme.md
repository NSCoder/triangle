# Triangle Solver

Small command line utility with a library to solve triangles using law of sines and cosines.

## Requirements

1.- Golang - https://golang.org
2.- Golang mod, `set -x GO111MODULE`

## Useful commands

- `make build`: Builds the library
- `make test`: Runs the unit tests
- `make run`: Runs a pre-defined triangle solve command

## Command Line Usage

```
NAME:
   Triangle Solver - solve triangles

USAGE:
   triangle [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     solve, s  solve a triangle
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Examples

```
./bin/triangle solve -a 1 -b 2 -c 2
Side-Side-Side Triangle
Classification by sides: Isosceles, by angles: Acute
Sides - A: 1, B: 2, C: 2
Angles - Alpha: 0.5053605102841574, Beta: 1.318116071652818, Gamma: 1.318116071652818
Perimeter: 5
Semiperimeter: 2.5
Area: 0.9682458365518543
Inradius: 2.581988897471611
Circumradius: 1.0327955589886444
```

```
./bin/triangle solve -alpha 1.0471975512 -b 2 -gamma 0.78
Angle-Side-Angle Triangle
Classification by sides: Scalene, by angles: Acute
Sides - A: 1.7905870806814888, B: 2, C: 1.4540948067158643
Angles - Alpha: 1.0471975512, Beta: 1.314395102389793, Gamma: 0.78
Perimeter: 5.244681887397353
Semiperimeter: 2.6223409436986764
Area: 1.259283042129435
Inradius: 2.082407890814065
Circumradius: 1.0337959330368929
```

```
./bin/triangle solve -alpha 1.0471975512 -beta 2 -a 0.78
Angle-Angle-Side Triangle
Classification by sides: Scalene, by angles: Acute
Sides - A: 0.78, B: 0.8189736580743276, C: 0.08489229652797967
Angles - Alpha: 1.0471975512, Beta: 2, Gamma: 0.0943951023897931
Perimeter: 1.6838659546023074
Semiperimeter: 0.8419329773011537
Area: 0.030105015248183757
Inradius: 27.966535487868512
Circumradius: 0.45033320996702353
```
