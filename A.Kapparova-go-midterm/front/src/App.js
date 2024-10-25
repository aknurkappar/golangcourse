import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import MainPage from './pages/MainPage';
import AddTaskPage from './pages/AddTaskPage';
import EditTaskPage from './pages/EditTaskPage';

const App = () => {
    return (
      <Router>
          <div>
              <Routes>
                  <Route path="/" element={<MainPage />} />
                  <Route path="/add" element={<AddTaskPage />} />
                  <Route path="/edit/:id" element={<EditTaskPage />} />
              </Routes>
          </div>
      </Router>
    );
};

export default App;