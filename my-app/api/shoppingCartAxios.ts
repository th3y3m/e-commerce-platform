import axios from "./customizeAxios";

interface CartItem {
    UserID: string;
    ProductID: string;
    Quantity: number;
}

interface Item {
    UserID: string;
    ProductID: string;
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
        const response = await axios.post('auth/shoppingCart', item);
        return response;
    } catch (error) {
        throw error;
    }
}

const RemoveProductFromCart = async (item: CartItem) => {
    try {
        const response = await axios.put('auth/shoppingCart', item);
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

const DeleteUnitItem = async (item: Item) => {
    try {
        const response = await axios.put(`cookie/deleteUnitItem`, { item });
        return response;
    } catch (error) {
        throw error;
    }
}

const RemoveFromCart = async (item: Item) => {
    try {
        const response = await axios.put(`cookie/removeFromCart`, item);
        return response;
    } catch (error) {
        throw error;
    }
}

const GetCartItems = async (userID: string) => {
    try {
        const response = await axios.get(`cookie/getCartItems/${userID}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const DeleteCartInCookie = async (userID: string) => {
    try {
        const response = await axios.delete(`cookie/deleteCartInCookie/${userID}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const NumberOfItemsInCartCookie = async (userID: string) => {
    try {
        const response = await axios.get(`cookie/numberOfItemsInCartCookie/${userID}`);
        return response;
    } catch (error) {
        throw error;
    }
}

const SaveCartToCookieHandler = async (item: Item) => {

    try {
        const response = await axios.post(`cookie/saveCartToCookieHandler`, item);
        return response;
    } catch (error) {
        throw error;
    }
}

export {
    GetShoppingCartByUserID,
    GetUserShoppingCart,
    AddProductToShoppingCart,
    RemoveProductFromCart,
    ClearShoppingCart,
    NumberOfItemsInCart,
    DeleteUnitItem,
    RemoveFromCart,
    GetCartItems,
    DeleteCartInCookie,
    NumberOfItemsInCartCookie,
    SaveCartToCookieHandler
};