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
	}

	defer file.Close()

	var Tasks []Task
	json.NewDecoder(file).Decode(&Tasks)

	NewTask := Task{
		ID:          len(Tasks) + 1,
		Name:        name,
		Description: description,
		Complete:    false,
	}

	Tasks = append(Tasks, NewTask)

	file.Seek(0, 0)
	file.Truncate(0)
	json.NewEncoder(file).Encode(Tasks)

	fmt.Println("Task creating with exit")

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
