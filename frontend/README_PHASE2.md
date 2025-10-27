# Angidi Frontend - Phase 2

Enterprise-quality Next.js frontend with complete authentication, product catalog, and admin management features.

## Features

### 🔐 Authentication
- User registration with validation
- Secure login with JWT tokens
- Automatic token refresh
- Protected routes
- Role-based access control (user/admin)

### 🛍️ Product Catalog
- Browse products with responsive grid layout
- Advanced filtering (search, category, price range)
- Pagination support
- Product detail pages
- Real-time stock display

### 👤 User Profile
- View and update profile information
- Account details and member status
- Secure profile updates

### ⚙️ Admin Panel
- Full CRUD operations for products
- Product management dashboard
- Create new products with validation
- Edit existing products
- Delete products with confirmation
- Admin-only access with role verification

## Tech Stack

- **Framework**: Next.js 16 with App Router
- **Language**: TypeScript
- **Styling**: Tailwind CSS 4
- **Forms**: React Hook Form + Zod validation
- **State**: React Context API
- **Notifications**: Sonner (toast notifications)
- **HTTP**: Fetch API with custom client

## Getting Started

### Prerequisites

- Node.js 20+ and npm
- Backend API running on `http://localhost:8080`

### Installation

```bash
# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build

# Start production server
npm start
```

### Environment Variables

Create a `.env.local` file:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## Project Structure

```
src/
├── app/                      # Next.js app router pages
│   ├── layout.tsx            # Root layout with providers
│   ├── page.tsx              # Home page
│   ├── login/                # Login page
│   ├── register/             # Registration page
│   ├── products/             # Product catalog
│   │   ├── page.tsx          # Product list
│   │   └── [id]/page.tsx     # Product detail
│   ├── profile/              # User profile
│   └── admin/
│       └── products/         # Admin product management
│           ├── page.tsx      # Product list
│           ├── new/          # Create product
│           └── [id]/edit/    # Edit product
├── components/
│   ├── Header.tsx            # Navigation header
│   └── ui/                   # Reusable UI components
│       ├── Button.tsx
│       └── Input.tsx
├── contexts/
│   └── AuthContext.tsx       # Authentication state management
└── lib/
    └── api/
        ├── client.ts         # API client with all endpoints
        └── types.ts          # TypeScript type definitions
```

## Pages

### Public Pages

#### Home (`/`)
- Welcome page with feature overview
- Links to products, registration, and backend health check
- Phase 2 highlights

#### Login (`/login`)
- Email and password fields
- Form validation
- Error handling
- Link to registration

#### Register (`/register`)
- Full registration form (name, email, password)
- Password confirmation
- Client-side validation
- Automatic login after registration

#### Products (`/products`)
- Product grid with responsive layout
- Advanced filters:
  - Search by name/description
  - Filter by category
  - Price range (min/max)
- Pagination
- Empty states
- Loading indicators

#### Product Detail (`/products/[id]`)
- Product image
- Full description
- Price and stock information
- Category display
- Availability status
- Back to products link

### Protected Pages

#### Profile (`/profile`)
- **Requires**: Authentication
- View account information
- Update user name
- Email (read-only)
- Role badge (user/admin)
- Member since date

### Admin Pages

#### Admin Products (`/admin/products`)
- **Requires**: Admin role
- Product table with all products
- Actions: Edit, Delete
- Create new product button
- Product thumbnails
- Stock and price display
- Empty state with CTA

#### Create Product (`/admin/products/new`)
- **Requires**: Admin role
- Product creation form
- Fields: name, description, price, stock, category, image URL
- Validation with helpful error messages
- Cancel button returns to list

#### Edit Product (`/admin/products/[id]/edit`)
- **Requires**: Admin role
- Pre-filled form with current product data
- Update any product field
- Validation
- Save/Cancel actions

## Authentication Flow

### Registration
1. User fills registration form
2. Client validates inputs (name, email, password)
3. API call to `/api/v1/users/register`
4. On success: JWT tokens stored, user logged in
5. Redirect to products page

### Login
1. User enters email and password
2. Client validates inputs
3. API call to `/api/v1/users/login`
4. On success: JWT tokens stored in localStorage
5. Redirect to products page

### Token Management
- Access token: 15 minutes expiration
- Refresh token: 7 days expiration
- Stored in localStorage
- Automatic refresh (can be implemented)
- Bearer token sent in Authorization header

### Protected Routes
- Check authentication status in `useEffect`
- Redirect to `/login` if not authenticated
- Admin routes check user role

## API Integration

### API Client (`lib/api/client.ts`)

All Phase 2 endpoints are integrated:

**Authentication**:
- `POST /api/v1/users/register` - User registration
- `POST /api/v1/users/login` - User login
- `POST /api/v1/users/refresh-token` - Token refresh

**User Profile**:
- `GET /api/v1/users/me` - Get profile
- `PUT /api/v1/users/me` - Update profile

**Products**:
- `GET /api/v1/products` - List products with filters
- `GET /api/v1/products/:id` - Get product
- `POST /api/v1/products` - Create product (admin)
- `PUT /api/v1/products/:id` - Update product (admin)
- `DELETE /api/v1/products/:id` - Delete product (admin)

### Type Safety

All API requests and responses are fully typed using TypeScript interfaces in `lib/api/types.ts`:

- User
- Product
- AuthResponse
- ProductListResponse
- Request/Response DTOs
- Error types

## Form Validation

Using Zod schemas for runtime validation:

### Registration
- Name: min 2 characters
- Email: valid email format
- Password: min 8 characters
- Confirm password: must match

### Login
- Email: valid email format
- Password: required

### Product Form
- Name: min 2 characters
- Description: min 10 characters
- Price: positive number
- Stock: non-negative integer
- Category: min 2 characters
- Image URL: valid URL or empty

## Components

### Header
- Navigation links (Home, Products)
- Authentication state display
- User menu (Profile, Admin, Logout)
- Login/Register buttons for guests
- Role-based menu items

### Button
- Variants: primary, secondary, danger
- Loading state with spinner
- Disabled state
- Full TypeScript support

### Input
- Label support
- Error message display
- Ref forwarding
- All HTML input attributes

## Error Handling

### API Errors
- Displayed via toast notifications
- User-friendly error messages
- Validation errors shown inline on forms

### Network Errors
- Caught and displayed to user
- Graceful fallback to error state

### Loading States
- Spinner animations
- Loading text
- Disabled buttons during operations

## Security

### Best Practices
- JWT tokens for authentication
- Bearer token in Authorization header
- Client-side validation
- Server-side validation via API
- Protected routes
- Role-based access control
- No sensitive data in client code

### Considerations
- Tokens in localStorage (can be moved to HTTP-only cookies)
- CORS configured on backend
- Input sanitization on server
- Rate limiting on API

## Testing

### Manual Testing Checklist

- [ ] User can register new account
- [ ] User can login with credentials
- [ ] Invalid credentials show error
- [ ] User can view products
- [ ] Filters work correctly
- [ ] Pagination works
- [ ] Product detail page loads
- [ ] Protected routes redirect if not authenticated
- [ ] User can view and update profile
- [ ] Admin can access admin panel
- [ ] Admin can create products
- [ ] Admin can edit products
- [ ] Admin can delete products
- [ ] Non-admin cannot access admin pages
- [ ] Logout works correctly
- [ ] Toast notifications appear

### Future Automated Tests

Consider adding:
- Component tests with React Testing Library
- E2E tests with Playwright
- API integration tests

## Development

### Scripts

```bash
# Development server with hot reload
npm run dev

# Type checking
npm run type-check

# Linting
npm run lint

# Auto-fix linting issues
npm run lint:fix

# Format code
npm run format

# Check formatting
npm run format:check

# Production build
npm run build

# Start production server
npm start
```

### Code Style

- TypeScript for type safety
- Functional components with hooks
- ESLint + Prettier for formatting
- Tailwind CSS for styling
- Client components marked with 'use client'

## Performance

### Optimizations
- Static page generation where possible
- Client-side navigation with Next.js Link
- Responsive images (consider next/image)
- Lazy loading for routes
- Minimal bundle size

### Considerations
- Consider implementing:
  - Image optimization with next/image
  - Code splitting
  - React Query for caching
  - Virtual scrolling for large lists

## Accessibility

### Current Implementation
- Semantic HTML
- Form labels
- Button states
- Color contrast
- Keyboard navigation (native)

### Future Improvements
- ARIA labels
- Focus management
- Screen reader testing
- Keyboard shortcuts

## Browser Support

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## Deployment

### Build

```bash
npm run build
```

### Environment Variables

Set in deployment platform:
```env
NEXT_PUBLIC_API_URL=https://api.yoursite.com
```

### Platforms

Compatible with:
- Vercel (recommended for Next.js)
- Netlify
- AWS Amplify
- Docker container

## Troubleshooting

### API Connection Issues
- Verify backend is running on correct port
- Check NEXT_PUBLIC_API_URL is set correctly
- Verify CORS is configured on backend

### Authentication Issues
- Clear localStorage and try again
- Check JWT secret matches backend
- Verify token expiration settings

### Build Errors
- Run `npm run type-check`
- Fix any TypeScript errors
- Clear `.next` folder and rebuild

## Future Enhancements

### Phase 3 Possibilities
- Shopping cart functionality
- Order management
- Payment integration
- Product reviews and ratings
- Wishlist
- Advanced search
- Product images upload
- Real-time notifications
- Multi-language support
- Dark mode

## Contributing

1. Create feature branch
2. Make changes
3. Run tests and linting
4. Submit pull request

## License

MIT

## Support

For issues or questions, please create a GitHub issue.
