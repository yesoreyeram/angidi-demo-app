export default function Home() {
  return (
    <div className="container mx-auto px-4 py-12">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-4xl font-bold mb-6">
          Welcome to Angidi E-Commerce Platform
        </h1>
        
        <div className="bg-blue-50 border-l-4 border-blue-600 p-6 mb-8">
          <h2 className="text-2xl font-semibold mb-2 text-blue-900">
            Phase 1: Repository Scaffolding
          </h2>
          <p className="text-gray-700">
            This is the initial setup phase of the Angidi platform. We've implemented
            the foundational project structure, development tooling, and CI/CD infrastructure.
          </p>
        </div>

        <div className="grid md:grid-cols-2 gap-6 mb-8">
          <div className="border rounded-lg p-6 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-3">Backend (Go)</h3>
            <ul className="list-disc list-inside text-gray-700 space-y-2">
              <li>Go module initialized</li>
              <li>Project structure established</li>
              <li>Configuration management</li>
              <li>Structured logging</li>
              <li>Health check endpoint</li>
            </ul>
          </div>

          <div className="border rounded-lg p-6 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-3">Frontend (Next.js)</h3>
            <ul className="list-disc list-inside text-gray-700 space-y-2">
              <li>Next.js 14 with App Router</li>
              <li>TypeScript configuration</li>
              <li>Tailwind CSS styling</li>
              <li>Responsive layout</li>
              <li>ESLint and Prettier</li>
            </ul>
          </div>
        </div>

        <div className="bg-gray-50 rounded-lg p-6">
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
            <a
              href="/products"
              className="bg-gray-600 text-white px-4 py-2 rounded hover:bg-gray-700 transition-colors"
            >
              Products (Coming Soon)
            </a>
            <a
              href="/cart"
              className="bg-gray-600 text-white px-4 py-2 rounded hover:bg-gray-700 transition-colors"
            >
              Shopping Cart (Coming Soon)
            </a>
          </div>
        </div>

        <div className="mt-8 text-center text-gray-600">
          <p>
            Built with ❤️ as a learning project for advanced system design concepts
          </p>
        </div>
      </div>
    </div>
  );
}
