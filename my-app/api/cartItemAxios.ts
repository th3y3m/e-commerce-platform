import axios from "./customizeAxios";

const fetchProductsFromCart = async (id: string) => {
    try {
        const response = await axios.get(`auth/cartItems/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    fetchProductsFromCart,
};