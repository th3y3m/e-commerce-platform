import { GetOrderDetailOfAOrder } from "@/api/orderDetailAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

export const fetchOrderDetailByOrderID = createAsyncThunk(
    "orderDetail/fetchOrderDetailByOrderID",
    async (order_id: string, { rejectWithValue }) => {
        try {
            const response = await GetOrderDetailOfAOrder(order_id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

interface OrderDetail {
    OrderID: string;
    ProductID: string;
    Quantity: number;
    UnitPrice: number;
}

const initialState = {
    orderDetails: [] as OrderDetail[],
    status: "",
};

const OrderDetailSlice = createSlice({
    name: "orderDetail",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchOrderDetailByOrderID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchOrderDetailByOrderID.fulfilled, (state, { payload }) => {
                state.orderDetails = payload;
                state.status = "success";
            })
            .addCase(fetchOrderDetailByOrderID.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export default OrderDetailSlice.reducer;