import { Card, Text, Group, Stack, ActionIcon, Tooltip, Progress, Badge } from '@mantine/core';
import {
  IconPlayerPlay,
  IconPlayerPause,
  IconPlayerStop,
  IconTrash,
} from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import {
  usePlayPlayerStream,
  usePausePlayerStream,
  useStopPlayerStream,
  useDeletePlayerStream,
} from '@/api/queries';
import type { PlayerStream } from '@/api/types';

interface PlayerStreamCardProps {
  stream: PlayerStream;
}

export function PlayerStreamCard({ stream }: PlayerStreamCardProps) {
  const playStream = usePlayPlayerStream();
  const pauseStream = usePausePlayerStream();
  const stopStream = useStopPlayerStream();
  const deleteStream = useDeletePlayerStream();

  const handlePlay = async () => {
    try {
      await playStream.mutateAsync(stream.id);
      notifications.show({
        title: 'Success',
        message: 'Playback started',
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to start playback',
        color: 'red',
      });
    }
  };

  const handlePause = async () => {
    try {
      await pauseStream.mutateAsync(stream.id);
      notifications.show({
        title: 'Success',
        message: 'Playback paused',
        color: 'blue',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to pause playback',
        color: 'red',
      });
    }
  };

  const handleStop = async () => {
    try {
      await stopStream.mutateAsync(stream.id);
      notifications.show({
        title: 'Success',
        message: 'Playback stopped',
        color: 'blue',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to stop playback',
        color: 'red',
      });
    }
  };

  const handleDelete = async () => {
    if (!confirm('Are you sure you want to delete this stream?')) {
      return;
    }

    try {
      await deleteStream.mutateAsync(stream.id);
      notifications.show({
        title: 'Success',
        message: 'Stream deleted',
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to delete stream',
        color: 'red',
      });
    }
  };

  const getStateColor = () => {
    switch (stream.state) {
      case 'playing':
        return 'green';
      case 'paused':
        return 'yellow';
      default:
        return 'gray';
    }
  };

  const getStateLabel = () => {
    switch (stream.state) {
      case 'playing':
        return 'Playing';
      case 'paused':
        return 'Paused';
      default:
        return 'Stopped';
    }
  };

  return (
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Stack gap="md">
        {/* Header */}
        <Group justify="space-between" wrap="nowrap">
          <Stack gap={4} style={{ flex: 1, minWidth: 0 }}>
            <Group gap="xs">
              <Text fw={600} size="md" lineClamp={1}>
                {stream.proxyName}
              </Text>
              <Badge size="sm" color={getStateColor()}>
                {getStateLabel()}
              </Badge>
            </Group>
            <Text size="xs" c="dimmed" lineClamp={1}>
              {stream.filename}
            </Text>
          </Stack>

          {/* Delete button */}
          <Tooltip label="Delete stream">
            <ActionIcon
              variant="light"
              color="red"
              size="sm"
              onClick={handleDelete}
              disabled={stream.state === 'playing'}
            >
              <IconTrash size={14} />
            </ActionIcon>
          </Tooltip>
        </Group>

        {/* Progress Bar */}
        <div>
          <Progress value={stream.progress * 100} size="sm" />
          <Group justify="space-between" mt={4}>
            <Text size="xs" c="dimmed">
              {stream.position} / {stream.totalEntries}
            </Text>
            <Text size="xs" c="dimmed">
              {(stream.progress * 100).toFixed(1)}%
            </Text>
          </Group>
        </div>

        {/* Settings */}
        <Group gap="lg">
          <Group gap={4}>
            <Text size="xs" c="dimmed">
              Speed:
            </Text>
            <Text size="xs" fw={500}>
              {stream.speed}x
            </Text>
          </Group>
          {stream.loop && (
            <Badge size="xs" variant="light" color="violet">
              Loop
            </Badge>
          )}
          {stream.direction && (
            <Badge size="xs" variant="light" color="blue">
              {stream.direction}
            </Badge>
          )}
        </Group>

        {/* Controls */}
        <Group gap="xs" justify="center">
          {stream.state === 'stopped' && (
            <Tooltip label="Play">
              <ActionIcon
                variant="filled"
                color="green"
                size="lg"
                onClick={handlePlay}
                loading={playStream.isPending}
              >
                <IconPlayerPlay size={20} />
              </ActionIcon>
            </Tooltip>
          )}
          {stream.state === 'playing' && (
            <Tooltip label="Pause">
              <ActionIcon
                variant="filled"
                color="yellow"
                size="lg"
                onClick={handlePause}
                loading={pauseStream.isPending}
              >
                <IconPlayerPause size={20} />
              </ActionIcon>
            </Tooltip>
          )}
          {stream.state === 'paused' && (
            <Tooltip label="Resume">
              <ActionIcon
                variant="filled"
                color="green"
                size="lg"
                onClick={handlePlay}
                loading={playStream.isPending}
              >
                <IconPlayerPlay size={20} />
              </ActionIcon>
            </Tooltip>
          )}
          {stream.state !== 'stopped' && (
            <Tooltip label="Stop">
              <ActionIcon
                variant="filled"
                color="red"
                size="lg"
                onClick={handleStop}
                loading={stopStream.isPending}
              >
                <IconPlayerStop size={20} />
              </ActionIcon>
            </Tooltip>
          )}
        </Group>
      </Stack>
    </Card>
  );
}
