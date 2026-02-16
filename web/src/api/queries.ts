import { useQuery } from '@tanstack/react-query';
import { apiClient } from './client';
import type { HealthResponse, StatusResponse } from './types';

export function useHealth() {
  return useQuery({
    queryKey: ['health'],
    queryFn: () => apiClient.get<HealthResponse>('/health'),
    refetchInterval: 30000, // Refetch every 30 seconds
  });
}

export function useStatus() {
  return useQuery({
    queryKey: ['status'],
    queryFn: () => apiClient.get<StatusResponse>('/status'),
    refetchInterval: 60000, // Refetch every 60 seconds
  });
}
