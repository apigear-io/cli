import { Stack, Title, Text, Paper } from '@mantine/core';

export function Monitor() {
  return (
    <Stack gap="lg">
      <Title order={2}>Monitor</Title>
      <Paper shadow="sm" p="xl" withBorder>
        <Stack gap="md" align="center" style={{ minHeight: 200 }} justify="center">
          <Title order={3} c="dimmed">Coming Soon</Title>
          <Text c="dimmed" ta="center">
            Real-time monitoring dashboard for your API traffic and performance metrics.
          </Text>
        </Stack>
      </Paper>
    </Stack>
  );
}
