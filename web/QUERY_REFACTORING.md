# Query Refactoring - useSuspenseQuery Migration

This document describes the migration from `useQuery` to `useSuspenseQuery` in the web UI.

## Changes Made

### 1. Query Key Factory (`src/api/queryKeys.ts`)

Created a centralized query key factory for type-safe and consistent query keys:

```typescript
export const queryKeys = {
  health: () => ['health'] as const,
  status: () => ['status'] as const,
  templates: {
    all: () => ['templates'] as const,
    registry: () => [...queryKeys.templates.all(), 'registry'] as const,
    cache: () => [...queryKeys.templates.all(), 'cache'] as const,
    detail: (id: string) => [...queryKeys.templates.all(), 'detail', id] as const,
    search: (query: string) => [...queryKeys.templates.all(), 'search', query] as const,
  },
} as const;
```

**Benefits:**
- Type-safe query keys
- Prevents typos and inconsistencies
- Easy to invalidate related queries (e.g., all templates with `queryKeys.templates.all()`)
- Self-documenting query structure

### 2. Migrated to useSuspenseQuery (`src/api/queries.ts`)

**Before:**
```typescript
export function useTemplates() {
  return useQuery({
    queryKey: ['templates'],
    queryFn: () => apiClient.get<TemplateListResponse>('/templates'),
  });
}
```

**After:**
```typescript
export function useTemplates() {
  return useSuspenseQuery({
    queryKey: queryKeys.templates.registry(),
    queryFn: () => apiClient.get<TemplateListResponse>('/templates'),
  });
}
```

**Benefits:**
- Data is guaranteed to be defined (no optional chaining needed)
- Better TypeScript inference
- Loading states handled by `<Suspense>`
- Error states handled by Error Boundaries

### 3. Error Boundary Component (`src/components/ErrorBoundary.tsx`)

Created a reusable error boundary component for centralized error handling:

```typescript
<ErrorBoundary>
  <YourComponent />
</ErrorBoundary>
```

### 4. Loading Fallback Component (`src/components/LoadingFallback.tsx`)

Created a consistent loading fallback component:

```typescript
<Suspense fallback={<LoadingFallback message="Loading..." />}>
  <YourComponent />
</Suspense>
```

### 5. Updated Components

#### Templates.tsx
**Before:**
```typescript
const { data, isLoading, error } = useTemplates();

if (isLoading) return <Loader />;
if (error) return <Alert>Error</Alert>;
if (!data?.templates) return null;

return <div>{data.templates.map(...)}</div>;
```

**After:**
```typescript
function TemplatesContent() {
  const { data } = useTemplates();
  // data.templates is guaranteed to exist!
  return <div>{data.templates.map(...)}</div>;
}

export function Templates() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback />}>
        <TemplatesContent />
      </Suspense>
    </ErrorBoundary>
  );
}
```

#### Child Components
Removed `isLoading` props from:
- `RegistryTemplateList`
- `CachedTemplateList`

Loading states are now handled at the parent level via Suspense.

### 6. Updated Test Utilities

Enhanced test utilities to support Suspense and Error Boundaries:

```typescript
// Test utilities now wrap components in Suspense automatically
const customRender = (ui, options) => {
  return render(ui, {
    wrapper: ({ children }) => (
      <AllProviders>
        <ErrorBoundary>
          <Suspense fallback={<div>Loading...</div>}>
            {children}
          </Suspense>
        </ErrorBoundary>
      </AllProviders>
    ),
  });
};
```

## Benefits Summary

### Code Quality
✅ **Cleaner components** - No manual loading/error handling
✅ **Better TypeScript** - Data is always defined
✅ **Less boilerplate** - No optional chaining (`data?.field`)
✅ **DRY principle** - Centralized loading/error states

### Developer Experience
✅ **Type-safe query keys** - Autocomplete and refactoring support
✅ **Easier testing** - Consistent wrapper setup
✅ **Better maintainability** - Single source of truth for query keys

### User Experience
✅ **Coordinated loading** - Multiple queries suspend together
✅ **Consistent UI** - Standardized loading/error states
✅ **Better error recovery** - Error boundaries with retry logic

## Migration Guide for Other Components

To migrate a component to use `useSuspenseQuery`:

1. Update the query hook import:
   ```typescript
   - import { useQuery } from '@tanstack/react-query';
   + import { useSuspenseQuery } from '@tanstack/react-query';
   ```

2. Use the query key factory:
   ```typescript
   - queryKey: ['myResource', id]
   + queryKey: queryKeys.myResource.detail(id)
   ```

3. Remove optional chaining:
   ```typescript
   - const items = data?.items ?? []
   + const items = data.items
   ```

4. Remove manual loading/error handling:
   ```typescript
   - if (isLoading) return <Loader />
   - if (error) return <Alert>Error</Alert>
   ```

5. Wrap the component in Suspense and ErrorBoundary:
   ```typescript
   export function MyFeature() {
     return (
       <ErrorBoundary>
         <Suspense fallback={<LoadingFallback />}>
           <MyFeatureContent />
         </Suspense>
       </ErrorBoundary>
     );
   }
   ```

## Testing

All existing tests continue to pass with the new setup. The test utilities automatically handle Suspense and Error Boundaries.

Run tests:
```bash
pnpm test              # Unit tests
task web:test          # Via task runner
```

## Next Steps

Consider migrating other pages to use this pattern:
- Dashboard
- Projects
- CodeGen
- Monitor

Each migration will further reduce code complexity and improve consistency.
