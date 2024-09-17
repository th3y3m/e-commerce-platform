import axios from "./customizeAxios";

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