import { createProduct, deleteProduct, getAllProducts, getProductByID, updateProduct } from "@/api/productAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

interface Params {
    searchValue?: string;
    sortBy?: string;
    productID?: string;
    sellerID?: string;
    categoryID?: string;
    pageIndex?: string;
    pageSize?: string;
    status?: string;
}

interface NewProduct {
    SellerID: string;
    ProductName: string;
    Description: string;
    CategoryID: string;
    Price: number;
    Stock: number;
    ImageURL: string;
}

export const fetchAllProducts = createAsyncThunk(
    "product/fetchAllProducts",
    async (params: Params, { rejectWithValue }) => {
        try {
            const response = await getAllProducts(params);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchProductById = createAsyncThunk(
    "product/fetchProductById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getProductByID(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const createNewProduct = createAsyncThunk(
    "product/createNewProduct",
    async (productData: NewProduct, { rejectWithValue }) => {
        try {
            const response = await createProduct(productData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const updateProductByID = createAsyncThunk(
    "product/updateProductByID",
    async ({ id, productData }: { id: string, productData: NewProduct }, { rejectWithValue }) => {
        try {
            const response = await updateProduct(id, productData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteProductByID = createAsyncThunk(
    "product/deleteProductByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteProduct(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const GetProductPriceAfterDiscount = createAsyncThunk(
    "product/GetProductPriceAfterDiscount",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getProductByID(id);
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
interface Product {
    ProductID: string
    SellerID: string
    ProductName: string
    Description: string
    Price: number
    Quantity: number
    CategoryID: string
    ImageURL: string
    CreatedAt: Date
    UpdatedAt: Date
    Status: boolean

}
const initialState = {
    products: {} as PaginatedList<Product>,
    product: {} as Product,
    status: "",
};

const ProductSlice = createSlice({
    name: "product",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAllProducts.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchAllProducts.fulfilled, (state, action) => {
                state.products = action.payload;
                state.status = "success";
            })
            .addCase(fetchAllProducts.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(fetchProductById.pending, (state) => {
                state.status = "loading";
            })
            .addCase(fetchProductById.fulfilled, (state, action) => {
                state.product = action.payload;
                state.status = "success";
            })
            .addCase(fetchProductById.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(createNewProduct.pending, (state) => {
                state.status = "loading";
            })
            .addCase(createNewProduct.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(createNewProduct.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(updateProductByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(updateProductByID.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(updateProductByID.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(deleteProductByID.pending, (state) => {
                state.status = "loading";
            })
            .addCase(deleteProductByID.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(deleteProductByID.rejected, (state) => {
                state.status = "failed";
            })
            .addCase(GetProductPriceAfterDiscount.pending, (state) => {
                state.status = "loading";
            })
            .addCase(GetProductPriceAfterDiscount.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(GetProductPriceAfterDiscount.rejected, (state) => {
                state.status = "failed";
            });
    },
});

export default ProductSlice.reducer;