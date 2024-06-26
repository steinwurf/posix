package posix

import "syscall"

//goland:noinspection GoSnakeCaseUsage
const (
	MADV_HUGEPAGE    = syscall.MADV_HUGEPAGE
	MADV_HWPOISON    = syscall.MADV_HWPOISON
	MADV_MERGEABLE   = syscall.MADV_MERGEABLE
	MADV_NOHUGEPAGE  = syscall.MADV_NOHUGEPAGE
	MADV_DOFORK      = syscall.MADV_DOFORK
	MADV_DONTFORK    = syscall.MADV_DONTFORK
	MADV_REMOVE      = syscall.MADV_REMOVE
	MADV_UNMERGEABLE = syscall.MADV_UNMERGEABLE
	MAP_DENYWRITE    = syscall.MAP_DENYWRITE
	MAP_EXECUTABLE   = syscall.MAP_EXECUTABLE
	MAP_GROWSDOWN    = syscall.MAP_GROWSDOWN
	MAP_HUGETLB      = syscall.MAP_HUGETLB
	MAP_LOCKED       = syscall.MAP_LOCKED
	MAP_NONBLOCK     = syscall.MAP_NONBLOCK
	MAP_POPULATE     = syscall.MAP_POPULATE
	MAP_STACK        = syscall.MAP_STACK
	MAP_TYPE         = syscall.MAP_TYPE
	MS_ACTIVE        = syscall.MS_ACTIVE
	MS_I_VERSION     = syscall.MS_I_VERSION
	MS_KERNMOUNT     = syscall.MS_KERNMOUNT
	MS_MANDLOCK      = syscall.MS_MANDLOCK
	MS_MGC_MSK       = syscall.MS_MGC_MSK
	MS_MGC_VAL       = syscall.MS_MGC_VAL
	MS_MOVE          = syscall.MS_MOVE
	MS_NOATIME       = syscall.MS_NOATIME
	MS_NODEV         = syscall.MS_NODEV
	MS_NODIRATIME    = syscall.MS_NODIRATIME
	MS_NOEXEC        = syscall.MS_NOEXEC
	MS_NOSUID        = syscall.MS_NOSUID
	MS_NOUSER        = syscall.MS_NOUSER
	MS_POSIXACL      = syscall.MS_POSIXACL
	MS_PRIVATE       = syscall.MS_PRIVATE
	MS_RDONLY        = syscall.MS_RDONLY
	MS_REC           = syscall.MS_REC
	MS_RELATIME      = syscall.MS_RELATIME
	MS_REMOUNT       = syscall.MS_REMOUNT
	MS_RMT_MASK      = syscall.MS_RMT_MASK
	MS_SHARED        = syscall.MS_SHARED
	MS_SILENT        = syscall.MS_SILENT
	MS_SLAVE         = syscall.MS_SLAVE
	MS_STRICTATIME   = syscall.MS_STRICTATIME
	MS_SYNCHRONOUS   = syscall.MS_SYNCHRONOUS
	MS_UNBINDABLE    = syscall.MS_UNBINDABLE
	MS_BIND          = syscall.MS_BIND
	MS_DIRSYNC       = syscall.MS_DIRSYNC
)
