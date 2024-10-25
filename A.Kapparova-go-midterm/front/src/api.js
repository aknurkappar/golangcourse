const API_URL = 'http://127.0.0.1:8000';

export const fetchTasks = async () => {
    const response = await fetch(`${API_URL}/tasks`);
    if (!response.ok) {
        throw new Error('Failed to fetch tasks');
    }
    return response.json();
};

export const fetchTask = async (id) => {
    try {
        const response = await fetch(`${API_URL}/tasks/${id}`);
        if (!response.ok) {
            throw new Error(`Error fetching task: ${response.statusText}`);
        }
        const task = await response.json();
        return task;
    } catch (error) {
        throw new Error(error.message);
    }
};

export const createTask = async (task) => {
    const response = await fetch(`${API_URL}/tasks/`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(task),
    });
    return response.json();
};

export const updateTask = async (id, task) => {
    const response = await fetch(`${API_URL}/tasks/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(task),
    });
    return response.json();
};

export const deleteTask = async (id) => {
    await fetch(`${API_URL}/tasks/${id}`, {
        method: 'DELETE',
    });
};

export const fetchCategories = async () => {
    const response = await fetch(`${API_URL}/categories`);
    if (!response.ok) {
        throw new Error('Failed to fetch categories');
    }
    return response.json();
};