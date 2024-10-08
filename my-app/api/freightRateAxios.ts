import axios from './customizeAxios';

interface NewFreightRate {
    CourierID?: string;
    DistanceMinKM?: number;
    DistanceMaxKM?: number;
    CostPerKM?: number;
    Status?: boolean;
}

const getAllFreightRates = async () => {
    try {
        const response = await axios.get(`auth/freightRates`);
        return response.data;
    } catch (error) {
        throw error;  // Re-throw the error to handle it outside this function
    }
}

const getFreightRateByID = async (id: string) => {
    try {
        const response = await axios.get(`auth/freightRates/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const createFreightRate = async (freightRateData: NewFreightRate) => {
    try {
        const response = await axios.post('auth/freightRates', freightRateData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const updateFreightRate = async (id: string, freightRateData: NewFreightRate) => {
    try {
        const response = await axios.post(`auth/freightRates/${id}`, freightRateData);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const deleteFreightRate = async (id: string) => {
    try {
        const response = await axios.delete(`auth/freightRates/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}


export {
    getAllFreightRates,
    getFreightRateByID,
    createFreightRate,
    updateFreightRate,
    deleteFreightRate
}