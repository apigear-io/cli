import { Card, Grid, Text, Title, Badge, Stack, Alert, Loader } from '@mantine/core';
import { IconInfoCircle } from '@tabler/icons-react';
import { useStatus, useHealth } from '@/api/queries';

export function Dashboard() {
  const { data: status, isLoading: statusLoading, error: statusError } = useStatus();
  const { data: health, isLoading: healthLoading, error: healthError } = useHealth();

  if (statusLoading || healthLoading) {
    return (
      <Stack align="center" justify="center" style={{ minHeight: 400 }}>
        <Loader size="lg" />
        <Text>Loading system status...</Text>
      </Stack>
    );
  }

  if (statusError || healthError) {
    return (
      <Alert
        icon={<IconInfoCircle size={16} />}
        title="Error"
        color="red"
      >
        Failed to load system status. Please ensure the ApiGear server is running.
      </Alert>
    );
  }

  return (
    <Stack gap="lg">
      <Title order={2}>System Dashboard</Title>

      <Grid>
        <Grid.Col span={{ base: 12, sm: 6, md: 4 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed" fw={500}>
                Health Status
              </Text>
              <Badge
                color={health?.status === 'ok' ? 'green' : 'red'}
                variant="light"
                size="lg"
              >
                {health?.status || 'Unknown'}
              </Badge>
            </Stack>
          </Card>
        </Grid.Col>

        <Grid.Col span={{ base: 12, sm: 6, md: 4 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed" fw={500}>
                Version
              </Text>
              <Text size="xl" fw={700}>
                {status?.version || 'N/A'}
              </Text>
            </Stack>
          </Card>
        </Grid.Col>

        <Grid.Col span={{ base: 12, sm: 6, md: 4 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed" fw={500}>
                Uptime
              </Text>
              <Text size="xl" fw={700}>
                {status?.uptime || 'N/A'}
              </Text>
            </Stack>
          </Card>
        </Grid.Col>

        <Grid.Col span={{ base: 12, sm: 6, md: 4 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed" fw={500}>
                Commit
              </Text>
              <Text size="sm" fw={500} style={{ fontFamily: 'monospace' }}>
                {status?.commit ? status.commit.substring(0, 8) : 'N/A'}
              </Text>
            </Stack>
          </Card>
        </Grid.Col>

        <Grid.Col span={{ base: 12, sm: 6, md: 4 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed" fw={500}>
                Build Date
              </Text>
              <Text size="sm" fw={500}>
                {status?.buildDate || 'N/A'}
              </Text>
            </Stack>
          </Card>
        </Grid.Col>

        <Grid.Col span={{ base: 12, sm: 6, md: 4 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed" fw={500}>
                Go Version
              </Text>
              <Text size="sm" fw={500}>
                {status?.goVersion || 'N/A'}
              </Text>
            </Stack>
          </Card>
        </Grid.Col>
      </Grid>
    </Stack>
  );
}
