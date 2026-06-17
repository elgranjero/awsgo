package rds

// PromoteReadReplica is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Promotes a read replica DB instance to a standalone DB instance.
//
// - Backup duration is a function of the amount of changes to the database
// since the previous backup. If you plan to promote a read replica to a standalone
// instance, we recommend that you enable backups and complete at least one backup
// prior to promotion. In addition, a read replica cannot be promoted to a
// standalone instance when it is in the backing-up status. If you have enabled
// backups on your read replica, configure the automated backup window so that
// daily backups do not interfere with read replica promotion.
//
// - This command doesn't apply to Aurora MySQL, Aurora PostgreSQL, or RDS
// Custom.
