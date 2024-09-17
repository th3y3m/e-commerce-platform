import axios from "./customizeAxios";

interface Params {
    sortBy?: string;
    reviewID?: string;
    userID?: string;
    productID?: string;
    pageIndex?: string;
    pageSize?: string;
    minRating?: string;
    maxRating?: string;
    status?: string;
}

interface NewReview {
    UserID: string;
    ProductID: string;
    Rating: number;
    Comment: string;
}

interface UpdateReview {
    Rating: number;
    Comment: string;
}

const fetchAllReviews = async (params: Params) => {
    try {
        const {
            sortBy = "",
            reviewID = "",
            userID = "",
            productID = "",
            pageIndex = "1",
            pageSize = "10",
            minRating = "",
            maxRating = "",
            status = ""
        } = params;

        const queryParams = new URLSearchParams();

        // Conditionally append query parameters if they have values
        if (sortBy) queryParams.append('sortBy', sortBy);
        if (reviewID) queryParams.append('reviewID', reviewID);
        if (userID) queryParams.append('userID', userID);
        if (productID) queryParams.append('productID', productID);
        if (status) queryParams.append('status', status);
        if (pageIndex) queryParams.append('pageIndex', pageIndex);
        if (pageSize) queryParams.append('pageSize', pageSize);
        if (minRating) queryParams.append('minRating', minRating);
        if (maxRating) queryParams.append('maxRating', maxRating);

        const response = await axios.get(`auth/reviews?${queryParams.toString()}`);
        return response;
    } catch (error) {
        throw error;  // Re-throw the error to handle it outside this function
    }
}

const getReviewByID = async (id: string) => {
    try {
        const response = await axios.get(`auth/reviews/${id}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const createReview = async (reviewData: NewReview) => {
    try {
        const response = await axios.post('auth/reviews', reviewData);
        return response;
    } catch (error) {
        throw error;
    }
}

const updateReview = async (id: string, reviewData: UpdateReview) => {
    try {
        const response = await axios.post(`auth/reviews/${id}`, reviewData);
        return response;
    } catch (error) {
        throw error;
    }
}

const deleteReview = async (id: string) => {
    try {
        const response = await axios.delete(`auth/reviews/${id}`);
        return response;
    } catch (error) {
        throw error;
    }
}

export {
    fetchAllReviews,
    getReviewByID,
    createReview,
    updateReview,
    deleteReview
}