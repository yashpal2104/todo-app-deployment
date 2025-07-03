<!-- my-todo-frontend/src/routes/+page.svelte -->

<script>
	import { onMount } from 'svelte';

	// The base URL for our Go backend API
	const API_URL = '/api';// The browser will handle the rest!

	// Reactive variables to hold our application state
	let tasks = [];
	let newTaskTitle = '';
	let isLoading = true;
	let error = null;

	// This function runs when the component is first mounted to the page
	onMount(async () => {
		fetchTasks();
	});

	// --- API Functions ---

	// Fetches all tasks from the backend
	async function fetchTasks() {
		try {
			const response = await fetch(`${API_URL}/tasks`);
			if (!response.ok) {
				throw new Error('Failed to fetch tasks');
			}
			tasks = await response.json();
			// Sort tasks by ID to keep the list stable
			tasks.sort((a, b) => a.id - b.id);
		} catch (e) {
			error = e.message;
		} finally {
			isLoading = false;
		}
	}

	// Adds a new task
	async function addTask(event) {
		event.preventDefault(); // Prevent form from reloading the page
		if (!newTaskTitle.trim()) return; // Don't add empty tasks

		try {
			const response = await fetch(`${API_URL}/tasks`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ title: newTaskTitle })
			});
			const addedTask = await response.json();
			
			// Add the new task to our local list to update the UI instantly
			tasks = [...tasks, addedTask];
			newTaskTitle = ''; // Clear the input field
		} catch (e) {
			error = 'Failed to add task.';
		}
	}

	// Toggles the 'completed' status of a task
	async function toggleTask(task) {
		const updatedTask = { ...task, completed: !task.completed };

		try {
			await fetch(`${API_URL}/tasks/${task.id}`, {
				method: 'PUT',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(updatedTask)
			});
			// Update the task in our local list
			const taskIndex = tasks.findIndex(t => t.id === task.id);
			tasks[taskIndex] = updatedTask;
		} catch (e) {
			error = 'Failed to update task.';
		}
	}

	// Deletes a task
	async function deleteTask(taskId) {
		try {
			await fetch(`${API_URL}/tasks/${taskId}`, {
				method: 'DELETE'
			});
			// Remove the task from our local list
			tasks = tasks.filter(t => t.id !== taskId);
		} catch (e) {
			error = 'Failed to delete task.';
		}
	}
</script>

<main>
	<h1>Go + Svelte TODO List</h1>

	<!-- Form for adding new tasks -->
	<form on:submit={addTask}>
		<input
			type="text"
			bind:value={newTaskTitle}
			placeholder="What needs to be done?"
			aria-label="New task title"
		/>
		<button type="submit">Add Task</button>
	</form>

	<!-- Display loading, error, or the task list -->
	{#if isLoading}
		<p>Loading tasks...</p>
	{:else if error}
		<p class="error">{error}</p>
	{:else}
		<!-- Task list -->
		<ul>
			{#each tasks as task (task.id)}
				<li class:completed={task.completed}>
					<input
						type="checkbox"
						checked={task.completed}
						on:change={() => toggleTask(task)}
						aria-label="Mark task as complete"
					/>
					<span>{task.title}</span>
					<button class="delete" on:click={() => deleteTask(task.id)}>
						×
					</button>
				</li>
			{/each}
		</ul>
	{/if}
</main>

<style>
	main {
		max-width: 600px;
		margin: 2rem auto;
		font-family: sans-serif;
		text-align: center;
	}
	form {
		display: flex;
		margin-bottom: 1.5rem;
	}
	input[type='text'] {
		flex-grow: 1;
		padding: 0.5rem;
		font-size: 1rem;
		border: 1px solid #ccc;
		border-radius: 4px 0 0 4px;
	}
	button {
		padding: 0.5rem 1rem;
		font-size: 1rem;
		border: 1px solid #007bff;
		background-color: #007bff;
		color: white;
		cursor: pointer;
		border-radius: 0 4px 4px 0;
	}
    button[type='submit'] {
        border-radius: 0 4px 4px 0;
    }
	ul {
		list-style: none;
		padding: 0;
	}
	li {
		display: flex;
		align-items: center;
		padding: 0.5rem;
		border-bottom: 1px solid #eee;
	}
	li.completed span {
		text-decoration: line-through;
		color: #999;
	}
	li span {
		flex-grow: 1;
		text-align: left;
		margin-left: 0.5rem;
	}
	.delete {
		background: none;
		border: none;
		color: red;
		font-size: 1.5rem;
		font-weight: bold;
		cursor: pointer;
		padding: 0 0.5rem;
	}
    .error {
        color: red;
    }
</style>