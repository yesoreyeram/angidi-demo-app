# Angidi Frontend

Next.js-based frontend for the Angidi e-commerce platform.

## Overview

This is the frontend application for Angidi, built with Next.js 14+, TypeScript, and Tailwind CSS. It provides a modern, responsive user interface for the e-commerce platform.

## Prerequisites

- Node.js 20 LTS or higher
- npm 10+ (or pnpm 8+)

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/yesoreyeram/angidi-demo-app.git
cd angidi-demo-app/frontend

# Install dependencies
npm install
```

### Development

```bash
# Start the development server
npm run dev

# The app will be available at http://localhost:3000
```

### Building

```bash
# Create a production build
npm run build

# Start the production server
npm start
```

### Code Quality

```bash
# Run linter
npm run lint

# Fix linting issues
npm run lint:fix

# Format code
npm run format

# Check formatting
npm run format:check

# Type check
npm run type-check
```

## Project Structure

```
frontend/
├── src/
│   ├── app/              # Next.js app directory
│   │   ├── layout.tsx    # Root layout
│   │   ├── page.tsx      # Home page
│   │   └── globals.css   # Global styles
│   ├── components/       # Reusable components
│   │   ├── ui/           # Base UI components
│   │   ├── layout/       # Layout components
│   │   └── features/     # Feature-specific components
│   └── lib/              # Utility functions
│       ├── api/          # API client
│       ├── hooks/        # Custom React hooks
│       ├── utils/        # Helper functions
│       └── types/        # TypeScript types
├── public/               # Static assets
├── tests/                # E2E tests (Playwright)
├── .env.example          # Environment variables template
├── .env.local            # Local environment variables (git-ignored)
├── package.json          # Dependencies and scripts
├── tsconfig.json         # TypeScript configuration
├── tailwind.config.ts    # Tailwind CSS configuration
└── README.md             # This file
```

## Environment Variables

Create a `.env.local` file in the frontend directory:

```bash
# API Configuration
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_APP_NAME="Angidi"
NEXT_PUBLIC_APP_VERSION="0.1.0"
```

See `.env.example` for all available variables.

## Available Scripts

- `dev` - Start development server
- `build` - Create production build
- `start` - Start production server
- `lint` - Run ESLint
- `lint:fix` - Fix ESLint errors
- `format` - Format code with Prettier
- `format:check` - Check code formatting
- `type-check` - Run TypeScript compiler

## Features (Phase 1)

- ✅ Next.js 14 with App Router
- ✅ TypeScript for type safety
- ✅ Tailwind CSS for styling
- ✅ Responsive layout with header and footer
- ✅ API client for backend communication
- ✅ ESLint and Prettier for code quality
- ✅ Production-ready build configuration

## License

MIT License - see LICENSE file for details.

---

**Current Phase**: Phase 1 - Repository Scaffolding  
**Status**: ✅ Complete  
**Last Updated**: 2025-10-26
