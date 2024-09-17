import axios from "./customizeAxios";

const fetchAllUsers = async () => {
    try {
        const response = await axios.get("auth/users");
        return response;
    } catch (error) {
        throw error;
    }
}

const fetchUserById = async (id: string) => {
    try {
        const response = await axios.get(`auth/users/${id}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const updateProfile = async (id: string, user: { fullName: string; phoneNumber: string; address: string }) => {
    try {
        const response = await axios.put(`auth/users/${id}`, user);
        return response;
    } catch (error) {
        throw error;
    }
};

const banUser = async (id: string) => {
    try {
        const response = await axios.put(`auth/users/Ban/${id}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const unbanUser = async (id: string) => {
    try {
        const response = await axios.put(`auth/users/Unban/${id}`);
        return response;
    } catch (error) {
        throw error;
    }
}

export {
    fetchAllUsers,
    fetchUserById,
    updateProfile,
    banUser,
    unbanUser,
};
