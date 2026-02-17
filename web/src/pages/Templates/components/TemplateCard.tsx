import { useState } from 'react';
import { Card, Stack, Group, Text, Badge, Button, Progress, ActionIcon, Tooltip, Menu, Select } from '@mantine/core';
import { notifications } from '@mantine/notifications';
import { IconCheck, IconAlertCircle, IconBrandGithub, IconChevronDown } from '@tabler/icons-react';
import { useInstallTemplate } from '@/api/queries';
import type { TemplateInfo, InstallProgressEvent } from '@/api/types';

interface TemplateCardProps {
  template: TemplateInfo;
}

export function TemplateCard({ template }: TemplateCardProps) {
  const [installing, setInstalling] = useState(false);
  const [progress, setProgress] = useState(0);
  const [progressMessage, setProgressMessage] = useState('');
  const [selectedVersion, setSelectedVersion] = useState<string>(template.latest || '');
  const installMutation = useInstallTemplate();

  const handleInstall = async (version?: string) => {
    setInstalling(true);
    setProgress(0);

    const versionToInstall = version || selectedVersion || template.latest;

    try {
      await installMutation.mutateAsync({
        id: template.name,
        version: versionToInstall,
        onProgress: (event: InstallProgressEvent) => {
          setProgress(event.progress);
          setProgressMessage(event.message);
        },
      });

      notifications.show({
        title: 'Success',
        message: `Template ${template.name} ${versionToInstall} installed successfully`,
        color: 'green',
        icon: <IconCheck size={18} />,
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Installation failed',
        color: 'red',
        icon: <IconAlertCircle size={18} />,
      });
    } finally {
      setInstalling(false);
      setProgress(0);
      setProgressMessage('');
    }
  };

  // Use server-calculated updateNeeded flag (based on semver comparison)
  const isUpToDate = template.inCache && !template.updateNeeded;
  const hasUpdate = template.updateNeeded;

  // Extract GitHub URL (remove .git suffix if present)
  const githubUrl = template.git ? template.git.replace(/\.git$/, '') : null;

  return (
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Stack gap="md">
        <Group justify="space-between" align="flex-start">
          <Group gap="xs" style={{ flex: 1, minWidth: 0 }}>
            <Text fw={600} size="sm" truncate>
              {template.name}
            </Text>
            {githubUrl && (
              <Tooltip label="View on GitHub">
                <ActionIcon
                  component="a"
                  href={githubUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  variant="subtle"
                  color="gray"
                  size="sm"
                >
                  <IconBrandGithub size={16} />
                </ActionIcon>
              </Tooltip>
            )}
          </Group>
          <Group gap="xs">
            {template.inCache && (
              <Badge color="green" size="sm" variant="light">
                Installed
              </Badge>
            )}
            {hasUpdate && (
              <Badge color="orange" size="sm" variant="light">
                Update Available
              </Badge>
            )}
          </Group>
        </Group>

        <Text size="sm" c="dimmed" lineClamp={2} style={{ minHeight: '2.5rem' }}>
          {template.description || 'No description available'}
        </Text>

        <Group gap="xs" wrap="nowrap">
          <Text size="xs" c="dimmed" truncate>
            Latest: {template.latest || 'N/A'}
          </Text>
          {template.inCache && template.version && (
            <Text size="xs" c="dimmed" truncate>
              • Installed: {template.version}
            </Text>
          )}
        </Group>

        {installing ? (
          <Stack gap="xs">
            <Progress value={progress} animated />
            <Text size="xs" c="dimmed" ta="center">
              {progressMessage}
            </Text>
          </Stack>
        ) : (
          <Stack gap="xs">
            {template.versions && template.versions.length > 1 && (
              <Select
                label="Version"
                placeholder="Select version"
                value={selectedVersion}
                onChange={(value) => setSelectedVersion(value || template.latest)}
                data={template.versions.map((v) => ({
                  value: v,
                  label: v === template.latest ? `${v} (Latest)` : v,
                }))}
                size="xs"
              />
            )}
            <Group gap="xs">
              <Button
                flex={1}
                onClick={() => handleInstall()}
                disabled={isUpToDate}
                variant={hasUpdate ? 'filled' : 'light'}
                loading={installMutation.isPending}
              >
                {isUpToDate ? 'Up to Date' : template.inCache ? 'Update' : 'Install'}
              </Button>
              {template.versions && template.versions.length > 1 && (
                <Menu shadow="md" width={200}>
                  <Menu.Target>
                    <ActionIcon
                      variant="light"
                      size="lg"
                      disabled={installMutation.isPending}
                    >
                      <IconChevronDown size={18} />
                    </ActionIcon>
                  </Menu.Target>
                  <Menu.Dropdown>
                    <Menu.Label>Install specific version</Menu.Label>
                    {template.versions.slice(0, 10).map((version) => (
                      <Menu.Item
                        key={version}
                        onClick={() => handleInstall(version)}
                      >
                        {version === template.latest ? `${version} (Latest)` : version}
                      </Menu.Item>
                    ))}
                    {template.versions.length > 10 && (
                      <Menu.Label>
                        +{template.versions.length - 10} more versions
                      </Menu.Label>
                    )}
                  </Menu.Dropdown>
                </Menu>
              )}
            </Group>
          </Stack>
        )}
      </Stack>
    </Card>
  );
}
