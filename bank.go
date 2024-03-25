package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("- - - - - - - - - ")
		panic("Can't continue, sorry ...")
	}

	for {
		presentOptions()

		var choise int
		fmt.Println("Your choice:")
		fmt.Scan(&choise)

		switch choise {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Println("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated. New amount:", accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Println("Withdrawal amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid withdrawal. Must be greater than 0.")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid withdrawal. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated. New amount:", accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
