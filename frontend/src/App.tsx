import React, { useEffect, useState } from 'react';
import './App.css';
import Todo, { TodoType } from './Todo';

function App() {
  const [todos, setTodos] = useState<TodoType[]>([]);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  // Initially fetch todo
  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const todos = await fetch('http://localhost:8080/');
        if (todos.status !== 200) {
          console.log('Error fetching data');
          return;
        }

        setTodos(await todos.json());
      } catch (e) {
        console.log('Could not connect to server. Ensure it is running. ' + e);
      }
    }

    fetchTodos()
  }, []);

  // handle posting a new To-Do
  const postToDo = async () => {
    if (!title || !description) {
      alert("Title and description are required!");
      return;
    }

    const newToDo = { title, description };

    try {
      const response = await fetch('http://localhost:8080/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newToDo),
    });

    if (!response.ok) {
      const errorMessage = await response.text();
      console.log(errorMessage);
    }

    // update the list with the new item
    const updatedTodos = await response.json();
    setTodos(updatedTodos); 

  } catch (error) {
    console.log(error);
  }
};


  return (
    <div className="app">
      <header className="app-header">
        <h1>TODO</h1>
      </header>

      <div className="todo-list">
        {todos.map((todo) =>
          <Todo
            key={todo.title + todo.description}
            title={todo.title}
            description={todo.description}
          />
        )}
      </div>

      <h2>Add a Todo</h2>
      <form>
        <input onChange={(e) => setTitle(e.target.value)} placeholder="Title" name="title" autoFocus={true} />
        <input onChange={(e) => setDescription(e.target.value)} placeholder="Description" name="description" />
        <button onClick={(e) => { e.preventDefault(); postToDo()}}>Add Todo</button>
      </form>
    </div>
  );
}

export default App;
