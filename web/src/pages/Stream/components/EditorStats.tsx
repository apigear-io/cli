import { Paper, Group, Text } from '@mantine/core';
import { useEditorContext } from './EditorContext';

export function EditorStats() {
  const { sessionStats } = useEditorContext();

  if (!sessionStats) return null;

  return (
    <Paper p="sm" withBorder style={{ backgroundColor: 'var(--mantine-color-gray-0)' }}>
      <Group gap="xl" wrap="nowrap">
        <Group gap={4}>
          <Text size="sm" c="dimmed">
            File
          </Text>
          <Text size="sm" fw={500}>
            {sessionStats.filename}
          </Text>
        </Group>

        <Group gap={4}>
          <Text size="sm" c="dimmed">
            Messages
          </Text>
          <Text size="sm" fw={500}>
            {sessionStats.totalCount.toLocaleString()}
          </Text>
        </Group>

        <Group gap={4}>
          <Text size="sm" c="dimmed">
            Time Range
          </Text>
          <Text size="sm" fw={500}>
            -
          </Text>
        </Group>

        <Group gap={4}>
          <Text size="sm" c="dimmed">
            Proxies
          </Text>
          <Text size="sm" fw={500}>
            {sessionStats.proxies.length > 0 ? sessionStats.proxies.join(', ') : '-'}
          </Text>
        </Group>

        <Group gap={4}>
          <Text size="sm" c="dimmed">
            Interfaces
          </Text>
          <Text size="sm" fw={500}>
            {sessionStats.interfaces.length > 0 ? sessionStats.interfaces.join(', ') : '-'}
          </Text>
        </Group>
      </Group>
    </Paper>
  );
}
