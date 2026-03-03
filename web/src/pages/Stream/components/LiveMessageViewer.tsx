import { useEffect, useState, useRef } from 'react';
import {
  Paper,
  ScrollArea,
  Group,
  Badge,
  Text,
  Stack,
  ActionIcon,
  Switch,
  Code,
  Divider,
} from '@mantine/core';
import { IconTrash, IconPlayerPause, IconPlayerPlay } from '@tabler/icons-react';
import type { ParsedMessageEvent } from '@/api/types';

interface LiveMessageViewerProps {
  proxyName: string;
  height?: string;
}

export function LiveMessageViewer({ proxyName, height = '600px' }: LiveMessageViewerProps) {
  const [messages, setMessages] = useState<ParsedMessageEvent[]>([]);
  const [isConnected, setIsConnected] = useState(false);
  const [isPaused, setIsPaused] = useState(false);
  const [autoScroll, setAutoScroll] = useState(true);
  const scrollAreaRef = useRef<HTMLDivElement>(null);
  const eventSourceRef = useRef<EventSource | null>(null);
  const pausedQueueRef = useRef<ParsedMessageEvent[]>([]);

  useEffect(() => {
    const eventSource = new EventSource(
      `/api/v1/stream/proxies/${encodeURIComponent(proxyName)}/events`
    );

    eventSourceRef.current = eventSource;

    eventSource.addEventListener('connected', () => {
      setIsConnected(true);
      setMessages([
        {
          type: 'info',
          proxy: proxyName,
          direction: 'INFO',
          timestamp: Date.now(),
          message: { text: `Connected to proxy: ${proxyName}` },
        },
      ]);
    });

    eventSource.addEventListener('message', (event) => {
      const messageEvent: ParsedMessageEvent = JSON.parse(event.data);

      if (isPaused) {
        pausedQueueRef.current.push(messageEvent);
      } else {
        setMessages((prev) => [...prev, messageEvent]);

        // Auto-scroll to bottom
        if (autoScroll) {
          setTimeout(() => {
            if (scrollAreaRef.current) {
              const viewport = scrollAreaRef.current.querySelector('[data-radix-scroll-area-viewport]');
              if (viewport) {
                viewport.scrollTop = viewport.scrollHeight;
              }
            }
          }, 10);
        }
      }
    });

    eventSource.onerror = () => {
      setIsConnected(false);
      setMessages((prev) => [
        ...prev,
        {
          type: 'error',
          proxy: proxyName,
          direction: 'ERROR',
          timestamp: Date.now(),
          message: { text: 'Connection lost' },
        },
      ]);
      eventSource.close();
    };

    return () => {
      eventSource.close();
      eventSourceRef.current = null;
    };
  }, [proxyName, isPaused, autoScroll]);

  const handleClear = () => {
    setMessages([]);
    pausedQueueRef.current = [];
  };

  const handleTogglePause = () => {
    if (isPaused) {
      // Resume: add queued messages
      setMessages((prev) => [...prev, ...pausedQueueRef.current]);
      pausedQueueRef.current = [];
    }
    setIsPaused(!isPaused);
  };

  const formatTimestamp = (ts: number): string => {
    const date = new Date(ts);
    return date.toLocaleTimeString() + '.' + (ts % 1000).toString().padStart(3, '0');
  };

  const formatMessage = (msg: unknown): string => {
    try {
      return JSON.stringify(msg, null, 2);
    } catch {
      return String(msg);
    }
  };

  const getDirectionColor = (direction: string): string => {
    switch (direction) {
      case 'SEND':
        return 'blue';
      case 'RECV':
        return 'green';
      case 'ERROR':
        return 'red';
      case 'INFO':
        return 'gray';
      default:
        return 'gray';
    }
  };

  return (
    <Paper withBorder p="md">
      <Stack gap="md">
        <Group justify="space-between">
          <Group gap="xs">
            <Text size="sm" fw={600}>
              Live Messages: {proxyName}
            </Text>
            {isConnected && <Badge size="xs" color="green">Connected</Badge>}
            {!isConnected && <Badge size="xs" color="red">Disconnected</Badge>}
            {isPaused && <Badge size="xs" color="yellow">Paused</Badge>}
          </Group>
          <Group gap="xs">
            <Switch
              label="Auto-scroll"
              checked={autoScroll}
              onChange={(e) => setAutoScroll(e.currentTarget.checked)}
              size="sm"
            />
            <ActionIcon
              variant="subtle"
              color={isPaused ? 'green' : 'yellow'}
              onClick={handleTogglePause}
              title={isPaused ? 'Resume' : 'Pause'}
            >
              {isPaused ? <IconPlayerPlay size={16} /> : <IconPlayerPause size={16} />}
            </ActionIcon>
            <ActionIcon
              variant="subtle"
              color="gray"
              onClick={handleClear}
              title="Clear messages"
            >
              <IconTrash size={16} />
            </ActionIcon>
          </Group>
        </Group>

        <Divider />

        <ScrollArea h={height} viewportRef={scrollAreaRef}>
          <Stack gap="xs">
            {messages.length === 0 && (
              <Text size="sm" c="dimmed" ta="center" py="xl">
                Waiting for messages...
              </Text>
            )}
            {messages.map((msg, i) => (
              <Paper key={i} withBorder p="sm" bg={msg.direction === 'SEND' ? 'blue.0' : msg.direction === 'RECV' ? 'green.0' : undefined}>
                <Stack gap="xs">
                  <Group gap="md" wrap="nowrap">
                    <Badge
                      color={getDirectionColor(msg.direction)}
                      variant="filled"
                      size="sm"
                      style={{ minWidth: 60 }}
                    >
                      {msg.direction}
                    </Badge>
                    <Text size="xs" c="dimmed" ff="monospace" style={{ flexShrink: 0 }}>
                      {formatTimestamp(msg.timestamp)}
                    </Text>
                    {msg.parsed?.msgTypeName && (
                      <Badge variant="light" size="xs">
                        {msg.parsed.msgTypeName}
                      </Badge>
                    )}
                    {msg.parsed?.symbol && (
                      <Text size="xs" c="dimmed" ff="monospace" truncate>
                        {msg.parsed.symbol}
                      </Text>
                    )}
                  </Group>
                  <Code block style={{ fontSize: '11px', maxHeight: '200px', overflow: 'auto' }}>
                    {formatMessage(msg.message)}
                  </Code>
                </Stack>
              </Paper>
            ))}
          </Stack>
        </ScrollArea>

        <Group justify="space-between">
          <Text size="xs" c="dimmed">
            {messages.length} messages
            {isPaused && pausedQueueRef.current.length > 0 && ` (${pausedQueueRef.current.length} queued)`}
          </Text>
        </Group>
      </Stack>
    </Paper>
  );
}
