import { Group, Select, Switch, Badge, Text } from '@mantine/core';
import { useEditorContext } from './EditorContext';

export function EditorFilters() {
  const {
    sessionStats,
    currentFilters,
    setCurrentFilters,
    hideDeleted,
    setHideDeleted,
    showMarkedOnly,
    setShowMarkedOnly,
    deletedIndices,
    markedIndices,
  } = useEditorContext();

  if (!sessionStats) return null;

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
    <Group p="xs" gap="md" style={{ borderBottom: '1px solid #dee2e6' }} wrap="wrap">
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

      <Switch
        label="Marked only"
        checked={showMarkedOnly}
        onChange={(e) => setShowMarkedOnly(e.currentTarget.checked)}
        size="sm"
      />

      <Badge>{markedIndices.size} marked</Badge>
      <Badge color="orange">{deletedIndices.size} deleted</Badge>
    </Group>
  );
}
