export interface HealthResponse {
  status: string;
  timestamp: string;
}

export interface StatusResponse {
  version: string;
  commit: string;
  buildDate: string;
  goVersion: string;
  uptime: string;
}

export interface TemplateInfo {
  name: string;
  description: string;
  author: string;
  git: string;
  version: string;
  latest: string;
  versions: string[];
  inCache: boolean;
  inRegistry: boolean;
  tags?: string[];
  updateNeeded: boolean; // True if cached version < latest version (semver comparison)
}

export interface TemplateListResponse {
  templates: TemplateInfo[];
  count: number;
}

export interface InstallProgressEvent {
  type: 'progress' | 'complete' | 'error';
  message: string;
  progress: number;
  error?: string;
}

// Stream types

export type ProxyStatus = 'stopped' | 'running' | 'error';
export type ProxyMode = 'proxy' | 'echo' | 'backend' | 'inbound-only';
export type ClientStatus = 'disconnected' | 'connecting' | 'connected' | 'error';

export interface ProxyInfo {
  name: string;
  listen: string;
  backend: string;
  mode: ProxyMode;
  status: ProxyStatus;
  messagesReceived: number;
  messagesSent: number;
  activeConnections: number;
  bytesReceived: number;
  bytesSent: number;
  uptime: number; // seconds
}

export interface ProxyConfig {
  listen: string;
  backend?: string;
  mode: ProxyMode;
  disabled?: boolean;
}

export interface CreateProxyRequest {
  name: string;
  config: ProxyConfig;
}

export interface ClientInfo {
  name: string;
  url: string;
  interfaces: string[];
  status: ClientStatus;
  autoReconnect: boolean;
  enabled: boolean;
  lastError?: string;
}

export interface ClientConfig {
  url: string;
  interfaces: string[];
  enabled: boolean;
  autoReconnect: boolean;
}

export interface CreateClientRequest {
  name: string;
  config: ClientConfig;
}

export interface StreamDashboardStats {
  proxies: {
    total: number;
    running: number;
    stopped: number;
  };
  clients: {
    total: number;
    connected: number;
    disconnected: number;
  };
  messages: {
    total: number;
    rate: number;
  };
}
