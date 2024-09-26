import axios from "./customizeAxios";

interface Params {
    searchQuery?: string;
    sortBy?: string;
    pageIndex?: string;
    pageSize?: string;
    newID?: string;
    authorID?: string;
    status?: string;
}

interface NewNews {
    title: string;
    content: string;
    authorID: string;
    Category: string;
    ImageURL: string;
}

const getAllNews = async (params: Params) => {
    try {
        const {
            searchQuery = "",
            sortBy = "",
            pageIndex = "1",
            pageSize = "10",
            newID = "",
            authorID = "",
            status = ""
        } = params;

        const queryParams = new URLSearchParams();

        // Conditionally append query parameters if they have values
        if (searchQuery) queryParams.append('searchQuery', searchQuery);
        if (sortBy) queryParams.append('sortBy', sortBy);
        if (status !== undefined) queryParams.append('status', status);
        if (pageIndex) queryParams.append('pageIndex', pageIndex);
        if (pageSize) queryParams.append('pageSize', pageSize);
        if (newID) queryParams.append('newID', newID);
        if (authorID) queryParams.append('authorID', authorID);

        const response = await axios.get(`auth/news?${queryParams.toString()}`);
        return response.data;
    } catch (error) {
        throw error;  // Re-throw the error to handle it outside this function
    }
}

const getNewsByID = async (id: string) => {
    try {
        const response = await axios.get(`auth/news/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const createNews = async (newsData: NewNews) => {
    try {
        const response = await axios.post('auth/news', newsData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const updateNews = async (id: string, newsData: NewNews) => {
    try {
        const response = await axios.post(`auth/news/${id}`, newsData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const deleteNews = async (id: string) => {
    try {
        const response = await axios.delete(`auth/news/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    getAllNews,
    getNewsByID,
    createNews,
    updateNews,
    deleteNews
}