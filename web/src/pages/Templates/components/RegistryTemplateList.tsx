import { Grid, Stack, Text, Center } from '@mantine/core';
import { IconMoodEmpty } from '@tabler/icons-react';
import { TemplateCard } from './TemplateCard';
import type { TemplateInfo } from '@/api/types';

interface RegistryTemplateListProps {
  templates: TemplateInfo[];
}

export function RegistryTemplateList({ templates }: RegistryTemplateListProps) {
  if (templates.length === 0) {
    return (
      <Center py="xl">
        <Stack align="center" gap="md">
          <IconMoodEmpty size={48} stroke={1.5} opacity={0.5} />
          <Text c="dimmed">No templates found</Text>
        </Stack>
      </Center>
    );
  }

  return (
    <Grid gutter="md">
      {templates.map((template) => (
        <Grid.Col key={template.name} span={{ base: 12, sm: 6, md: 4 }}>
          <TemplateCard template={template} />
        </Grid.Col>
      ))}
    </Grid>
  );
}
