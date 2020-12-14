package states

import (
	"fmt"
	"strings"
	"time"
)

func init() {
	loadDerivedStateChunkLookUpMap()
}

func FindOccupiedSpaces(layout []string) (occupiedSeats uint16) {
	start := time.Now()

	numStateChanges := 0
	var grid, prevGrid *StateGrid = StateGridFromLayout(layout), nil
	fmt.Println("iterating through grid states")
	for !grid.Equals(prevGrid) {
		numStateChanges++
		prevGrid, grid = grid, grid.NextState()
	}
	occupiedSeats = prevGrid.OccupiedSeats()
	fmt.Printf("  found %d occupied seats after %d state changes in %v.\n", occupiedSeats, numStateChanges, time.Since(start))
	return
}

// State chunk lookup is a map of 4x4 state chunks mapped to their next inner state
const stateChunkLen int = 4
const derivedStateChunkLen int = stateChunkLen - 2

// Represents a state chunk to provide context for the next state
type stateChunk [stateChunkLen][stateChunkLen]State // 3^16 or 43,046,721 variations
// Represents the inner derivable portion of a state chunk
type derivedStateChunk [stateChunkLen - 2][stateChunkLen - 2]State // 3^4 or 81 variations
type serializedStateChunk string                                   // [0-2]{16}, represents a base 3 number

func (sc stateChunk) String() string {
	var builder strings.Builder
	builder.WriteString("\n")
	for _, row := range sc {
		for _, el := range row {
			builder.WriteString(el.String())
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (dsc derivedStateChunk) String() string {
	var builder strings.Builder
	builder.WriteString("\n")
	for _, row := range dsc {
		for _, el := range row {
			builder.WriteString(el.String())
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (ssc serializedStateChunk) PrettyPrint() string {
	var builder strings.Builder
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("%s\n", string(ssc[:4])))
	builder.WriteString(fmt.Sprintf("%s\n", string(ssc[4:8])))
	builder.WriteString(fmt.Sprintf("%s\n", string(ssc[8:12])))
	builder.WriteString(fmt.Sprintf("%s\n", string(ssc[12:])))
	return builder.String()
}

func (sc *stateChunk) toSerializedStateChunk() serializedStateChunk {
	var serializedStateBuilder strings.Builder
	for _, row := range sc {
		for _, state := range row {
			serializedStateBuilder.WriteString(state.String())
		}
	}
	fmt.Println("generated serialized state chunk:", serializedStateBuilder.String())
	return serializedStateChunk(serializedStateBuilder.String())
}

func (sc *stateChunk) deriveState() derivedStateChunk {
	return derivedStateChunkFromSerializedStateChunk(sc.toSerializedStateChunk())
}

