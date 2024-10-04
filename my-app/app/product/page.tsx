// app/Product/page.tsx
"use client";

import { useAppDispatch, useAppSelector } from '@/redux/hooks';
import { fetchCategories, fetchCategoryById } from '@/redux/slice/CategorySlice';
import React, { useEffect, useState } from 'react';

const ProductPage = () => {
    const dispatch = useAppDispatch();

    const categories = useAppSelector(state => state.category.categories);
    const category = useAppSelector(state => state.category.category);

    useEffect(() => {
        dispatch(fetchCategories());
        dispatch(fetchCategoryById("CAT9e1f6ee0g1"));

    }, [dispatch]);
    const [n, setN] = useState(60);

    // countdown n
    setTimeout(() => {
        setN(n - 1);
    }, 1000);

    return (
        <div>
            <h1>Product Page</h1>
            <p>Welcome to the Product page!</p>
            {categories && categories.map((category) => (
                <div key={category.CategoryID}>
                    <h3>{category.CategoryName}</h3>
                </div>
            ))}
            {n}
        </div>
    );
};

export default ProductPage;
