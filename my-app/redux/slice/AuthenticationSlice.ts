import { login, register } from "@/api/authenticationAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

interface LoginData {
    email: string;
    password: string;
}

interface RegisternData {
    email: string;
    password: string;
    confirmPassword: string;
}

export const Login = createAsyncThunk(
    "authentication/login",
    async (data: LoginData, { rejectWithValue }) => {
        try {
            const response = await login(data.email, data.password);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const Register = createAsyncThunk(
    "authentication/register",
    async (data: RegisternData, { rejectWithValue }) => {
        try {
            const response = await register(data.email, data.password, data.confirmPassword);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

const token = localStorage.getItem('authToken');
const initialState = {
    token: token !== null ? token : "",
    status: "",
};

const AuthenticationSlice = createSlice({
    name: "authentication",
    initialState,
    reducers: {
        logout: (state) => {
            state.token = "";
            localStorage.removeItem('authToken');
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(Login.pending, (state) => {
                state.status = "loading";
            })
            .addCase(Login.fulfilled, (state, action) => {
                state.token = action.payload.token;
                state.status = "success";
                localStorage.setItem('authToken', JSON.stringify(action.payload.token));
            })
            .addCase(Login.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(Register.pending, (state) => {
                state.status = "loading";
            })
            .addCase(Register.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(Register.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export const { logout } = AuthenticationSlice.actions;
export default AuthenticationSlice.reducer;