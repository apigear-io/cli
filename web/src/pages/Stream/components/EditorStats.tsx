import { Paper, Group, Stack, Text, Badge } from '@mantine/core';
import { useEditorContext } from './EditorContext';

export function EditorStats() {
  const { sessionStats } = useEditorContext();

  if (!sessionStats) return null;

  const formatTimestamp = (ts: number) => {
    return new Date(ts).toLocaleString();
  };

  const formatDuration = () => {
    const durationMs = sessionStats.timeRange.end - sessionStats.timeRange.start;
    const seconds = Math.floor(durationMs / 1000);
    const minutes = Math.floor(seconds / 60);
    const hours = Math.floor(minutes / 60);

    if (hours > 0) {
      return `${hours}h ${minutes % 60}m`;
    } else if (minutes > 0) {
      return `${minutes}m ${seconds % 60}s`;
    }
    return `${seconds}s`;
  };

  return (
    <Paper p="md" withBorder>
      <Group justify="space-between" wrap="wrap" gap="md">
        <Stack gap={4}>
          <Text size="xs" c="dimmed">
            File
          </Text>
          <Text fw={500}>{sessionStats.filename}</Text>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">
            Messages
          </Text>
          <Text fw={500}>{sessionStats.totalCount.toLocaleString()}</Text>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">
            Duration
          </Text>
          <Text fw={500}>{formatDuration()}</Text>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">
            Time Range
          </Text>
          <Text fw={500} size="sm">
            {formatTimestamp(sessionStats.timeRange.start)}
          </Text>
          <Text fw={500} size="sm">
            {formatTimestamp(sessionStats.timeRange.end)}
          </Text>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">
            Proxies
          </Text>
          <Group gap="xs">
            {sessionStats.proxies.map((p) => (
              <Badge key={p} size="sm" variant="light">
                {p}
              </Badge>
            ))}
          </Group>
        </Stack>

        <Stack gap={4}>
          <Text size="xs" c="dimmed">
            Interfaces
          </Text>
          <Group gap="xs">
            {sessionStats.interfaces.slice(0, 3).map((i) => (
              <Badge key={i} size="sm" variant="light" color="green">
                {i}
              </Badge>
            ))}
            {sessionStats.interfaces.length > 3 && (
              <Badge size="sm" variant="light" color="gray">
                +{sessionStats.interfaces.length - 3}
              </Badge>
            )}
          </Group>
        </Stack>
      </Group>
    </Paper>
  );
}
