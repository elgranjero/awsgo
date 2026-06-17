package transfer

// StartDirectoryListing is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Retrieves a list of the contents of a directory from a remote SFTP server. You
// specify the connector ID, the output path, and the remote directory path. You
// can also specify the optional MaxItems value to control the maximum number of
// items that are listed from the remote directory. This API returns a list of all
// files and directories in the remote directory (up to the maximum value), but
// does not return files or folders in sub-directories. That is, it only returns a
// list of files and directories one-level deep.
//
// After you receive the listing file, you can provide the files that you want to
// transfer to the RetrieveFilePaths parameter of the StartFileTransfer API call.
//
// The naming convention for the output file is  connector-ID-listing-ID.json . The
// output file contains the following information:
//
// - filePath : the complete path of a remote file, relative to the directory of
// the listing request for your SFTP connector on the remote server.
//
// - modifiedTimestamp : the last time the file was modified, in UTC time format.
// This field is optional. If the remote file attributes don't contain a timestamp,
// it is omitted from the file listing.
//
// - size : the size of the file, in bytes. This field is optional. If the remote
// file attributes don't contain a file size, it is omitted from the file listing.
//
// - path : the complete path of a remote directory, relative to the directory of
// the listing request for your SFTP connector on the remote server.
//
// - truncated : a flag indicating whether the list output contains all of the
// items contained in the remote directory or not. If your Truncated output value
// is true, you can increase the value provided in the optional max-items input
// attribute to be able to list more items (up to the maximum allowed list size of
// 10,000 items).
