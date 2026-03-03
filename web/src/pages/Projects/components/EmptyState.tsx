import { Stack, Title, Text, Button, Paper } from '@mantine/core';
import { IconFolderPlus } from '@tabler/icons-react';

interface EmptyStateProps {
  onCreate: () => void;
}

export function EmptyState({ onCreate }: EmptyStateProps) {
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
        <Button onClick={onCreate} leftSection={<IconFolderPlus size={16} />}>
          Create Your First Project
        </Button>
      </Stack>
    </Paper>
  );
}
