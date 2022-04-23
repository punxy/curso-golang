package main

import "fmt"

// TaskList
type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) addTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *TaskList) printTasks() {
	for i, task := range tl.tasks {
		fmt.Printf("\nTarea numero %d\n", i+1)
		task.toPrint()
	}
}

// Task
type Task struct {
	name string
	desc string
	done bool
}

func (p *Task) check() {
	p.done = !p.done
}

func (p *Task) toPrint() {
	fmt.Printf("Nombre:\t%s\nDesc:\t%s\nDone:\t%t\n", p.name, p.desc, p.done)
}

func main() {
	task1 := Task{
		name: "Aprender GO",
		desc: "terminar curso para aprender Golang",
		done: false,
	}

	task1.toPrint()
	task1.check()
	task1.toPrint()

	task2 := Task{
		name: "Aprender React",
		desc: "Temrinar curso para aprender React",
		done: false,
	}

	listado := TaskList{}
	listado.addTask(&task1)
	listado.addTask(&task2)

	// fmt.Println(listado)       // muestra las refreencias a la memoria
	// listado.tasks[0].toPrint() // muestra la primera tarea

	listado.printTasks()
	listado.removeTask(0)
	listado.printTasks()

}
