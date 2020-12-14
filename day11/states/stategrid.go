package states

import (
	"fmt"
	"math"
	"strings"
)

type StateGrid [][]State

func StateGridFromLayout(layout []string) *StateGrid {
	nRowsLayout, nColsLayout := len(layout), len(layout[0])
	//nRowsLayout, nColsLayout = 90, 92
	nRowsGrid, nColsGrid := 92, 96
	nRowsGrid, nColsGrid = int(math.Ceil(float64(nRowsLayout+2)/float64(stateChunkLen)))*stateChunkLen, int(math.Ceil(float64(nColsLayout+2)/float64(stateChunkLen)))*stateChunkLen
	fmt.Printf("loading layout with %d x %d into %d x %d grid\n", nRowsLayout, nColsLayout, nRowsGrid, nColsGrid)

	grid := make(StateGrid, nRowsGrid)
	for ix := range grid {
		grid[ix] = make([]State, nColsGrid)
	}

	// Pad the first row of grid
	for colNum := range grid[0] {
		grid[0][colNum] = Floor
	}

	// Transfer layout into grid
	for layoutRowNum, layoutRow := range layout {
		gridRowNum := layoutRowNum + 1
		// Pad first column of row
		grid[gridRowNum][0] = Floor
		for layoutColNum, stateChar := range layoutRow {
			gridColNum := layoutColNum + 1
			grid[gridRowNum][gridColNum] = State(byte(stateChar))
		}
		// Pad remaining columns of row
		for i := nColsLayout + 1; i < nColsGrid; i++ {
			grid[gridRowNum][i] = Floor
		}
	}

	// Pad remaining rows of grid
	for i := nRowsLayout + 1; i < nRowsGrid; i++ {
		grid[i] = make([]State, nColsGrid)
		for rowNum := range grid[i] {
			grid[i][rowNum] = Floor
		}
	}
	fmt.Println("  done.")

	return &grid
}

func (sg StateGrid) String() string {
	var builder strings.Builder
	builder.WriteString("\n")
	for _, row := range sg {
		for _, el := range row {
			builder.WriteString(el.String())
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (sg *StateGrid) Equals(rhs *StateGrid) bool {
	if rhs == nil {
		return sg == nil
	}
	for ix := 0; ix < len(*sg); ix++ {
		for jx := 0; jx < len((*sg)[ix]); jx++ {
			if (*sg)[ix][jx] != (*rhs)[ix][jx] {
				return false
			}
		}
	}
	return true
}

func (sg *StateGrid) NextState() *StateGrid {
	grid := *sg

	nRowsGrid, nColsGrid := len(grid), len(grid[0])

	nextGrid := make(StateGrid, nRowsGrid)
	for ix := range nextGrid {
		nextGrid[ix] = make([]State, nColsGrid)
		for jx := range nextGrid[ix] {
			nextGrid[ix][jx] = grid[ix][jx]
		}
	}

	// Remove padding rows, then find out how many derived state chunks fit
	nRowsStateChunks := (nRowsGrid - 2) / derivedStateChunkLen
	nColsStateChunks := (nColsGrid - 2) / derivedStateChunkLen

	serializedStateChunks := make([][]serializedStateChunk, nRowsStateChunks)
	for ix := 0; ix < nRowsStateChunks; ix++ {
		serializedStateChunks[ix] = make([]serializedStateChunk, nColsStateChunks)
	}

	// Get serialized state chunks
	for stateRowIx, stateChunkRowIx := 0, 0; stateRowIx <= nRowsGrid-stateChunkLen;
	stateRowIx, stateChunkRowIx = stateRowIx+derivedStateChunkLen, stateChunkRowIx+1 {

		stateChunkStartRow, stateChunkEndRow := stateRowIx, stateRowIx+stateChunkLen
		stateChunkRow := grid[stateChunkStartRow:stateChunkEndRow]

		for stateColIx, stateChunkColIx := 0, 0; stateColIx <= nColsGrid-stateChunkLen;
		stateColIx, stateChunkColIx = stateColIx+derivedStateChunkLen, stateChunkColIx+1 {

			var serializedStateChunkBuilder strings.Builder
			stateChunkStartCol, stateChunkEndCol := stateColIx, stateColIx+stateChunkLen

			for interStateChunkRowIx := 0; interStateChunkRowIx < stateChunkLen; interStateChunkRowIx++ {
				for _, state := range stateChunkRow[interStateChunkRowIx][stateChunkStartCol:stateChunkEndCol] {
					serializedStateChunkBuilder.WriteString(state.String())
				}
			}

			serializedStateChunks[stateChunkRowIx][stateChunkColIx] = serializedStateChunk(serializedStateChunkBuilder.String())
		}
	}

	// Derive state chunks from serialized state chunks
	for stateChunkRowIx, serializedStateChunkRow := range serializedStateChunks {
		for stateChunkColIx, serializedStateChunk := range serializedStateChunkRow {
			derivedStateChunk := derivedStateChunkFromSerializedStateChunk(serializedStateChunk)
			stateChunkInGridRowIx, stateChunkInGridColIx := 1+stateChunkRowIx*derivedStateChunkLen, 1+stateChunkColIx*derivedStateChunkLen
			nextGrid[stateChunkInGridRowIx][stateChunkInGridColIx] = derivedStateChunk[0][0]
			nextGrid[stateChunkInGridRowIx+1][stateChunkInGridColIx] = derivedStateChunk[0+1][0]
			nextGrid[stateChunkInGridRowIx][stateChunkInGridColIx+1] = derivedStateChunk[0][0+1]
			nextGrid[stateChunkInGridRowIx+1][stateChunkInGridColIx+1] = derivedStateChunk[0+1][0+1]
		}
	}

	return &nextGrid
}

func (sg *StateGrid) OccupiedSeats() (seatCount uint16) {
	grid := *sg
	for _, row := range grid {
		for _, seat := range row {
			if seat == Occupied {
				seatCount += 1
			}
		}
	}
	return
}

