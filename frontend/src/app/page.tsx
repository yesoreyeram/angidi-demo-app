import Link from 'next/link';

export default function Home() {
  return (
    <div className="container mx-auto px-4 py-12">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-4xl font-bold mb-6">
          Welcome to Angidi E-Commerce Platform
        </h1>
        
        <div className="bg-blue-50 border-l-4 border-blue-600 p-6 mb-8">
          <h2 className="text-2xl font-semibold mb-2 text-blue-900">
            Phase 2: Core Services
          </h2>
          <p className="text-gray-700">
            Full-featured e-commerce platform with user authentication, product catalog,
            and admin management. Experience enterprise-quality microservices architecture
            with JWT authentication, role-based access control, and comprehensive API documentation.
          </p>
        </div>

        <div className="grid md:grid-cols-2 gap-6 mb-8">
          <div className="border rounded-lg p-6 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-3">Backend (Go)</h3>
            <ul className="list-disc list-inside text-gray-700 space-y-2">
              <li>User Service (JWT auth, registration, profile)</li>
              <li>Product Service (CRUD, pagination, filters)</li>
              <li>API Gateway with middleware</li>
              <li>Rate limiting & CORS</li>
              <li>Comprehensive testing (47 tests)</li>
              <li>OpenAPI 3.0.3 specification</li>
            </ul>
          </div>

          <div className="border rounded-lg p-6 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-3">Frontend (Next.js)</h3>
            <ul className="list-disc list-inside text-gray-700 space-y-2">
              <li>User registration & login</li>
              <li>Protected routes & auth state</li>
              <li>Product catalog with filters</li>
              <li>User profile management</li>
              <li>Admin product management</li>
              <li>Toast notifications & error handling</li>
            </ul>
          </div>
        </div>

        <div className="bg-gray-50 rounded-lg p-6 mb-6">
          <h3 className="text-xl font-semibold mb-3">Quick Links</h3>
          <div className="flex flex-wrap gap-4">
            <a
              href="http://localhost:8080/health"
              target="_blank"
              rel="noopener noreferrer"
              className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition-colors"
            >
              Backend Health Check
            </a>
            <Link
              href="/products"
              className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition-colors"
            >
              Browse Products
            </Link>
            <a
              href="/register"
              className="bg-purple-600 text-white px-4 py-2 rounded hover:bg-purple-700 transition-colors"
            >
              Create Account
            </a>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow p-6 mb-6">
          <h3 className="text-xl font-semibold mb-3">Features</h3>
          <div className="grid md:grid-cols-2 gap-4">
            <div>
              <h4 className="font-semibold text-blue-600 mb-2">🔐 Authentication</h4>
              <p className="text-sm text-gray-600">
                Secure JWT-based authentication with bcrypt password hashing
              </p>
            </div>
            <div>
              <h4 className="font-semibold text-blue-600 mb-2">🛍️ Product Catalog</h4>
              <p className="text-sm text-gray-600">
                Browse products with search, category, and price filters
              </p>
            </div>
            <div>
              <h4 className="font-semibold text-blue-600 mb-2">👤 User Profiles</h4>
              <p className="text-sm text-gray-600">
                Manage your account information and preferences
              </p>
            </div>
            <div>
              <h4 className="font-semibold text-blue-600 mb-2">⚙️ Admin Panel</h4>
              <p className="text-sm text-gray-600">
                Full CRUD operations for product management (admin only)
              </p>
            </div>
          </div>
        </div>

        <div className="mt-8 text-center text-gray-600">
          <p>
            Built with ❤️ using Go, Next.js, TypeScript, and Tailwind CSS
          </p>
          <p className="text-sm mt-2">
            Enterprise-quality microservices architecture
          </p>
        </div>
      </div>
    </div>
  );
}
