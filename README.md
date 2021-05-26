This is one of the Advent of Code 2020 solutions that I'm most proud of. I loved working through this problem, discovering new optimizations, and ultimately, explaining it all in an understandable way. 

# Day 11 Write-Up

The prompt for this challenge can be found at [Advent of Code Day 11](https://adventofcode.com/2020/day/11).

## Challenge Description

The input for this Advent of Code challenge consists of a grid of states, referring to whether or not a space is a `Floor` (or `.`), an `Empty` seat (or `L`), or an `Occupied` seat (or `#`). 

The goal is to look for a stable equilibrium of grid state (i.e. the grid stops changing) and report the number of `Occupied` seats at the equilibrium. The three heuristics for seat state changes between iterations are:

- An `Empty` seat becomes `Occupied` if there are no adjacent `Occupied` seats.
- An `Occupied` seat becomes `Empty` if there are **four or more** `Occupied` adjacent seats.
- `Floor` never changes.

A seat is considered adjacent to another seat if it is located at one of the eight positions immediately up, down, left, right, or diagonal from the other.

The goal is to keep finding the next state of the grid until the grid stops changing. Then once the grid stops changing, we count how many `Occupied` seats there are.

# Approach

## Gotta go fast

The challenge description mentioned that:

> All decisions are based on the number of occupied seats adjacent to a given seat.

so, the brute-force method to deriving the grid's next state would be to iterate through every seat and check if it will change based on the seats adjacent to it, all at run time.

However, there's opportunity for some pre-processing, since a seat's state change is only and directly related to its adjacent seats and not the rest of the grid. Given any 3x3 chunk of states, we can derive the next state of the center seat. 

For example, given the following chunk,

```
#.#
##L
L.#
```

we can derive that the center seat, currently `Occupied` (`#`), will become `Empty` (`L`) in the next state since there are four `Occupied` seats adjacent to it.  

This means that we can iterate through every seat in a grid, take its adjacent seats, and determine what its next state will be. 

## What about edge seats? 

For seats on the edges, we can add missing adjacent seats as `Floor` since `Floor` doesn't affect what the next state of a seat will be. 

For example, the two `Empty` seats in chunk 1 and chunk 2 would be derived to the same states. (Unknown state is denoted as `?`.)

```
  L??       ...
  ???       .L?
  ???       .??
  
chunk 1   chunk 2
```

## Pre-processing time

Since seats can only be in three states, we can note every 3x3 chunk to see what its center seat will be derived to. In total, that would be 3<sup>9</sup>, or 19,683 state chunks to derive. (There are 9 spaces with 3 potential states each.) We can serialize (turn into a string) all these state chunks and put them into a mapping from state chunk to its next center seat. Here's what serializing a state chunk could look like:

```
#.#
##L -> #.###LL.#
L.#
```

We're just listing out every seat from left-to-right then top-to-bottom.

If we made a dictionary of these mappings, the entry would look like the middle one, since `#.###LL.#` would mean the center seat would turn into `L`:

```
...
#.###LL.L -> #

#.###LL.# -> L // Nice, we know really quick that the next center seat will be an "L"

#.###LLL. -> #
...
```

The algorithm for finding the next state of a grid would then require iterating through every seat, finding its adjacent seats, serializing the chunk of seats, and then lookup up the seat in the mapping. Hm... That doesn't seem very much faster than the original brute force approach since we still have to grab all of a seat's adjacent seats. What could we improve? 

## Let me do you one better. 3<sup>16</sup> = 43,046,721

Rather than deriving all the possible 3x3 state chunks, why not do all the possible 4x4 state chunks? The principle is still the same - if we can split the grid into state chunks of 2x2 states and then get the surrounding adjacent seats, we can derive the next state of every 2x2 state chunk just by looking it up in a mapping. 

For example, given this state chunk, 

```
###.
.L.#
.LL#
LL.#
```

we can derive that the next 2x2 center will be:

```
L.
#L
```

We can do this derivation for every possible 4x4 state chunk, and serialize them into a mapping. The entry might look like:

```
###..L.#.LL#LL.# -> L.#L
```

In total, the number of entries we would have would be 3<sup>16</sup>, or 43,046,721 entries; there are 3 possible states for 16 seats. That's a lot! 

## Let me do you one better?? 3<sup>25</sup> = 847,288,609,443...!

Woah, woah, hold up. Computers can have a lot of memory, but not that much memory (as of 2020 at least). In order to load all the 4x4 state chunks (43 million) into memory, it would take approximately 1 GB. However, 3<sup>25</sup> or 847 billion, the number of possible 5x5 state chunks, is nearly 20,000x larger than 43 million, and each entry size would be approximately at least 25 + 9 bytes (5x5 chunk and its 3x3 derived chunk) rather than 16 + 4 bytes which would result in about a 1.7x increase in entry size.

Instead, the next places for optimization would probably compressing the mappings (to reduce the 1.7x increase) and/or getting two terabytes of RAM. We could also eliminate all rotations of a state chunk and check if a given state chunk is in the mapping in any of its rotations. For all possible 5x5 state chunks, that would reduce the number of entries by more than 25% of 847 billion (still about 350 GB). 

In any case, we'll just stick with the 4x4 state chunks without compression for now. 

## Calculate all the things

We can use this line of Python to generate our table in the form of a CSV (comma-separated values): 

```bash
$ python -c 'open("./state_chunk_map.csv", "w").write("\n".join(["{state},{derived_state}".format(state=state, derived_state="".join([(lambda state, pos: "." if state[pos] == "." else ("L" if 0 < len([1 for adj_pos in [pos-5, pos-4, pos-3, pos-1, pos+1, pos+3, pos+4, pos+5] if state[adj_pos] == "#"]) else "#") if state[pos] == "L" else ("L" if 4 <= len([1 for adj_pos in [pos-5, pos-4, pos-3, pos-1, pos+1, pos+3, pos+4, pos+5] if state[adj_pos] == "#"]) else "#"))(state, pos) for pos in [5, 6, 9, 10]])) for state in (lambda f: (lambda x: f(f, x)))(lambda gen_states, iter: [""] if iter == 0 else [state for list_of_states in [[char + next_char for next_char in gen_states(gen_states, iter - 1)] for char in ".L#"] for state in list_of_states])(16)]))'
```

See [this file](state_chunk_map.csv) for the result. (~~Be warned - it might take a lot of data to load~~ GitHub isn't letting me upload this even with git-lfs so... you can generate it if you want.) (Also, a cool Where's Waldo type of thing: [Anonymous recursion is pretty cool](https://en.wikipedia.org/wiki/Anonymous_recursion).)

This will calculate all 43 million 4x4 state chunks, and their derived 2x2 state chunks, keeping in mind all the heuristics of the state changes. Locally, it takes me about 5.5 minutes to run, since it's generating about 1 GB worth of mappings and writing them to disk.

The top of our CSV will end up looking like this: 

```
................,....
...............L,....
...............#,....
..............L.,....
..............LL,....
..............L#,....
..............#.,....
..............#L,....
..............##,....
.............L..,....

...
```

This is denoting the serialized state chunk to its serialized derived state chunks (separated by the comma). We're starting by generating all the state chunks that look like, 

```
.... .... .... .... .... .... .... .... .... ....
.... .... .... .... .... .... .... .... .... ....
.... .... .... .... .... .... .... .... .... ....
.... ...L ...# ..L. ..LL ..L# ..#. ..#L ..## .L.. ...
```

and for these first few entries, they all map to `....` since the center of the state chunks are all `Floor` and `Floor` does not change. 

Later on, we get entries that look more like this, which will probably be used a lot more:

```
...

#.L#L#..L###..LL,#.##
#.L#L#..L###..L#,#.#L
#.L#L#..L###..#.,#.#L
#.L#L#..L###..#L,#.#L                          #.L#
#.L#L#..L###..##,#.#L <- this line represents: L#.. -> #.
#.L#L#..L###.L..,#.##                          L###    #L
#.L#L#..L###.L.L,#.##                          ..##
#.L#L#..L###.L.#,#.#L
#.L#L#..L###.LL.,#.##

...
```

The file ends up being about 1 GB, but now our program will be a matter of doing some mapping lookups rather than having to calculate every seat's adjacent seat states. Much faster! (One in-memory lookup for four seats rather than 9 lookups per seat, or 36 lookups for the equivalent four seats.)

## Steps

Here is the algorithm we'll be following: 

0.  (As a pre-processing step) Parse and load the mapping into a lookup table.
    

1. Iterate through every 2x2 state chunk and serialize it its adjacent neighbors and any additional `Floor` to get a 4x4 state chunk.
2. Look up the serialized state chunk in the look up table and write the result to a new grid.
3. Check to see if the new grid is the same as the old grid. 
    1. If so, repeat from step 1
4. Count how many `Occupied` seats there are. 

And that's that! Oh, the joys of pre-processing. I love the idea of pre-processing because we're doing all the work beforehand. It's like driving a car with a 10 gallon tank as opposed to the same car with a 2 gallon tank. Sure, the fill-up time takes longer, but once we're on our way, there's no stopping us! 

## Next Steps

We can introduce some parallelism into this now. Since each 4x4 state chunk has enough information to calculate its derived 2x2 state chunk, we can start to delegate batches of calculations to threads that can run in parallel to each other. 

For example, we can split out the top half of the grid to one thread and the bottom half to another. Or maybe have each row of state chunks split into its own thread. Then once each of the threads is complete, they circle back to the main thread and update the new grid with the derived 2x2 state chunks. This might not see a performance benefit with a grid of 90 x 92 like we have in the input, but we'll see a massive performance gain in grids with thousands of rows and columns. 

## Downsides to this approach

This approach worked really well because the heuristics of state changes were very localized to the seat that was changing. That is, we could grab a chunk of states and that was all we needed to find the next state. However, part two of the problem (which I won't be doing but instead crying over) includes a heuristic that looks beyond the immediately adjacent seats further into the row or column of the seat.

There are plenty of ways to preprocess for this step as well though, for example, putting an additional state `Unknown` (denoted as `?`) if the state of a center piece can't be immediately determined to mark it for further derivation.

# Result

## Implementation Setup Notes

I ended up parsing the initial layout and then putting it into a grid where it would be padded top and left with a row and column of `Empty` and right and bottom with as many rows and columns of `Empty` as necessary to reach a multiple of four in grid row length and column length. 

Then I followed the algorithm above. Much of the writing above is framed as parsing non-overlapping 2x2 state chunks, but I approached the implementation thinking about 4x4 overlapping chunks. I think it definitely helped with writing the implementation and slicing the 2D array cake. 

## Output
```
loading state chunk look up map into memory from: ./day11/state_chunk_map.csv
  expecting 43046721 entries...
  loaded 10 million entries
  loaded 20 million entries
  loaded 30 million entries
  loaded 40 million entries
  loaded 43046721 entries in 13.423943286s.
loading layout with 90 x 92 into 92 x 96 grid
  done.
iterating through grid states
  found 2126 occupied seats after 116 state changes in 76.049754ms.
```
