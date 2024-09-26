import axios from "./customizeAxios";

const GetOrderDetailOfAOrder = async (orderID: string) => {
    try {
        const response = await axios.get(`auth/orderDetails/${orderID}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    GetOrderDetailOfAOrder
}