import { Suspense } from 'react';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { PlayerContent } from './components/PlayerContent';

export function Player() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading player..." />}>
        <PlayerContent />
      </Suspense>
    </ErrorBoundary>
  );
}
