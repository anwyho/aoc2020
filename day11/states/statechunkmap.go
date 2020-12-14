package states

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

const stateChunkMapFileName = "./day11/state_chunk_map.csv"
const pythonSrcToGenerateLookUpMap = `generate_states = lambda iter: [''] if iter == 0 else [state for list_of_states in [[c + n for n in generate_states(iter - 1)] for c in ('.', 'L', '#')] for state in list_of_states]; open('/Users/alo/go/src/aoc/state_chunk_map_2.csv', 'w').write('\n'.join(["{state},{derived_state}".format(state=state, derived_state=''.join([(lambda state, pos: '.' if state[pos] == '.' else ('L' if 0 < len([1 for adj_pos in [pos-5, pos-4, pos-3, pos-1, pos+1, pos+3, pos+4, pos+5] if state[adj_pos] == '#']) else '#') if state[pos] == 'L' else ('L' if 4 <= len([1 for adj_pos in [pos-5, pos-4, pos-3, pos-1, pos+1, pos+3, pos+4, pos+5] if state[adj_pos] == '#']) else '#'))(state, pos) for pos in [5, 6, 9, 10]])) for state in generate_states(iter=16)]))`
// Each state chunk is 16 states, with 3 kinds of states
var numPossibleStateChunks int = int(math.Pow(3, 16))                  // 43,046,721
var derivedStateChunkLookUp map[serializedStateChunk]derivedStateChunk // = GetDerivedStateChunkLookUpMap()

func derivedStateChunkFromSerializedStateChunk(ssc serializedStateChunk) derivedStateChunk {
	if derivedStateChunkLookUp == nil { loadDerivedStateChunkLookUpMap() }
	if dsc, found := derivedStateChunkLookUp[ssc]; found {
		return dsc
	} else {
		panic(fmt.Sprintf("couldn't find state chunk in look up map: %s", ssc))
	}
}

// Returns the look up table for a serialized state chunk to a derived state chunk.
//
// There are (3^16 or 43,046,721 variations) to a derived state chunk (3^4
// or 81 variations). In total, there should be 3^16 entries, with each
// entry consisting of a serialized state chunk (uint16) mapped to a
func loadDerivedStateChunkLookUpMap() {
	start := time.Now()
	derivedStateChunkLookUp = make(map[serializedStateChunk]derivedStateChunk, numPossibleStateChunks)
	input, err := ioutil.ReadFile(stateChunkMapFileName)
	fmt.Printf("loading state chunk look up map into memory from: %s\n", stateChunkMapFileName)
	if err != nil {
		panic(fmt.Sprintf("couldn't read file: %s", stateChunkMapFileName))
	}
	fmt.Printf("  expecting %d entries...\n", numPossibleStateChunks)
	for entryIx, entry := range strings.Split(string(input), "\n") {
		// entries are in the format "serialized_state_chunk,serialized_derived_state_chunk"
		entryList := strings.Split(entry, ",")
		if len(entryList) != 2 {
			if len(entryList) == 1 && entryList[0] == "" {
				continue
			}
			panic(fmt.Sprintf("state chunk map file must be a CSV with two columns, e.g. \"L.#..L#LL##.#..#,#L.L\"; found: %s", entryList))
		}
		stateChunk := serializedStateChunk(entryList[0])
		derivedStateChunk := parseDerivedStateChunk(entryList[1])
		derivedStateChunkLookUp[stateChunk] = derivedStateChunk

		if (entryIx+1)%10_000_000 == 0 {
			fmt.Printf("  loaded %d million entries\n", (entryIx+1)/1_000_000)
		}
	}
	fmt.Printf("  loaded %d entries in %v.\n", len(derivedStateChunkLookUp), time.Since(start))
	if len(derivedStateChunkLookUp) != numPossibleStateChunks {
		panic(fmt.Sprintf("missing or invalid state chunks. expected %d, found: %d",
			numPossibleStateChunks, len(derivedStateChunkLookUp)))
	}
}

func parseDerivedStateChunk(s string) derivedStateChunk {
	return derivedStateChunk{
		{State(s[0]), State(s[1])},
		{State(s[2]), State(s[3])},
	}
}

