import React, { useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { fetchProfile } from './slices/profileSlice';
import MainPage from './pages/MainPage';
import AddUserPage from './pages/AddUserPage';
import RegisterPage from './pages/RegisterPage';
import LoginPage from './pages/LoginPage';

const App = () => {
    const dispatch = useDispatch();

    useEffect(() => {
        dispatch(fetchProfile());
    }, [dispatch]);

    return (
        <Router>
            <div>
                <Routes>
                    <Route path="/register" element={<RegisterPage />} />
                    <Route path="/login" element={<LoginPage />} />
                    <Route path="/" element={<MainPage />} />
                    <Route path="/add-user" element={<AddUserPage />} />
                    <Route path="*" element={<LoginPage />} />
                </Routes>
            </div>
        </Router>
    );
};

export default App;
