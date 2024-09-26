import { createFreightRate, deleteFreightRate, getAllFreightRates, getFreightRateByID, updateFreightRate } from "@/api/freightRateAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

interface NewFreightRate {
    CourierID?: string;
    DistanceMinKM?: number;
    DistanceMaxKM?: number;
    CostPerKM?: number;
    Status?: boolean;
}

export const fetchAllFreightRates = createAsyncThunk(
    "freightRate/fetchAllFreightRates",
    async (_, { rejectWithValue }) => {
        try {
            const response = await getAllFreightRates();
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchFreightRateById = createAsyncThunk(
    "freightRate/fetchFreightRateById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getFreightRateByID(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const createNewFreightRate = createAsyncThunk(
    "freightRate/createNewFreightRate",
    async (freightRateData: NewFreightRate, { rejectWithValue }) => {
        try {
            const response = await createFreightRate(freightRateData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const updateFreightRateByID = createAsyncThunk(
    "freightRate/updateFreightRateByID",
    async ({ id, freightRateData }: { id: string, freightRateData: NewFreightRate }, { rejectWithValue }) => {
        try {
            const response = await updateFreightRate(id, freightRateData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteFreightRateByID = createAsyncThunk(
    "freightRate/deleteFreightRateByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteFreightRate(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

interface FreightRate {
    RateID: string;
    CourierID: string;
    DistanceMinKM: number;
    DistanceMaxKM: number;
    CostPerKM: number;
    Status: boolean;
}

const initialState = {
    freightRates: [] as FreightRate[],
    freightRate: {} as FreightRate,
    status: "",
};

const FreightRateSlice = createSlice({
    name: "freightRate",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAllFreightRates.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchAllFreightRates.fulfilled, (state, { payload }) => {
                state.freightRates = payload;
                state.status = "success";
            })
            .addCase(fetchAllFreightRates.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(fetchFreightRateById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchFreightRateById.fulfilled, (state, { payload }) => {
                state.freightRate = payload;
                state.status = "success";
            })
            .addCase(fetchFreightRateById.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(createNewFreightRate.pending, (state) => {
                state.status = "loading";
            })
            .addCase(createNewFreightRate.fulfilled, (state, { payload }) => {
                state.freightRate = payload;
                state.freightRates.push(payload);
                state.status = "success";
            })
            .addCase(createNewFreightRate.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(updateFreightRateByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(updateFreightRateByID.fulfilled, (state, { payload }) => {
                state.freightRate = payload;
                state.freightRates = state.freightRates.map((freightRate) => {
                    if (freightRate.CourierID === payload.CourierID) {
                        return payload;
                    }
                    return freightRate;
                });
                state.status = "success";
            })
            .addCase(updateFreightRateByID.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(deleteFreightRateByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(deleteFreightRateByID.fulfilled, (state, { payload }) => {
                state.freightRate = payload;
                state.freightRates = state.freightRates.filter((freightRate) => freightRate.CourierID !== payload.CourierID);
                state.status = "success";
            })
            .addCase(deleteFreightRateByID.rejected, (state) => {
                state.status = "failed";
            });
    }
});

export default FreightRateSlice.reducer;