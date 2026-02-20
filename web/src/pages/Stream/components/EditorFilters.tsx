import { Group, Select, Switch, Text } from '@mantine/core';
import { IconStar } from '@tabler/icons-react';
import { useEditorContext } from './EditorContext';
import { useEditorMessages } from '@/api/queries';

export function EditorFilters() {
  const {
    sessionStats,
    currentFilters,
    setCurrentFilters,
    hideDeleted,
    setHideDeleted,
    showMarkedOnly,
    setShowMarkedOnly,
    markedIndices,
  } = useEditorContext();

  const { data: messagesData } = useEditorMessages(
    sessionStats?.sessionId || null,
    0,
    1, // Just need count, not actual messages
    currentFilters
  );

  if (!sessionStats) return null;

  const filteredCount = messagesData?.total || 0;
  const totalCount = sessionStats.totalCount;

  const proxyOptions = ['All Proxies', ...sessionStats.proxies].map((p) => ({ value: p, label: p }));
  const interfaceOptions = ['All Interfaces', ...sessionStats.interfaces].map((i) => ({
    value: i,
    label: i,
  }));
  const directionOptions = [
    { value: '', label: 'All Directions' },
    { value: 'SEND', label: 'SEND' },
    { value: 'RECV', label: 'RECV' },
  ];
  const typeOptions = [
    { value: '', label: 'All Types' },
    { value: 'LINK', label: 'LINK' },
    { value: 'INIT', label: 'INIT' },
    { value: 'INVOKE', label: 'INVOKE' },
    { value: 'INVOKE_REPLY', label: 'INVOKE_REPLY' },
    { value: 'SIGNAL', label: 'SIGNAL' },
    { value: 'PROPERTY_CHANGE', label: 'PROPERTY_CHANGE' },
  ];

  return (
    <Group p="xs" gap="md" style={{ borderBottom: '1px solid #dee2e6' }} wrap="wrap" justify="space-between">
      <Group gap="md">
        <Text size="sm" fw={500}>
          Filters:
        </Text>

        <Select
          size="xs"
          data={proxyOptions}
          value={currentFilters.proxy || 'All Proxies'}
          onChange={(value) =>
            setCurrentFilters({ ...currentFilters, proxy: value === 'All Proxies' ? undefined : value || undefined })
          }
          w={150}
        />

        <Select
          size="xs"
          data={interfaceOptions}
          value={currentFilters.interface || 'All Interfaces'}
          onChange={(value) =>
            setCurrentFilters({
              ...currentFilters,
              interface: value === 'All Interfaces' ? undefined : value || undefined,
            })
          }
          w={150}
        />

        <Select
          size="xs"
          data={directionOptions}
          value={currentFilters.direction || ''}
          onChange={(value) => setCurrentFilters({ ...currentFilters, direction: value || undefined })}
          w={130}
        />

        <Select
          size="xs"
          data={typeOptions}
          value={currentFilters.type || ''}
          onChange={(value) => setCurrentFilters({ ...currentFilters, type: value || undefined })}
          w={150}
        />

        <Switch
          label="Hide deleted"
          checked={hideDeleted}
          onChange={(e) => setHideDeleted(e.currentTarget.checked)}
          size="sm"
        />

        <Group gap={4}>
          <Switch
            checked={showMarkedOnly}
            onChange={(e) => setShowMarkedOnly(e.currentTarget.checked)}
            size="sm"
          />
          <IconStar size={16} color="var(--mantine-color-yellow-6)" fill="var(--mantine-color-yellow-6)" />
          <Text size="sm">Marked only</Text>
        </Group>
      </Group>

      <Text size="sm" c="dimmed">
        {markedIndices.size} marked | {filteredCount} / {totalCount} messages
      </Text>
    </Group>
  );
}
