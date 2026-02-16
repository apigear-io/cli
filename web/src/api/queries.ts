import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { apiClient } from './client';
import type {
  HealthResponse,
  StatusResponse,
  TemplateListResponse,
  TemplateInfo,
  InstallProgressEvent,
} from './types';

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

// Template queries
export function useTemplates() {
  return useQuery({
    queryKey: ['templates'],
    queryFn: () => apiClient.get<TemplateListResponse>('/templates'),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}

export function useTemplate(id: string) {
  return useQuery({
    queryKey: ['templates', id],
    queryFn: () => apiClient.get<TemplateInfo>(`/templates/get?id=${encodeURIComponent(id)}`),
    enabled: !!id,
  });
}

export function useCachedTemplates() {
  return useQuery({
    queryKey: ['templates', 'cache'],
    queryFn: () => apiClient.get<TemplateListResponse>('/templates/cache'),
    refetchInterval: 30000, // Refresh every 30s
  });
}

export function useSearchTemplates(query: string) {
  return useQuery({
    queryKey: ['templates', 'search', query],
    queryFn: () => apiClient.get<TemplateListResponse>(`/templates/search?q=${encodeURIComponent(query)}`),
    enabled: !!query,
  });
}

// Template mutations
export function useInstallTemplate() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({
      id,
      version,
      onProgress,
    }: {
      id: string;
      version?: string;
      onProgress?: (event: InstallProgressEvent) => void;
    }) => {
      const response = await fetch(`/api/v1/templates/install?id=${encodeURIComponent(id)}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'text/event-stream',
        },
        body: version ? JSON.stringify({ version }) : '{}',
      });

      if (!response.ok) {
        throw new Error(`Installation failed: ${response.statusText}`);
      }

      const reader = response.body?.getReader();
      if (!reader) {
        throw new Error('No response body');
      }

      const decoder = new TextDecoder();
      let buffer = '';

      while (true) {
        const { done, value } = await reader.read();
        if (done) break;

        buffer += decoder.decode(value, { stream: true });
        const lines = buffer.split('\n\n');
        buffer = lines.pop() || '';

        for (const line of lines) {
          if (line.startsWith('data: ')) {
            const data: InstallProgressEvent = JSON.parse(line.slice(6));

            if (data.type === 'progress' && onProgress) {
              onProgress(data);
            } else if (data.type === 'complete') {
              return data;
            } else if (data.type === 'error') {
              throw new Error(data.error || 'Installation failed');
            }
          }
        }
      }

      throw new Error('Installation stream ended unexpectedly');
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['templates'] });
      queryClient.invalidateQueries({ queryKey: ['templates', 'cache'] });
    },
  });
}

export function useRemoveTemplate() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => apiClient.delete<{ message: string }>(`/templates/cache/remove?id=${encodeURIComponent(id)}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['templates'] });
      queryClient.invalidateQueries({ queryKey: ['templates', 'cache'] });
    },
  });
}

export function useUpdateRegistry() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: () => apiClient.post<{ message: string }>('/templates/registry/update', {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['templates'] });
    },
  });
}

export function useCleanCache() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: () => apiClient.post<{ message: string }>('/templates/cache/clean', {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['templates'] });
      queryClient.invalidateQueries({ queryKey: ['templates', 'cache'] });
    },
  });
}
