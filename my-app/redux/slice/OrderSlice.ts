import { createOrder, deleteOrder, getAllOrders, getOrderById, updateOrder } from "@/api/orderAxioz";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

interface Params {
    sortBy?: string;
    orderID?: string;
    customerID?: string;
    courierId?: string;
    voucherId?: string;
    pageIndex?: string;
    pageSize?: string;
    minPrice?: string;
    maxPrice?: string;
    status?: string;
    startDate?: Date;
    endDate?: Date;
}

interface NewOrder {
    userId?: string;
    CourierID?: string;
    VoucherID?: string;
    cartId?: number;
    shipAddress?: string;
}

export const fetchAllOrders = createAsyncThunk(
    "order/fetchAllOrders",
    async (params: Params, { rejectWithValue }) => {
        try {
            const response = await getAllOrders(params);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchOrderById = createAsyncThunk(
    "order/fetchOrderById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getOrderById(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const createNewOrder = createAsyncThunk(
    "order/createNewOrder",
    async (orderData: NewOrder, { rejectWithValue }) => {
        try {
            const response = await createOrder(orderData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const updateOrderById = createAsyncThunk(
    "order/updateOrderById",
    async ({ id, orderData }: { id: string, orderData: NewOrder }, { rejectWithValue }) => {
        try {
            const response = await updateOrder(id, orderData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteOrderById = createAsyncThunk(
    "order/deleteOrderById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteOrder(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);
interface PaginatedList<T> {
    Items: T[];
    TotalCount: number;
    PageIndex: number;
    PageSize: number;
    TotalPages: number;
}
interface Order {
    OrderID: string
    CustomerID: string
    OrderDate: Date
    TotalAmount: number
    OrderStatus: string
    ShippingAddress: string
    CourierID: string
    FreightPrice: number
    EstimatedDeliveryDate: Date
    ActualDeliveryDate: Date
    PaymentMethod: string
    PaymentStatus: string
    VoucherID: string
}

const initialState = {
    orders: {} as PaginatedList<Order>,
    order: {} as Order,
    status: "",
};

const OrderSlice = createSlice({
    name: "order",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAllOrders.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchAllOrders.fulfilled, (state, action) => {
                state.status = "success";
                state.orders = action.payload;
            })
            .addCase(fetchAllOrders.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(fetchOrderById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchOrderById.fulfilled, (state, action) => {
                state.status = "success";
                state.order = action.payload;
            })
            .addCase(fetchOrderById.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(createNewOrder.pending, (state) => {
                state.status = "loading";
            })
            .addCase(createNewOrder.fulfilled, (state, action) => {
                state.order = action.payload;
                state.status = "success";
            })
            .addCase(createNewOrder.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(updateOrderById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(updateOrderById.fulfilled, (state, action) => {
                state.status = "success";
                state.order = action.payload;
            })
            .addCase(updateOrderById.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(deleteOrderById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(deleteOrderById.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(deleteOrderById.rejected, (state) => {
                state.status = "failed";
            });
    }
});

export default OrderSlice.reducer;