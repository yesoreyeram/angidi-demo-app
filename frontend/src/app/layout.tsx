import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Angidi - E-Commerce Platform",
  description: "A comprehensive learning project for advanced system design concepts",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="font-sans antialiased">
        <div className="min-h-screen flex flex-col">
          <header className="bg-blue-600 text-white shadow-md">
            <div className="container mx-auto px-4 py-4">
              <nav className="flex items-center justify-between">
                <h1 className="text-2xl font-bold">Angidi</h1>
                <ul className="flex gap-6">
                  <li>
                    <a href="/" className="hover:underline">
                      Home
                    </a>
                  </li>
                  <li>
                    <a href="/products" className="hover:underline">
                      Products
                    </a>
                  </li>
                  <li>
                    <a href="/cart" className="hover:underline">
                      Cart
                    </a>
                  </li>
                </ul>
              </nav>
            </div>
          </header>
          <main className="flex-1">{children}</main>
          <footer className="bg-gray-800 text-white py-8">
            <div className="container mx-auto px-4 text-center">
              <p>&copy; 2025 Angidi. Built with Next.js and TypeScript.</p>
              <p className="text-sm text-gray-400 mt-2">
                Phase 1 - Repository Scaffolding
              </p>
            </div>
          </footer>
        </div>
      </body>
    </html>
  );
}
