import { useState, useMemo } from 'react';
import { Stack, Title, Tabs, TextInput, Button, Group, Alert } from '@mantine/core';
import { IconSearch, IconRefresh, IconAlertCircle } from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { useTemplates, useCachedTemplates, useUpdateRegistry } from '@/api/queries';
import { RegistryTemplateList } from './components/RegistryTemplateList';
import { CachedTemplateList } from './components/CachedTemplateList';

export function Templates() {
  const [searchQuery, setSearchQuery] = useState('');
  const [activeTab, setActiveTab] = useState<string | null>('registry');

  const { data: registryData, isLoading: registryLoading, error: registryError } = useTemplates();
  const { data: cacheData, isLoading: cacheLoading, error: cacheError } = useCachedTemplates();
  const updateRegistry = useUpdateRegistry();

  const handleUpdateRegistry = async () => {
    try {
      await updateRegistry.mutateAsync();
      notifications.show({
        title: 'Success',
        message: 'Registry updated successfully',
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to update registry',
        color: 'red',
      });
    }
  };

  const filteredTemplates = useMemo(() => {
    if (!registryData?.templates) return [];
    if (!searchQuery) return registryData.templates;

    const queryLower = searchQuery.toLowerCase();
    return registryData.templates.filter(
      (t) =>
        t.name.toLowerCase().includes(queryLower) ||
        t.description?.toLowerCase().includes(queryLower)
    );
  }, [registryData, searchQuery]);

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="center">
        <Title order={2}>Templates</Title>
        <Button
          leftSection={<IconRefresh size={16} />}
          onClick={handleUpdateRegistry}
          loading={updateRegistry.isPending}
          variant="light"
        >
          Update Registry
        </Button>
      </Group>

      {registryError && (
        <Alert icon={<IconAlertCircle size={16} />} title="Error loading registry" color="red">
          {registryError instanceof Error ? registryError.message : 'Failed to load templates'}
        </Alert>
      )}

      {cacheError && (
        <Alert icon={<IconAlertCircle size={16} />} title="Error loading cache" color="yellow">
          {cacheError instanceof Error ? cacheError.message : 'Failed to load installed templates'}
        </Alert>
      )}

      <TextInput
        placeholder="Search templates..."
        leftSection={<IconSearch size={16} />}
        value={searchQuery}
        onChange={(e) => setSearchQuery(e.target.value)}
      />

      <Tabs value={activeTab} onChange={setActiveTab}>
        <Tabs.List>
          <Tabs.Tab value="registry">
            Registry ({registryData?.count ?? 0})
          </Tabs.Tab>
          <Tabs.Tab value="installed">
            Installed ({cacheData?.count ?? 0})
          </Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="registry" pt="md">
          <RegistryTemplateList
            templates={filteredTemplates}
            isLoading={registryLoading}
          />
        </Tabs.Panel>

        <Tabs.Panel value="installed" pt="md">
          <CachedTemplateList
            templates={cacheData?.templates ?? []}
            isLoading={cacheLoading}
          />
        </Tabs.Panel>
      </Tabs>
    </Stack>
  );
}
