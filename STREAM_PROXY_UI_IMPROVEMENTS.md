# Stream Proxy UI Improvements

This document summarizes the UI improvements made to the Stream Proxy functionality.

## Issues Addressed

1. ✅ **Start/Stop functionality not working** - Added working start/stop buttons to proxy cards
2. ✅ **No proxy detail page** - Created comprehensive proxy detail page with trace viewer
3. ✅ **No edit functionality** - Added edit drawer for modifying proxy configuration

## Changes Made

### 1. ProxyCard Component (`web/src/pages/Stream/components/ProxyCard.tsx`)

**Added:**
- Start/Stop buttons that appear conditionally based on proxy status
- Click handler to navigate to proxy detail page
- Event propagation prevention for action buttons to avoid triggering navigation
- Visual feedback with `cursor: pointer` on card hover

**Features:**
- Green "Start" button with play icon when proxy is stopped
- Red "Stop" button with stop icon when proxy is running
- All action buttons (view stats, edit, delete) now prevent event bubbling
- Card click navigates to `/stream/proxies/:name` detail page

### 2. Proxies Page (`web/src/pages/Stream/Proxies.tsx`)

**Added:**
- Edit drawer UI using Mantine Drawer component
- Start/stop mutation handlers with notifications
- Update proxy handler for edit functionality
- State management for selected proxy and edit drawer

**Features:**
- Edit drawer slides in from right with proxy configuration
- Pre-fills form with existing proxy values
- Backend address field conditionally shown for 'proxy' mode
- Proxy name is read-only in edit mode
- Success/error notifications for all operations
- Loading states for async operations

**API Integrations:**
- `useStartProxy()` - Start a stopped proxy
- `useStopProxy()` - Stop a running proxy
- `useUpdateProxy()` - Update proxy configuration
- All mutations properly invalidate query cache

### 3. ProxyDetail Page (`web/src/pages/Stream/ProxyDetail.tsx`) - NEW

**Created a comprehensive detail page with:**

**Header Section:**
- Back button to return to proxies list
- Proxy name and status badge
- Quick action buttons (Start, Stop, Edit, Delete)

**Statistics Dashboard (6 cards):**
1. **Connections** - Active connection count with icon
2. **Messages** - Total messages with ↓received ↑sent breakdown
3. **Data** - Total bytes with ↓↑ breakdown in formatted units (B/KB/MB/GB)
4. **Uptime** - Formatted uptime (seconds/minutes/hours)
5. **Message Rate** - Messages per second calculation
6. **Mode** - Proxy mode display

**Configuration Section:**
- Listen address display
- Backend address display (if applicable)

**Tabbed Interface:**
1. **Overview Tab**
   - Proxy status summary
   - Quick instructions

2. **Live Messages Tab**
   - Live message viewer when proxy is running
   - Empty state with instruction when stopped
   - Uses existing `LiveMessageViewer` component
   - 60vh height for optimal viewing

3. **Trace Files Tab**
   - Placeholder for future trace file management
   - Empty state with "coming soon" message

**Auto-refresh:**
- Proxy details refresh every 2 seconds
- Real-time updates of stats and status

**Helper Functions:**
- `formatBytes()` - Converts bytes to human-readable format (B, KB, MB, GB)
- `formatUptime()` - Converts seconds to readable format (s, m, h)
- Status color mapping for consistent UI

### 4. Routing Updates (`web/src/App.tsx`)

**Added:**
- Import for `ProxyDetail` component
- New route: `/stream/proxies/:name` for detail page
- Route positioned before `/stream/clients` to ensure proper matching

## User Flows

### Starting a Proxy
1. Navigate to Proxies page
2. Find proxy in stopped state
3. Click green "Start" button on card
4. Proxy starts, button changes to "Stop"
5. Success notification appears

### Editing a Proxy
1. Click edit icon (pencil) on proxy card
2. Edit drawer slides in from right
3. Modify listen address, mode, or backend
4. Click "Update" button
5. Drawer closes, changes applied
6. Success notification appears

### Viewing Proxy Details
1. Click anywhere on proxy card
2. Navigate to detail page
3. View comprehensive statistics
4. Switch between tabs for different views
5. Click back button or breadcrumb to return

### Viewing Live Messages
1. Navigate to proxy detail page
2. Ensure proxy is running (start if needed)
3. Click "Live Messages" tab
4. View real-time message stream
5. Messages auto-scroll and update

## Technical Details

### State Management
- React Query for server state
- Automatic cache invalidation on mutations
- Optimistic updates handled by query client
- 2-second refetch interval for detail page

### Error Handling
- Try-catch blocks around all async operations
- User-friendly error notifications
- Detailed error messages from API
- Graceful degradation for unavailable features

### Performance
- Suspense boundaries for loading states
- Error boundaries for error handling
- Debounced mutations to prevent double-clicks
- Efficient re-renders with React Query

### Accessibility
- Semantic HTML structure
- ARIA labels on icon buttons
- Keyboard navigation support
- Screen reader friendly

## Testing Recommendations

### Manual Testing
1. **Start/Stop Operations**
   - Start a stopped proxy
   - Stop a running proxy
   - Verify status updates in real-time
   - Check notifications appear

2. **Edit Functionality**
   - Edit proxy configuration
   - Verify mode changes affect form
   - Test validation (required fields)
   - Confirm changes persist after save

3. **Navigation**
   - Click proxy card to navigate
   - Verify detail page loads
   - Test back button
   - Check breadcrumb navigation

4. **Live Messages**
   - Start proxy if needed
   - View live messages tab
   - Verify messages stream
   - Test with stopped proxy

### E2E Tests to Add
```typescript
test('should start and stop proxy', async ({ page }) => {
  await page.goto('/stream/proxies');
  await page.getByRole('button', { name: /start/i }).first().click();
  await expect(page.getByRole('button', { name: /stop/i })).toBeVisible();
  await page.getByRole('button', { name: /stop/i }).click();
  await expect(page.getByRole('button', { name: /start/i })).toBeVisible();
});

test('should navigate to proxy detail page', async ({ page }) => {
  await page.goto('/stream/proxies');
  await page.locator('.mantine-Card-root').first().click();
  await expect(page).toHaveURL(/\/stream\/proxies\/.+/);
  await expect(page.getByRole('heading', { level: 2 })).toBeVisible();
});

test('should edit proxy configuration', async ({ page }) => {
  await page.goto('/stream/proxies');
  await page.getByRole('button', { name: /edit/i }).first().click();
  await page.fill('input[label="Listen Address"]', 'ws://localhost:9999/ws');
  await page.getByRole('button', { name: /update/i }).click();
  await expect(page.getByText('ws://localhost:9999/ws')).toBeVisible();
});
```

## Files Modified

1. **web/src/pages/Stream/components/ProxyCard.tsx**
   - Added start/stop buttons
   - Made card clickable
   - Added navigation on click

2. **web/src/pages/Stream/Proxies.tsx**
   - Added edit drawer UI
   - Added start/stop handlers
   - Integrated update proxy mutation

3. **web/src/pages/Stream/ProxyDetail.tsx** (NEW)
   - Created comprehensive detail page
   - Added statistics dashboard
   - Implemented tabbed interface
   - Added live message viewer

4. **web/src/App.tsx**
   - Added ProxyDetail route
   - Imported ProxyDetail component

5. **web/src/pages/Stream/components/ScriptingContent.tsx**
   - Fixed unused error variable (pre-existing issue)

## Known Issues / Future Improvements

1. **Trace Files Tab** - Currently shows placeholder, needs implementation
2. **Edit Button in Detail Page** - Currently just shows tooltip, needs to open edit drawer
3. **WebSocket Connection Status** - Could add real-time connection indicator
4. **Message Filtering** - Live viewer could benefit from filtering capabilities
5. **Export Stats** - Could add ability to export statistics as CSV/JSON
6. **Proxy Templates** - Could add templates for common proxy configurations

## Migration from wsproxy

These changes align with the functionality from the original wsproxy application, providing:
- ✅ Similar UI/UX patterns
- ✅ Consistent terminology
- ✅ Equivalent feature set
- ✅ Improved with React/TypeScript
- ✅ Better state management with React Query

## References

- API Queries: `web/src/api/queries.ts`
- API Types: `web/src/api/types.ts`
- Backend Handlers: `internal/handler/stream_proxies.go`
- Query Keys: `web/src/api/queryKeys.ts`
