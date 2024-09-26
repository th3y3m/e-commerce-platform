import axios from './customizeAxios';

interface Params {
    sortBy?: string;
    orderID?: string;
    customerID?: string;
    courierId?: string;
    voucherId?: string;
    pageIndex?: string;
    pageSize?: string;
    minPrice?: string;
    maxPrice?: string;
    status?: string;
    startDate?: Date;
    endDate?: Date;
}

interface NewOrder {
    userId?: string;
    CourierID?: string;
    VoucherID?: string;
    cartId?: number;
    shipAddress?: string;
}

const getAllOrders = async (params: Params) => {
    try {
        const {
            sortBy = "",
            orderID = "",
            customerID = "",
            courierId = "",
            voucherId = "",
            pageIndex = "1",
            pageSize = "10",
            minPrice = "",
            maxPrice = "",
            status = "",
            startDate,
            endDate
        } = params;

        const queryParams = new URLSearchParams();

        // Conditionally append query parameters if they have values
        if (sortBy) queryParams.append('sortBy', sortBy);
        if (orderID) queryParams.append('orderID', orderID);
        if (customerID) queryParams.append('customerID', customerID);
        if (courierId) queryParams.append('courierId', courierId);
        if (voucherId) queryParams.append('voucherId', voucherId);
        if (status) queryParams.append('status', status);
        if (pageIndex) queryParams.append('pageIndex', pageIndex);
        if (pageSize) queryParams.append('pageSize', pageSize);
        if (minPrice) queryParams.append('minPrice', minPrice);
        if (maxPrice) queryParams.append('maxPrice', maxPrice);

        // Check if startDate and endDate are valid Date objects
        if (startDate instanceof Date && !isNaN(startDate.getTime())) {
            queryParams.append('startDate', startDate.toISOString());
        }
        if (endDate instanceof Date && !isNaN(endDate.getTime())) {
            queryParams.append('endDate', endDate.toISOString());
        }

        const response = await axios.get(`auth/orders?${queryParams.toString()}`);
        return response.data;
    } catch (error) {
        throw error;  // Re-throw the error to handle it outside this function
    }
}

const getOrderById = async (id: string) => {
    try {
        const response = await axios.get(`auth/orders/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const createOrder = async (orderData: NewOrder) => {
    try {
        const response = await axios.post('auth/orders', orderData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const updateOrder = async (id: string, orderData: NewOrder) => {
    try {
        const response = await axios.put(`auth/orders/${id}`, orderData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const deleteOrder = async (id: string) => {
    try {
        const response = await axios.delete(`auth/orders/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    getAllOrders,
    getOrderById,
    createOrder,
    updateOrder,
    deleteOrder
}