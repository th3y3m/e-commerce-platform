import instance from "./customizeAxios";

const getAllCategories = async () => {
    try {
        const response = await instance.get("categories");
        return response.data;
    } catch (error) {
        throw error;
    }
}

const getCategoryById = async (id: string) => {
    try {
        const response = await instance.get(`categories/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const createCategory = async (category_name: string) => {
    try {
        const response = await instance.post("auth/categories", category_name);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const updateCategory = async (id: string, category_name: string) => {
    try {
        const response = await instance.put(`auth/categories/${id}`, category_name);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const deleteCategory = async (id: string) => {
    try {
        const response = await instance.delete(`auth/categories/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    getAllCategories,
    getCategoryById,
    createCategory,
    updateCategory,
    deleteCategory
}