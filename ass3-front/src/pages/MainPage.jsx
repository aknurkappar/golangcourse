import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { fetchUsers } from '../api';

const Main = () => {
    const dispatch = useDispatch();
    const profile = useSelector(state => state.profile.currentUser);
    const [users, setUsers] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        if (profile && profile.role && profile.role.name === 'admin') {
            const getUsers = async () => {
                try {
                    const fetchedData = await fetchUsers();
                    setUsers(fetchedData);
                } catch (error) {
                    setError('Failed to fetch users: ' + error.message);
                }
            };
            getUsers()
        }
    }, [dispatch, profile]);

    return (
        <div>
            {profile ? (
                <div>
                    <h1>Hello, {profile.name}!</h1>
                    <p><strong>Email:</strong> {profile.email}</p>
                    <p><strong>Role:</strong> {profile.role && profile.role.name}</p>

                    {profile.role.name === 'admin' && (
                        <div>
                            <h2>User List</h2>
                            <ul>
                                {users && users.length > 0 && users.map(user => (
                                    <li key={user.id}>
                                        {user.name} - {user.email} ({user.role && user.role.name})
                                    </li>
                                ))}
                            </ul>
                            <Link to="/add-user">Add user</Link>
                        </div>
                    )}
                </div>
            ) : (
                <h1>Welcome, please <Link to="/login">login!</Link></h1>
            )}
        </div>
    );
};

export default Main;
