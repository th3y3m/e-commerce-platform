import axios from "./customizeAxios";

const GetOrderDetailOfAOrder = async (orderID: string) => {
    try {
        const response = await axios.get(`auth/orderDetails/${orderID}`);
        return response;
    } catch (error) {
        throw error;
    }
}

export {
    GetOrderDetailOfAOrder
}