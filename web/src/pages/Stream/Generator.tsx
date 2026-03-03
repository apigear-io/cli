import { Suspense } from 'react';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { GeneratorContent } from './components/GeneratorContent';

export function Generator() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading generator..." />}>
        <GeneratorContent />
      </Suspense>
    </ErrorBoundary>
  );
}
