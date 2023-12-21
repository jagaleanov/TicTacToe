package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// var board [3][3]rune
	board := [3][3]rune{
		{' ', ' ', ' '}, 
		{' ', ' ', ' '}, 
		{' ', ' ', ' '}}

	printBoard(&board)
	play(&board)

}

func play(board *[3][3]rune) {
	turn := 1
	for {
		if turn%2 == 0 {
			machineTurn(board)
		} else {
			userTurn(board)
		}
		printBoard(board)

		winner := checkWinner(board)

		if winner != ' ' {
			fmt.Printf("El ganador es %c \n", winner)
			break
		}

		if !checkForSpace(board) {
			fmt.Println("El juego finalizó sin ganador")
			break
		}

		turn++
	}
}

func checkForSpace(board *[3][3]rune) bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == ' ' {
                return true
            }
        }
    }
    return false
}

func printBoard(board *[3][3]rune) {
	for i := 0; i < 3; i++ {
		fmt.Printf("| %c | %c | %c |\n", board[i][0], board[i][1], board[i][2])
	}
	fmt.Println(" ")
}

func checkWinner(board *[3][3]rune) rune {
	// Verificar filas y columnas
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != ' ' {
			return board[i][0]
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != ' ' {
			return board[0][i]
		}
	}

	// Verificar diagonales
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != ' ' {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != ' ' {
		return board[0][2]
	}

	// No hay ganador todavía
	return ' '
}

// MACHINE TURN
func machineTurn(board *[3][3]rune) {
	machineRune := 'O'
	gamerRune := 'X'
	row, col := findOpportunity(board, machineRune)
	if row > -1 && col > -1 {
		// fmt.Printf("oportunidad ofensiva [%d,%d]\n",row,col)
		setPosition(board, row, col, machineRune)
	} else {
		row, col := findOpportunity(board, gamerRune)
		if row > -1 && col > -1 {
			// fmt.Printf("oportunidad defensiva [%d,%d]\n",row,col)
			setPosition(board, row, col, machineRune)
		} else {
			row, col := findSecondOpportunity(board, machineRune)
			if row > -1 && col > -1 {
				// fmt.Printf("oportunidad secundaria [%d,%d]\n",row,col)
				setPosition(board, row, col, machineRune)
			} else {

				for {
					// fmt.Println("Generando rand")
					randomNumber := rand.Intn(9) + 1
					row, col := numberToPosition(randomNumber)
					// fmt.Printf("Posicion generada [%d,%d]\n",row,col)

					if isEmpty(board, row, col) {
						// fmt.Println("Rand aceptado")
						setPosition(board, row, col, machineRune)
						break
					}
				}
			}
		}
	}
}

func findOpportunity(board *[3][3]rune, playerRune rune) (row int, col int) {
	// fmt.Printf("buscando oportunidad con la runa %c \n",playerRune)
	emptyRune := ' '
	// Verificar filas y columnas
	for i := 0; i < 3; i++ {
		// fmt.Println("condicional horizontal")
		if board[i][0] == board[i][1] && board[i][0] == playerRune && board[i][2] == emptyRune {
			return i, 2
		} else if board[i][0] == board[i][2] && board[i][0] == playerRune && board[i][1] == emptyRune {
			return i, 1

		} else if board[i][1] == board[i][2] && board[i][1] == playerRune && board[i][0] == emptyRune {
			return i, 0
		}

		// fmt.Println("condicional vertical")
		if board[0][i] == board[1][i] && board[0][i] == playerRune && board[2][i] == emptyRune {
			return 2, i
		} else if board[0][i] == board[2][i] && board[0][i] == playerRune && board[1][i] == emptyRune {
			return 1, i
		} else if board[1][i] == board[2][i] && board[1][i] == playerRune && board[0][i] == emptyRune {
			return 0, i
		}
	}

	//Verificar diagonales
		// fmt.Println("condicional diagonal 1")
	if board[0][0] == board[1][1] && board[0][0] == playerRune && board[2][2] == emptyRune {
		return 2, 2
	} else if board[0][0] == board[2][2] && board[0][0] == playerRune && board[1][1] == emptyRune {
		return 2, 1

	} else if board[1][1] == board[2][2] && board[1][1] == playerRune && board[0][0] == emptyRune {
		return 0, 0
	}

	// fmt.Println("condicional diagonal 2")
	if board[0][2] == board[1][1] && board[0][2] == playerRune && board[2][0] == emptyRune {
		return 2, 0
	} else if board[0][2] == board[2][0] && board[0][2] == playerRune && board[1][1] == emptyRune {
		return 1, 1

	} else if board[1][1] == board[2][0] && board[1][1] == playerRune && board[0][2] == emptyRune {
		return 0, 2
	}
	return -1, -1

}

func findSecondOpportunity(board *[3][3]rune, playerRune rune) (row int, col int) {
	emptyRune := ' '
	// Verificar filas y columnas
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][0] == emptyRune && board[i][2] == playerRune {
			return i, 2
		} else if board[i][0] == board[i][2] && board[i][0] == emptyRune && board[i][1] == playerRune {
			return i, 1

		} else if board[i][1] == board[i][2] && board[i][1] == emptyRune && board[i][0] == playerRune {
			return i, 0
		}

		if board[0][i] == board[1][i] && board[i][0] == emptyRune && board[2][i] == playerRune {
			return 2, i
		} else if board[0][i] == board[2][i] && board[0][i] == emptyRune && board[1][i] == playerRune {
			return 1, i
		} else if board[1][i] == board[2][i] && board[1][i] == emptyRune && board[0][i] == playerRune {
			return 0, i
		}
	}

	//Verificar diagonales
	if board[0][0] == board[1][1] && board[0][0] == emptyRune && board[2][2] == playerRune {
		return 2, 2
	} else if board[0][0] == board[2][2] && board[0][0] == emptyRune && board[1][1] == playerRune {
		return 2, 1

	} else if board[1][1] == board[2][2] && board[1][1] == emptyRune && board[0][0] == playerRune {
		return 0, 0
	}

	if board[0][2] == board[1][1] && board[0][2] == emptyRune && board[2][0] == playerRune {
		return 2, 0
	} else if board[0][2] == board[2][0] && board[0][2] == emptyRune && board[1][1] == playerRune {
		return 1, 1

	} else if board[1][1] == board[2][0] && board[1][1] == emptyRune && board[0][2] == playerRune {
		return 0, 2
	}
	return -1, -1

}

// USER TURN
func userTurn(board *[3][3]rune) {
	var number int
	for {
		fmt.Print("Ingrese un número entre 1 y 9: ")
		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("Por favor, ingrese un número válido.")
			continue
		}

		if number < 1 || number > 9 {
			fmt.Println("El número debe estar entre 1 y 9.")
			continue
		}

		row, col := numberToPosition(number)
		if !isEmpty(board, row, col) {
			fmt.Println("La casilla no esta disponible.")
			continue
		} else {
			fmt.Println("Número ingresado:", number)
			setPosition(board, row, col, 'X')
			break
		}
	}
}

func numberToPosition(number int) (row int, col int) {
	row = (number - 1) / 3
	col = (number - 1) % 3
	// fmt.Print("number")
	// fmt.Println(number)
	// fmt.Print("row")
	// fmt.Println(row)
	// fmt.Print("col")
	// fmt.Println(col)
	return row, col
}

func getPosition(board *[3][3]rune, row int, col int) rune {
	return board[row][col]
}

func isEmpty(board *[3][3]rune, row int, col int) bool {
	position := getPosition(board, row, col)
	return position == ' '
}

func setPosition(board *[3][3]rune, row int, col int, value rune) {
	board[row][col] = value
}
