package main

import (
	"fmt"
	"learnGo/prototype"
	"log"
	"strings"
)

func twoSums(arr []int, target int) (int, int) {

	arrSize := len(arr)
	numbMap := make(map[int]int)
	for i := 0; i < arrSize; i++ {
		competence := target - arr[i]
		val, ok := numbMap[competence]
		if ok {
			log.Printf("%d %d", val, i)
			return val, i
		} else {
			numbMap[arr[i]] = i
		}
	}

	return -1, -2
}

func palindrome(word string) bool {
	reverseWord := ""
	wordLen := len(word)

	for i := wordLen - 1; i >= 0; i-- {
		reverseWord = reverseWord + string(word[i])
	}

	log.Println("Reverse word ", reverseWord)
	log.Println("word ", word)

	same := strings.EqualFold(reverseWord, word)
	if same {
		log.Println("is palindrome")
	} else {
		log.Println("nah")
	}

	return same
}

func fibonacci(size int) []int {
	var fibonaccis []int
	var first int
	var second int
	var result int
	if size <= 1 {
		fibonaccis[0] = 0
	} else {
		first = 0
		second = 1
		fibonaccis = append(fibonaccis, first)
		fibonaccis = append(fibonaccis, second)

		for j := 0; j < size; j++ {
			result = first + second
			fibonaccis = append(fibonaccis, result)
			first = second
			second = result
			result = first
		}
	}

	for i := 0; i < len(fibonaccis); i++ {
		log.Println(fibonaccis[i])
	}

	return fibonaccis
}

func main() {
	//db := &singleton.Database{}
	//for i := 0; i < 10; i++ {
	//	fmt.Println(db.GetInstance())
	// }
	//x := []int{
	//	1, 1, 3, 4, 7, 6,
	//}
	//index1, index2 := twoSums(x, 8)
	//log.Printf("index at %d the value is %d", index1, x[index1])
	//log.Printf("index at %d the value is %d", index2, x[index2])
	//
	//isPalindrome := palindrome("Ana")
	//log.Println(isPalindrome)

	fibonaccis := fibonacci(10)
	log.Println(fibonaccis)

	//book := &prototype.Book{
	//	"Tidak tau siapa aku",
	//	20,
	//	"Aku",
	//}
	//
	//novel := &prototype.Novel{
	//	book,
	//	"Psychology",
	//}
	//
	//novel.Read()
	//novel.CurrentPage(novel.Book.Page)
	//var input int
	//for {
	//	input = menu(input)
	//	if input == 3 {
	//		break
	//	}
	//}

}

func menu(input int) int {
	var employee prototype.Employee
	var employees prototype.Employees

	fmt.Println("1. Register User")
	fmt.Println("2. Login User")
	fmt.Println("3. Exit")
	fmt.Print(">>")
	fmt.Scanln(&input)

	if input == 1 {
		var ID int
		var FullName string
		var Email string
		var Password string

		fmt.Print("ID : ")
		fmt.Scanln(&ID)
		fmt.Print("Fullname : ")
		fmt.Scanln(&FullName)
		fmt.Print("Email : ")
		fmt.Scanln(&Email)
		fmt.Print("Password : ")
		fmt.Scanln(&Password)
		var positionChooser int
		positions := []string{
			"1. Manager",
			"2. Senior",
			"3. Junior",
		}
		for _, position := range positions {
			fmt.Println(position)
		}
		fmt.Print("Position : ")
		fmt.Scanln(&positionChooser)

		if positionChooser == 1 {
			position := "Manager"
			employee = prototype.Employee{ID, FullName, Email, Password, 0}
			employee, employees = employee.PostEmployee(employee, position)
			fmt.Println("Employee : ", employee)
		} else if positionChooser == 2 {
			position := "Senior"
			employee = prototype.Employee{ID, FullName, Email, Password, 0}
			employee, employees = employee.PostEmployee(employee, position)
			fmt.Println("Employee : ", employee)
		} else {
			position := "Junior"
			employee = prototype.Employee{ID, FullName, Email, Password, 0}
			employee, employees = employee.PostEmployee(employee, position)
			fmt.Println("Employee : ", employee)
		}

	} else if input == 2 {
		var email string
		var password string
		success := false
		fmt.Print("email : ")
		fmt.Scanln(&email)
		fmt.Print("password : ")
		fmt.Scanln(&password)

		employees = employee.GetAllEmployees()
		index := -1
		var searchEmployee prototype.Employee
		for i, employee := range employees {
			if email == employee.Email && password == employee.Password {
				success = true
				index = i
			}
			searchEmployee = employees[index]
		}

		if success == true {
			fmt.Println("User berhasil ditemukan")
			mainMenu(searchEmployee)
		} else {
			fmt.Println("User tidak ditemukan")
		}

	} else {
		fmt.Println("Bye")
	}

	return input
}

func mainMenu(employee prototype.Employee) {
	var input int
	fmt.Println("Welcome : ", employee.FullName)
	fmt.Println("1. See all users")
	fmt.Println("2. Logout")
	fmt.Print(">>")
	fmt.Scanln(&input)

}
