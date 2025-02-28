package c

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

func CreateFile() {
	_, err := os.Stat("task.json")

	if os.IsNotExist(err) {
		file, err := os.Create("task.json")
		if err != nil {
			fmt.Println("Error Creating the file task.json", err)
			return
		}
		defer file.Close()

		Tasks := []Task{}

		json.NewEncoder(file).Encode(Tasks)

		fmt.Println("The file task.json was created successfully")
	}
}

func AddTask(name, description string) {
	file, err := os.OpenFile("task.json", os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Could not open file")
		return
	}

	defer file.Close()

	var Tasks []Task
	err = json.NewDecoder(file).Decode(&Tasks)

	if err != nil {
		fmt.Println("Warning: Could not decode JSON, initializing empty list.")
		Tasks = []Task{}
	}

	maxID := 0
	for _, task := range Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	NewTask := Task{
		ID:          maxID + 1,
		Name:        name,
		Description: description,
		Complete:    false,
	}

	Tasks = append(Tasks, NewTask)

	file.Seek(0, 0)
	file.Truncate(0)
	json.NewEncoder(file).Encode(Tasks)

	fmt.Println("Task created successfully")

}

func LisTask() {
	file, err := os.OpenFile("task.json", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("error opening task.json", err)
		return
	}
	defer file.Close()

	var Tasks []Task
	err = json.NewDecoder(file).Decode(&Tasks)

	if err != nil {
		fmt.Println("Erorr decoding json:", err)
		return
	}

	if len(Tasks) == 0 {
		fmt.Println("No task found")
		return
	}
	fmt.Println("📋 Task List:")
	for _, task := range Tasks {
		status := "❌"
		if task.Complete {
			status = "✅"
		}
		fmt.Printf("[%s] %d. %s - %s\n", status, task.ID, task.Name, task.Description)
	}

}

func DeleteTask(id int) {
	file, err := os.OpenFile("Task.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("error opening Task.json")
		return
	}
	defer file.Close()

	var Tasks []Task
	err = json.NewDecoder(file).Decode(&Tasks)
	if err != nil {
		fmt.Println("Error decoding json")
		return
	}

	var UpdateTask []Task
	found := false

	for _, task := range Tasks {
		if task.ID != id {
			UpdateTask = append(UpdateTask, task)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found.\n", id)
		return
	}
	file.Truncate(0)
	file.Seek(0, 0)

	err = json.NewEncoder(file).Encode(UpdateTask)
	if err != nil {
		fmt.Println("error encoding update task:", err)
		return
	}
	fmt.Printf("Task with ID %d deleted successfully.\n", id)

}

func UpdateTask(id int) {
	file, err := os.OpenFile("Task.json", os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Error opening task")
		return
	}
	defer file.Close()

	var Tasks []Task
	err = json.NewDecoder(file).Decode(&Tasks)
	if err != nil {
		fmt.Println("error decoding task.json")
		return
	}

	found := false

	for i, task := range Tasks {
		if task.ID == id {
			Tasks[i].Complete = true
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Task with ID %d not found.\n", id)
		return
	}

	file.Seek(0, 0)
	file.Truncate(0)

	err = json.NewEncoder(file).Encode(Tasks)

	if err != nil {
		fmt.Println("error encoding task.json")
		return
	}
	LisTask()
}
