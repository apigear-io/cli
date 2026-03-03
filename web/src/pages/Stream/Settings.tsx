import { Suspense } from 'react';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { SettingsContent } from './components/SettingsContent';

export function Settings() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading settings..." />}>
        <SettingsContent />
      </Suspense>
    </ErrorBoundary>
  );
}
