import { useState } from 'react';
import { Stack, Group, Badge, Text, Select, ScrollArea, Paper, Code } from '@mantine/core';
import { useTraceFile } from '@/api/queries';

interface TraceViewerProps {
  filename: string;
}

export function TraceViewer({ filename }: TraceViewerProps) {
  const [direction, setDirection] = useState<string>('');
  const [limit, setLimit] = useState<number>(100);

  const { data: trace } = useTraceFile(filename, {
    direction: direction || undefined,
    limit,
  });

  const formatTimestamp = (ts: number): string => {
    const date = new Date(ts);
    return date.toLocaleTimeString() + '.' + (ts % 1000).toString().padStart(3, '0');
  };

  const formatMessage = (msg: unknown): string => {
    try {
      return JSON.stringify(msg, null, 2);
    } catch {
      return String(msg);
    }
  };

  return (
    <Stack gap="md">
      <Group>
        <Select
          label="Direction"
          placeholder="All"
          value={direction}
          onChange={(value) => setDirection(value || '')}
          data={[
            { value: '', label: 'All' },
            { value: 'SEND', label: 'Send' },
            { value: 'RECV', label: 'Receive' },
          ]}
          style={{ width: 150 }}
        />
        <Select
          label="Limit"
          value={limit.toString()}
          onChange={(value) => setLimit(parseInt(value || '100', 10))}
          data={[
            { value: '50', label: '50 entries' },
            { value: '100', label: '100 entries' },
            { value: '500', label: '500 entries' },
            { value: '1000', label: '1000 entries' },
          ]}
          style={{ width: 150 }}
        />
        <Text size="sm" c="dimmed" pt={25}>
          Showing {trace.count} of {trace.count} entries
        </Text>
      </Group>

      <ScrollArea h={500}>
        <Stack gap="xs">
          {trace.entries.length === 0 && (
            <Text ta="center" c="dimmed" py="xl">
              No entries found
            </Text>
          )}
          {trace.entries.map((entry, index) => (
            <Paper key={index} withBorder p="sm">
              <Stack gap="xs">
                <Group gap="md">
                  <Badge
                    color={entry.dir === 'SEND' ? 'blue' : 'green'}
                    variant="light"
                    size="sm"
                  >
                    {entry.dir}
                  </Badge>
                  <Text size="xs" c="dimmed" ff="monospace">
                    {formatTimestamp(entry.ts)}
                  </Text>
                  <Badge variant="light" size="sm">
                    {entry.proxy}
                  </Badge>
                </Group>
                <Code block>{formatMessage(entry.msg)}</Code>
              </Stack>
            </Paper>
          ))}
        </Stack>
      </ScrollArea>
    </Stack>
  );
}
