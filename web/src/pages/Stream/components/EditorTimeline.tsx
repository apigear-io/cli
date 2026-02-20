import { useRef, useEffect, useState, MouseEvent } from 'react';
import { Box, Group, Text, Button } from '@mantine/core';
import { IconX } from '@tabler/icons-react';
import { useEditorContext } from './EditorContext';
import { useEditorTimeline, useEditorSeek } from '@/api/queries';

const TIMELINE_HEIGHT = 80;
const NUM_BUCKETS = 200;

export function EditorTimeline() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  const [canvasWidth, setCanvasWidth] = useState(800);
  const {
    sessionStats,
    markedIndices,
    timelineSelection,
    setTimelineSelection,
    currentFilters
  } = useEditorContext();

  const { data: timelineData } = useEditorTimeline(sessionStats?.sessionId || null);
  const seekMutation = useEditorSeek();

  const [hoveredBucket, setHoveredBucket] = useState<number | null>(null);
  const [isDragging, setIsDragging] = useState(false);
  const [dragStart, setDragStart] = useState<number | null>(null);

  // Responsive canvas width
  useEffect(() => {
    if (!containerRef.current) return;

    const observer = new ResizeObserver((entries) => {
      const entry = entries[0];
      if (entry) {
        setCanvasWidth(entry.contentRect.width);
      }
    });

    observer.observe(containerRef.current);
    return () => observer.disconnect();
  }, []);

  // Draw canvas
  useEffect(() => {
    if (!canvasRef.current || !timelineData) return;

    const canvas = canvasRef.current;
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    // Set canvas size
    canvas.width = canvasWidth;
    canvas.height = TIMELINE_HEIGHT;

    // Clear
    ctx.fillStyle = '#f8f9fa';
    ctx.fillRect(0, 0, canvasWidth, TIMELINE_HEIGHT);

    // Draw buckets
    const bucketWidth = canvasWidth / NUM_BUCKETS;
    const maxCount = Math.max(...timelineData.buckets.map(b => b.sendCount + b.recvCount), 1);

    timelineData.buckets.forEach((bucket, i) => {
      const x = i * bucketWidth;
      const sendHeight = (bucket.sendCount / maxCount) * (TIMELINE_HEIGHT - 20);
      const recvHeight = (bucket.recvCount / maxCount) * (TIMELINE_HEIGHT - 20);

      // Draw SEND (blue)
      ctx.fillStyle = '#228be6';
      ctx.fillRect(x, TIMELINE_HEIGHT - sendHeight - 10, bucketWidth - 1, sendHeight);

      // Draw RECV (green) on top
      ctx.fillStyle = '#40c057';
      ctx.fillRect(x, TIMELINE_HEIGHT - sendHeight - recvHeight - 10, bucketWidth - 1, recvHeight);

      // Highlight hovered bucket
      if (i === hoveredBucket) {
        ctx.strokeStyle = '#000';
        ctx.lineWidth = 2;
        ctx.strokeRect(x, 0, bucketWidth, TIMELINE_HEIGHT);
      }
    });

    // Draw selection range
    if (timelineSelection) {
      const startX = timelineSelection.start * bucketWidth;
      const endX = (timelineSelection.end + 1) * bucketWidth;
      ctx.fillStyle = 'rgba(34, 139, 230, 0.2)';
      ctx.fillRect(startX, 0, endX - startX, TIMELINE_HEIGHT);
      ctx.strokeStyle = 'rgba(34, 139, 230, 0.6)';
      ctx.lineWidth = 2;
      ctx.strokeRect(startX, 0, endX - startX, TIMELINE_HEIGHT);
    }

    // TODO: Draw gold flags for marked messages
    // This requires mapping message indices to time buckets, which needs message timestamp data
    // For now, we skip this feature
  }, [timelineData, canvasWidth, hoveredBucket, timelineSelection, markedIndices]);

  const handleMouseMove = (e: MouseEvent<HTMLCanvasElement>) => {
    if (!canvasRef.current || !timelineData) return;
    const rect = canvasRef.current.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const bucketIndex = Math.floor((x / canvasWidth) * NUM_BUCKETS);

    if (isDragging && dragStart !== null) {
      setTimelineSelection({
        start: Math.min(dragStart, bucketIndex),
        end: Math.max(dragStart, bucketIndex),
      });
    } else {
      setHoveredBucket(bucketIndex);
    }
  };

  const handleMouseDown = (e: MouseEvent<HTMLCanvasElement>) => {
    if (!canvasRef.current) return;
    const rect = canvasRef.current.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const bucketIndex = Math.floor((x / canvasWidth) * NUM_BUCKETS);

    setIsDragging(true);
    setDragStart(bucketIndex);
    setTimelineSelection(null);
  };

  const handleMouseUp = async (e: MouseEvent<HTMLCanvasElement>) => {
    if (!canvasRef.current || !isDragging || dragStart === null) return;
    const rect = canvasRef.current.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const bucketIndex = Math.floor((x / canvasWidth) * NUM_BUCKETS);

    if (dragStart === bucketIndex) {
      // Single click - seek to this timestamp
      if (timelineData && sessionStats) {
        const bucket = timelineData.buckets[bucketIndex];
        if (bucket) {
          try {
            await seekMutation.mutateAsync({
              sessionId: sessionStats.sessionId,
              timestamp: bucket.startTime,
              filters: currentFilters,
            });
          } catch (error) {
            console.error('Seek failed:', error);
          }
        }
      }
      setTimelineSelection(null);
    } else {
      // Drag selection
      setTimelineSelection({
        start: Math.min(dragStart, bucketIndex),
        end: Math.max(dragStart, bucketIndex),
      });
    }

    setIsDragging(false);
    setDragStart(null);
  };

  if (!timelineData) return null;

  return (
    <Box p="md" style={{ backgroundColor: '#f8f9fa', borderBottom: '1px solid #dee2e6' }} ref={containerRef}>
      <Group justify="space-between" mb="xs">
        <Text size="sm" c="dimmed">
          Timeline (click to jump, drag to select range)
        </Text>
        {timelineSelection && (
          <Button
            size="xs"
            variant="subtle"
            leftSection={<IconX size={14} />}
            onClick={() => setTimelineSelection(null)}
          >
            Clear Selection
          </Button>
        )}
      </Group>

      <canvas
        ref={canvasRef}
        onMouseMove={handleMouseMove}
        onMouseDown={handleMouseDown}
        onMouseUp={handleMouseUp}
        onMouseLeave={() => {
          setHoveredBucket(null);
          if (isDragging) {
            setIsDragging(false);
            setDragStart(null);
          }
        }}
        style={{
          width: '100%',
          height: TIMELINE_HEIGHT,
          cursor: isDragging ? 'crosshair' : 'pointer',
          display: 'block',
          border: '1px solid #dee2e6',
          borderRadius: '4px',
        }}
      />

      {hoveredBucket !== null && timelineData.buckets[hoveredBucket] && (
        <Text size="xs" c="dimmed" mt="xs">
          Bucket {hoveredBucket}: {timelineData.buckets[hoveredBucket].sendCount} sent, {timelineData.buckets[hoveredBucket].recvCount} received
        </Text>
      )}
    </Box>
  );
}
