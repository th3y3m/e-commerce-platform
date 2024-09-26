import axios from './customizeAxios';

const getAllCouriers = async () => {
    try {
        const response = await axios.get('auth/couriers');
        return response.data;
    } catch (error) {
        throw error;
    }
}

const getCourierById = async (id: string) => {
    try {
        const response = await axios.get(`auth/couriers/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const createCourier = async (courier_name: string) => {
    try {
        const response = await axios.post('auth/couriers', courier_name);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const updateCourier = async (id: string, courier_name: string) => {
    try {
        const response = await axios.put(`auth/couriers/${id}`, courier_name);
        return response.data;
    } catch (error) {
        throw error;
    }
}

const deleteCourier = async (id: string) => {
    try {
        const response = await axios.delete(`auth/couriers/${id}`);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export {
    getAllCouriers,
    getCourierById,
    createCourier,
    updateCourier,
    deleteCourier
}
