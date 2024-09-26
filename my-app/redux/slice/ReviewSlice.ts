import { createReview, deleteReview, getAllReviews, getReviewByID, updateReview } from "@/api/reviewAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { create } from "domain";

interface Params {
    sortBy?: string;
    reviewID?: string;
    userID?: string;
    productID?: string;
    pageIndex?: string;
    pageSize?: string;
    minRating?: string;
    maxRating?: string;
    status?: string;
}

interface NewReview {
    UserID: string;
    ProductID: string;
    Rating: number;
    Comment: string;
}

interface UpdateReview {
    Rating: number;
    Comment: string;
}

export const fetchAllReviews = createAsyncThunk(
    "review/fetchAllReviews",
    async (params: Params, { rejectWithValue }) => {
        try {
            const response = await getAllReviews(params);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchReviewById = createAsyncThunk(
    "review/fetchReviewById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getReviewByID(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const createNewReview = createAsyncThunk(
    "review/createNewReview",
    async (reviewData: NewReview, { rejectWithValue }) => {
        try {
            const response = await createReview(reviewData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const updateReviewByID = createAsyncThunk(
    "review/updateReviewByID",
    async ({ id, reviewData }: { id: string, reviewData: UpdateReview }, { rejectWithValue }) => {
        try {
            const response = await updateReview(id, reviewData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteReviewByID = createAsyncThunk(
    "review/deleteReviewByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteReview(id);
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

interface Review {
    ReviewID: string
    ProductID: string
    UserID: string
    Rating: number
    Comment: string
    CreatedAt: Date
    Status: boolean
}

const initialState = {
    reviews: {} as PaginatedList<Review>,
    review: {} as Review,
    status: "",
}

const ReviewSlice = createSlice({
    name: "review",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAllReviews.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchAllReviews.fulfilled, (state, action) => {
                state.reviews = action.payload;
                state.status = "success";
            })
            .addCase(fetchAllReviews.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(fetchReviewById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchReviewById.fulfilled, (state, action) => {
                state.review = action.payload;
                state.status = "success";
            })
            .addCase(fetchReviewById.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(createNewReview.pending, (state) => {
                state.status = "loading";
            })
            .addCase(createNewReview.fulfilled, (state, action) => {
                state.review = action.payload;
                state.status = "success";
            })
            .addCase(createNewReview.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(updateReviewByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(updateReviewByID.fulfilled, (state, action) => {
                state.review = action.payload;
                state.status = "success";
            })
            .addCase(updateReviewByID.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(deleteReviewByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(deleteReviewByID.fulfilled, (state, action) => {
                state.review = action.payload;
                state.status = "success";
            })
            .addCase(deleteReviewByID.rejected, (state, action) => {
                state.status = "failed";
            });
    }
});

export default ReviewSlice.reducer;