import axios from "./customizeAxios";

const login = async (email: string, password: string) => {
    try {
        const response = await axios.post("login", { email, password });
        return response;
    } catch (error) {
        throw error;
    }
}

const register = async (email: string, password: string, confirmPassword: string) => {
    try {
        const response = await axios.post("register", { email, password, confirmPassword });
        return response;
    } catch (error) {
        throw error;
    }
}

export {
    login,
    register,
};