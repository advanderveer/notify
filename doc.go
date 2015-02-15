// Copyright (c) 2014-2015 The Notify Authors. All rights reserved.
// Use of this source code is governed by the MIT license that can be
// found in the LICENSE file.

// Package notify implements access to filesystem events.
//
// Notify is a high-level abstraction over filesystem watchers like inotify,
// kqueue, FSEvents or ReadDirectoryChangesW. Watcher implementations are
// split into two groups: ones that natively support recursive notifications
// (FSEvents and ReadDirectoryChangesW) and ones that do not (inotify and kqueue).
// For more details see watcher and recursiveWatcher interfaces in watcher.go
// source file.
//
// On top of filesystem watchers notify maintains a watchpoint tree, which provides
// strategy for creating and closing filesystem watches and dispatching filesystem
// events to user channels.
//
// An event set is just an event list joint using bitwise and operator
// into a single event value.
//
// A filesystem watch or just a watch is platform-specific entity which represents
// single path registered for notifications for specific event set. Setting a watch
// means using platform-specific API call for creating / initialising said watch.
// For each watcher the API call is:
//
//   - inotify:  notify_add_watch
//   - kqueue:   kevent
//   - FSEvents: FSEventStreamCreate
//   - ReadDirectoryChangesW: CreateFile+ReadDirectoryChangesW
//
// To rewatch means to either shrink or expand an event set that was previously
// registered during watch operation for particular filesystem watch.
//
// A watchpoint is a list of user channel and event set pairs for particular
// path (watchpoint tree's node). A single watchpoint can contain multiple
// different user channels registered to listen for one or more events. A single
// user channel can be registered in one or more watchpoints, recurisve and
// non-recursive ones as well.
package notify
