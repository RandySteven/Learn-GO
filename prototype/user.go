package prototype

type Employee struct {
	ID       int
	FullName string
	Email    string
	Password string
	Salary   int
}

type Employees []Employee

type IEmployee interface {
	GetAllEmployees() Employees
	PostEmployee(employee Employee, position string) (Employee, Employees)
	GetEmployeeById(ID int) Employee
}

const (
	JUNIOR  string = "Junior"
	SENIOR  string = "Senior"
	MANAGER string = "Manager"
)

type Junior struct {
	Employee Employee
}

type Senior struct {
	Employee Employee
}

type Manager struct {
	Employee Employee
}

var employees Employees

func (e *Employee) PostEmployee(employee Employee, position string) (Employee, Employees) {
	var junior Junior
	var senior Senior
	var manager Manager
	var postEmployee *Employee
	if position == JUNIOR {
		employee.Salary = 80000
		junior.Employee = employee
		postEmployee = &junior.Employee
	} else if position == SENIOR {
		employee.Salary = 90000
		senior.Employee = employee
		postEmployee = &senior.Employee
	} else {
		employee.Salary = 100000
		manager.Employee = employee
		postEmployee = &manager.Employee
	}
	employees = append(employees, *postEmployee)
	return *postEmployee, employees
}

func (e *Employee) GetAllEmployees() Employees {
	if len(employees) <= 0 {
		return nil
	}
	return employees
}
