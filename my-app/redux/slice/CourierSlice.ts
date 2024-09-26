import { createCourier, deleteCourier, getAllCouriers, getCourierById, updateCourier } from "@/api/courierAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

export const fetchAllCouriers = createAsyncThunk(
    "courier/fetchAllCouriers",
    async (_, { rejectWithValue }) => {
        try {
            const response = await getAllCouriers();
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchCourierById = createAsyncThunk(
    "courier/fetchCourierById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getCourierById(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const createNewCourier = createAsyncThunk(
    "courier/createNewCourier",
    async (courier_name: string, { rejectWithValue }) => {
        try {
            const response = await createCourier(courier_name);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const updateCourierByID = createAsyncThunk(
    "courier/updateCourierByID",
    async ({ id, courier_name }: { id: string, courier_name: string }, { rejectWithValue }) => {
        try {
            const response = await updateCourier(id, courier_name);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteCourierByID = createAsyncThunk(
    "courier/deleteCourierByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteCourier(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

interface Courier {
    CourierID: string
    Courier: string
    Status: boolean
}

const initialState = {
    couriers: [] as Courier[],
    courier: {} as Courier,
    status: "",
};

export const CourierSlice = createSlice({
    name: "courier",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAllCouriers.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchAllCouriers.fulfilled, (state, action) => {
                state.couriers = action.payload;
                state.status = "success";
            })
            .addCase(fetchAllCouriers.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(fetchCourierById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchCourierById.fulfilled, (state, action) => {
                state.courier = action.payload;
                state.status = "success";
            })
            .addCase(fetchCourierById.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(createNewCourier.pending, (state) => {
                state.status = "loading";
            })
            .addCase(createNewCourier.fulfilled, (state, action) => {
                state.courier = action.payload;
                state.couriers.push(action.payload);
                state.status = "success";
            })
            .addCase(createNewCourier.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(updateCourierByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(updateCourierByID.fulfilled, (state, action) => {
                state.courier = action.payload;
                state.couriers = state.couriers.map((courier) => {
                    if (courier.CourierID === action.payload.CourierID) {
                        return action.payload;
                    }
                    return courier;
                });
                state.status = "success";
            })
            .addCase(updateCourierByID.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(deleteCourierByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(deleteCourierByID.fulfilled, (state, action) => {
                state.courier = action.payload;
                state.couriers = state.couriers.filter((courier) => courier.CourierID !== state.courier.CourierID);
                state.status = "success";
            })
            .addCase(deleteCourierByID.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export default CourierSlice.reducer;