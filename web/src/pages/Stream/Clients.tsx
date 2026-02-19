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
  TagsInput,
  ActionIcon,
  Tooltip,
  Switch,
} from '@mantine/core';
import {
  IconUsers,
  IconPlugConnected,
  IconPlugConnectedX,
  IconTrash,
  IconPlus,
  IconRefresh,
  IconAlertCircle,
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

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'connected':
        return 'green';
      case 'connecting':
        return 'yellow';
      case 'disconnected':
        return 'gray';
      case 'error':
        return 'red';
      default:
        return 'gray';
    }
  };

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="center">
        <Group>
          <Title order={2}>Clients</Title>
          <Badge size="lg" variant="light" color="cyan">
            {clients.length} total
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
            Create Client
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
              Create Client
            </Button>
          </Stack>
        </Card>
      ) : (
        <Grid>
          {clients.map((client) => (
            <Grid.Col key={client.name} span={{ base: 12, md: 6, lg: 4 }}>
              <Card shadow="sm" padding="lg" radius="md" withBorder>
                <Stack gap="md">
                  <Group justify="space-between" align="flex-start">
                    <Stack gap={4}>
                      <Group gap="xs">
                        <IconUsers size={20} color="var(--mantine-color-cyan-6)" />
                        <Text fw={600} size="lg">
                          {client.name}
                        </Text>
                      </Group>
                      <Group gap="xs">
                        <Badge size="sm" color={getStatusColor(client.status)}>
                          {client.status}
                        </Badge>
                        {!client.enabled && (
                          <Badge size="sm" color="gray" variant="light">
                            disabled
                          </Badge>
                        )}
                        {client.autoReconnect && (
                          <Badge size="sm" color="blue" variant="light">
                            auto-reconnect
                          </Badge>
                        )}
                      </Group>
                    </Stack>
                    <Tooltip label="Delete client">
                      <ActionIcon
                        color="red"
                        variant="subtle"
                        onClick={() => handleDelete(client.name)}
                        disabled={client.status === 'connected'}
                      >
                        <IconTrash size={18} />
                      </ActionIcon>
                    </Tooltip>
                  </Group>

                  <Stack gap="xs">
                    <Group gap="xs">
                      <Text size="xs" c="dimmed" fw={500}>
                        URL:
                      </Text>
                      <Text size="xs" fw={500}>
                        {client.url}
                      </Text>
                    </Group>
                  </Stack>

                  {client.interfaces.length > 0 && (
                    <Stack gap={4}>
                      <Text size="xs" c="dimmed" fw={500}>
                        Interfaces:
                      </Text>
                      <Group gap="xs">
                        {client.interfaces.map((iface) => (
                          <Badge key={iface} size="xs" variant="light" color="violet">
                            {iface}
                          </Badge>
                        ))}
                      </Group>
                    </Stack>
                  )}

                  {client.lastError && (
                    <Group gap="xs" wrap="nowrap">
                      <IconAlertCircle size={16} color="var(--mantine-color-red-6)" />
                      <Text size="xs" c="red" lineClamp={2}>
                        {client.lastError}
                      </Text>
                    </Group>
                  )}

                  <Group grow>
                    {client.status === 'connected' ? (
                      <Button
                        leftSection={<IconPlugConnectedX size={16} />}
                        color="orange"
                        variant="light"
                        size="sm"
                        onClick={() => handleDisconnect(client.name)}
                        loading={disconnectClient.isPending}
                      >
                        Disconnect
                      </Button>
                    ) : (
                      <Button
                        leftSection={<IconPlugConnected size={16} />}
                        color="green"
                        variant="light"
                        size="sm"
                        onClick={() => handleConnect(client.name)}
                        loading={connectClient.isPending}
                        disabled={!client.enabled}
                      >
                        Connect
                      </Button>
                    )}
                  </Group>
                </Stack>
              </Card>
            </Grid.Col>
          ))}
        </Grid>
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
