import { Card, Text, Badge, Group, Stack, ActionIcon, Tooltip } from '@mantine/core';
import {
  IconEdit,
  IconTrash,
  IconPlugConnected,
  IconPlugConnectedX,
  IconRefresh,
} from '@tabler/icons-react';
import type { ClientInfo } from '@/api/types';

interface ClientCardProps {
  client: ClientInfo;
  onConnect?: (name: string) => void;
  onDisconnect?: (name: string) => void;
  onRetry?: (name: string) => void;
  onEdit?: (client: ClientInfo) => void;
  onDelete?: (name: string) => void;
}

export function ClientCard({
  client,
  onConnect,
  onDisconnect,
  onRetry,
  onEdit,
  onDelete,
}: ClientCardProps) {
  const getStatusDotColor = () => {
    switch (client.status) {
      case 'connected':
        return 'var(--mantine-color-green-6)';
      case 'connecting':
        return 'var(--mantine-color-yellow-6)';
      default:
        return 'var(--mantine-color-red-6)';
    }
  };

  const showRetryBadge = client.status === 'error' || (client.status === 'disconnected' && client.autoReconnect);

  return (
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Stack gap="md">
        {/* Header with status dot, name, and badges */}
        <Group justify="space-between" wrap="nowrap">
          <Group gap="sm" wrap="nowrap" style={{ flex: 1, minWidth: 0 }}>
            {/* Status dot */}
            <div
              style={{
                width: 8,
                height: 8,
                borderRadius: '50%',
                backgroundColor: getStatusDotColor(),
                flexShrink: 0,
              }}
            />
            <Stack gap={4} style={{ flex: 1, minWidth: 0 }}>
              <Text fw={600} size="md" lineClamp={1}>
                {client.name}
              </Text>
              <Group gap="xs">
                <Badge size="sm" color="violet" variant="filled">
                  DriverAssistance
                </Badge>
                {showRetryBadge && (
                  <Badge size="sm" color="orange" variant="light">
                    Retry #1
                  </Badge>
                )}
              </Group>
            </Stack>
          </Group>

          {/* Action buttons */}
          <Group gap={4} wrap="nowrap">
            {showRetryBadge && onRetry && (
              <Tooltip label="Retry connection">
                <ActionIcon
                  variant="light"
                  color="orange"
                  size="sm"
                  onClick={() => onRetry(client.name)}
                >
                  <IconRefresh size={14} />
                </ActionIcon>
              </Tooltip>
            )}
            {client.status === 'connected' ? (
              <Tooltip label="Disconnect">
                <ActionIcon
                  variant="light"
                  color="orange"
                  size="sm"
                  onClick={() => onDisconnect?.(client.name)}
                >
                  <IconPlugConnectedX size={14} />
                </ActionIcon>
              </Tooltip>
            ) : (
              <Tooltip label="Connect">
                <ActionIcon
                  variant="light"
                  color="green"
                  size="sm"
                  onClick={() => onConnect?.(client.name)}
                  disabled={!client.enabled}
                >
                  <IconPlugConnected size={14} />
                </ActionIcon>
              </Tooltip>
            )}
            {onEdit && (
              <Tooltip label="Edit">
                <ActionIcon
                  variant="light"
                  color="gray"
                  size="sm"
                  onClick={() => onEdit(client)}
                >
                  <IconEdit size={14} />
                </ActionIcon>
              </Tooltip>
            )}
            {onDelete && (
              <Tooltip label="Delete">
                <ActionIcon
                  variant="light"
                  color="red"
                  size="sm"
                  onClick={() => onDelete(client.name)}
                  disabled={client.status === 'connected'}
                >
                  <IconTrash size={14} />
                </ActionIcon>
              </Tooltip>
            )}
          </Group>
        </Group>

        {/* WebSocket URL */}
        <Text size="xs" c="dimmed" lineClamp={1}>
          {client.url}
        </Text>

        {/* Interfaces */}
        {client.interfaces.length > 0 && (
          <Group gap="xs">
            {client.interfaces.map((iface) => (
              <Badge key={iface} size="sm" variant="light" color="violet">
                {iface}
              </Badge>
            ))}
          </Group>
        )}
      </Stack>
    </Card>
  );
}
