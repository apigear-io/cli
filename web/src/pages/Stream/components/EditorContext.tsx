/**
 * Editor Context - Shared state for the stream editor feature
 *
 * Manages both server state (session stats) and client state (selections, marks, filters)
 * Uses Sets for efficient O(1) lookups on large datasets.
 */

import { createContext, useContext, useState, type ReactNode } from 'react';
import type { EditorStats, EditorFilters } from '@/api/types';

export type TimelineSelection = {
  start: number;
  end: number;
};

export type EditorContextValue = {
  // Session state (from server)
  sessionStats: EditorStats | null;
  setSessionStats: (stats: EditorStats | null) => void;

  // Client-side selection state
  selectedIndices: Set<number>;
  setSelectedIndices: (indices: Set<number>) => void;
  toggleSelection: (index: number) => void;
  clearSelection: () => void;

  // Marked messages (starred)
  markedIndices: Set<number>;
  setMarkedIndices: (indices: Set<number>) => void;
  toggleMarked: (index: number) => void;

  // Soft-deleted messages (cut)
  deletedIndices: Set<number>;
  setDeletedIndices: (indices: Set<number>) => void;

  // Filter state
  currentFilters: EditorFilters;
  setCurrentFilters: (filters: EditorFilters) => void;

  // Timeline selection (drag selection on timeline)
  timelineSelection: TimelineSelection | null;
  setTimelineSelection: (selection: TimelineSelection | null) => void;

  // Scroll target (for jump to message feature)
  scrollToIndex: number | null;
  setScrollToIndex: (index: number | null) => void;

  // UI state
  hideDeleted: boolean;
  setHideDeleted: (hide: boolean) => void;
  showMarkedOnly: boolean;
  setShowMarkedOnly: (show: boolean) => void;

  // Helpers
  isSelected: (index: number) => boolean;
  isMarked: (index: number) => boolean;
  isDeleted: (index: number) => boolean;
};

const EditorContext = createContext<EditorContextValue | null>(null);

export function EditorProvider({ children }: { children: ReactNode }) {
  // Session state
  const [sessionStats, setSessionStats] = useState<EditorStats | null>(null);

  // Selection state
  const [selectedIndices, setSelectedIndices] = useState<Set<number>>(new Set());
  const [markedIndices, setMarkedIndices] = useState<Set<number>>(new Set());
  const [deletedIndices, setDeletedIndices] = useState<Set<number>>(new Set());

  // Filter state
  const [currentFilters, setCurrentFilters] = useState<EditorFilters>({});

  // Timeline selection
  const [timelineSelection, setTimelineSelection] = useState<TimelineSelection | null>(null);

  // Scroll target for jump to message
  const [scrollToIndex, setScrollToIndex] = useState<number | null>(null);

  // UI state
  const [hideDeleted, setHideDeleted] = useState(false);
  const [showMarkedOnly, setShowMarkedOnly] = useState(false);

  // Helper functions
  const toggleSelection = (index: number) => {
    const newSet = new Set(selectedIndices);
    if (newSet.has(index)) {
      newSet.delete(index);
    } else {
      newSet.add(index);
    }
    setSelectedIndices(newSet);
  };

  const clearSelection = () => {
    setSelectedIndices(new Set());
  };

  const toggleMarked = (index: number) => {
    const newSet = new Set(markedIndices);
    if (newSet.has(index)) {
      newSet.delete(index);
    } else {
      newSet.add(index);
    }
    setMarkedIndices(newSet);
  };

  const isSelected = (index: number) => selectedIndices.has(index);
  const isMarked = (index: number) => markedIndices.has(index);
  const isDeleted = (index: number) => deletedIndices.has(index);

  const value: EditorContextValue = {
    sessionStats,
    setSessionStats,
    selectedIndices,
    setSelectedIndices,
    toggleSelection,
    clearSelection,
    markedIndices,
    setMarkedIndices,
    toggleMarked,
    deletedIndices,
    setDeletedIndices,
    currentFilters,
    setCurrentFilters,
    timelineSelection,
    setTimelineSelection,
    scrollToIndex,
    setScrollToIndex,
    hideDeleted,
    setHideDeleted,
    showMarkedOnly,
    setShowMarkedOnly,
    isSelected,
    isMarked,
    isDeleted,
  };

  return <EditorContext.Provider value={value}>{children}</EditorContext.Provider>;
}

// eslint-disable-next-line react-refresh/only-export-components
export function useEditorContext() {
  const context = useContext(EditorContext);
  if (!context) {
    throw new Error('useEditorContext must be used within EditorProvider');
  }
  return context;
}
