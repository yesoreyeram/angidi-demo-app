'use client';

import { useState, useEffect } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { useAuth } from '@/contexts/AuthContext';
import { apiClient } from '@/lib/api/client';
import type { Product } from '@/lib/api/types';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';
import { toast } from 'sonner';
import Link from 'next/link';

const productSchema = z.object({
  name: z.string().min(2, 'Name must be at least 2 characters'),
  description: z.string().min(10, 'Description must be at least 10 characters'),
  price: z.string().refine((val) => !isNaN(parseFloat(val)) && parseFloat(val) > 0, {
    message: 'Price must be a positive number',
  }),
  stock: z.string().refine((val) => !isNaN(parseInt(val)) && parseInt(val) >= 0, {
    message: 'Stock must be a non-negative number',
  }),
  category: z.string().min(2, 'Category must be at least 2 characters'),
  imageURL: z.string().url('Must be a valid URL').or(z.literal('')),
});

type ProductForm = z.infer<typeof productSchema>;

export default function EditProductPage() {
  const params = useParams();
  const router = useRouter();
  const { user, isAuthenticated, isLoading: authLoading } = useAuth();
  const [product, setProduct] = useState<Product | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [isFetching, setIsFetching] = useState(true);

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<ProductForm>({
    resolver: zodResolver(productSchema),
  });

  useEffect(() => {
    if (!authLoading && (!isAuthenticated || user?.role !== 'admin')) {
      toast.error('Access denied. Admin privileges required.');
      router.push('/');
    }
  }, [isAuthenticated, user, authLoading, router]);

  useEffect(() => {
    if (user?.role === 'admin' && params.id && typeof params.id === 'string') {
      fetchProduct(params.id);
    }
  }, [user, params.id]);

  const fetchProduct = async (id: string) => {
    setIsFetching(true);
    try {
      const response = await apiClient.getProduct(id);

      if (response.error || !response.data) {
        toast.error(response.error || 'Failed to load product');
        router.push('/admin/products');
        return;
      }

      const productData = response.data;
      setProduct(productData);

      // Pre-fill form
      setValue('name', productData.name);
      setValue('description', productData.description);
      setValue('price', productData.price.toString());
      setValue('stock', productData.stock.toString());
      setValue('category', productData.category);
      setValue('imageURL', productData.imageURL);
    } catch (error) {
      toast.error('Failed to load product');
      router.push('/admin/products');
    } finally {
      setIsFetching(false);
    }
  };

  if (authLoading || !isAuthenticated || user?.role !== 'admin' || isFetching) {
    return (
      <div className="container mx-auto px-4 py-12">
        <div className="text-center">
          <div className="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
          <p className="mt-4 text-gray-600">Loading...</p>
        </div>
      </div>
    );
  }

  if (!product) {
    return null;
  }

  const onSubmit = async (data: ProductForm) => {
    if (!params.id || typeof params.id !== 'string') return;

    setIsLoading(true);
    try {
      const response = await apiClient.updateProduct(params.id, {
        name: data.name,
        description: data.description,
        price: parseFloat(data.price),
        stock: parseInt(data.stock),
        category: data.category,
        imageURL: data.imageURL || '',
      });

      if (response.error || !response.data) {
        toast.error(response.error || 'Failed to update product');
        return;
      }

      toast.success('Product updated successfully');
      router.push('/admin/products');
    } catch (error) {
      toast.error('An unexpected error occurred');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <Link
        href="/admin/products"
        className="inline-flex items-center text-blue-600 hover:underline mb-6"
      >
        ← Back to Products
      </Link>

      <div className="max-w-2xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">Edit Product</h1>

        <form
          onSubmit={handleSubmit(onSubmit)}
          className="bg-white rounded-lg shadow-md p-8 space-y-6"
        >
          <Input
            label="Product Name"
            type="text"
            placeholder="e.g., iPhone 15 Pro"
            error={errors.name?.message}
            {...register('name')}
          />

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Description
            </label>
            <textarea
              className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 ${
                errors.description ? 'border-red-500' : 'border-gray-300'
              }`}
              rows={4}
              placeholder="Detailed product description..."
              {...register('description')}
            />
            {errors.description && (
              <p className="mt-1 text-sm text-red-600">{errors.description.message}</p>
            )}
          </div>

          <div className="grid md:grid-cols-2 gap-6">
            <Input
              label="Price ($)"
              type="number"
              step="0.01"
              placeholder="99.99"
              error={errors.price?.message}
              {...register('price')}
            />

            <Input
              label="Stock"
              type="number"
              placeholder="100"
              error={errors.stock?.message}
              {...register('stock')}
            />
          </div>

          <Input
            label="Category"
            type="text"
            placeholder="e.g., Electronics"
            error={errors.category?.message}
            {...register('category')}
          />

          <Input
            label="Image URL (optional)"
            type="text"
            placeholder="https://example.com/image.jpg"
            error={errors.imageURL?.message}
            {...register('imageURL')}
          />

          <div className="flex gap-4 pt-6 border-t">
            <Button type="submit" variant="primary" isLoading={isLoading}>
              Save Changes
            </Button>
            <Link href="/admin/products">
              <Button type="button" variant="secondary">
                Cancel
              </Button>
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
}
