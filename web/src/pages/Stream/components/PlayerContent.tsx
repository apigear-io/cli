import { useState } from 'react';
import { Stack, Title, Group, Button, Card, Text, SimpleGrid } from '@mantine/core';
import { IconPlus, IconPlayerPlay } from '@tabler/icons-react';
import { usePlayerStreams } from '@/api/queries';
import { PlayerNewDrawer } from './PlayerNewDrawer';
import { PlayerStreamCard } from './PlayerStreamCard';

export function PlayerContent() {
  const { data: streams } = usePlayerStreams();
  const [newDrawerOpen, setNewDrawerOpen] = useState(false);

  return (
    <>
      <Stack gap="md">
        {/* Header */}
        <Group justify="space-between">
          <Title order={2}>Stream Player</Title>
          <Button leftSection={<IconPlus size={16} />} onClick={() => setNewDrawerOpen(true)}>
            New Stream
          </Button>
        </Group>

        {/* Streams or Welcome */}
        {streams.length === 0 ? (
          <Card shadow="sm" padding="xl" radius="md" withBorder>
            <Stack align="center" gap="md">
              <IconPlayerPlay size={48} color="var(--mantine-color-gray-5)" />
              <Text size="lg" fw={500} c="dimmed">
                No Active Streams
              </Text>
              <Text size="sm" c="dimmed" ta="center">
                Load a trace file to start streaming messages to a proxy
              </Text>
              <Button leftSection={<IconPlus size={16} />} onClick={() => setNewDrawerOpen(true)}>
                New Stream
              </Button>
            </Stack>
          </Card>
        ) : (
          <SimpleGrid cols={{ base: 1, md: 2, lg: 3 }} spacing="md">
            {streams.map((stream) => (
              <PlayerStreamCard key={stream.id} stream={stream} />
            ))}
          </SimpleGrid>
        )}
      </Stack>

      {/* New Stream Drawer */}
      <PlayerNewDrawer opened={newDrawerOpen} onClose={() => setNewDrawerOpen(false)} />
    </>
  );
}
