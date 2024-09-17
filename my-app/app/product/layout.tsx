// app/Product/layout.tsx

export default function ProductLayout({ children }: { children: React.ReactNode }) {
    return (
        <div>
            <nav>Product Navigation</nav>
            <main>{children}</main>
        </div>
    );
}
