import { Paper, Stack, Text, Group, ActionIcon, Badge, ScrollArea } from '@mantine/core';
import { IconTrash, IconFileCode, IconPlayerPlay } from '@tabler/icons-react';
import type { ScriptInfo } from '@/api/types';

interface ScriptListProps {
  scripts: string[];
  currentScript: string | null;
  onSelect: (name: string) => void;
  onDelete: (name: string) => void;
  onRun: (name: string) => void;
}

export function ScriptList({ scripts, currentScript, onSelect, onDelete, onRun }: ScriptListProps) {
  return (
    <Paper withBorder p="md">
      <Stack gap="md">
        <Text size="sm" fw={600}>
          Scripts
        </Text>

        <ScrollArea h={300}>
          <Stack gap="xs">
            {scripts.length === 0 && (
              <Text size="sm" c="dimmed" fs="italic">
                No saved scripts yet
              </Text>
            )}
            {scripts.map((name) => (
              <Group
                key={name}
                justify="space-between"
                p="xs"
                style={{
                  borderRadius: 4,
                  backgroundColor: name === currentScript ? 'var(--mantine-color-blue-light)' : 'transparent',
                  cursor: 'pointer',
                }}
                onClick={() => onSelect(name)}
              >
                <Group gap="xs">
                  <IconFileCode size={16} />
                  <Text size="sm">{name}</Text>
                </Group>
                <Group gap={4}>
                  <ActionIcon
                    size="sm"
                    variant="subtle"
                    color="blue"
                    onClick={(e) => {
                      e.stopPropagation();
                      onRun(name);
                    }}
                    title="Run script"
                  >
                    <IconPlayerPlay size={14} />
                  </ActionIcon>
                  <ActionIcon
                    size="sm"
                    variant="subtle"
                    color="red"
                    onClick={(e) => {
                      e.stopPropagation();
                      onDelete(name);
                    }}
                    title="Delete script"
                  >
                    <IconTrash size={14} />
                  </ActionIcon>
                </Group>
              </Group>
            ))}
          </Stack>
        </ScrollArea>
      </Stack>
    </Paper>
  );
}

interface RunningScriptsProps {
  scripts: ScriptInfo[];
  onStop: (id: string) => void;
}

export function RunningScripts({ scripts, onStop }: RunningScriptsProps) {
  return (
    <Paper withBorder p="md" mt="md">
      <Stack gap="md">
        <Group justify="space-between">
          <Text size="sm" fw={600}>
            Running Scripts
          </Text>
          {scripts.length > 0 && (
            <Badge size="sm" color="green">
              {scripts.length}
            </Badge>
          )}
        </Group>

        <ScrollArea h={200}>
          <Stack gap="xs">
            {scripts.length === 0 && (
              <Text size="sm" c="dimmed" fs="italic">
                No running scripts
              </Text>
            )}
            {scripts.map((script) => (
              <Group key={script.id} justify="space-between" p="xs">
                <Stack gap={0}>
                  <Text size="sm">{script.name}</Text>
                  <Text size="xs" c="dimmed">
                    {script.type}
                  </Text>
                </Stack>
                <ActionIcon
                  size="sm"
                  variant="subtle"
                  color="red"
                  onClick={() => onStop(script.id)}
                  title="Stop script"
                >
                  <IconTrash size={14} />
                </ActionIcon>
              </Group>
            ))}
          </Stack>
        </ScrollArea>
      </Stack>
    </Paper>
  );
}
