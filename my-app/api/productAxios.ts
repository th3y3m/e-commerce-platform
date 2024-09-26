import axios from "./customizeAxios";

interface Params {
    searchValue?: string;
    sortBy?: string;
    productID?: string;
    sellerID?: string;
    categoryID?: string;
    pageIndex?: string;
    pageSize?: string;
    status?: string;
}

interface NewProduct {
    SellerID: string;
    ProductName: string;
    Description: string;
    CategoryID: string;
    Price: number;
    Stock: number;
    ImageURL: string;
}

const getAllProducts = async (params: Params) => {
    try {
        const {
            searchValue = "",
            sortBy = "",
            productID = "",
            sellerID = "",
            categoryID = "",
            pageIndex = "1",
            pageSize = "10",
            status = "",
        } = params;

        const queryParams = new URLSearchParams();

        // Conditionally append query parameters if they have values
        if (searchValue) queryParams.append("searchValue", searchValue);
        if (sortBy) queryParams.append("sortBy", sortBy);
        if (productID) queryParams.append("productID", productID);
        if (sellerID) queryParams.append("sellerID", sellerID);
        if (categoryID) queryParams.append("categoryID", categoryID);
        if (status) queryParams.append("status", status);
        if (pageIndex) queryParams.append("pageIndex", pageIndex);
        if (pageSize) queryParams.append("pageSize", pageSize);

        const response = await axios.get(`/products?${queryParams.toString()}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const getProductByID = async (id: string) => {
    try {
        const response = await axios.get(`/product/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const createProduct = async (productData: NewProduct) => {
    try {
        const response = await axios.post("auth/product", productData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const updateProduct = async (id: string, productData: NewProduct) => {
    try {
        const response = await axios.put(`auth/product/${id}`, productData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const deleteProduct = async (id: string) => {
    try {
        const response = await axios.delete(`auth/product/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const GetProductPriceAfterDiscount = async (id: string) => {
    try {
        const response = await axios.get(`/product/${id}/price`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    getAllProducts,
    getProductByID,
    createProduct,
    updateProduct,
    deleteProduct,
    GetProductPriceAfterDiscount
};