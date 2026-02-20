/**
 * EditorWelcome - Welcome screen for stream editor
 *
 * Shown when no session is loaded. Provides overview of features
 * and prompts user to load a trace file.
 */

import { Stack, Title, Text, Button, Paper, SimpleGrid } from '@mantine/core';
import { IconFileText, IconFilter, IconCode, IconDownload } from '@tabler/icons-react';

interface EditorWelcomeProps {
  onOpenLoad?: () => void;
}

export function EditorWelcome({ onOpenLoad }: EditorWelcomeProps) {
  return (
    <Stack align="center" justify="center" gap="xl" style={{ minHeight: '60vh' }}>
      <Stack align="center" gap="md">
        <IconFileText size={64} color="var(--mantine-color-blue-6)" stroke={1.5} />
        <Title order={2}>Welcome to Stream Editor</Title>
        <Text c="dimmed" ta="center" maw={600}>
          Analyze and edit WebSocket message logs. Filter, search with JQ queries, select messages,
          and export refined datasets.
        </Text>
        {onOpenLoad && (
          <Button
            size="lg"
            leftSection={<IconFileText size={20} />}
            onClick={onOpenLoad}
          >
            Set Stream
          </Button>
        )}
      </Stack>

      <SimpleGrid cols={{ base: 1, sm: 3 }} spacing="lg" mt="xl" w="100%" maw={900}>
        <Paper p="md" withBorder>
          <Stack gap="sm" align="center">
            <IconFilter size={32} color="var(--mantine-color-blue-6)" />
            <Title order={4}>Filter Messages</Title>
            <Text size="sm" c="dimmed" ta="center">
              Filter by proxy, interface, direction, and message type
            </Text>
          </Stack>
        </Paper>

        <Paper p="md" withBorder>
          <Stack gap="sm" align="center">
            <IconCode size={32} color="var(--mantine-color-blue-6)" />
            <Title order={4}>JQ Queries</Title>
            <Text size="sm" c="dimmed" ta="center">
              Use powerful JQ expressions to search and transform
            </Text>
          </Stack>
        </Paper>

        <Paper p="md" withBorder>
          <Stack gap="sm" align="center">
            <IconDownload size={32} color="var(--mantine-color-blue-6)" />
            <Title order={4}>Export</Title>
            <Text size="sm" c="dimmed" ta="center">
              Export all, selected, or non-deleted messages
            </Text>
          </Stack>
        </Paper>
      </SimpleGrid>
    </Stack>
  );
}
