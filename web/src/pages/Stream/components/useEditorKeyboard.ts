import { useEffect } from 'react';
import { useEditorContext } from './EditorContext';
import { useEditorMessages, useEditorExport } from '@/api/queries';
import { notifications } from '@mantine/notifications';

export function useEditorKeyboard() {
  const {
    sessionStats,
    clearSelection,
    setSelectedIndices,
    currentFilters,
  } = useEditorContext();

  const { data: messagesData } = useEditorMessages(
    sessionStats?.sessionId || null,
    0,
    10000,
    currentFilters
  );

  const exportMutation = useEditorExport();

  useEffect(() => {
    if (!sessionStats) return;

    const handleKeyDown = (e: KeyboardEvent) => {
      // Ctrl/Cmd + A - Select All
      if ((e.ctrlKey || e.metaKey) && e.key === 'a') {
        e.preventDefault();
        if (messagesData) {
          const allIndices = new Set(messagesData.messages.map((m) => m.index));
          setSelectedIndices(allIndices);
          notifications.show({
            title: 'All Selected',
            message: `Selected ${allIndices.size} messages`,
            color: 'blue',
          });
        }
      }

      // Escape - Clear Selection
      if (e.key === 'Escape') {
        clearSelection();
      }

      // Ctrl/Cmd + E - Export
      if ((e.ctrlKey || e.metaKey) && e.key === 'e') {
        e.preventDefault();
        if (sessionStats) {
          exportMutation.mutateAsync({
            sessionId: sessionStats.sessionId,
          }).then((blob) => {
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.href = url;
            a.download = sessionStats.filename;
            document.body.appendChild(a);
            a.click();
            window.URL.revokeObjectURL(url);
            document.body.removeChild(a);

            notifications.show({
              title: 'Export Complete',
              message: `Downloaded ${sessionStats.filename}`,
              color: 'green',
            });
          }).catch((error) => {
            notifications.show({
              title: 'Export Failed',
              message: error instanceof Error ? error.message : 'Export failed',
              color: 'red',
            });
          });
        }
      }
    };

    document.addEventListener('keydown', handleKeyDown);
    return () => document.removeEventListener('keydown', handleKeyDown);
  }, [sessionStats, clearSelection, messagesData, setSelectedIndices, exportMutation]);
}
