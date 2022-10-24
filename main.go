package main

import (
	s "WB_L1/Solutions"
	"fmt"
)

func main() {
	fmt.Println("Введите номер задания (1-26): ")
	var taskNum int

	fmt.Scan(&taskNum)

	switch taskNum {
	case 1:
		s.Task1()
		break
	case 2:
		s.Task2()
		break
	case 3:
		s.Task3()
		break
	case 4:
		s.Task4()
		break
	case 5:
		s.Task5()
		break
	case 6:
		fmt.Println("Введите номер подзадания (1-3): ")
		var subTask int

		fmt.Scan(&subTask)
		switch subTask {
		case 1:
			s.Task6_1()
			break
		case 2:
			s.Task6_2()
			break
		case 3:
			s.Task6_3()
			break
		default:
			fmt.Println("Error: Подзадачи с номером > 3 не существует")
		}
		break
	case 7:
		s.Task7()
		break
	case 8:
		s.Task8()
		break
	case 9:
		s.Task9()
		break
	case 10:
		s.Task10()
		break
	case 11:
		s.Task11()
		break
	case 12:
		s.Task12()
		break
	case 13:
		fmt.Println("Введите номер подзадания (1-2): ")
		var subTask int

		fmt.Scan(&subTask)
		switch subTask {
		case 1:
			s.Task13_1()
			break
		case 2:
			s.Task13_2()
			break
		default:
			fmt.Println("Error: Подзадачи с номером > 2 не существует")
		}
		break
	case 14:
		s.Task14()
		break
	case 15:
		s.Task15()
		break
	case 16:
		s.Task16()
		break
	case 17:
		s.Task17()
		break
	case 18:
		s.Task18()
		break
	case 19:
		s.Task19()
		break
	case 20:
		s.Task20()
		break
	case 21:
		s.Task21()
		break
	case 22:
		s.Task22()
	case 23:
		s.Task23()
		break
	case 24:
		s.Task24()
		break
	case 25:
		fmt.Println("Введите номер подзадания (1-2): ")
		var subTask int

		fmt.Scan(&subTask)
		switch subTask {
		case 1:
			s.Task25_1()
			break
		case 2:
			s.Task25_2()
			break
		default:
			fmt.Println("Error: Подзадачи с номером > 2 не существует")
		}
		break
	case 26:
		s.Task26()
		break

	default:
		fmt.Println("You input wrong number!")
	}
}
