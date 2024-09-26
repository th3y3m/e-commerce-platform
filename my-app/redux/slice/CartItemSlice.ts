import { fetchProductsFromCart } from "@/api/cartItemAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

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

interface CartItem {
    CartID: string;
    ProductID: string;
    Quantity: number;
}

const initialState = {
    cartItem: [] as CartItem[],
    status: "",
};

export const CartItemSlice = createSlice({
    name: "cartItem",
    initialState,
    reducers: {
        // clearProducts: (state) => {
        //     state.cartItem = [];
        // },
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchProductsFromShoppingCart.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchProductsFromShoppingCart.fulfilled, (state, action) => {
                state.cartItem = action.payload;
                state.status = "success";
            })
            .addCase(fetchProductsFromShoppingCart.rejected, (state) => {
                state.status = "failed";
            });
    },
});


export default CartItemSlice.reducer;