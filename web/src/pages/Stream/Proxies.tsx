import { Suspense, useState } from 'react';
import {
  Card,
  Text,
  Title,
  Stack,
  Group,
  Button,
  Modal,
  TextInput,
  Select,
  SimpleGrid,
  Drawer,
} from '@mantine/core';
import {
  IconServer,
  IconPlus,
  IconRefresh,
} from '@tabler/icons-react';
import {
  useProxies,
  useCreateProxy,
  useUpdateProxy,
  useDeleteProxy,
  useStartProxy,
  useStopProxy,
} from '@/api/queries';
import type { ProxyMode, CreateProxyRequest, ProxyInfo } from '@/api/types';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { LiveMessageViewer } from './components/LiveMessageViewer';
import { ProxyCard } from './components/ProxyCard';
import { notifications } from '@mantine/notifications';

function ProxiesContent() {
  const { data: proxies } = useProxies();
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [editDrawerOpen, setEditDrawerOpen] = useState(false);
  const [viewerOpen, setViewerOpen] = useState(false);
  const [selectedProxy, setSelectedProxy] = useState<ProxyInfo | null>(null);
  const [formData, setFormData] = useState({
    name: '',
    listen: 'ws://localhost:5550/ws',
    backend: 'ws://localhost:5560/ws',
    mode: 'proxy' as ProxyMode,
  });

  const createProxy = useCreateProxy();
  const updateProxy = useUpdateProxy();
  const deleteProxy = useDeleteProxy();
  const startProxy = useStartProxy();
  const stopProxy = useStopProxy();

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

  const handleEdit = (proxy: ProxyInfo) => {
    setSelectedProxy(proxy);
    setFormData({
      name: proxy.name,
      listen: proxy.listen,
      backend: proxy.backend || '',
      mode: proxy.mode,
    });
    setEditDrawerOpen(true);
  };

  const handleUpdate = async () => {
    if (!selectedProxy) return;

    try {
      await updateProxy.mutateAsync({
        name: selectedProxy.name,
        config: {
          listen: formData.listen,
          backend: formData.mode === 'proxy' ? formData.backend : undefined,
          mode: formData.mode,
        },
      });
      notifications.show({
        title: 'Success',
        message: `Proxy "${selectedProxy.name}" updated successfully`,
        color: 'green',
      });
      setEditDrawerOpen(false);
      setSelectedProxy(null);
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to update proxy',
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

  const handleViewStats = (name: string) => {
    const proxy = proxies.find((p) => p.name === name);
    if (proxy) {
      setSelectedProxy(proxy);
      setViewerOpen(true);
    }
  };

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="center">
        <Group>
          <Title order={2}>Proxies</Title>
        </Group>
        <Group>
          <Button
            leftSection={<IconPlus size={16} />}
            onClick={() => setCreateModalOpen(true)}
          >
            Add Proxy
          </Button>
          <Button
            leftSection={<IconRefresh size={16} />}
            variant="light"
            onClick={() => window.location.reload()}
          >
            Refresh
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
              Add Proxy
            </Button>
          </Stack>
        </Card>
      ) : (
        <SimpleGrid cols={{ base: 1, md: 2, lg: 3 }} spacing="md">
          {proxies.map((proxy) => (
            <ProxyCard
              key={proxy.name}
              proxy={proxy}
              onViewStats={handleViewStats}
              onEdit={handleEdit}
              onDelete={handleDelete}
              onStart={handleStart}
              onStop={handleStop}
            />
          ))}
        </SimpleGrid>
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

      {/* Edit Proxy Drawer */}
      <Drawer
        opened={editDrawerOpen}
        onClose={() => {
          setEditDrawerOpen(false);
          setSelectedProxy(null);
        }}
        title="Edit Proxy"
        position="right"
        size="md"
      >
        <Stack gap="md">
          <TextInput
            label="Name"
            value={formData.name}
            disabled
            description="Proxy name cannot be changed"
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

          <Group justify="flex-end" gap="xs" mt="md">
            <Button
              variant="light"
              onClick={() => {
                setEditDrawerOpen(false);
                setSelectedProxy(null);
              }}
            >
              Cancel
            </Button>
            <Button
              onClick={handleUpdate}
              loading={updateProxy.isPending}
              disabled={!formData.listen}
            >
              Update
            </Button>
          </Group>
        </Stack>
      </Drawer>

      {/* Live Message Viewer Modal */}
      <Modal
        opened={viewerOpen}
        onClose={() => {
          setViewerOpen(false);
          setSelectedProxy(null);
        }}
        title={`Live Messages: ${selectedProxy?.name}`}
        size="xl"
      >
        {selectedProxy && <LiveMessageViewer proxyName={selectedProxy.name} height="70vh" />}
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
