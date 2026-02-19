import { Suspense } from 'react';
import { Card, Grid, Text, Title, Stack, Group, Badge, Button } from '@mantine/core';
import { IconActivity, IconServer, IconUsers, IconMessages } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';
import { useStreamDashboard } from '@/api/queries';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';

function DashboardContent() {
  const { data: stats } = useStreamDashboard();
  const navigate = useNavigate();

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="center">
        <Title order={2}>Stream Dashboard</Title>
        <Badge size="lg" variant="light" color="blue">
          <IconActivity size={14} style={{ marginRight: 4 }} />
          Live
        </Badge>
      </Group>

      <Grid>
        {/* Proxies Card */}
        <Grid.Col span={{ base: 12, sm: 6, lg: 3 }}>
          <Card
            shadow="sm"
            padding="lg"
            radius="md"
            withBorder
            style={{ cursor: 'pointer' }}
            onClick={() => navigate('/stream/proxies')}
          >
            <Stack gap="xs">
              <Group justify="space-between">
                <Text size="sm" c="dimmed" fw={500}>
                  Proxies
                </Text>
                <IconServer size={20} color="var(--mantine-color-blue-6)" />
              </Group>
              <Text size="xl" fw={700}>
                {stats.proxies.total}
              </Text>
              <Group gap="xs">
                <Badge size="sm" color="green" variant="light">
                  {stats.proxies.running} running
                </Badge>
                <Badge size="sm" color="gray" variant="light">
                  {stats.proxies.stopped} stopped
                </Badge>
              </Group>
            </Stack>
          </Card>
        </Grid.Col>

        {/* Clients Card */}
        <Grid.Col span={{ base: 12, sm: 6, lg: 3 }}>
          <Card
            shadow="sm"
            padding="lg"
            radius="md"
            withBorder
            style={{ cursor: 'pointer' }}
            onClick={() => navigate('/stream/clients')}
          >
            <Stack gap="xs">
              <Group justify="space-between">
                <Text size="sm" c="dimmed" fw={500}>
                  Clients
                </Text>
                <IconUsers size={20} color="var(--mantine-color-cyan-6)" />
              </Group>
              <Text size="xl" fw={700}>
                {stats.clients.total}
              </Text>
              <Group gap="xs">
                <Badge size="sm" color="green" variant="light">
                  {stats.clients.connected} connected
                </Badge>
                <Badge size="sm" color="gray" variant="light">
                  {stats.clients.disconnected} offline
                </Badge>
              </Group>
            </Stack>
          </Card>
        </Grid.Col>

        {/* Messages Card */}
        <Grid.Col span={{ base: 12, sm: 6, lg: 3 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="xs">
              <Group justify="space-between">
                <Text size="sm" c="dimmed" fw={500}>
                  Messages
                </Text>
                <IconMessages size={20} color="var(--mantine-color-violet-6)" />
              </Group>
              <Text size="xl" fw={700}>
                {stats.messages.total.toLocaleString()}
              </Text>
              <Text size="xs" c="dimmed">
                {stats.messages.rate.toFixed(2)} msg/s
              </Text>
            </Stack>
          </Card>
        </Grid.Col>

        {/* Quick Actions Card */}
        <Grid.Col span={{ base: 12, sm: 6, lg: 3 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed" fw={500}>
                Quick Actions
              </Text>
              <Stack gap="xs">
                <Button
                  variant="light"
                  size="xs"
                  fullWidth
                  onClick={() => navigate('/stream/proxies')}
                >
                  Manage Proxies
                </Button>
                <Button
                  variant="light"
                  size="xs"
                  fullWidth
                  onClick={() => navigate('/stream/clients')}
                >
                  Manage Clients
                </Button>
              </Stack>
            </Stack>
          </Card>
        </Grid.Col>
      </Grid>

      {/* Info Cards */}
      <Grid>
        <Grid.Col span={{ base: 12, md: 6 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="md">
              <Title order={4}>About WebSocket Streaming</Title>
              <Text size="sm" c="dimmed">
                The Stream module provides WebSocket proxy and client management capabilities with
                real-time message monitoring. Use proxies to forward WebSocket connections and
                clients to connect to ObjectLink backends.
              </Text>
              <Group gap="xs">
                <Badge variant="light">Proxy Modes</Badge>
                <Badge variant="light">ObjectLink Protocol</Badge>
                <Badge variant="light">Message Tracing</Badge>
              </Group>
            </Stack>
          </Card>
        </Grid.Col>

        <Grid.Col span={{ base: 12, md: 6 }}>
          <Card shadow="sm" padding="lg" radius="md" withBorder>
            <Stack gap="md">
              <Title order={4}>Getting Started</Title>
              <Stack gap="xs">
                <Text size="sm">
                  1. Create a proxy to forward WebSocket connections
                </Text>
                <Text size="sm">
                  2. Start the proxy to begin accepting connections
                </Text>
                <Text size="sm">
                  3. Create clients to connect to ObjectLink backends
                </Text>
                <Text size="sm">
                  4. Monitor messages and connection status in real-time
                </Text>
              </Stack>
            </Stack>
          </Card>
        </Grid.Col>
      </Grid>
    </Stack>
  );
}

export function StreamDashboard() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading stream dashboard..." />}>
        <DashboardContent />
      </Suspense>
    </ErrorBoundary>
  );
}
