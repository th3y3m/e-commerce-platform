import React from 'react';

interface ProductCardProps {
    product: {
        ProductID: string;
        ProductName: string;
        Description: string;
        Price: number;
        Quantity: number;
    };
}

const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
    return (
        <div className="border rounded-lg p-4 shadow-md">
            <h3 className="text-xl font-bold mb-2">{product.ProductName}</h3>
            <p className="text-gray-700 mb-2">{product.Description}</p>
            <p className="text-gray-900 font-semibold mb-2">${product.Price}</p>
            <p className="text-gray-600">Quantity: {product.Quantity}</p>
        </div>
    );
};

export default ProductCard;