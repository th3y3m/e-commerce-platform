import axios from "./customizeAxios";

interface Params {
    searchValue?: string;
    sortBy?: string;
    pageIndex?: string;
    pageSize?: string;
    status?: string;
}

const getAllUsers = async (params: Params) => {
    try {
        const {
            sortBy = "",
            pageIndex = "1",
            pageSize = "10",
            status = ""
        } = params;
        const queryParams = new URLSearchParams();
        // Conditionally append query parameters if they have values
        if (sortBy) queryParams.append('sortBy', sortBy);
        if (status) queryParams.append('status', status);
        if (pageIndex) queryParams.append('pageIndex', pageIndex);
        if (pageSize) queryParams.append('pageSize', pageSize);

        const response = await axios.get(`auth/users?${queryParams.toString()}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const getUserById = async (id: string) => {
    try {
        const response = await axios.get(`auth/users/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const updateProfile = async (id: string, user: { fullName: string; phoneNumber: string; address: string }) => {
    try {
        const response = await axios.put(`auth/users/${id}`, user);
        return response.data;
    } catch (error) {
        throw error;
    }
};

const banUser = async (id: string) => {
    try {
        const response = await axios.put(`auth/users/Ban/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const unbanUser = async (id: string) => {
    try {
        const response = await axios.put(`auth/users/Unban/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    getAllUsers,
    getUserById,
    updateProfile,
    banUser,
    unbanUser,
};
