import { Stack, Title, Text, Paper } from '@mantine/core';

export function Templates() {
  return (
    <Stack gap="lg">
      <Title order={2}>Templates</Title>
      <Paper shadow="sm" p="xl" withBorder>
        <Stack gap="md" align="center" style={{ minHeight: 200 }} justify="center">
          <Title order={3} c="dimmed">Coming Soon</Title>
          <Text c="dimmed" ta="center">
            Browse and install code generation templates for different languages and frameworks.
          </Text>
        </Stack>
      </Paper>
    </Stack>
  );
}
