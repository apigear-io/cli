import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Stack, Paper, Text, Group, Badge, Button, ScrollArea } from '@mantine/core';
import { IconCheck, IconX, IconLoader, IconArrowLeft } from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { Breadcrumbs } from '@/components/Breadcrumbs';
import type { TaskEvent } from '@/api/types';

interface LogEntry {
  timestamp: Date;
  type: 'connected' | 'task' | 'error' | 'completed';
  message: string;
  data?: unknown;
}

export function CodeGeneration() {
  const { encodedSolutionPath } = useParams<{ encodedSolutionPath: string }>();
  const navigate = useNavigate();

  const [logs, setLogs] = useState<LogEntry[]>([]);
  const [status, setStatus] = useState<'connecting' | 'running' | 'completed' | 'error'>('connecting');
  const [currentTask, setCurrentTask] = useState<string>('');
  const [isDone, setIsDone] = useState(false);

  useEffect(() => {
    if (!encodedSolutionPath) {
      navigate('/codegen/projects');
      return;
    }

    // Prevent reconnection if already done
    if (isDone) {
      return;
    }

    const solutionPath = decodeURIComponent(encodedSolutionPath);
    const url = new URL('/api/v1/projects/generate', window.location.origin);
    url.searchParams.set('path', solutionPath);
    url.searchParams.set('force', 'false');

    const eventSource = new EventSource(url.toString());

    eventSource.addEventListener('connected', (event: MessageEvent) => {
      const data = JSON.parse(event.data);
      addLog('connected', 'Connected to code generation service', data);
      setStatus('running');
    });

    eventSource.addEventListener('task', (event: MessageEvent) => {
      const data = JSON.parse(event.data) as TaskEvent;
      addLog('task', `Task: ${data.name} - ${data.state}`, data);
      setCurrentTask(`${data.name} (${data.state})`);
    });

    eventSource.addEventListener('error', (event: MessageEvent) => {
      const data = JSON.parse(event.data);
      addLog('error', data.message || 'An error occurred', data);
      setStatus('error');
      setIsDone(true);
      eventSource.close();

      notifications.show({
        title: 'Error',
        message: data.message || 'Code generation failed',
        color: 'red',
      });
    });

    eventSource.addEventListener('completed', (event: MessageEvent) => {
      const data = JSON.parse(event.data);
      addLog('completed', data.message || 'Code generation completed', data);
      setStatus('completed');
      setCurrentTask('');
      setIsDone(true);
      eventSource.close();

      notifications.show({
        title: 'Success',
        message: 'Code generation completed successfully',
        color: 'green',
      });
    });

    eventSource.onerror = () => {
      // Only treat as error if we're still connecting or running
      if (status === 'connecting' || status === 'running') {
        addLog('error', 'Connection lost or failed');
        setStatus('error');
        setIsDone(true);

        notifications.show({
          title: 'Connection Error',
          message: 'Connection to code generation service was lost',
          color: 'red',
        });
      }
      eventSource.close();
    };

    return () => {
      eventSource.close();
    };
  }, [encodedSolutionPath, navigate, isDone, status]);

  const addLog = (type: LogEntry['type'], message: string, data?: unknown) => {
    setLogs((prev) => [
      ...prev,
      {
        timestamp: new Date(),
        type,
        message,
        data,
      },
    ]);
  };

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
