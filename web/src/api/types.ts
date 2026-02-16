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
