package server

import pb "RouteRaydar/api"

type Matrix struct {
	Rows int64
	Cols int64
	Data [][]int64
}

func NewMatrix(rows, cols int64) *Matrix {
	data := make([][]int64, rows)
	for i := range data {
		data[i] = make([]int64, cols)
	}
	return &Matrix{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func (m *Matrix) Set(row, col int, value int64) {
	m.Data[row][col] = value
}

func (m *Matrix) Get(row, col int) int64 {
	return m.Data[row][col]
}

// Any error checks needed here?
func (m *Matrix) ToProto() *pb.Grid {
	grid := &pb.Grid{}
	for _, row := range m.Data {
		pbRow := &pb.Row{}
		pbRow.Values = append(pbRow.Values, row...)
		grid.Rows = append(grid.Rows, pbRow)
	}
	return grid
}
