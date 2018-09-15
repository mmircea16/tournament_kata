## Disclaimer
This is kata from exercism.io that was extended with some new requirements. The original kata can be found [here](https://exercism.io/my/tracks/go) under the name Tournament. The code is [here](https://github.com/exercism/go/tree/master/exercises/tournament). Kudos to all contributors to the kata.
 
# Tournament

Tally the results of a small football competition.

Based on an input file containing which team played against which and what the
outcome was, create a file with a table like this:

```text
Team                           | MP |  W |  D |  L | GS | GR | P
Devastating Donkeys            |  3 |  2 |  1 |  0 | 10 |  7 | 7
Allegoric Alaskans             |  3 |  2 |  0 |  1 | 12 |  5 | 6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3 |  7 | 3
Courageous Californians        |  3 |  0 |  1 |  2 |  3 |  9 | 1
```

What do those abbreviations mean?

- MP: Matches Played
- W: Matches Won
- D: Matches Drawn (Tied)
- L: Matches Lost
- GS: Goals scored
- GR: Goals received
- P: Points

A win earns a team 3 points. A draw earns 1. A loss earns 0.

The outcome should be ordered by points, descending. In case of a tie, teams are ordered on goal difference (ie. GS - GR). In case of a tie on goal difference, teams are ordered on goals scored. In case of a further tie, the teams are ordered alphabetically 

###

Input

Your tallying program will receive input that looks like:

```text
Allegoric Alaskans;Blithering Badgers;3-0
Devastating Donkeys;Courageous Californians;2-2
Devastating Donkeys;Allegoric Alaskans;4-3
Courageous Californians;Blithering Badgers;0-1
Blithering Badgers;Devastating Donkeys;2-4
Allegoric Alaskans;Courageous Californians;6-1
```

The result of the match refers to the first team listed. So this line

```text
Allegoric Alaskans;Blithering Badgers;3-0
```

Means that the Allegoric Alaskans scored 3 goals beat the Blithering Badgers who have not scored any.

## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem`
flags:

    go test -v --bench . --benchmem

Keep in mind that each reviewer will run benchmarks on a different machine, with
different specs, so the results from these benchmark tests may vary.

## Further information

For more detailed information about the Go track, including how to get help if
you're having trouble, please visit the exercism.io [Go language page](http://exercism.io/languages/go/resources).

