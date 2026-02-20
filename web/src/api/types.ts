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

// Stream Editor types

export interface EditTraceRequest {
  sourceFile: string;
  outputFile: string;
  direction?: string;
  startTime?: number;
  endTime?: number;
  proxyNames?: string[];
  messageTypes?: number[];
  containsText?: string;
  normalizeTime?: boolean;
  remapProxyName?: string;
  timestampOffset?: number;
}

export interface MergeTracesRequest {
  sourceFiles: string[];
  outputFile: string;
  sortByTime?: boolean;
  normalize?: boolean;
}

export interface ExportTraceRequest {
  sourceFile: string;
  format: 'json' | 'jsonl';
  direction?: string;
  startTime?: number;
  endTime?: number;
  limit?: number;
}

// Stream Editor types

export interface EditorMessage {
  index: number;
  timestamp: number;
  direction: string;
  proxy: string;
  raw: Record<string, unknown>;
  parsed: ParsedObjectLink;
}

export interface ParsedObjectLink {
  msgType: number;
  msgTypeName: string;
  symbol?: string;
  objectId?: string;
  requestId?: number;
  args?: unknown;
}

export interface EditorStats {
  sessionId: string;
  filename: string;
  totalCount: number;
  timeRange: {
    start: number;
    end: number;
  };
  proxies: string[];
  interfaces: string[];
}

export interface EditorMessagesResponse {
  messages: EditorMessage[];
  total: number;
  offset: number;
  limit: number;
}

export interface EditorBucket {
  startTime: number;
  endTime: number;
  sendCount: number;
  recvCount: number;
}

export interface EditorTimelineResponse {
  buckets: EditorBucket[];
  timeRange: {
    start: number;
    end: number;
  };
}

export interface EditorFilters {
  proxy?: string;
  interface?: string;
  direction?: string;
  type?: string;
}

export interface EditorLoadRequest {
  filename?: string;
}

export interface EditorSeekResponse {
  offset: number;
  messageIndex: number;
}

export interface EditorJQMatch {
  index: number;
  result: unknown;
}

export interface EditorJQResponse {
  matches: EditorJQMatch[];
  totalMatches: number;
}

export interface EditorExportRequest {
  sessionId: string;
  indices?: number[];
}

// Stream Player types

export type PlayerState = 'stopped' | 'playing' | 'paused';

export interface PlayerStream {
  id: string;
  proxyName: string;
  filename: string;
  speed: number;
  loop: boolean;
  direction: string; // "", "SEND", "RECV"
  state: PlayerState;
  position: number;
  totalEntries: number;
  progress: number; // 0.0 to 1.0
}

export interface CreatePlayerStreamRequest {
  proxyName: string;
  filename: string;
  speed: number;
  initialDelay: number; // ms
  loop: boolean;
  direction: string; // "", "SEND", "RECV"
}

// Application Logs types

export type LogLevel = 'DEBUG' | 'INFO' | 'WARN' | 'ERROR';

export interface LogEntry {
  timestamp: string;
  level: LogLevel;
  message: string;
  fields?: Record<string, unknown>;
}

export interface LogsResponse {
  entries: LogEntry[];
  count: number;
}

// Trace Generator types

export interface GenerateRequest {
  template: string;
  count: number;
}

export interface GenerateResult {
  entries: unknown[];
  count: number;
}

export interface GeneratorSaveRequest {
  template: string;
  count: number;
  proxyName: string;
  filename: string;
}

export interface GeneratorSaveResponse {
  filename: string;
  count: number;
}

export interface GeneratorSaveTemplateRequest {
  name: string;
  template: string;
}

export interface GeneratorLoadTemplateResponse {
  name: string;
  template: string;
}

export interface GeneratorListTemplatesResponse {
  templates: string[];
}
