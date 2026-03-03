import { Suspense } from 'react';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { LogsContent } from './components/LogsContent';

export function Logs() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading logs..." />}>
        <LogsContent />
      </Suspense>
    </ErrorBoundary>
  );
}
