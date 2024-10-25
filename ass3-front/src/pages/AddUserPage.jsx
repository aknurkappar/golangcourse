import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { createUser, fetchRoles } from '../api';

const AddUserPage = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [role_id, setRoleId] = useState(null);
    const [roles, setRoles] = useState([]);
    const [error, setError] = useState(null);
    const [notFound, setNotFound] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        const getRoles = async () => {
            try {
                const fetchedRoles = await fetchRoles();
                setRoles(fetchedRoles);
                setNotFound(false);
            } catch (error) {
                if (error.response && error.response.status === 403) {
                    setNotFound(true);
                } else {
                    setError('Failed to fetch roles: ' + error.message);
                }
            }
        };

        getRoles();
    }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const newUser = { name, email, password, role_id: parseInt(role_id) || null };
            await createUser(newUser);
            setName('');
            setEmail('');
            setPassword('');
            setRoleId(null);
            navigate('/');
        } catch (error) {
            setError('Failed to add user: ' + error.message);
        }
    };

    return (
        <div>
            {notFound ? (
                <div style={{ color: 'red' }}>Page not found: You do not have permission to access this page.</div>
            ) : (
                <>
                    <h2>Add User</h2>
                    {error && <div style={{ color: 'red' }}>{error}</div>}
                    <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '8px' }}>
                        <label>
                            Name: 
                            <input
                                type="text"
                                value={name}
                                onChange={(e) => setName(e.target.value)}
                                required
                            />
                        </label>
                        <label>
                            Email: 
                            <input
                                type="text"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                required
                            />
                        </label>
                        <label>
                            Password: 
                            <input
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                required
                            />
                        </label>
                        <label>
                            Role: 
                            <select value={role_id} onChange={(e) => setRoleId(e.target.value)}>
                                <option value="">Select a role</option>
                                {roles.map((role) => (
                                    <option key={role.id} value={role.id}>
                                        {role.name}
                                    </option>
                                ))}
                            </select>
                        </label>
                        <button style={{ width: '50px' }} type="submit">Add</button>
                    </form>
                </>
            )}
        </div>
    );
};

export default AddUserPage;
