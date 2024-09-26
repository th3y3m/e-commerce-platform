import { configureStore } from "@reduxjs/toolkit";
import AuthenticationSlice from "./slice/AuthenticationSlice";
import CartItemSlice from "./slice/CartItemSlice";
import CourierSlice from "./slice/CourierSlice";
import DiscountSlice from "./slice/DiscountSlice";
import NewsSlice from "./slice/NewsSlice";
import FreightRateSlice from "./slice/FreightRateSlice";
import OrderDetailSlice from "./slice/OrderDetailSlice";
import OrderSlice from "./slice/OrderSlice";
import ProductDiscountSlice from "./slice/ProductDiscountSlice";
import ProductSlice from "./slice/ProductSlice";
import ReviewSlice from "./slice/ReviewSlice";
import ShoppingCartSlice from "./slice/ShoppingCartSlice";
import UserSlice from "./slice/UserSlice";
import CategorySlice from "./slice/CategorySlice";

export const makeStore = () => {
    return configureStore({
        reducer: {
            auth: AuthenticationSlice,
            cartItem: CartItemSlice,
            category: CategorySlice,
            courier: CourierSlice,
            discount: DiscountSlice,
            freightRate: FreightRateSlice,
            news: NewsSlice,
            orderDetail: OrderDetailSlice,
            order: OrderSlice,
            productDiscount: ProductDiscountSlice,
            product: ProductSlice,
            review: ReviewSlice,
            shoppingCart: ShoppingCartSlice,
            user: UserSlice,
        }
    })
}

export type AppStore = ReturnType<typeof makeStore>
export type RootState = ReturnType<AppStore['getState']>
export type AppDispatch = AppStore['dispatch']