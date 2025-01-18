package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Todo struct {
	Task  string `json:"task"`
	IsCompleted   bool    `json:"isCompleted"`
}


func createFile() {
	filePtr, err := os.Create("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
}


func openFile() {
	filePtr, err :=os.Open("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
}

func readFile(){
// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file's content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}



	var todo []Todo

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(byteValue, &todo)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the result
	for idx, t := range todo {
		fmt.Printf("ID: %s, Name: %s, Age: %t\n", idx, t.Task, t.IsCompleted)
	}
}




func addTodo(todo string , status bool) error {
	var todos []Todo
	file, err := os.Open("data.json")
	if err !=nil {
		return fmt.Errorf("error in openieng file %")
	}
	
	defer file.Close()

	newTodo := Todo{
		Task: todo, 
		IsCompleted: status,
	}
	fmt.Println(file)

	if file != nil {
	
		data,_ := io.ReadAll(file)

		if(len(data)>0){
		err = json.Unmarshal(data,&todos)
		if err != nil {
			return fmt.Errorf("error unmarshalling JSON: %w", err)
			}
		} 
		todos = append(todos,newTodo)
		fmt.Println("🚀 ~ todos:", todos)
		// Marshal the updated todos back to JSON
	jsonData, err := json.MarshalIndent(todos, "", " ")
	fmt.Println("🚀 ~ todos:", jsonData)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	err = os.WriteFile("data.json",jsonData,0644)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}
	return nil
		}
		return fmt.Errorf("file emp:")

}



func readDataAndUnmarshalJSON() []Todo {
	file, err := os.Open("data.json")

	if err !=nil {
		fmt.Println("Error")
	}

	  data, err := io.ReadAll(file)
	  if err !=nil {
		fmt.Println("Error")
	}
	  var todos []Todo 

	  err = json.Unmarshal(data,&todos)
	  if err !=nil {
		fmt.Println("Error")
	}

	return todos

}

// func updateTask(id int, updatedTask string){
// 	var found = false
// 	todos := readDataAndUnmarshalJSON()

// 	for i,todo := range todos {
// 		if(todo.Id == id){
// 			todos[i].Task = updatedTask
// 			found = true
// 			break
// 		}
// 	}

// 	if found != true {
// 		 fmt.Println("No such task exists")
// 	} else {
// 		jsonData, err := json.MarshalIndent(todos, "", " ")
// 		fmt.Println("🚀>> ~ funcupdateTask ~ todos:", todos)
// 		if err !=nil {
// 			fmt.Println("Error")
// 		}
// 		os.WriteFile("data.json",jsonData,0644)
// 	}

// }


// func updateTaskStatus(id int, status bool){
// 	var found = false
// 	todos := readDataAndUnmarshalJSON()

// 	for i,todo := range todos {
// 		fmt.Println("🚀 ~ funcupdateTask ~ todo:", todo, todo.Id == id)
// 		if(todo.Id == id){
// 			todos[i].IsCompleted = status
// 			found = true
// 			break
// 		}
// 	}

// 	if found != true {
// 		 fmt.Println("No such task exists")
// 	} else {
// 		jsonData, err := json.MarshalIndent(todos, "", " ")
// 		fmt.Println("🚀>> ~ funcupdateTask ~ todos:", todos)
// 		if err !=nil {
// 			fmt.Println("Error")
// 		}
// 		os.WriteFile("data.json",jsonData,0644)
// 	}

// }


// func deleteTask(id int){
// 	var found = false
// 	todos := readDataAndUnmarshalJSON()

// 	for i,todo := range todos {
// 		fmt.Println("🚀 ~ funcupdateTask ~ todo:", todo, todo.Id == id)
// 		if(todo.Id == id){
// 			todos = append(todos[:i], todos[i+1:]...)
// 			found = true
// 			break
// 		}
// 	}

// 	if found != true {
// 		 fmt.Println("No such task exists")
// 	} else {
// 		jsonData, err := json.MarshalIndent(todos, "", " ")
// 		fmt.Println("🚀>> ~ funcupdateTask ~ todos:", todos)
// 		if err !=nil {
// 			fmt.Println("Error")
// 		}
// 		os.WriteFile("data.json",jsonData,0644)
// 	}

// }

// func main(){
// 	createFile()
// 	initLastID()
// 	addTodo("Had breakfast", false)
// 	addTodo("Had lunch", false)
// 	addTodo("Exercise", false)
// 	updateTask(4,"Instead have bruch i guesss")
// 	// writeToFile()
// }