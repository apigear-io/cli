import { useEffect, useState, useRef } from 'react';
import { Paper, ScrollArea, Group, Badge, Text, Stack, ActionIcon } from '@mantine/core';
import { IconTrash } from '@tabler/icons-react';
import type { ScriptOutputEntry, ScriptOutputLevel } from '@/api/types';

interface ConsoleOutputProps {
  scriptId: string;
  height?: string;
}

function getLevelColor(level: ScriptOutputLevel): string {
  switch (level) {
    case 'error':
      return 'red';
    case 'warn':
      return 'yellow';
    case 'info':
      return 'blue';
    case 'debug':
      return 'gray';
    case 'log':
    default:
      return 'gray';
  }
}

export function ConsoleOutput({ scriptId, height = '300px' }: ConsoleOutputProps) {
  const [entries, setEntries] = useState<ScriptOutputEntry[]>([]);
  const [isConnected, setIsConnected] = useState(false);
  const scrollAreaRef = useRef<HTMLDivElement>(null);
  const eventSourceRef = useRef<EventSource | null>(null);

  useEffect(() => {
    const eventSource = new EventSource(
      `/api/v1/stream/scripts/output?id=${encodeURIComponent(scriptId)}`
    );

    eventSourceRef.current = eventSource;

    eventSource.addEventListener('connected', () => {
      setIsConnected(true);
      setEntries([
        {
          level: 'info',
          message: 'Connected to script output stream',
        },
      ]);
    });

    eventSource.addEventListener('output', (event) => {
      const entry: ScriptOutputEntry = JSON.parse(event.data);
      setEntries((prev) => [...prev, entry]);

      // Auto-scroll to bottom
      setTimeout(() => {
        if (scrollAreaRef.current) {
          const viewport = scrollAreaRef.current.querySelector('[data-radix-scroll-area-viewport]');
          if (viewport) {
            viewport.scrollTop = viewport.scrollHeight;
          }
        }
      }, 10);
    });

    eventSource.addEventListener('closed', () => {
      setIsConnected(false);
      setEntries((prev) => [
        ...prev,
        {
          level: 'info',
          message: 'Script stopped',
        },
      ]);
      eventSource.close();
    });

    eventSource.onerror = () => {
      setIsConnected(false);
      setEntries((prev) => [
        ...prev,
        {
          level: 'error',
          message: 'Connection to script output lost',
        },
      ]);
      eventSource.close();
    };

    return () => {
      eventSource.close();
      eventSourceRef.current = null;
    };
  }, [scriptId]);

  const handleClear = () => {
    setEntries([]);
  };

  return (
    <Paper withBorder p="md">
      <Stack gap="xs">
        <Group justify="space-between">
          <Group gap="xs">
            <Text size="sm" fw={600}>
              Console Output
            </Text>
            {isConnected && <Badge size="xs" color="green">Connected</Badge>}
          </Group>
          <ActionIcon
            variant="subtle"
            color="gray"
            onClick={handleClear}
            title="Clear console"
          >
            <IconTrash size={16} />
          </ActionIcon>
        </Group>

        <ScrollArea h={height} viewportRef={scrollAreaRef}>
          <Stack gap={4}>
            {entries.length === 0 && (
              <Text size="sm" c="dimmed" fs="italic">
                No output yet...
              </Text>
            )}
            {entries.map((entry, i) => (
              <Group key={i} gap="xs" wrap="nowrap" align="flex-start">
                <Badge
                  size="xs"
                  color={getLevelColor(entry.level)}
                  style={{ flexShrink: 0, marginTop: 2 }}
                >
                  {entry.level}
                </Badge>
                <Text
                  size="sm"
                  ff="monospace"
                  style={{ wordBreak: 'break-word', whiteSpace: 'pre-wrap' }}
                >
                  {entry.message}
                </Text>
              </Group>
            ))}
          </Stack>
        </ScrollArea>
      </Stack>
    </Paper>
  );
}
