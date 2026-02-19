import { Suspense } from 'react';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { ScriptingContent } from './components/ScriptingContent';

export function Scripting() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading scripting environment..." />}>
        <ScriptingContent />
      </Suspense>
    </ErrorBoundary>
  );
}
