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

// Script types

export type ScriptType = 'client' | 'backend';
export type ScriptOutputLevel = 'log' | 'info' | 'warn' | 'error' | 'debug';

export interface ScriptFileInfo {
  name: string;
  modTime: number;
}

export interface ScriptInfo {
  id: string;
  name: string;
  type: ScriptType;
}

export interface ScriptFile {
  name: string;
  code: string;
  modTime: number;
}

export interface ScriptOutputEntry {
  level: ScriptOutputLevel;
  message: string;
}

export interface SaveScriptRequest {
  name: string;
  code: string;
  expectedModTime?: number;
}

export interface SaveScriptResponse {
  name: string;
  modTime: number;
  message: string;
}

export interface RunScriptResponse {
  id: string;
  name: string;
  message: string;
}

export interface RunCodeRequest {
  name?: string;
  code: string;
}

// Trace types

export interface TraceFileInfo {
  name: string;
  path: string;
  size: number;
  modTime: string;
  proxyName: string;
}

export interface TraceEntry {
  ts: number; // Timestamp in milliseconds
  dir: string; // Direction: "SEND" or "RECV"
  proxy: string;
  msg: unknown; // Raw JSON message
}

export interface TraceStats {
  fileCount: number;
  totalBytes: number;
  totalMB: number;
  traceDir: string;
}

export interface TraceFileResponse {
  filename: string;
  entries: TraceEntry[];
  count: number;
}

export interface SearchTracesRequest {
  proxyName?: string;
  direction?: string;
  startTime?: number;
  endTime?: number;
  maxFiles?: number;
  maxEntries?: number;
}

// Live Message types

export interface ParsedMessage {
  msgType: number;
  msgTypeName: string;
  symbol?: string;
  objectId?: string;
  requestId?: number;
  args?: unknown;
}

export interface ParsedMessageEvent {
  type: string;
  proxy: string;
  direction: string; // "SEND" or "RECV"
  timestamp: number;
  message: unknown; // Raw JSON message
  parsed?: ParsedMessage;
}
