import axios from "./customizeAxios";

const fetchAllCategories = async () => {
    try {
        const response = await axios.get("auth/categories");
        return response;
    } catch (error) {
        throw error;
    }
}

const fetchCategoryById = async (id: string) => {
    try {
        const response = await axios.get(`auth/categories/${id}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const createCategory = async (category_name: string) => {
    try {
        const response = await axios.post("auth/categories", category_name);
        return response;
    } catch (error) {
        throw error;
    }
}

const updateCategory = async (id: string, category_name: string) => {
    try {
        const response = await axios.put(`auth/categories/${id}`, category_name);
        return response;
    } catch (error) {
        throw error;
    }
}

const deleteCategory = async (id: string) => {
    try {
        const response = await axios.delete(`auth/categories/${id}`);
        return response;
    } catch (error) {
        throw error;
    }
}

export {
    fetchAllCategories,
    fetchCategoryById,
    createCategory,
    updateCategory,
    deleteCategory
}