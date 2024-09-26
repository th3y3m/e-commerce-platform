import axios from "./customizeAxios";

const GetProductDiscountByID = async (id: string) => {
    try {
        const response = await axios.get(`auth/productDiscounts/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}


export {
    GetProductDiscountByID
}