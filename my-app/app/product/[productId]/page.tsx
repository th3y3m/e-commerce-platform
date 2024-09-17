// Import the necessary hook
'use client';

import { useParams } from 'next/navigation';

const ProductPage = () => {
    // Get the dynamic parameter from the route
    const { productId } = useParams();

    return (
        <div>
            <h1>Product ID: {productId}</h1>
        </div>
    );
};

export default ProductPage;
