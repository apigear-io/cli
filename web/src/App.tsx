import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { CodeGenAppShell } from './components/Layout/CodeGenAppShell';
import { StreamAppShell } from './components/Layout/StreamAppShell';
import { Dashboard } from './pages/Dashboard/Dashboard';
import { Templates } from './pages/Templates/Templates';
import { Projects } from './pages/Projects/Projects';
import { ProjectDetail } from './pages/Projects/ProjectDetail';
import { CodeGeneration } from './pages/Projects/CodeGeneration';
import { CodeGen } from './pages/CodeGen/CodeGen';
import { Monitor } from './pages/Monitor/Monitor';
import { StreamDashboard } from './pages/Stream/Dashboard';
import { Proxies } from './pages/Stream/Proxies';
import { ProxyDetail } from './pages/Stream/ProxyDetail';
import { Clients } from './pages/Stream/Clients';
import { Scripting } from './pages/Stream/Scripting';
import { Traces } from './pages/Stream/Traces';
import { StreamEditor } from './pages/Stream/StreamEditor';
import { Player } from './pages/Stream/Player';
import { Logs } from './pages/Stream/Logs';
import { Generator } from './pages/Stream/Generator';
import { Settings } from './pages/Stream/Settings';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* Root redirect */}
        <Route path="/" element={<Navigate to="/codegen/dashboard" replace />} />

        {/* CodeGen routes with CodeGen AppShell */}
        <Route path="/codegen" element={<CodeGenAppShell />}>
          <Route index element={<Navigate to="/codegen/dashboard" replace />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="templates" element={<Templates />} />
          <Route path="projects" element={<Projects />} />
          <Route path="projects/:encodedPath" element={<ProjectDetail />} />
          <Route path="projects/generate/:encodedSolutionPath" element={<CodeGeneration />} />
          <Route path="generate" element={<CodeGen />} />
          <Route path="monitor" element={<Monitor />} />
        </Route>

        {/* Stream routes with Stream AppShell */}
        <Route path="/stream" element={<StreamAppShell />}>
          <Route index element={<Navigate to="/stream/dashboard" replace />} />
          <Route path="dashboard" element={<StreamDashboard />} />
          <Route path="proxies" element={<Proxies />} />
          <Route path="proxies/:name" element={<ProxyDetail />} />
          <Route path="clients" element={<Clients />} />
          <Route path="scripting" element={<Scripting />} />
          <Route path="traces" element={<Traces />} />
          <Route path="editor" element={<StreamEditor />} />
          <Route path="player" element={<Player />} />
          <Route path="generator" element={<Generator />} />
          <Route path="logs" element={<Logs />} />
          <Route path="settings" element={<Settings />} />
        </Route>

        {/* Fallback for old routes - redirect to new structure */}
        <Route path="/dashboard" element={<Navigate to="/codegen/dashboard" replace />} />
        <Route path="/templates" element={<Navigate to="/codegen/templates" replace />} />
        <Route path="/projects" element={<Navigate to="/codegen/projects" replace />} />
        <Route path="/monitor" element={<Navigate to="/codegen/monitor" replace />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
