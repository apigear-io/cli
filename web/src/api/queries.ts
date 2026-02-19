import { useSuspenseQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { apiClient } from './client';
import { queryKeys } from './queryKeys';
import type {
  HealthResponse,
  StatusResponse,
  TemplateListResponse,
  TemplateInfo,
  InstallProgressEvent,
  StreamDashboardStats,
  ProxyInfo,
  ProxyConfig,
  CreateProxyRequest,
  ClientInfo,
  ClientConfig,
  CreateClientRequest,
} from './types';

export function useHealth() {
  return useSuspenseQuery({
    queryKey: queryKeys.health(),
    queryFn: () => apiClient.get<HealthResponse>('/health'),
    refetchInterval: 30000, // Refetch every 30 seconds
  });
}

export function useStatus() {
  return useSuspenseQuery({
    queryKey: queryKeys.status(),
    queryFn: () => apiClient.get<StatusResponse>('/status'),
    refetchInterval: 60000, // Refetch every 60 seconds
  });
}

// Template queries
export function useTemplates() {
  return useSuspenseQuery({
    queryKey: queryKeys.templates.registry(),
    queryFn: () => apiClient.get<TemplateListResponse>('/templates'),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}

export function useTemplate(id: string) {
  return useSuspenseQuery({
    queryKey: queryKeys.templates.detail(id),
    queryFn: () => apiClient.get<TemplateInfo>(`/templates/get?id=${encodeURIComponent(id)}`),
  });
}

export function useCachedTemplates() {
  return useSuspenseQuery({
    queryKey: queryKeys.templates.cache(),
    queryFn: () => apiClient.get<TemplateListResponse>('/templates/cache'),
    refetchInterval: 30000, // Refresh every 30s
  });
}

export function useSearchTemplates(query: string) {
  return useSuspenseQuery({
    queryKey: queryKeys.templates.search(query),
    queryFn: () => apiClient.get<TemplateListResponse>(`/templates/search?q=${encodeURIComponent(query)}`),
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
      queryClient.invalidateQueries({ queryKey: queryKeys.templates.all() });
    },
  });
}

export function useRemoveTemplate() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => apiClient.delete<{ message: string }>(`/templates/cache/remove?id=${encodeURIComponent(id)}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.templates.all() });
    },
  });
}

export function useUpdateRegistry() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: () => apiClient.post<{ message: string }>('/templates/registry/update', {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.templates.all() });
    },
  });
}

export function useCleanCache() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: () => apiClient.post<{ message: string }>('/templates/cache/clean', {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.templates.all() });
    },
  });
}

// Stream queries

export function useStreamDashboard() {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.dashboard(),
    queryFn: () => apiClient.get<StreamDashboardStats>('/stream/dashboard'),
    refetchInterval: 5000, // Refresh every 5 seconds
  });
}

export function useProxies() {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.proxies.list(),
    queryFn: () => apiClient.get<ProxyInfo[]>('/stream/proxies'),
    refetchInterval: 3000, // Refresh every 3 seconds
  });
}

export function useProxy(name: string) {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.proxies.detail(name),
    queryFn: () => apiClient.get<ProxyInfo>(`/stream/proxies/${encodeURIComponent(name)}`),
    refetchInterval: 2000, // Refresh every 2 seconds
  });
}

export function useProxyStats(name: string) {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.proxies.stats(name),
    queryFn: () => apiClient.get<ProxyInfo>(`/stream/proxies/${encodeURIComponent(name)}/stats`),
    refetchInterval: 1000, // Refresh every second
  });
}

export function useClients() {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.clients.list(),
    queryFn: () => apiClient.get<ClientInfo[]>('/stream/clients'),
    refetchInterval: 3000, // Refresh every 3 seconds
  });
}

export function useClient(name: string) {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.clients.detail(name),
    queryFn: () => apiClient.get<ClientInfo>(`/stream/clients/${encodeURIComponent(name)}`),
    refetchInterval: 2000, // Refresh every 2 seconds
  });
}

// Stream mutations - Proxies

export function useCreateProxy() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: CreateProxyRequest) =>
      apiClient.post<ProxyInfo>('/stream/proxies', request),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.proxies.all() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

export function useUpdateProxy() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ name, config }: { name: string; config: ProxyConfig }) =>
      apiClient.put<ProxyInfo>(`/stream/proxies/${encodeURIComponent(name)}`, config),
    onSuccess: (_, { name }) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.proxies.detail(name) });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.proxies.list() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

export function useDeleteProxy() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.delete(`/stream/proxies/${encodeURIComponent(name)}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.proxies.all() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

export function useStartProxy() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.post<ProxyInfo>(`/stream/proxies/${encodeURIComponent(name)}/start`, {}),
    onSuccess: (_, name) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.proxies.detail(name) });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.proxies.list() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

export function useStopProxy() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.post<ProxyInfo>(`/stream/proxies/${encodeURIComponent(name)}/stop`, {}),
    onSuccess: (_, name) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.proxies.detail(name) });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.proxies.list() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

// Stream mutations - Clients

export function useCreateClient() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: CreateClientRequest) =>
      apiClient.post<ClientInfo>('/stream/clients', request),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.clients.all() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

export function useUpdateClient() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ name, config }: { name: string; config: ClientConfig }) =>
      apiClient.put<ClientInfo>(`/stream/clients/${encodeURIComponent(name)}`, config),
    onSuccess: (_, { name }) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.clients.detail(name) });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.clients.list() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

export function useDeleteClient() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.delete(`/stream/clients/${encodeURIComponent(name)}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.clients.all() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

export function useConnectClient() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.post<ClientInfo>(`/stream/clients/${encodeURIComponent(name)}/connect`, {}),
    onSuccess: (_, name) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.clients.detail(name) });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.clients.list() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}

export function useDisconnectClient() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.post<ClientInfo>(`/stream/clients/${encodeURIComponent(name)}/disconnect`, {}),
    onSuccess: (_, name) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.clients.detail(name) });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.clients.list() });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.dashboard() });
    },
  });
}
