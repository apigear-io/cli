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
  TagsInput,
  Switch,
  SimpleGrid,
} from '@mantine/core';
import {
  IconUsers,
  IconPlus,
  IconRefresh,
} from '@tabler/icons-react';
import {
  useClients,
  useCreateClient,
  useConnectClient,
  useDisconnectClient,
  useDeleteClient,
} from '@/api/queries';
import type { CreateClientRequest } from '@/api/types';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { ClientCard } from './components/ClientCard';
import { notifications } from '@mantine/notifications';

function ClientsContent() {
  const { data: clients } = useClients();
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    url: 'ws://localhost:5560/ws',
    interfaces: [] as string[],
    enabled: true,
    autoReconnect: true,
  });

  const createClient = useCreateClient();
  const connectClient = useConnectClient();
  const disconnectClient = useDisconnectClient();
  const deleteClient = useDeleteClient();

  const handleCreate = async () => {
    try {
      const request: CreateClientRequest = {
        name: formData.name,
        config: {
          url: formData.url,
          interfaces: formData.interfaces,
          enabled: formData.enabled,
          autoReconnect: formData.autoReconnect,
        },
      };

      await createClient.mutateAsync(request);
      notifications.show({
        title: 'Success',
        message: `Client "${formData.name}" created successfully`,
        color: 'green',
      });
      setCreateModalOpen(false);
      setFormData({
        name: '',
        url: 'ws://localhost:5560/ws',
        interfaces: [],
        enabled: true,
        autoReconnect: true,
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to create client',
        color: 'red',
      });
    }
  };

  const handleConnect = async (name: string) => {
    try {
      await connectClient.mutateAsync(name);
      notifications.show({
        title: 'Success',
        message: `Client "${name}" connecting...`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to connect client',
        color: 'red',
      });
    }
  };

  const handleDisconnect = async (name: string) => {
    try {
      await disconnectClient.mutateAsync(name);
      notifications.show({
        title: 'Success',
        message: `Client "${name}" disconnected`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to disconnect client',
        color: 'red',
      });
    }
  };

  const handleDelete = async (name: string) => {
    if (!confirm(`Are you sure you want to delete client "${name}"?`)) {
      return;
    }

    try {
      await deleteClient.mutateAsync(name);
      notifications.show({
        title: 'Success',
        message: `Client "${name}" deleted`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to delete client',
        color: 'red',
      });
    }
  };

  const handleRetry = async (name: string) => {
    // Retry is just reconnecting
    await handleConnect(name);
  };

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="center">
        <Group>
          <Title order={2}>Clients</Title>
        </Group>
        <Group>
          <Button
            leftSection={<IconPlus size={16} />}
            onClick={() => setCreateModalOpen(true)}
          >
            Add Client
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

      {clients.length === 0 ? (
        <Card shadow="sm" padding="xl" radius="md" withBorder>
          <Stack align="center" gap="md">
            <IconUsers size={48} color="var(--mantine-color-gray-5)" />
            <Text size="lg" fw={500} c="dimmed">
              No clients configured
            </Text>
            <Text size="sm" c="dimmed" ta="center">
              Create your first client to connect to ObjectLink backends
            </Text>
            <Button
              leftSection={<IconPlus size={16} />}
              onClick={() => setCreateModalOpen(true)}
            >
              Add Client
            </Button>
          </Stack>
        </Card>
      ) : (
        <SimpleGrid cols={{ base: 1, md: 2, lg: 3 }} spacing="md">
          {clients.map((client) => (
            <ClientCard
              key={client.name}
              client={client}
              onConnect={handleConnect}
              onDisconnect={handleDisconnect}
              onRetry={handleRetry}
              onDelete={handleDelete}
            />
          ))}
        </SimpleGrid>
      )}

      <Modal
        opened={createModalOpen}
        onClose={() => setCreateModalOpen(false)}
        title="Create Client"
        size="md"
      >
        <Stack gap="md">
          <TextInput
            label="Name"
            placeholder="my-client"
            value={formData.name}
            onChange={(e) => setFormData({ ...formData, name: e.target.value })}
            required
          />

          <TextInput
            label="WebSocket URL"
            placeholder="ws://localhost:5560/ws"
            value={formData.url}
            onChange={(e) => setFormData({ ...formData, url: e.target.value })}
            required
          />

          <TagsInput
            label="ObjectLink Interfaces"
            placeholder="Enter interface names"
            value={formData.interfaces}
            onChange={(value) => setFormData({ ...formData, interfaces: value })}
            description="Leave empty to link all available interfaces"
          />

          <Switch
            label="Enabled"
            checked={formData.enabled}
            onChange={(e) => setFormData({ ...formData, enabled: e.target.checked })}
          />

          <Switch
            label="Auto-reconnect"
            checked={formData.autoReconnect}
            onChange={(e) => setFormData({ ...formData, autoReconnect: e.target.checked })}
            description="Automatically reconnect on connection loss"
          />

          <Group justify="flex-end" gap="xs">
            <Button variant="light" onClick={() => setCreateModalOpen(false)}>
              Cancel
            </Button>
            <Button
              onClick={handleCreate}
              loading={createClient.isPending}
              disabled={!formData.name || !formData.url}
            >
              Create
            </Button>
          </Group>
        </Stack>
      </Modal>
    </Stack>
  );
}

export function Clients() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading clients..." />}>
        <ClientsContent />
      </Suspense>
    </ErrorBoundary>
  );
}
