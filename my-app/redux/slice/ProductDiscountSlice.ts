import { GetProductDiscountByID } from "@/api/productDiscountAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

export const fetchProductDiscountByID = createAsyncThunk(
    "productDiscount/fetchProductDiscountByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await GetProductDiscountByID(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

const initialState = {
    productDiscount: [],
    status: "",
};

const ProductDiscountSlice = createSlice({
    name: "productDiscount",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchProductDiscountByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchProductDiscountByID.fulfilled, (state, { payload }) => {
                state.productDiscount = payload;
                state.status = "success";
            })
            .addCase(fetchProductDiscountByID.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export default ProductDiscountSlice.reducer;