import { Stack, Title, Text, Paper } from '@mantine/core';

export function CodeGen() {
  return (
    <Stack gap="lg">
      <Title order={2}>Code Generation</Title>
      <Paper shadow="sm" p="xl" withBorder>
        <Stack gap="md" align="center" style={{ minHeight: 200 }} justify="center">
          <Title order={3} c="dimmed">Coming Soon</Title>
          <Text c="dimmed" ta="center">
            Generate SDKs from your API specifications with drag-and-drop support.
          </Text>
        </Stack>
      </Paper>
    </Stack>
  );
}
