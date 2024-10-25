import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { fetchTask, updateTask, fetchCategories } from '../api';

const EditTaskPage = () => {
    const { id } = useParams();
    const [title, setTitle] = useState('');
    const [status, setStatus] = useState('todo');
    const [category_id, setCategory] = useState('');
    const [categories, setCategories] = useState([]);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const getTask = async () => {
            try {
                const task = await fetchTask(id);
                setTitle(task.title);
                setStatus(task.status);
                setCategory(task.category_id);
            } catch (error) {
                setError('Failed to fetch task: ' + error.message);
            }
        };

        const getCategories = async () => {
            try {
                const fetchedCategories = await fetchCategories();
                setCategories(fetchedCategories);
            } catch (error) {
                setError('Failed to fetch categories: ' + error.message);
            }
        };

        getTask();
        getCategories();
    }, [id]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const updatedTask = { title, status, category_id: parseInt(category_id) || null };
        try {
            await updateTask(id, updatedTask);
            navigate('/');
        } catch (error) {
            setError('Failed to update task: ' + error.message);
        }
    };

    return (
        <div>
            <h2>Edit Task</h2>
            {error && <div style={{ color: 'red' }}>{error}</div>}
            <form onSubmit={handleSubmit}>
                <label>
                    Title:
                    <input
                        type="text"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        required
                    />
                </label>
                <label>
                    Status:
                    <select value={status} onChange={(e) => setStatus(e.target.value)} required>
                        <option value="todo">Todo</option>
                        <option value="in process">In Process</option>
                        <option value="done">Done</option>
                    </select>
                </label>
                <label>
                    Category:
                    <select value={category_id} onChange={(e) => setCategory(e.target.value)} required>
                        <option value="">Select a category</option>
                        {categories.map((cat) => (
                            <option key={cat.id} value={cat.id}>
                                {cat.name}
                            </option>
                        ))}
                    </select>
                </label>
                <button type="submit">Update</button>
            </form>
        </div>
    );
};

export default EditTaskPage;
