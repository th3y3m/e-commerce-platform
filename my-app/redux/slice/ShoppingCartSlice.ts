import { fetchProductsFromCart } from "@/api/cartItemAxios";
import { AddProductToShoppingCart, ClearShoppingCart, DeleteUnitItem, GetCartItems, GetShoppingCartByUserID, GetUserShoppingCart, NumberOfItemsInCart, NumberOfItemsInCartCookie, RemoveFromCart, RemoveProductFromCart } from "@/api/shoppingCartAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

interface CartItem {
    UserID: string;
    ProductID: string;
    Quantity: number;
}

interface Item {
    UserID: string;
    ProductID: string;
}

export const fetchUserShoppingCart = createAsyncThunk(
    "shoppingCart/fetchUserShoppingCart",
    async (userID: string, { rejectWithValue }) => {
        try {
            const response = await GetUserShoppingCart(userID);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const addProductToShoppingCart = createAsyncThunk(
    "shoppingCart/addProductToShoppingCart",
    async (item: CartItem, { rejectWithValue }) => {
        try {
            const response = await AddProductToShoppingCart(item);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const removeProductFromCart = createAsyncThunk(
    "shoppingCart/removeProductFromCart",
    async (item: CartItem, { rejectWithValue }) => {
        try {
            const response = await RemoveProductFromCart(item);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const clearShoppingCart = createAsyncThunk(
    "shoppingCart/clearShoppingCart",
    async (userID: string, { rejectWithValue }) => {
        try {
            const response = await ClearShoppingCart(userID);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const numberOfItemsInCart = createAsyncThunk(
    "shoppingCart/numberOfItemsInCart",
    async (userID: string, { rejectWithValue }) => {
        try {
            const response = await NumberOfItemsInCart(userID);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteUnitItem = createAsyncThunk(
    "shoppingCart/deleteUnitItem",
    async (item: Item, { rejectWithValue }) => {
        try {
            const response = await DeleteUnitItem(item);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const removeFromCart = createAsyncThunk(
    "shoppingCart/removeFromCart",
    async (item: Item, { rejectWithValue }) => {
        try {
            const response = await RemoveFromCart(item);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const getCart = createAsyncThunk(
    "shoppingCart/getCart",
    async (userID: string, { rejectWithValue }) => {
        try {
            const response = await GetCartItems(userID);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchNumberOfItemsInCartCookie = createAsyncThunk(
    "shoppingCart/fetchNumberOfItemsInCartCookie",
    async (userID: string, { rejectWithValue }) => {
        try {
            const response = await NumberOfItemsInCartCookie(userID);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);
export const fetchProductsFromShoppingCart = createAsyncThunk(
    "cartItem/fetchProductsFromCart",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await fetchProductsFromCart(id)
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

interface ShoppingCart {
    CartID: string
    UserID: string
    CreatedAt: Date
    Status: boolean
}
const initialState = {
    cart: {} as ShoppingCart,
    cartItem: [] as CartItem[],
    numberOfItemsInCart: 0,
    status: "",
};

const ShoppingCartSlice = createSlice({
    name: "shoppingCart",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchUserShoppingCart.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchUserShoppingCart.fulfilled, (state, { payload }) => {
                state.cart = payload;
                state.status = "success";
            })
            .addCase(fetchUserShoppingCart.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(fetchProductsFromShoppingCart.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchProductsFromShoppingCart.fulfilled, (state, { payload }) => {
                state.cartItem = payload;
                state.status = "success";
            })
            .addCase(fetchProductsFromShoppingCart.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(addProductToShoppingCart.pending, (state) => {
                state.status = "loading";
            })
            .addCase(addProductToShoppingCart.fulfilled, (state, { payload }) => {
                state.cartItem = payload;
                state.status = "success";
            })
            .addCase(addProductToShoppingCart.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(removeProductFromCart.pending, (state) => {
                state.status = "loading";
            })
            .addCase(removeProductFromCart.fulfilled, (state, { payload }) => {
                state.cartItem = payload;
                state.status = "success";
            })
            .addCase(removeProductFromCart.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(clearShoppingCart.pending, (state) => {
                state.status = "loading";
            })
            .addCase(clearShoppingCart.fulfilled, (state) => {
                state.status = "success";
            })
            .addCase(clearShoppingCart.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(numberOfItemsInCart.pending, (state) => {
                state.status = "loading";
            })
            .addCase(numberOfItemsInCart.fulfilled, (state, { payload }) => {
                state.numberOfItemsInCart = payload;
                state.status = "success";
            })
            .addCase(numberOfItemsInCart.rejected, (state) => {
                state.status = "failed";
            });
    }
});

export default ShoppingCartSlice.reducer;