import { Stack, Title, Text, Button, Paper, Group } from '@mantine/core';
import { IconFolderPlus, IconFolderOpen } from '@tabler/icons-react';

interface EmptyStateProps {
  onCreate: () => void;
  onOpen: () => void;
}

export function EmptyState({ onCreate, onOpen }: EmptyStateProps) {
  return (
    <Paper shadow="sm" p="xl" withBorder>
      <Stack gap="md" align="center" style={{ minHeight: 300 }} justify="center">
        <IconFolderPlus size={48} stroke={1.5} style={{ opacity: 0.5 }} />
        <Title order={3} c="dimmed">
          No Projects Yet
        </Title>
        <Text c="dimmed" ta="center" maw={400}>
          Create your first ApiGear project to get started with API development and code
          generation.
        </Text>
        <Group>
          <Button onClick={onCreate} leftSection={<IconFolderPlus size={16} />}>
            Create New Project
          </Button>
          <Button
            variant="light"
            onClick={onOpen}
            leftSection={<IconFolderOpen size={16} />}
          >
            Open Existing Project
          </Button>
        </Group>
      </Stack>
    </Paper>
  );
}
