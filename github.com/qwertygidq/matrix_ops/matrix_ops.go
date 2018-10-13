package matrix_ops

import (
	"errors"
	"fmt"
)

type Matrix [][]float64

func CreateEmptyMatrix(rows, cols int) Matrix { // TODO: check for valid rows, cols
	var mat Matrix = make(Matrix, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, cols)
	}

	return mat
}

func CopyMatrix(mat Matrix) Matrix {
	var new_mat Matrix = CreateEmptyMatrix(Size(mat))
	for row_k := range mat {
		copy(new_mat[row_k], mat[row_k])
	}

	return new_mat
}

func CreateIdentityMatrix(size int) Matrix {
	var mat Matrix = CreateEmptyMatrix(size, size)
	for row_k := range mat {
		mat[row_k][row_k] = 1
	}

	return mat
}

func CheckMatrix(mat Matrix) bool {
	rows, cols := Size(mat)
	if rows == 0 && cols == 0 {
		return true
	} else if rows == 0 || cols == 0 {
		return false
	} else {
		for row := 1; row < rows; row++ {
			if len(mat[row]) != cols {
				return false
			}
		}

		return true
	}
}

func Size(mat Matrix) (int, int) {
	if len(mat) == 0 {
		return 0, 0
	} else {
		return len(mat), len(mat[0])
	}
}

func PrintMatrix(mat Matrix) {
	for _, row := range mat {
		for _, val := range row {
			fmt.Printf("%v ", val)
		}

		fmt.Printf("\n")
	}
}

func TransposeMatrix(mat Matrix) (Matrix, error) {
	if !CheckMatrix(mat) {
		return nil, errors.New("Bad matrix")
	}

	rows, cols := Size(mat)
	var new_mat Matrix = CreateEmptyMatrix(cols, rows)
	for row_k, row := range mat {
		for col_k := range row {
			new_mat[col_k][row_k] = mat[row_k][col_k]
		}
	}

	return new_mat, nil
}

func MultiplyMatrixScalar(mat Matrix, scalar float64) (Matrix, error) {
	if !CheckMatrix(mat) {
		return nil, errors.New("Bad matrix")
	}

	var new_mat Matrix = CopyMatrix(mat)

	for _, row := range new_mat {
		for col_k := range row {
			row[col_k] *= scalar
		}
	}

	return new_mat, nil
}

func MultiplyMatrices(mat1, mat2 Matrix) (Matrix, error) {
	if !CheckMatrix(mat1) {
		return nil, errors.New("Bad first matrix")
	} else if !CheckMatrix(mat2) {
		return nil, errors.New("Bad second matrix")
	}

	rows1, cols1 := Size(mat1)
	rows2, cols2 := Size(mat2)
	if cols1 != rows2 {
		return nil, errors.New("Number of columns of the first matrix should " +
			"be equal to the number of rows of the second " +
			"matrix")
	} else {
		var new_mat Matrix = CreateEmptyMatrix(rows1, cols2)
		for row_k, row := range new_mat {
			for col_k := range row {
				var sum float64
				for i := 0; i < cols1; i++ {
					sum += mat1[row_k][i] * mat2[i][col_k]
				}

				new_mat[row_k][col_k] = sum
			}
		}

		return new_mat, nil
	}
}

func SumMatrices(mat1, mat2 Matrix) (Matrix, error) {
	if !CheckMatrix(mat1) {
		return nil, errors.New("Bad first matrix")
	} else if !CheckMatrix(mat2) {
		return nil, errors.New("Bad second matrix")
	}

	rows1, cols1 := Size(mat1)
	rows2, cols2 := Size(mat2)

	if rows1 != rows2 || cols1 != cols2 {
		return nil, errors.New("Matrix sizes should be equal")
	} else {
		var new_mat = CopyMatrix(mat1)
		for row_k, row := range new_mat {
			for col_k := range row {
				new_mat[row_k][col_k] += mat2[row_k][col_k]
			}
		}

		return new_mat, nil
	}
}

func ScalarMultiplyVectors(vec1, vec2 Matrix) (float64, error) {
	// First we check for errors
	if !CheckMatrix(vec1) {
		return 0, errors.New("Bad first vector")
	} else if !CheckMatrix(vec2) {
		return 0, errors.New("Bad second vector")
	}

	rows1, cols1 := Size(vec1)
	rows2, cols2 := Size(vec2)

	// Then we check if these VALID matrices are vectors
	if !(rows1 == 1 || cols1 == 1) {
		return 0, errors.New("First matrix is not a vector")
	} else if !(rows2 == 1 || cols2 == 1) {
		return 0, errors.New("Second matrix is not a vector")
	}

	var new_vec1, new_vec2 Matrix = vec1, vec2
	var err error
	if rows1 != 1 {
		new_vec1, err = TransposeMatrix(CopyMatrix(vec1))
		if err != nil {
			return 0, errors.New("Couldn't transpose the first matrix")
		}
	}
	if rows2 != 1 {
		new_vec2, err = TransposeMatrix(CopyMatrix(vec2))
		if err != nil {
			return 0, errors.New("Couldn't transpose the second matrix")
		}
	}

	_, cols1 = Size(new_vec1)
	_, cols2 = Size(new_vec2)

	if cols1 != cols2 {
		return 0, errors.New("Vectors should be equal in length")
	}

	var sum float64
	for col := 0; col < cols1; col++ {
		sum += new_vec1[0][col] * new_vec2[0][col]
	}

	return sum, nil
}
