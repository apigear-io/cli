import { useState } from 'react';
import { Group, TextInput, Button, Collapse, Text, Badge, Paper } from '@mantine/core';
import { IconCode, IconPlayerPlay, IconChevronDown, IconChevronUp } from '@tabler/icons-react';
import { useEditorContext } from './EditorContext';
import { useEditorJQ } from '@/api/queries';
import { notifications } from '@mantine/notifications';

export function EditorJQPanel() {
  const { sessionStats, setSelectedIndices } = useEditorContext();
  const [query, setQuery] = useState('');
  const [showExamples, setShowExamples] = useState(false);
  const jqMutation = useEditorJQ();

  if (!sessionStats) return null;

  const handleRun = async () => {
    if (!query.trim()) return;

    try {
      const result = await jqMutation.mutateAsync({
        sessionId: sessionStats.sessionId,
        query,
        limit: 100,
      });

      notifications.show({
        title: 'JQ Query Complete',
        message: `Found ${result.totalMatches} matches`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'JQ Query Failed',
        message: error instanceof Error ? error.message : 'Query failed',
        color: 'red',
      });
    }
  };

  const handleSelectMatches = async () => {
    if (!query.trim()) return;

    try {
      const result = await jqMutation.mutateAsync({
        sessionId: sessionStats.sessionId,
        query,
        limit: 1000,
      });

      const indices = new Set(result.matches.map((m) => m.index));
      setSelectedIndices(indices);

      notifications.show({
        title: 'Matches Selected',
        message: `Selected ${indices.size} messages`,
        color: 'blue',
      });
    } catch (error) {
      notifications.show({
        title: 'Selection Failed',
        message: error instanceof Error ? error.message : 'Failed to select matches',
        color: 'red',
      });
    }
  };

  const examples = [
    { label: 'Messages with args', query: 'select(.msg[1] | has("args"))' },
    { label: 'INVOKE messages', query: 'select(.msg[1].msgType == 30)' },
    { label: 'Specific interface', query: 'select(.msg[1].symbol == "demo.Counter")' },
  ];

  return (
    <Paper p="xs" style={{ borderBottom: '1px solid #dee2e6' }}>
      <Group gap="md" align="flex-start" wrap="nowrap">
        <IconCode size={20} style={{ marginTop: 6 }} />
        <div style={{ flex: 1 }}>
          <TextInput
            placeholder='JQ query (e.g., select(.msg[1].msgType == 30))'
            value={query}
            onChange={(e) => setQuery(e.currentTarget.value)}
            onKeyDown={(e) => {
              if (e.key === 'Enter' && !e.shiftKey) {
                e.preventDefault();
                handleRun();
              }
            }}
          />

          <Collapse in={showExamples}>
            <Group gap="xs" mt="xs">
              <Text size="xs" c="dimmed">
                Examples:
              </Text>
              {examples.map((ex) => (
                <Badge
                  key={ex.label}
                  variant="light"
                  style={{ cursor: 'pointer' }}
                  onClick={() => setQuery(ex.query)}
                >
                  {ex.label}
                </Badge>
              ))}
            </Group>
          </Collapse>
        </div>

        <Button
          size="xs"
          variant="subtle"
          onClick={() => setShowExamples(!showExamples)}
          rightSection={showExamples ? <IconChevronUp size={14} /> : <IconChevronDown size={14} />}
        >
          Examples
        </Button>

        <Button
          size="xs"
          leftSection={<IconPlayerPlay size={14} />}
          onClick={handleRun}
          loading={jqMutation.isPending}
        >
          Run
        </Button>

        <Button
          size="xs"
          variant="light"
          onClick={handleSelectMatches}
          loading={jqMutation.isPending}
        >
          Select Matches
        </Button>
      </Group>
    </Paper>
  );
}
