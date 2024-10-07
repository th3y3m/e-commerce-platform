"use client";

import ProductCard from '@/components/ProductCard';
import { useAppDispatch, useAppSelector } from '@/redux/hooks';
import { fetchAllProducts } from '@/redux/slice/ProductSlice';
import React, { useEffect, useState } from 'react';

const ProductPage = () => {
    const dispatch = useAppDispatch();

    const products = useAppSelector(state => state.product.products);

    // Fetch products when the component mounts
    useEffect(() => {
        dispatch(fetchAllProducts({
            pageIndex: "1",
            pageSize: "10"
        }));
    }, [dispatch]);

    return (
        <div>
            <h1>Product Page</h1>
            <p>Welcome to the Product page!</p>

            {products && products.Items && products.Items.length > 0 && products.Items.map((product) => (
                <ProductCard key={product.ProductID} product={product} />
            ))}
        </div>
    );
};

export default ProductPage;
