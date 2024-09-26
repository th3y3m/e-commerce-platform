import { createCategory, deleteCategory, getAllCategories, getCategoryById, updateCategory } from "@/api/categoryAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

export const fetchCategories = createAsyncThunk(
    "category/fetchCategories",
    async (_, { rejectWithValue }) => {
        try {
            const response = await getAllCategories();

            return response.categories;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchCategoryById = createAsyncThunk(
    "category/fetchCategoryById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getCategoryById(id);
            return response.category;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const createNewCategory = createAsyncThunk(
    "category/createNewCategory",
    async (category_name: string, { rejectWithValue }) => {
        try {
            const response = await createCategory(category_name);
            return response; // response is already returned by the interceptor
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const updateCategoryByID = createAsyncThunk(
    "category/updateCategoryByID",
    async ({ id, category_name }: { id: string, category_name: string }, { rejectWithValue }) => {
        try {
            const response = await updateCategory(id, category_name);
            return response; // response is already returned by the interceptor
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteCategoryByID = createAsyncThunk(
    "category/deleteCategoryByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteCategory(id);
            return response; // response is already returned by the interceptor
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

interface Category {
    CategoryID: string;
    CategoryName: string;
}

const initialState = {
    categories: [] as Category[],
    category: {} as Category,
    status: "",
};

const CategorySlice = createSlice({
    name: "category",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            // Fetch all categories
            .addCase(fetchCategories.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchCategories.fulfilled, (state, action) => {
                console.log(action.payload);
                state.categories = action.payload; // action.payload is the data directly
                state.status = "success";
            })
            .addCase(fetchCategories.rejected, (state) => {
                state.status = "failed";
            })

            // Fetch category by ID
            .addCase(fetchCategoryById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchCategoryById.fulfilled, (state, action) => {
                console.log(action.payload);
                state.category = action.payload; // action.payload is the data directly
                state.status = "success";
            })
            .addCase(fetchCategoryById.rejected, (state) => {
                state.status = "failed";
            })

            // Create new category
            .addCase(createNewCategory.pending, (state) => {
                state.status = "loading";
            })
            .addCase(createNewCategory.fulfilled, (state, action) => {
                state.category = action.payload;
                state.categories.push(action.payload);
                state.status = "success";
            })
            .addCase(createNewCategory.rejected, (state) => {
                state.status = "failed";
            })

            // Update category by ID
            .addCase(updateCategoryByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(updateCategoryByID.fulfilled, (state, action) => {
                state.category = action.payload;
                state.categories = state.categories.map((category) =>
                    category.CategoryID === action.payload.CategoryID ? action.payload : category
                );
                state.status = "success";
            })
            .addCase(updateCategoryByID.rejected, (state) => {
                state.status = "failed";
            })

            // Delete category by ID
            .addCase(deleteCategoryByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(deleteCategoryByID.fulfilled, (state, action) => {
                state.categories = state.categories.filter(
                    (category) => category.CategoryID !== action.payload.CategoryID
                );
                state.status = "success";
            })
            .addCase(deleteCategoryByID.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export default CategorySlice.reducer;