import { useState } from 'react';
import { Container, Stack, Group, Title, Button } from '@mantine/core';
import { IconFilePlus } from '@tabler/icons-react';
import { EditorProvider, useEditorContext } from './components/EditorContext';
import { EditorWelcome } from './components/EditorWelcome';
import { EditorLoadDrawer } from './components/EditorLoadDrawer';
import { EditorStats } from './components/EditorStats';
import { EditorTimeline } from './components/EditorTimeline';
import { EditorFilters } from './components/EditorFilters';
import { EditorJQPanel } from './components/EditorJQPanel';
import { EditorTable } from './components/EditorTable';
import { EditorToolbar } from './components/EditorToolbar';
import { useEditorKeyboard } from './components/useEditorKeyboard';

function EditorContent() {
  const { sessionStats } = useEditorContext();
  const [loadDrawerOpen, setLoadDrawerOpen] = useState(false);

  return (
    <>
      {!sessionStats ? (
        <EditorWelcome onOpenLoad={() => setLoadDrawerOpen(true)} />
      ) : (
        <EditorContentWithSession loadDrawerOpen={loadDrawerOpen} setLoadDrawerOpen={setLoadDrawerOpen} />
      )}

      <EditorLoadDrawer opened={loadDrawerOpen} onClose={() => setLoadDrawerOpen(false)} />
    </>
  );
}

function EditorContentWithSession({
  setLoadDrawerOpen,
}: {
  loadDrawerOpen: boolean;
  setLoadDrawerOpen: (open: boolean) => void;
}) {
  useEditorKeyboard();

  return (
    <Stack gap="md">
      {/* Header with Set Stream button */}
      <Group justify="space-between">
        <Title order={2}>Stream Editor</Title>
        <Button leftSection={<IconFilePlus size={16} />} onClick={() => setLoadDrawerOpen(true)}>
          Set Stream
        </Button>
      </Group>

      <EditorStats />
      <EditorTimeline />
      <EditorFilters />
      <EditorJQPanel />
      <EditorTable />
      <EditorToolbar />
    </Stack>
  );
}

export function StreamEditor() {
  return (
    <EditorProvider>
      <Container size="xl" py="md">
        <Stack gap="md">
          <EditorContent />
        </Stack>
      </Container>
    </EditorProvider>
  );
}
