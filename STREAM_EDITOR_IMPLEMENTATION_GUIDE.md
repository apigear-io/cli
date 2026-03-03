# Stream Editor Implementation Guide

## Overview

This guide provides step-by-step instructions to complete the Stream Editor feature. The backend is **100% complete** and working. Frontend foundation is in place. This guide covers the remaining ~8 frontend components.

## What's Complete ✅

### Backend (100% Done)
- **File**: `internal/handler/stream_editor.go`
- Session management with 30min TTL
- All 6 API endpoints working:
  - `POST /api/v1/stream/editor/load` - Load trace (upload or server file)
  - `GET /api/v1/stream/editor/messages?sessionId=...&offset=0&limit=100` - Paginated messages
  - `GET /api/v1/stream/editor/timeline?sessionId=...` - 200 time buckets
  - `GET /api/v1/stream/editor/seek?sessionId=...&timestamp=...` - Jump to timestamp
  - `POST /api/v1/stream/editor/export` - Export selected indices
  - `POST /api/v1/stream/editor/jq` - JQ query (using gojq library)
- Routes registered in `internal/handler/router.go`
- Compiles successfully

### Frontend Foundation (70% Done)
- **Types**: `web/src/api/types.ts` - All editor types added
- **Context**: `web/src/pages/Stream/components/EditorContext.tsx` - State management
- **Welcome**: `web/src/pages/Stream/components/EditorWelcome.tsx` - Landing screen
- **Page**: `web/src/pages/Stream/StreamEditor.tsx` - Main container
- **Route**: Added to `web/src/App.tsx` at `/stream/editor`
- **Navigation**: Added to sidebar

## What's Remaining (8 Components + Hooks)

1. EditorLoadDrawer - File upload UI
2. Editor API hooks - React Query integration
3. EditorStats - Stats bar
4. EditorTimeline - Canvas visualization
5. EditorFilters - Filter controls
6. EditorJQPanel - JQ query interface
7. EditorTable - Virtual scrolling table
8. EditorToolbar - Action buttons
9. Keyboard shortcuts
10. Integration & Testing

---

## Component 1: EditorLoadDrawer

**File**: `web/src/pages/Stream/components/EditorLoadDrawer.tsx`

**Purpose**: Modal/Drawer for loading trace files - either upload or select from server.

**Reference**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/editor/components/EditorLoadDrawer.tsx`

**Dependencies**:
```typescript
import { Drawer, Stack, FileButton, Button, Text, ScrollArea, Paper, Group, Badge } from '@mantine/core';
import { IconUpload, IconFileText, IconRefresh } from '@tabler/icons-react';
import { useTraceFiles } from '@/api/queries'; // Already exists
import { useEditorLoad } from '@/api/queries'; // Need to create
```

**Key Features**:
1. Drag & drop zone for file upload
2. List of recent trace files from server (use `useTraceFiles()`)
3. Show file size and date
4. On selection, call load API and update context

**Structure**:
```typescript
interface EditorLoadDrawerProps {
  opened: boolean;
  onClose: () => void;
}

export function EditorLoadDrawer({ opened, onClose }: EditorLoadDrawerProps) {
  const { data: traceFiles } = useTraceFiles();
  const loadMutation = useEditorLoad();
  const { setSessionStats } = useEditorContext();

  const handleFileUpload = async (file: File) => {
    const stats = await loadMutation.mutateAsync({ file });
    setSessionStats(stats);
    onClose();
  };

  const handleServerFile = async (filename: string) => {
    const stats = await loadMutation.mutateAsync({ name: filename });
    setSessionStats(stats);
    onClose();
  };

  return (
    <Drawer opened={opened} onClose={onClose} title="Set Stream" size="lg">
      <Stack gap="lg">
        {/* Upload section */}
        <Paper withBorder p="xl" style={{ borderStyle: 'dashed' }}>
          <Stack align="center" gap="md">
            <IconUpload size={48} />
            <Text>Drop JSONL log file here or click to browse</Text>
            <FileButton onChange={handleFileUpload} accept=".jsonl,.jsonl.gz">
              {(props) => <Button {...props}>Browse</Button>}
            </FileButton>
          </Stack>
        </Paper>

        {/* Divider */}
        <Text c="dimmed" ta="center">OR SELECT FROM TRACES</Text>

        {/* Server files list */}
        <Stack gap="xs">
          <Group justify="space-between">
            <Text fw={500}>Recent Log Files</Text>
            <Button variant="subtle" size="xs" leftSection={<IconRefresh size={14} />}>
              Refresh
            </Button>
          </Group>
          <ScrollArea h={400}>
            {traceFiles.map((file) => (
              <Paper
                key={file.name}
                p="md"
                mb="xs"
                withBorder
                style={{ cursor: 'pointer' }}
                onClick={() => handleServerFile(file.name)}
              >
                <Group justify="space-between">
                  <Group>
                    <IconFileText size={20} />
                    <Stack gap={0}>
                      <Text size="sm" fw={500}>{file.name}</Text>
                      <Text size="xs" c="dimmed">{file.proxyName}</Text>
                    </Stack>
                  </Group>
                  <Badge>{(file.size / 1024).toFixed(1)} KB</Badge>
                </Group>
              </Paper>
            ))}
          </ScrollArea>
        </Stack>
      </Stack>
    </Drawer>
  );
}
```

**Integration**: Add to `StreamEditor.tsx`:
```typescript
const [loadDrawerOpen, setLoadDrawerOpen] = useState(false);

// In EditorContent:
<Button onClick={() => setLoadDrawerOpen(true)}>Set Stream</Button>
<EditorLoadDrawer opened={loadDrawerOpen} onClose={() => setLoadDrawerOpen(false)} />
```

---

## Component 2: Editor API Hooks

**File**: `web/src/api/queries.ts` (add to existing file)

**Add these hooks**:

```typescript
import type {
  EditorStats,
  EditorLoadRequest,
  EditorMessagesResponse,
  EditorTimelineResponse,
  EditorSeekResponse,
  EditorJQResponse,
  EditorFilters,
} from './types';

// Load trace file
export function useEditorLoad() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({ file, name }: { file?: File; name?: string }) => {
      if (file) {
        const formData = new FormData();
        formData.append('file', file);
        const response = await fetch('/api/v1/stream/editor/load', {
          method: 'POST',
          body: formData,
        });
        if (!response.ok) throw new Error('Upload failed');
        return response.json() as Promise<EditorStats>;
      } else if (name) {
        return apiClient.post<EditorStats>('/stream/editor/load', { filename: name });
      }
      throw new Error('Either file or name must be provided');
    },
  });
}

// Get paginated messages
export function useEditorMessages(
  sessionId: string | null,
  offset: number,
  limit: number,
  filters?: EditorFilters
) {
  return useQuery({
    queryKey: ['editor', 'messages', sessionId, offset, limit, filters],
    queryFn: () => {
      const params = new URLSearchParams({
        sessionId: sessionId!,
        offset: offset.toString(),
        limit: limit.toString(),
      });
      if (filters?.proxy) params.append('proxy', filters.proxy);
      if (filters?.interface) params.append('interface', filters.interface);
      if (filters?.direction) params.append('direction', filters.direction);
      if (filters?.type) params.append('type', filters.type);

      return apiClient.get<EditorMessagesResponse>(`/stream/editor/messages?${params}`);
    },
    enabled: !!sessionId,
  });
}

// Get timeline buckets
export function useEditorTimeline(sessionId: string | null) {
  return useQuery({
    queryKey: ['editor', 'timeline', sessionId],
    queryFn: () =>
      apiClient.get<EditorTimelineResponse>(`/stream/editor/timeline?sessionId=${sessionId}`),
    enabled: !!sessionId,
  });
}

// Seek to timestamp
export function useEditorSeek() {
  return useMutation({
    mutationFn: async ({
      sessionId,
      timestamp,
      filters,
    }: {
      sessionId: string;
      timestamp: number;
      filters?: EditorFilters;
    }) => {
      const params = new URLSearchParams({
        sessionId,
        timestamp: timestamp.toString(),
      });
      if (filters?.proxy) params.append('proxy', filters.proxy);
      if (filters?.interface) params.append('interface', filters.interface);
      if (filters?.direction) params.append('direction', filters.direction);
      if (filters?.type) params.append('type', filters.type);

      return apiClient.get<EditorSeekResponse>(`/stream/editor/seek?${params}`);
    },
  });
}

// Run JQ query
export function useEditorJQ() {
  return useMutation({
    mutationFn: async ({
      sessionId,
      query,
      limit = 100,
    }: {
      sessionId: string;
      query: string;
      limit?: number;
    }) => {
      return apiClient.post<EditorJQResponse>('/stream/editor/jq', {
        sessionId,
        query,
        limit,
      });
    },
  });
}

// Export messages
export function useEditorExport() {
  return useMutation({
    mutationFn: async ({ sessionId, indices }: { sessionId: string; indices?: number[] }) => {
      const response = await fetch('/api/v1/stream/editor/export', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ sessionId, indices }),
      });
      if (!response.ok) throw new Error('Export failed');
      return response.blob();
    },
  });
}
```

---

## Component 3: EditorStats

**File**: `web/src/pages/Stream/components/EditorStats.tsx`

**Purpose**: Display session statistics (filename, message count, time range, proxies, interfaces)

**Reference**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/editor/components/EditorStats.tsx`

**Structure**:
```typescript
import { Paper, Group, Stack, Text, Badge } from '@mantine/core';
import { useEditorContext } from './EditorContext';

export function EditorStats() {
  const { sessionStats } = useEditorContext();

  if (!sessionStats) return null;

  const formatTimestamp = (ts: number) => {
    return new Date(ts).toLocaleString();
  };

  return (
    <Paper p="md" withBorder>
      <Group justify="space-between" wrap="wrap">
        <Stack gap={4}>
          <Text size="xs" c="dimmed">File</Text>
          <Text fw={500}>{sessionStats.filename}</Text>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">Messages</Text>
          <Text fw={500}>{sessionStats.totalCount.toLocaleString()}</Text>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">Time Range</Text>
          <Text fw={500} size="sm">
            {formatTimestamp(sessionStats.timeRange.start)} -{' '}
            {formatTimestamp(sessionStats.timeRange.end)}
          </Text>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">Proxies</Text>
          <Group gap="xs">
            {sessionStats.proxies.map((p) => (
              <Badge key={p} size="sm" variant="light">
                {p}
              </Badge>
            ))}
          </Group>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">Interfaces</Text>
          <Group gap="xs">
            {sessionStats.interfaces.slice(0, 3).map((i) => (
              <Badge key={i} size="sm" variant="light" color="green">
                {i}
              </Badge>
            ))}
            {sessionStats.interfaces.length > 3 && (
              <Badge size="sm" variant="light" color="gray">
                +{sessionStats.interfaces.length - 3}
              </Badge>
            )}
          </Group>
        </Stack>
      </Group>
    </Paper>
  );
}
```

---

## Component 4: EditorTimeline (Canvas)

**File**: `web/src/pages/Stream/components/EditorTimeline.tsx`

**Purpose**: Canvas-based timeline with 200 buckets, click to jump, drag to select range.

**Reference**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/editor/components/EditorTimeline.tsx` and `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/editor/utils/timeline.ts`

**This is the most complex component**. Key features:
- 200 time buckets (blue bars for SEND, green for RECV)
- Gold triangle flags for marked messages
- Click to seek to timestamp
- Drag to select time range
- Hover to highlight bucket

**Simplified Structure**:
```typescript
import { useRef, useEffect, useState, MouseEvent } from 'react';
import { Box, Group, Text, Button } from '@mantine/core';
import { IconX } from '@tabler/icons-react';
import { useEditorContext } from './EditorContext';
import { useEditorTimeline } from '@/api/queries';

const TIMELINE_HEIGHT = 80;
const NUM_BUCKETS = 200;

export function EditorTimeline() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [canvasWidth, setCanvasWidth] = useState(800);
  const { sessionStats, markedIndices, timelineSelection, setTimelineSelection } = useEditorContext();

  const { data: timelineData } = useEditorTimeline(sessionStats?.sessionId || null);

  const [hoveredBucket, setHoveredBucket] = useState<number | null>(null);
  const [isDragging, setIsDragging] = useState(false);
  const [dragStart, setDragStart] = useState<number | null>(null);

  // Draw canvas
  useEffect(() => {
    if (!canvasRef.current || !timelineData) return;

    const canvas = canvasRef.current;
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    // Set canvas size
    canvas.width = canvasWidth;
    canvas.height = TIMELINE_HEIGHT;

    // Clear
    ctx.fillStyle = '#f8f9fa';
    ctx.fillRect(0, 0, canvasWidth, TIMELINE_HEIGHT);

    // Draw buckets
    const bucketWidth = canvasWidth / NUM_BUCKETS;
    const maxCount = Math.max(...timelineData.buckets.map(b => b.sendCount + b.recvCount), 1);

    timelineData.buckets.forEach((bucket, i) => {
      const x = i * bucketWidth;
      const sendHeight = (bucket.sendCount / maxCount) * (TIMELINE_HEIGHT - 20);
      const recvHeight = (bucket.recvCount / maxCount) * (TIMELINE_HEIGHT - 20);

      // Draw SEND (blue)
      ctx.fillStyle = '#228be6';
      ctx.fillRect(x, TIMELINE_HEIGHT - sendHeight - 10, bucketWidth - 1, sendHeight);

      // Draw RECV (green) on top
      ctx.fillStyle = '#40c057';
      ctx.fillRect(x, TIMELINE_HEIGHT - sendHeight - recvHeight - 10, bucketWidth - 1, recvHeight);

      // Highlight hovered bucket
      if (i === hoveredBucket) {
        ctx.strokeStyle = '#000';
        ctx.lineWidth = 2;
        ctx.strokeRect(x, 0, bucketWidth, TIMELINE_HEIGHT);
      }
    });

    // Draw selection range
    if (timelineSelection) {
      const startX = timelineSelection.start * bucketWidth;
      const endX = (timelineSelection.end + 1) * bucketWidth;
      ctx.fillStyle = 'rgba(34, 139, 230, 0.2)';
      ctx.fillRect(startX, 0, endX - startX, TIMELINE_HEIGHT);
    }

    // TODO: Draw gold flags for marked messages (requires mapping indices to buckets)

  }, [timelineData, canvasWidth, hoveredBucket, timelineSelection, markedIndices]);

  const handleMouseMove = (e: MouseEvent<HTMLCanvasElement>) => {
    if (!canvasRef.current || !timelineData) return;
    const rect = canvasRef.current.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const bucketIndex = Math.floor((x / canvasWidth) * NUM_BUCKETS);

    if (isDragging && dragStart !== null) {
      setTimelineSelection({
        start: Math.min(dragStart, bucketIndex),
        end: Math.max(dragStart, bucketIndex),
      });
    } else {
      setHoveredBucket(bucketIndex);
    }
  };

  const handleMouseDown = (e: MouseEvent<HTMLCanvasElement>) => {
    if (!canvasRef.current) return;
    const rect = canvasRef.current.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const bucketIndex = Math.floor((x / canvasWidth) * NUM_BUCKETS);

    setIsDragging(true);
    setDragStart(bucketIndex);
    setTimelineSelection(null);
  };

  const handleMouseUp = (e: MouseEvent<HTMLCanvasElement>) => {
    if (!canvasRef.current || !isDragging || dragStart === null) return;
    const rect = canvasRef.current.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const bucketIndex = Math.floor((x / canvasWidth) * NUM_BUCKETS);

    if (dragStart === bucketIndex) {
      // Single click - seek to this timestamp
      if (timelineData) {
        const bucket = timelineData.buckets[bucketIndex];
        if (bucket) {
          // TODO: Use useEditorSeek to jump to this timestamp
          console.log('Seek to:', bucket.startTime);
        }
      }
      setTimelineSelection(null);
    } else {
      // Drag selection
      setTimelineSelection({
        start: Math.min(dragStart, bucketIndex),
        end: Math.max(dragStart, bucketIndex),
      });
    }

    setIsDragging(false);
    setDragStart(null);
  };

  if (!timelineData) return null;

  return (
    <Box p="md" style={{ backgroundColor: '#f8f9fa', borderBottom: '1px solid #dee2e6' }}>
      <Group justify="space-between" mb="xs">
        <Text size="sm" c="dimmed">
          Timeline (click to jump, drag to select range)
        </Text>
        {timelineSelection && (
          <Button
            size="xs"
            variant="subtle"
            leftSection={<IconX size={14} />}
            onClick={() => setTimelineSelection(null)}
          >
            Clear Selection
          </Button>
        )}
      </Group>

      <canvas
        ref={canvasRef}
        onMouseMove={handleMouseMove}
        onMouseDown={handleMouseDown}
        onMouseUp={handleMouseUp}
        onMouseLeave={() => setHoveredBucket(null)}
        style={{
          width: '100%',
          height: TIMELINE_HEIGHT,
          cursor: isDragging ? 'crosshair' : 'pointer',
          display: 'block',
        }}
      />
    </Box>
  );
}
```

**Note**: This is a simplified version. Full implementation includes:
- Resize observer for responsive canvas
- Better timestamp formatting
- Mapping marked messages to buckets for gold flags
- Smoother seek integration

---

## Component 5: EditorFilters

**File**: `web/src/pages/Stream/components/EditorFilters.tsx`

**Purpose**: Filter dropdowns and checkboxes

**Structure**:
```typescript
import { Group, Select, Switch, Badge, Text } from '@mantine/core';
import { useEditorContext } from './EditorContext';

export function EditorFilters() {
  const {
    sessionStats,
    currentFilters,
    setCurrentFilters,
    hideDeleted,
    setHideDeleted,
    showMarkedOnly,
    setShowMarkedOnly,
    deletedIndices,
    markedIndices,
  } = useEditorContext();

  if (!sessionStats) return null;

  const proxyOptions = ['All Proxies', ...sessionStats.proxies].map((p) => ({ value: p, label: p }));
  const interfaceOptions = ['All Interfaces', ...sessionStats.interfaces].map((i) => ({
    value: i,
    label: i,
  }));
  const directionOptions = [
    { value: '', label: 'All Directions' },
    { value: 'SEND', label: 'SEND' },
    { value: 'RECV', label: 'RECV' },
  ];
  const typeOptions = [
    { value: '', label: 'All Types' },
    { value: 'LINK', label: 'LINK' },
    { value: 'INIT', label: 'INIT' },
    { value: 'INVOKE', label: 'INVOKE' },
    { value: 'INVOKE_REPLY', label: 'INVOKE_REPLY' },
    { value: 'SIGNAL', label: 'SIGNAL' },
    { value: 'PROPERTY_CHANGE', label: 'PROPERTY_CHANGE' },
  ];

  return (
    <Group p="xs" gap="md" style={{ borderBottom: '1px solid #dee2e6' }}>
      <Text size="sm" fw={500}>Filters:</Text>

      <Select
        size="xs"
        data={proxyOptions}
        value={currentFilters.proxy || 'All Proxies'}
        onChange={(value) =>
          setCurrentFilters({ ...currentFilters, proxy: value === 'All Proxies' ? undefined : value })
        }
        w={150}
      />

      <Select
        size="xs"
        data={interfaceOptions}
        value={currentFilters.interface || 'All Interfaces'}
        onChange={(value) =>
          setCurrentFilters({
            ...currentFilters,
            interface: value === 'All Interfaces' ? undefined : value,
          })
        }
        w={150}
      />

      <Select
        size="xs"
        data={directionOptions}
        value={currentFilters.direction || ''}
        onChange={(value) => setCurrentFilters({ ...currentFilters, direction: value || undefined })}
        w={130}
      />

      <Select
        size="xs"
        data={typeOptions}
        value={currentFilters.type || ''}
        onChange={(value) => setCurrentFilters({ ...currentFilters, type: value || undefined })}
        w={150}
      />

      <Switch
        label="Hide deleted"
        checked={hideDeleted}
        onChange={(e) => setHideDeleted(e.currentTarget.checked)}
        size="sm"
      />

      <Switch
        label="Marked only"
        checked={showMarkedOnly}
        onChange={(e) => setShowMarkedOnly(e.currentTarget.checked)}
        size="sm"
      />

      <Badge>{markedIndices.size} marked</Badge>
      <Badge color="orange">{deletedIndices.size} deleted</Badge>
    </Group>
  );
}
```

---

## Component 6: EditorJQPanel

**File**: `web/src/pages/Stream/components/EditorJQPanel.tsx`

**Purpose**: JQ query input with examples

**Structure**:
```typescript
import { useState } from 'react';
import { Group, TextInput, Button, Collapse, Text, Badge } from '@mantine/core';
import { IconCode, IconPlayerPlay, IconChevronDown, IconChevronUp } from '@tabler/icons-react';
import { useEditorContext } from './EditorContext';
import { useEditorJQ } from '@/api/queries';
import { notifications } from '@mantine/notifications';

export function EditorJQPanel() {
  const { sessionStats, setSelectedIndices } = useEditorContext();
  const [query, setQuery] = useState('');
  const [showExamples, setShowExamples] = useState(false);
  const jqMutation = useEditorJQ();

  if (!sessionStats) return null;

  const handleRun = async () => {
    if (!query.trim()) return;

    try {
      const result = await jqMutation.mutateAsync({
        sessionId: sessionStats.sessionId,
        query,
        limit: 100,
      });

      notifications.show({
        title: 'JQ Query Complete',
        message: `Found ${result.totalMatches} matches`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'JQ Query Failed',
        message: error instanceof Error ? error.message : 'Query failed',
        color: 'red',
      });
    }
  };

  const handleSelectMatches = async () => {
    if (!query.trim()) return;

    try {
      const result = await jqMutation.mutateAsync({
        sessionId: sessionStats.sessionId,
        query,
        limit: 1000,
      });

      const indices = new Set(result.matches.map((m) => m.index));
      setSelectedIndices(indices);

      notifications.show({
        title: 'Matches Selected',
        message: `Selected ${indices.size} messages`,
        color: 'blue',
      });
    } catch (error) {
      notifications.show({
        title: 'Selection Failed',
        message: error instanceof Error ? error.message : 'Failed to select matches',
        color: 'red',
      });
    }
  };

  const examples = [
    { label: 'Messages with args', query: 'select(.msg[1] | has("args"))' },
    { label: 'INVOKE messages', query: 'select(.msg[1].msgType == 30)' },
    { label: 'Specific interface', query: 'select(.msg[1].symbol == "demo.Counter")' },
  ];

  return (
    <Group p="xs" gap="md" align="flex-start" style={{ borderBottom: '1px solid #dee2e6' }}>
      <IconCode size={20} />
      <div style={{ flex: 1 }}>
        <TextInput
          placeholder="JQ query (e.g., select(.msg[1].msgType == 30))"
          value={query}
          onChange={(e) => setQuery(e.currentTarget.value)}
          rightSection={
            <Button
              size="xs"
              variant="subtle"
              onClick={() => setShowExamples(!showExamples)}
              rightSection={showExamples ? <IconChevronUp size={14} /> : <IconChevronDown size={14} />}
            >
              Examples
            </Button>
          }
        />

        <Collapse in={showExamples}>
          <Group gap="xs" mt="xs">
            <Text size="xs" c="dimmed">Examples:</Text>
            {examples.map((ex) => (
              <Badge
                key={ex.label}
                variant="light"
                style={{ cursor: 'pointer' }}
                onClick={() => setQuery(ex.query)}
              >
                {ex.label}
              </Badge>
            ))}
          </Group>
        </Collapse>
      </div>

      <Button
        size="xs"
        leftSection={<IconPlayerPlay size={14} />}
        onClick={handleRun}
        loading={jqMutation.isPending}
      >
        Run
      </Button>

      <Button
        size="xs"
        variant="light"
        onClick={handleSelectMatches}
        loading={jqMutation.isPending}
      >
        Select Matches
      </Button>
    </Group>
  );
}
```

---

## Component 7: EditorTable (Virtual Scrolling)

**File**: `web/src/pages/Stream/components/EditorTable.tsx`

**Purpose**: Display messages with virtual scrolling, checkboxes, star buttons

**This requires a virtual scrolling library**:
```bash
cd web
pnpm add react-window @types/react-window
```

**Simplified Structure** (without full virtual scrolling):
```typescript
import { useState } from 'react';
import { Table, Checkbox, ActionIcon, Text } from '@mantine/core';
import { IconStar, IconStarFilled } from '@tabler/icons-react';
import { useEditorContext } from './EditorContext';
import { useEditorMessages } from '@/api/queries';

export function EditorTable() {
  const {
    sessionStats,
    currentFilters,
    selectedIndices,
    toggleSelection,
    markedIndices,
    toggleMarked,
    deletedIndices,
    hideDeleted,
    showMarkedOnly,
  } = useEditorContext();

  const [offset, setOffset] = useState(0);
  const limit = 100;

  const { data: messagesData } = useEditorMessages(
    sessionStats?.sessionId || null,
    offset,
    limit,
    currentFilters
  );

  if (!messagesData) return null;

  // Filter client-side for deleted/marked
  let visibleMessages = messagesData.messages;
  if (hideDeleted) {
    visibleMessages = visibleMessages.filter((m) => !deletedIndices.has(m.index));
  }
  if (showMarkedOnly) {
    visibleMessages = visibleMessages.filter((m) => markedIndices.has(m.index));
  }

  const formatTimestamp = (ts: number) => {
    const date = new Date(ts);
    return date.toLocaleTimeString();
  };

  return (
    <div style={{ overflowX: 'auto' }}>
      <Table highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th style={{ width: 40 }}>
              <Checkbox
                checked={visibleMessages.every((m) => selectedIndices.has(m.index))}
                onChange={(e) => {
                  const newSet = new Set(selectedIndices);
                  visibleMessages.forEach((m) => {
                    if (e.currentTarget.checked) {
                      newSet.add(m.index);
                    } else {
                      newSet.delete(m.index);
                    }
                  });
                  setSelectedIndices(newSet);
                }}
              />
            </Table.Th>
            <Table.Th style={{ width: 40 }}>★</Table.Th>
            <Table.Th style={{ width: 60 }}>#</Table.Th>
            <Table.Th>Time</Table.Th>
            <Table.Th>Proxy</Table.Th>
            <Table.Th>Dir</Table.Th>
            <Table.Th>Type</Table.Th>
            <Table.Th>Symbol</Table.Th>
            <Table.Th>ReqID</Table.Th>
            <Table.Th>Args</Table.Th>
          </Table.Tr>
        </Table.Thead>
        <Table.Tbody>
          {visibleMessages.map((message) => (
            <Table.Tr
              key={message.index}
              style={{
                backgroundColor: deletedIndices.has(message.index)
                  ? '#ffe0e0'
                  : selectedIndices.has(message.index)
                  ? '#e7f5ff'
                  : undefined,
              }}
            >
              <Table.Td>
                <Checkbox
                  checked={selectedIndices.has(message.index)}
                  onChange={() => toggleSelection(message.index)}
                />
              </Table.Td>
              <Table.Td>
                <ActionIcon
                  variant="subtle"
                  color={markedIndices.has(message.index) ? 'yellow' : 'gray'}
                  onClick={() => toggleMarked(message.index)}
                >
                  {markedIndices.has(message.index) ? (
                    <IconStarFilled size={16} />
                  ) : (
                    <IconStar size={16} />
                  )}
                </ActionIcon>
              </Table.Td>
              <Table.Td>
                <Text size="sm" c="dimmed">
                  {message.index}
                </Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm" ff="monospace">
                  {formatTimestamp(message.timestamp)}
                </Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{message.proxy}</Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm" c={message.direction === 'SEND' ? 'blue' : 'green'}>
                  {message.direction}
                </Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{message.parsed.msgTypeName}</Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{message.parsed.symbol || '-'}</Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{message.parsed.requestId || '-'}</Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm" lineClamp={1}>
                  {JSON.stringify(message.parsed.args)}
                </Text>
              </Table.Td>
            </Table.Tr>
          ))}
        </Table.Tbody>
      </Table>

      {/* Pagination controls */}
      <Group p="md" justify="center">
        <Button size="xs" onClick={() => setOffset(Math.max(0, offset - limit))} disabled={offset === 0}>
          Previous
        </Button>
        <Text size="sm">
          {offset + 1} - {Math.min(offset + limit, messagesData.total)} of {messagesData.total}
        </Text>
        <Button
          size="xs"
          onClick={() => setOffset(offset + limit)}
          disabled={offset + limit >= messagesData.total}
        >
          Next
        </Button>
      </Group>
    </div>
  );
}
```

**For production**: Replace with `react-window` for virtual scrolling. See wsproxy example.

---

## Component 8: EditorToolbar

**File**: `web/src/pages/Stream/components/EditorToolbar.tsx`

**Purpose**: Bottom toolbar with selection controls, mark/unmark, cut operations, export

**Reference**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/editor/components/EditorToolbar.tsx`

**Structure**:
```typescript
import { Group, Button, Text, Badge, Menu } from '@mantine/core';
import {
  IconSelect,
  IconSelectAll,
  IconX,
  IconCut,
  IconRestore,
  IconStar,
  IconStarOff,
  IconDownload,
} from '@tabler/icons-react';
import { useEditorContext } from './EditorContext';
import { useEditorMessages, useEditorExport } from '@/api/queries';
import { notifications } from '@mantine/notifications';

export function EditorToolbar() {
  const {
    sessionStats,
    selectedIndices,
    clearSelection,
    setSelectedIndices,
    deletedIndices,
    setDeletedIndices,
    markedIndices,
    setMarkedIndices,
    currentFilters,
  } = useEditorContext();

  const { data: messagesData } = useEditorMessages(
    sessionStats?.sessionId || null,
    0,
    10000, // Load all for operations
    currentFilters
  );

  const exportMutation = useEditorExport();

  const selectedCount = selectedIndices.size;
  const deletedCount = deletedIndices.size;
  const markedCount = markedIndices.size;

  const handleSelectAll = () => {
    if (!messagesData) return;
    const allIndices = new Set(messagesData.messages.map((m) => m.index));
    setSelectedIndices(allIndices);
  };

  const handleCutSelected = () => {
    const newDeleted = new Set(deletedIndices);
    selectedIndices.forEach((idx) => newDeleted.add(idx));
    setDeletedIndices(newDeleted);
    clearSelection();
  };

  const handleCutBefore = () => {
    if (selectedIndices.size === 0 || !messagesData) return;
    const minSelected = Math.min(...Array.from(selectedIndices));
    const newDeleted = new Set(deletedIndices);
    messagesData.messages.forEach((m) => {
      if (m.index < minSelected) newDeleted.add(m.index);
    });
    setDeletedIndices(newDeleted);
  };

  const handleCutAfter = () => {
    if (selectedIndices.size === 0 || !messagesData) return;
    const maxSelected = Math.max(...Array.from(selectedIndices));
    const newDeleted = new Set(deletedIndices);
    messagesData.messages.forEach((m) => {
      if (m.index > maxSelected) newDeleted.add(m.index);
    });
    setDeletedIndices(newDeleted);
  };

  const handleUndoAllCuts = () => {
    setDeletedIndices(new Set());
  };

  const handleMarkSelected = () => {
    const newMarked = new Set(markedIndices);
    selectedIndices.forEach((idx) => newMarked.add(idx));
    setMarkedIndices(newMarked);
  };

  const handleUnmarkSelected = () => {
    const newMarked = new Set(markedIndices);
    selectedIndices.forEach((idx) => newMarked.delete(idx));
    setMarkedIndices(newMarked);
  };

  const handleExport = async (indices?: number[]) => {
    if (!sessionStats) return;

    try {
      const blob = await exportMutation.mutateAsync({
        sessionId: sessionStats.sessionId,
        indices,
      });

      // Download file
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = sessionStats.filename;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);

      notifications.show({
        title: 'Export Complete',
        message: `Downloaded ${sessionStats.filename}`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Export Failed',
        message: error instanceof Error ? error.message : 'Export failed',
        color: 'red',
      });
    }
  };

  if (!sessionStats) return null;

  return (
    <Group gap="md" p="xs" style={{ borderTop: '1px solid #dee2e6' }}>
      <Group gap="xs">
        <IconSelect size={16} />
        <Text size="sm" fw={500}>
          Selection:
        </Text>
        {selectedCount > 0 ? (
          <Badge size="md" color="blue">
            {selectedCount.toLocaleString()} selected
          </Badge>
        ) : (
          <Text size="sm" c="dimmed">
            None
          </Text>
        )}
      </Group>

      {selectedCount > 0 && (
        <Button
          size="xs"
          variant="subtle"
          color="gray"
          leftSection={<IconX size={14} />}
          onClick={clearSelection}
        >
          Clear
        </Button>
      )}

      <Button
        size="xs"
        variant="light"
        leftSection={<IconSelectAll size={14} />}
        onClick={handleSelectAll}
      >
        Select All
      </Button>

      {/* Cut operations */}
      {selectedCount > 0 && (
        <Menu shadow="md" width={200}>
          <Menu.Target>
            <Button size="xs" variant="light" color="orange" leftSection={<IconCut size={14} />}>
              Cut
            </Button>
          </Menu.Target>

          <Menu.Dropdown>
            <Menu.Item onClick={handleCutBefore}>Cut Before Selected</Menu.Item>
            <Menu.Item onClick={handleCutSelected}>Cut Selected</Menu.Item>
            <Menu.Item onClick={handleCutAfter}>Cut After Selected</Menu.Item>
          </Menu.Dropdown>
        </Menu>
      )}

      {deletedCount > 0 && (
        <>
          <Badge size="md" color="orange">
            {deletedCount.toLocaleString()} deleted
          </Badge>
          <Button
            size="xs"
            variant="subtle"
            color="orange"
            leftSection={<IconRestore size={14} />}
            onClick={handleUndoAllCuts}
          >
            Undo All Cuts
          </Button>
        </>
      )}

      {/* Mark operations */}
      {selectedCount > 0 && (
        <>
          <Button
            size="xs"
            variant="light"
            color="yellow"
            leftSection={<IconStar size={14} />}
            onClick={handleMarkSelected}
          >
            Mark Selected
          </Button>
          <Button
            size="xs"
            variant="subtle"
            color="yellow"
            leftSection={<IconStarOff size={14} />}
            onClick={handleUnmarkSelected}
          >
            Unmark
          </Button>
        </>
      )}

      {markedCount > 0 && (
        <Badge size="md" color="yellow">
          {markedCount.toLocaleString()} marked
        </Badge>
      )}

      {/* Export menu */}
      <Menu shadow="md" width={200}>
        <Menu.Target>
          <Button
            size="xs"
            variant="filled"
            color="blue"
            leftSection={<IconDownload size={14} />}
            loading={exportMutation.isPending}
          >
            Export
          </Button>
        </Menu.Target>

        <Menu.Dropdown>
          <Menu.Item onClick={() => handleExport()}>Export All</Menu.Item>
          <Menu.Item onClick={() => handleExport(Array.from(selectedIndices))} disabled={selectedCount === 0}>
            Export Selected ({selectedCount})
          </Menu.Item>
          <Menu.Item onClick={() => handleExport(Array.from(markedIndices))} disabled={markedCount === 0}>
            Export Marked ({markedCount})
          </Menu.Item>
          <Menu.Item
            onClick={() => {
              if (!messagesData) return;
              const undeleted = messagesData.messages
                .filter((m) => !deletedIndices.has(m.index))
                .map((m) => m.index);
              handleExport(undeleted);
            }}
            disabled={deletedCount === 0}
          >
            Export Undeleted
          </Menu.Item>
        </Menu.Dropdown>
      </Menu>
    </Group>
  );
}
```

---

## Component 9: Keyboard Shortcuts

**File**: `web/src/pages/Stream/components/useEditorKeyboard.ts`

**Purpose**: Global keyboard shortcuts for editor

**Structure**:
```typescript
import { useEffect } from 'react';
import { useEditorContext } from './EditorContext';

export function useEditorKeyboard() {
  const { clearSelection, sessionStats } = useEditorContext();

  useEffect(() => {
    if (!sessionStats) return;

    const handleKeyDown = (e: KeyboardEvent) => {
      // Ctrl/Cmd + A - Select All
      if ((e.ctrlKey || e.metaKey) && e.key === 'a') {
        e.preventDefault();
        // TODO: Implement select all
      }

      // Escape - Clear Selection
      if (e.key === 'Escape') {
        clearSelection();
      }

      // Ctrl/Cmd + E - Export
      if ((e.ctrlKey || e.metaKey) && e.key === 'e') {
        e.preventDefault();
        // TODO: Trigger export
      }
    };

    document.addEventListener('keydown', handleKeyDown);
    return () => document.removeEventListener('keydown', handleKeyDown);
  }, [sessionStats, clearSelection]);
}
```

**Usage**: Call in `StreamEditor.tsx` when session is loaded:
```typescript
function EditorContentWithSession() {
  useEditorKeyboard(); // Add this
  return <>...</>;
}
```

---

## Step 10: Integration

**Update** `web/src/pages/Stream/StreamEditor.tsx`:

```typescript
import { useState } from 'react';
import { Container, Stack, Group, Title, Button } from '@mantine/core';
import { IconFilePlus } from '@tabler/icons-react';
import { EditorProvider, useEditorContext } from './components/EditorContext';
import { EditorWelcome } from './components/EditorWelcome';
import { EditorLoadDrawer } from './components/EditorLoadDrawer';
import { EditorStats } from './components/EditorStats';
import { EditorTimeline } from './components/EditorTimeline';
import { EditorFilters } from './components/EditorFilters';
import { EditorJQPanel } from './components/EditorJQPanel';
import { EditorTable } from './components/EditorTable';
import { EditorToolbar } from './components/EditorToolbar';
import { useEditorKeyboard } from './components/useEditorKeyboard';

function EditorContent() {
  const { sessionStats } = useEditorContext();
  const [loadDrawerOpen, setLoadDrawerOpen] = useState(false);

  return (
    <>
      {!sessionStats ? (
        <EditorWelcome />
      ) : (
        <EditorContentWithSession loadDrawerOpen={loadDrawerOpen} setLoadDrawerOpen={setLoadDrawerOpen} />
      )}

      <EditorLoadDrawer opened={loadDrawerOpen} onClose={() => setLoadDrawerOpen(false)} />
    </>
  );
}

function EditorContentWithSession({
  loadDrawerOpen,
  setLoadDrawerOpen,
}: {
  loadDrawerOpen: boolean;
  setLoadDrawerOpen: (open: boolean) => void;
}) {
  useEditorKeyboard();

  return (
    <Stack gap="md">
      {/* Header with Set Stream button */}
      <Group justify="space-between">
        <Title order={2}>Stream Editor</Title>
        <Button leftSection={<IconFilePlus size={16} />} onClick={() => setLoadDrawerOpen(true)}>
          Set Stream
        </Button>
      </Group>

      <EditorStats />
      <EditorTimeline />
      <EditorFilters />
      <EditorJQPanel />
      <EditorTable />
      <EditorToolbar />
    </Stack>
  );
}

export function StreamEditor() {
  return (
    <EditorProvider>
      <Container size="xl" py="md">
        <EditorContent />
      </Container>
    </EditorProvider>
  );
}
```

---

## Testing

### Backend Testing
```bash
# Build and run
go build -o /tmp/apigear ./cmd/apigear
/tmp/apigear stream

# Test endpoints
# 1. Load a trace file
curl -X POST http://localhost:8080/api/v1/stream/editor/load \
  -H "Content-Type: application/json" \
  -d '{"filename":"proxy1.jsonl"}'
# Save sessionId from response

# 2. Get timeline
curl "http://localhost:8080/api/v1/stream/editor/timeline?sessionId=<SESSION_ID>"

# 3. Get messages
curl "http://localhost:8080/api/v1/stream/editor/messages?sessionId=<SESSION_ID>&offset=0&limit=10"

# 4. Run JQ query
curl -X POST http://localhost:8080/api/v1/stream/editor/jq \
  -H "Content-Type: application/json" \
  -d '{"sessionId":"<SESSION_ID>","query":"select(.msg[1].msgType == 30)","limit":10}'
```

### Frontend Testing
```bash
cd web
pnpm dev
```

1. Navigate to http://localhost:5173/stream/editor
2. Should see welcome screen
3. Click "Set Stream"
4. Upload a trace file or select from list
5. Verify stats bar appears
6. Verify timeline renders
7. Test filters
8. Test JQ queries
9. Test selection (checkboxes)
10. Test mark/unmark (stars)
11. Test cut operations
12. Test export

---

## Known Issues & Improvements

### Issues to Fix
1. **Canvas resize**: EditorTimeline needs ResizeObserver for responsive sizing
2. **Virtual scrolling**: EditorTable should use `react-window` for large datasets
3. **Marked message flags**: Timeline doesn't show gold flags yet (requires bucket mapping)
4. **Seek integration**: Timeline click doesn't scroll table yet
5. **Filter reset**: No "Clear All Filters" button

### Performance Optimizations
1. Use `React.memo` on table rows
2. Debounce filter changes
3. Add loading states
4. Add error boundaries per component

### UX Improvements
1. Add keyboard shortcut help modal (? key)
2. Add confirmation dialogs for destructive actions
3. Add undo/redo for cut operations
4. Add saved filter presets
5. Add column sorting in table
6. Add message detail modal (click row to expand)

---

## Reference Files

All wsproxy source code is available at:
- `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/editor/`
- `/Users/jryannel/dev/tmp/wsproxy/pkg/web/editor.go`

Screenshots are at:
- `/Users/jryannel/dev/github.com/apigear-io/cli/wsproxy_screenshots/wsproxy_stream_editor*.png`

---

## Summary

**Completed**: Backend (100%), Frontend Foundation (30%)
**Remaining**: 8 components (~500-700 lines each)
**Estimated Time**: 4-6 hours for experienced developer
**Total Lines**: ~3000-4000 lines of frontend code

The architecture is sound, patterns are established, and all API endpoints are working. The remaining work is primarily UI components following the established patterns.

Good luck! 🚀
