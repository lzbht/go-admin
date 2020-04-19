// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package language

import "strings"

var en = LangSet{
	"managers":         "Managers",
	"name":             "Name",
	"nickname":         "Nickname",
	"role":             "Role",
	"createdat":        "createdAt",
	"updatedat":        "updatedAt",
	"path":             "path",
	"new":              "New",
	"action":           "Action",
	"toggle dropdown":  "Toggle Dropdown",
	"delete":           "Delete",
	"refresh":          "Refresh",
	"expand":           "Expand",
	"collapse":         "Collapse",
	"back":             "Back",
	"reset":            "Reset",
	"save":             "Save",
	"edit":             "Edit",
	"operation":        "Operation",
	"method":           "Method",
	"input":            "input",
	"online":           "Online",
	"setting":          "Setting",
	"sign out":         "Sign out",
	"all":              "All",
	"confirm password": "Confirm Password",
	"search":           "Search",

	"are you sure to delete":    "Are you sure to delete",
	"yes":                       "yes",
	"cancel":                    "cancel",
	"refresh succeeded":         "Refresh succeeded",
	"reload succeeded":          "Reload succeeded",
	"all method if empty":       "All method if empty",
	"password does not match":   "Password does not match",
	"should be unique":          "Should be unique",
	"slug exists":               "Slug exists",
	"no corresponding options?": "No corresponding options?",
	"create here.":              "Create here.",
	"use for login":             "Use for login",
	"use to display":            "Use to display",
	"a path a line":             "A path a line",
	"slug or http_path or name should not be empty": "slug or http_path or name should not be empty",
	"no roles":          "no roles",
	"fixed the sidebar": "Fixed the sidebar",
	"enter fullscreen":  "Enter fullscreen",
	"exit fullscreen":   "Exit fullscreen",

	"permission manage": "Permission Manage",
	"menus manage":      "Menus Manage",
	"roles manage":      "Roles manage",
	"operation log":     "Operation log",

	"continue editing":  "Continue editing",
	"continue creating": "Continue creating",

	"browse":     "Browse",
	"avatar":     "Avatar",
	"password":   "Password",
	"username":   "Username",
	"slug":       "Slug",
	"permission": "Permission",
	"userid":     "UserID",
	"content":    "Content",
	"parent":     "Parent",
	"icon":       "Icon",
	"uri":        "Uri",

	"detail": "Detail",

	"admin":     "Admin",
	"users":     "Users",
	"roles":     "Roles",
	"menu":      "Menu",
	"dashboard": "Dashboard",
	"home":      "Home",

	"config.domain":          "Website Domain",
	"config.language":        "Website Language",
	"config.url prefix":      "URL Prefix",
	"config.theme":           "Theme",
	"config.title":           "Title",
	"config.index url":       "Home URL",
	"config.login url":       "Login URL",
	"config.env":             "Env",
	"config.color scheme":    "Color Scheme",
	"config.cdn url":         "CDN Asset URL",
	"config.login title":     "Login Title",
	"config.auth user table": "Auth User Table",
	"config.extra":           "Extra Configuration",
	"config.store":           "File Store Setting",
	"config.databases":       "Database Setting",

	"config.general":      "General",
	"config.log":          "Log",
	"config.site setting": "Site Settings",
	"config.custom":       "Customize",
	"config.debug":        "Debug Mode",
	"config.site off":     "Site Offline",
	"config.true":         "On",
	"config.false":        "Off",

	"config.logo":               "Logo",
	"config.mini logo":          "Mini Logo",
	"config.session life time":  "Session Life Time",
	"config.custom head html":   "Head HTML",
	"config.custom foot html":   "Foot HTML",
	"config.footer info":        "Footer Info",
	"config.login logo":         "Login Logo",
	"config.no limit login ip":  "No Limit Login Multi IPs",
	"config.animation type":     "Animation Type",
	"config.animation duration": "Animation Duration(s)",
	"config.animation delay":    "Animation Delay(s)",
	"config.file upload engine": "File Upload Engine",

	"config.logger rotate":             "Log Rotate Settings",
	"config.logger rotate max size":    "Max Size（m）",
	"config.logger rotate max backups": "Max Buckups",
	"config.logger rotate max age":     "Max Age（day）",
	"config.logger rotate compress":    "Compress",

	"config.info log path":         "Info Log File Path",
	"config.error log path":        "Error Log File Path",
	"config.access log path":       "Access Log File Path",
	"config.info log off":          "Info Log Off",
	"config.error log off":         "Error Log Off",
	"config.access log off":        "Access Log Off",
	"config.access assets log off": "Access Assets Log Off",
	"config.sql log on":            "Open SQL Log",
	"config.log level":             "Level",

	"config.logger rotate encoder":                "Log Encoder Settings",
	"config.logger rotate encoder time key":       "Time Key",
	"config.logger rotate encoder level key":      "Level Key",
	"config.logger rotate encoder name key":       "Name Key",
	"config.logger rotate encoder caller key":     "Caller Key",
	"config.logger rotate encoder message key":    "Message Key",
	"config.logger rotate encoder stacktrace key": "Stacktrace Key",
	"config.logger rotate encoder level":          "Level Encoder",
	"config.logger rotate encoder time":           "Time Encoder",
	"config.logger rotate encoder duration":       "Duration Encoder",
	"config.logger rotate encoder caller":         "Caller Encoder",
	"config.logger rotate encoder encoding":       "Output Format",

	"config.capital":        "Capital",
	"config.capitalcolor":   "Capital with color",
	"config.lowercase":      "Lowercase",
	"config.lowercasecolor": "Lowercase with color",

	"config.seconds":     "Seconds",
	"config.nanosecond":  "Nanosecond",
	"config.microsecond": "Microsecond",
	"config.millisecond": "Millisecond",

	"config.full path":  "Full path",
	"config.short path": "Short path",

	"config.do not modify when you have not set up all assets": "Do not modify when you have not set up all assets",
	"config.it will work when theme is adminlte":               "It will work when theme is adminlte",

	"config.language." + CN:                  "Chinese",
	"config.language." + EN:                  "English",
	"config.language." + JP:                  "Japanese",
	"config.language." + strings.ToLower(TC): "Traditional Chinese",

	"system.system info":     "System And Application Info",
	"system.application":     "Application Info",
	"system.application run": "Applications Running Info",
	"system.system":          "System Info",

	"system.server_uptime":                        "Server Uptime",
	"system.current_goroutine":                    "Current Goroutines",
	"system.current_memory_usage":                 "Current Memory Usage",
	"system.total_memory_allocated":               "Total Memory Allocated",
	"system.memory_obtained":                      "Memory Obtained",
	"system.pointer_lookup_times":                 "Pointer Lookup Times",
	"system.memory_allocate_times":                "Memory Allocate Times",
	"system.memory_free_times":                    "Memory Free Times",
	"system.current_heap_usage":                   "Current Heap Usage",
	"system.heap_memory_obtained":                 "Heap Memory Obtained",
	"system.heap_memory_idle":                     "Heap Memory Idle",
	"system.heap_memory_in_use":                   "Heap Memory In Use",
	"system.heap_memory_released":                 "Heap Memory Released",
	"system.heap_objects":                         "Heap Objects",
	"system.bootstrap_stack_usage":                "Bootstrap Stack Usage",
	"system.stack_memory_obtained":                "Stack Memory Obtained",
	"system.mspan_structures_usage":               "MSpan Structures Usage",
	"system.mspan_structures_obtained":            "MSpan Structures Obtained",
	"system.mcache_structures_usage":              "MCache Structures Usage",
	"system.mcache_structures_obtained":           "MCache Structures Obtained",
	"system.profiling_bucket_hash_table_obtained": "Profiling Bucket Hash Table Obtained",
	"system.gc_metadata_obtained":                 "GC Metadata Obtained",
	"system.other_system_allocation_obtained":     "Other System Allocation Obtained",
	"system.next_gc_recycle":                      "Next GC Recycle",
	"system.last_gc_time":                         "Since Last GC Time",
	"system.total_gc_time":                        "Total GC Pause",
	"system.total_gc_pause":                       "Total GC Pause",
	"system.last_gc_pause":                        "Last GC Pause",
	"system.gc_times":                             "GC Times",

	"system.cpu_logical_core": "CPU Logical Core",
	"system.cpu_core":         "CPU Physical Core",
	"system.os_platform":      "OS Platform",
	"system.os_family":        "OS Family",
	"system.os_version":       "OS Version",
	"system.load1":            "Load1",
	"system.load5":            "Load5",
	"system.load15":           "Load15",
	"system.mem_total":        "Total Memory",
	"system.mem_available":    "Available Memory",
	"system.mem_used":         "Used Memory",

	"system.app_name":         "App Name",
	"system.go_admin_version": "App Version",
	"system.theme_name":       "Theme",
	"system.theme_version":    "Theme Version",
}
