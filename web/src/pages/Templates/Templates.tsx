import { Suspense, useState, useMemo } from 'react';
import { Stack, Title, Tabs, TextInput, Button, Group } from '@mantine/core';
import { IconSearch, IconRefresh } from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { useTemplates, useCachedTemplates, useUpdateRegistry } from '@/api/queries';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { RegistryTemplateList } from './components/RegistryTemplateList';
import { CachedTemplateList } from './components/CachedTemplateList';

function TemplatesContent() {
  const [searchQuery, setSearchQuery] = useState('');
  const [activeTab, setActiveTab] = useState<string | null>('registry');

  // No need for optional chaining - data is guaranteed to exist with useSuspenseQuery
  const { data: registryData } = useTemplates();
  const { data: cacheData } = useCachedTemplates();
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
    if (!searchQuery) return registryData.templates;

    const queryLower = searchQuery.toLowerCase();
    return registryData.templates.filter(
      (t) =>
        t.name.toLowerCase().includes(queryLower) ||
        t.description?.toLowerCase().includes(queryLower)
    );
  }, [registryData.templates, searchQuery]);

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

      <TextInput
        placeholder="Search templates..."
        leftSection={<IconSearch size={16} />}
        value={searchQuery}
        onChange={(e) => setSearchQuery(e.target.value)}
      />

      <Tabs value={activeTab} onChange={setActiveTab}>
        <Tabs.List>
          <Tabs.Tab value="registry">
            Registry ({registryData.count})
          </Tabs.Tab>
          <Tabs.Tab value="installed">
            Installed ({cacheData.count})
          </Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="registry" pt="md">
          <RegistryTemplateList templates={filteredTemplates} />
        </Tabs.Panel>

        <Tabs.Panel value="installed" pt="md">
          <CachedTemplateList templates={cacheData.templates} />
        </Tabs.Panel>
      </Tabs>
    </Stack>
  );
}

export function Templates() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading templates..." />}>
        <TemplatesContent />
      </Suspense>
    </ErrorBoundary>
  );
}
