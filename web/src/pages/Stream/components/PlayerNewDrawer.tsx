import { useState } from 'react';
import {
  Drawer,
  Stack,
  Select,
  NumberInput,
  Switch,
  Group,
  Button,
  SegmentedControl,
  Text,
} from '@mantine/core';
import { notifications } from '@mantine/notifications';
import { useProxies, useTraceFiles, useCreatePlayerStream } from '@/api/queries';
import type { CreatePlayerStreamRequest, TraceFileInfo } from '@/api/types';

interface PlayerNewDrawerProps {
  opened: boolean;
  onClose: () => void;
}

export function PlayerNewDrawer({ opened, onClose }: PlayerNewDrawerProps) {
  const { data: proxies } = useProxies();
  const { data: filesData } = useTraceFiles();
  const createStream = useCreatePlayerStream();

  const [formData, setFormData] = useState({
    proxyName: '',
    filename: '',
    fileSource: 'directory', // 'directory' or 'upload'
    speed: 1.0,
    initialDelay: 0,
    loop: false,
    direction: '',
  });

  const proxyOptions = proxies.map((p) => ({ value: p.name, label: p.name }));
  const fileOptions = filesData.map((f: TraceFileInfo) => ({ value: f.name, label: f.name }));

  const speedOptions = [
    { value: '0.5', label: '0.5x' },
    { value: '1', label: '1x' },
    { value: '2', label: '2x' },
    { value: '5', label: '5x' },
  ];

  const directionOptions = [
    { value: '', label: 'Both' },
    { value: 'SEND', label: 'SEND' },
    { value: 'RECV', label: 'RECV' },
  ];

  const handleCreate = async () => {
    if (!formData.proxyName) {
      notifications.show({
        title: 'Error',
        message: 'Please select a target proxy',
        color: 'red',
      });
      return;
    }

    if (!formData.filename) {
      notifications.show({
        title: 'Error',
        message: 'Please select a trace file',
        color: 'red',
      });
      return;
    }

    try {
      const request: CreatePlayerStreamRequest = {
        proxyName: formData.proxyName,
        filename: formData.filename,
        speed: formData.speed,
        initialDelay: formData.initialDelay,
        loop: formData.loop,
        direction: formData.direction,
      };

      await createStream.mutateAsync(request);
      notifications.show({
        title: 'Success',
        message: 'Stream created successfully',
        color: 'green',
      });
      onClose();

      // Reset form
      setFormData({
        proxyName: '',
        filename: '',
        fileSource: 'directory',
        speed: 1.0,
        initialDelay: 0,
        loop: false,
        direction: '',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to create stream',
        color: 'red',
      });
    }
  };

  return (
    <Drawer
      opened={opened}
      onClose={onClose}
      title="New Stream"
      position="right"
      size="md"
      padding="lg"
    >
      <Stack gap="md">
        {/* Target Proxy */}
        <Select
          label="Target Proxy"
          placeholder="Select a proxy..."
          description="Messages will be sent to this proxy's backend"
          data={proxyOptions}
          value={formData.proxyName}
          onChange={(value) => setFormData({ ...formData, proxyName: value || '' })}
          required
        />

        {/* File Source */}
        <div>
          <Text size="sm" fw={500} mb="xs">
            File Source
          </Text>
          <SegmentedControl
            fullWidth
            value={formData.fileSource}
            onChange={(value) => setFormData({ ...formData, fileSource: value })}
            data={[
              { value: 'directory', label: 'Traces Directory' },
              { value: 'upload', label: 'Upload File' },
            ]}
          />
        </div>

        {/* Trace File */}
        {formData.fileSource === 'directory' && (
          <Select
            label="Trace File"
            placeholder="Select a trace file..."
            description="JSONL file from the traces directory"
            data={fileOptions}
            value={formData.filename}
            onChange={(value) => setFormData({ ...formData, filename: value || '' })}
            required
          />
        )}

        {/* Speed */}
        <div>
          <Text size="sm" fw={500} mb="xs">
            Speed
          </Text>
          <SegmentedControl
            fullWidth
            value={formData.speed.toString()}
            onChange={(value) => setFormData({ ...formData, speed: parseFloat(value) })}
            data={speedOptions}
          />
        </div>

        {/* Initial Delay */}
        <NumberInput
          label="Initial Delay (ms)"
          description="Wait before starting playback"
          value={formData.initialDelay}
          onChange={(value) => setFormData({ ...formData, initialDelay: Number(value) || 0 })}
          min={0}
          max={10000}
          step={100}
        />

        {/* Loop Playback */}
        <Switch
          label="Loop playback"
          checked={formData.loop}
          onChange={(e) => setFormData({ ...formData, loop: e.currentTarget.checked })}
        />

        {/* Direction Filter */}
        <div>
          <Text size="sm" fw={500} mb="xs">
            Direction Filter
          </Text>
          <Text size="xs" c="dimmed" mb="xs">
            Filter messages by direction
          </Text>
          <SegmentedControl
            fullWidth
            value={formData.direction}
            onChange={(value) => setFormData({ ...formData, direction: value })}
            data={directionOptions}
          />
        </div>

        {/* Actions */}
        <Group justify="flex-end" mt="md">
          <Button variant="light" onClick={onClose}>
            Cancel
          </Button>
          <Button onClick={handleCreate} loading={createStream.isPending}>
            Load Stream
          </Button>
        </Group>
      </Stack>
    </Drawer>
  );
}
