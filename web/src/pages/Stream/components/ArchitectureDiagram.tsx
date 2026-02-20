import { Paper, Stack, Group, Box, Text, Badge } from '@mantine/core';
import { IconArrowRight } from '@tabler/icons-react';

interface ComponentBoxProps {
  label: string;
  items: string[];
  color: string;
}

function ComponentBox({ label, items, color }: ComponentBoxProps) {
  return (
    <Box>
      <Text size="xs" c="dimmed" mb={4} ta="center" fw={500}>
        {label}
      </Text>
      <Stack gap={4}>
        {items.map((item) => (
          <Badge key={item} variant="light" color={color} size="md" style={{ width: '100%' }}>
            {item}
          </Badge>
        ))}
      </Stack>
    </Box>
  );
}

export function ArchitectureDiagram() {
  return (
    <Paper p="xl" withBorder>
      <Stack gap="lg">
        <Text size="lg" fw={600}>
          Architecture
        </Text>

        <Group justify="center" align="center" gap="xl" wrap="nowrap">
          {/* INBOUND */}
          <ComponentBox
            label="INBOUND"
            items={['External Client', 'Client', 'Client Script']}
            color="blue"
          />

          <IconArrowRight size={32} color="var(--mantine-color-gray-5)" />

          {/* PROXY */}
          <Box>
            <Text size="xs" c="dimmed" mb={4} ta="center" fw={500}>
              PROXY
            </Text>
            <Paper
              withBorder
              p="md"
              style={{
                backgroundColor: 'var(--mantine-color-violet-0)',
                borderColor: 'var(--mantine-color-violet-3)',
              }}
            >
              <Stack gap="xs" align="center">
                <Badge variant="filled" color="violet" size="lg">
                  WSProxy
                </Badge>
                <Group gap="xs">
                  <Badge variant="light" color="yellow" size="sm">
                    Stream
                  </Badge>
                  <Badge variant="light" color="orange" size="sm">
                    Traces
                  </Badge>
                </Group>
              </Stack>
            </Paper>
          </Box>

          <IconArrowRight size={32} color="var(--mantine-color-gray-5)" />

          {/* OUTBOUND */}
          <ComponentBox
            label="OUTBOUND"
            items={['External Service', 'Backend Script']}
            color="green"
          />
        </Group>
      </Stack>
    </Paper>
  );
}
