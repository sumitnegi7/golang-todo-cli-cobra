package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Todo struct {
	// Id int `json:"id"`
	Task        string    `json:"task"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
}


func createFile() {
	filePtr, err := os.Create("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
}


func listTodo() (err error){
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		return fmt.Errorf("Error opening file:", err)
	}
	defer file.Close()

	// Read the file's content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Error reading file:", err)
	}

	var todos []Todo

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(byteValue, &todos)
	if err != nil {
		return fmt.Errorf("Error reading file:", err)
	}

	// Print the result
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Todo", "Completed","CreatedAt"})
	
	for i, todo := range todos {
		var status string
		if todo.IsCompleted {
			status = " ✅  True"
		} else {
			status = " ❌  False"
		}
		t.AppendRow(table.Row{i + 1, todo.Task, status, todo.CreatedAt.Format("2006-01-02 15:04:05")})
	}
	t.SetStyle(table.StyleColoredBlackOnRedWhite)
    t.Render()
	
	return nil
}

func addTodo(todo string, Status string) (id int, err error) {
	var todos []Todo
	status, er := strconv.ParseBool(Status)
	if er != nil {
		return 0, fmt.Errorf("invalid data type %v", err)
	}

	file, err := os.Open("data.json")
	if err != nil {
		createFile()
		file, err = os.Open("data.json")
		if err != nil {
			return 0, fmt.Errorf("error in openieng file ")
		}
	}

	defer file.Close()

	newTodo := Todo{
		Task:        todo,
		IsCompleted: status,
		CreatedAt:   time.Now(),
	}

	if file != nil {

		data, _ := io.ReadAll(file)

		if len(data) > 0 {
			err = json.Unmarshal(data, &todos)
			if err != nil {
				return 0, fmt.Errorf("error unmarshalling JSON: %w", err)
			}
		}
		fmt.Println("🚀 ~ files:", todos)
		todos = append(todos, newTodo)
		fmt.Println("🚀 ~ todos:", todos)

		jsonData, err := json.MarshalIndent(todos, "", " ")
		if err != nil {
			return 0, fmt.Errorf("error marshalling JSON: %w", err)
		}

		err = os.WriteFile("data.json", jsonData, 0644)
		if err != nil {
			return 0, fmt.Errorf("error marshalling JSON: %w", err)
		}
	}
	return len(todos), nil

}



func readDataAndUnmarshalJSON() []Todo {
	file, err := os.Open("data.json")

	if err != nil {
		fmt.Println("Error")
	}

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error")
	}
	var todos []Todo

	err = json.Unmarshal(data, &todos)
	if err != nil {
		fmt.Println("Error")
	}

	return todos

}

func updateTask(Id string, updatedTask string)(err error) {

	id,err := strconv.Atoi(Id);
	if err != nil {
		return fmt.Errorf("invalid data type %v", err)
	} 
	var found = false
	todos := readDataAndUnmarshalJSON()

	for i, _ := range todos {
		if(i+1 == id){
			todos[i].Task = updatedTask 
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no such task exists %v", err)
	} else {
		jsonData, err := json.MarshalIndent(todos, "", " ")
		if err != nil {
			fmt.Println("Error")
		}
		os.WriteFile("data.json", jsonData, 0644)
	}
	return nil 
}

func deleteTask(Id string) (err error) {

	id,err := strconv.Atoi(Id);
	if err != nil {
		return fmt.Errorf("invalid data type %v", err)
	} 
	var found = false
	todos := readDataAndUnmarshalJSON()

	for i, _ := range todos {
		if(i+1 == id){
			todos = append(todos[0:i],todos[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no such task exists %v", err)
	} else {
		jsonData, err := json.MarshalIndent(todos, "", " ")
		if err != nil {
			fmt.Println("Error")
		}
		os.WriteFile("data.json", jsonData, 0644)
	}
	return nil 
}


func markDone(Id string) (err error) {

	id,err := strconv.Atoi(Id);
	if err != nil {
		return fmt.Errorf("invalid data type %v", err)
	} 
	var found = false
	todos := readDataAndUnmarshalJSON()

	for i, _ := range todos {
		if(i+1 == id){
			todos[i].IsCompleted = true 
			found = true
			break
		}
	}

	if !found {
		fmt.Println("No such task exists")
	} else {
		jsonData, err := json.MarshalIndent(todos, "", " ")
		if err != nil {
			fmt.Println("Error")
		}
		os.WriteFile("data.json", jsonData, 0644)
	}

	return nil

}