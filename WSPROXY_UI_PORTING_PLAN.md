# WSProxy UI Porting Plan

This document outlines the plan to port all UI screens from WSProxy to ApiGear CLI's Stream module.

## Overview

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/`
**Screenshots**: `/Users/jryannel/dev/github.com/apigear-io/cli/wsproxy_screenshots/`
**Target**: `/Users/jryannel/dev/github.com/apigear-io/cli/web/src/pages/Stream/`

## Status Summary

| Feature | Status | Priority | Complexity | Screenshot |
|---------|--------|----------|------------|------------|
| **Stream Editor** | ✅ Implemented | HIGH | High | wsproxy_stream_editor.png |
| **Dashboard** | 🟡 Needs Enhancement | HIGH | Medium | wsproxy_dashboard.png |
| **Proxies** | 🟡 Needs Enhancement | HIGH | Low | wsproxy_proxy.png |
| **Clients** | 🟡 Needs Enhancement | HIGH | Low | wsproxy_clients.png |
| **Scripting** | 🟡 Needs Enhancement | HIGH | Medium | wsproxy_scripting.png |
| **Traces (Files)** | 🟡 Needs Enhancement | MEDIUM | Low | wsproxy_stream_files.png |
| **Stream Player** | ❌ Not Implemented | MEDIUM | High | wsproxy_stream_player.png |
| **Trace Generator** | ❌ Not Implemented | LOW | High | wsproxy_stream_generator.png |
| **Application Logs** | ❌ Not Implemented | LOW | Medium | wsproxy_logs.png |
| **Proxy Stream (Live)** | ❌ Not Implemented | MEDIUM | High | (part of dashboard) |
| **Settings** | ❌ Not Implemented | LOW | Medium | wsproxy_settings_*.png |

## Detailed Breakdown

### ✅ Phase 1: Stream Editor (COMPLETED)
**Status**: Fully implemented with all 8 components
**Files**:
- `StreamEditor.tsx` - Main page
- `EditorContext.tsx` - State management
- `EditorWelcome.tsx` - Welcome screen
- `EditorLoadDrawer.tsx` - File upload/selection
- `EditorStats.tsx` - Session stats
- `EditorTimeline.tsx` - Canvas timeline
- `EditorFilters.tsx` - Filter controls
- `EditorJQPanel.tsx` - JQ queries
- `EditorTable.tsx` - Message table
- `EditorToolbar.tsx` - Actions toolbar
- `useEditorKeyboard.ts` - Keyboard shortcuts

**Improvements Needed**:
- Visual polish to exactly match wsproxy design
- Stats bar should be inline not stacked
- Toolbar button grouping
- Select controls (None, Filtered, Invert)

---

### 🟡 Phase 2: Visual Enhancements (HIGH Priority)

#### 2.1 Dashboard Enhancement
**Current**: Basic stats display
**Target**: Rich analytics dashboard with:
- Quick Action cards (6 cards: Proxy Stream, Stream Editor, Stream Player, Scripting, Proxies, Settings)
- Architecture flow diagram (INBOUND → PROXY → OUTBOUND)
- Component boxes showing data flow
- Active Connections, Messages In/Out, Total Throughput metrics
- Proxy Statistics table

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/dashboard/`
**Effort**: 4-6 hours
**Backend**: No new APIs needed

#### 2.2 Proxies List Enhancement
**Current**: Basic table view
**Target**: Card-based layout with:
- Status badges and colored dots
- IN → OUT address display
- Stats icons (connections, messages, bytes)
- Action buttons (view stats, edit, delete)

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/proxies/components/ProxyCard.tsx`
**Effort**: 2-3 hours
**Backend**: No new APIs needed

#### 2.3 Clients List Enhancement
**Current**: Basic table view
**Target**: Card-based layout with:
- Status badges and colored dots
- WebSocket URL display
- Interface badges
- Action buttons (retry, connect/disconnect, edit, delete)

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/clients/components/ClientCard.tsx`
**Effort**: 2-3 hours
**Backend**: No new APIs needed

#### 2.4 Scripting Page Enhancement
**Current**: Basic editor with script list
**Target**: Full IDE-like layout with:
- Script list sidebar
- Monaco editor with syntax highlighting
- Toolbar actions (New, Generate from Module, Insert Type, Insert Faker, Save)
- Console/Messages tabs at bottom
- Start/Stop script controls

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/scripting/`
**Effort**: 3-4 hours
**Backend**: Already implemented

#### 2.5 Traces Page Enhancement
**Current**: Basic file list
**Target**: Clean table design with:
- Directory path display
- Filter dropdowns
- File count badge
- Icon buttons for actions
- Compact table layout

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/traces/`
**Effort**: 1-2 hours
**Backend**: No new APIs needed

---

### ❌ Phase 3: New Features (MEDIUM/LOW Priority)

#### 3.1 Stream Player (MEDIUM Priority)
**Purpose**: Replay trace files to live proxies
**Features**:
- Select target proxy
- Choose trace file from directory or upload
- Playback speed control (0.5x, 1x, 2x, 5x)
- Initial delay setting
- Loop playback option
- Direction filter (Both, SEND, RECV)
- Play/Pause/Stop controls
- Progress indicator

**Backend Needed**:
- Stream playback API endpoint
- WebSocket or HTTP streaming for playback
- Playback session management

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/player/`
**Effort**: 8-12 hours (6h frontend, 4-6h backend)

#### 3.2 Trace Generator (LOW Priority)
**Purpose**: Generate synthetic trace files using Go templates
**Features**:
- Monaco editor for Go templates
- Template management (load, save)
- Insert Type helper (ObjectLink message types)
- Insert Faker helper (faker functions)
- Example templates
- Preview output
- Lines to generate setting
- Save as trace file

**Backend Needed**:
- Template execution API
- Faker library integration
- Template preview endpoint
- Template save/load endpoints

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/generator/`
**Effort**: 10-15 hours (6h frontend, 6-9h backend)

#### 3.3 Application Logs (LOW Priority)
**Purpose**: View application logs in UI
**Features**:
- Log level filtering
- Search/filter messages
- Timestamp display
- Level badges (WARN, INFO, ERROR)
- JSON fields expansion
- Auto-refresh
- Export logs
- Clear logs

**Backend Needed**:
- Logging API endpoint
- Log streaming (WebSocket or SSE)
- Log filtering backend

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/logs/`
**Effort**: 6-8 hours (4h frontend, 2-4h backend)

#### 3.4 Proxy Stream / Live Viewer (MEDIUM Priority)
**Purpose**: Real-time message viewing
**Features**:
- Live message streaming from proxies
- Filter by proxy, direction, type
- Auto-scroll to new messages
- Pause/resume
- Clear messages
- Message detail view

**Backend Needed**:
- WebSocket endpoint for live message streaming
- Message buffering
- Filter support in stream

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/stream/`
**Effort**: 8-12 hours (5h frontend, 5-7h backend)

#### 3.5 Settings Pages (LOW Priority)
**Purpose**: Configure application settings
**Features**:
- Tabs: General, Buffer, Traces, Advanced
- Form inputs for each setting category
- Save/reset functionality
- Validation

**Backend Needed**:
- Settings API (GET/PUT)
- Settings persistence
- Settings validation

**Source**: `/Users/jryannel/dev/tmp/wsproxy/web2/src/features/settings/`
**Effort**: 6-8 hours (4h frontend, 2-4h backend)

---

## Implementation Priority

### Sprint 1: Visual Polish (HIGH Priority - Week 1)
1. ✅ Stream Editor improvements (Task #22) - 2h
2. Dashboard enhancement (Task #12) - 6h
3. Proxies enhancement (Task #13) - 3h
4. Clients enhancement (Task #14) - 3h
5. Scripting enhancement (Task #15) - 4h
6. Traces enhancement (Task #16) - 2h

**Total Effort**: ~20 hours
**Outcome**: Professional, polished UI matching wsproxy design

### Sprint 2: Stream Player (MEDIUM Priority - Week 2)
1. Backend: Playback API (Task #17) - 6h
2. Frontend: Stream Player UI (Task #17) - 6h
3. Testing and integration - 2h

**Total Effort**: ~14 hours
**Outcome**: Trace file playback capability

### Sprint 3: Nice-to-Have Features (LOW Priority - Week 3-4)
1. Application Logs (Task #19) - 8h
2. Proxy Stream / Live Viewer (Task #20) - 12h
3. Trace Generator (Task #18) - 15h
4. Settings Pages (Task #21) - 8h

**Total Effort**: ~43 hours
**Outcome**: Feature parity with wsproxy

---

## Code Reuse Strategy

### Direct Copy (Minimal Changes)
These components can be copied almost directly:
- Card layouts (ProxyCard, ClientCard)
- Monaco editor setup
- Architecture diagram component
- Table styling and layouts

### Adapt (Moderate Changes)
These need adaptation for our API structure:
- API hooks (different endpoint structure)
- State management (we use React Query, not custom)
- Routing (React Router v7 vs v6)

### Rewrite (Significant Changes)
These require significant changes:
- Backend API implementations (Go vs existing backend)
- WebSocket integration (different message format)
- Configuration management (different settings structure)

---

## Navigation Structure

```
Stream Module
├── Dashboard (Analytics)
├── Proxies
│   └── Proxy Detail (with live stream)
├── Clients
│   └── Client Detail
├── Scripting
├── Stream Editor (Trace Analysis)
├── Stream Files (Traces)
├── Stream Player (Playback)
├── Generator (Create Traces)
├── Logs (Application Logs)
└── Settings
    ├── General
    ├── Buffer
    ├── Traces
    └── Advanced
```

---

## Technical Considerations

### Dependencies to Add
```json
{
  "@monaco-editor/react": "^4.6.0",  // Code editor (already have?)
  "react-window": "^1.8.10",          // Virtual scrolling (if not added)
}
```

### Backend APIs Needed

**Already Implemented**:
- ✅ Proxy CRUD
- ✅ Client CRUD
- ✅ Script CRUD & execution
- ✅ Trace files CRUD
- ✅ Stream Editor (load, messages, timeline, jq, export)

**Need Implementation**:
- ❌ Stream Player (playback API)
- ❌ Trace Generator (template execution API)
- ❌ Application Logs (logging API)
- ❌ Live Message Stream (WebSocket endpoint)
- ❌ Settings (configuration API)

---

## Success Criteria

### Phase 1 (Visual Polish)
- [ ] Dashboard matches wsproxy design
- [ ] Proxies use card layout with proper styling
- [ ] Clients use card layout with proper styling
- [ ] Scripting has full toolbar and console
- [ ] Traces has clean table design
- [ ] Stream Editor matches pixel-perfect design

### Phase 2 (Stream Player)
- [ ] Can load trace file to player
- [ ] Can select target proxy
- [ ] Playback speed controls work
- [ ] Direction filter works
- [ ] Play/Pause/Stop controls work

### Phase 3 (Complete Feature Parity)
- [ ] Trace generator works with templates
- [ ] Application logs display correctly
- [ ] Live message viewer works
- [ ] Settings pages functional
- [ ] All features match wsproxy functionality

---

## Next Steps

1. **Review this plan** - Get approval on priorities
2. **Start Sprint 1** - Begin with Dashboard enhancement
3. **Iterative development** - Complete one task at a time
4. **Test each feature** - Ensure quality before moving forward
5. **Document as we go** - Update this plan with learnings

---

## Notes

- All wsproxy source code is available at `/Users/jryannel/dev/tmp/wsproxy/web2/`
- We can freely copy and adapt code since it's our own codebase
- Focus on visual consistency and UX quality
- Backend APIs can be implemented progressively as needed
- Stream Editor is already a strong foundation - use as reference for other pages
