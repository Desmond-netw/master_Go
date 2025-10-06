## CREATING DIRECTORY IN GO

In Go, use the os.Mkdir() method to create a single directory. Use 
os.MkdirAll() to establish a folder hierarchy (nested directories). Both 
methods need a path and the folder’s permission bits as parameters.

- This program attempts to create a directory in the current working directory.
- os.ModePerm " is a function that set the dirc permission mode to 0777 for all users, group and others.
- Note: The os.Mkdir() make you create a directory without a subdirectory.
- The os.MkdirAll() function make your create nested directories.
