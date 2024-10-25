import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';

const API_URL = 'http://127.0.0.1:8080';

export const fetchProfile = createAsyncThunk('profile/fetchProfile', async () => {
    const token = localStorage.getItem('access-token');
    const response = await fetch(`${API_URL}/profile`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`
        }
    });
    if (response.ok) {
        const data = await response.json();
        return data;
    }
    throw new Error('Failed to fetch profile');
});

const profileSlice = createSlice({
    name: 'profile',
    initialState: { user: null, loading: false, error: null },
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchProfile.pending, (state) => {
                state.loading = true;
            })
            .addCase(fetchProfile.fulfilled, (state, action) => {
                state.loading = false;                
                state.currentUser = action.payload;
            })
            .addCase(fetchProfile.rejected, (state, action) => {
                state.loading = false;
                state.error = action.error.message;
            });
    },
});

export default profileSlice.reducer;
