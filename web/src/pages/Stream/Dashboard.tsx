import { Suspense } from 'react';
import { Card, Text, Title, Stack, Group, Badge, SimpleGrid } from '@mantine/core';
import { IconActivity } from '@tabler/icons-react';
import { useStreamDashboard } from '@/api/queries';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { QuickActions } from './components/QuickActions';
import { ArchitectureDiagram } from './components/ArchitectureDiagram';
import { ProxyStatsTable } from './components/ProxyStatsTable';

function DashboardContent() {
  const { data: stats } = useStreamDashboard();

  return (
    <Stack gap="xl">
      {/* Header */}
      <Group justify="space-between" align="center">
        <div>
          <Title order={1} size="h2">
            Analytics
          </Title>
        </div>
        <Badge size="lg" variant="dot" color="green">
          <Group gap={4}>
            <IconActivity size={14} />
            <Text size="sm">00:00:35</Text>
          </Group>
        </Badge>
      </Group>

      {/* Quick Actions */}
      <div>
        <Text size="md" fw={600} mb="md">
          Quick Actions
        </Text>
        <QuickActions />
      </div>

      {/* Architecture */}
      <ArchitectureDiagram />

      {/* Statistics Cards */}
      <SimpleGrid cols={{ base: 1, sm: 2, md: 4 }} spacing="md">
        <Card shadow="sm" padding="lg" radius="md" withBorder>
          <Stack gap="xs">
            <Text size="sm" c="dimmed" fw={500}>
              ACTIVE CONNECTIONS
            </Text>
            <Text size="2rem" fw={700} c="blue">
              {stats.proxies.running * 2}
            </Text>
            <Text size="xs" c="dimmed">
              / {stats.proxies.total * 2} total
            </Text>
            <Text size="xs" c="dimmed">
              0 failed
            </Text>
          </Stack>
        </Card>

        <Card shadow="sm" padding="lg" radius="md" withBorder>
          <Stack gap="xs">
            <Text size="sm" c="dimmed" fw={500}>
              MESSAGES IN
            </Text>
            <Text size="2rem" fw={700} c="blue">
              {stats.messages.total}
            </Text>
            <Text size="xs" c="dimmed">
              0/s
            </Text>
            <Text size="xs" c="dimmed">
              0 B
            </Text>
          </Stack>
        </Card>

        <Card shadow="sm" padding="lg" radius="md" withBorder>
          <Stack gap="xs">
            <Text size="sm" c="dimmed" fw={500}>
              MESSAGES OUT
            </Text>
            <Text size="2rem" fw={700} c="blue">
              {stats.messages.total}
            </Text>
            <Text size="xs" c="dimmed">
              0/s
            </Text>
            <Text size="xs" c="dimmed">
              0 B
            </Text>
          </Stack>
        </Card>

        <Card shadow="sm" padding="lg" radius="md" withBorder>
          <Stack gap="xs">
            <Text size="sm" c="dimmed" fw={500}>
              TOTAL THROUGHPUT
            </Text>
            <Text size="2rem" fw={700} c="blue">
              0
            </Text>
            <Text size="xs" c="dimmed">
              0/s
            </Text>
            <Text size="xs" c="dimmed">
              0 B total
            </Text>
          </Stack>
        </Card>
      </SimpleGrid>

      {/* Proxy Statistics */}
      <div>
        <Group justify="space-between" mb="md">
          <Text size="md" fw={600}>
            Proxy Statistics
          </Text>
          <Badge>{stats.proxies.total} proxies</Badge>
        </Group>
        <ProxyStatsTable />
      </div>
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
