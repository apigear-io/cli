import { Suspense, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import {
  Stack,
  Group,
  Title,
  Button,
  Card,
  Text,
  Badge,
  SimpleGrid,
  ActionIcon,
  Tooltip,
  Tabs,
} from '@mantine/core';
import {
  IconArrowLeft,
  IconPlayerPlay,
  IconPlayerStop,
  IconEdit,
  IconTrash,
  IconUsers,
  IconMessages,
  IconDatabase,
  IconClock,
  IconActivity,
  IconCode,
} from '@tabler/icons-react';
import { useSuspenseQuery } from '@tanstack/react-query';
import { apiClient } from '@/api/client';
import { queryKeys } from '@/api/queryKeys';
import { useStartProxy, useStopProxy, useDeleteProxy } from '@/api/queries';
import type { ProxyInfo } from '@/api/types';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { LiveMessageViewer } from './components/LiveMessageViewer';
import { notifications } from '@mantine/notifications';

function useProxy(name: string) {
  return useSuspenseQuery({
    queryKey: queryKeys.stream.proxies.detail(name),
    queryFn: () => apiClient.get<ProxyInfo>(`/stream/proxies/${encodeURIComponent(name)}`),
    refetchInterval: 2000, // Refresh every 2 seconds
  });
}

function formatBytes(bytes: number): string {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / Math.pow(k, i)).toFixed(1)} ${sizes[i] ?? 'TB'}`;
}

function formatUptime(seconds: number): string {
  if (seconds < 60) return `${seconds}s`;
  if (seconds < 3600) return `${Math.floor(seconds / 60)}m ${seconds % 60}s`;
  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  return `${hours}h ${minutes}m`;
}

function ProxyDetailInner({ name }: { name: string }) {
  const navigate = useNavigate();
  const [activeTab, setActiveTab] = useState<string | null>('overview');

  const { data: proxy } = useProxy(name);
  const startProxy = useStartProxy();
  const stopProxy = useStopProxy();
  const deleteProxy = useDeleteProxy();

  const getStatusColor = () => {
    switch (proxy.status) {
      case 'running':
        return 'green';
      case 'error':
        return 'orange';
      default:
        return 'gray';
    }
  };

  const handleStart = async () => {
    try {
      await startProxy.mutateAsync(name);
      notifications.show({
        title: 'Success',
        message: `Proxy "${name}" started`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to start proxy',
        color: 'red',
      });
    }
  };

  const handleStop = async () => {
    try {
      await stopProxy.mutateAsync(name);
      notifications.show({
        title: 'Success',
        message: `Proxy "${name}" stopped`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to stop proxy',
        color: 'red',
      });
    }
  };

  const handleDelete = async () => {
    if (!confirm(`Are you sure you want to delete proxy "${name}"?`)) {
      return;
    }

    try {
      await deleteProxy.mutateAsync(name);
      notifications.show({
        title: 'Success',
        message: `Proxy "${name}" deleted`,
        color: 'green',
      });
      navigate('/stream/proxies');
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to delete proxy',
        color: 'red',
      });
    }
  };

  return (
    <Stack gap="lg">
      {/* Header */}
      <Group justify="space-between">
        <Group gap="md">
          <ActionIcon
            variant="subtle"
            size="lg"
            onClick={() => navigate('/stream/proxies')}
          >
            <IconArrowLeft size={20} />
          </ActionIcon>
          <div>
            <Group gap="sm" align="center">
              <Title order={2}>{proxy.name}</Title>
              <Badge size="lg" color={getStatusColor()}>
                {proxy.status}
              </Badge>
            </Group>
            <Text size="sm" c="dimmed" mt={4}>
              {proxy.mode} mode
            </Text>
          </div>
        </Group>
        <Group gap="xs">
          {proxy.status === 'running' ? (
            <Button
              leftSection={<IconPlayerStop size={16} />}
              color="red"
              onClick={handleStop}
              loading={stopProxy.isPending}
            >
              Stop
            </Button>
          ) : (
            <Button
              leftSection={<IconPlayerPlay size={16} />}
              color="green"
              onClick={handleStart}
              loading={startProxy.isPending}
            >
              Start
            </Button>
          )}
          <Tooltip label="Edit">
            <ActionIcon variant="light" size="lg" color="gray">
              <IconEdit size={18} />
            </ActionIcon>
          </Tooltip>
          <Tooltip label="Delete">
            <ActionIcon
              variant="light"
              size="lg"
              color="red"
              onClick={handleDelete}
              disabled={proxy.status === 'running'}
            >
              <IconTrash size={18} />
            </ActionIcon>
          </Tooltip>
        </Group>
      </Group>

      {/* Stats Cards */}
      <SimpleGrid cols={{ base: 2, sm: 3, lg: 6 }} spacing="md">
        <Card shadow="sm" padding="md" radius="md" withBorder>
          <Stack gap="xs">
            <Group gap="xs">
              <IconUsers size={16} color="var(--mantine-color-blue-6)" />
              <Text size="xs" c="dimmed" fw={500}>
                Connections
              </Text>
            </Group>
            <Text size="xl" fw={700}>
              {proxy.activeConnections}
            </Text>
          </Stack>
        </Card>

        <Card shadow="sm" padding="md" radius="md" withBorder>
          <Stack gap="xs">
            <Group gap="xs">
              <IconMessages size={16} color="var(--mantine-color-green-6)" />
              <Text size="xs" c="dimmed" fw={500}>
                Messages
              </Text>
            </Group>
            <Text size="xl" fw={700}>
              {proxy.messagesReceived + proxy.messagesSent}
            </Text>
            <Text size="xs" c="dimmed">
              ↓{proxy.messagesReceived} ↑{proxy.messagesSent}
            </Text>
          </Stack>
        </Card>

        <Card shadow="sm" padding="md" radius="md" withBorder>
          <Stack gap="xs">
            <Group gap="xs">
              <IconDatabase size={16} color="var(--mantine-color-violet-6)" />
              <Text size="xs" c="dimmed" fw={500}>
                Data
              </Text>
            </Group>
            <Text size="xl" fw={700}>
              {formatBytes(proxy.bytesReceived + proxy.bytesSent)}
            </Text>
            <Text size="xs" c="dimmed">
              ↓{formatBytes(proxy.bytesReceived)} ↑{formatBytes(proxy.bytesSent)}
            </Text>
          </Stack>
        </Card>

        <Card shadow="sm" padding="md" radius="md" withBorder>
          <Stack gap="xs">
            <Group gap="xs">
              <IconClock size={16} color="var(--mantine-color-orange-6)" />
              <Text size="xs" c="dimmed" fw={500}>
                Uptime
              </Text>
            </Group>
            <Text size="xl" fw={700}>
              {formatUptime(proxy.uptime)}
            </Text>
          </Stack>
        </Card>

        <Card shadow="sm" padding="md" radius="md" withBorder>
          <Stack gap="xs">
            <Group gap="xs">
              <IconActivity size={16} color="var(--mantine-color-teal-6)" />
              <Text size="xs" c="dimmed" fw={500}>
                Msg Rate
              </Text>
            </Group>
            <Text size="xl" fw={700}>
              {proxy.uptime > 0
                ? ((proxy.messagesReceived + proxy.messagesSent) / proxy.uptime).toFixed(1)
                : '0'}
            </Text>
            <Text size="xs" c="dimmed">
              msg/s
            </Text>
          </Stack>
        </Card>

        <Card shadow="sm" padding="md" radius="md" withBorder>
          <Stack gap="xs">
            <Group gap="xs">
              <IconCode size={16} color="var(--mantine-color-pink-6)" />
              <Text size="xs" c="dimmed" fw={500}>
                Mode
              </Text>
            </Group>
            <Text size="lg" fw={700} tt="capitalize">
              {proxy.mode}
            </Text>
          </Stack>
        </Card>
      </SimpleGrid>

      {/* Configuration */}
      <Card shadow="sm" padding="lg" radius="md" withBorder>
        <Stack gap="md">
          <Title order={4}>Configuration</Title>
          <SimpleGrid cols={{ base: 1, sm: 2 }} spacing="md">
            <div>
              <Text size="xs" c="dimmed" fw={500}>
                Listen Address
              </Text>
              <Text size="sm" mt={4}>
                {proxy.listen}
              </Text>
            </div>
            {proxy.backend && (
              <div>
                <Text size="xs" c="dimmed" fw={500}>
                  Backend Address
                </Text>
                <Text size="sm" mt={4}>
                  {proxy.backend}
                </Text>
              </div>
            )}
          </SimpleGrid>
        </Stack>
      </Card>

      {/* Tabs for Traces and Live Messages */}
      <Tabs value={activeTab} onChange={setActiveTab}>
        <Tabs.List>
          <Tabs.Tab value="overview">Overview</Tabs.Tab>
          <Tabs.Tab value="live">Live Messages</Tabs.Tab>
          <Tabs.Tab value="traces">Trace Files</Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="overview" pt="lg">
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="md">
              <Title order={4}>Proxy Overview</Title>
              <Text size="sm" c="dimmed">
                This proxy is configured in {proxy.mode} mode and is currently {proxy.status}.
              </Text>
              {proxy.status === 'stopped' && (
                <Text size="sm" c="dimmed">
                  Click the Start button above to begin proxying WebSocket connections.
                </Text>
              )}
            </Stack>
          </Card>
        </Tabs.Panel>

        <Tabs.Panel value="live" pt="lg">
          {proxy.status === 'running' ? (
            <Card shadow="sm" padding="md" radius="md" withBorder>
              <LiveMessageViewer proxyName={proxy.name} height="60vh" />
            </Card>
          ) : (
            <Card shadow="sm" padding="lg" radius="md" withBorder>
              <Stack align="center" gap="md">
                <IconMessages size={48} color="var(--mantine-color-gray-5)" />
                <Text size="lg" fw={500} c="dimmed">
                  Proxy is not running
                </Text>
                <Text size="sm" c="dimmed" ta="center">
                  Start the proxy to view live messages
                </Text>
              </Stack>
            </Card>
          )}
        </Tabs.Panel>

        <Tabs.Panel value="traces" pt="lg">
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack align="center" gap="md">
              <IconDatabase size={48} color="var(--mantine-color-gray-5)" />
              <Text size="lg" fw={500} c="dimmed">
                Trace files coming soon
              </Text>
              <Text size="sm" c="dimmed" ta="center">
                Trace file management will be available in a future update
              </Text>
            </Stack>
          </Card>
        </Tabs.Panel>
      </Tabs>
    </Stack>
  );
}

function ProxyDetailContent() {
  const { name } = useParams<{ name: string }>();
  const navigate = useNavigate();

  if (!name) {
    navigate('/stream/proxies');
    return null;
  }

  return <ProxyDetailInner name={name} />;
}

export function ProxyDetail() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading proxy details..." />}>
        <ProxyDetailContent />
      </Suspense>
    </ErrorBoundary>
  );
}
