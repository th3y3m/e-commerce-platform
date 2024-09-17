import axios from "./customizeAxios";

const fetchProductsFromCart = async (id: string) => {
    try {
        const response = await axios.get(`auth/cartItems/${id}`);
        return response;
    } catch (error) {
        throw error;
    }
}

export {
    fetchProductsFromCart,
};