import { useSuspenseQuery, useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
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
  ScriptFile,
  ScriptInfo,
  SaveScriptRequest,
  SaveScriptResponse,
  RunScriptResponse,
  RunCodeRequest,
  TraceFileInfo,
  TraceStats,
  TraceFileResponse,
  EditTraceRequest,
  MergeTracesRequest,
  ExportTraceRequest,
  EditorStats,
  EditorMessagesResponse,
  EditorTimelineResponse,
  EditorSeekResponse,
  EditorJQResponse,
  EditorFilters,
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

// Stream queries - Scripts

export function useScripts() {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.scripts.list(),
    queryFn: async () => {
      const response = await apiClient.get<{ scripts: string[] | null }>('/stream/scripts');
      return response.scripts || [];
    },
    refetchInterval: 10000, // Refresh every 10 seconds
  });
}

export function useScript(name: string) {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.scripts.detail(name),
    queryFn: () => apiClient.get<ScriptFile>(`/stream/scripts/${encodeURIComponent(name)}`),
  });
}

export function useRunningScripts() {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.scripts.running(),
    queryFn: async () => {
      const response = await apiClient.get<{ scripts: ScriptInfo[] | null }>('/stream/scripts/running');
      return response.scripts || [];
    },
    refetchInterval: 3000, // Refresh every 3 seconds
  });
}

// Stream mutations - Scripts

export function useSaveScript() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: SaveScriptRequest) =>
      apiClient.post<SaveScriptResponse>('/stream/scripts', request),
    onSuccess: (_, { name }) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.scripts.detail(name) });
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.scripts.list() });
    },
  });
}

export function useDeleteScript() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.delete(`/stream/scripts/${encodeURIComponent(name)}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.scripts.all() });
    },
  });
}

export function useRunScript() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.post<RunScriptResponse>(`/stream/scripts/${encodeURIComponent(name)}/run`, {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.scripts.running() });
    },
  });
}

export function useRunCode() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: RunCodeRequest) =>
      apiClient.post<RunScriptResponse>('/stream/scripts/run', request),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.scripts.running() });
    },
  });
}

export function useStopScript() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) =>
      apiClient.post(`/stream/scripts/stop/${encodeURIComponent(id)}`, {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.scripts.running() });
    },
  });
}

// Stream queries - Traces

export function useTraceFiles() {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.traces.list(),
    queryFn: async () => {
      const response = await apiClient.get<{ files: TraceFileInfo[] | null }>('/stream/traces');
      return response.files || [];
    },
    refetchInterval: 10000, // Refresh every 10 seconds
  });
}

export function useTraceStats() {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.traces.stats(),
    queryFn: () => apiClient.get<TraceStats>('/stream/traces/stats'),
    refetchInterval: 10000, // Refresh every 10 seconds
  });
}

export function useTraceFile(name: string, options?: { direction?: string; limit?: number }) {
  return useSuspenseQuery({
    queryKey: [...queryKeys.stream.traces.detail(name), options],
    queryFn: () => {
      const params = new URLSearchParams();
      if (options?.direction) params.append('direction', options.direction);
      if (options?.limit) params.append('limit', options.limit.toString());
      const query = params.toString() ? `?${params.toString()}` : '';
      return apiClient.get<TraceFileResponse>(`/stream/traces/${encodeURIComponent(name)}${query}`);
    },
  });
}

export function useTraceFilePreview(
  name: string | null,
  options?: { direction?: string; limit?: number; enabled?: boolean }
) {
  return useQuery({
    queryKey: [...queryKeys.stream.traces.detail(name || ''), options],
    queryFn: () => {
      if (!name) throw new Error('No file selected');
      const params = new URLSearchParams();
      if (options?.direction) params.append('direction', options.direction);
      if (options?.limit) params.append('limit', options.limit.toString());
      const query = params.toString() ? `?${params.toString()}` : '';
      return apiClient.get<TraceFileResponse>(`/stream/traces/${encodeURIComponent(name)}${query}`);
    },
    enabled: options?.enabled ?? false,
  });
}

// Stream mutations - Traces

export function useDeleteTraceFile() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (name: string) =>
      apiClient.delete(`/stream/traces/${encodeURIComponent(name)}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.traces.all() });
    },
  });
}

export function useEditTrace() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: EditTraceRequest) =>
      apiClient.post('/stream/traces/edit', request),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.traces.all() });
    },
  });
}

export function useMergeTraces() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: MergeTracesRequest) =>
      apiClient.post('/stream/traces/merge', request),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.stream.traces.all() });
    },
  });
}

export function useExportTrace() {
  return useMutation({
    mutationFn: async (request: ExportTraceRequest) => {
      const response = await fetch('/api/v1/stream/traces/export', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(request),
      });

      if (!response.ok) {
        throw new Error('Export failed');
      }

      return await response.blob();
    },
  });
}

// Stream Editor queries and mutations

export function useEditorLoad() {
  return useMutation({
    mutationFn: async ({ file, name }: { file?: File; name?: string }) => {
      if (file) {
        const formData = new FormData();
        formData.append('file', file);
        const response = await fetch('/api/v1/stream/editor/load', {
          method: 'POST',
          body: formData,
        });
        if (!response.ok) throw new Error('Upload failed');
        return response.json() as Promise<EditorStats>;
      } else if (name) {
        return apiClient.post<EditorStats>('/stream/editor/load', { filename: name });
      }
      throw new Error('Either file or name must be provided');
    },
  });
}

export function useEditorMessages(
  sessionId: string | null,
  offset: number,
  limit: number,
  filters?: EditorFilters
) {
  return useQuery({
    queryKey: ['editor', 'messages', sessionId, offset, limit, filters],
    queryFn: () => {
      const params = new URLSearchParams({
        sessionId: sessionId!,
        offset: offset.toString(),
        limit: limit.toString(),
      });
      if (filters?.proxy) params.append('proxy', filters.proxy);
      if (filters?.interface) params.append('interface', filters.interface);
      if (filters?.direction) params.append('direction', filters.direction);
      if (filters?.type) params.append('type', filters.type);

      return apiClient.get<EditorMessagesResponse>(`/stream/editor/messages?${params}`);
    },
    enabled: !!sessionId,
  });
}

export function useEditorTimeline(sessionId: string | null) {
  return useQuery({
    queryKey: ['editor', 'timeline', sessionId],
    queryFn: () =>
      apiClient.get<EditorTimelineResponse>(`/stream/editor/timeline?sessionId=${sessionId}`),
    enabled: !!sessionId,
  });
}

export function useEditorSeek() {
  return useMutation({
    mutationFn: async ({
      sessionId,
      timestamp,
      filters,
    }: {
      sessionId: string;
      timestamp: number;
      filters?: EditorFilters;
    }) => {
      const params = new URLSearchParams({
        sessionId,
        timestamp: timestamp.toString(),
      });
      if (filters?.proxy) params.append('proxy', filters.proxy);
      if (filters?.interface) params.append('interface', filters.interface);
      if (filters?.direction) params.append('direction', filters.direction);
      if (filters?.type) params.append('type', filters.type);

      return apiClient.get<EditorSeekResponse>(`/stream/editor/seek?${params}`);
    },
  });
}

export function useEditorJQ() {
  return useMutation({
    mutationFn: async ({
      sessionId,
      query,
      limit = 100,
    }: {
      sessionId: string;
      query: string;
      limit?: number;
    }) => {
      return apiClient.post<EditorJQResponse>('/stream/editor/jq', {
        sessionId,
        query,
        limit,
      });
    },
  });
}

export function useEditorExport() {
  return useMutation({
    mutationFn: async ({ sessionId, indices }: { sessionId: string; indices?: number[] }) => {
      const response = await fetch('/api/v1/stream/editor/export', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ sessionId, indices }),
      });
      if (!response.ok) throw new Error('Export failed');
      return response.blob();
    },
  });
}
