import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { createTask, fetchCategories } from '../api';

const AddTaskPage = () => {
    const [title, setTitle] = useState('');
    const [category_id, setCategory] = useState('');
    const [categories, setCategories] = useState([]);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const getCategories = async () => {
            try {
                const fetchedCategories = await fetchCategories();
                setCategories(fetchedCategories);
            } catch (error) {
                setError('Failed to fetch categories: ' + error.message);
            }
        };

        getCategories();
    }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const newTask = { title, category_id: parseInt(category_id) || null };
            await createTask(newTask);
            setTitle('');
            setCategory('');
            navigate('/');
        } catch (error) {
            console.log(error);
            
            setError('Failed to add task: ' + error.message);
        }
    };

    return (
        <div>
            <h2>Add Task</h2>
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
                    Category:
                    <select value={category_id} onChange={(e) => setCategory(e.target.value)}>
                        <option value="">Select a category (optional)</option>
                        {categories.map((cat) => (
                            <option key={cat.id} value={cat.id}>
                                {cat.name}
                            </option>
                        ))}
                    </select>
                </label>
                <button type="submit">Submit</button>
            </form>
        </div>
    );
};

export default AddTaskPage;
