package main

import (
	"fmt"
	"encoding/json"
)

type Matrix struct {
    row,col int
	mat [][]int
}
func GetRows(row,col int) int{
    getrow := Matrix{row: row, col: col}
    return getrow.row
}

func GetColumns(row,col int) int{
    getcol := Matrix{row: row, col: col}
    return getcol.col
}
func AddingMatrices(row,col int) *Matrix {
    mat1:= make([][]int, row, col)
    mat2:= make([][]int, row, col)
     m := Matrix{row: row, col: col, mat: mat1}
    n := Matrix{row: row, col: col, mat: mat1}
     add :=  Matrix{row: row, col: col, mat:mat2}
    for i := 0; i < row; i++ {
        m.mat[i] = make([]int, col)
        for j := 0; j < col; j++ {
         m.mat[i][j] = i+j
        }
    }
    for i := 0; i < row; i++ {
        n.mat[i] = make([]int, col)
        for j := 0; j < col; j++ {
           n.mat[i][j] = i+j
        }
    }
    for i := 0; i < row; i++ {
        add.mat[i] = make([]int, col)
        for j := 0; j < col; j++ {
            add.mat[i][j] = m.mat[i][j] + n.mat[i][j]
        }
    }
    return &add
}
func SetMatrixElements(row,col int) *Matrix {
    mat4 :=make([][]int, row, col)
    p := Matrix{row: row, col: col,mat: mat4}
    for i := 0; i < row; i++ {
        p.mat[i] = make([]int, 3)
        for j := 0; j < col; j++ {
            fmt.Scanln(&p.mat[i][j])
        }
    }
  return &p
}

func Printinjson(row,col int) {
    roww := row
    coll := col
   l :=SetMatrixElements(roww,coll)
   hey,err := json.MarshalIndent(l,"","  ")
   if err!= nil {
       fmt.Println(err)
   }
   fmt.Println(hey)

fmt.Println("Elements", l)
}


func main() {
 fmt.Println("Enter tne number of rows ")
 var r int
 fmt.Scanln(&r)
 fmt.Println("Enter tne number of columns ")
 var c int
 fmt.Scanln(&c)
 fmt.Println("Number of Rows", GetRows(r,c))
 fmt.Println("Number of columns", GetColumns(r,c))
 set := SetMatrixElements(r,c)
 fmt.Println(set.mat)
 added := AddingMatrices(r,c)
 fmt.Println(added.mat)
 Printinjson(r,c)
}