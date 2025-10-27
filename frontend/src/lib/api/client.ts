import type {
  User,
  RegisterRequest,
  LoginRequest,
  AuthResponse,
  RefreshTokenRequest,
  UpdateProfileRequest,
  Product,
  ProductListResponse,
  ProductFilters,
  CreateProductRequest,
  UpdateProductRequest,
  HealthCheck,
  ApiError,
} from './types';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export interface ApiResponse<T> {
  data?: T;
  error?: string;
  details?: Record<string, string>;
}

class ApiClient {
  private baseUrl: string;
  private accessToken: string | null = null;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  setAccessToken(token: string | null) {
    this.accessToken = token;
  }

  getAccessToken(): string | null {
    return this.accessToken;
  }

  private async request<T>(
    endpoint: string,
    options?: RequestInit
  ): Promise<ApiResponse<T>> {
    try {
      const headers: Record<string, string> = {
        'Content-Type': 'application/json',
      };

      if (this.accessToken) {
        headers['Authorization'] = `Bearer ${this.accessToken}`;
      }

      // Merge with any additional headers from options
      if (options?.headers) {
        Object.entries(options.headers).forEach(([key, value]) => {
          if (typeof value === 'string') {
            headers[key] = value;
          }
        });
      }

      const response = await fetch(`${this.baseUrl}${endpoint}`, {
        ...options,
        headers,
      });

      const contentType = response.headers.get('content-type');
      const isJson = contentType?.includes('application/json');

      if (!response.ok) {
        if (isJson) {
          const errorData = (await response.json()) as ApiError;
          return {
            error: errorData.error || `HTTP error! status: ${response.status}`,
            details: errorData.details,
          };
        }
        return {
          error: `HTTP error! status: ${response.status}`,
        };
      }

      if (isJson) {
        const data = await response.json();
        return { data };
      }

      return { data: {} as T };
    } catch (error) {
      console.error('API request failed:', error);
      return {
        error: error instanceof Error ? error.message : 'An error occurred',
      };
    }
  }

  // Health Check
  async healthCheck(): Promise<ApiResponse<HealthCheck>> {
    return this.request<HealthCheck>('/health');
  }

  // Authentication endpoints
  async register(data: RegisterRequest): Promise<ApiResponse<AuthResponse>> {
    return this.request<AuthResponse>('/api/v1/users/register', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async login(data: LoginRequest): Promise<ApiResponse<AuthResponse>> {
    return this.request<AuthResponse>('/api/v1/users/login', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async refreshToken(
    data: RefreshTokenRequest
  ): Promise<ApiResponse<AuthResponse>> {
    return this.request<AuthResponse>('/api/v1/users/refresh-token', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // User profile endpoints
  async getProfile(): Promise<ApiResponse<User>> {
    return this.request<User>('/api/v1/users/me');
  }

  async updateProfile(
    data: UpdateProfileRequest
  ): Promise<ApiResponse<User>> {
    return this.request<User>('/api/v1/users/me', {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  // Product endpoints
  async getProducts(
    filters?: ProductFilters
  ): Promise<ApiResponse<ProductListResponse>> {
    const params = new URLSearchParams();
    if (filters?.page) params.append('page', filters.page.toString());
    if (filters?.perPage) params.append('perPage', filters.perPage.toString());
    if (filters?.category) params.append('category', filters.category);
    if (filters?.minPrice)
      params.append('minPrice', filters.minPrice.toString());
    if (filters?.maxPrice)
      params.append('maxPrice', filters.maxPrice.toString());
    if (filters?.search) params.append('search', filters.search);

    const query = params.toString();
    const endpoint = query ? `/api/v1/products?${query}` : '/api/v1/products';

    return this.request<ProductListResponse>(endpoint);
  }

  async getProduct(id: string): Promise<ApiResponse<Product>> {
    return this.request<Product>(`/api/v1/products/${id}`);
  }

  async createProduct(
    data: CreateProductRequest
  ): Promise<ApiResponse<Product>> {
    return this.request<Product>('/api/v1/products', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateProduct(
    id: string,
    data: UpdateProductRequest
  ): Promise<ApiResponse<Product>> {
    return this.request<Product>(`/api/v1/products/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteProduct(id: string): Promise<ApiResponse<void>> {
    return this.request<void>(`/api/v1/products/${id}`, {
      method: 'DELETE',
    });
  }
}

export const apiClient = new ApiClient(API_URL);
