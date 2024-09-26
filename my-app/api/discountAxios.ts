import axios from './customizeAxios';

interface Params {
    searchQuery?: string;
    sortBy?: string;
    pageIndex?: string;
    pageSize?: string;
    status?: string;
}
interface NewDiscount {
    DiscountType?: string;
    DiscountValue?: number;
    StartDate?: Date;
    EndDate?: Date;
}

const getAllDiscounts = async (params: Params) => {
    try {
        const {
            searchQuery = "",
            sortBy = "",
            pageIndex = "1",
            pageSize = "10",
            status = ""
        } = params;

        const queryParams = new URLSearchParams();

        // Conditionally append query parameters if they have values
        if (searchQuery) queryParams.append('searchQuery', searchQuery);
        if (sortBy) queryParams.append('sortBy', sortBy);
        if (status !== undefined) queryParams.append('status', status);
        if (pageIndex) queryParams.append('pageIndex', pageIndex);
        if (pageSize) queryParams.append('pageSize', pageSize);

        const response = await axios.get(`auth/discounts?${queryParams.toString()}`);
        return response.data;
    } catch (error) {
        throw error;  // Re-throw the error to handle it outside this function
    }
}

const getDiscountByID = async (id: string) => {
    try {
        const response = await axios.get(`auth/discounts/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const createDiscount = async (discountData: NewDiscount) => {
    try {
        const response = await axios.post('auth/discounts', discountData);
        return response.data;
    } catch (error) {
        throw error;
    }
}
const updateDiscount = async (id: string, discountData: NewDiscount) => {
    try {
        const response = await axios.post(`auth/discounts/${id}`, discountData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const deleteDiscount = async (id: string) => {
    try {
        const response = await axios.delete(`auth/discounts/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    getAllDiscounts,
    getDiscountByID,
    createDiscount,
    updateDiscount,
    deleteDiscount
}
