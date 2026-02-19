import { Suspense, useState } from 'react';
import {
  Card,
  Grid,
  Text,
  Title,
  Stack,
  Group,
  Badge,
  Button,
  Modal,
  TextInput,
  Select,
  ActionIcon,
  Tooltip,
} from '@mantine/core';
import {
  IconServer,
  IconPlayerPlay,
  IconPlayerStop,
  IconTrash,
  IconPlus,
  IconRefresh,
  IconEye,
} from '@tabler/icons-react';
import {
  useProxies,
  useCreateProxy,
  useStartProxy,
  useStopProxy,
  useDeleteProxy,
} from '@/api/queries';
import type { ProxyMode, CreateProxyRequest } from '@/api/types';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { LiveMessageViewer } from './components/LiveMessageViewer';
import { notifications } from '@mantine/notifications';

function ProxiesContent() {
  const { data: proxies } = useProxies();
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [viewerOpen, setViewerOpen] = useState(false);
  const [selectedProxy, setSelectedProxy] = useState<string | null>(null);
  const [formData, setFormData] = useState({
    name: '',
    listen: 'ws://localhost:5550/ws',
    backend: 'ws://localhost:5560/ws',
    mode: 'proxy' as ProxyMode,
  });

  const createProxy = useCreateProxy();
  const startProxy = useStartProxy();
  const stopProxy = useStopProxy();
  const deleteProxy = useDeleteProxy();

  const handleCreate = async () => {
    try {
      const request: CreateProxyRequest = {
        name: formData.name,
        config: {
          listen: formData.listen,
          backend: formData.mode === 'proxy' ? formData.backend : undefined,
          mode: formData.mode,
        },
      };

      await createProxy.mutateAsync(request);
      notifications.show({
        title: 'Success',
        message: `Proxy "${formData.name}" created successfully`,
        color: 'green',
      });
      setCreateModalOpen(false);
      setFormData({
        name: '',
        listen: 'ws://localhost:5550/ws',
        backend: 'ws://localhost:5560/ws',
        mode: 'proxy',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to create proxy',
        color: 'red',
      });
    }
  };

  const handleStart = async (name: string) => {
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

  const handleStop = async (name: string) => {
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

  const handleDelete = async (name: string) => {
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
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to delete proxy',
        color: 'red',
      });
    }
  };

  const handleViewMessages = (name: string) => {
    setSelectedProxy(name);
    setViewerOpen(true);
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'running':
        return 'green';
      case 'stopped':
        return 'gray';
      case 'error':
        return 'red';
      default:
        return 'gray';
    }
  };

  const getModeColor = (mode: string) => {
    switch (mode) {
      case 'proxy':
        return 'blue';
      case 'echo':
        return 'cyan';
      case 'backend':
        return 'violet';
      case 'inbound-only':
        return 'orange';
      default:
        return 'gray';
    }
  };

  const formatUptime = (seconds: number) => {
    if (seconds < 60) return `${seconds}s`;
    if (seconds < 3600) return `${Math.floor(seconds / 60)}m`;
    const hours = Math.floor(seconds / 3600);
    const mins = Math.floor((seconds % 3600) / 60);
    return `${hours}h ${mins}m`;
  };

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="center">
        <Group>
          <Title order={2}>Proxies</Title>
          <Badge size="lg" variant="light" color="blue">
            {proxies.length} total
          </Badge>
        </Group>
        <Group>
          <Button
            leftSection={<IconRefresh size={16} />}
            variant="light"
            onClick={() => window.location.reload()}
          >
            Refresh
          </Button>
          <Button
            leftSection={<IconPlus size={16} />}
            onClick={() => setCreateModalOpen(true)}
          >
            Create Proxy
          </Button>
        </Group>
      </Group>

      {proxies.length === 0 ? (
        <Card shadow="sm" padding="xl" radius="md" withBorder>
          <Stack align="center" gap="md">
            <IconServer size={48} color="var(--mantine-color-gray-5)" />
            <Text size="lg" fw={500} c="dimmed">
              No proxies configured
            </Text>
            <Text size="sm" c="dimmed" ta="center">
              Create your first proxy to start forwarding WebSocket connections
            </Text>
            <Button
              leftSection={<IconPlus size={16} />}
              onClick={() => setCreateModalOpen(true)}
            >
              Create Proxy
            </Button>
          </Stack>
        </Card>
      ) : (
        <Grid>
          {proxies.map((proxy) => (
            <Grid.Col key={proxy.name} span={{ base: 12, md: 6, lg: 4 }}>
              <Card shadow="sm" padding="lg" radius="md" withBorder>
                <Stack gap="md">
                  <Group justify="space-between" align="flex-start">
                    <Stack gap={4}>
                      <Group gap="xs">
                        <IconServer size={20} color="var(--mantine-color-blue-6)" />
                        <Text fw={600} size="lg">
                          {proxy.name}
                        </Text>
                      </Group>
                      <Group gap="xs">
                        <Badge size="sm" color={getStatusColor(proxy.status)}>
                          {proxy.status}
                        </Badge>
                        <Badge size="sm" color={getModeColor(proxy.mode)} variant="light">
                          {proxy.mode}
                        </Badge>
                      </Group>
                    </Stack>
                    <Tooltip label="Delete proxy">
                      <ActionIcon
                        color="red"
                        variant="subtle"
                        onClick={() => handleDelete(proxy.name)}
                        disabled={proxy.status === 'running'}
                      >
                        <IconTrash size={18} />
                      </ActionIcon>
                    </Tooltip>
                  </Group>

                  <Stack gap="xs">
                    <Group gap="xs">
                      <Text size="xs" c="dimmed" fw={500}>
                        Listen:
                      </Text>
                      <Text size="xs" fw={500}>
                        {proxy.listen}
                      </Text>
                    </Group>
                    {proxy.backend && (
                      <Group gap="xs">
                        <Text size="xs" c="dimmed" fw={500}>
                          Backend:
                        </Text>
                        <Text size="xs" fw={500}>
                          {proxy.backend}
                        </Text>
                      </Group>
                    )}
                  </Stack>

                  <Group grow>
                    <Stack gap={4}>
                      <Text size="xs" c="dimmed" ta="center">
                        Messages
                      </Text>
                      <Text size="sm" fw={600} ta="center">
                        ↓{proxy.messagesReceived} ↑{proxy.messagesSent}
                      </Text>
                    </Stack>
                    <Stack gap={4}>
                      <Text size="xs" c="dimmed" ta="center">
                        Connections
                      </Text>
                      <Text size="sm" fw={600} ta="center">
                        {proxy.activeConnections}
                      </Text>
                    </Stack>
                    <Stack gap={4}>
                      <Text size="xs" c="dimmed" ta="center">
                        Uptime
                      </Text>
                      <Text size="sm" fw={600} ta="center">
                        {proxy.status === 'running' ? formatUptime(proxy.uptime) : '-'}
                      </Text>
                    </Stack>
                  </Group>

                  <Group grow>
                    {proxy.status === 'running' ? (
                      <Button
                        leftSection={<IconPlayerStop size={16} />}
                        color="orange"
                        variant="light"
                        size="sm"
                        onClick={() => handleStop(proxy.name)}
                        loading={stopProxy.isPending}
                      >
                        Stop
                      </Button>
                    ) : (
                      <Button
                        leftSection={<IconPlayerPlay size={16} />}
                        color="green"
                        variant="light"
                        size="sm"
                        onClick={() => handleStart(proxy.name)}
                        loading={startProxy.isPending}
                      >
                        Start
                      </Button>
                    )}
                  </Group>

                  <Button
                    leftSection={<IconEye size={16} />}
                    variant="light"
                    size="sm"
                    onClick={() => handleViewMessages(proxy.name)}
                    disabled={proxy.status !== 'running'}
                  >
                    View Messages
                  </Button>
                </Stack>
              </Card>
            </Grid.Col>
          ))}
        </Grid>
      )}

      <Modal
        opened={createModalOpen}
        onClose={() => setCreateModalOpen(false)}
        title="Create Proxy"
        size="md"
      >
        <Stack gap="md">
          <TextInput
            label="Name"
            placeholder="my-proxy"
            value={formData.name}
            onChange={(e) => setFormData({ ...formData, name: e.target.value })}
            required
          />

          <Select
            label="Mode"
            value={formData.mode}
            onChange={(value) => setFormData({ ...formData, mode: value as ProxyMode })}
            data={[
              { value: 'proxy', label: 'Proxy - Forward to backend' },
              { value: 'echo', label: 'Echo - Echo back messages' },
              { value: 'backend', label: 'Backend - ObjectLink backend' },
              { value: 'inbound-only', label: 'Inbound Only - Accept only' },
            ]}
            required
          />

          <TextInput
            label="Listen Address"
            placeholder="ws://localhost:5550/ws"
            value={formData.listen}
            onChange={(e) => setFormData({ ...formData, listen: e.target.value })}
            required
          />

          {formData.mode === 'proxy' && (
            <TextInput
              label="Backend Address"
              placeholder="ws://localhost:5560/ws"
              value={formData.backend}
              onChange={(e) => setFormData({ ...formData, backend: e.target.value })}
              required
            />
          )}

          <Group justify="flex-end" gap="xs">
            <Button variant="light" onClick={() => setCreateModalOpen(false)}>
              Cancel
            </Button>
            <Button
              onClick={handleCreate}
              loading={createProxy.isPending}
              disabled={!formData.name || !formData.listen}
            >
              Create
            </Button>
          </Group>
        </Stack>
      </Modal>

      {/* Live Message Viewer Modal */}
      <Modal
        opened={viewerOpen}
        onClose={() => setViewerOpen(false)}
        title={`Live Messages: ${selectedProxy}`}
        size="xl"
      >
        {selectedProxy && <LiveMessageViewer proxyName={selectedProxy} height="70vh" />}
      </Modal>
    </Stack>
  );
}

export function Proxies() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading proxies..." />}>
        <ProxiesContent />
      </Suspense>
    </ErrorBoundary>
  );
}
