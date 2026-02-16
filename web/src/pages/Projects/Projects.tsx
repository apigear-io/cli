import { Stack, Title, Text, Paper } from '@mantine/core';

export function Projects() {
  return (
    <Stack gap="lg">
      <Title order={2}>Projects</Title>
      <Paper shadow="sm" p="xl" withBorder>
        <Stack gap="md" align="center" style={{ minHeight: 200 }} justify="center">
          <Title order={3} c="dimmed">Coming Soon</Title>
          <Text c="dimmed" ta="center">
            Manage your ApiGear projects and configurations.
          </Text>
        </Stack>
      </Paper>
    </Stack>
  );
}
