import { Suspense } from 'react';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { TracesContent } from './components/TracesContent';

export function Traces() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading traces..." />}>
        <TracesContent />
      </Suspense>
    </ErrorBoundary>
  );
}
