import { useState, useEffect, useCallback } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Stack, Paper, Text, Group, Badge, Button, ScrollArea, SimpleGrid, ThemeIcon } from '@mantine/core';
import { IconCheck, IconX, IconLoader, IconArrowLeft, IconFileCheck, IconFileOff, IconCopy, IconFiles } from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { Breadcrumbs } from '@/components/Breadcrumbs';
import type { TaskEvent, CodeGenerationSummary } from '@/api/types';

interface LogEntry {
  timestamp: Date;
  type: 'connected' | 'task' | 'error' | 'completed';
  message: string;
  data?: unknown;
}

function parseSSEEvents(text: string): { eventType: string; data: string }[] {
  const events: { eventType: string; data: string }[] = [];
  const blocks = text.split('\n\n');
  for (const block of blocks) {
    if (!block.trim()) continue;
    let eventType = '';
    let data = '';
    for (const line of block.split('\n')) {
      if (line.startsWith('event: ')) eventType = line.slice(7);
      else if (line.startsWith('data: ')) data = line.slice(6);
    }
    if (eventType && data) {
      events.push({ eventType, data });
    }
  }
  return events;
}

export function CodeGeneration() {
  const { encodedSolutionPath } = useParams<{ encodedSolutionPath: string }>();
  const navigate = useNavigate();

  const [logs, setLogs] = useState<LogEntry[]>([]);
  const [status, setStatus] = useState<'connecting' | 'running' | 'completed' | 'error'>('connecting');
  const [currentTask, setCurrentTask] = useState<string>('');
  const [isDone, setIsDone] = useState(false);
  const [summary, setSummary] = useState<CodeGenerationSummary | null>(null);

  const addLog = useCallback((type: LogEntry['type'], message: string, data?: unknown) => {
    setLogs((prev) => [
      ...prev,
      { timestamp: new Date(), type, message, data },
    ]);
  }, []);

  useEffect(() => {
    if (!encodedSolutionPath) {
      navigate('/codegen/projects');
      return;
    }

    if (isDone) return;

    const abortController = new AbortController();

    const solutionPath = decodeURIComponent(encodedSolutionPath);
    const url = new URL('/api/v1/projects/generate', window.location.origin);
    url.searchParams.set('path', solutionPath);
    url.searchParams.set('force', 'false');

    const runGeneration = async () => {
      let response: Response;
      try {
        response = await fetch(url.toString(), { signal: abortController.signal });
      } catch {
        if (abortController.signal.aborted) return;
        addLog('error', 'Failed to connect to code generation service');
        setStatus('error');
        setIsDone(true);
        notifications.show({
          title: 'Connection Error',
          message: 'Failed to connect to code generation service',
          color: 'red',
        });
        return;
      }

      // Handle HTTP errors (e.g., template not found) — returned as JSON before SSE starts
      if (!response.ok) {
        let errorMessage = 'Code generation failed';
        try {
          const errorBody = await response.json();
          errorMessage = errorBody.message || errorBody.error || errorMessage;
        } catch {
          errorMessage = `Server error: ${response.status} ${response.statusText}`;
        }
        addLog('error', errorMessage);
        setStatus('error');
        setIsDone(true);
        notifications.show({
          title: 'Error',
          message: errorMessage,
          color: 'red',
        });
        return;
      }

      // Stream SSE events from response body
      const reader = response.body?.getReader();
      if (!reader) {
        addLog('error', 'No response body');
        setStatus('error');
        setIsDone(true);
        return;
      }

      const decoder = new TextDecoder();
      let buffer = '';

      try {
        while (true) {
          const { done: readerDone, value } = await reader.read();
          if (readerDone) break;

          buffer += decoder.decode(value, { stream: true });

          // Split on double newlines (SSE event separator)
          const parts = buffer.split('\n\n');
          buffer = parts.pop() || '';

          for (const part of parts) {
            const events = parseSSEEvents(part + '\n\n');
            for (const { eventType, data: eventData } of events) {
              try {
                const data = JSON.parse(eventData);

                switch (eventType) {
                  case 'connected':
                    addLog('connected', 'Connected to code generation service', data);
                    setStatus('running');
                    break;
                  case 'task': {
                    const taskData = data as TaskEvent;
                    addLog('task', `Task: ${taskData.name} - ${taskData.state}`, data);
                    setCurrentTask(`${taskData.name} (${taskData.state})`);
                    break;
                  }
                  case 'error':
                    addLog('error', data.message || 'An error occurred', data);
                    setStatus('error');
                    setIsDone(true);
                    notifications.show({
                      title: 'Error',
                      message: data.message || 'Code generation failed',
                      color: 'red',
                    });
                    return;
                  case 'completed':
                    addLog('completed', data.message || 'Code generation completed', data);
                    setStatus('completed');
                    setCurrentTask('');
                    setIsDone(true);
                    if (data.totalFiles !== undefined) {
                      setSummary({
                        filesWritten: data.filesWritten ?? 0,
                        filesSkipped: data.filesSkipped ?? 0,
                        filesCopied: data.filesCopied ?? 0,
                        totalFiles: data.totalFiles ?? 0,
                        targetCount: data.targetCount ?? 0,
                        durationMs: data.durationMs ?? 0,
                      });
                    }
                    notifications.show({
                      title: 'Success',
                      message: 'Code generation completed successfully',
                      color: 'green',
                    });
                    return;
                }
              } catch {
                // Skip unparseable events
              }
            }
          }
        }
      } catch {
        if (abortController.signal.aborted) return;
        addLog('error', 'Connection lost');
        setStatus('error');
        setIsDone(true);
        notifications.show({
          title: 'Connection Error',
          message: 'Connection to code generation service was lost',
          color: 'red',
        });
      }
    };

    runGeneration();

    return () => {
      abortController.abort();
    };
  }, [encodedSolutionPath, navigate, isDone, addLog]);

  const getStatusBadge = () => {
    switch (status) {
      case 'connecting':
        return (
          <Badge color="blue" leftSection={<IconLoader size={12} />}>
            Connecting
          </Badge>
        );
      case 'running':
        return (
          <Badge color="blue" leftSection={<IconLoader size={12} />}>
            Running
          </Badge>
        );
      case 'completed':
        return (
          <Badge color="green" leftSection={<IconCheck size={12} />}>
            Completed
          </Badge>
        );
      case 'error':
        return (
          <Badge color="red" leftSection={<IconX size={12} />}>
            Error
          </Badge>
        );
    }
  };

  const getLogColor = (type: LogEntry['type']) => {
    switch (type) {
      case 'connected':
        return 'blue';
      case 'task':
        return 'dimmed';
      case 'error':
        return 'red';
      case 'completed':
        return 'green';
    }
  };

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="flex-start">
        <div>
          <Breadcrumbs
            items={[
              { label: 'Projects', href: '/codegen/projects' },
              { label: 'Code Generation' },
            ]}
          />
        </div>
        <Group>
          {getStatusBadge()}
          <Button
            variant="light"
            leftSection={<IconArrowLeft size={16} />}
            onClick={() => navigate(-1)}
          >
            Back
          </Button>
        </Group>
      </Group>

      {/* Current Task */}
      {currentTask && status === 'running' && (
        <Paper shadow="xs" p="md" withBorder>
          <Group gap="xs">
            <IconLoader size={20} style={{ animation: 'spin 1s linear infinite' }} />
            <Text size="sm" fw={500}>
              {currentTask}
            </Text>
          </Group>
        </Paper>
      )}

      {/* Summary */}
      {summary && status === 'completed' && (
        <Paper shadow="xs" p="md" withBorder>
          <Group gap="xs" mb="md">
            <ThemeIcon color="green" size="lg" radius="xl">
              <IconCheck size={18} />
            </ThemeIcon>
            <Text fw={600} size="lg">
              Generation Complete
            </Text>
            <Badge color="gray" variant="light" size="lg">
              {summary.durationMs < 1000
                ? `${summary.durationMs}ms`
                : `${(summary.durationMs / 1000).toFixed(1)}s`}
            </Badge>
          </Group>
          <SimpleGrid cols={{ base: 2, sm: 4 }} spacing="md">
            <Paper p="sm" withBorder radius="md">
              <Group gap="xs">
                <IconFileCheck size={18} color="var(--mantine-color-green-6)" />
                <div>
                  <Text size="xl" fw={700}>{summary.filesWritten}</Text>
                  <Text size="xs" c="dimmed">Written</Text>
                </div>
              </Group>
            </Paper>
            <Paper p="sm" withBorder radius="md">
              <Group gap="xs">
                <IconFileOff size={18} color="var(--mantine-color-yellow-6)" />
                <div>
                  <Text size="xl" fw={700}>{summary.filesSkipped}</Text>
                  <Text size="xs" c="dimmed">Skipped</Text>
                </div>
              </Group>
            </Paper>
            <Paper p="sm" withBorder radius="md">
              <Group gap="xs">
                <IconCopy size={18} color="var(--mantine-color-blue-6)" />
                <div>
                  <Text size="xl" fw={700}>{summary.filesCopied}</Text>
                  <Text size="xs" c="dimmed">Copied</Text>
                </div>
              </Group>
            </Paper>
            <Paper p="sm" withBorder radius="md">
              <Group gap="xs">
                <IconFiles size={18} color="var(--mantine-color-gray-6)" />
                <div>
                  <Text size="xl" fw={700}>{summary.totalFiles}</Text>
                  <Text size="xs" c="dimmed">Total</Text>
                </div>
              </Group>
            </Paper>
          </SimpleGrid>
        </Paper>
      )}

      {/* Logs */}
      <Paper shadow="xs" p="md" withBorder>
        <Text size="sm" fw={500} mb="md">
          Generation Log
        </Text>
        <ScrollArea h={500} type="auto">
          <Stack gap="xs">
            {logs.length === 0 ? (
              <Text size="sm" c="dimmed" ta="center" py="xl">
                Waiting for events...
              </Text>
            ) : (
              logs.map((log, index) => (
                <Group key={index} gap="xs" wrap="nowrap" align="flex-start">
                  <Text size="xs" c="dimmed" style={{ minWidth: 80 }}>
                    {log.timestamp.toLocaleTimeString()}
                  </Text>
                  <Text size="sm" c={getLogColor(log.type)} style={{ flex: 1 }}>
                    {log.message}
                  </Text>
                </Group>
              ))
            )}
          </Stack>
        </ScrollArea>
      </Paper>

      {/* Action Buttons */}
      {(status === 'completed' || status === 'error') && (
        <Group justify="flex-end">
          <Button
            onClick={() => navigate(-1)}
          >
            Done
          </Button>
        </Group>
      )}

      <style>{`
        @keyframes spin {
          from {
            transform: rotate(0deg);
          }
          to {
            transform: rotate(360deg);
          }
        }
      `}</style>
    </Stack>
  );
}
