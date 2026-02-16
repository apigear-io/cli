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
