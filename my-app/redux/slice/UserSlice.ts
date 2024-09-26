import { banUser, getAllUsers, getUserById, unbanUser, updateProfile } from "@/api/userAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

interface Params {
    searchValue?: string;
    sortBy?: string;
    pageIndex?: string;
    pageSize?: string;
    status?: string;
}

interface Profile {
    fullName: string;
    phoneNumber: string;
    address: string
}

const fetchAllUsers = createAsyncThunk(
    'users/fetchAllUsers',
    async (params: Params, { rejectWithValue }) => {
        try {
            const response = await getAllUsers(params);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

const fetchUserById = createAsyncThunk(
    'users/fetchUserById',
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getUserById(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

const updateUserProfile = createAsyncThunk(
    'users/updateUserProfile',
    async ({ id, user }: { id: string, user: Profile }, { rejectWithValue }) => {
        try {
            const response = await updateProfile(id, user);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

const banAccount = createAsyncThunk(
    'users/banAccount',
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await banUser(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

const unbanAccount = createAsyncThunk(
    'users/unbanAccount',
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await unbanUser(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);
interface User {
    UserID: string
    Username: string
    PasswordHash: string
    Email: string
    FullName: string
    PhoneNumber: string
    Address: string
    UserType: string
    ImageURL: string
    CreatedAt: Date
    Token: string
    TokenExpires: Date
    Status: boolean
}
interface PaginatedList<T> {
    Items: T[];
    TotalCount: number;
    PageIndex: number;
    PageSize: number;
    TotalPages: number;
}
const initialState = {
    users: {} as PaginatedList<User>,
    user: {} as User,
    status: "",
};

const UserSlice = createSlice({
    name: 'users',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAllUsers.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchAllUsers.fulfilled, (state, action) => {
                state.status = "success";
                state.users = action.payload;
            })
            .addCase(fetchAllUsers.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(fetchUserById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchUserById.fulfilled, (state, action) => {
                state.status = "success";
                state.user = action.payload;
            })
            .addCase(fetchUserById.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(updateUserProfile.pending, (state) => {
                state.status = "loading";
            })
            .addCase(updateUserProfile.fulfilled, (state) => {
                state.status = "success";
            })
            .addCase(updateUserProfile.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(banAccount.pending, (state) => {
                state.status = "loading";
            })
            .addCase(banAccount.fulfilled, (state) => {
                state.status = "success";
            })
            .addCase(banAccount.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(unbanAccount.pending, (state) => {
                state.status = "loading";
            })
            .addCase(unbanAccount.fulfilled, (state) => {
                state.status = "success";
            })
            .addCase(unbanAccount.rejected, (state) => {
                state.status = "failed";
            });
    }
});

export default UserSlice.reducer;
