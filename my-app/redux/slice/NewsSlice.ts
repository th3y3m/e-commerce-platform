import { createNews, deleteNews, getAllNews, getNewsByID, updateNews } from "@/api/newsAxios";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

interface Params {
    searchQuery?: string;
    sortBy?: string;
    pageIndex?: string;
    pageSize?: string;
    newID?: string;
    authorID?: string;
    status?: string;
}

interface NewNews {
    title: string;
    content: string;
    authorID: string;
    Category: string;
    ImageURL: string;
}

export const fetchAllNews = createAsyncThunk(
    "news/fetchAllNews",
    async (params: Params, { rejectWithValue }) => {
        try {
            const response = await getAllNews(params);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const fetchNewsById = createAsyncThunk(
    "news/fetchNewsById",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await getNewsByID(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const createNewNews = createAsyncThunk(
    "news/createNewNews",
    async (newsData: NewNews, { rejectWithValue }) => {
        try {
            const response = await createNews(newsData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const updateNewsByID = createAsyncThunk(
    "news/updateNewsByID",
    async ({ id, newsData }: { id: string, newsData: NewNews }, { rejectWithValue }) => {
        try {
            const response = await updateNews(id, newsData);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

export const deleteNewsByID = createAsyncThunk(
    "news/deleteNewsByID",
    async (id: string, { rejectWithValue }) => {
        try {
            const response = await deleteNews(id);
            return response;
        } catch (error) {
            return rejectWithValue(error);
        }
    }
);

interface News {
    NewsID: string
    Title: string
    Content: string
    PublishedDate: Date
    AuthorID: string
    Status: boolean
    ImageURL: string
    Category: string
}

interface PaginatedList<T> {
    Items: T[];
    TotalCount: number;
    PageIndex: number;
    PageSize: number;
    TotalPages: number;
}

const initialState = {
    news: {} as PaginatedList<News>,
    newsDetail: {} as News,
    status: "",
};

const NewsSlice = createSlice({
    name: "news",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchAllNews.pending, (state, action) => {
                state.status = "loading";
            })
            .addCase(fetchAllNews.fulfilled, (state, action) => {
                state.news = action.payload;
                state.status = "success";
            })
            .addCase(fetchAllNews.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(fetchNewsById.pending, (state, action) => {
                state.status = "loading";
            })
            .addCase(fetchNewsById.fulfilled, (state, action) => {
                state.newsDetail = action.payload;
                state.status = "success";
            })
            .addCase(fetchNewsById.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(createNewNews.pending, (state, action) => {
                state.status = "loading";
            })
            .addCase(createNewNews.fulfilled, (state, action) => {
                state.newsDetail = action.payload;
                state.status = "success";
            })
            .addCase(createNewNews.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(updateNewsByID.pending, (state, action) => {
                state.status = "loading";
            })
            .addCase(updateNewsByID.fulfilled, (state, action) => {
                state.newsDetail = action.payload;
                state.status = "success";
            })
            .addCase(updateNewsByID.rejected, (state, action) => {
                state.status = "failed";
            })
            .addCase(deleteNewsByID.pending, (state, action) => {
                state.status = "loading";
            })
            .addCase(deleteNewsByID.fulfilled, (state, action) => {
                state.status = "success";
            })
            .addCase(deleteNewsByID.rejected, (state, action) => {
                state.status = "failed";
            });
    },
});


export default NewsSlice.reducer;