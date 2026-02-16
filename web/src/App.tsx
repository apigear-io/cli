import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { AppLayout } from './components/Layout/AppLayout';
import { Dashboard } from './pages/Dashboard/Dashboard';
import { Templates } from './pages/Templates/Templates';
import { Projects } from './pages/Projects/Projects';
import { CodeGen } from './pages/CodeGen/CodeGen';
import { Monitor } from './pages/Monitor/Monitor';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<AppLayout />}>
          <Route index element={<Navigate to="/dashboard" replace />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="templates" element={<Templates />} />
          <Route path="projects" element={<Projects />} />
          <Route path="codegen" element={<CodeGen />} />
          <Route path="monitor" element={<Monitor />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
