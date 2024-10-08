"use client"

import ProductCard from '@/components/ProductCard';
import { useAppDispatch, useAppSelector } from '@/redux/hooks';
import { fetchAllProducts, updateProductQuantity } from '@/redux/slice/ProductSlice'; // Add action to update product quantity
import React, { useEffect, useState } from 'react';

const ProductPage = () => {
    const dispatch = useAppDispatch();

    const products = useAppSelector(state => state.product.products);
    const [socket, setSocket] = useState<WebSocket | null>(null);

    // Fetch products when the component mounts
    useEffect(() => {
        dispatch(fetchAllProducts({
            pageIndex: "1",
            pageSize: "10"
        }));
    }, [dispatch]);

    useEffect(() => {
        // Initialize WebSocket connection
        const ws = new WebSocket('ws://localhost:8080/ws');

        ws.onopen = () => {
            console.log('Connected to WebSocket server');
        };

        // Handle incoming WebSocket messages
        ws.onmessage = (event) => {
            const updatedProduct = JSON.parse(event.data);
            dispatch(updateProductQuantity(updatedProduct));
        };

        ws.onclose = () => {
            console.log('WebSocket connection closed');
        };

        // Clean up WebSocket connection on component unmount
        return () => {
            ws.close();
        };
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
