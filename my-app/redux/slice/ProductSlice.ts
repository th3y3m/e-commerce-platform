import { createProduct, deleteProduct, getAllProducts, getProductByID, updateProduct } from "@/api/productAxios";
import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";

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

interface PaginatedList<T> {
    Items: T[];
    TotalCount: number;
    PageIndex: number;
    PageSize: number;
    TotalPages: number;
}

interface Product {
    ProductID: string;
    SellerID: string;
    ProductName: string;
    Description: string;
    Price: number;
    Quantity: number;
    CategoryID: string;
    ImageURL: string;
    CreatedAt: Date;
    UpdatedAt: Date;
    Status: boolean;
}

interface ProductState {
    products: PaginatedList<Product>;
    product: Product | null;
    status: string;
    error: string | null;
}

const initialState: ProductState = {
    products: { Items: [], TotalCount: 0, PageIndex: 1, PageSize: 10, TotalPages: 0 },
    product: null,
    status: "",
    error: null,
};

// Async Thunks
export const fetchAllProducts = createAsyncThunk(
    "product/fetchAllProducts",
    async (params: Params, { rejectWithValue }) => {
        try {
            const response = await getAllProducts(params);
            return response;
        } catch (error: unknown) {
            return rejectWithValue((error as Error).message);
        }
    }
);

export const fetchProductById = createAsyncThunk(
    "product/fetchProductById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getProductByID(id);
            return response;
        } catch (error: unknown) {
            return rejectWithValue((error as Error).message);
        }
    }
);

export const createNewProduct = createAsyncThunk(
    "product/createNewProduct",
    async (productData: NewProduct, { rejectWithValue }) => {
        try {
            const response = await createProduct(productData);
            return response;
        } catch (error: unknown) {
            return rejectWithValue((error as Error).message);
        }
    }
);

export const updateProductByID = createAsyncThunk(
    "product/updateProductByID",
    async ({ id, productData }: { id: string; productData: NewProduct }, { rejectWithValue }) => {
        try {
            const response = await updateProduct(id, productData);
            return response;
        } catch (error: unknown) {
            return rejectWithValue((error as Error).message);
        }
    }
);

export const deleteProductByID = createAsyncThunk(
    "product/deleteProductByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteProduct(id);
            return response;
        } catch (error: unknown) {
            return rejectWithValue((error as Error).message);
        }
    }
);

export const updateProductQuantity = createAsyncThunk(
    "product/updateProductQuantity",
    async ({ productId, quantity }: { productId: string, quantity: number }, { rejectWithValue }) => {
        try {
            return { productId, quantity };
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

// Product Slice
const ProductSlice = createSlice({
    name: "product",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(updateProductQuantity.fulfilled, (state, action) => {
                const { productId, quantity } = action.payload;
                const product = state.products.Items.find(item => item.ProductID === productId);
                if (product) {
                    product.Quantity = quantity;
                }
            })
            // Fetch All Products
            .addCase(fetchAllProducts.pending, (state) => {
                state.status = "loading";
                state.error = null;
            })
            .addCase(fetchAllProducts.fulfilled, (state, action: PayloadAction<PaginatedList<Product>>) => {
                state.products = action.payload;
                state.status = "success";
            })
            .addCase(fetchAllProducts.rejected, (state, action) => {
                state.status = "failed";
                state.error = action.payload as string;
            })
            // Fetch Product by ID
            .addCase(fetchProductById.pending, (state) => {
                state.status = "loading";
                state.error = null;
            })
            .addCase(fetchProductById.fulfilled, (state, action: PayloadAction<Product>) => {
                state.product = action.payload;
                state.status = "success";
            })
            .addCase(fetchProductById.rejected, (state, action) => {
                state.status = "failed";
                state.error = action.payload as string;
            })
            // Create Product
            .addCase(createNewProduct.pending, (state) => {
                state.status = "loading";
                state.error = null;
            })
            .addCase(createNewProduct.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(createNewProduct.rejected, (state, action) => {
                state.status = "failed";
                state.error = action.payload as string;
            })
            // Update Product
            .addCase(updateProductByID.pending, (state) => {
                state.status = "loading";
                state.error = null;
            })
            .addCase(updateProductByID.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(updateProductByID.rejected, (state, action) => {
                state.status = "failed";
                state.error = action.payload as string;
            })
            // Delete Product
            .addCase(deleteProductByID.pending, (state) => {
                state.status = "loading";
                state.error = null;
            })
            .addCase(deleteProductByID.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(deleteProductByID.rejected, (state, action) => {
                state.status = "failed";
                state.error = action.payload as string;
            });
    },
});

export default ProductSlice.reducer;
