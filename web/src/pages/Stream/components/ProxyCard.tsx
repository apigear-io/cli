import { Card, Text, Badge, Group, Stack, ActionIcon, Tooltip } from '@mantine/core';
import {
  IconEdit,
  IconTrash,
  IconChartLine,
  IconArrowRight,
  IconUsers,
  IconMessages,
  IconDatabase,
} from '@tabler/icons-react';
import type { ProxyInfo } from '@/api/types';

interface ProxyCardProps {
  proxy: ProxyInfo;
  onViewStats?: (name: string) => void;
  onEdit?: (proxy: ProxyInfo) => void;
  onDelete?: (name: string) => void;
}

function formatBytes(bytes: number): string {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / Math.pow(k, i)).toFixed(1)} ${sizes[i] ?? 'TB'}`;
}

export function ProxyCard({ proxy, onViewStats, onEdit, onDelete }: ProxyCardProps) {
  const getStatusDotColor = () => {
    switch (proxy.status) {
      case 'running':
        return 'var(--mantine-color-green-6)';
      case 'error':
        return 'var(--mantine-color-orange-6)';
      default:
        return 'var(--mantine-color-red-6)';
    }
  };

  const getStatusBadgeColor = () => {
    switch (proxy.status) {
      case 'running':
        return 'green';
      case 'error':
        return 'orange';
      default:
        return 'gray';
    }
  };

  const getStatusLabel = () => {
    switch (proxy.status) {
      case 'running':
        return 'Running';
      case 'error':
        return 'Retrying';
      default:
        return 'Stopped';
    }
  };

  const getStatusMessage = () => {
    if (proxy.status === 'running') {
      return `${proxy.activeConnections} active connections`;
    }
    if (proxy.status === 'error') {
      return 'Retry #13 in 2s';
    }
    return 'Proxy not started';
  };

  return (
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Stack gap="md">
        {/* Header with status dot and badges */}
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
              <Group gap="xs">
                <Text fw={600} size="md" lineClamp={1}>
                  {proxy.name}
                </Text>
                <Badge size="sm" color={getStatusBadgeColor()}>
                  {getStatusLabel()}
                </Badge>
              </Group>
              <Text size="xs" c="dimmed" lineClamp={1}>
                {getStatusMessage()}
              </Text>
            </Stack>
          </Group>

          {/* Action buttons */}
          <Group gap={4} wrap="nowrap">
            {onViewStats && (
              <Tooltip label="View statistics">
                <ActionIcon
                  variant="light"
                  color="blue"
                  size="sm"
                  onClick={() => onViewStats(proxy.name)}
                >
                  <IconChartLine size={14} />
                </ActionIcon>
              </Tooltip>
            )}
            {onEdit && (
              <Tooltip label="Edit">
                <ActionIcon
                  variant="light"
                  color="gray"
                  size="sm"
                  onClick={() => onEdit(proxy)}
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
                  onClick={() => onDelete(proxy.name)}
                  disabled={proxy.status === 'running'}
                >
                  <IconTrash size={14} />
                </ActionIcon>
              </Tooltip>
            )}
          </Group>
        </Group>

        {/* IN → OUT addresses */}
        <Group gap="md" wrap="nowrap" align="center">
          <Stack gap={0} style={{ flex: 1, minWidth: 0 }}>
            <Text size="xs" c="green" fw={500}>
              IN
            </Text>
            <Text size="xs" c="dimmed" lineClamp={1}>
              {proxy.listen || 'none'}
            </Text>
          </Stack>

          <IconArrowRight size={16} color="var(--mantine-color-gray-6)" style={{ flexShrink: 0 }} />

          <Stack gap={0} style={{ flex: 1, minWidth: 0 }}>
            <Text size="xs" c="blue" fw={500}>
              OUT
            </Text>
            <Text size="xs" c="dimmed" lineClamp={1}>
              {proxy.backend || 'none'}
            </Text>
          </Stack>
        </Group>

        {/* Stats icons */}
        <Group gap="lg" justify="space-between">
          <Group gap={4}>
            <IconUsers size={14} color="var(--mantine-color-gray-6)" />
            <Text size="xs" c="dimmed">
              {proxy.activeConnections} / {proxy.activeConnections}
            </Text>
          </Group>
          <Group gap={4}>
            <IconMessages size={14} color="var(--mantine-color-gray-6)" />
            <Text size="xs" c="dimmed">
              {proxy.messagesReceived + proxy.messagesSent}
            </Text>
          </Group>
          <Group gap={4}>
            <IconDatabase size={14} color="var(--mantine-color-gray-6)" />
            <Text size="xs" c="dimmed">
              {formatBytes(proxy.bytesReceived + proxy.bytesSent)}
            </Text>
          </Group>
        </Group>
      </Stack>
    </Card>
  );
}
