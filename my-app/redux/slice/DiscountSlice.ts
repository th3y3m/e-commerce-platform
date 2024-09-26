import { createDiscount, deleteDiscount, getAllDiscounts, getDiscountByID, updateDiscount } from "@/api/discountAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

interface NewDiscount {
    DiscountType?: string;
    DiscountValue?: number;
    StartDate?: Date;
    EndDate?: Date;
}

interface Params {
    searchQuery?: string;
    sortBy?: string;
    pageIndex?: string;
    pageSize?: string;
    status?: string;
}

export const fetchDiscounts = createAsyncThunk(
    "courier/fetchDiscounts",
    async (params: Params, { rejectWithValue }) => {
        try {
            const response = await getAllDiscounts(params);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchDiscountById = createAsyncThunk(
    "courier/fetchDiscountById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getDiscountByID(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const createNewDiscount = createAsyncThunk(
    "courier/createNewDiscount",
    async (discountData: NewDiscount, { rejectWithValue }) => {
        try {
            const response = await createDiscount(discountData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const updateDiscountByID = createAsyncThunk(
    "courier/updateDiscountByID",
    async ({ id, discountData }: { id: string, discountData: NewDiscount }, { rejectWithValue }) => {
        try {
            const response = await updateDiscount(id, discountData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteDiscountByID = createAsyncThunk(
    "courier/deleteDiscountByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteDiscount(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

const initialState = {
    discounts: {},
    discount: {},
    status: "",
};

const DiscountSlice = createSlice({
    name: "discount",
    initialState,
    reducers: {
        clearDiscounts: (state) => {
            state.discounts = [];
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchDiscounts.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchDiscounts.fulfilled, (state, action) => {
                state.discounts = action.payload;
                state.status = "success";
            })
            .addCase(fetchDiscounts.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(fetchDiscountById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchDiscountById.fulfilled, (state, action) => {
                state.discount = action.payload;
                state.status = "success";
            })
            .addCase(fetchDiscountById.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(createNewDiscount.pending, (state) => {
                state.status = "loading";
            })
            .addCase(createNewDiscount.fulfilled, (state, action) => {
                state.discount = action.payload;
                state.status = "success";
            })
            .addCase(createNewDiscount.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(updateDiscountByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(updateDiscountByID.fulfilled, (state, action) => {
                state.discount = action.payload;
                state.status = "success";
            })
            .addCase(updateDiscountByID.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(deleteDiscountByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(deleteDiscountByID.fulfilled, (state, action) => {
                state.discount = action.payload;
                state.status = "success";
            })
            .addCase(deleteDiscountByID.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export default DiscountSlice.reducer;