const API_URL = 'http://127.0.0.1:8080';

export const fetchUsers = async () => {
    const token = localStorage.getItem('access-token');

    const response = await fetch(`${API_URL}/users`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`,
        },
    });
    if (!response.ok) {
        throw new Error('Failed to fetch users');
    }
    return response.json();
};

export const createUser = async (user) => {
    const token = localStorage.getItem('access-token');

    const response = await fetch(`${API_URL}/users/`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`,
        },
        body: JSON.stringify(user),
    });
    return response.json();
};

export const fetchRoles = async () => {
    const token = localStorage.getItem('access-token');

    const response = await fetch(`${API_URL}/roles`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`,
        },
    });
    if (!response.ok) {
        throw new Error('Failed to fetch roles');
    }
    return response.json();
};

export const register = async (user) => {
    const response = await fetch(`${API_URL}/register`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
    });
    const data = await response.json();
    if (data.token) {
        localStorage.setItem('access-token', data.token);
    }
    return data;
};

export const login = async (user) => {
    const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
    });
    const data = await response.json();
    if (data.token) {
        localStorage.setItem('access-token', data.token);
    }
    return data;
};

export const fetchProfile = async () => {
    const token = localStorage.getItem('access-token');

    const response = await fetch(`${API_URL}/profile/`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`,
        },
    });
    if (!response.ok) {
        throw new Error('Failed to fetch profile');
    }
    return response.json();
};