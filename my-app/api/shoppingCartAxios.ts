import axios from "./customizeAxios";

interface CartItem {
    UserID: string;
    ProductID: string;
    Quantity: number;
}

const GetShoppingCartByUserID = async (userID: string) => {
    try {
        const response = await axios.get(`auth/shoppingCart/${userID}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const GetUserShoppingCart = async (userID: string) => {
    try {
        const response = await axios.get(`auth/shoppingCart/${userID}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const AddProductToShoppingCart = async (item: CartItem) => {
    try {
        // Send data in the body of the POST request
        const response = await axios.post('auth/shoppingCart', item);

        return response;
    } catch (error) {
        throw error;
    }
}

const RemoveProductFromCart = async (item: CartItem) => {
    try {
        const response = await axios.put(`auth/shoppingCart`, item);
        return response;
    } catch (error) {
        throw error;
    }
}

const ClearShoppingCart = async (userID: string) => {
    try {
        const response = await axios.delete(`auth/shoppingCart/${userID}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const NumberOfItemsInCart = async (userID: string) => {
    try {
        const response = await axios.get(`auth/shoppingCart/numberofitems/${userID}`);
        return response;
    } catch (error) {
        throw error;
    }
}